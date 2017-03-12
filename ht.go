package aqara

import (
	"github.com/brutella/hc/accessory"
	"github.com/xuebing1110/gohomekit/myaccessory"
)

func (this *Aqara) InitTempHumi() {
	for _, ht := range this.manager.SensorHTs {
		info := accessory.Info{
			Name:         "Temperature",
			Manufacturer: "Aqara",
			SerialNumber: ht.Sid,
			Model:        ht.Model,
		}
		acc := accessory.NewTemperatureSensor(info, 0, -20, 100, 0.1)
		acc.TempSensor.CurrentTemperature.SetValue(ht.Temperature)
		this.AddAcc(acc.Accessory)

		info2 := accessory.Info{
			Name:         "Humidity",
			Manufacturer: "Aqara",
			SerialNumber: ht.Sid,
			Model:        ht.Model,
		}
		acc2 := myaccessory.NewHumiditySensor(info2, 0, 0, 100, 0.1)
		acc2.HumiditySensor.CurrentRelativeHumidity.SetValue(ht.Humidity)
		this.AddAcc(acc2.Accessory)

		//refresh
		go func() {
			for {
				<-ht.ReportChan
				LOGGER.Info("Temperature:%f, Humidity:%f", ht.Temperature, ht.Humidity)
				acc.TempSensor.CurrentTemperature.SetValue(ht.Temperature)
				acc2.HumiditySensor.CurrentRelativeHumidity.SetValue(ht.Humidity)
			}
		}()
	}

	return
}
