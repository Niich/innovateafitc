package modelsx

import (
	"innovateafitc/models"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
)

// Sensor is an object representing the database table.
type Sensor struct {
	ID              string     `boil:"id" json:"id" toml:"id" yaml:"id"`
	TrackedDeviceID null.Int64 `boil:"tracked_device_id" json:"tracked_device_id,omitempty" toml:"tracked_device_id" yaml:"tracked_device_id,omitempty"`
	BatStatus       null.Int   `boil:"bat_status" json:"bat_status,omitempty" toml:"bat_status" yaml:"bat_status,omitempty"`
}

func (r *Sensor) ToModel() *models.Sensor {
	ret := models.Sensor{} // create a blank item to store our info
	copyToMatchingFields(r, &ret)
	return &ret
}

func SensorFromModel(data *models.Sensor) *Sensor {
	ret := Sensor{}
	copyToMatchingFields(data, &ret)
	return &ret
}

type SensorSlice []Sensor

func SensorFromModelBatch(data ...*models.Sensor) SensorSlice {
	var res SensorSlice

	for _, m := range data {
		res = append(res, *SensorFromModel(m))
	}

	return res
}

func (r *Sensor) Send(c *gin.Context) {
	SendModelsxObj(c, r)
}

func (r *SensorSlice) Send(c *gin.Context) {
	SendModelsxObj(c, r)
}
