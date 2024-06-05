package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/documents/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

// https://localhost:8081/github/advanced-go/documents:v1/resiliency?reg=region1&az=zone1&host=www.google.com

const (
	resiliencyPath = "resiliency"
)

var authorityResponse = httpx.NewAuthorityResponse(module.Authority)

// Exchange - HTTP exchange
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if r == nil {
		return httpx.NewResponse[core.Log](http.StatusBadRequest, h2, nil)
	}
	p, status := httpx.ValidateURL(r.URL, module.Authority)
	if !status.OK() {
		return httpx.NewResponse[core.Log](status.HttpCode(), h2, status.Err)
	}
	r.Header.Add(core.XVersion, p.Version)
	core.AddRequestId(r.Header)
	switch p.Resource {
	case resiliencyPath:
		return resiliencyExchange[core.Log](r, p)
	case core.VersionPath:
		return httpx.NewVersionResponse(module.Version), core.StatusOK()
	case core.AuthorityPath:
		return authorityResponse, core.StatusOK()
	case core.HealthReadinessPath, core.HealthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status = core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", p.Resource)))
		return httpx.NewResponse[core.Log](status.HttpCode(), h2, status.Err)
	}
}
