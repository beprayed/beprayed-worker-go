package model

import (
	"beprayed-worker-go/db"
	"database/sql"
	"fmt"
	"time"
)

type PrayRecord struct {
	// ID          int64  `db:"id"`
	PrayerID    string `db:"prayer_id"`
	Contributor string `db:"contributor"`
	Time        int64  `db:"time"`
}

type PrayRecordModel struct{}

const interval = "DAY"

func (m *PrayRecordModel) Insert(r PrayRecord) error {
	if m.IsOld(r) {
		_, err := db.GetDB().Exec("INSERT INTO public.pray_record (prayer_id, contributor, time) VALUES ($1, $2, $3)", r.PrayerID, r.Contributor, r.Time)
		if err != nil {
			return err
		}
		return nil
	} else {
		// we just ignore the insert since it's redundant
		return nil
	}
}

func (m *PrayRecordModel) IsOld(r PrayRecord) bool {
	last, err := lastRecord(r)

	if err == sql.ErrNoRows || r == (PrayRecord{}) {
		// if there is no record, then allow insert
		return true
	}

	// compare the time, return false if the date is the same

	lastPrayTime := time.Unix(last.Time, 0)
	thisPrayTime := time.Unix(r.Time, 0)

	fmt.Println("last pray time", lastPrayTime)
	fmt.Println("this pray time", thisPrayTime)

	if interval == "DAY" {
		return !(lastPrayTime.Year() == thisPrayTime.Year() && lastPrayTime.Month() == thisPrayTime.Month() && lastPrayTime.Day() == thisPrayTime.Day())
	}

	return false
}

func lastRecord(r PrayRecord) (PrayRecord, error) {
	var lastRecord PrayRecord
	err := db.GetDB().SelectOne(&lastRecord, "SELECT * FROM public.pray_record WHERE prayer_id = $1 AND contributor = $2 ORDER BY time DESC LIMIT 1", r.PrayerID, r.Contributor)
	if err != nil {
		return lastRecord, err
	}

	return lastRecord, nil
}
