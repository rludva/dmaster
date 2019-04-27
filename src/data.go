package dmaster

type data struct {
	dataid int
	value string
	rowid int
}

func newData(dataid int, value string, rowid int) *data {
	return &data{dataid, value, rowid}
}
