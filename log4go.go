package main

import (
	"time"

	"git.corp.plu.cn/micro/utils"

	log "git.corp.plu.cn/plugo/log4go"
)

func main() {
	if err := utils.InitLog4goDefault(); err != nil {
		panic(err)
	}

	//log something
	log.Finest("Everything is created now (notice that I will not be printing to the file)")
	log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	log.Error("aa: %v", 111)
	log.Critical("Time to close out!")

	time.Sleep(50000000000)
}
