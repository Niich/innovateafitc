package modelsx

import (
	"innovateafitc/models"

	"github.com/gin-gonic/gin"
)

// DeviceType is an object representing the database table.
type DeviceType struct {
	ID         int64  `in"id"     out:"id"`
	DeviceName string `in:"name"  out:"name"`
}

func (r *DeviceType) ToModel() *models.DeviceType {
	ret := models.DeviceType{} // create a blank item to store our info
	copyToMatchingFields(r, &ret)
	return &ret
}

func DeviceTypeFromModel(data *models.DeviceType) *DeviceType {
	ret := DeviceType{}
	copyToMatchingFields(data, &ret)
	return &ret
}

type DeviceTypeSlice []DeviceType

func DeviceTypeFromModelBatch(data ...*models.DeviceType) DeviceTypeSlice {
	var res DeviceTypeSlice

	for _, m := range data {
		res = append(res, *DeviceTypeFromModel(m))
	}

	return res
}
func (r *DeviceType) Send(c *gin.Context) {
	SendModelsxObj(c, r)
}

func (r *DeviceTypeSlice) Send(c *gin.Context) {
	SendModelsxObj(c, r)
}
