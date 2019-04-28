package dmaster

type column struct {
	columnid     int
	name         string
	mtypeid      int
	collectionid int
}

func newRow(columnid int, name string, mtypeid int, collectionid int) column {
	return column{columnid, name, mtypeid, collectionid}
}
