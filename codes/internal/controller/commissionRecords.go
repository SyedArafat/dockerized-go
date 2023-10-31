package controller

import (
	"fmt"
	"log"
)

type Settings struct {
	ID  int
	Key string
}

func GetCommissionRecords() {
	res, err := db.Query("select sid, settings_key from settings")
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {

		var settings Settings
		err := res.Scan(&settings.ID, &settings.Key)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", settings)
	}

}
