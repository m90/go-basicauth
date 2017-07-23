package basicauth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWith(t *testing.T) {
	tests := []struct {
		name               string
		auth               Credentials
		handlerFunc        http.HandlerFunc
		user               string
		pass               string
		expectedStatusCode int
	}{
		{
			"authorized",
			Credentials{"user", "tops3cr3tp455w0rd"},
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello Gophers!"))
			},
			"user",
			"tops3cr3tp455w0rd",
			200,
		},
		{
			"disabled",
			Credentials{"user", ""},
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello Gophers!"))
			},
			"foo",
			"bar",
			200,
		},
		{
			"incorrect user",
			Credentials{"user", "tops3cr3tp455w0rd"},
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello Gophers!"))
			},
			"zuck",
			"tops3cr3tp455w0rd",
			401,
		},
		{
			"incorrect password",
			Credentials{"user", "tops3cr3tp455w0rd"},
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello Gophers!"))
			},
			"user",
			"password123",
			401,
		},
		{
			"propagation",
			Credentials{"user", "tops3cr3tp455w0rd"},
			func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "OH NOES", http.StatusInternalServerError)
			},
			"user",
			"tops3cr3tp455w0rd",
			500,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wrapped := With(test.auth)(test.handlerFunc)

			ts := httptest.NewServer(wrapped)
			defer ts.Close()

			req, _ := http.NewRequest(http.MethodGet, ts.URL, nil)
			req.SetBasicAuth(test.user, test.pass)
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("Got error %v", err)
			}
			if res.StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %v response, got %v", test.expectedStatusCode, res.StatusCode)
			}
		})
	}
}
