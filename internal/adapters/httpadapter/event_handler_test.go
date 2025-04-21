package httpadapter_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"event-manager-go/internal/adapters/httpadapter"

	"github.com/gorilla/mux"
)

func TestCreateEventHandler_InvalidJSON(t *testing.T) {
	// Invalid JSON body should return 400.
	req := httptest.NewRequest(http.MethodPost, "/events", bytes.NewBufferString("invalid-json"))
	w := httptest.NewRecorder()

	httpadapter.CreateEventHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", resp.StatusCode)
	}
}

func TestCreateEventHandler_DBConnectionError(t *testing.T) {
	// Set valid JSON that passes DTO decoding,
	// then force a repository connection error by clearing env values.
	// (Assumes that empty env variables cause an error on NewMongoEventRepository.)
	payload := map[string]interface{}{
		"Name":                     "Test Event",
		"Description":              "A test event",
		"Address":                  "123 test st",
		"MapUrl":                   "http://maps.example.com",
		"Date":                     "2023-10-10",
		"Modality":                 "online",
		"CancellationPolicy":       "none",
		"ParticipantEditionPolicy": "open",
		"TicketType":               "free",
		"TicketPrice":              0,
		"TicketQuantity":           100,
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/events", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	// Clear env so that NewMongoEventRepository fails.
	os.Setenv("MONGO_URI", "")
	os.Setenv("MONGO_DB", "")
	os.Setenv("MONGO_COLLECTION", "")

	httpadapter.CreateEventHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", resp.StatusCode)
	}
}

func TestUpdateEventHandler_MissingID(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/events", bytes.NewBufferString("{}"))
	w := httptest.NewRecorder()

	httpadapter.UpdateEventHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 for missing ID, got %d", resp.StatusCode)
	}
}

func TestUpdateEventHandler_InvalidJSON(t *testing.T) {
	// Provide an ID using mux vars but an invalid JSON body.
	req := httptest.NewRequest(http.MethodPut, "/events/some-id", bytes.NewBufferString("invalid-json"))
	// Set mux vars.
	req = mux.SetURLVars(req, map[string]string{"id": "some-id"})
	w := httptest.NewRecorder()

	httpadapter.UpdateEventHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 for invalid JSON, got %d", resp.StatusCode)
	}
}

func TestFindAllEventHandler_DBConnectionError(t *testing.T) {
	// For FindAllEventHandler, if repository connection fails,
	// the handler should return 500.
	req := httptest.NewRequest(http.MethodGet, "/events", nil)
	w := httptest.NewRecorder()

	// Clear environment so repository connection returns error.
	os.Setenv("MONGO_URI", "")
	os.Setenv("MONGO_DB", "")
	os.Setenv("MONGO_COLLECTION", "")

	httpadapter.FindAllEventHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status 500 for DB connection error, got %d", resp.StatusCode)
	}
}

func TestGetEventHandler_MissingID(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/events", nil)
	w := httptest.NewRecorder()

	httpadapter.GetEventHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 for missing ID, got %d", resp.StatusCode)
	}
}

func TestDeleteEventHandler_MissingID(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/events", nil)
	w := httptest.NewRecorder()

	httpadapter.DeleteEventHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 for missing ID, got %d", resp.StatusCode)
	}
}
