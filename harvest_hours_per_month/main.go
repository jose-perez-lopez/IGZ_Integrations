package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

// HoursPerMonth Hours resported by people in harvest
type HoursPerMonth struct {
	Year          int
	Month         int
	Hours         int
	BillableHours int
}

func main() {
	//fmt.Println("1H")
	//fmt.Print(string(getHoursFromHarvers()))
	lambda.Start(HandleRequest)
}

// HandleRequest  ... the real thing :-D
func HandleRequest(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: string(getHoursFromHarvers()), StatusCode: 200}, nil
}

func getHoursFromHarvers() []byte {
	fmt.Println("2H")
	t := time.Now()
	var hours []HoursPerMonth

	for year := 2019; year <= t.Year(); year++ {
		endMonth := 12
		if year == t.Year() {
			endMonth = int(t.Month())
		}
		for month := 1; month <= endMonth; month++ {

			initDate := getFirstDay(year, month)
			lastDate := getLastDay(year, month)
			listSlice := callHarvestHoursAPI(initDate, lastDate)

			var billableHours float64 = 0
			var totalHours float64 = 0
			for _, value := range listSlice {
				billableHours += value.(map[string]interface{})["billable_hours"].(float64)
				totalHours += value.(map[string]interface{})["total_hours"].(float64)
			}
			hoursPerMonth := HoursPerMonth{year,
				month,
				int(math.Round(totalHours)),
				int(math.Round(billableHours))}
			hours = append(hours, hoursPerMonth)
		}
	}

	jsonHours, _ := json.Marshal(hours)
	return jsonHours
}
