package main

import (
	"testing"
	"time"

	"github.com/researchlab/crawler/engine"
	"github.com/researchlab/crawler/model"
	"github.com/researchlab/crawler_distributed/config"
	"github.com/researchlab/crawler_distributed/rpcsupport"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	//start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)
	//start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		t.Fatal(err)
	}

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/91162834",
		Type: "zhenai",
		Id:   "91162834",
		Payload: model.Profile{
			Name:      "蜗牛漫步",
			Gender:    "男士",
			Age:       29,
			Height:    175,
			Weight:    59,
			Income:    "5-8千",
			Marriage:  "未婚",
			Education: "中专",
			Hokou:     "四川绵阳",
			Xinzuo:    "魔羯座(12.22-01.19)",
		},
	}
	result := ""
	//call save
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil {
		t.Error(err)
	}
	if err != nil || result != "ok" {
		t.Errorf("result: %v, err: %v", result, err)
	}
}
