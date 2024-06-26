package http

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func Test_get(t *testing.T) {
	type args struct {
		ctx    context.Context
		h      http.Header
		values url.Values
	}
	tests := []struct {
		name       string
		args       args
		wantResp   *http.Response
		wantStatus *core.Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, gotStatus := get[core.Output](tt.args.ctx, tt.args.h, nil, "")
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("get() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if !reflect.DeepEqual(gotStatus, tt.wantStatus) {
				t.Errorf("get() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func Test_put(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name       string
		args       args
		wantResp   *http.Response
		wantStatus *core.Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, gotStatus := put[core.Output](tt.args.r, "")
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("put() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if !reflect.DeepEqual(gotStatus, tt.wantStatus) {
				t.Errorf("put() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func Test_resiliencyExchange(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name  string
		args  args
		want  *http.Response
		want1 *core.Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := resiliencyExchange[core.Output](tt.args.r, nil)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resiliencyExchange() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("resiliencyExchange() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
