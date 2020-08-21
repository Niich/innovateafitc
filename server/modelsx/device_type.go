package modelsx

import "innovateafitc/models"

// DeviceType is an object representing the database table.
type DeviceType struct {
	ID         int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	DeviceName string `boil:"device_name" json:"device_name" toml:"device_name" yaml:"device_name"`
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
