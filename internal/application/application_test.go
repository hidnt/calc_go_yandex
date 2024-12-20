package application

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testCase struct {
	name               string
	expression         string
	exceptedStatusCode int
	exceptedRes        float64
	wantError          bool
}

func TestCalc(t *testing.T) {
	testCases := []testCase{
		{
			name:               "minus",
			expression:         "-10*(-5+2)",
			exceptedStatusCode: http.StatusOK,
			exceptedRes:        30,
			wantError:          false,
		},
		{
			name:               "priority",
			expression:         "(2+2-(-2+7)*2)/2",
			exceptedStatusCode: http.StatusOK,
			exceptedRes:        -3,
			wantError:          false,
		},
		{
			name:               "fractional numbers 2",
			expression:         "140.5*12.5/10",
			exceptedStatusCode: http.StatusOK,
			exceptedRes:        175.625,
			wantError:          false,
		},
		{
			name:               "incorrect expression 1",
			expression:         "1238)",
			exceptedStatusCode: http.StatusUnprocessableEntity,
			exceptedRes:        0,
			wantError:          true,
		},
		{
			name:               "incorrect expression 2",
			expression:         "124+2-",
			exceptedStatusCode: http.StatusUnprocessableEntity,
			exceptedRes:        0,
			wantError:          true,
		},
		{
			name:               "incorrect expression 3",
			expression:         "incorrect",
			exceptedStatusCode: http.StatusUnprocessableEntity,
			exceptedRes:        0,
			wantError:          true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			jsonExpression := fmt.Sprintf("{\"expression\": \"%s\"}", testCase.expression)
			request, _ := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte(jsonExpression)))
			response := httptest.NewRecorder()
			CalcHandler(response, request)
			excepted := fmt.Sprintf("{\n\t\"result\": \"%f\"\n}", testCase.exceptedRes)
			if testCase.wantError {
				if testCase.exceptedStatusCode != response.Code {
					t.Errorf("successful case is %d, but returns error: %d", testCase.exceptedStatusCode, response.Code)
				}
			} else {
				if response.Body.String() != excepted {
					t.Errorf("Want %s, got %s", excepted, response.Body.String())
				}
			}
		})
	}
}
