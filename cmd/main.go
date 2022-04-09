package main

import "time"

type CodeLink struct {
	grail       []string
	singleStock []string
}

func NewCodeLink() *CodeLink {
	task := &CodeLink{}
	task.grail = []string{
		"http://hq.sinajs.cn/list=s_sh000001",
		"http://hq.sinajs.cn/list=s_sz399001",
		"http://hq.sinajs.cn/list=s_sz399006",
	}
	task.singleStock = []string{
		"http://hq.sinajs.cn/list=sz150304",
		"http://hq.sinajs.cn/list=sz000582",
		"http://hq.sinajs.cn/list=sz150153",
		"http://hq.sinajs.cn/list=sz150224",
		"http://hq.sinajs.cn/list=sz150282",
		"http://hq.sinajs.cn/list=sz150201",
		"http://hq.sinajs.cn/list=sz150224",
		"http://hq.sinajs.cn/list=sz150236",
		"http://hq.sinajs.cn/list=sh502012",
		"http://hq.sinajs.cn/list=sh502055",
	}
	return task
}

func (c *CodeLink) Task() {
	u := NewUser()
	for _, grail := range c.grail {
		u.body += crawl.Fetch_CompositeIndex(grail)
	}
	for _, singleStock := range c.singleStock {
		u.body += crawl.Fetch_shsz(singleStock)
	}
	u.subject = "股票当前情况"
	u.mailType = "plain"
	gemail.SendToMail(u.host, u.user, u.password, u.to, u.mailType, u.subject, u.body)
}

func (c *CodeLink) IsStockTime() bool {
	if !gtime.IsWorkDay() {
		return false
	}
	now := time.Now()
	minute := now.Hour()*60 + now.Minute()
	if (minute >= 570 && minute <= 690) || (minute >= 780 && minute <= 900) {
		return true
	}
	return false
}

func main() {
	codeLink := NewCodeLink()
	ticker := time.NewTicker(time.Minute * 10)
	for _ = range ticker.C {
		if codeLink.IsStockTime() {
			codeLink.Task()
		}
	}
}
