package handler_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	. "go-web-api/src/handler"
	"go-web-api/src/handler/mocks"
)

func TestStatusHandler(t *testing.T) {
	// ARRANGE
	uID, err := uuid.Parse("5187924f-ac60-43a7-af0c-877f21026cce")
	if err != nil {
		t.Fatal(err)
	}
	oID, err := uuid.Parse("4187924f-ac60-43a7-af0c-877f21026cce")
	if err != nil {
		t.Fatal(err)
	}
	url := fmt.Sprintf("/order/status?userId=%s&orderId=%s", uID, oID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	// mock the repository for testing
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)

	rr := httptest.NewRecorder()
	h := NewHandler(mockRepo)
	handler := http.HandlerFunc(h.StatusHandler())
	ctx := context.Background()

	// ACT
	mockRepo.EXPECT().GetStatus(ctx, uID, oID).Return("starting", nil)

	handler.ServeHTTP(rr, req)

	// ASSERT
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}

	// read the response
	var actual map[string]interface{}
	resp := rr.Result()
	readableResp, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = json.Unmarshal(readableResp, &actual)
	if err != nil {
		t.Error("json did not unmarshal correctly")
	}

	if actual == nil {
		t.Error("no response")
	}

	if actual["status"] == "" {
		t.Error("status should not be nil")
	}
}
