package validation

import (
	"ivf-calculator-backend/internal/calculator"
)

// ValidateCalculateRequest validates the calculate request and returns errors if any
func ValidateCalculateRequest(req calculator.CalculateRequest) map[string]string {
	errors := make(map[string]string)

	if req.Age < 20 || req.Age > 50 {
		errors["age"] = "must be between 20 and 50"
	}

	if req.WeightLbs < 80 || req.WeightLbs > 300 {
		errors["weightLbs"] = "must be between 80 and 300"
	}

	if req.HeightFt < 4 || req.HeightFt > 6 {
		errors["heightFt"] = "must be between 4 and 7"
	}

	if req.HeightIn < 0 || req.HeightIn > 11 {
		errors["heightIn"] = "must be between 0 and 12"
	}

	if req.PriorPregnancies < 0 || req.PriorPregnancies > 2 {
		errors["priorPregnancies"] = "must be 0, 1, or 2+"
	}

	if req.PriorBirths < 0 || req.PriorBirths > 2 {
		errors["priorBirths"] = "must be 0, 1, or 2+"
	}

	if req.PriorBirths > req.PriorPregnancies {
		if (errors["priorBirths"] != "") {
			errors["priorBirths"] += "; "
		}
		errors["priorBirths"] += "cannot exceed the number of prior pregnancies (even in the case of twins)"
	}

	if req.EggSource != "own" && req.EggSource != "donor" {
		errors["eggSource"] = "must be 'own' or 'donor'"
	}

	if req.EggSource == "own" && req.PriorIvfCycles == "" {
		errors["priorIvfCycles"] = "must be 'yes' or 'no' when planning to use 'own' eggs"
	}

	if len(req.Reasons) == 0 {
		errors["reasons"] = "at least one reason must be selected"
	}

	validReasons := map[string]bool{
		"male_factor_infertility":  true,
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
