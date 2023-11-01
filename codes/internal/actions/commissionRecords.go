package actions

import (
	"export_nrt_report/internal/model"
	"export_nrt_report/internal/repository"
	"strconv"
	"time"
)

var nrtRecords []model.NrtCommission

func GenerateCSV() {
	nrtRecords = repository.GetEligibleRecords()
	data := formatRecordForCSV()
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + ".csv"
	generateAndStore(data, fileName)
}

func formatRecordForCSV() [][]string {
	var data [][]string
	for _, record := range nrtRecords {
		row := []string{strconv.Itoa(record.ID), record.Msisdn, record.InstallationTime}
		data = append(data, row)
	}

	return data
}
