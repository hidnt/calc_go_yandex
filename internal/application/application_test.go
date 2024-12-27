package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testCase struct {
	name               string
	expression         string
	exceptedStatusCode int
	expectedAns        Answer
}

type CalcRequest struct {
	Expression string `json:"expression"`
}

func TestCalc(t *testing.T) {
	testCases := []testCase{
		{
			name:               "minus",
			expression:         "-10*(-5+2)",
			exceptedStatusCode: http.StatusOK,
			expectedAns:        Answer{Result: fmt.Sprintf("%f", float64(30))},
		},
		{
			name:               "priority",
			expression:         "(2+2-(-2+7)*2)/2",
			exceptedStatusCode: http.StatusOK,
			expectedAns:        Answer{Result: fmt.Sprintf("%f", float64(-3))},
		},
		{
			name:               "fractional numbers 2",
			expression:         "140.5*12.5/10",
			exceptedStatusCode: http.StatusOK,
			expectedAns:        Answer{Result: fmt.Sprintf("%f", float64(175.625))},
		},
		{
			name:               "incorrect expression 1",
			expression:         "1238)",
			exceptedStatusCode: http.StatusUnprocessableEntity,
			expectedAns:        Answer{Error: "Expression is not valid"},
		},
		{
			name:               "incorrect expression 2",
			expression:         "124+2-",
			exceptedStatusCode: http.StatusUnprocessableEntity,
			expectedAns:        Answer{Error: "Expression is not valid"},
		},
		{
			name:               "incorrect expression 3",
			expression:         "incorrect",
			exceptedStatusCode: http.StatusUnprocessableEntity,
			expectedAns:        Answer{Error: "Expression is not valid"},
		},
	}
	for _, testCase := range testCases {
		requestBody, _ := json.Marshal(CalcRequest{Expression: testCase.expression})
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		w := httptest.NewRecorder()

		CalcHandler(w, r)

		if w.Code != testCase.exceptedStatusCode {
			t.Errorf("%s returned wrong status code: got %v want %v", testCase.name, w.Code, testCase.exceptedStatusCode)
		}

		var ans Answer
		err := json.NewDecoder(w.Body).Decode(&ans)

		if err != nil {
			t.Errorf("%s returned error when decoding response body: %v", testCase.name, err)
		}

		if ans != testCase.expectedAns {
			t.Errorf("%s returned incorrect body: got %#v want %#v", testCase.name, ans, testCase.expectedAns)
		}
	}
}
