package webquic

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/devsisters/goquic"
)

func TestDial(t *testing.T) {
	t.Run(`success`, func(t *testing.T) {

		// Data
		type payload struct {
			Message string
		}
		expected := payload{Message: `alive`}

		// Server
		s := Service{}
		// goquic.SetLogLevel(-1)
		http.HandleFunc(`/`, func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set(`Content-Type`, `application/json`)
			w.WriteHeader(http.StatusOK)

			raw, _ := json.Marshal(expected)
			io.WriteString(w, string(raw))
		})
		err := s.Dial(Config{
			Address:    `0.0.0.0:8080`,
			CertKey:    `bin/key.pem`,
			CertPem:    `bin/cert.pem`,
			Dispatcher: 1,
		})
		if err != nil {
			t.Error(err)
		}

		// Client
		client := &http.Client{
			Transport: goquic.NewRoundTripper(true),
		}
		r, err := client.Get(`http://0.0.0.0:8080/`)
		if err != nil {
			t.Error(err)
		}

		// Results
		raw, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			t.Error(err)
		}
		var actual payload
		if err := json.Unmarshal(raw, &actual); err != nil {
			t.Error(err)
		}
		if actual.Message != expected.Message {
			t.Error(fmt.Errorf("expected=%s|actual=%s", expected.Message, actual.Message))
		}
	})
}
