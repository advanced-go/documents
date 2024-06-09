package resiliency1

import (
	"encoding/json"
	"fmt"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func ExampleGetDocuments() {
	values := make(url.Values)
	values.Add(core.RegionKey, "*")
	req := access.NewRequest(http.MethodGet, "", nil)
	//req,_ := http.NewRequest(http.MethodGet, "test", module.RouteName, nil, timeout)
	docs, h, status1 := getDocuments(nil, req, values)
	fmt.Printf("test: getDocuments(nil) -> [status:%v] [header:%v] [count:%v]\n", status1, h, len(docs))

	values = make(url.Values)
	values.Add(core.RegionKey, "reGion2")
	docs, h, status1 = getDocuments(nil, req, values)
	fmt.Printf("test: getDocuments(\"region2\") -> [status:%v] [header:%v] [count:%v]\n", status1, h, len(docs))

	//Output:
	//test: getDocuments(nil) -> [status:OK] [header:map[]] [count:4]
	//test: getDocuments("region2") -> [status:OK] [header:map[]] [count:2]

}

func ExampleAddDocuments() {
	//req := NewRequest(http.MethodPut, "test", module.RouteName, nil, timeout)
	req := access.NewRequest(http.MethodPut, "", nil)

	h, status := addDocuments(nil, req, storage)
	fmt.Printf("test: addDocuments() -> [status:%v] [header:%v] [count:%v]\n", status, h, len(storage))

	req = access.NewRequest(http.MethodGet, "test", nil)
	values := make(url.Values)
	values.Add(core.RegionKey, "reGion2")
	docs, h, status1 := getDocuments(nil, req, values)
	fmt.Printf("test: getDocuments(\"reGion2\") -> [status:%v] [header:%v] [count:%v]\n", status1, h, len(docs))

	//Output:
	//test: addDocuments() -> [status:OK] [header:map[]] [count:8]
	//test: getDocuments("reGion2") -> [status:OK] [header:map[]] [count:4]

}

func _ExampleOutput() {
	buff, err := json.Marshal(storage)
	fmt.Printf("%v\n", err)
	fmt.Printf("%v\n", string(buff))

	//Output:
	//fail
}
