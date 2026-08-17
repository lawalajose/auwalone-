package main

import (
	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/aggregate"
	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/parser"
	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/report"
	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/validator"
)

func main() {

	file := "shipments_2000.csv"

	rawData, _ := parser.ParserFxn(file)

	cleanData, errr := validator.Validator(rawData)

	reportStruct := aggregate.Aggregator(cleanData, errr, file, rawData)

	report.Generate(reportStruct)

}
