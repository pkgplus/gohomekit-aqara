package aqara

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

type Motion struct {
	*accessory.Accessory
	Motion *service.MotionSensor
}

func NewMotion(info accessory.Info) *Motion {
	acc := Motion{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.Motion = service.NewMotionSensor()

	acc.Motion.MotionDetected.SetValue(false)

	acc.AddService(acc.Motion.Service)

	return &acc
}

func (this *Aqara) InitMotion() {
	for _, mt := range this.manager.Magnets {
		info := accessory.Info{
			Name:         "OpenStatus",
			Manufacturer: "Aqara",
			SerialNumber: mt.Sid,
			Model:        mt.Model,
		}

		acc := NewMotion(info)
		acc.Motion.MotionDetected.SetValue(mt.Opened)
		this.AddAcc(acc.Accessory)

		//refresh
		go func() {
			for {
				<-mt.ReportChan
				LOGGER.Info("Magnet Detected:: open:%v", mt.Opened)
				acc.Motion.MotionDetected.SetValue(mt.Opened)
			}
		}()
	}

	return
}
