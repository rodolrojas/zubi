package gateway

import (
	"fmt"
	"net/http"

	"rodolrojas.com/zubi/internal/common/structs"
	"rodolrojas.com/zubi/internal/config"
	"rodolrojas.com/zubi/internal/router"

	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type Gateway struct {
	config *structs.BaseServerConfig
	router *router.Router
}

func NewGateway() *Gateway {
	cfg, err := config.LoadConfig()
	rtr, err := router.NewRouter(cfg)
	if err != nil {
		logrus.Fatal("Cannot load configuration: ", err)
	}

	return &Gateway{
		config: cfg,
		router: rtr,
	}
}

func (g *Gateway) Start() error {

	logrus.Infof("Starting gateway server on port %d", g.config.Port)
	
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(g.router.GetHandler())
	
	srv := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf(":%d", g.config.Port),
		WriteTimeout: g.config.WriteTimeout,
		ReadTimeout:  g.config.ReadTimeout,
	}
	// appCommon.DebugStruct(&g)
	logrus.Infof("Gateway ready and listening at %s", srv.Addr)
	return srv.ListenAndServe()
}