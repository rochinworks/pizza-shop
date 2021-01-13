package handler_test

import (
	"encoding/json"
	. "go-web-api/src/handler"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBaseHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h := NewHandler()
	handler := http.HandlerFunc(h.BaseHandler())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}

	expected := map[string]interface{}{"alive": true}
	var actual map[string]interface{}
	err = json.Unmarshal([]byte(rr.Body.String()), &actual)
	if err != nil {
		t.Error("json did not unmarshal correctly")
	}

	if actual["alive"] != expected["alive"] {
		t.Errorf("handler returned unexpected body: got %v want %v\n", actual, expected)
	}
}
