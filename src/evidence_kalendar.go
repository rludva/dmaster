package dmaster

func createEvidenceForKalendar() {
	// Evidence..
	addCollection(2, "Kalendář")

	// Create columns..
	addColumn(2, 1, "#", 1)
	addColumn(2, 2, "rok", 1)
	addColumn(2, 3, "měsíc", 1)
	addColumn(2, 4, "den", 1)
	addColumn(2, 5, "dnazev", 4)	// Week name of day..
	addColumn(2, 6, "title", 4)		// Title of special day..
}

func addDataForEvidenceKalendar() {
  addValue(1,1, "1")
	addValue(1,2, "2019")
	addValue(1,3, "01")
	addValue(1,4, "01")
	addValue(1,5, "Úterý")
	addValue(1,6, "")
}
