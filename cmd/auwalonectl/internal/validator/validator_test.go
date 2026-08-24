package validator

import (
	"testing"
	"time"

	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/model"
)

func TestValidator(t *testing.T) {
	tests := []struct {
		name              string
		content           []model.RawShipment
		expectedShipments []model.CleanedShipment
		expectedErrors    []model.ShipmentErrorReport
	}{
		{
			name: "valid shipment",
			content: []model.RawShipment{
				{
					ShipmentID:           "SHP00001",
					OriginRegion:         "East",
					DestinationRegion:    "West",
					ShipmentDate:         "2026-08-01",
					ExpectedDeliveryDate: "2026-08-05",
					ActualDeliveryDate:   "2026-08-04",
					ShipmentStatus:       "Delivered",
					Carrier:              "DHL",
				},
			},
			expectedShipments: []model.CleanedShipment{
				{
					ShipmentID:           "SHP00001",
					OriginRegion:         "East",
					DestinationRegion:    "West",
					ShipmentDate:         time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC),
					ExpectedDeliveryDate: time.Date(2026, 8, 5, 0, 0, 0, 0, time.UTC),
					ActualDeliveryDate:   time.Date(2026, 8, 4, 0, 0, 0, 0, time.UTC),
					ShipmentStatus:       "Delivered",
					Carrier:              "DHL",
				},
			},
			expectedErrors: []model.ShipmentErrorReport{},
		},

		{
			name: "invalid shipment",
			content: []model.RawShipment{
				{
					ShipmentID:           "INVALID",
					OriginRegion:         "InvalidRegion",
					DestinationRegion:    "InvalidRegion",
					ShipmentDate:         "invalid-date",
					ExpectedDeliveryDate: "invalid-date",
					ActualDeliveryDate:   "invalid-date",
					ShipmentStatus:       "InvalidStatus",
					Carrier:              "InvalidCarrier",
				},
			},
			expectedShipments: []model.CleanedShipment{},
			expectedErrors: []model.ShipmentErrorReport{
				{
					RawShipment: model.RawShipment{
						ShipmentID:           "INVALID",
						OriginRegion:         "InvalidRegion",
						DestinationRegion:    "InvalidRegion",
						ShipmentDate:         "invalid-date",
						ExpectedDeliveryDate: "invalid-date",
						ActualDeliveryDate:   "invalid-date",
						ShipmentStatus:       "InvalidStatus",
						Carrier:              "InvalidCarrier",
					},
					Errors: []model.ValidationError{
						{
							Row:     1,
							Column:  "ShipmentID",
							Value:   "INVALID",
							Message: "Invalid shipment ID",
						},
						{
							Row:     1,
							Column:  "Origin Region",
							Value:   "InvalidRegion",
							Message: "Invalid Origin Region, Expected (East, West, South, Central, North)",
						},
						{
							Row:     1,
							Column:  "Destination Region",
							Value:   "InvalidRegion",
							Message: "Invalid Origin Region, Expected (East, West, South, Central, North)",
						},
						{
							Row:     1,
							Column:  "Shipment Date",
							Value:   "invalid-date",
							Message: "Invalid Date Format, expected YYYY-MM-DD format",
						},
						{
							Row:     1,
							Column:  "Expected Delivery Date",
							Value:   "invalid-date",
							Message: "Invalid Date Format, expected YYYY-MM-DD format",
						},
						{
							Row:     1,
							Column:  "Actual Delivery Date",
							Value:   "invalid-date",
							Message: "Invalid Date Format, expected YYYY-MM-DD format",
						},
						{
							Row:     1,
							Column:  "Shipment Status",
							Value:   "InvalidStatus",
							Message: "Invalid Shipment Staus, expected (Delivered, In Transit, Delayed, Cancelled)",
						},
						{
							Row:     1,
							Column:  "Carrier",
							Value:   "InvalidCarrier",
							Message: "Invalid Carrier, expected (UPS, FedEx, DHL)",
						},
					},
				},
			},
		},

		{
			name:              "empty input",
			content:           []model.RawShipment{},
			expectedShipments: []model.CleanedShipment{},
			expectedErrors:    []model.ShipmentErrorReport{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShipments, gotErrors := Validator(tt.content)

			if len(gotShipments) != len(tt.expectedShipments) {
				t.Errorf(
					"expected %d cleaned shipments, got %d",
					len(tt.expectedShipments),
					len(gotShipments),
				)
			}

			if len(gotErrors) != len(tt.expectedErrors) {
				t.Errorf(
					"expected %d error reports, got %d",
					len(tt.expectedErrors),
					len(gotErrors),
				)
			}

			if len(tt.expectedShipments) > 0 {
				got := gotShipments[0]
				expected := tt.expectedShipments[0]

				if got.ShipmentID != expected.ShipmentID {
					t.Errorf(
						"expected ShipmentID %q, got %q",
						expected.ShipmentID,
						got.ShipmentID,
					)
				}

				if got.OriginRegion != expected.OriginRegion {
					t.Errorf(
						"expected OriginRegion %q, got %q",
						expected.OriginRegion,
						got.OriginRegion,
					)
				}

				if got.DestinationRegion != expected.DestinationRegion {
					t.Errorf(
						"expected DestinationRegion %q, got %q",
						expected.DestinationRegion,
						got.DestinationRegion,
					)
				}

				if !got.ShipmentDate.Equal(expected.ShipmentDate) {
					t.Errorf(
						"expected ShipmentDate %v, got %v",
						expected.ShipmentDate,
						got.ShipmentDate,
					)
				}

				if !got.ExpectedDeliveryDate.Equal(expected.ExpectedDeliveryDate) {
					t.Errorf(
						"expected ExpectedDeliveryDate %v, got %v",
						expected.ExpectedDeliveryDate,
						got.ExpectedDeliveryDate,
					)
				}

				if !got.ActualDeliveryDate.Equal(expected.ActualDeliveryDate) {
					t.Errorf(
						"expected ActualDeliveryDate %v, got %v",
						expected.ActualDeliveryDate,
						got.ActualDeliveryDate,
					)
				}

				if got.ShipmentStatus != expected.ShipmentStatus {
					t.Errorf(
						"expected ShipmentStatus %q, got %q",
						expected.ShipmentStatus,
						got.ShipmentStatus,
					)
				}

				if got.Carrier != expected.Carrier {
					t.Errorf(
						"expected Carrier %q, got %q",
						expected.Carrier,
						got.Carrier,
					)
				}
			}

			if len(tt.expectedErrors) > 0 {
				if len(gotErrors[0].Errors) != len(tt.expectedErrors[0].Errors) {
					t.Errorf(
						"expected %d validation errors, got %d",
						len(tt.expectedErrors[0].Errors),
						len(gotErrors[0].Errors),
					)
				}
			}
		})
	}
}
