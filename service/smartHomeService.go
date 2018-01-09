package service


import (
	"homebridge/models"
	"encoding/json"
	"net/http"
	"homebridge/config"
	"io/ioutil"
	"homebridge/httpUtil"
	"net/url"
	"strconv"
	"strings"
	"fmt"
)



func FetchDevicesInfo(typeId string ) ([]models.Position,[]models.Device,error) {

	url := fmt.Sprintf("%v?typeId=%v",config.Host+":"+config.Port+config.FetchDeviceInfo,typeId)
	body,err:= httpUtil.DoGet(url,"application/x-www-form-urlencoded")
	//responseData:= []byte(``)
	var resp models.DataInfo
	err = json.Unmarshal(body, &resp)
	return resp.Data.Positions,resp.Data.Devices,err
}


func TurnOnLight(light *models.Light,position *models.Position) ([]byte,error) {
	parms := url.Values{}
	parms.Set("code",light.OpenCode)
	parms.Add("controlCode",light.ControlCode)
	parms.Add("deviceId",light.DeviceId)
	parms.Add("deviceType",config.TestDeviceType)
	parms.Add("mac",config.TestMacAddress)
	parms.Add("positionId",light.PositonId)
	parms.Add("typeId",position.TypeId)
	parmsStr :=  parms.Encode()
	resp,err := http.Post(config.Host+config.ControlUrl,"application/x-www-form-urlencoded",strings.NewReader(parmsStr))
	defer resp.Body.Close()
	if err != nil {
		return nil,err
	}
	body,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}

	return body,nil
}

func TurnOffLight(light *models.Light,position *models.Position) ([]byte,error) {


	parms := url.Values{}
	parms.Set("code",light.CloseCode)
	parms.Add("controlCode",light.ControlCode)
	parms.Add("deviceId",light.DeviceId)
	parms.Add("deviceType",config.TestDeviceType)
	parms.Add("mac",config.TestMacAddress)
	parms.Add("positionId",light.PositonId)
	parms.Add("typeId",position.TypeId)
	return httpUtil.DoPost(config.Host+config.ControlUrl,&parms,"application/x-www-form-urlencoded")
}

func AdjustLightBrightness(light *models.Light,position *models.Position,value int) ([]byte,error) {
	parms := url.Values{}
	parms.Add("controlCode",light.AdjustCode)
	parms.Add("data",strconv.Itoa(value))
	parms.Add("deviceId",light.DeviceId)
	parms.Add("deviceType",config.TestDeviceType)
	parms.Add("mac",config.TestMacAddress)
	parms.Add("positionId",light.PositonId)
	parms.Add("typeId",position.TypeId)
	parmsStr :=  parms.Encode()
	resp,err := http.Post(config.Host+config.ControlUrl,"application/x-www-form-urlencoded",strings.NewReader(parmsStr))
	defer resp.Body.Close()
	if err != nil {
		return nil,err
	}
	body,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}

	return body,nil
}

func ControlWindowStatus(window *models.Curtain,position *models.Position,isOpen bool) ([]byte,error) {
	parms := url.Values{}
	var code string
	if isOpen {
		code = window.OpenCode
	} else {
		code = window.CloseCode
	}
	parms.Add("controlCode",window.ControlCode)
	parms.Add("controlCode",code)
	parms.Add("deviceId",window.DeviceId)
	parms.Add("deviceType",config.TestDeviceType)
	parms.Add("mac",config.TestMacAddress)
	parms.Add("positionId",window.PositionId)
	parms.Add("typeId",position.TypeId)
	parmsStr :=  parms.Encode()
	resp,err := http.Post(config.Host+config.ControlUrl,"application/x-www-form-urlencoded",strings.NewReader(parmsStr))
	defer resp.Body.Close()
	if err != nil {
		return nil,err
	}
	body,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}

	return body,nil
}