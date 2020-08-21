package routes

import (
	"database/sql"
	"innovateafitc/config"
	"innovateafitc/models"
	"innovateafitc/modelsx"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type Routes struct {
	cfg    *config.Config
	db     *sql.DB
	Router *gin.Engine
}

func New(cfg *config.Config, db *sql.DB) (*Routes, error) {
	r := &Routes{
		cfg: cfg,
		db:  db,
	}

	router := gin.Default()

	router.POST("/report", r.PostReport)
	router.POST("/add-sensor", r.AddUpdateSensor)
	router.POST("/add-tracked-item", r.AddUpdateTrackedItem)
	router.POST("/add-item-type", r.AddUpdateDeviceType)

	r.Router = router

	return r, nil
}

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
