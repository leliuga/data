package database

import (
	"sync"

	"github.com/leliuga/data"
	"github.com/leliuga/validation"
)

type (
	// Schema defines a single schema structure.
	Schema struct {
		validation.Validatable `json:"-"`
		Name                   string         `json:"name"            yaml:"Name"`
		Description            string         `json:"description"     yaml:"Description"`
		Deprecated             string         `json:"deprecated"      yaml:"Deprecated"`
		Documentation          string         `json:"documentation"   yaml:"Documentation"`
		Encoding               string         `json:"encoding"        yaml:"Encoding"`
		Tables                 []*Table       `json:"tables"          yaml:"Tables"`
		NamingStrategy         NamingStrategy `json:"naming_strategy" yaml:"NamingStrategy"`
		mutex                  sync.Mutex     `json:"-"`
	}

	// Table defines a single Schema table structure.
	Table struct {
		validation.Validatable `json:"-"`
		Name                   string     `json:"name"        yaml:"Name"`
		Description            string     `json:"description" yaml:"Description"`
		Deprecated             string     `json:"deprecated"  yaml:"Deprecated"`
		Engine                 string     `json:"engine"      yaml:"Engine"`
		Codec                  string     `json:"codec"       yaml:"Codec"`
		Charset                string     `json:"charset"     yaml:"Charset"`
		ReadOnly               bool       `json:"read_only"   yaml:"ReadOnly"`
		Columns                []*Column  `json:"columns"     yaml:"Columns"`
		mutex                  sync.Mutex `json:"-"`
	}

	// Column defines a single Table column structure.
	Column struct {
		validation.Validatable `json:"-"`
		Kind                   data.Kind `json:"kind"                yaml:"Kind"`
		Name                   string    `json:"name"                yaml:"Name"`
		Description            string    `json:"description"         yaml:"Description"`
		Deprecated             string    `json:"deprecated"          yaml:"Deprecated"`
		NativeKind             string    `json:"native_kind"         yaml:"NativeKind"`
		Length                 int       `json:"length"              yaml:"Length"`
		NumericPrecision       int       `json:"numeric_precision"   yaml:"NumericPrecision"`
		NumericScale           int       `json:"numeric_scale"       yaml:"NumericScale"`
		DateTimePrecision      int       `json:"date_time_precision" yaml:"DateTimePrecision"`
		Codec                  string    `json:"codec"               yaml:"Codec"`
		Charset                string    `json:"charset"             yaml:"Charset"`
		Default                string    `json:"default"             yaml:"Default"`
		Validation             string    `json:"validation"          yaml:"Validation"`
		Replacement            string    `json:"replacement"         yaml:"Replacement"`
		Sensitive              bool      `json:"sensitive"           yaml:"Sensitive"`
		AutoIncrement          bool      `json:"auto_increment"      yaml:"AutoIncrement"`
		Primary                bool      `json:"primary"             yaml:"Primary"`
		Index                  bool      `json:"index"               yaml:"Index"`
		Unique                 bool      `json:"unique"              yaml:"Unique"`
		Nullable               bool      `json:"nullable"            yaml:"Nullable"`
		Creatable              bool      `json:"creatable"           yaml:"Creatable"`
		Updatable              bool      `json:"updatable"           yaml:"Updatable"`
		Readable               bool      `json:"readable"            yaml:"Readable"`
	}

	// NamingStrategy naming strategy
	NamingStrategy struct {
		SchemaNameLength int        `json:"schema_name_length" yaml:"SchemaNameLength"`
		TableNameLength  int        `json:"table_name_length"  yaml:"TableNameLength"`
		ColumnNameLength int        `json:"column_name_length" yaml:"ColumnNameLength"`
		TablePrefix      string     `json:"table_prefix"       yaml:"TablePrefix"`
		ColumnPrefix     string     `json:"column_prefix"      yaml:"ColumnPrefix"`
		SingularTable    bool       `json:"singular_table"     yaml:"SingularTable"`
		NoLowerCase      bool       `json:"no_lower_case"      yaml:"NoLowerCase"`
		NameReplacer     IReplacing `json:"-"`
	}

	// IReplacing replacing interface like strings.Replacer
	IReplacing interface {
		Replace(string) string
	}
)
