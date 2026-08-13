package report

import (
	"fmt"
	"io"

	m "github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/model"
)

func JSON(w io.Writer, r m.ReportFormat) string {

	report := fmt.Sprintf(`
	Shipment Report ──────────────────────────────────────────────────────────────────────
	%-24s %d
	%-24s %d


	SUMMARY ──────────────────────────────────────────────────────────────────────
	%-24s %d
	%-24s %d
	%-24s %d
	%-24s %.1f%%
	`,
		"Generated", r.GeneratedAt,
		"Input", r.InputFile,
		"Total rows", r.TotalRows,
		"Valid shipments", r.ValidShipments,
		"Invalid rows", r.InvalidRows,
		"Validation rate", r.ValidationRate,
	)

	return report
}
