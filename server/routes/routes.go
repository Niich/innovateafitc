package routes

import (
	"database/sql"
	"innovateafitc/config"
	"net/http"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

type Routes struct {
	db     *sql.DB
	Router *gin.Engine
}

func New(cfg *config.Config) (*Routes, error) {
	r := &Routes{}

	router := gin.Default()

	router.POST("/report", r.PostReport)

	r.Router = router

	return r, nil
}

func (r *Routes) PostReport(c *gin.Context) {
	data, err := ParseReport(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	println(data.IdMac, data.Value)
}

type Report struct {
	IdMac string `in:"id" out:"id"`
	Value int    `in:"value" out:"value"`
}

func ParseReport(c *gin.Context) (*Report, error) {
	data, err := c.GetRawData()

	if err != nil {
		return nil, err
	}

	api := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 string("in"),
	}.Froze()

	a := &Report{}

	if err := api.Unmarshal(data, a); err != nil {
		return nil, err
	}

	return a, nil
}
