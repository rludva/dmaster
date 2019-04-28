package dmaster

type collection struct {
	collectionid int
	name         string
}

func newCollection(collectionid int, name string) *collection {
	return &collection{collectionid: collectionid, name: name}
}
