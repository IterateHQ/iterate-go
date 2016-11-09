package iterate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const accessToken = "abc123"

type Results struct {
	Token string
}

func TestList(t *testing.T) {
	_, server := mockServer(Response{
		Results: []Survey{{Id: "123"}},
	})
	defer server.Close()

	client := New(accessToken)
	client.host = server.URL

	surveys, err := client.ListSurveys()
	if err != nil {
		t.Error(err)
	}

	if len(surveys) == 0 {
		t.Error("Should have listed surveys")
	}
}

func TestSend(t *testing.T) {
	requests, server := mockServer(Response{})
	defer server.Close()

	surveyId := "58223a83d5167e00010012d1"
	email := "art@vandelayindustries.com"
	firstName := "Art"
	lastName := "Vandelay"
	delay := 7 * 24 * time.Hour
	date := time.Now().Add(24 * time.Hour)

	client := New(accessToken)
	client.host = server.URL

	t.Run("with basic send parameters", func(t *testing.T) {
		client.EmailSurvey(surveyId, SendParams{Email: email})
		r := <-requests

		if r.FormValue("access_token") != accessToken {
			t.Error("Request should contain the api key")
		}

		if r.FormValue("v") != client.version {
			t.Error("Request should contain the api version")
		}

		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			t.Error("Request should have content-type application/x-www-form-urlencoded")
		}

		if !strings.Contains(r.RequestURI, surveyId) {
			t.Error("Request url should contain the survey id")
		}

		if r.FormValue("email") != email {
			t.Error("Request should have the email")
		}
	})

	t.Run("with delay", func(t *testing.T) {
		client.EmailSurvey(surveyId, SendParams{
			Email:     email,
			FirstName: firstName,
			LastName:  lastName,
			Delay:     delay,
		})
		r := <-requests

		if r.FormValue("email") != email {
			t.Error("Request should have the email")
		}

		if r.FormValue("first_name") != firstName {
			t.Error("Request should have the first name")
		}

		if r.FormValue("last_name") != lastName {
			t.Error("Request should have the last name")
		}

		if r.FormValue("delay") != "604800" {
			t.Error("Request should have the delay")
		}
	})

	t.Run("with date", func(t *testing.T) {
		client.EmailSurvey(surveyId, SendParams{
			Email: email,
			Date:  date,
		})
		r := <-requests

		if r.FormValue("email") != email {
			t.Error("Request should have the email")
		}

		if r.FormValue("date") != fmt.Sprintf("%v", date.Unix()) {
			t.Error("Request should have the date")
		}
	})
}

func mockServer(response interface{}) (chan *http.Request, *httptest.Server) {
	requests := make(chan *http.Request, 1)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		if response != nil {
			json, _ := json.Marshal(response)
			w.Header().Set("Content-Type", "application/json")
			w.Write(json)
		}

		requests <- r
	}))

	return requests, server
}
