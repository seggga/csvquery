package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/seggga/csvquery/parse"
	"github.com/seggga/csvquery/rpn"
)

func scanCSV(lm *parse.LexMachine, errorChan chan error, finishChan chan struct{}, ctx context.Context) {

	defer func() {
		close(errorChan)
		close(finishChan)
	}()

	lm.Where = rpn.ConvertToRPN(lm.Where)

	loopFiles := true   // a flag to indicate that the loop on files in FROM statement should be stopped
	printHeader := true // a flag to show the need to print the table's header

	//	scan csv-files
	for _, fileNameLex := range lm.From {

		if !loopFiles {
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
		reader := csv.NewReader(file)     // Считываем файл с помощью библиотеки encoding/csv
		tableHeader, err := reader.Read() //  Считываем шапку таблицы
		if err != nil {
			errorChan <- fmt.Errorf("cannot read file %s: %v", fileName, err)
			return
		}

		// compare columns sets from the query and the file
		err = parse.CheckCols(tableHeader, lm)
		if err != nil {
			errorChan <- fmt.Errorf("query does not fit the file data: %v", err)
			return
		}

		// print header if has not been printed yet
		if printHeader {
			parse.PrintHeader(lm)
			printHeader = false
		}

		loop := true
		for loop {

			select {
			case <-ctx.Done():
				loopFiles = false
				loop = false
			default:

				row, err := reader.Read()
				if err == io.EOF {
					break
				}

				if err != nil {
					errorChan <- fmt.Errorf("error reading csv-file %s: %v", fileName, err)
					loopFiles = false
					break
				}
				// compose a map with data of the current row: map[column_name]column_value
				valuesMap := parse.FillTheMap(tableHeader, row, lm)
				// create a slice based on the conditions in WHERE-statement
				lexSlice := rpn.InsertValues(valuesMap, lm.Where)

				result, err := rpn.CalculateRPN(lexSlice)
				if err != nil {
					errorChan <- fmt.Errorf("%v", err)
					loopFiles = false
					break
				}

				if result {
					parse.PrintLine(valuesMap, lm)
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
