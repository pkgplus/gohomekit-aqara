package aqara

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

type Magnet struct {
	*accessory.Accessory
	Magnet *service.OccupancySensor
}

func NewMagnet(info accessory.Info) *Magnet {
	acc := Magnet{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.Magnet = service.NewOccupancySensor()
	acc.Magnet.OccupancyDetected.SetValue(0)

	acc.AddService(acc.Magnet.Service)

	return &acc
}

func (this *Aqara) InitMagnet() {
	for _, mt := range this.manager.Motions {
		info := accessory.Info{
			Name:         "Occupancy",
			Manufacturer: "Aqara",
			SerialNumber: mt.Sid,
			Model:        mt.Model,
		}

		acc := NewMagnet(info)
		if mt.IsMotorial {
			acc.Magnet.OccupancyDetected.SetValue(1)
		} else {
			acc.Magnet.OccupancyDetected.SetValue(0)
		}

		this.AddAcc(acc.Accessory)

		//refresh
		go func() {
			for {
				<-mt.ReportChan
				LOGGER.Info("Occupancy Sensor: %v", mt.IsMotorial)
				if mt.IsMotorial {
					acc.Magnet.OccupancyDetected.SetValue(1)
				} else {
					acc.Magnet.OccupancyDetected.SetValue(0)
				}

			}
		}()
	}

	return
}
