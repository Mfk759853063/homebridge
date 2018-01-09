package models

type UserInfo struct{
	ApiBaseResp
	Data Data `json:"data"`
}

type Data struct {
	AppUserInfo User `json:appUserInfo`
	RoomList []Room `json:roomList`
}

type User struct {
	Token string `json:token`
	UserName string `json:userName`
	NickName string `json:nickName`
	PhoneNumber string `json:phoneNumber`
	RoomList []Room
}


type Room struct {
	RoomName string `json:roomName`
	IsDefault string `json:isDefault`
	TypeId string `json:typeId`
}


func (user *User) GetDefaultRoom() *Room {
	for _,room := range user.RoomList {
		if room.IsDefault == "1" {
			return &room
			break
		}
	}
	return nil
}