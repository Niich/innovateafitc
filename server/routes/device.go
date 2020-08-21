package routes

import (
	"innovateafitc/models"
	"innovateafitc/modelsx"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

func (r *Routes) AddUpdateDeviceType(c *gin.Context) {
	var payload struct {
		ID   int64  `in:"id"`
		Name string `in:"name"`
	}

	err := modelsx.ParseToObj(c, &payload)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	println(payload.ID)

	newData := models.DeviceType{
		ID:         payload.ID,
		DeviceName: payload.Name,
	}

	err = newData.Upsert(c.Request.Context(), r.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
}

func (r *Routes) AddUpdateTrackedItem(c *gin.Context) {
	var payload struct {
		ID           int64       `in:"id,omitempty"`
		DeviceTypeId int64       `in:"typeId"`
		DevLocation  null.String `in:"location"`
	}

	err := modelsx.ParseToObj(c, &payload)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	println(payload.ID)

	newData := models.TrackedDevice{
		ID:             payload.ID,
		DeviceLocation: payload.DevLocation,
		DeviceTypeID:   payload.DeviceTypeId,
	}

	err = newData.Upsert(c.Request.Context(), r.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
}
