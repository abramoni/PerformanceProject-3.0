package mypackages

import (
	"context"
	"encoding/json"
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

func GetESClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}

func StoreInElasticDb(body []byte, IndexName string) {
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing elastcidb client : ", err)
		panic("Client fail ")
	}

	jsondata_string := string(body)
	ind, err := esclient.Index().
		Index(IndexName).
		BodyJson(jsondata_string).
		Do(ctx)
	_ = ind

	if err != nil {
		panic(err)
	}

}

func RetriveFronElasticDb(time_stamp int64) {
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("timeStamp", time_stamp))
	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}
	fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))

	searchService := esclient.Search().Index("longevitycluster").SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		
		err := json.Unmarshal(hit.Source, &Response_for_elastic)
		if err != nil {
			fmt.Println("Searching Result Err = ", err)
		}

	}

	if err != nil {
		fmt.Println("Fetching from elastic DB failed: ", err)
	} else {
		fmt.Println("id = ", Response_for_elastic.Id)
		fmt.Println("Timestamp = ", Response_for_elastic.TimeStamp)
		fmt.Println("fileCreateRate = ", Response_for_elastic.FileCreateRate)
		fmt.Println("utilizationChangeRate = ", Response_for_elastic.UtilizationChangeRate)
		fmt.Println("utilizationChangeRateUnit = ",Response_for_elastic.UtilizationChangeRateUnit)
		fmt.Println("garbageCollection = ",Response_for_elastic.GarbageCollection)
		fmt.Println("garbageCollectionUnit = ",Response_for_elastic.GarbageCollectionUnit)
		fmt.Println( "clusterUsage = ",Response_for_elastic.ClusterUsage)
		fmt.Println("clusterUsageUnits = ",Response_for_elastic.ClusterUsageUnits)
		fmt.Println( "metaDataUtilization = ",Response_for_elastic.MetaDataUtilization)
		fmt.Println("metaDataUtilizationUnits = ",Response_for_elastic.MetaDataUtilizationUnits)
		fmt.Println( "meta_data_space_percentage = ",Response_for_elastic.MetaDataSpacePercentage)
		}
	}