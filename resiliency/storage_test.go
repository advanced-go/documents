package resiliency

import (
	"encoding/json"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

var testDocs = []Document{
	{Origin: core.Origin{Region: "region1", Zone: "Zone1", Host: "www.host1.com"}, Version: "v2", Content: "[]"},
	{Origin: core.Origin{Region: "region1", Zone: "Zone2", Host: "www.host1.com"}, Version: "v2", Content: "{ \"status\" : \"active\" }"},
	{Origin: core.Origin{Region: "region2", Zone: "Zone1", Host: "www.google.com"}, Version: "v2", Content: "{ \"status\" : \"inactive\" }"},
	{Origin: core.Origin{Region: "region2", Zone: "Zone1", Host: "www.yahoo.com"}, Version: "v2", Content: "{ \"status\" : \"removed\" }"},
}

func ExampleAddDocuments() {

	status := addDocuments(nil, nil, testDocs)
	fmt.Printf("test: addDocuments() -> [status:%v] [count:%v]\n", status, len(storage))

	docs, status1 := getDocuments(nil, nil, nil)
	fmt.Printf("test: getDocuments(nil) -> [status:%v] [count:%v]\n", status1, len(docs))

	values := make(url.Values)
	values.Add(core.RegionKey, "reGion2")
	docs, status1 = getDocuments(nil, nil, values)
	fmt.Printf("test: getDocuments(\"region2\") -> [status:%v] [count:%v]\n", status1, len(docs))

	//Output:
	//test: addDocuments() -> [status:OK] [count:4]
	//test: getDocuments(nil) -> [status:OK] [count:4]
	//test: getDocuments("region2") -> [status:OK] [count:2]

}

func _ExampleOutput() {
	buff, err := json.Marshal(testDocs)
	fmt.Printf("%v\n", err)
	fmt.Printf("%v\n", string(buff))

	//Output:
	//fail
}
