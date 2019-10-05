package dbs

// ForeignKey of table
type ForeignKey struct {
	referer   string
	reference string
}

// Table defined db table structure
type Table struct {
	name        string
	primaryKey  []string
	columns     []*Column
	checks      []string
	comment     string
	foreignKeys []*ForeignKey
}

// WithName set name for table
func (table *Table) WithName(name string) *Table {
	table.name = name

	return table
}

// WithComment set comment for table
func (table *Table) WithComment(comment string) *Table {
	table.comment = comment

	return table
}

// AddPrimaryKey defined primary of table
func (table *Table) AddPrimaryKey(key []string) *Table {
	table.primaryKey = key

	return table
}

// AddColumn add defined column into table
func (table *Table) AddColumn(col *Column) *Table {
	table.columns = append(table.columns, col)

	return table
}

// AddColumns add a list of defined columns into table
func (table *Table) AddColumns(cols []*Column) *Table {
	table.columns = append(table.columns, cols...)

	return table
}

// AddCheck to table
// eg: table.AddCheck("age > 10")
func (table *Table) AddCheck(check string) *Table {
	table.checks = append(table.checks, check)

	return table
}

// AddChecks to table
// eg: table.AddCheck([]string{"age > 10", "age < 100"})
func (table *Table) AddChecks(checks []string) *Table {
	table.checks = append(table.checks, checks...)

	return table
}

// AddForeignKey create a ForeignKey object and add to table
func (table *Table) AddForeignKey(referer string, reference string) *Table {
	table.foreignKeys = append(table.foreignKeys, &ForeignKey{referer: referer, reference: reference})

	return table
}
