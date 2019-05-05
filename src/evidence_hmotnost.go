package dmaster

func createEvidenceForHmotnost()  {
	// Evidence..
	addCollection(1, "Osobní hmotnost")

	// Create columns..
	addColumn(1, 1, "#", 1)
	addColumn(1, 2, "datum", 2)
	addColumn(1, 3, "hmotnost", 3)
	addColumn(1, 4, "tuky", 5)
	addColumn(1, 5, "kosti", 5)
	addColumn(1, 6, "voda", 5)
	addColumn(1, 7, "svaly", 5)
	addColumn(1, 8, "poznámka", 4)
}

func addDataForEvidenceHmotnost() {
	addValue(1,1, "1")
	addValue(1,2, "2019-05-04")
	addValue(1,3, "112,4")
	addValue(1,4, "53,4")
	addValue(1,5, "2,4")
	addValue(1,6, "34,3")
	addValue(1,7, "38,0")
}
