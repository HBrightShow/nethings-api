package machine

import (
	"github.com/gin-gonic/gin"
	"github.com/nethings/internal-api/db"
)

type Machine struct {
	db *db.DB
}

func New(db *db.DB) *Machine {
	return &Machine{db: db}
}

func (g *Machine) SetRoutes(engine *gin.Engine) {
	machine := engine.Group("/api/machine")
	machine.GET("/status", g.GetStatus)

}
