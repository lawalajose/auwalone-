package parser

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/model"
)

func TestParseShipments(t *testing.T) {

	tests := []struct {
		name     string
		content  string
		csvPath  string
		expected []model.RawShipment
		err      error
	}{
		{
			name: "valid csv",
			content: `Shipment_ID,Origin_Region,Destination_Region,Shipment_Date,Expected_Delivery_Date,Actual_Delivery_Date,Shipment_Status,Carrier
SHP00001,East,Central,2026-04-02,2026-04-09,2026-04-08,Delivered,UPS
SHP00002,Central,West,2026-05-13,2026-05-19,2026-05-19,Delivered,UPS
SHP00003,North,South,2026-02-08,2026-02-15,,In Transit,FedEx
SHP00004,North,West,2026-04-15,2026-04-22,,In Transit,FedEx
SHP00005,North,South,2026-04-16,2026-04-20,2026-04-18,Delivered,DHL
SHP00006,East,East,2026-03-26,2026-04-01,2026-04-04,Delivered,UPS
SHP00007,South,South,2026-01-25,2026-01-31,2026-01-31,Delivered,FedEx
SHP00008,North,West,2026-05-19,2026-05-24,2026-05-24,Delivered,UPS
SHP00009,North,East,2026-06-06,2026-06-13,,Delayed,FedEx
SHP00010,East,West,2026-01-17,2026-01-21,,In Transit,FedEx
SHP00011,Central,North,2026-01-31,2026-02-05,2026-02-05,Delivered,DHL
SHP00012,East,North,2026-02-06,2026-02-12,,Delayed,FedEx
SHP00013,North,North,2026-01-16,2026-01-23,2026-01-23,Delivered,FedEx
SHP00014,West,South,2026-04-24,2026-04-29,2026-05-02,Delivered,FedEx
SHP00015,West,West,2026-05-26,2026-05-29,2026-05-30,Delivered,DHL
SHP00016,West,Central,2026-04-19,2026-04-25,,In Transit,UPS
SHP00017,Central,North,2026-06-03,2026-06-09,,Delayed,UPS
SHP00018,North,West,2026-05-29,2026-06-04,,In Transit,DHL
SHP00019,Central,South,2026-01-02,2026-01-08,2026-01-09,Delivered,UPS
SHP00020,North,West,2026-05-19,2026-05-27,2026-05-27,Delivered,DHL`,
			csvPath: "/shipments.csv",
			expected: []model.RawShipment{
				{
					ShipmentID:           "SHP00001",
					OriginRegion:         "East",
					DestinationRegion:    "Central",
					ShipmentDate:         "2026-04-02",
					ExpectedDeliveryDate: "2026-04-09",
					ActualDeliveryDate:   "2026-04-08",
					ShipmentStatus:       "Delivered",
					Carrier:              "UPS",
				},
				{
					ShipmentID:           "SHP00002",
					OriginRegion:         "Central",
					DestinationRegion:    "West",
					ShipmentDate:         "2026-05-13",
					ExpectedDeliveryDate: "2026-05-19",
					ActualDeliveryDate:   "2026-05-19",
					ShipmentStatus:       "Delivered",
					Carrier:              "UPS",
				},
				{
					ShipmentID:           "SHP00003",
					OriginRegion:         "North",
					DestinationRegion:    "South",
					ShipmentDate:         "2026-02-08",
					ExpectedDeliveryDate: "2026-02-15",
					ActualDeliveryDate:   "",
					ShipmentStatus:       "In Transit",
					Carrier:              "FedEx",
				},
				{
					ShipmentID:           "SHP00004",
					OriginRegion:         "North",
					DestinationRegion:    "West",
					ShipmentDate:         "2026-04-15",
					ExpectedDeliveryDate: "2026-04-22",
					ActualDeliveryDate:   "",
					ShipmentStatus:       "In Transit",
					Carrier:              "FedEx",
				},
				{
					ShipmentID:           "SHP00005",
					OriginRegion:         "North",
					DestinationRegion:    "South",
					ShipmentDate:         "2026-04-16",
					ExpectedDeliveryDate: "2026-04-20",
					ActualDeliveryDate:   "2026-04-18",
					ShipmentStatus:       "Delivered",
					Carrier:              "DHL",
				},
				{
					ShipmentID:           "SHP00006",
					OriginRegion:         "East",
					DestinationRegion:    "East",
					ShipmentDate:         "2026-03-26",
					ExpectedDeliveryDate: "2026-04-01",
					ActualDeliveryDate:   "2026-04-04",
					ShipmentStatus:       "Delivered",
					Carrier:              "UPS",
				},
				{
					ShipmentID:           "SHP00007",
					OriginRegion:         "South",
					DestinationRegion:    "South",
					ShipmentDate:         "2026-01-25",
					ExpectedDeliveryDate: "2026-01-31",
					ActualDeliveryDate:   "2026-01-31",
					ShipmentStatus:       "Delivered",
					Carrier:              "FedEx",
				},
				{
					ShipmentID:           "SHP00008",
					OriginRegion:         "North",
					DestinationRegion:    "West",
					ShipmentDate:         "2026-05-19",
					ExpectedDeliveryDate: "2026-05-24",
					ActualDeliveryDate:   "2026-05-24",
					ShipmentStatus:       "Delivered",
					Carrier:              "UPS",
				},
				{
					ShipmentID:           "SHP00009",
					OriginRegion:         "North",
					DestinationRegion:    "East",
					ShipmentDate:         "2026-06-06",
					ExpectedDeliveryDate: "2026-06-13",
					ActualDeliveryDate:   "",
					ShipmentStatus:       "Delayed",
					Carrier:              "FedEx",
				},
				{
					ShipmentID:           "SHP00010",
					OriginRegion:         "East",
					DestinationRegion:    "West",
					ShipmentDate:         "2026-01-17",
					ExpectedDeliveryDate: "2026-01-21",
					ActualDeliveryDate:   "",
					ShipmentStatus:       "In Transit",
					Carrier:              "FedEx",
				},
				{
					ShipmentID:           "SHP00011",
					OriginRegion:         "Central",
					DestinationRegion:    "North",
					ShipmentDate:         "2026-01-31",
					ExpectedDeliveryDate: "2026-02-05",
					ActualDeliveryDate:   "2026-02-05",
					ShipmentStatus:       "Delivered",
					Carrier:              "DHL",
				},
				{
					ShipmentID:           "SHP00012",
					OriginRegion:         "East",
					DestinationRegion:    "North",
					ShipmentDate:         "2026-02-06",
					ExpectedDeliveryDate: "2026-02-12",
					ActualDeliveryDate:   "",
					ShipmentStatus:       "Delayed",
					Carrier:              "FedEx",
				},
				{
					ShipmentID:           "SHP00013",
					OriginRegion:         "North",
					DestinationRegion:    "North",
					ShipmentDate:         "2026-01-16",
					ExpectedDeliveryDate: "2026-01-23",
					ActualDeliveryDate:   "2026-01-23",
					ShipmentStatus:       "Delivered",
					Carrier:              "FedEx",
				},
				{
					ShipmentID:           "SHP00014",
					OriginRegion:         "West",
					DestinationRegion:    "South",
					ShipmentDate:         "2026-04-24",
					ExpectedDeliveryDate: "2026-04-29",
					ActualDeliveryDate:   "2026-05-02",
					ShipmentStatus:       "Delivered",
					Carrier:              "FedEx",
				},
				{
					ShipmentID:           "SHP00015",
					OriginRegion:         "West",
					DestinationRegion:    "West",
					ShipmentDate:         "2026-05-26",
					ExpectedDeliveryDate: "2026-05-29",
					ActualDeliveryDate:   "2026-05-30",
					ShipmentStatus:       "Delivered",
					Carrier:              "DHL",
				},
				{
					ShipmentID:           "SHP00016",
					OriginRegion:         "West",
					DestinationRegion:    "Central",
					ShipmentDate:         "2026-04-19",
					ExpectedDeliveryDate: "2026-04-25",
					ActualDeliveryDate:   "",
					ShipmentStatus:       "In Transit",
					Carrier:              "UPS",
				},
				{
					ShipmentID:           "SHP00017",
					OriginRegion:         "Central",
					DestinationRegion:    "North",
					ShipmentDate:         "2026-06-03",
					ExpectedDeliveryDate: "2026-06-09",
					ActualDeliveryDate:   "",
					ShipmentStatus:       "Delayed",
					Carrier:              "UPS",
				},
				{
					ShipmentID:           "SHP00018",
					OriginRegion:         "North",
					DestinationRegion:    "West",
					ShipmentDate:         "2026-05-29",
					ExpectedDeliveryDate: "2026-06-04",
					ActualDeliveryDate:   "",
					ShipmentStatus:       "In Transit",
					Carrier:              "DHL",
				},
				{
					ShipmentID:           "SHP00019",
					OriginRegion:         "Central",
					DestinationRegion:    "South",
					ShipmentDate:         "2026-01-02",
					ExpectedDeliveryDate: "2026-01-08",
					ActualDeliveryDate:   "2026-01-09",
					ShipmentStatus:       "Delivered",
					Carrier:              "UPS",
				},
				{
					ShipmentID:           "SHP00020",
					OriginRegion:         "North",
					DestinationRegion:    "West",
					ShipmentDate:         "2026-05-19",
					ExpectedDeliveryDate: "2026-05-27",
					ActualDeliveryDate:   "2026-05-27",
					ShipmentStatus:       "Delivered",
					Carrier:              "DHL",
				},
			},
			err: nil,
		}, {
			name: "wrong number of columns",
			content: `Shipment_ID,Origin_Region,Destination_Region,Shipment_Date,Expected_Delivery_Date
SHP00001,East,Central,2026-04-02,2026-04-09`,
			csvPath: "/shipments.csv",
			err:     errors.New("invalid record: expected 8 columns, got 5"), // replace with the specific error your parser returns
		}, {
			name: "malformed CSV",
			content: `shipment_id,origin_region,destination_region,shipment_date,expected_delivery_date,actual_delivery_date,shipment_status,carrier
"SHP001,Lagos,Abuja
`,
			csvPath: "/shipments.csv",
			err:     errors.New("read CSV record"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := t.TempDir() + tt.csvPath

			if err := os.WriteFile(file, []byte(tt.content), 0644); err != nil {
				t.Fatal(err)
			}

			got, err := ParseShipments(file)

			if tt.err != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tt.err)
				}

				// if err.Error() != tt.err.Error() {
				// 	t.Errorf("expected error %q, got %q", tt.err.Error(), err.Error())
				// }
				if !strings.Contains(err.Error(), tt.err.Error()) {
					t.Errorf("expected error containing %q, got %q", tt.err.Error(), err.Error())
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("ParseShipments() mismatch (-want +got):\n%s", diff)
			}
		})

	}

}
