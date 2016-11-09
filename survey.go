package iterate

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Survey represents a survey.
type Survey struct {
	Id        string        `json:"id,omitempty"`
	Name      string        `json:"id,omitempty"`
	Questions []interface{} `json:"questions,omitempty"`
}

// SendParams is the set of parameters that can be used when
// sending an email survey.
// For more details see: http://docs.iterate.apiary.io/#reference/0/surveysidsend/post
type SendParams struct {
	Email     string
	FirstName string
	LastName  string
	Delay     time.Duration
	Date      time.Time
}

// ListSurveys returns a list of your surveys.
// For more details see: http://docs.iterate.apiary.io/#reference/0/surveys/get
func (c Client) ListSurveys() (surveys []Survey, err error) {
	results, err := c.get("/surveys", url.Values{})

	if err != nil {
		return
	}

	err = json.Unmarshal(results, &surveys)
	return
}

// EmailSurvey emails a survey to the specified email address with
// optional additional send parameters.
// For more details see: http://docs.iterate.apiary.io/#reference/0/surveysidsend/post
func (c Client) EmailSurvey(surveyId string, params SendParams) error {
	path := fmt.Sprintf("/surveys/%s/send", surveyId)

	values := url.Values{}
	if params.Email != "" {
		values.Add("email", params.Email)
	}
	if params.FirstName != "" {
		values.Add("first_name", params.FirstName)
	}
	if params.LastName != "" {
		values.Add("last_name", params.LastName)
	}

	delay := int(params.Delay.Seconds())
	if delay > 0 {
		values.Add("delay", fmt.Sprintf("%v", delay))
	}

	date := params.Date.Unix()
	if date > 0 {
		values.Add("date", fmt.Sprintf("%v", date))
	}

	_, err := c.post(path, values)
	return err
}
