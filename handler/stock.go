package handler

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"klik/controler"
	"klik/model"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mnurfilza/lib"
)

var db *sql.DB

func RegisDB(sqlDB *sql.DB) {
	if db != nil {
		panic("db telah terdaftar")
	}
	db = sqlDB
}

func Adjustment(c *gin.Context) {
	adjust := &model.AdjustStock{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, adjust); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	adj := make(map[string]interface{})
	adj["adjustment"] = adjust.Adjustment

	stock := &controler.Stock{
		ID: adjust.ID,
	}
	_, err := stock.Update(db, adj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if err := stock.Get(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	result := &model.AdjustResult{}
	result.Adjustment = adjust.Adjustment
	result.LocationID = strconv.Itoa(adjust.ID)
	result.StockQuantity = stock.Quantity
	result.Product = stock.Product
	result.Status = "Success"
	response := &model.AdjustResponse{}
	response.Status = "200"
	response.Result = *result
	c.JSON(200, response)
}

func GetAllstock(c *gin.Context) {
	reqParams := lib.RequestParams{OrderBy: "id"}
	data, err := controler.Stocks(db, reqParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	response := &model.ResponseStocks{}
	response.Status = 200
	response.StatusMessage = "Success"
	response.Stocks = data
	c.JSON(200, response)
}

func GetLogStock(c *gin.Context) {
	params := GetParams(c)
	locID := LocationIDParams(c, params)
	id, _ := strconv.Atoi(locID)
	stock := &controler.Stock{ID: id}
	if err := stock.Get(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	p := []lib.Params{
		{Field: "locationID", Op: "=", Value: stock.ID},
	}

	rp := lib.RequestParams{Param: p}

	data, err := controler.Logs(db, rp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	response := &model.ResponseStockLog{}
	response.StatusCode = 200
	response.Status = "Success Logs Found"
	response.Product = stock.Product
	response.LocationID = stock.ID
	response.Location = stock.Location
	response.CurrentQuantity = stock.Quantity
	response.Logs = data
	if len(data) == 0 {
		response.Status = "There is no log found"

	}

	c.JSON(200, response)
}

func LocationIDParams(c *gin.Context, params url.Values) string {
	var locationID = ""
	if params.Get("location_id") != "" {
		locationID = params.Get("location_id")
	}
	return locationID
}

func GetParams(c *gin.Context) url.Values {
	return c.Request.URL.Query()
}
