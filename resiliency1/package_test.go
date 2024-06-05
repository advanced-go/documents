package resiliency1

import (
	"fmt"
	"net/http"
	url2 "net/url"
)

func ExamplePut() {
	uri := "http://localhost:8081/github/advanced-go/documents:resiliency"
	req, _ := http.NewRequest(http.MethodPut, uri, nil)
	h, status := Put(req, testDocs)
	fmt.Printf("test: Put() -> [status:%v] [header:%v]\n", status, h)

	//Output:
	//test: Put() -> [status:OK] [header:map[]]

}

func ExampleGet() {
	url, _ := url2.Parse("http://localhost:8081/github/advanced-go/documents:resiliency?region=*")
	entries, h, status := Get(nil, nil, url.Query())
	fmt.Printf("test: Get() -> [status:%v] [header:%v] [entries:%v]\n", status, h, len(entries))

	//Output:
	//test: Get() -> [status:OK] [header:map[]] [entries:4]

}
