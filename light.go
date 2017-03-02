package aqara

import (
	"github.com/bingbaba/tool/color"
	"github.com/brutella/hc/accessory"
	"log"
)

func (this *Aqara) InitLight() (err error) {
	gw := this.manager.GateWay

	info := accessory.Info{
		Name:         "AqaraLight",
		Manufacturer: "Aqara",
	}
	acc := accessory.NewLightbulb(info)
	acc.Lightbulb.Brightness.SetValue(100)

	//gateway light
	acc.Lightbulb.Hue.OnValueRemoteUpdate(func(hue float64) {
		log.Printf("Hue %f \n", hue)
		acc.Lightbulb.Hue.SetValue(hue)
		sat := acc.Lightbulb.Saturation.GetValue()
		bri := acc.Lightbulb.Brightness.GetValue()

		hsb := &color.HSV{uint(hue), uint(sat), uint(bri)}
		log.Printf("%+v", hsb)
		gw.ChangeColor(hsb)
		acc.Lightbulb.On.SetValue(true)
	})

	acc.Lightbulb.Saturation.OnValueRemoteUpdate(func(sat float64) {
		log.Printf("Saturation %f \n", sat)
		hue := acc.Lightbulb.Hue.GetValue()
		bri := acc.Lightbulb.Brightness.GetValue()

		hsb := &color.HSV{uint(hue), uint(sat), uint(bri)}
		log.Printf("%+v", hsb)
		gw.ChangeColor(hsb)
		acc.Lightbulb.On.SetValue(true)
	})

	acc.Lightbulb.Brightness.OnValueRemoteUpdate(func(bri int) {
		log.Printf("Brightness %d \n", bri)
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

	return
}
