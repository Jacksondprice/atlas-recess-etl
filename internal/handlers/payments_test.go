package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListPayments(t *testing.T) {
	rec := httptest.NewRecorder()
	ListPayments(rec, httptest.NewRequest("GET", "/api/payments", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}
