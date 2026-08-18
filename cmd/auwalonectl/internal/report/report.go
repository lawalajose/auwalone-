package report

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/model"
)

const (
	textReportFile = "report.txt"
	jsonReportFile = "report.json"
)

// Generate creates both report.txt and report.json.
func Generate(r model.ReportFormat) error {
	if err := writeTextReport(r); err != nil {
		return err
	}

	if err := writeJSONReport(r); err != nil {
		return err
	}

	return nil
}

func writeTextReport(r model.ReportFormat) error {
	var b strings.Builder

	b.WriteString("RawShipment Report\n")
	b.WriteString("──────────────────────────────────────────────────────────────────────\n")
	fmt.Fprintf(&b, "%-24s %s\n", "Generated", r.GeneratedAt)
	fmt.Fprintf(&b, "%-24s %s\n", "Input", r.InputFile)

	b.WriteString("\n")
	b.WriteString("SUMMARY\n")
	b.WriteString("──────────────────────────────────────────────────────────────────────\n")
	fmt.Fprintf(&b, "%-24s %d\n", "Total rows", r.TotalRows)
	fmt.Fprintf(&b, "%-24s %d\n", "Valid shipments", r.ValidShipments)
	fmt.Fprintf(&b, "%-24s %d\n", "Invalid rows", r.InvalidRows)
	fmt.Fprintf(&b, "%-24s %.1f%%\n", "Validation rate", r.ValidationRate)

	b.WriteString("\n")
	b.WriteString("ON-TIME PERFORMANCE\n")
	b.WriteString("──────────────────────────────────────────────────────────────────────\n")
	fmt.Fprintf(&b, "%-24s %d\n", "On-time shipments", r.OnTimeShipments)
	fmt.Fprintf(&b, "%-24s %d\n", "Late shipments", r.LateShipments)
	fmt.Fprintf(&b, "%-24s %.1f%%\n", "On-time rate", r.OnTimePercentage)
	fmt.Fprintf(&b, "%-24s %.1f\n", "Average delivery days", r.AverageDeliveryDays)

	b.WriteString("\n")
	b.WriteString("DELIVERY PERFORMANCE BY REGION\n")
	b.WriteString("──────────────────────────────────────────────────────────────────────\n")
	b.WriteString("Region             Shipments    On-time %    Avg Delivery\n")
	b.WriteString("──────────────────────────────────────────────────────────────────────\n")
	writeRegionPerformance(&b, r.Regions)
	b.WriteString("──────────────────────────────────────────────────────────────────────\n")

	b.WriteString("\n\n\n\n")
	b.WriteString("SHIPMENT VALIDATION ERROR REPORT\n")
	b.WriteString("======================================================================\n")

	if len(r.ShipmentError) == 0 {
		b.WriteString("No validation errors found.\n")
	} else {
		for _, report := range r.ShipmentError {
			writeShipmentError(&b, report)
		}
	}

	if err := os.WriteFile(textReportFile, []byte(b.String()), 0644); err != nil {
		return fmt.Errorf("write %s: %w", textReportFile, err)
	}

	return nil
}

func writeShipmentError(b *strings.Builder, report model.ShipmentErrorReport) {
	if len(report.Errors) == 0 {
		return
	}

	s := report.RawShipment

	fmt.Fprintf(b, "\nRow %d\n", report.Errors[0].Row)
	b.WriteString("----------------------------------------\n")
	fmt.Fprintf(b, "Shipment ID:             %s\n", s.ShipmentID)
	fmt.Fprintf(b, "Origin Region:           %s\n", s.OriginRegion)
	fmt.Fprintf(b, "Destination Region:      %s\n", s.DestinationRegion)
	fmt.Fprintf(b, "Shipment Date:           %s\n", s.ShipmentDate)
	fmt.Fprintf(b, "Expected Delivery Date:  %s\n", s.ExpectedDeliveryDate)
	fmt.Fprintf(b, "Actual Delivery Date:    %s\n", s.ActualDeliveryDate)
	fmt.Fprintf(b, "Shipment Status:         %s\n", s.ShipmentStatus)
	fmt.Fprintf(b, "Carrier:                 %s\n", s.Carrier)

	b.WriteString("\nErrors:\n")

	for _, err := range report.Errors {
		fmt.Fprintf(b, "- %s: %s\n", err.Column, err.Message)
		b.WriteString("\n\n")
	}
}

func writeJSONReport(r model.ReportFormat) error {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal report to JSON: %w", err)
	}

	data = append(data, '\n')

	if err := os.WriteFile(jsonReportFile, data, 0644); err != nil {
		return fmt.Errorf("write %s: %w", jsonReportFile, err)
	}

	return nil
}

func writeRegionPerformance(b *strings.Builder, regions []model.RegionPerformance) {
	for _, d := range regions {
		fmt.Fprintf(b, "%-16s %10d %10.1f%% %12.1f days\n", d.Region, d.Shipments, d.OnTimePercentage, d.AverageDeliveryDays)
	}
}
