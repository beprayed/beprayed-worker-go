package test

import (
	"time"

	"beprayed-worker-go/db"
	. "beprayed-worker-go/model"

	"testing"

	"github.com/stretchr/testify/assert"
	uuid "github.com/twinj/uuid"
)

const duplicateError = "pq: duplicate key value violates unique constraint \"pray_record_prayer_id_contributor_time_key\""

func TestPrayRecord(t *testing.T) {
	var u = &TestUtil{}
	u.InitDB()
	u.ClearDataInPostgres()

	t.Run("should insert a new pray record", func(t *testing.T) {
		// generate a uuid
		prayerID := uuid.NewV4().String()
		contributorID := uuid.NewV4().String()

		record := PrayRecord{
			PrayerID:    prayerID,
			Contributor: contributorID,
			Time:        time.Now().Unix(),
		}

		m := PrayRecordModel{}
		err := m.Insert(record)

		if err != nil {
			t.Errorf("error: failed to insert pray record")
		}

		assert.Nil(t, err)

		// there should be 1 record in the database
		count, err := db.GetDB().SelectInt("SELECT COUNT(*) FROM public.pray_record")
		assert.Nil(t, err)
		assert.Equal(t, 1, int(count))
	})

	t.Run("should reject record when prayer id, contributor already exist within the past 24 hours if interval is set to 1 day", func(t *testing.T) {
		u.ClearDataInPostgres()
		prayerID := uuid.NewV4().String()
		contributorID := uuid.NewV4().String()
		time := time.Now().Unix()

		r := PrayRecord{
			PrayerID:    prayerID,
			Contributor: contributorID,
			Time:        time,
		}

		m := PrayRecordModel{}
		err := m.Insert(r)

		if err != nil {
			t.Errorf("error: failed to insert pray record")
		}

		assert.Nil(t, err)

		count, _ := db.GetDB().SelectInt("SELECT COUNT(*) FROM public.pray_record")
		assert.Equal(t, 1, int(count))

		err = m.Insert(r)
		assert.Nil(t, err)

		// should still have 1 record
		assert.Equal(t, 1, int(count))
	})

	t.Run("should accept record when prayer id, contributor already exist within the past 24 hours if interval is set to 1 day", func(t *testing.T) {
		u.ClearDataInPostgres()
		prayerID := uuid.NewV4().String()
		contributorID := uuid.NewV4().String()
		// yesterdays time
		time := time.Now().Unix() - 86400

		r := PrayRecord{
			PrayerID:    prayerID,
			Contributor: contributorID,
			Time:        time,
		}

		m := PrayRecordModel{}
		err := m.Insert(r)

		if err != nil {
			t.Errorf("error: failed to insert pray record")
		}

		assert.Nil(t, err)

		count, _ := db.GetDB().SelectInt("SELECT COUNT(*) FROM public.pray_record")
		assert.Equal(t, 1, int(count))

		// update time to today
		r.Time = time + 86401

		err = m.Insert(r)
		assert.Nil(t, err)

		count, _ = db.GetDB().SelectInt("SELECT COUNT(*) FROM public.pray_record")
		// should have 2 records
		assert.Equal(t, 2, int(count))
	})
}
