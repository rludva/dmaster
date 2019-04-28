package dmaster

var (
	columns     []column
	collections []collection
	values      []data
	mtypes      []mtype
)

func addMType(typeid int, name string) {
	mtypes = append(mtypes, newMType(typeid, name))
}

func addCollection(collectionid int, name string) {
	collections = append(collections, collection{collectionid: collectionid, name: name})
}

func addColumn(columnid int, name string, mtypeid int, collectionid int) {
	columns = append(columns, column{columnid, name, mtypeid, collectionid})
}

// Init - Initializace basic stuff..
func Init() {

	// Add Types..
	addMType(1, "serial")
	addMType(2, "date")
	addMType(3, "decimal")
	addMType(4, "string")

	// Create collection of evidence..
	addCollection(1, "Osobní hmotnost")

	// Create evidence..
	addColumn(1, "#", 1, 1)
	addColumn(2, "datum", 2, 1)
	addColumn(3, "hmotnost", 3, 1)
	addColumn(4, "poznámka", 4, 1)
}
