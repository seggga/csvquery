package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/seggga/csvquery/parse"
)

func scanCSV(lm *parse.LexMachine, errorChan chan error, finishChan chan struct{}, ctx context.Context) {

	defer func() {
		close(errorChan)
		close(finishChan)
	}()

	stopCycle := false  // a flag to show that ctx-channel is closed
	printHeader := true // a flag to show the need to print the table's header

	//	scan csv-files
	for _, fileNameLex := range lm.From {

		if stopCycle {
			break
		}

		fileName := fileNameLex.Litera
		// check file existense
		if _, err := os.Stat(fileName); err != nil {
			errorChan <- fmt.Errorf("file %s was not found. %w", fileName, err)
			return
		}

		// file opening
		file, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
		if err != nil {
			errorChan <- fmt.Errorf("unable to read file %s. %w", fileName, err)
			return
		}

		// read the header of the csv-file
		reader := csv.NewReader(file)  // Считываем файл с помощью библиотеки encoding/csv
		fileCols, err := reader.Read() //  Считываем шапку таблицы
		if err != nil {
			errorChan <- fmt.Errorf("Cannot read file %s: %v", fileName, err)
			return
		}

		// compare columns sets from the query and the file
		err = parse.CheckCols(fileCols, lm)
		if err != nil {
			errorChan <- fmt.Errorf("query does not fit the file data: %v", err)
			return
		}

		if printHeader {
			parse.PrintTheLine(fileCols, lm)
			printHeader = false
		}

		for {

			select {
			case <-ctx.Done():
				stopCycle = true
				break
			default:

				row, err := reader.Read()
				if err == io.EOF {
					break
				}

				if err != nil {
					errorChan <- fmt.Errorf("Error reading csv-file %s: %v", fileName, err)
					stopCycle = true
					break
				}
				// compose a map holding data of the current row
				rowData := mylexer.FillTheMap(fileCols, row, lm)
				// create a slice based on the conditions in WHERE-statement
				lexSlice := mylexer.MakeSlice(rowData, lm)

				if mylexer.Execute(lexSlice) {
					_ = mylexer.PrintTheRow(rowData, lm)
				}
			}
		}

		// closing file
		if err := file.Close(); err != nil {
			errorChan <- fmt.Errorf("error while closing a file %s %v", fileName, err)
			return
		}

	}
}
