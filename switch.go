package aqara

import (
	"github.com/brutella/hc/accessory"
	//"github.com/xuebing1110/gohomekit/myaccessory"
	"github.com/xuebing1110/migateway"
	"time"
)

/*
func (this *Aqara) InitSwitchTest() {
	for _, dev := range this.manager.Switchs {
		info := accessory.Info{
			Name:         "Switch",
			Manufacturer: "Aqara",
			SerialNumber: dev.Sid,
			Model:        dev.Model,
		}
		acc := myaccessory.NewSwitch(info)
		this.AddAcc(acc.Accessory)

		//refresh
		go func() {
			for {
				msg := <-dev.ReportChan
				status, ok := msg.(string)
				if !ok {
					continue
				}
				LOGGER.Info("switch:: %s", status)

				if status == migateway.SWITCH_STATUS_CLICK {
					acc.Switch.SetValue(1)
				} else if status == migateway.SWITCH_STATUS_DOUBLECLICK {
					acc.Switch.SetValue(2)
				}
			}
		}()
	}

	return
}*/

func (this *Aqara) InitSwitch() {
	for _, dev := range this.manager.Switchs {
		info := accessory.Info{
			Name:         "Switch_Click",
			Manufacturer: "Aqara",
			SerialNumber: dev.Sid,
			Model:        dev.Model,
		}
		acc := accessory.NewSwitch(info)
		acc.Switch.On.OnValueRemoteUpdate(func(on bool) {
			LOGGER.Warn("switch click: %v", on)
		})
		this.AddAcc(acc.Accessory)

		info2 := accessory.Info{
			Name:         "Switch_DoubleClick",
			Manufacturer: "Aqara",
			SerialNumber: dev.Sid,
			Model:        dev.Model,
		}
		acc2 := accessory.NewSwitch(info2)
		acc2.Switch.On.OnValueRemoteUpdate(func(on bool) {
			LOGGER.Warn("switch double click: %v", on)
		})
		this.AddAcc(acc2.Accessory)

		//refresh
		go func() {
			for {
				msg := <-dev.ReportChan
				status, ok := msg.(string)
				if !ok {
					continue
				}
				LOGGER.Info("switch:: %s", status)

				if status == migateway.SWITCH_STATUS_CLICK {
					acc.Switch.On.SetValue(true)
					go func() {
						time.Sleep(30 * time.Second)
						acc.Switch.On.SetValue(false)
					}()
				} else if status == migateway.SWITCH_STATUS_DOUBLECLICK {
					acc2.Switch.On.SetValue(true)
					go func() {
						time.Sleep(30 * time.Second)
						acc2.Switch.On.SetValue(false)
					}()
				}
			}
		}()
	}

	return
}
