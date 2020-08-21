package routes

import (
	"innovateafitc/models"
	"innovateafitc/modelsx"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

func (r *Routes) AddUpdateSensor(c *gin.Context) {
	var payload struct {
		SensorID     string     `in:"id"`
		TrackedDevID null.Int64 `in:"trackedId"`
		BatStatus    null.Int   `in:"batteryLevel"`
	}

	err := modelsx.ParseToObj(c, &payload)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	println(payload.SensorID)

	newData := models.Sensor{
		ID:              payload.SensorID,
		BatStatus:       payload.BatStatus,
		TrackedDeviceID: payload.TrackedDevID,
	}

	err = newData.Upsert(c.Request.Context(), r.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
}
