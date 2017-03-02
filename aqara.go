package aqara

import (
	"github.com/brutella/hc/accessory"
	"github.com/xuebing1110/gohomekit/bridge"
	"github.com/xuebing1110/gohomekit/myaccessory"
	"github.com/xuebing1110/migateway"
	"time"
)

type Aqara struct {
	*bridge.BasePlatForm
	sid      string
	password string
	manager  *migateway.AqaraManager
}

func (this *Aqara) New(sid, pwd string) bridge.PlatForm {
	return &Aqara{
		BasePlatForm: bridge.NewBasePlatForm(),
		sid:          sid,
		password:     pwd,
	}
}

func (this *Aqara) Init() (err error) {
	mgwConf := migateway.NewConfig()
	mgwConf.AESKey = this.password
	this.manager, err = migateway.NewAqaraManager(mgwConf)
	if err != nil {
		return
	}

	//gateway online status
	//this.initGateWay()

	//light
	err = this.InitLight()
	if err != nil {
		return
	}

	//temp && humi sensor
	this.InitTempHumi()

	return
}

func (this *Aqara) GetName() string {
	return "AqaraPlatform"
}

func (this *Aqara) Start() error {
	return nil
}

func (this *Aqara) Stop() error {
	return nil
}

func (this *Aqara) initGateWay() {
	//accessory
	info := accessory.Info{
		Name:         "GateWay",
		Manufacturer: "Aqara",
	}
	acc := myaccessory.NewBridgeStatus(info)
	this.AddAcc(acc.Accessory)

	//check heartbeat
	go func() {
		for {
			time.Sleep(time.Minute)
			if this.manager.GateWay.GetHeartTime() < time.Now().Unix() {
				acc.BridgingState.Reachable.SetValue(false)
			} else {
				acc.BridgingState.Reachable.SetValue(true)
			}
		}
	}()
}
