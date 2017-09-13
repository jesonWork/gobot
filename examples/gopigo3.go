// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	g "gobot.io/x/gobot/platforms/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	raspiAdaptor := raspi.NewAdaptor()
	gopigo3 := g.NewGoPiGo3Driver(raspiAdaptor)

	work := func() {
		on := uint8(0xFF)
		gobot.Every(1000*time.Millisecond, func() {
			err := gopigo3.SetLED(g.LED_EYE_RIGHT, 0x00, 0x00, on)
			if err != nil {
				fmt.Println(err)
			}
			err = gopigo3.SetLED(g.LED_EYE_LEFT, ^on, 0x00, 0x00)
			if err != nil {
				fmt.Println(err)
			}
			on = ^on
		})
	}

	robot := gobot.NewRobot("gopigo3",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{gopigo3},
		work,
	)

	robot.Start()
}
