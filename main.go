// Code generated by hertz generator.

package main

import (
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/joho/godotenv"
	"github.com/lyleshaw/mlops/biz/dal"
	"github.com/lyleshaw/mlops/biz/mw"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatalf("Some error occurred. Err: %s", err)
	}

	dal.Init()
	mw.InitJwt()
	h := server.New(server.WithHostPorts(":7654"))
	h.Use(mw.Cors())
	register(h)
	h.Spin()
}
