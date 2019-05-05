package dmaster

import (
	"os"
	"bufio"
	"io"
	"log"
	"encoding/csv"
)

func createEvidenceForKalendar() {
	// Evidence..
	addCollection(2, "Kalendář")

	// Create columns..
	addColumn(2, 1, "#", 1)
	addColumn(2, 2, "rok", 1)
	addColumn(2, 3, "měsíc", 1)
	addColumn(2, 4, "den", 1)
	addColumn(2, 5, "wday", 1)		// week day
	addColumn(2, 6, "title", 4)		// Title of special day..
}

func addDataForEvidenceKalendar() {
  csvFile, _ := os.Open("../../data/kalendar.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
    line, error := reader.Read()
		if error == io.EOF {
      break
    }
    if error != nil {
      log.Fatal(error)
	  }
		addValue(1,1, line[0])
		addValue(1,2, line[1])
		addValue(1,3, line[2])
		addValue(1,4, line[3])
		addValue(1,5, line[4])
		addValue(1,6, line[5])
	}
}
