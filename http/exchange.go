package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/documents/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	url2 "net/url"
	"strings"
)

// https://localhost:8081/github/advanced-go/documents:v1/resiliency?reg=region1&az=zone1&host=www.google.com

const (
	resiliencyPath = "resiliency"
)

var authorityResponse = httpx.NewAuthorityResponse(module.Authority)

// Exchange - HTTP exchange
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	p, status := httpx.ValidateURL(r.URL, module.Authority)
	if !status.OK() {
		return httpx.NewResponseWithStatus(status, status.Err)
	}
	r.Header.Add(core.XVersion, p.Version)
	core.AddRequestId(r.Header)
	switch strings.ToLower(p.Resource) {
	case resiliencyPath:
		return resiliencyExchange(r, extract(r.URL, p))
	case core.VersionPath:
		return httpx.NewVersionResponse(module.Version), core.StatusOK()
	case core.AuthorityPath:
		return authorityResponse, core.StatusOK()
	case core.HealthReadinessPath, core.HealthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status = core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", p.Resource)))
		return httpx.NewResponseWithStatus(status, status.Err)
	}
}

func extract(u *url2.URL, p *uri.Parsed) *url2.URL {
	if u == nil || p == nil {
		return u
	}
	raw := p.Path
	if u.RawQuery != "" {
		raw += "?" + u.RawQuery
	}
	newUrl, _ := url2.Parse(raw)
	return newUrl
}
