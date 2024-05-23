package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/io"
	"github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

var storage []Document

func getDocuments(_ context.Context, h http.Header, values url.Values) (docs []Document, status *core.Status) {
	var buf []byte

	list := storage
	if h != nil {
		location := h.Get(httpx.ContentLocation)
		if location != "" {
			buf, status = io.ReadFile(location)
			if !status.OK() {
				return nil, status
			}
			if len(buf) == 0 {
				return nil, core.StatusNotFound()
			}
			list, status = json.New[[]Document](buf, nil)
			if !status.OK() {
				return nil, status
			}

		}
	}
	if len(list) == 0 {
		return nil, core.StatusNotFound()
	}
	filter := core.NewOrigin(values)
	for _, target := range list {
		if core.OriginMatch(target.Origin, filter) {
			docs = append(docs, target)
		}
	}
	if len(docs) == 0 {
		return nil, core.StatusNotFound()
	}
	return docs, core.StatusOK()
}

func addDocuments(_ context.Context, _ http.Header, docs []Document) *core.Status {
	if len(docs) > 0 {
		storage = append(storage, docs...)
	}
	return core.StatusOK()
}
