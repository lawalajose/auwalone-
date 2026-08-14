package aggregate

import (
	"time"

	m "github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/model"
)

func Aggregator(cleanData []m.CleanedShipment, errors []m.ValidationError, file string, lenRawaData int) m.ReportFormat {

	data := m.ReportFormat{}
	data.GeneratedAt = time.Now().UTC().Format(time.RFC3339)
	data.InputFile = file

	data.TotalRows = totalRows(cleanData, errors)
	data.ValidShipments = len(cleanData)
	data.InvalidRows = lenRawaData - len(cleanData)
	data.ValidationRate = (float64(data.ValidShipments) / float64(data.TotalRows)) * 100

	return data
}

func totalRows(cleanData []m.CleanedShipment, errors []m.ValidationError) int {
	return len(cleanData) + len(errors)
}
