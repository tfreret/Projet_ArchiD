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

type WindSensor struct {
	sensor.Sensor
}

func (wSensor *WindSensor) GetActualizeMeasure() (sensor.Measurement, error) {
	if wSensor.Api.Use {
		apiResponse, err := apiClient.GetApiResponse(config.CHECKWX_URL+wSensor.Params.Airport+"/decoded", wSensor.Api.Key)
		if err != nil {
			log.Printf("Erreur lors de l'obtention de la réponse de l'API : %v", err)

			return sensor.Measurement{}, fmt.Errorf("échec lors de l'obtention de la mesure : %w", err)
		}
		if len(apiResponse.Data) == 0 {
			return sensor.Measurement{}, fmt.Errorf("réponse de l'API invalide")
		}
		return sensor.Measurement{TypeMesure: "Wind", Value: apiResponse.Data[0].Wind.SpeedKph, Timestamp: time.Now().UTC().Format(time.RFC3339)}, nil
	} else {
		return sensor.Measurement{TypeMesure: "Wind", Value: wSensor.NumberGenerator.GenerateRandomNumber(), Timestamp: time.Now().UTC().Format(time.RFC3339)}, nil
	}
}

func NewWindSensor(configSensor sensor.ConfigSensor) *WindSensor {
	wSensor := &WindSensor{}
	var min, max float64 = 0, 100
	start := min + rand.Float64()*(max-min)
	nbGenerator := randomSensor.NewNumberGenerator(start, min, max)
	wSensor.Sensor = sensor.NewSensor(wSensor, configSensor, *nbGenerator)
	return wSensor
}
