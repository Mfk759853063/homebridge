package main

import (
	"fmt"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"homebridge/service"
	"homebridge/models"
	"github.com/brutella/hc/util"
	"homebridge/syncSocket"
	"homebridge/config"
)

var user *models.User

func main() {
	socket:= &syncSocket.SyncSocket{
		Host:config.SmartSocketUrl,
		Port:config.SmartSocketPort,
	}
	socket.Start()
	//service.Login("13714834809","111111",func(loginData *models.User,err error) {
	//	if loginData != nil {
	//		fmt.Printf("登录成功 %v",loginData.UserName)
	//		user = loginData
	//		typeId := loginData.GetDefaultRoom().TypeId
	//		positions, _, err := service.FetchDevicesInfo(typeId)
	//		if err != nil {
	//			fmt.Println(err)
	//			panic(err)
	//		}
	//
	//		err = startHomeBridgeWithPosition(&positions)
	//		if err != nil {
	//			panic(err)
	//		}
	//
	//	} else {
	//		fmt.Printf("登录失败%v",err)
	//		panic(err)
	//	}
	//})
}

func startHomeBridgeWithPosition(positions *[]models.Position) error {

	fmt.Println("服务启动")
	bridgeInfo := accessory.Info{
		Name:         "homebridge",
		SerialNumber: util.MAC48Address(util.RandomHexString()),
		Manufacturer: "bridge",
		Model:        user.GetDefaultRoom().RoomName,
	}
	bridge := accessory.NewBridge(bridgeInfo)

	accs := make([]*accessory.Accessory, 0, 20)


	for _, position := range *positions {
		fmt.Println(position.PositionName)
		// 灯光
		for _, light := range position.DeviceLight {
			if light.DeviceId == "" {
				continue
			}
			//fmt.Printf("%v\n", position.PositionName, light.Name)
			name := fmt.Sprintf("%v", light.Name)
			lightInfo := accessory.Info{
				Name:         name,
				Manufacturer: position.PositionName,
				Model:        util.MAC48Address(util.RandomHexString()),
				SerialNumber: util.MAC48Address(util.RandomHexString()),
			}
			lightAccessory := accessory.NewLightbulb(lightInfo)
			setLightCallback(lightAccessory,&light,&position)
			accs = append(accs, lightAccessory.Accessory)
		}
		// 窗帘
		for _,curtain := range position.DeviceCurtain {
			name := fmt.Sprintf("%v", curtain.Name)
			windowInfo := accessory.Info{
				Name:         name,
				Manufacturer: position.PositionName,
				Model:        util.MAC48Address(util.RandomHexString()),
				SerialNumber: util.MAC48Address(util.RandomHexString()),
			}
			windowAccessory := accessory.NewWindow(windowInfo)
			accs = append(accs, windowAccessory.Accessory)
			setWindowCallback(windowAccessory,&curtain,&position)
		}
		// 空调
		for index,deviceAir := range position.DeviceAir {
			fmt.Println("空调 index  %d",index)
			airConditionInfo := accessory.Info{
				Name:"空调",
				Manufacturer: position.PositionName,
				Model:        util.MAC48Address(util.RandomHexString()),
				SerialNumber: util.MAC48Address(util.RandomHexString()),
			}
			airConditionAccessory := accessory.NewThermostat(airConditionInfo,16,16,32,1)
			accs = append(accs,airConditionAccessory.Accessory)
			setAirConditionCallback(airConditionAccessory, &deviceAir, &position)
		}


	}

	// 大门
	doorInfo := accessory.Info{
		Name:"大门",
		Manufacturer: "",
		Model:        util.MAC48Address(util.RandomHexString()),
		SerialNumber: util.MAC48Address(util.RandomHexString()),
	}
	doorAccessory := accessory.NewDoor(doorInfo)
	accs = append(accs,doorAccessory.Accessory)
	setDoorCallback(doorAccessory)

	t, err := hc.NewIPTransport(hc.Config{Pin: "85296312",StoragePath:"homekitPairDataBase"},bridge.Accessory, accs...)

	if err != nil {
		return err
	}
	hc.OnTermination(func() {
		fmt.Println("停止homebridge")
		t.Stop()
	})

	t.Start()

	return nil
}

func setDoorCallback(doorAccessory *accessory.Door) {
	doorAccessory.Door.TargetPosition.OnValueRemoteUpdate(func(value int) {
		fmt.Println("门目标状态 %d",value)
		maxValue := doorAccessory.Door.TargetPosition.MaxValue.(int)
		minValue := doorAccessory.Door.TargetPosition.MinValue.(int)
		if value == maxValue {
			doorAccessory.Door.CurrentPosition.SetValue(maxValue)
		} else {
			doorAccessory.Door.CurrentPosition.SetValue(minValue)
		}
	})
}

func setAirConditionCallback(airConditionAccessory *accessory.Thermostat,airCondition *models.DeviceAir, position *models.Position ) {


	airConditionAccessory.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(value float64) {
		fmt.Println("目标温度%s",value)
	})

	airConditionAccessory.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(value int) {
		fmt.Println("目标空调模式 %d",value)
	})
}

func setWindowCallback(windowAccessory *accessory.Window, curtain *models.Curtain, position *models.Position) {
	windowAccessory.Window.TargetPosition.OnValueRemoteUpdate(func(value int) {
		fmt.Println("TargetPosition ",value)
		maxValue := windowAccessory.Window.TargetPosition.MaxValue.(int)
		minValue := windowAccessory.Window.TargetPosition.MinValue.(int)
		if maxValue == value {
			fmt.Println("打开窗帘 ",value)
			_,err:=service.ControlWindowStatus(curtain,position,true)
			if err != nil {
				fmt.Println(err)
			} else {
				//windowAccessory.Window.CurrentPosition.SetValue(0)
				//PositionStateDecreasing int = 0
				//PositionStateIncreasing int = 1
				//PositionStateStopped    int = 2
				windowAccessory.Window.PositionState.SetValue(1)
				windowAccessory.Window.TargetPosition.SetValue(maxValue)
				windowAccessory.Window.CurrentPosition.SetValue(maxValue)
			}
		} else if minValue == value {
			fmt.Println("关闭窗帘 ",value)
			_,err := service.ControlWindowStatus(curtain,position,false)
			if err != nil {
				fmt.Println(err)
			} else {
				windowAccessory.Window.CurrentPosition.SetValue(0)
				windowAccessory.Window.TargetPosition.SetValue(minValue)
				windowAccessory.Window.CurrentPosition.SetValue(minValue)
			}

		}
	})
}

func setLightCallback(lightAccessroy *accessory.Lightbulb,light *models.Light, position *models.Position) {
	// 设置灯的开关
	lightAccessroy.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		if on {
			fmt.Println("开灯")
			_,err := service.TurnOnLight(light,position)
			if err != nil {
				fmt.Println(err)
			} else {
				lightAccessroy.Lightbulb.On.SetValue(true)
			}
		} else {
			fmt.Println("关灯")
			_,err := service.TurnOffLight(light,position)
			if err != nil {
				fmt.Println(err)
			} else {
				lightAccessroy.Lightbulb.On.SetValue(false)
			}
		}
	})
	// 设置灯的亮度
	lightAccessroy.Lightbulb.Brightness.OnValueRemoteUpdate(func(value int) {
		fmt.Println("灯亮度改变", value)
		_,err :=service.AdjustLightBrightness(light,position,value)
		if err != nil {
			fmt.Println(err)
		} else {
			lightAccessroy.Lightbulb.Brightness.SetValue(value)
		}
	})
}