package dmaster

type data struct {
	dataid int
	columnid  int
	value  string
}

func newData(dataid int, columnid int, value string) *data {
	return &data{dataid, columnid, value}
}
