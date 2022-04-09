// Package cron 定时任务
package cron

import (
	"context"

	"github.com/axiaoxin-com/goutils"
	"github.com/reaperhero/stockfund/datacenter"
	"github.com/reaperhero/stockfund/models"
)

// SyncBond 同步债券
func SyncBond() {
	if !goutils.IsTradingDay() {
		return
	}
	ctx := context.Background()
	syl := datacenter.ChinaBond.QueryAAACompanyBondSyl(ctx)
	if syl != 0 {
		models.AAACompanyBondSyl = syl
	}
}
