package dmaster

type data struct {
	rowid int
	columnid  int
	value  string
}

func newData(rowid int, value string, columnid int) *data {
	return &data{rowid, value, columnid}
}
