package domain

import "time"

type Shipment struct {
	ShipmentID           string
	OriginRegion         string
	DestinationRegion    string
	ShipmentDate         time.Time
	ExpectedDeliveryDate time.Time
	ActualDeliveryDate   time.Time
	ShipmentStatus       string
	Carrier              string
}
