package database

import (
	"github.com/leliuga/data/constants"
	"github.com/leliuga/validation"
)

func NewSchema(name, description string, tables ...*Table) *Schema {
	schema := &Schema{
		Name:        name,
		Description: description,
		Encoding:    constants.DefaultCharset,
		NamingStrategy: NamingStrategy{
			SchemaNameLength: 64,
			TableNameLength:  64,
			ColumnNameLength: 64,
		},
	}

	for _, table := range tables {
		schema.CreateTable(table)
	}

	return schema
}

// Validate makes Schema validatable by implementing [validation.Validatable] interface.
func (s *Schema) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Name, validation.Required, validation.Length(1, 63), validation.Match(constants.NameRegex).Error(constants.InvalidName)),
		validation.Field(&s.Tables, validation.Required.Error("At least one table must be defined.")),
	)
}

// CreateTable add a single table to the schema.
func (s *Schema) CreateTable(table *Table) {
	if index := s.TablePosition(s.Name); index > -1 {
		return
	}

	s.Tables = append(s.Tables, table)
}

// DropTable delete a single table by its name.
func (s *Schema) DropTable(name string) {
	if index := s.TablePosition(name); index > -1 {
		s.Tables = append(s.Tables[:index], s.Tables[index+1:]...)
	}
}

// Table returns a single table by its name.
func (s *Schema) Table(name string) *Table {
	if index := s.TablePosition(name); index > -1 {
		return s.Tables[index]
	}

	return nil
}

// TablePosition returns the index of a table by its name.
func (s *Schema) TablePosition(name string) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for index, f := range s.Tables {
		if f.Name == name {
			return index
		}
	}

	return -1
}
