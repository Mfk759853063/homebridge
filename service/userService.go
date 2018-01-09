package service

import (
	"homebridge/httpUtil"
	"homebridge/config"
	"fmt"
	"crypto/md5"
	"encoding/json"
	"homebridge/models"
)

func Login(name string, pwd string, callback func(userInfo *models.User, err error)) {
	data := []byte(pwd)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	url := fmt.Sprintf("%v?proprietorKey=%v&pwd=%v",config.Host+config.LoginUrl,name,md5str)
	body,err:= httpUtil.DoGet(url,"application/x-www-form-urlencoded")
	if err != nil {
		if callback != nil {
			callback(nil,err)
		}
	} else {
		var resp models.UserInfo
		err := json.Unmarshal(body, &resp)
		if err != nil	 {
			if callback != nil {
				callback(nil,err)
			}
		} else {
			user:=resp.Data.AppUserInfo
			user.RoomList = resp.Data.RoomList
			if callback != nil {
				callback(&user,nil)
			}

		}

	}
}