package aggregate

import (
	"time"

	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/model"
)

func Aggregator(cleanData []model.CleanedShipment, errors []model.ShipmentErrorReport, file string, rawaData []model.RawShipment) model.ReportFormat {

	data := model.ReportFormat{}
	data.GeneratedAt = time.Now().UTC().Format(time.RFC3339)
	data.InputFile = file

	data.TotalRows = len(rawaData)
	data.ValidShipments = len(cleanData)
	data.InvalidRows = len(rawaData) - len(cleanData)
	data.ValidationRate = (float64(data.ValidShipments) / float64(data.TotalRows)) * 100

	a, b, c, d := onTimePerformance(cleanData)
	data.OnTimeShipments = a
	data.LateShipments = b
	data.OnTimePercentage = c
	data.AverageDeliveryDays = d

	data.Regions, data.Overall = deliveryRegion(cleanData)
	data.ShipmentError = errors

	return data
}

func onTimePerformance(cleanData []model.CleanedShipment) (int, int, float64, float64) {

	var delivered, onTime, totalDays int

	for _, data := range cleanData {

		if data.ShipmentStatus != "Delivered" {
			continue
		}

		delivered++
		if !data.ActualDeliveryDate.After(data.ExpectedDeliveryDate) {
			onTime++

		}

		days := int(data.ActualDeliveryDate.Sub(data.ShipmentDate).Hours() / 24)
		totalDays += days

	}

	averageDeliveryDays := float64(totalDays) / float64(delivered)
	rate := (float64(onTime) / float64(delivered)) * 100
	return onTime, (delivered - onTime), rate, averageDeliveryDays

}

func performaceByRegion(cleanData []model.CleanedShipment, region string) (int, float64, float64) {

	var delivered, onTime, totalDays int

	for _, data := range cleanData {

		if data.ShipmentStatus != "Delivered" || data.DestinationRegion != region {
			continue
		}

		delivered++
		if !data.ActualDeliveryDate.After(data.ExpectedDeliveryDate) {
			onTime++

		}

		days := int(data.ActualDeliveryDate.Sub(data.ShipmentDate).Hours() / 24)
		totalDays += days

	}

	averageDeliveryDays := float64(totalDays) / float64(delivered)
	rate := (float64(onTime) / float64(delivered)) * 100
	return delivered, rate, averageDeliveryDays

}

func deliveryRegion(cleanData []model.CleanedShipment) ([]model.RegionPerformance, model.OverAllPerformance) {
	var sliceRegions []model.RegionPerformance
	var overall model.OverAllPerformance
	var sliceRegion model.RegionPerformance
	var totalShipment int
	var totalOntime, totalAverageDays float64

	for k := range model.ValidRegion {

		a, b, c := performaceByRegion(cleanData, k)
		totalShipment += a
		totalOntime += b
		totalAverageDays += c

		sliceRegion.Region = k
		sliceRegion.Shipments = a
		sliceRegion.OnTimePercentage = b
		sliceRegion.AverageDeliveryDays = c

		sliceRegions = append(sliceRegions, sliceRegion)
	}
	overall.OverAllShipments = totalShipment
	overall.OverAllOnTimePercentage = totalOntime / float64(len(model.ValidRegion))
	overall.OverAllAverageDeliveryDays = totalAverageDays / float64(len(model.ValidRegion))

	return sliceRegions, overall
}
