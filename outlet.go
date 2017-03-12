package aqara

import (
	"github.com/brutella/hc/accessory"
)

func (this *Aqara) InitPlug() {
	for _, plug := range this.manager.Plugs {
		info := accessory.Info{
			Name:         "Outlet",
			Manufacturer: "Aqara",
			SerialNumber: plug.Sid,
			Model:        plug.Model,
		}

		acc := accessory.NewOutlet(info)
		acc.Outlet.OutletInUse.SetValue(plug.InUse)
		acc.Outlet.On.SetValue(plug.IsOn)

		acc.Outlet.On.OnValueRemoteUpdate(func(on bool) {
			var err error
			if on {
				err = plug.TurnOn()
			} else {
				err = plug.TurnOff()
			}

			if err != nil {
				LOGGER.Error("Write to plug err: %v", err)
			}
		})

		this.AddAcc(acc.Accessory)

		//refresh
		go func() {
			for {
				<-plug.ReportChan
				LOGGER.Info("Plug:: isUse:%v isOn:%v", plug.InUse, plug.IsOn)
				acc.Outlet.OutletInUse.SetValue(plug.InUse)
				acc.Outlet.On.SetValue(plug.IsOn)
			}
		}()
	}

	return
}
