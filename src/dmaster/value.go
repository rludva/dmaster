package dmaster

import "fmt"

type value struct {
	collectionid int
	columnid     int
	content      string
}

var values []value

func AddValue(collectionid, columnid int, content string) {
	values = append(values, value{collectionid, columnid, content})
}

func PrintValues(collectionid int) {
	fmt.Printf("%v", values)

	for _, v := range values {
		if v.collectionid != collectionid {
			//continue
		}

		numberOfColumns := NumberOfColumns(collectionid)
		for c := 1; c < numberOfColumns; c++ {
			fmt.Printf("column_%v = \"\"", c)
		}
	}

}
