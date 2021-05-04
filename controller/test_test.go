package controller

import (
	"testing"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/golang/mock/gomock"

	"github.com/zeal-haven/message-publisher/mock"
)

func TestController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		path = "/test"
		mockWriter = mock.NewMockResponseWriter(ctrl)
	)

	r := chi.NewRouter()
	r.Get(path, Test)

	t.Run("test page returns test text", func(t *testing.T) {
		mockWriter.EXPECT().WriteHeader(http.StatusOK)
		mockWriter.EXPECT().Write([]byte("Test")).Return(0, nil)

		request, _ := http.NewRequest("GET", path, nil)
		r.ServeHTTP(mockWriter, request)
	})
}