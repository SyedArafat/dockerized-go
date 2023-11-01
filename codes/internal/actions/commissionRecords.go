package actions

import (
	"export_nrt_report/internal/model"
	"export_nrt_report/internal/repository"
	"fmt"
)

var nrtRecords []model.NrtCommission

func GetCommissionRecords() {
	nrtRecords = repository.GetEligibleRecords()
	fmt.Println(nrtRecords)
}
