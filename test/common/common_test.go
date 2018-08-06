package common_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/convee/golab/test/common"
)

func init() {
	common.Routes()
}

func TestSendJSON(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "sendjson", nil)
	if err != nil {
		t.Fatal("创建Request失败")
	}
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)

	t.Log("code:", rw.Code)
	t.Log("body:", rw.Body.String())
}
