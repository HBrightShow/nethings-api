package api

import (
	"github.com/nethings/internal-api/governance"
	"github.com/nethings/internal-api/notifications"
	"github.com/nethings/internal-api/smartalpha"
	"github.com/nethings/internal-api/smartexposure"
	"github.com/nethings/internal-api/smartyield"
	"github.com/nethings/internal-api/yieldfarming"
)

func (a *API) registerPackages() {
	a.packages = append(a.packages, governance.New(a.db))
	a.packages = append(a.packages, smartexposure.New(a.db))
	a.packages = append(a.packages, yieldfarming.New(a.db))
	a.packages = append(a.packages, smartyield.New(a.db))
	a.packages = append(a.packages, smartalpha.New(a.db))
	a.packages = append(a.packages, notifications.New(a.db))
}
