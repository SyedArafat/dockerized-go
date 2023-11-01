package actions

import (
	"encoding/csv"
	"log"
	"os"
)

func generateAndStore(data [][]string, fileName string) {
	err := os.MkdirAll(os.Getenv("CSV_EXPORT_DIRECTORY"), 0700)
	if err != nil {
		log.Printf("Directory creation failed %s", err)
		return
	}
	file, err := os.Create(os.Getenv("CSV_EXPORT_DIRECTORY") + fileName)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	fileSaveErr := w.WriteAll(data)
	if fileSaveErr != nil {
		return
	}
	defer w.Flush()
}
