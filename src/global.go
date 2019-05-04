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

	// Create collection of evidence..
	addCollection(1, "Osobní hmotnost")

	// Create evidence..
	addColumn(1, "#", 1, 1)
	addColumn(2, "datum", 2, 1)
	addColumn(3, "hmotnost", 3, 1)
	addColumn(4, "tuky", 5, 1)
	addColumn(5, "kosti", 5, 1)
	addColumn(6, "voda", 5, 1)
	addColumn(7, "svaly", 5, 1)
	addColumn(8, "poznámka", 4, 1)
}

func Data() {
	addValue(1,1, "1")
	addValue(1,2, "2019-05-04")
	addValue(1,3, "112,4")
	addValue(1,4, "53,4")
	addValue(1,5, "2,4")
	addValue(1,6, "34,3")
	addValue(1,7, "38,0")
}
