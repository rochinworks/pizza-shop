package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"go-web-api/src/handler/mocks"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	. "go-web-api/src/handler"
	"go-web-api/src/pg"
)

func TestPizzaHandler(t *testing.T) {
	// ARRANGE
	jsonStr := []byte(`
		{
			"style": "hawaiian",
			"userId": "5187924f-ac60-43a7-af0c-877f21026cce"
		}
	`)
	req, err := http.NewRequest("POST", "/order/start", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	// mock the repository for testing
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)

	rr := httptest.NewRecorder()
	h := NewHandler(mockRepo)
	handler := http.HandlerFunc(h.PizzaHandler())
	ctx := context.Background()
	id := uuid.New()

	uID, err := uuid.Parse("5187924f-ac60-43a7-af0c-877f21026cce")
	if err != nil {
		t.Error("uuid did not parse correctly")
	}

	pizza := pg.Pizza{
		Style:  "hawaiian",
		Status: "starting",
		UserID: uID,
	}

	// ACT
	mockRepo.EXPECT().StoreOrder(ctx, pizza).Return(&id, nil)

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
}
