package models

type DataInfo struct {
	ApiBaseResp
	Data FetchData `json:"data"`
}

type FetchData struct {
	Devices []Device `json:"devices"`
	Positions []Position `json:"position"`
}

type Device struct {
	DeviceType string `json:"deviceType"`
	ControlCode string `json:"controlCode"`
	TypeId string `json:"typeId"`
	Id string `json:"id"`
	DeviceName string `json:"deviceName"`
}

type Position struct {
	PositionName string `json:"positionName"`
	TypeId string `json:"typeId"`
	DeviceLight []Light `json:deviceLight`
	DeviceCurtain []Curtain `json:deviceCurtain`
	DeviceAir []DeviceAir `json:deviceAir`
}

type Light struct {
	StopCode string `json:"stopCode"`
	DeviceId string `json:"deviceId"`
	ControlCode string `json:"controlCode"`
	OpenCode string `json:"openCode"`
	CloseCode string `json:"closeCode"`
	AdjustCode string `json:"adjustCode"`
	Name string `json:"name"`
	PositonId string `json:"positionId"`
}

type Curtain struct {
	StopCode string `json:"stopCode"`
	ControlCode string `json:"controlCode"`
	OpenCode string `json:"openCode"`
	CloseCode string `json:"closeCode"`
	PositionId string `json:"positionId"`
	Name string `json:"name"`
	DeviceId string `json:"deviceId"`

}

type DeviceAir struct {
	AirEconomize []AirEconomize `json:"airEconomize"`
	AirWind []AirWind `json:"airWind"`
	AirMode []AirMode `json:"airMode"`
	Id string `json:"id"`
	DeviceId string `json:"deviceId"`
	ControlCode string `json:"controlCode"`
	OpenCode string `json:"openCode"`
	CloseCode string `json:"closeCode"`
	PositionId string `json:"positionId"`
	Name string `json:"name"`
	AdjustCode string `json:"adjustCode"`
}

type AirEconomize struct {

	ControlCode string `json:"controlCode"`
	Code string `json:"code"`
	Id string `json:"id"`
	AirId string `json:"airId"`
	EconomizeName string `json:"economizeName"`
	Status string `json:"status"`
}

type AirWind struct {
	Status string `json:"status"`
	//"controlCode": "53"
	//"code": "00"
	//"createTime": "2017-11-10 15:56:10"
	//"windName": "低速"
	//"modeIcon": "icon_speed_slow_normal"
	//"id": "1"
	//"modeIconSelected": "icon_speed_slow_selected"
	//"airId": "1"
	//"status": "1"
}

type AirMode struct {

}