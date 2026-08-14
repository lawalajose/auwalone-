package main

import (
	"fmt"
	"os"

	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/aggregate"
	p "github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/parser"
	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/report"
	v "github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/validator"
)

func main() {

	file := "shipments_2000.csv"

	rawData, _ := p.ParserFxn(file)

	cleanData, errr := v.Validator(rawData)

	reportStruct := aggregate.Aggregator(cleanData, errr, file, len(rawData))

	data, err := os.Create("report.json")

	if err != nil {
		return
	}
	defer data.Close()

	result := report.JSON(data, reportStruct)
	fmt.Println(result)
}
