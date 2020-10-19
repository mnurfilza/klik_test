package model

import "klik/controler"

type AdjustStock struct {
	ID         int    `json:"id"`
	Product    string `json:"product"`
	Adjustment int    `json:"adjustment"`
}

type AdjustResponse struct {
	Status string       `json:"status_code"`
	Result AdjustResult `json:"result"`
}

type AdjustResult struct {
	Status        string `json:"status"`
	LocationID    string `json:"location_id"`
	Product       string `json:"product"`
	Adjustment    int    `json:"adjustment"`
	StockQuantity int    `json:"stock_quantity"`
}

type ResponseStocks struct {
	Status        int                `json:"status"`
	StatusMessage string             `json:"status_message"`
	Stocks        []*controler.Stock `json:"stocks"`
}

type ResponseStockLog struct {
	StatusCode      int              `json:"status_code"`
	Status          string           `json:"status"`
	LocationID      int              `json:"location_id"`
	Location        string           `json:"location"`
	Product         string           `json:"product"`
	CurrentQuantity int              `json"current_qty"`
	Logs            []*controler.Log `json:"logs"`
}
