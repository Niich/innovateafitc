package modelsx

import (
	"innovateafitc/models"

	"github.com/volatiletech/null"
)

// DataLog is an object representing the database table.
type DataLog struct {
	ID              int64     `in:"-"  out:"id"     `
	TrackedDeviceID int64     `in:"trackedDeviceId"  out:"trackedDeviceId"     `
	LogTime         null.Time `in:"-"  out:"time"     `
	SensorValue     null.Int  `in:"value"  out:"value"     `
}

func (r *DataLog) ToModel() *models.DataLog {
	ret := models.DataLog{} // create a blank item to store our info
	copyToMatchingFields(r, &ret)
	return &ret
}

func DataLogFromModel(data *models.DataLog) *DataLog {
	ret := DataLog{}
	copyToMatchingFields(data, &ret)
	return &ret
}

type DataLogSlice []DataLog

func DataLogFromModelBatch(data ...*models.DataLog) DataLogSlice {
	var res DataLogSlice

	for _, m := range data {
		res = append(res, *DataLogFromModel(m))
	}

	return res
}
