package main

import (
	"fmt"
	"os"

	a "github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/aggregate"
	p "github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/parser"
	"github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/report"
	v "github.com/lawalajose/auwalone-/cmd/auwalonectl/internal/validator"
)

func main() {

	file := "shipments_2000.csv"

	rawData, err := p.ParserFxn(file)
	fmt.Println(err)

	cleanData, errr := v.Validator(rawData)
	reportStruct := a.Aggregator(cleanData, errr, file)
	data, err := os.Create("report.json")

	if err != nil {
		return
	}
	defer data.Close()

	result := report.JSON(data, reportStruct)
	fmt.Println(result)
}
