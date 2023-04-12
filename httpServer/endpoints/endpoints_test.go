package endpoints

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestArticlesCategoryHandler(t *testing.T) {
	tt := []struct {
		routeVariable string
		shouldPass    bool
	}{
		{"system/ip", true},
		{"system/sdfsdf", true},
		{"cat", false},
		{"system/id:", true},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/api/%s", tc.routeVariable)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/api/{system}/{config}", ApiHandler)
		router.ServeHTTP(rr, req)

		if rr.Code == http.StatusOK && !tc.shouldPass {
			t.Errorf("handler should have failed on routeVariable %s: got %v want %v shouldPass: %v",
				tc.routeVariable, rr.Code, http.StatusOK, tc.shouldPass)
		}
	}
}
