package krollparser

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

//Record is a struct which holds one dispensing record
type Record struct {
	DIN            int
	Quantity       float64
	RxNum          string
	DispensingDate time.Time
}

// ReadKrollCSV reads the RxForDrugDoctorGroups report
// and loads all the data into a slice of structs
func ReadKrollCSV(file string) ([]Record, error) {
	var dinCol, qtyCol, dispDateCol, rxNumCol int = 0, 0, 0, 0
	var foundDin, foundQty, foundDispDate, foundRxNum, foundAllCols bool = false, false, false, false, false

	//Lets make our returnable slice
	records := make([]Record, 0)

	csvFile, err := os.Open(file)
	if err != nil {
		return records, err
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	//Now we iterate through the file and look at the records
	for {
		//Read the next line
		line, err := reader.Read()
		if err == io.EOF {
			//We've reached the end of the file
			break
		} else if err != nil {
			//Some other error occurred
			return records, err
		}

		din, dinErr := strconv.Atoi(line[dinCol])
		qty, qtyErr := strconv.ParseFloat(strings.Replace(line[qtyCol], ",", "", -1), 64) //If dispensed more than 1000, remove the comma from the dispensed qty
		dispDate, dispDateErr := dateparse.ParseAny(line[dispDateCol])

		if foundAllCols && dinErr == nil && qtyErr == nil && dispDateErr == nil {
			//Process Data Record

			//Now we add the record to records
			records = append(records, Record{
				DIN:            din,
				Quantity:       qty,
				DispensingDate: dispDate,
				RxNum:          line[rxNumCol],
			})
		} else {
			//Process Header Column

			//Iterate through columns to check headings
			for c := 0; c < int(len(line)); c++ {
				if line[c] == "Rx" {
					//This is the Rx # column
					rxNumCol = c
					foundRxNum = true
				} else if line[c] == "DIN" {
					//We have found the column for DIN
					dinCol = c
					foundDin = true
				} else if line[c] == "FillDate" {
					// We have found the dispensing/fill date
					dispDateCol = c
					foundDispDate = true
				} else if line[c] == "Qty" {
					// We found the quantity column
					qtyCol = c
					foundQty = true
				}
			}

			if foundRxNum && foundDin && foundDispDate && foundQty {
				foundAllCols = true
			}
		}
	}

	return records, nil
}
