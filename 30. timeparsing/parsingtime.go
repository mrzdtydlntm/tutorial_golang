package main

import (
	"fmt"
	"time"
)

func main() {
	//parsing time
	/*
		var layoutFormat, value string
		var date time.Time

		layoutFormat = "2006-01-02 15:04:05"
		value = "2015-09-02 08:04:00"
		date, _ = time.Parse(layoutFormat, value)
		fmt.Println(value, "\t->", date.String())
		// 2015-09-02 08:04:00 +0000 UTC

		layoutFormat = "02/01/2006 MST"
		value = "02/09/2015 WIB"
		date, _ = time.Parse(layoutFormat, value)
		fmt.Println(value, "\t\t->", date.String())
		// 2015-09-02 00:00:00 +0700 WIB
	*/

	//parsed using predefined time
	/*
		var date, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
		fmt.Println(date.String())
	*/

	//time.Time ke string
	var date, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")

	var dateS1 = date.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1", dateS1)
	// Wednesday 02, September 2015 08:00 WIB

	var dateS2 = date.Format(time.RFC3339)
	fmt.Println("dateS2", dateS2)
	// 2015-09-02T08:00:00+07:00
}
