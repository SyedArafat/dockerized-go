package repository

import "export_nrt_report/internal/model"

func GetEligibleRecords() []model.NrtCommission {
	var commissionRecords []model.NrtCommission
	response, err := db.Query("select * from nrt_commission_records")
	defer response.Close()
	for response.Next() {
		var commissionRecord model.NrtCommission
		err = response.Scan(&commissionRecord.ID, &commissionRecord.Msisdn, &commissionRecord.DeviceID,
			&commissionRecord.ActivationDate, &commissionRecord.InstallationTime, &commissionRecord.Shared,
			&commissionRecord.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		commissionRecords = append(commissionRecords, commissionRecord)
	}

	return commissionRecords
}
