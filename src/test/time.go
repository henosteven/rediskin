package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println(int64(time.Second / time.Millisecond)) //1000
	fmt.Println(time.Duration(10) * time.Second)       //10s
	fmt.Println(time.Duration(100000000000).Minutes()) //1.6666666666666665

	//time-> Now() Unix() Parse() -> time.Time->Unix() UnixNano() Format()
	//Now()
	//Unix(15123123123, 0) 时间戳转为time.Time
	//Parse('2016-01-02', '2017-11-11') 字符串转为time.Time

	fmt.Println(time.Now())                 //2017-11-22 19:16:43.391598563 +0800 CST
	fmt.Println(reflect.TypeOf(time.Now())) //time.Time

	fmt.Println(time.Now().Unix())                 //1511349453
	fmt.Println(reflect.TypeOf(time.Now().Unix())) //int64

	fmt.Println(time.Now().UnixNano())                 //1511349486973089316
	fmt.Println(reflect.TypeOf(time.Now().UnixNano())) //int64

	var timestamp int64 = 1511349453
	tm2 := time.Unix(timestamp, 0)
	fmt.Println(tm2)                               //2017-11-22 19:17:33 +0800 CST
	fmt.Println(reflect.TypeOf(tm2))               //time.Time
	fmt.Println(tm2.Format("2006-01-02 03:04:05")) //2017-11-22 07:17:33
	fmt.Println(tm2.Format("2006-01-02 15:04:05")) //2017-11-22 19:17:33

	//需要转换的参考样例 01/02/2006  | 2006-Jan-02  | Jan 2, 2006 at 3:04pm (MST)
	t, _ := time.Parse("01/02/2006", "06/21/2017")
	//t, _ := time.Parse("2006-Jan-02", "2017-Jun-21") //2017-06-21 00:00:00 +0000 UTC
	fmt.Println(t)                        //2017-06-21 00:00:00 +0000 UTC
	fmt.Println(reflect.TypeOf(t))        //time.Time
	fmt.Println(t.Unix())                 //1498003200
	fmt.Println(reflect.TypeOf(t.Unix())) //int64

}
