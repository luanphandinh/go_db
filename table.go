package dbs

import "fmt"

type Table struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

func (table *Table) ValidateName() error {
	if table.Name == "" {
		return fmt.Errorf("table name should not empty")
	}

	return nil
}

func (table *Table) ValidateColumns() error  {
	for _, col := range table.Columns {
		if err := col.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (table *Table) Validate() error  {
	if err := table.ValidateName(); err != nil {
		return err
	}

	if err := table.ValidateColumns(); err != nil {
		return err
	}

	return nil
}

func (table *Table) ToString() (tableString string) {
	tableString = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", table.Name)
	cols := table.Columns
	for i := 0; i < len(cols); i++ {
		if i == 0 {
			tableString += fmt.Sprintf("%s", cols[i].ToString())
		} else {
			tableString += fmt.Sprintf(", %s", cols[i].ToString())
		}
	}

	return tableString + ")"
}
