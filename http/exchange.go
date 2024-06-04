package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/documents/module"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
	"time"
)

// https://localhost:8081/github/advanced-go/documents:v1/resiliency?reg=region1&az=zone1&host=www.google.com

const (
	resiliencyPath = "resiliency1"
)

var authorityResponse = httpx.NewAuthorityResponse(module.Authority)

// Controllers - egress controllers
func Controllers() []*controller.Controller {
	return []*controller.Controller{
		controller.NewController("google-search", controller.NewPrimaryResource("www.google.com", "", time.Second*2, "", nil), nil),
	}
}

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
		return resiliencyExchange(r)
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
