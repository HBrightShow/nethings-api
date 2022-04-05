package api

import (
	"github.com/nethings/internal-api/machine"
)

func (a *API) registerPackages() {
	a.packages = append(a.packages, machine.New(a.db))
}
