package main

import (
	"fmt"
	"myblob/pkg/setting"
	"myblob/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run(fmt.Sprintf(":%d", setting.Port))
}