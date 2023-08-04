package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tealeg/xlsx"

	// imported by me
	"io/fs"
	"path/filepath"
)

func usage() {
	fmt.Printf(`%s: -f=<Directory holding CSV> -d=<Delimiter> -o=<Output Directory>

`,
		os.Args[0])
}

func generateXLSXFromCSV(csvPath string, XLSXPath string, delimiter string) error {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	if len(delimiter) > 0 {
		reader.Comma = rune(delimiter[0])
	} else {
		reader.Comma = rune(';')
	}
	xlsxFile := xlsx.NewFile()
	sheet, err := xlsxFile.AddSheet(csvPath)
	if err != nil {
		return err
	}
	fields, err := reader.Read()
	for err == nil {
		row := sheet.AddRow()
		for _, field := range fields {
			cell := row.AddCell()
			cell.Value = field
		}
		fields, err = reader.Read()
	}
	if err != nil {
		fmt.Printf(err.Error())
	}
	return xlsxFile.Save(XLSXPath)
}

/*
listCSV

lists the paths of all csv files in a partcular directory
*/
func listAllCSV(directory *string) []string {
	var paths []string

	// goes through the directory
	// and will add paths of all csv files
	filepath.WalkDir(*directory, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ".csv" {
			paths = append(paths, s)
		}
		return nil
	})

	return paths
}

func main() {

	var cwd, _ = os.Getwd()

	var csvPath = flag.String("f", "", "Path to the CSV input file")
	var delimiter = flag.String("d", ",", "Delimiter for felds in the CSV input.")
	var outputPath = flag.String("o", cwd, "Output folder for xlsx files")

	flag.Parse()
	if len(os.Args) < 3 {
		usage()
		return
	}
	flag.Parse()

	var files []string = listAllCSV(csvPath) // all csv files

	for _, path := range files {

		// if the name of the file is longer than 31 char (which the library we're using has an issue with)
		if len(filepath.Base(path)) > 31 {
			var newPath = filepath.Join(*csvPath, strings.TrimSuffix(filepath.Base(path), ".csv")[len(path)-31:]+".csv")
			/*
				^^ this line renames them strangely because of the fact that the naming scheme would be identical
			*/
			err := os.Rename(path, newPath)
			if err != nil {
				fmt.Printf(err.Error())
				return
			}
			path = newPath
		}

		err := generateXLSXFromCSV(
			path,
			filepath.Join(*outputPath, strings.TrimSuffix(filepath.Base(path), ".csv")+".xlsx"),
			*delimiter,
		)

		if err != nil {
			fmt.Printf(err.Error())
			return
		}

	}

}
