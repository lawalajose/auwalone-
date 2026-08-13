package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	m "github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/model"
)

func ParserFxn(filePath string) ([]m.RawShipment, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open shipments file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip header
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("read CSV header: %w", err)
	}

	var shipments []m.RawShipment

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("read CSV record: %w", err)
		}

		if len(record) != 8 {
			return nil, fmt.Errorf(
				"invalid record: expected 8 columns, got %d",
				len(record),
			)
		}

		shipment := m.RawShipment{
			ShipmentID:           record[0],
			OriginRegion:         record[1],
			DestinationRegion:    record[2],
			ShipmentDate:         record[3],
			ExpectedDeliveryDate: record[4],
			ActualDeliveryDate:   record[5],
			ShipmentStatus:       record[6],
			Carrier:              record[7],
		}

		shipments = append(shipments, shipment)
	}

	return shipments, nil
}
