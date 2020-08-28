package main

import (
	"time"
)

func getFirstDay(year int, month int) string {
	t := time.Date(year, time.Month(month), 01, 00, 00, 00, 00, time.Now().Location())
	result := t.Format("20060102")
	return result
}

func getLastDay(year int, month int) string {
	t := time.Date(year, time.Month(month), 01, 00, 00, 00, 00, time.Now().Location())	
	t = t.AddDate(0, 1, -1)	
	result := t.Format("20060102")
	return result
}
