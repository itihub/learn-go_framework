package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

func main() {

	// 初始化es连接
	host := "http://local.docker.node1.com:9200"
	/*
		这里必须将sniff设置为false。因为使用olivere/elastic连接elasticsearch时，发现连接地址明明输入的时网络地址，
		但是连接时会自动转换成内网地址或者docker中的ip地址，导致服务连接不上。
	*/
	logger := log.New(os.Stdout, "shop", log.LstdFlags)
	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	//search(client)
	//addIndex(client)
	createIndex(client)

}

func search(client *elastic.Client) {
	// 构建http请求报文
	q := elastic.NewMatchQuery("address", "street")
	// 打印生成的请求报文
	//src, err := q.Source()
	//if err != nil {
	//	panic(err)
	//}
	//data, err := json.Marshal(src)
	//got := string(data)
	//fmt.Println(got)

	// 发起http请求
	result, err := client.Search().Index("user").Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}
	total := result.Hits.TotalHits.Value
	fmt.Printf("搜索结果数量:%d\n", total)
	for _, value := range result.Hits.Hits {

		// 将es中的对象转换为struct类型
		account := Account{}
		_ = json.Unmarshal(value.Source, &account)
		fmt.Println(account)

		// 使用json打印es中文档
		if jsonData, err := value.Source.MarshalJSON(); err == nil {
			fmt.Println(string(jsonData))
		} else {
			panic(err)
		}
	}
}

func addIndex(client *elastic.Client) {
	// 添加数据到es
	account1 := Account{AccountNumber: 15468, Firstname: "Jimmy", Lastname: "Ji"}
	put, err := client.Index().
		Index("user").
		BodyJson(account1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed account %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}

func createIndex(client *elastic.Client) {
	// 添加索引

	const (
		goodsMapping = `
			{
			  "settings": {
				"number_of_shards": 1,
				"number_of_replicas": 0
			  },
			  "mappings": {
				  "properties": {
					"name": {
					  "type": "text",
					  "analyzer": "ik_max_word"
					},
					"id": {
					  "type": "integer"
					}
   				  }
			  }
			}
		`
	)

	createResult, err := client.CreateIndex("mygoods").BodyString(goodsMapping).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(createResult.Acknowledged)

}

type Account struct {
	AccountNumber int32  `json:"account_number"`
	Balance       int32  `json:"balance"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
}
