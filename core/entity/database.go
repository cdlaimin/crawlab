package entity

type DatabaseMetadata struct {
	Databases []Database `json:"databases"`
}

type Database struct {
	Name   string          `json:"name"`
	Tables []DatabaseTable `json:"tables"`
}

type DatabaseTable struct {
	Name    string           `json:"name"`
	Columns []DatabaseColumn `json:"columns"`
	Indexes []DatabaseIndex  `json:"indexes"`
}

type DatabaseColumn struct {
	Name          string           `json:"name"`
	Type          string           `json:"type"`
	PrimaryKey    bool             `json:"primary_key,omitempty"`
	NotNull       bool             `json:"not_null"`
	Key           string           `json:"key,omitempty"`
	Default       string           `json:"default"`
	Extra         string           `json:"extra,omitempty"`
	AutoIncrement bool             `json:"auto_increment,omitempty"`
	Children      []DatabaseColumn `json:"children,omitempty"`
	Hash          string           `json:"hash,omitempty"`
	OriginalName  string           `json:"original_name,omitempty"`
	Status        string           `json:"status,omitempty"`
}

type DatabaseIndex struct {
	Name         string                `json:"name"`
	Type         string                `json:"type,omitempty"`
	Columns      []DatabaseIndexColumn `json:"columns"`
	Unique       bool                  `json:"unique"`
	Hash         string                `json:"hash,omitempty"`
	OriginalName string                `json:"original_name,omitempty"`
	Status       string                `json:"status,omitempty"`
}

type DatabaseIndexColumn struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}

func (col *DatabaseIndexColumn) OrderString() string {
	if col.Order < 0 {
		return "DESC"
	} else {
		return "ASC"
	}
}
