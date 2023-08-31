package database

import (
	"github.com/leliuga/data/constants"
	"github.com/leliuga/validation"
)

func NewTable(name, description string, columns ...*Column) *Table {
	table := &Table{
		Name:        name,
		Description: description,
	}

	for _, column := range columns {
		table.CreateColumn(column)
	}

	return table
}

// Validate makes Table validatable by implementing [validation.Validatable] interface.
func (t *Table) Validate() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.Name, validation.Required, validation.Length(1, 63), validation.Match(constants.NameRegex).Error(constants.InvalidName)),
		validation.Field(&t.Columns, validation.Required.Error("At least one column must be defined.")),
	)
}

// CreateColumn add a single column to the table.
func (t *Table) CreateColumn(column *Column) {
	if index := t.ColumnPosition(t.Name); index > -1 {
		return
	}

	t.Columns = append(t.Columns, column)
}

// DropColumn delete a single column by its name.
func (t *Table) DropColumn(name string) {
	if index := t.ColumnPosition(name); index > -1 {
		t.Columns = append(t.Columns[:index], t.Columns[index+1:]...)
	}
}

// Column returns a single column by its name.
func (t *Table) Column(name string) *Column {
	if index := t.ColumnPosition(name); index > -1 {
		return t.Columns[index]
	}

	return nil
}

// ColumnPosition returns the index of a column by its name.
func (t *Table) ColumnPosition(name string) int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	for index, f := range t.Columns {
		if f.Name == name {
			return index
		}
	}

	return -1
}
