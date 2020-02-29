package libtz

import (
	"log"
	"time"
)

func Auckland() *time.Location {
	loc, err := time.LoadLocation("Pacific/Auckland")
	if err != nil {
		log.Fatalf("load timezone#%s error, system breakdown", "Pacific/Auckland")
	}
	return loc
}

func Tokyo() *time.Location {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalf("load timezone#%s error, system breakdown", "Asia/Tokyo")
	}
	return loc
}
