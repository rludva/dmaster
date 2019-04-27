package dmaster

type row struct {
	rowid int
	name string
	mtypeid int
	collectionid int
}

func newRow(rowid int, name string, mtypeid int, collectionid int) *row {
	return &row{rowid, name, mtypeid, collectionid}
}

