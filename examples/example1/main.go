package main

import (
	"fmt"
	lsm303 "github.com/bskari/go-lsm303/pkg/lsm303"
	"log"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
	"time"
)

func main() {
	// Load i2c drivers
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	accelerometer, err := lsm303.NewAccelerometer(bus, &lsm303.DefaultAccelerometerOpts)
	if err != nil {
		log.Fatal("Couldn't connect to accelerometer")
	}

	magnetometer, err := lsm303.NewMagnetometer(bus, &lsm303.DefaultMagnetometerOpts)
	if err != nil {
		log.Fatal("Couldn't connect to magnetometer")
	}

	for {
		xa, ya, za, err := accelerometer.Sense()
		if err != nil {
			log.Fatalf("Got error %v", err)
		}
		fmt.Printf("accel x:%v y:%v z:%v\n", xa, ya, za)
		xr, yr, zr, err := accelerometer.SenseRaw()
		if err != nil {
			log.Fatalf("Got error %v", err)
		}
		fmt.Printf("raw accel x:%v y:%v z:%v\n", xr, yr, zr)

		// The periph.io has units defined for many things, but not for
		// magnetometer flux, so we only have SenseRaw
		xm, ym, zm, err := magnetometer.SenseRaw()
		if err != nil {
			log.Fatalf("Got error %v", err)
		}
		fmt.Printf("raw mag x:%v y:%v z:%v\n", xm, ym, zm)

		time.Sleep(time.Second * 1)
	}

	// Examples for setting options
	/*
		accelerometer.SetRange(ACCELEROMETER_RANGE_16G)
		accelerometer.SetMode(ACCELEROMETER_MODE_LOW_POWER)
		magnetometer.SetGain(MAGNETOMETER_GAIN_5_6)
		magnetometer.SetRate(MAGNETOMETER_RATE_75)
	*/
}
