package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qday-wallet/bridge-analysis/config"
	"github.com/qday-wallet/bridge-analysis/server"
	"github.com/sunjiangjun/xlog"
)

func main() {

	var configPath string
	flag.StringVar(&configPath, "config", "./config.json", "The system file of config")
	flag.Parse()
	if len(configPath) < 1 {
		panic("can not find config file")
	}
	cfg := config.LoadConfig(configPath)

	if cfg.Log == nil {
		cfg.Log.Delay = 2
		cfg.Log.Path = "./log/analysis"
	}

	LOG := xlog.NewXLogger().BuildOutType(xlog.FILE).BuildLevel(xlog.InfoLevel).BuildFormatter(xlog.FORMAT_JSON).BuildFile(cfg.Log.Path, 24*time.Hour)

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	root := e.Group(cfg.Root)
	root.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: LOG.Out}))

	srv := server.NewHandler(cfg.DB, LOG)
	root.POST("/monitor", srv.Monitor)
	root.POST("/query", srv.QueryTxs)
	root.POST("/analysis/income", srv.QueryTxs)
	root.POST("/analysis/pay", srv.QueryTxs)

	err := e.Run(fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		panic(err)
	}

}
