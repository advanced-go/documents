package resiliency

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/json"
	"net/http"
)

// Put - resource PUT
func Put[T PutBodyConstraints](ctx context.Context, h http.Header, body T) *core.Status {
	if body == nil {
		return core.NewStatus(http.StatusBadRequest)
	}
	return put[core.Log](ctx, core.AddRequestId(h), body)
}

func put[E core.ErrorHandler](ctx context.Context, h http.Header, body any) *core.Status {
	var e E

	docs, status := json.New[[]Document](body, nil)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return status
	}
	status = addDocuments(ctx, h, docs)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return status
	}
	return status
}
