package krollparser

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/tushar2708/altcsv"
)

//MakeCSV will write the
func MakeCSV(records []Record) error {

	newFileName := fmt.Sprintf("ProcessedKrollFile-%s.csv", time.Now().Format("2006-01-02-3-04-05-PM"))
	csvfile, err := os.Create(newFileName)

	if err != nil {
		return err
	}

	csvwriter := altcsv.NewWriter(csvfile)
	csvwriter.AllQuotes = true

	headerRow := make([]string, 0)
	headerRow = append(headerRow, "DIN")
	headerRow = append(headerRow, "FillDate")
	headerRow = append(headerRow, "Qty")
	headerRow = append(headerRow, "Rx")

	err = csvwriter.Write(headerRow)
	if err != nil {
		return err
	}

	for s := 0; s < len(records); s++ {
		dataRow := make([]string, 0)
		dataRow = append(dataRow, strconv.Itoa(records[s].DIN))
		dataRow = append(dataRow, records[s].DispensingDate.Format("2006-01-02"))
		dataRow = append(dataRow, strconv.FormatFloat(records[s].Quantity, 'f', 1, 64))
		dataRow = append(dataRow, records[s].RxNum)
		err = csvwriter.Write(dataRow)
		if err != nil {
			return err
		}
	}
	csvwriter.Flush()

	csvfile.Close()

	return nil
}
