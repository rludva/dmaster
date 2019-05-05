package dmaster

import "fmt"

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

func Print(evidenceid int) {
	for _,v := range collections {
		if v.collectionid == evidenceid {
			mcolumns := []int{}
			for _,vv := range columns {
				if vv.collectionid == evidenceid {
					mcolumns = append(mcolumns, vv.columnid)
				}
			}

			data := []string{}
			for _,vv := range values {
				if vv.columnid == mcolumns[0] {
					data = append(data, vv.value)
				}
			}

			fmt.Printf("%v, %v", v, mcolumns)
			fmt.Printf("%v", data)
		}
	}
}
