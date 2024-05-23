package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/documents/resiliency"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

func resiliencyExchange(r *http.Request) (*http.Response, *core.Status) {
	switch r.Method {
	case http.MethodGet:
		return get(r.Context(), r.Header, r.URL.Query())
	case http.MethodPut:
		return nil, nil
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
}

func get(ctx context.Context, h http.Header, values url.Values) (resp *http.Response, status *core.Status) {
	var docs any

	docs, status = resiliency.Get(ctx, h, values)
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
	status = resiliency.Put[*http.Request](r.Context(), r.Header, r)
	return httpx.NewResponseWithStatus(status, status.Err)
}