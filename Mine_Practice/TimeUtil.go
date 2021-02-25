package main

import (
	"fmt"
	"strconv"
	"time"
)

var timeLayoutStr = "2021-01-02 15:04:05" //go中的时间格式化必须是这个时间

/*

解读：格式化他的月份的时候只要在Format的传参中输入1、01、Jan、January等三个字符的时候就会将时间在这几个值的对应位置将时间对应的月份替换掉的意思。
举例，当前时间是2020-08-01 11:46:30按照以下的Format传参将1、01、Jan、January转换为当前月份对应的格式比如Jan转为为Aug

	fmt.Println(time.Now().Format("现在是1月份"))
	fmt.Println(time.Now().Format("现在是01月份"))
	fmt.Println(time.Now().Format("现在是Jan月份"))
	fmt.Println(time.Now().Format("现在是January月份"))

	现在是8月份
	现在是08月份
	现在是Aug月份
	现在是August月

时间格式对照表:
月份 1,01,Jan,January
日　 2,02,_2
时　 3,03,15,PM,pm,AM,am
分　 4,04
秒　 5,05
年　 06,2006
时区 -07,-0700,Z0700,Z07:00,-07:00,MST
周几 Mon,Monday

*/

func NowTimeFun() {
	//1、获取当前时间

	currentTime := time.Now() //获取当前时间，类型是Go的时间类型Time
	fmt.Println("当前时间：", currentTime)

	t1 := time.Now().Year() //年
	fmt.Println("年", t1)

	t2 := time.Now().Format("01") //月
	fmt.Println("月", t2)

	t3 := time.Now().Day() //日
	fmt.Println("日", t3)

	t4 := time.Now().Hour() //小时
	fmt.Println("小时", t4)

	t5 := time.Now().Minute() //分钟
	fmt.Println("分钟", t5)

	t6 := time.Now().Second() //秒
	fmt.Println("秒", t6)

	t7 := time.Now().Nanosecond() //纳秒
	fmt.Println("纳秒", t7)
}

func SubDateFun() {

	// Add 时间相加
	now := time.Now()
	fmt.Println("当前时间：", now)
	fmt.Println("当前时间-毫秒s：", strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
	fmt.Println("当前时间-毫秒：", time.Now().UnixNano()/1e6)

	// 10分钟前
	m, _ := time.ParseDuration("-10m")
	m1 := now.Add(m)
	fmt.Println("10分钟前：", m1)

	// 8个小时前
	h, _ := time.ParseDuration("-8h")
	h1 := now.Add(8 * h)
	fmt.Println("8个小时前：", h1)

	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println("一天前：", d1)

	// 10分钟后
	mm, _ := time.ParseDuration("10m")
	mm1 := now.Add(mm)
	fmt.Println("10分钟后：", mm1)

	// 8小时后
	hh, _ := time.ParseDuration("8h")
	hh1 := now.Add(hh)
	fmt.Println(hh1)

	// 一天后
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	fmt.Println("一天后：", dd1)

	// Sub 计算两个时间差
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes(), "分钟-差")

	sumH := now.Sub(h1)
	fmt.Println(sumH.Hours(), "小时-差")

	sumD := now.Sub(d1)
	fmt.Printf("%v 天-差\n", sumD.Hours()/24)

}

func main() {

	// 当前日期测试[年-月-日 小时-分钟-秒-毫秒]
	NowTimeFun()

	fmt.Println("\n\n==============================================分割==================================================== \n\n")

	// 时间加减[日期相减，小时相减，月份相减]
	SubDateFun()

}
