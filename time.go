package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("")
	fmt.Println("现在的时间是\n")
	fmt.Println(now.String())
	location, _ := time.LoadLocation("Local")
	fmt.Println("\n命运开始的时间是\n")
	past := time.Date(2013, 11, 2, 0, 22, 0, 0, location)
	fmt.Println(past.String())
	sub := now.Sub(past)
	hours := sub.Hours()
	minutes := sub.Minutes()
	fmt.Println("\n我们已经度过了", int(hours/24/7), "个星期\n")
	fmt.Println("		", int((hours/24))%7, " 天")
	fmt.Println("	", int(hours)%24, " 小时")
	fmt.Println("		", int(minutes)%60, " 分")
	fmt.Println("	", int(sub.Seconds())%60, " 秒")
	fmt.Println("\n	", int(sub.Minutes()), "次心跳的爱你")
	for {

	}
}
