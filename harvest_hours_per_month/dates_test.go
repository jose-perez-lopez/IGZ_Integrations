package main

import (
	"testing"
)

func TestGetFirstDay(t *testing.T) {
	firstDayOfJan2020 := "20200101"
	firstDayOfDic2021 := "20211201"
	if getFirstDay(2020, 1) != firstDayOfJan2020 {
		t.Errorf("Expected %s have %s", firstDayOfJan2020, getFirstDay(2020, 1))
	}
	if getFirstDay(2021, 12) != firstDayOfDic2021 {
		t.Errorf("Expected %s have %s", firstDayOfDic2021, getFirstDay(2021, 12))
	}
}

func TestGetLastDay(t *testing.T) {
	lastDayOfJan2020 := "20200131"
	lastDayOfDic2021 := "20211231"
	if getLastDay(2020, 1) != lastDayOfJan2020 {
		t.Errorf("Expected %s have %s", lastDayOfJan2020, getLastDay(2020, 1))
	}
	if getLastDay(2021, 12) != lastDayOfDic2021 {
		t.Errorf("Expected %s have %s", lastDayOfDic2021, getLastDay(2021, 12))
	}
}
