package models

// Dummy is the base model for dummy records.
type Dummy struct {
	BaseTableModel
	Name string `pg:"name" json:"name"`
}

// Columns returns the columns used by the model
func (Dummy) Columns() []string {
	return []string{"guid", "name"}
}
