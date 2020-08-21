package modelsx

import (
	"innovateafitc/models"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
)

// DataLog is an object representing the database table.
type DataLog struct {
	ID              int64          `in:"-"                out:"id"     `
	TrackedDeviceID int64          `in:"trackedDeviceId"  out:"-"     `
	LogTime         null.Time      `in:"-"                out:"time"     `
	SensorValue     null.Int       `in:"value"            out:"value"     `
	TrackedDevice   *TrackedDevice `in:"-"                out:"trackedDevice,omitempty"     `
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
		v := DataLogFromModel(m)
		if v == nil {
			println("nil responce: ", m.ID)
		} else {
			res = append(res, *v)
		}
	}

	return res
}

func (r *DataLog) Send(c *gin.Context) {
	SendModelsxObj(c, r)
}

func (r *DataLogSlice) Send(c *gin.Context) {
	SendModelsxObj(c, r)
}
