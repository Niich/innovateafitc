package routes

import (
	"innovateafitc/models"
	"innovateafitc/modelsx"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (r *Routes) PostReport(c *gin.Context) {
	var payload struct {
		SensorID string   `in:"id"`
		Value    null.Int `in:"value"`
	}

	err := modelsx.ParseToObj(c, &payload)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	println(payload.SensorID)

	mods := []qm.QueryMod{}

	mods = append(mods, models.SensorWhere.ID.EQ(payload.SensorID))

	dbSensor, err := models.Sensors(mods...).One(c.Request.Context(), r.db)

	newData := models.DataLog{
		TrackedDeviceID: dbSensor.TrackedDeviceID.Int64,
		SensorValue:     payload.Value,
	}

	err = newData.Insert(c.Request.Context(), r.db, boil.Infer())
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
}

func (r *Routes) GetAllLogs(c *gin.Context) {
	//

	logs, err := models.DataLogs(qm.Load(models.DataLogRels.TrackedDevice), qm.Load(qm.Rels(models.DataLogRels.TrackedDevice, models.TrackedDeviceRels.DeviceType))).All(c.Request.Context(), r.db)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	items := modelsx.DataLogFromModelBatch(logs...)

	items.Send(c)

}
