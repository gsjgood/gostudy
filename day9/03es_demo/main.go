package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
)

// Elasticsearch demo

type Question struct {
	category_name string `json:"category_name"`
	mechanic_name string `json:"mechanic_name"`
}

func main() {

	client, err := elastic.NewClient(elastic.SetURL("http://elk.test.qcds.com:9233"), elastic.SetBasicAuth("elastic", "123456"), elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")

	//查询test

	boolSearch := elastic.NewBoolQuery()
	boolSearch.Should(elastic.NewMatchQuery("mechanic_name", "秦"))

	//temp := client.Get().Index("qcds_question").Id("100030540")
	//get, err := temp.Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//if get.Found {
	//	fmt.Printf("Got document %s in version %d from index %s, type %s\n", get.Id, get.Version, get.Index, get.Type)
	//}
	//source, err := get.Source.MarshalJSON()
	//if err != nil {
	//	fmt.Printf("byte convert string failed, err: %v", err)
	//}
	//var question Question
	//if err := json.Unmarshal(source, &question); err == nil {
	//	fmt.Printf("data: %v\n", question)
	//}

	q, _ := boolSearch.Source()
	PrintQuery(q)
	result, err := client.Search().
		Index("qcds_question").
		Type("questions").
		Query(boolSearch).
		//From(0).Size(10).  //查询起始位和查询个数
		//Pretty(true).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

//自定义打印函数
func PrintQuery(src interface{}) {
	fmt.Println("*****")
	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
