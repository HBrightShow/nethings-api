package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/nethings/internal-api/config"
	"github.com/nethings/internal-api/db"
	"fmt"
)

var log = logrus.WithField("module", "api")

type Config struct {
	Port           string
	DevCorsEnabled bool
	DevCorsHost    string
}

type API struct {
	engine *gin.Engine
	db     *db.DB

	packages []Package
}

type Package interface {
	SetRoutes(engine *gin.Engine)
}

func New(db *db.DB) *API {
	fmt.Println("8+++++++++++")
	return &API{
		db: db,
	}
}

func (a *API) Run() {

	fmt.Println("9+++++++++++")
	a.registerPackages()

	a.engine = gin.Default()

	if config.Store.API.DevCors {
		a.engine.Use(cors.New(cors.Config{
			AllowOrigins:     []string{config.Store.API.DevCorsHost},
			AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		}))
	}

	a.setRoutes()
	for _, p := range a.packages {
		p.SetRoutes(a.engine)
	}

	logrus.Infof("starting api on port %s", config.Store.API.Port)

	err := a.engine.Run(":" + config.Store.API.Port)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *API) Close() {
}
