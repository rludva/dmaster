package dmaster

type view struct {
	viewid int
	name   string
	dataid []int
}

func newView(viewid int, name string, dataid []int) *view {
	return &view{viewid, name, dataid}
}
