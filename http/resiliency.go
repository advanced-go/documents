package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/documents/resiliency1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

func resiliencyExchange(r *http.Request, url *url.URL) (*http.Response, *core.Status) {
	switch r.Method {
	case http.MethodGet:
		return get(r.Context(), r.Header, r.URL)
	case http.MethodPut:
		return nil, nil
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
}

func get(ctx context.Context, h http.Header, url *url.URL) (resp *http.Response, status *core.Status) {
	var docs any

	docs, status = resiliency1.Get(ctx, h, url)
	if !status.OK() {
		return httpx.NewResponseWithStatus(status, status.Err)
	}
	resp, status = httpx.NewJsonResponse(docs, nil)
	if !status.OK() {
		var e core.Log
		e.Handle(status, core.RequestId(h))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
	return
}

func put(r *http.Request) (resp *http.Response, status *core.Status) {
	status = resiliency1.Put(r, nil)
	return httpx.NewResponseWithStatus(status, status.Err)
}
