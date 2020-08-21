package modelsx

import (
	"innovateafitc/models"

	"github.com/volatiletech/null"
)

// TrackedDevice is an object representing the database table.
type TrackedDevice struct {
	ID             int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	DeviceTypeID   int64       `boil:"device_type_id" json:"device_type_id" toml:"device_type_id" yaml:"device_type_id"`
	DeviceLocation null.String `boil:"device_location" json:"device_location,omitempty" toml:"device_location" yaml:"device_location,omitempty"`
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
