package time_utils

import (
	"fmt"
	"time"
)

func DemoTime(strDeadline string) {

	fmt.Println(time.Now().Format("2006-01-02T15:04:05Z"))

	layout := "2006-01-02T15:04:05Z"

	tDeadline, err := time.Parse(layout, strDeadline)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("tDeadline: ", tDeadline.String())

	workorderDeadline := tDeadline.Add(-(time.Duration(48) * time.Hour))
	fmt.Println("workorderDeadline: ", workorderDeadline.String())

	currentTime := time.Now()
	currentTimePlusHour := currentTime.Add((time.Duration(1) * time.Hour))

	fmt.Println("Current Time plus 1 hour in String: ", currentTimePlusHour.String())

	if workorderDeadline.Before(currentTimePlusHour) {
		fmt.Println("deadline is in past: ", workorderDeadline.String())
		workorderDeadline = currentTimePlusHour
	}

	fmt.Println("target deadline is: ", workorderDeadline.Format("layout"))
	fmt.Println("target deadline is: ", workorderDeadline.String())

}
