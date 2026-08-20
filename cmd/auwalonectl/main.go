package main

import (
	"fmt"

	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/aggregate"
	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/parser"
	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/report"
	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/validator"
)

func main() {

	file := "shipmenttest.csv"

	rawData, err := parser.ParseShipments(file)
	fmt.Printf("%#v\n", rawData)
	fmt.Printf("%#v\n", err)

	cleanData, errr := validator.Validator(rawData)

	reportStruct := aggregate.Aggregator(cleanData, errr, file, rawData)

	report.Generate(reportStruct)

}
