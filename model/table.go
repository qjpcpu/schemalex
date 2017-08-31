package model

// NewTable create a new table with the given name
func NewTable(name string) Table {
	return &table{
		name: name,
	}
}

func (t *table) ID() string {
	return "table#" + t.name
}

func (t *table) LookupColumn(name string) (TableColumn, bool) {
	for col := range t.Columns() {
		if col.ID() == name {
			return col, true
		}
	}
	return nil, false
}

func (t *table) LookupIndex(id string) (Index, bool) {
	for idx := range t.Indexes() {
		if idx.ID() == id {
			return idx, true
		}
	}
	return nil, false
}

func (t *table) AddColumn(v TableColumn) {
	t.columns = append(t.columns, v)
}

func (t *table) AddIndex(v Index) {
	if v.IsPrimaryKey() {
		pcols := v.Columns()
		for pcol := range pcols {
			for i, col := range t.columns {
				if col.Name() == pcol {
					t.columns[i].SetPrimary(true)
				}
			}
		}
	}
	t.indexes = append(t.indexes, v)
}

func (t *table) AddOption(v TableOption) {
	t.options = append(t.options, v)
}

func (t *table) Name() string {
	return t.name
}

func (t *table) IsIfNotExists() bool {
	return t.ifnotexists
}

func (t *table) IsTemporary() bool {
	return t.temporary
}

func (t *table) SetIfNotExists(v bool) {
	t.ifnotexists = v
}

func (t *table) SetTemporary(v bool) {
	t.temporary = v
}

func (t *table) Columns() chan TableColumn {
	ch := make(chan TableColumn, len(t.columns))
	for _, col := range t.columns {
		ch <- col
	}
	close(ch)
	return ch
}

func (t *table) Indexes() chan Index {
	ch := make(chan Index, len(t.indexes))
	for _, idx := range t.indexes {
		ch <- idx
	}
	close(ch)
	return ch
}

func (t *table) Options() chan TableOption {
	ch := make(chan TableOption, len(t.options))
	for _, idx := range t.options {
		ch <- idx
	}
	close(ch)
	return ch
}

// NewTableOption creates a new table option with the given name, value, and a flag indicating if quoting is necessary
func NewTableOption(k, v string, q bool) TableOption {
	return &tableopt{
		key:        k,
		value:      v,
		needQuotes: q,
	}
}

func (t *tableopt) ID() string       { return "tableopt#" + t.key }
func (t *tableopt) Key() string      { return t.key }
func (t *tableopt) Value() string    { return t.value }
func (t *tableopt) NeedQuotes() bool { return t.needQuotes }
