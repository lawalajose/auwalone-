package validator

import (
	"time"

	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/model"
)

var validRegion = map[string]bool{
	"East":    true,
	"West":    true,
	"South":   true,
	"Central": true,
	"North":   true,
}

var validCarriers = map[string]bool{
	"UPS":   true,
	"FedEx": true,
	"DHL":   true,
}

var validStatus = map[string]bool{
	"Delivered":  true,
	"In Transit": true,
	"Delayed":    true,
	"Cancelled":  true,
}

func Validator(rawData []model.RawShipment) ([]model.CleanedShipment, []model.ShipmentErrorReport) {

	var shipmentError []model.ShipmentErrorReport

	var clean_shipment []model.CleanedShipment
	count := 0

	for i, data := range rawData {
		var errors []model.ValidationError

		if !validShipmentID(data.ShipmentID) {
			count++

			errr := model.ValidationError{
				Row:     i + 1,
				Column:  "ShipmentID",
				Value:   data.ShipmentID,
				Message: "Invalid shipment ID",
			}

			errors = append(errors, errr)
		}

		if !validRegionn(data.OriginRegion) {
			count++
			errr := model.ValidationError{
				Row:     i + 1,
				Column:  "Origin Region",
				Value:   data.OriginRegion,
				Message: "Invalid Origin Region, Expected (East, West, South, Central, North)",
			}

			errors = append(errors, errr)

		}

		if !validRegionn(data.DestinationRegion) {
			count++
			errr := model.ValidationError{
				Row:     i + 1,
				Column:  "Destination Region",
				Value:   data.DestinationRegion,
				Message: "Invalid Origin Region, Expected (East, West, South, Central, North)",
			}

			errors = append(errors, errr)

		}

		shipment_date, isDate := validDate(data.ShipmentDate)
		if !isDate {
			count++
			errr := model.ValidationError{

				Row:     i + 1,
				Column:  "Shipment Date",
				Value:   data.ShipmentDate,
				Message: "Invalid Date Format, expected 2006-01-02 format",
			}

			errors = append(errors, errr)
		}

		expDelivery_date, isDate := validDate(data.ExpectedDeliveryDate)
		if !isDate {
			count++
			errr := model.ValidationError{
				Row:     i + 1,
				Column:  "Expected Delivery Date",
				Value:   data.ExpectedDeliveryDate,
				Message: "Invalid Date Format, expected 2006-01-02 format",
			}

			errors = append(errors, errr)
		}

		actDelivery_date, isDate := validDate(data.ActualDeliveryDate)
		if !isDate {
			count++
			errr := model.ValidationError{
				Row:     i + 1,
				Column:  "Actual Delivery Date",
				Value:   data.ActualDeliveryDate,
				Message: "Invalid Date Format, expected 2006-01-02 format",
			}

			errors = append(errors, errr)
		}

		if !validStatus[data.ShipmentStatus] {
			count++
			errr := model.ValidationError{
				Row:     i + 1,
				Column:  "Shipment Status",
				Value:   data.ShipmentStatus,
				Message: "Invalid Shipment Staus, expected (Delivered, In Transit, Delayed, Cancelled)",
			}

			errors = append(errors, errr)

		}

		if !validCarriers[data.Carrier] {
			count++
			errr := model.ValidationError{
				Row:     i + 1,
				Column:  "Carrier",
				Value:   data.Carrier,
				Message: "Invalid Carrier, expected (UPS, FedEx, DHL)",
			}
			errors = append(errors, errr)
		}

		shipment_err := model.ShipmentErrorReport{
			RawShipment: data,
			Errors:      errors,
		}

		clnshipment := model.CleanedShipment{
			ShipmentID:           data.ShipmentID,
			OriginRegion:         data.OriginRegion,
			DestinationRegion:    data.DestinationRegion,
			ShipmentDate:         shipment_date,
			ExpectedDeliveryDate: expDelivery_date,
			ActualDeliveryDate:   actDelivery_date, // use IsZero() to check if missing
			ShipmentStatus:       data.ShipmentStatus,
			Carrier:              data.Carrier,
		}

		if count != 0 {
			count = 0
			continue
		}

		clean_shipment = append(clean_shipment, clnshipment)
		shipmentError = append(shipmentError, shipment_err)
	}

	return clean_shipment, shipmentError
}

func validShipmentID(id string) bool {
	return !(len(id) != 8 || id == "")
}

func validRegionn(reg string) bool {
	return validRegion[reg]
}

func validDate(date string) (time.Time, bool) {
	layout := "2006-01-02"

	newDate, err := time.Parse(layout, date)
	if err != nil {
		return newDate, false
	}
	return newDate, true
}
