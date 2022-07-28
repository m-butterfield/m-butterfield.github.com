package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCSound(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/csound", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	testRouter().ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
}
