package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"reflect"
	"time"
)

var client *elastic.Client

var host = "http://elk.test.qcds.com:9233"

//type Employee struct {
//	Mechanic_name string   `json:"mechanic_name"`
//}

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

type SearchRecord struct {
	CreatedAt int    `json:"created_at"`
	IsShow    int    `json:"is_show"`
	Keyword   string `json:"keyword"`
	Num       int    `json:"num"`
	UpdatedAt int    `json:"updated_at"`
}

//初始化
func init() {
	//errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error

	//这个地方有个小坑 不加上elastic.SetSniff(false) 会连接不上
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host), elastic.SetBasicAuth("elastic", "123456"))
	if err != nil {
		panic(err)
	}
	_, _, err = client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	_, err = client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Elasticsearch version %s\n", esversion)
}

//创建
func create() {

	fmt.Println(time.Now().Unix())
	return
	//使用结构体
	//e1 := Employee{"Jane", "Smith", 32, "I like to collect rock albums", []string{"music"}}
	e1 := SearchRecord{1111, 1, "关键字测试", 1, 11}

	put1, err := client.Index().
		Index("search_record").
		Type("questions").
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	return
	//使用字符串
	e2 := `{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	put2, err := client.Index().
		Index("megacorp").
		Type("employee").
		Id("2").
		BodyJson(e2).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put2.Id, put2.Index, put2.Type)

	//e3 := `{"first_name":"Douglas","last_name":"Fir","age":35,"about":"I like to build cabinets","interests":["forestry"]}`
	//put3, err := client.Index().
	//	Index("megacorp").
	//	Type("employee").
	//	Id("3").
	//	BodyJson(e3).
	//	Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put3.Id, put3.Index, put3.Type)
}

//查找
func gets() {
	//通过id查找
	get1, err := client.Get().Index("megacorp").Type("employee").Id("1").Do(context.Background())
	if err != nil {
		panic(err)
	}

	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
		var bb Employee
		data, _ := get1.Source.MarshalJSON()
		err := json.Unmarshal(data, &bb)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(bb)
		//fmt.Println(string(get1.Source))
	}
}

func main() {
	//create()
	//gets()
	//query()
	update()
}

func update() {

	doc := map[string]interface{}{
		"num": "宝马发动机异响",
	}
	res, err := client.Update().
		Index("search_record").Type("questions").Id("_ACOJn0BncnMHLjopPTL").
		Doc(doc).
		FetchSource(true).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

////搜索
func query() {

	//var res *elastic.SearchResult
	//var err error
	//path,err := client.IndexAnalyze().Analyzer("ik_max_word").Text("宝马发动机异响").Do(context.TODO())
	//
	//for _,v := range path.Tokens {
	//	fmt.Println(v.Token)
	//}
	//
	//return
	////取所有
	//res, err = client.Search("megacorp").Type("employee").Do(context.Background())
	//printEmployee(res, err)
	keyword := "宝马发动机异响"
	qy := elastic.NewBoolQuery()
	qy.Should(elastic.NewMatchQuery("brand_name.pinyin", keyword))
	qy.Should(elastic.NewMatchQuery("series_name.pinyin", keyword))
	qy.Should(elastic.NewMatchQuery("mechanic_name", keyword))
	qy.Should(elastic.NewMatchQuery("model_name", keyword))
	qy.Should(elastic.NewMatchQuery("content.pinyin", keyword).Analyzer("ik_max_word"))
	qy.Should(elastic.NewMatchQuery("summary", keyword))
	qy.Must(elastic.NewRangeQuery("num_useful_sub").Gte("-5"))
	qy.MinimumNumberShouldMatch(1)

	q := elastic.NewFunctionScoreQuery().
		Query(qy).
		Add(elastic.NewMatchQuery("content.pinyin", keyword).MinimumShouldMatch("100%"), elastic.NewWeightFactorFunction(10)).
		Add(elastic.NewMatchQuery("summary", keyword).MinimumShouldMatch("100%"), elastic.NewWeightFactorFunction(9)).
		Add(elastic.NewMatchQuery("keywords", keyword), elastic.NewWeightFactorFunction(9)).
		Add(elastic.NewMatchQuery("mechanic_name", keyword), elastic.NewWeightFactorFunction(4)).
		Add(elastic.NewMatchQuery("series_name.pinyin", keyword), elastic.NewWeightFactorFunction(3)).
		AddScoreFunc(elastic.NewScriptFunction(elastic.NewScript("0.01 * doc['num_useful'].value"))).
		BoostMode("sum").
		ScoreMode("sum")

	src, err := q.Source()
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	res, err := client.Search("qcds_question").
		Type("questions").
		From(0).
		Size(10).
		FetchSourceContext(elastic.NewFetchSourceContext(true).Include("content", "mechanic_name", "id")).
		Query(q).
		Sort("id", false).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	printEmployee(res, err)

	////条件查询
	////年龄大于30岁的
	//boolQ := elastic.NewBoolQuery()
	//boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	//boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	//res, err = client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
	//printEmployee(res, err)
	//
	////短语搜索 搜索about字段中有 rock climbing
	//matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
	//res, err = client.Search("megacorp").Type("employee").Query(matchPhraseQuery).Do(context.Background())
	//printEmployee(res, err)
	//
	////分析 interests
	//aggs := elastic.NewTermsAggregation().Field("interests")
	//res, err = client.Search("megacorp").Type("employee").Aggregation("all_interests", aggs).Do(context.Background())
	//printEmployee(res, err)

}

//
////简单分页
func list(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := client.Search("megacorp").
		Type("employee").
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	printEmployee(res, err)

}

//
//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) []interface{} {
	if err != nil {
		print(err.Error())
	}
	//var typ = make(map[string]interface{})
	var typ = make(map[string]interface{})
	//var typ Employee

	//var ress []Employee
	var ress []interface{}
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		//t := item.(Employee)
		ccc, _ := json.Marshal(item)
		fmt.Println(string(ccc))
		ress = append(ress, item)
		//t := item.(map[string]interface{})
		//ts["aa"] = t
	}
	//fmt.Printf("%#v\n", ress)
	fmt.Println("------")
	ccc, _ := json.Marshal(ress)
	fmt.Println(string(ccc))

	return ress
}
