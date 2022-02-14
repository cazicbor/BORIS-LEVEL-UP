package main

import (
	"errors"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Tests struct {
	name          string
	server        *httptest.Server
	response      *Task
	expectedError error
}

func TestGetHandler(t *testing.T) {
	tasks := index()

	/* tests := []Tests{
		{
			name: "req-1",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"id" : 6, "description" : "Finir partie tech", "Deadline" : "11/02/2022", "Status" : "Ongoing"}`))
			})),
			response: &Task{
				ID:          "6",
				Description: "Finir partie tech",
				Deadline:    "11/02/2022",
				Status:      "Ongoing",
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()

			resp, err := index()

			if !reflect.DeepEqual(resp, test.response) {
				t.Errorf("FAILED : expected %v, got %v", test.response, resp)
			}
			if !errors.Is(err, test.expectedError) {
				t.Errorf("FAILED : expected %v, got %v", test.expectedError, err)
			}
		})
	} */

}

func TestPostHandler(t *testing.T) {

}

func PutHandler(t *testing.T) {

}

func DeleteHandler(t *testing.T) {

}
