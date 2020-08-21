package routes

import (
	"database/sql"
	"innovateafitc/config"

	"github.com/gin-gonic/gin"
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

	router.GET("/get-logs", r.GetAllLogs)
	router.POST("/report", r.PostReport)
	router.POST("/add-sensor", r.AddUpdateSensor)
	router.POST("/add-tracked-item", r.AddUpdateTrackedItem)
	router.POST("/add-item-type", r.AddUpdateDeviceType)

	r.Router = router

	return r, nil
}
