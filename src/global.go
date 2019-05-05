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

func addColumn(collectionid, columnid int, name string, mtypeid int) {
	columns = append(columns, column{columnid, name, mtypeid, collectionid})
}

func addValue(rowid, columnid int, value string) {
	values = append(values, data{rowid, columnid, value})
}

// Init - Initializace basic stuff..
func Init() {

	// Add Types..
	addMType(1, "serial")
	addMType(2, "date")
	addMType(3, "decimal")
	addMType(4, "string")
	addMType(5, "percent")

	//
	createEvidenceForHmotnost()
	createEvidenceForKalendar()
}
