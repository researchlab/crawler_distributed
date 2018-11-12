package main

import (
	"fmt"

	"github.com/researchlab/crawler/engine"
	"github.com/researchlab/crawler/scheduler"
	"github.com/researchlab/crawler/zhenai/parser"
	"github.com/researchlab/crawler_distributed/config"
	"github.com/researchlab/crawler_distributed/persist/client"
)

func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	e := (&engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: config.WorkerCount,
		ItemChan:    itemChan,
	})
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
