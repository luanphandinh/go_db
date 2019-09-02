package dbs

import "testing"

func TestSchemaValidate(t *testing.T) {
	tables := []Table{
		{
			Columns: []Column{
				{Name: "id", Type: INT, Primary: true, NotNull: true, AutoIncrement: true},
				{Name: "name", Type: TEXT, NotNull: true},
			},
		},
	}

	dbSchema := &Schema{
		Name: "workspace",
		Tables: tables,
	}

	assertHasErrorMessage(t, "table name should not empty", dbSchema.Validate())

	tables[0].Name = "user"
	assertNotHasError(t, dbSchema.Validate())
}
