package resiliency1

import (
	"context"
	"github.com/advanced-go/documents/module"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
	"time"
)

const (
	entriesJson = "file://[cwd]/resiliency1test/documents-v1.json"
)

var storage = []Entry{
	{Region: "region1", Zone: "Zone1", Host: "www.host1.com", Status: "active", Timeout: 100, RateLimit: 125, RateBurst: 25},
	{Region: "region1", Zone: "Zone2", Host: "www.host2.com", Status: "inactive", Timeout: 250, RateLimit: 100, RateBurst: 10},
	{Region: "region2", Zone: "Zone1", Host: "www.google.com", Status: "removed", Timeout: 3000, RateLimit: 50, RateBurst: 5},
	{Region: "region2", Zone: "Zone1", Host: "www.yahoo.com", Status: "removed", Timeout: 1500, RateLimit: 30, RateBurst: 2},
}

func getDocuments(_ context.Context, req access.Request, values url.Values) (docs []Entry, h2 http.Header, status *core.Status) {
	if len(storage) == 0 {
		return nil, nil, core.StatusNotFound()
	}
	var start = time.Now().UTC()

	filter := core.NewOrigin(values)
	for _, item := range storage {
		target := core.Origin{Region: item.Region, Zone: item.Zone, SubZone: item.SubZone, Host: item.Host}
		if core.OriginMatch(target, filter) {
			docs = append(docs, item)
		}
	}
	if len(docs) == 0 {
		status = core.StatusNotFound()
	} else {
		status = core.StatusOK()
	}
	access.LogEgress(start, time.Since(start), req, status, module.Authority, routeName, "", timeout, 0, 0, "")
	return docs, h2, status
}

func addDocuments(_ context.Context, req access.Request, docs []Entry) (http.Header, *core.Status) {
	var start = time.Now().UTC()

	if len(docs) > 0 {
		storage = append(storage, docs...)
	}
	access.LogEgress(start, time.Since(start), req, core.StatusOK(), module.Authority, routeName, "", timeout, 0, 0, "")
	return nil, core.StatusOK()
}

func setTimeout(ctx context.Context, duration time.Duration) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	}
	if _, ok := ctx.Deadline(); ok {
		return ctx, nil
	}
	return context.WithTimeout(ctx, duration)
}
