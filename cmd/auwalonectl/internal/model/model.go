package model

import "time"

type RawShipment struct {
	ShipmentID           string
	OriginRegion         string
	DestinationRegion    string
	ShipmentDate         string
	ExpectedDeliveryDate string
	ActualDeliveryDate   string
	ShipmentStatus       string
	Carrier              string
}
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

type ValidationError struct {
	Row     int
	Column  string
	Value   string
	Message string
}
type ShipmentErrorReport struct {
	RawShipment RawShipment
	Errors      []ValidationError
}

type ReportFormat struct {
	// Metadata
	GeneratedAt string `json:"generatedAt"`
	InputFile   string `json:"inputFile"`

	// Summary
	TotalRows      int     `json:"totalRows"`
	ValidShipments int     `json:"validShipments"`
	InvalidRows    int     `json:"invalidRows"`
	ValidationRate float64 `json:"validationRate"`

	// Delivery performance
	OnTimeShipments     int     `json:"onTimeShipments"`
	LateShipments       int     `json:"lateShipments"`
	OnTimePercentage    float64 `json:"onTimePercentage"`
	AverageDeliveryDays float64 `json:"averageDeliveryDays"`

	// Regional performance
	Regions []RegionPerformance `json:"regions"`

	// Data quality
	DataQuality DataQuality `json:"dataQuality"`

	//Error report
	ShipmentError []ShipmentErrorReport
}

type RegionPerformance struct {
	Region              string  `json:"region"`
	Shipments           int     `json:"shipments"`
	OnTimePercentage    float64 `json:"onTimePercentage"`
	AverageDeliveryDays float64 `json:"averageDeliveryDays"`
}

type DataQuality struct {
	MissingShipmentIDs   int `json:"missingShipmentIds"`
	InvalidDates         int `json:"invalidDates"`
	MissingRegions       int `json:"missingRegions"`
	InvalidDeliveryTimes int `json:"invalidDeliveryTimes"`
	DuplicateShipmentIDs int `json:"duplicateShipmentIds"`
}
