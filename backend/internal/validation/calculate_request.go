package validation

import (
	"ivf-calculator-backend/internal/http/handlers"
)

// ValidateCalculateRequest validates the calculate request and returns errors if any
func ValidateCalculateRequest(req handlers.CalculateRequest) map[string]string {
	errors := make(map[string]string)

	if req.Age < 20 || req.Age > 50 {
		errors["age"] = "must be between 20 and 50"
	}

	if req.WeightLbs < 80 || req.WeightLbs > 300 {
		errors["weightLbs"] = "must be between 80 and 300"
	}

	if req.HeightIn < 55 || req.HeightIn > 78 {
		errors["heightIn"] = "must be between 55 and 78"
	}

	if req.PriorIvfCycles < 0 || req.PriorIvfCycles > 3 {
		errors["priorIvfCycles"] = "must be 0-3"
	}

	if req.PriorPregnancies < 0 || req.PriorPregnancies > 2 {
		errors["priorPregnancies"] = "must be 0, 1, or 2+"
	}

	if req.PriorBirths < 0 || req.PriorBirths > 2 {
		errors["priorBirths"] = "must be 0, 1, or 2+"
	}

	if req.EggSource != "own" && req.EggSource != "donor" {
		errors["eggSource"] = "must be 'own' or 'donor'"
	}

	if req.Retrievals < 1 || req.Retrievals > 3 {
		errors["retrievals"] = "must be 1-3"
	}

	if len(req.Reasons) == 0 {
		errors["reasons"] = "at least one reason must be selected"
	}

	validReasons := map[string]bool{
		"male_factor":              true,
		"endometriosis":            true,
		"tubal_factor":             true,
		"ovulatory_disorder":       true,
		"diminished_ovarian_reserve": true,
		"uterine_factor":           true,
		"other":                    true,
		"unexplained":              true,
		"unknown":                  true,
	}

	for _, reason := range req.Reasons {
		if !validReasons[reason] {
			errors["reasons"] = "invalid reason: " + reason
			break
		}
	}

	return errors
}
