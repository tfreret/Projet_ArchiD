package main

import (
	"airport/internal/apiClient"
	"airport/internal/config"
	"airport/internal/randomSensor"
	"airport/internal/sensor"
	"fmt"
	"math/rand"
	"time"
)

type TempSensor struct {
	sensor.Sensor
}

func (tSensor *TempSensor) GetActualizeMeasure() (sensor.Measurement, error) {
	if tSensor.ConfigSensor.Api.Use {
		apiResponse, err := apiClient.GetApiResponse(config.CHECKWX_URL+tSensor.Params.Airport+"/decoded", tSensor.Api.Key)
		if err != nil {
			log.Printf("Erreur lors de l'obtention de la réponse de l'API : %v", err)

			return sensor.Measurement{}, fmt.Errorf("échec lors de l'obtention de la mesure : %w", err)
		}
		if len(apiResponse.Data) == 0 {
			return sensor.Measurement{}, fmt.Errorf("réponse de l'API invalide")
		}
		return sensor.Measurement{TypeMesure: "Temp", Value: apiResponse.Data[0].Temperature.Celsius, Timestamp: time.Now().UTC().Format(time.RFC3339)}, nil
	} else {
		return sensor.Measurement{TypeMesure: "Temp", Value: tSensor.NumberGenerator.GenerateRandomNumber(), Timestamp: time.Now().UTC().Format(time.RFC3339)}, nil
	}
}

func NewTempSensor(configSensor sensor.ConfigSensor) *TempSensor {
	tSensor := &TempSensor{}
	var min, max float64 = -10, 45
	start := min + rand.Float64()*(max-min)
	nbGenerator := randomSensor.NewNumberGenerator(start, min, max)
	tSensor.Sensor = sensor.NewSensor(tSensor, configSensor, *nbGenerator)
	return tSensor
}
