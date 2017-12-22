package main

import (
	"fmt"
)

func showAppend() {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("len:%d, cap:%d, %v\n", len(s), cap(s), s)
	}
}

const (
	_ int = iota
	JAN
	FEB
	MAR
	APR
	MAY
	JUN
	JUL
	AGU
	SEP
	OCT
	NUM
	DEC
)

func appendPrefix(s []string) {
	for i, _ := range s {
		s[i] += "_"
	}
}

func appendPrefixArr(s [13]string) {
	for i, _ := range s {
		s[i] += "_"
	}
}

func displayCompositeType(input interface{}) {
	switch input.(type) {
	case []string:
		newinput := input.([]string)
		for i, v := range newinput {
			fmt.Println(i, v)
		}
	default:
		fmt.Println("un-support-type")
		fmt.Println(input)
	}
}

func main() {

	var month = [...]string{JAN: "Jan",
		FEB: "Feb",
		MAR: "Mar",
		APR: "Apr",
		MAY: "May",
		JUN: "Jun",
		JUL: "Jul",
		AGU: "Agu",
		SEP: "Sep",
		OCT: "Oct",
		NUM: "Num",
		DEC: "Dec"}
	//[ Jan Feb Mar Apr May Jun Jul Agu Sep Oct Num Dec]
	displayCompositeType(month)
	appendPrefixArr(month) //数组
	//[ Jan Feb Mar Apr May Jun Jul Agu Sep Oct Num Dec]
	displayCompositeType(month)
	appendPrefix(month[:]) //slice
	//[_ Jan_ Feb_ Mar_ Apr_ May_ Jun_ Jul_ Agu_ Sep_ Oct_ Num_ Dec_]
	displayCompositeType(month)
	showAppend()
}
