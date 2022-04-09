// Package datacenter 数据来源
package datacenter

import (
	"github.com/reaperhero/stockfund/datacenter/chinabond"
	"github.com/reaperhero/stockfund/datacenter/eastmoney"
	"github.com/reaperhero/stockfund/datacenter/eniu"
	"github.com/reaperhero/stockfund/datacenter/sina"
	"github.com/reaperhero/stockfund/datacenter/zszx"
)

var (
	// EastMoney 东方财富
	EastMoney eastmoney.EastMoney
	// Eniu 亿牛网
	Eniu eniu.Eniu
	// Sina 新浪财经
	Sina sina.Sina
	// Zszx 招商证券
	Zszx zszx.Zszx
	// ChinaBond 中国债券信息网
	ChinaBond chinabond.ChinaBond
)

func init() {
	EastMoney = eastmoney.NewEastMoney()
	Eniu = eniu.NewEniu()
	Sina = sina.NewSina()
	Zszx = zszx.NewZszx()
	ChinaBond = chinabond.NewChinaBond()
}
