package dmaster

type collection struct {
	collectionid    int
	name            string
	numberOfColumns int
}

var collections []collection

func AddCollection(collectionid int, name string) {
	collections = append(collections, collection{collectionid: collectionid, name: name})
}

func NumberOfColumns(collectionid int) int {
	return 2
}
