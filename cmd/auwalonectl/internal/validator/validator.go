package validator

import (
	"fmt"
	"time"

	p "github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/parser"
)

type CleanedShipment struct {
	ShipmentID           string
	OriginRegion         string
	DestinationRegion    string
	ShipmentDate         time.Time
	ExpectedDeliveryDate time.Time
	ActualDeliveryDate   time.Time // use IsZero() to check if missing
	ShipmentStatus       string
	Carrier              string
}

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

type validationError struct {
	Row     int
	Column  string
	Value   string
	Message string
}

func Validator() error {

	var errors []validationError

	var clean_shipment []CleanedShipment

	rawData, err := p.ParserFxn()
	if err != nil {
		return fmt.Errorf("Unable to Parse data: %w", err)
	}

	for i, data := range rawData {

		if !validShipmentID(data.ShipmentID) {
			errr := validationError{
				Row:     i + 1,
				Column:  "ShipmentID",
				Value:   data.ShipmentID,
				Message: "Invalid shipment ID",
			}

			errors = append(errors, errr)
		}

		if !validRegionn(data.OriginRegion) {
			errr := validationError{
				Row:     i + 1,
				Column:  "Origin Region",
				Value:   data.OriginRegion,
				Message: "Invalid Origin Region, Expected (East, West, South, Central, North)",
			}

			errors = append(errors, errr)

		}

		if !validRegionn(data.DestinationRegion) {
			errr := validationError{
				Row:     i + 1,
				Column:  "Destination Region",
				Value:   data.DestinationRegion,
				Message: "Invalid Origin Region, Expected (East, West, South, Central, North)",
			}

			errors = append(errors, errr)

		}

		shipment_date, isDate := validDate(data.ShipmentDate)
		if !isDate {
			errr := validationError{
				Row:     i + 1,
				Column:  "Shipment Date",
				Value:   data.ShipmentDate,
				Message: "Invalid Date Format, expected 2006-01-02 format",
			}

			errors = append(errors, errr)
		}

		expDelivery_date, isDate := validDate(data.ExpectedDeliveryDate)
		if !isDate {
			errr := validationError{
				Row:     i + 1,
				Column:  "Expected Delivery Date",
				Value:   data.ExpectedDeliveryDate,
				Message: "Invalid Date Format, expected 2006-01-02 format",
			}

			errors = append(errors, errr)
		}

		actDelivery_date, isDate := validDate(data.ExpectedDeliveryDate)
		if !isDate {
			errr := validationError{
				Row:     i + 1,
				Column:  "Actual Delivery Date",
				Value:   data.ActualDeliveryDate,
				Message: "Invalid Date Format, expected 2006-01-02 format",
			}

			errors = append(errors, errr)
		}

		if !validStatus[data.ShipmentStatus] {
			errr := validationError{
				Row:     i + 1,
				Column:  "Shipment Status",
				Value:   data.ShipmentStatus,
				Message: "Invalid Shipment Staus, expected (Delivered, In Transit, Delayed, Cancelled)",
			}

			errors = append(errors, errr)

		}

		if !validCarriers[data.Carrier] {
			errr := validationError{
				Row:     i + 1,
				Column:  "Carrier",
				Value:   data.ShipmentStatus,
				Message: "Invalid Carrier, expected (UPS, FedEx, DHL)",
			}

			errors = append(errors, errr)

		}

		clnshipment := CleanedShipment{
			ShipmentID:           data.ShipmentID,
			OriginRegion:         data.OriginRegion,
			DestinationRegion:    data.DestinationRegion,
			ShipmentDate:         shipment_date,
			ExpectedDeliveryDate: expDelivery_date,
			ActualDeliveryDate:   actDelivery_date, // use IsZero() to check if missing
			ShipmentStatus:       data.ShipmentStatus,
			Carrier:              data.Carrier,
		}

		if len(errors) != 0 {
			continue

		}

		clean_shipment = append(clean_shipment, clnshipment)

	}
	for _, e := range errors {
		fmt.Println(e.Value)

	}

	return nil
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
