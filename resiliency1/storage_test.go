package resiliency1

import (
	"encoding/json"
	"fmt"
	"github.com/advanced-go/documents/module"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

var testDocs = []Entry{
	{Region: "region1", Zone: "Zone1", Host: "www.host1.com", Status: "active", Timeout: 500, RateLimit: 100, RateBurst: 10},
	{Region: "region1", Zone: "Zone2", Host: "www.host1.com", Status: "inactive", Timeout: 1000, RateLimit: 100, RateBurst: 10},
	{Region: "region2", Zone: "Zone1", Host: "www.google.com", Status: "active", Timeout: 800, RateLimit: 100, RateBurst: 10},
	{Region: "region2", Zone: "Zone1", Host: "www.yahoo.com", Status: "active", Timeout: 2000, RateLimit: 100, RateBurst: 10},
}

func ExampleAddDocuments() {
	req := NewRequest(http.MethodPut, "test", module.RouteName, nil, timeout)
	status := addDocuments(nil, req, testDocs)
	fmt.Printf("test: addDocuments() -> [status:%v] [count:%v]\n", status, len(storage))

	req = NewRequest(http.MethodGet, "test", module.RouteName, nil, timeout)
	docs, status1 := getDocuments(nil, req, nil)
	fmt.Printf("test: getDocuments(nil) -> [status:%v] [count:%v]\n", status1, len(docs))

	values := make(url.Values)
	values.Add(core.RegionKey, "reGion2")
	docs, status1 = getDocuments(nil, req, values)
	fmt.Printf("test: getDocuments(\"region2\") -> [status:%v] [count:%v]\n", status1, len(docs))

	//Output:
	//test: addDocuments() -> [status:OK] [count:4]
	//test: getDocuments(nil) -> [status:OK] [count:0]
	//test: getDocuments("region2") -> [status:OK] [count:2]

}

func _ExampleOutput() {
	buff, err := json.Marshal(testDocs)
	fmt.Printf("%v\n", err)
	fmt.Printf("%v\n", string(buff))

	//Output:
	//fail
}
