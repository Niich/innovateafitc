package modelsx

import (
	"innovateafitc/models"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
)

// TrackedDevice is an object representing the database table.
type TrackedDevice struct {
	ID             int64       `in:"id" out:"id"`
	DeviceTypeID   int64       `in:"deviceTypeID" out:"deviceTypeID"`
	DeviceType     *DeviceType `in:"deviceType" out:"deviceType,omitempty"`
	DeviceLocation null.String `in:"location" out:"location"`
}

func (r *TrackedDevice) ToModel() *models.TrackedDevice {
	ret := models.TrackedDevice{} // create a blank item to store our info
	copyToMatchingFields(r, &ret)
	return &ret
}

func TrackedDeviceFromModel(data *models.TrackedDevice) *TrackedDevice {
	ret := TrackedDevice{}
	copyToMatchingFields(data, &ret)
	return &ret
}

type TrackedDeviceSlice []TrackedDevice

func TrackedDeviceFromModelBatch(data ...*models.TrackedDevice) TrackedDeviceSlice {
	var res TrackedDeviceSlice

	for _, m := range data {
		res = append(res, *TrackedDeviceFromModel(m))
	}

	return res
}

func (r *TrackedDevice) Send(c *gin.Context) {
	SendModelsxObj(c, r)
}

func (r *TrackedDeviceSlice) Send(c *gin.Context) {
	SendModelsxObj(c, r)
}
