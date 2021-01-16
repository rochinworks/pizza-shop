package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "go-web-api/src/handler"
	"go-web-api/src/handler/mocks"
	"go-web-api/src/pg"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestAddUser(t *testing.T) {

	// ARRANGE
	jsonStr := []byte(`{"username": "testing123"}`)
	req, err := http.NewRequest("POST", "/user/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	// mock the repository for testing
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)
	mockUser := pg.User{
		Name: "testing123",
	}

	rr := httptest.NewRecorder()
	h := NewHandler(mockRepo)
	handler := http.HandlerFunc(h.AddUserHandler())
	ctx := context.Background()
	id := uuid.New()

	// ACT
	mockRepo.EXPECT().Store(ctx, mockUser).Return(&id, nil)

	handler.ServeHTTP(rr, req)

	// ASSERT
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v\n", status, http.StatusOK)
	}

	var actual map[string]interface{}
	resp := rr.Result()
	readableResp, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = json.Unmarshal(readableResp, &actual)
	if err != nil {
		t.Error("json did not unmarshal correctly")
	}

	if actual["id"] == id {
		t.Error("no id in response")
	}
}
