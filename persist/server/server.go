package main

import (
	"fmt"
	"log"

	"github.com/olivere/elastic"
	"github.com/researchlab/crawler_distributed/config"
	"github.com/researchlab/crawler_distributed/persist"
	"github.com/researchlab/crawler_distributed/rpcsupport"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
