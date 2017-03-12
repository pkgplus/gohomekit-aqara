package aqara

import (
	"github.com/bingbaba/tool/color"
	"github.com/brutella/hc/accessory"
)

func (this *Aqara) InitLight() (err error) {
	gw := this.manager.GateWay

	info := accessory.Info{
		Name:         "AqaraGateWay",
		Manufacturer: "Aqara",
		SerialNumber: gw.Sid,
		Model:        gw.Model,
	}
	acc := accessory.NewLightbulb(info)
	acc.Lightbulb.Brightness.SetValue(100)

	//gateway light
	acc.Lightbulb.Hue.OnValueRemoteUpdate(func(hue float64) {
		LOGGER.Debug("Hue %f", hue)
		acc.Lightbulb.Hue.SetValue(hue)
		sat := acc.Lightbulb.Saturation.GetValue()
		bri := acc.Lightbulb.Brightness.GetValue()

		hsb := &color.HSV{uint(hue), uint(sat), uint(bri)}
		gw.ChangeColor(hsb)
		acc.Lightbulb.On.SetValue(true)
	})

	acc.Lightbulb.Saturation.OnValueRemoteUpdate(func(sat float64) {
		LOGGER.Debug("Saturation %f", sat)
		hue := acc.Lightbulb.Hue.GetValue()
		bri := acc.Lightbulb.Brightness.GetValue()

		hsb := &color.HSV{uint(hue), uint(sat), uint(bri)}
		gw.ChangeColor(hsb)
		acc.Lightbulb.On.SetValue(true)
	})

	acc.Lightbulb.Brightness.OnValueRemoteUpdate(func(bri int) {
		LOGGER.Debug("Brightness %d", bri)
		gw.ChangeBrightness(bri)
		acc.Lightbulb.Brightness.SetValue(bri)
		acc.Lightbulb.On.SetValue(true)
	})

	acc.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		var err error
		if on {
			err = gw.TurnOn()
		} else {
			err = gw.TurnOff()
		}
		if err != nil {
			this.OnError(err)
		}
	})

	this.AddAcc(acc.Accessory)

	//refresh
	go func() {
		for {
			<-gw.ReportChan
			if gw.RGB == 0 {
				acc.Lightbulb.On.SetValue(false)
			} else {
				acc.Lightbulb.On.SetValue(true)
			}
		}
	}()

	return
}
