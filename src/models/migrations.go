package models

// SchemaMigrations is the base model for schema_migrations
type SchemaMigrations struct {
	Version   int    `pg:"version,pk" json:"version"`
	Dirty     bool   `pg:"dirty" json:"dirty"`
	UpdatedAt string `pg:"updatedat" json:"updatedAt"`
}

// TableName defines the table name
func (schemaMigrations SchemaMigrations) TableName() string {
	return "schema_migrations"
}
