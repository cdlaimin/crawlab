package notification

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/apex/log"
	"github.com/crawlab-team/crawlab/core/models/models/v2"
	"github.com/crawlab-team/crawlab/trace"
	"gopkg.in/gomail.v2"
	"net/mail"
	"regexp"
	"strings"
)

func SendMail(s *models.NotificationSettingV2, ch *models.NotificationChannelV2, to, cc, bcc []string, title, content string) error {
	// compatibility for different providers
	var auth *XOAuth2Auth
	if ch.Provider == ChannelMailProviderOutlook {
		token, err := getOutlookToken(ch.OutlookTenantId, ch.OutlookClientId, ch.OutlookClientSecret)
		if err != nil {
			log.Errorf("failed to get outlook token: %v", err)
			return err
		}
		auth = &XOAuth2Auth{
			Username: ch.SMTPUsername,
			Token:    token,
		}
	}

	// config
	smtpConfig := smtpAuthentication{
		Server:         ch.SMTPServer,
		Port:           ch.SMTPPort,
		SenderIdentity: s.SenderName,
		SenderEmail:    s.SenderEmail,
		SMTPUser:       ch.SMTPUsername,
		SMTPPassword:   ch.SMTPPassword,
	}

	options := sendOptions{
		Subject: title,
		To:      to,
		Cc:      cc,
		Bcc:     bcc,
	}

	// convert html to text
	text := content
	if isHtml(text) {
		text = convertHtmlToText(text)
	}

	// apply theme
	if isHtml(content) {
		content = GetTheme() + content
	}

	// send the email
	if err := sendMail(smtpConfig, options, content, text, auth); err != nil {
		log.Errorf("failed to send email: %v", err)
		trace.PrintError(err)
		return err
	}

	return nil
}

func isHtml(content string) bool {
	regex := regexp.MustCompile(`(?i)<\s*(html|head|body|div|span|p|a|img|table|tr|td|th|tbody|thead|tfoot|ul|ol|li|dl|dt|dd|form|input|textarea|button|select|option|optgroup|fieldset|legend|label|iframe|embed|object|param|video|audio|source|canvas|svg|math|style|link|script|meta|base|title|br|hr|b|strong|i|em|u|s|strike|del|ins|mark|small|sub|sup|big|pre|code|q|blockquote|abbr|address|bdo|cite|dfn|kbd|var|samp|ruby|rt|rp|time|progress|meter|output|area|map)`)
	return regex.MatchString(content)
}

func convertHtmlToText(content string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Errorf("failed to convert html to text: %v", err)
		trace.PrintError(err)
		return ""
	}
	return doc.Text()
}

type smtpAuthentication struct {
	Server         string
	Port           int
	SenderEmail    string
	SenderIdentity string
	SMTPUser       string
	SMTPPassword   string
}

// sendOptions are options for sending an email
type sendOptions struct {
	Subject string
	To      []string
	Cc      []string
	Bcc     []string
}

// send email
func sendMail(smtpConfig smtpAuthentication, options sendOptions, htmlBody string, txtBody string, auth *XOAuth2Auth) error {
	if smtpConfig.Server == "" {
		return errors.New("SMTP server config is empty")
	}

	if smtpConfig.Port == 0 {
		return errors.New("SMTP port config is empty")
	}

	if smtpConfig.SMTPUser == "" {
		return errors.New("SMTP user is empty")
	}

	if len(options.To) == 0 {
		return errors.New("no receiver emails configured")
	}

	from := mail.Address{
		Name:    smtpConfig.SenderIdentity,
		Address: smtpConfig.SenderEmail,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from.String())
	m.SetHeader("To", options.To...)
	m.SetHeader("Subject", options.Subject)
	if len(options.Cc) > 0 {
		m.SetHeader("Cc", options.Cc...)
	}
	if len(options.Bcc) > 0 {
		m.SetHeader("Bcc", options.Bcc...)
	}

	m.SetBody("text/plain", txtBody)
	m.AddAlternative("text/html", htmlBody)

	d := gomail.NewDialer(smtpConfig.Server, smtpConfig.Port, smtpConfig.SMTPUser, smtpConfig.SMTPPassword)
	if auth != nil {
		d.Auth = auth
	}

	return d.DialAndSend(m)
}
