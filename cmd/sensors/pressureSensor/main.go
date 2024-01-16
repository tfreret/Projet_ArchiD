package main

import (
	"airport/internal/sensor"
	"airport/internal/config"
	"flag"
	"fmt"
)

func main() {
	var (
		configFile = flag.String("config", "sensor-pressure-config.yaml", "Config file of the sensor")
	)
	flag.Parse()
	configSensor := config.ReadConfig[sensor.ConfigSensor](*configFile)

	fmt.Println("Using config :", configSensor)
	NewPressureSensor(configSensor).StartSendingData()
}
