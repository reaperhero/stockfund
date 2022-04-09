package service

import (
	"log"
	"math"
	"strings"
	"time"
)

func ParseStr(hqStr string) []string {
	baseData := strings.Split(hqStr, "\"")
	base := strings.Split(baseData[1], ",")
	return base
}

func Fetch_shsz(url string) string {
	client := ghttp.HttpClient(time.Second * 10)
	if client == nil {
		log.Println("get http client error")
		return ""
	}
	data := ghttp.HttpGet(client, url, "")
	base := ParseStr(data)
	str := gconvert.ConvertCharacterEncoding(base[0], "gb2312", "utf-8") + "\n"
	str += " 当前最新价：" + base[3] + " | 当前涨跌率："
	agoPrice := gconvert.StringToFloat64(base[2])
	nowPrice := gconvert.StringToFloat64(base[3])
	lv := ""
	if agoPrice > nowPrice {
		lv = "-"
	}
	upDown := math.Abs(nowPrice-agoPrice) / agoPrice * 100
	str += lv + gconvert.Float64ToString(upDown, 2) + "%\n"
	str += " 今日最高价：" + base[4] + " | 今日最低价：" + base[5] + "\n"
	str += " 今日开盘价：" + base[1] + " | 昨日收盘价：" + base[2] + "\n\n"
	return str
}

func Fetch_CompositeIndex(url string) string {
	client := ghttp.HttpClient(time.Second * 10)
	if client == nil {
		log.Println("get http client error")
		return ""
	}
	data := ghttp.HttpGet(client, url, "")
	base := ParseStr(data)
	str := gconvert.ConvertCharacterEncoding(base[0], "gb2312", "utf-8") + " 当前指数：" + base[1] + " | 涨跌额：" + base[2] + " | 涨跌率：" + base[3] + "%\n\n"
	return str
}
