package resiliency1

import (
	"fmt"
	"net/http"
	url2 "net/url"
)

func ExamplePut() {
	uri := "http://localhost:8081/github/advanced-go/documents:resiliency?region=*"
	req, _ := http.NewRequest(http.MethodPut, uri, nil)
	status := Put(req, testDocs)
	fmt.Printf("test: Put() -> [status:%v]\n", status)

	//Output:
	//test: Put() -> [status:OK]

}

func ExampleGet() {
	url, _ := url2.Parse("http://localhost:8081/github/advanced-go/documents:resiliency?region=*")
	entries, status := Get(nil, nil, url)
	fmt.Printf("test: Get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: Get() -> [status:OK] [entries:4]

}
