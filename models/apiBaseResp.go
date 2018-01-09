package models

import "encoding/json"

type ApiBaseResp struct {
	Msg string `json:"msg"`
	Status json.Number `json:"status"`
	Data interface{} `json:"data"`
}
