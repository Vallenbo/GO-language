package main

import (
	"fmt"
	"time"
)

func timeDemo(now time.Time) { // 1-timeDemo 时间对象的年月日时分秒
	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Println(year, month, day, hour, minute, second)
}

func main() {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)

	time.Now().Unix()                          // 返回自 1970 年 1 月 1 日 UTC 以来经过的秒数
	time.Sleep(time.Second * time.Duration(3)) //休眠3秒
	//timeDemo(now)
	//timestampDemo(now)
	//timesOperation(now)
	//TimeTick()
	//TimeFormat(now)
}

func timestampDemo(now time.Time) { // 2-timestampDemo 时间戳
	timestamp := now.Unix()  // 秒级时间戳
	milli := now.UnixMilli() // 毫秒时间戳 Go1.17+
	micro := now.UnixMicro() // 微秒时间戳 Go1.17+
	nano := now.UnixNano()   // 纳秒时间戳
	fmt.Println(timestamp, milli, micro, nano)

	ret := time.Unix(timestamp, 0) //将时间戳转为time时间对象
	fmt.Println(ret, ret.Year(), ret.YearDay())
}

func timesOperation(now time.Time) { //3-时间间隔常用操作
	now.Add(time.Hour) //在现在的时间基础上，添加一小时 //add()参数为d Duration：如time.Hour
	parse, err := time.Parse("2020-01-02 15:04:15", "2023-01-02 15:04:15")
	if err != nil {
		return
	}
	d := parse.UTC().Sub(now.UTC()) //减：时间差值计算
	fmt.Println(d)
}

func TimeTick() { //定时器
	ticker := time.Tick(time.Second) //定义一个定时器,1秒间隔的
	for i := range ticker {
		fmt.Println(i) //每过1秒执行一次
	}
}

func TimeFormat(now time.Time) { // formatDemo 时间格式化
	// 格式化的模板为 2006-01-02 15:04:05
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))    // 24小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan")) // 12小时制

	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
	fmt.Println(now.Format("2006/01/02 15:04:05.000")) // 2022/02/27 00:10:42.960
	// 小数点后写9，会省略末尾可能出现的0
	fmt.Println(now.Format("2006/01/02 15:04:05.999")) // 2022/02/27 00:10:42.96
	fmt.Println(now.Format("15:04:05"))                // 只格式化时分秒部分
	fmt.Println(now.Format("2006.01.02"))              // 只格式化日期部分
}

func parseTime() { // parseDemo 指定时区解析时间
	timeObj, err := time.Parse("2006/01/02 15:04:05", "2022/10/05 11:25:20") // 在没有时区指示符的情况下，time.Parse 返回UTC时间
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) // 2022-10-05 11:25:20 +0000 UTC

	// 在有时区指示符的情况下，time.Parse 返回对应时区的时间表示
	timeObj, err = time.Parse(time.RFC3339, "2022-10-05T11:25:20+08:00") // RFC3339     = "2006-01-02T15:04:05Z07:00"
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) // 2022-10-05 11:25:20 +0800 CST
}

func parseDemo(now time.Time) { // parseDemo 解析时间
	fmt.Println(now)
	loc, err := time.LoadLocation("Asia/Shanghai") // 加载时区
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2022/10/05 11:25:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}
