package api

import (
// "encoding/json"
// "net/http"
)

type CoinBalanceParams struct {
	Username string `json:"username"`
}

type CoinBalanceResponse struct {
	Code    int `json:"code"`
	Balance int `json:"balance"`
}
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
