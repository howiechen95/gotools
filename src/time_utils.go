package main

import (
	"fmt"
	"time"
)

import "strings"

func test_000() {
	format := "2006-01-02 15:04:05"
	now := time.Now()
	//now, _ := time.Parse(format, time.Now().Format(format))
	a, _ := time.Parse(format, "2015-03-10 11:00:00")
	b, _ := time.Parse(format, "2015-03-10 16:mc00:00")
	fmt.Println(now.Unix(), a.Unix(), b.Unix())

	fmt.Println(now.Format(format), a.Format(format), b.Format(format))
	fmt.Println(now.After(a))
	fmt.Println(now.Before(a))
	fmt.Println(now.After(b))
	fmt.Println(now.Before(b))
	fmt.Println(a.After(b))
	fmt.Println(a.Before(b))
	fmt.Println(now.Format(format), a.Format(format), b.Format(format))
	fmt.Println(now.Unix(), a.Unix(), b.Unix())
}

func test_001() {
	// Add 时间相加
	now := time.Now()
	// ParseDuration parses a duration string.
	// A duration string is a possibly signed sequence of decimal numbers,
	// each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	//  Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// 10分钟前
	m, _ := time.ParseDuration("-1m")
	fmt.Println(m)
	m1 := now.Add(m)
	fmt.Println(m1)
	fmt.Println("----------------------------")
	// 8个小时前
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(8 * h)
	fmt.Println(h1)
	fmt.Println("----------------------------")
	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1)
	fmt.Println("----------------------------")
	printSplit(50)
	fmt.Println("----------------------------")
	// 10分钟后
	mm, _ := time.ParseDuration("1m")
	mm1 := now.Add(mm)
	fmt.Println(mm1)
	fmt.Println("----------------------------")
	// 8小时后
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(hh)
	fmt.Println(hh1)
	fmt.Println("----------------------------")
	// 一天后
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	fmt.Println(dd1)

	printSplit(50)
	fmt.Println("----------------------------")
	// Sub 计算两个时间差
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes(), "分钟")
	fmt.Println("----------------------------")
	sumH := now.Sub(h1)
	fmt.Println(sumH.Hours(), "小时")
	fmt.Println("----------------------------")
	sumD := now.Sub(d1)
	fmt.Printf("%v 天\n", sumD.Hours()/24)

}

func printSplit(count int) {
	fmt.Println(strings.Repeat("#", count))
}
func main() {
	d := time.Now().AddDate(1, 0, 0).Unix()
	fmt.Println(time.Now().Unix(), d)
}
