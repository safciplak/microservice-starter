package models

import "strings"

// BaseTableModel is the base model for all table models
type BaseTableModel struct {
	ID        uint   `pg:"id,pk" json:"-"`
	CreatedBy uint   `pg:"createdby" json:"-"`
	CreatedAt string `pg:"createdat" json:"-"`
	UpdatedBy uint   `pg:"updatedby" json:"-"`
	UpdatedAt string `pg:"updatedat" json:"-"`
	GUID      string `pg:"guid" json:"guid"`
	IsDeleted bool   `pg:"isdeleted" json:"-"`
}

// Returning makes sure only certain tables are returned and updated on the model
func (BaseTableModel) Returning() string {
	columns := []string{"id", "createdby", "createdat", "updatedby", "updatedat", "guid", "isdeleted"}

	return strings.Join(columns, ", ")
}
