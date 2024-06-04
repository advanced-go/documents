package http

import (
	"fmt"
	"github.com/advanced-go/documents/resiliency1"
	"github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	url2 "net/url"
)

func ExampleExtract() {
	s := "http://localhost:8080/github/advanced-go/documents:v1/resiliency/google-search?region=*"
	url, _ := url2.Parse(s)
	p := uri.Uproot(s)

	u2 := extract(url, p)
	fmt.Printf("test: extract(\"%v\") -> [url:%v] [path:%v] [query:%v] [resource:%v]\n", s, u2, u2.Path, u2.RawQuery, p.Resource)

	//Output:
	//fail
}

func ExampleExchange_Resiliency() {
	uri := "http://localhost:8081/github/advanced-go/documents:resiliency?region=region1"
	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	resp, status := Exchange(req)
	if !status.OK() {
		fmt.Printf("test: Exchange() -> [status:%v]\n", status)
	} else {
		entries, status1 := json.New[[]resiliency1.Entry](resp.Body, resp.Header)
		fmt.Printf("test: Exchange() -> [status:%v] [status-code:%v] [bytes:%v] [content:%v]\n", status1, resp.StatusCode, resp.ContentLength, entries)
	}

	//Output:
	//test: Exchange() -> [status:Not Found]

}
