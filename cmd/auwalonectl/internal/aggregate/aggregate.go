package aggregate

import (
	"fmt"
	"time"

	m "github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/model"
)

func Aggregator(cleanData []m.CleanedShipment, errors []m.ValidationError, file string) m.ReportFormat {

	data := m.ReportFormat{}
	data.GeneratedAt = time.Now().UTC().Format(time.RFC3339)
	data.InputFile = file

	data.TotalRows = totalRows(cleanData, errors)
	data.ValidShipments = len(cleanData)
	data.InvalidRows = len(errors)
	data.ValidationRate = validationRate(cleanData, errors)

	return data
}

func totalRows(cleanData []m.CleanedShipment, errors []m.ValidationError) int {
	return len(cleanData) + len(errors)
}

func validationRate(cleanData []m.CleanedShipment, errors []m.ValidationError) float64 {
	fmt.Println(len(cleanData))
	fmt.Println(len(errors))
	fmt.Println(float64(len(cleanData) / totalRows(cleanData, errors)))
	return float64(len(cleanData)/totalRows(cleanData, errors)) * 100
}
