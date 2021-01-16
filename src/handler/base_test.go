package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	. "go-web-api/src/handler"
	"go-web-api/src/handler/mocks"
)

func TestBaseHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// mock the repository for testing
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)

	rr := httptest.NewRecorder()
	h := NewHandler(mockRepo)
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
