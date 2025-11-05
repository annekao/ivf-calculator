package validation

import (
	"ivf-calculator-backend/internal/calculator"
	"reflect"
	"testing"
)

func TestValidateCalculateRequest(t *testing.T) {
	tests := []struct {
		name     string
		req      calculator.CalculateRequest
		wantErrs map[string]string
	}{
		{
			name: "valid request",
			req: calculator.CalculateRequest{
				Age:              30,
				WeightLbs:        150,
				HeightFt:         5,
				HeightIn:         6,
				EggSource:        "own",
				PriorIvfCycles:   "no",
				PriorPregnancies: 1,
				PriorBirths:      1,
				Reasons:          []string{"male_factor_infertility"},
			},
			wantErrs: map[string]string{},
		},
		{
			name: "invalid age, weight, height, and heightIn",
			req: calculator.CalculateRequest{
				Age:      19,
				WeightLbs: 301,
				HeightFt:  3,
				HeightIn:  12,
			},
			wantErrs: map[string]string{
				"age":      "must be between 20 and 50",
				"weightLbs": "must be between 80 and 300",
				"heightFt":  "must be between 4 and 7",
				"heightIn":  "must be between 0 and 12",
				"eggSource": "must be 'own' or 'donor'",
				"reasons":   "at least one reason must be selected",
			},
		},
		{
			name: "eggSource own without prior IVF cycles",
			req: calculator.CalculateRequest{
				Age:       35,
				WeightLbs: 140,
				HeightFt:  5,
				HeightIn:  5,
				EggSource: "own",
				Reasons:   []string{"other"},
			},
			wantErrs: map[string]string{
				"priorIvfCycles": "must be 'yes' or 'no' when planning to use 'own' eggs",
			},
		},
		{
			name: "invalid pregnancies and births relationship",
			req: calculator.CalculateRequest{
				Age:       35,
				WeightLbs: 140,
				HeightFt:  5,
				HeightIn:  5,
				EggSource: "donor",
				PriorPregnancies: 1,
				PriorBirths:      2,
				Reasons:          []string{"other"},
			},
			wantErrs: map[string]string{
				"priorBirths": "cannot exceed the number of prior pregnancies (even in the case of twins)",
			},
		},
		{
			name: "invalid reasons",
			req: calculator.CalculateRequest{
				Age:       35,
				WeightLbs: 140,
				HeightFt:  5,
				HeightIn:  5,
				EggSource: "donor",
				PriorPregnancies: 0,
				PriorBirths:      0,
				Reasons: []string{"invalid_reason"},
			},
			wantErrs: map[string]string{
				"reasons": "invalid reason: invalid_reason",
			},
		},
		{
			name: "unexplained reason must be alone",
			req: calculator.CalculateRequest{
				Age:       35,
				WeightLbs: 140,
				HeightFt:  5,
				HeightIn:  5,
				EggSource: "donor",
				PriorPregnancies: 0,
				PriorBirths:      0,
				Reasons: []string{"unexplained", "male_factor_infertility"},
			},
			wantErrs: map[string]string{
				"reasons": "'Unexplained (Idiopathic) infertility' must be selected by itself",
			},
		},
		{
			name: "unknown reason must be alone",
			req: calculator.CalculateRequest{
				Age:       35,
				WeightLbs: 140,
				HeightFt:  5,
				HeightIn:  5,
				EggSource: "donor",
				PriorPregnancies: 0,
				PriorBirths:      0,
				Reasons: []string{"unknown", "endometriosis"},
			},
			wantErrs: map[string]string{
				"reasons": "'I don't know/no reason' must be selected by itself",
			},
		},
		{
			name: "heightFt upper boundary",
			req: calculator.CalculateRequest{
				Age:       35,
				WeightLbs: 140,
				HeightFt:  7,
				HeightIn:  5,
				EggSource: "donor",
				PriorPregnancies: 0,
				PriorBirths:      0,
				Reasons:  []string{"other"},
			},
			wantErrs: map[string]string{
				"heightFt": "must be between 4 and 7",
			},
		},
		{
			name: "empty reasons",
			req: calculator.CalculateRequest{
				Age:       35,
				WeightLbs: 140,
				HeightFt:  5,
				HeightIn:  5,
				EggSource: "donor",
				PriorPregnancies: 0,
				PriorBirths:      0,
				Reasons: []string{},
			},
			wantErrs: map[string]string{
				"reasons": "at least one reason must be selected",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErrs := ValidateCalculateRequest(tt.req)

			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("ValidateCalculateRequest() = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}