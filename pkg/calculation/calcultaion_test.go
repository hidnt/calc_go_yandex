package calculation

import "testing"

type testCase struct {
	name        string
	expression  string
	exceptedRes float64
	wantError   bool
}

func TestCalc(t *testing.T) {
	testCases := []testCase{
		{
			name:        "one number 1",
			expression:  "1012",
			exceptedRes: 1012,
			wantError:   false,
		},
		{
			name:        "one number 2",
			expression:  "-1024",
			exceptedRes: -1024,
			wantError:   false,
		},
		{
			name:        "minus 1",
			expression:  "-10-5",
			exceptedRes: -15,
			wantError:   false,
		},
		{
			name:        "minus 2",
			expression:  "-10--5",
			exceptedRes: -5,
			wantError:   false,
		},
		{
			name:        "minus 3",
			expression:  "-10*(-5+2)",
			exceptedRes: 30,
			wantError:   false,
		},
		{
			name:        "minus 3",
			expression:  "-10---5",
			exceptedRes: 0,
			wantError:   true,
		},
		{
			name:        "priority 1",
			expression:  "2+2*2",
			exceptedRes: 6,
			wantError:   false,
		},
		{
			name:        "priority 2",
			expression:  "(2+2-(-2+7)*2)/2",
			exceptedRes: -3,
			wantError:   false,
		},
		{
			name:        "div by zero",
			expression:  "2/0",
			exceptedRes: 0,
			wantError:   true,
		},
		{
			name:        "fractional numbers 1",
			expression:  "140/12.5",
			exceptedRes: 11.2,
			wantError:   false,
		},
		{
			name:        "fractional numbers 2",
			expression:  "140.5*12.5/10",
			exceptedRes: 175.625,
			wantError:   false,
		},
		{
			name:        "incorrect expression 1",
			expression:  "1238)",
			exceptedRes: 0,
			wantError:   true,
		},
		{
			name:        "incorrect expression 2",
			expression:  "124+2-",
			exceptedRes: 0,
			wantError:   true,
		},
		{
			name:        "incorrect expression 3",
			expression:  "/",
			exceptedRes: 0,
			wantError:   true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ans, err := Calc(testCase.expression)
			if testCase.wantError {
				if err == nil {
					t.Fatalf("Excepted an err")
				}
			} else {
				if err != nil {
					t.Fatalf("Successful case is %s, but returns error: %s", testCase.expression, err.Error())
				}
				if ans != testCase.exceptedRes {
					t.Fatalf("%f should be equal %f", ans, testCase.exceptedRes)
				}
			}
		})
	}
}
