package dmaster

type mtype struct {
	mtypeid int
	name    string
}

func newMType(mtypeid int, name string) mtype {
	return mtype{mtypeid, name}
}
