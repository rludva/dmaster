package dmaster

var (
	mtypes []mtype
)

func addMType(typeid int, name string) {
	mtypes = append(mtypes, newMType(typeid, name))
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
	for _, v := range collections {
		if v.collectionid == evidenceid {
			mcolumns := []int{}
			for _, vv := range columns {
				if vv.collectionid == evidenceid {
					mcolumns = append(mcolumns, vv.columnid)
				}
			}

			PrintValues(evidenceid)

			//
			//fmt.Printf("%v, %v", v, mcolumns)
		}
	}
}
