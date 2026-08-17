package main

import (
	"fmt"

	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/parser"
	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/validator"
)

func main() {

	file := "shipments_2000.csv"

	rawData, _ := parser.ParserFxn(file)

	_, errr := validator.Validator(rawData)
	fmt.Println(errr)

	// reportStruct := aggregate.Aggregator(cleanData, errr, file, rawData)

	// report.Generate(reportStruct)

}
