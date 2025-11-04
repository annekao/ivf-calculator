package calculator

import (
	"math"
	"testing"
)

// Helper function to calculate weight and height from BMI
// BMI = weight(lbs)/703 * (height(in))^2
// For BMI = 22.8 and height = 66 inches:
// weight = 22.8 * 66^2 = 141 lbs
func getWeightHeightForBMI(bmi float64) (weightLbs, heightFt, heightIn int) {
	// Using height of 5'6" as standard
	heightFt = 5
	heightIn = 6
	weightLbs = int(bmi/703.0 * math.Pow(float64(heightFt * 12 + heightIn), 2.0))
	return weightLbs, heightFt, heightIn
}

func TestCalculate_OwnEggs_NoPriorIVF_KnownReason_Scenario1(t *testing.T) {
	// Scenario 1: Using Own Eggs / Did Not Previously Attempt IVF / Known Infertility Reason
	// Age: 32, BMI: 22.8
	// Endometriosis: TRUE, Ovulatory Disorder: TRUE
	// Prior Pregnancies: 1, Prior Live Births: 1

	weightLbs, heightFt, heightIn := getWeightHeightForBMI(22.8)

	req := CalculateRequest{
		Age:              32,
		WeightLbs:        weightLbs,
		HeightFt:		  heightFt,
		HeightIn:         heightIn,
		PriorIvfCycles:   "no",
		PriorPregnancies: 1,
		PriorBirths:      1,
		Reasons: []string{
			"endometriosis",
			"ovulatory_disorder",
		},
		EggSource: "own", // TRUE
	}

	result := Calculate(req)

	// Validate that we got a result
	if result.CumulativeChancePercent != 62.21 {
		t.Errorf("Expected cumulative chance to be 62.21, got %f", result.CumulativeChancePercent)
	}

	t.Logf("Scenario 1 Result: %.2f%% chance of success", result.CumulativeChancePercent)
	t.Logf("BMI: %.2f (Weight: %d lbs, Height: %d in)", 22.8, weightLbs, heightIn)
}

func TestCalculate_OwnEggs_NoPriorIVF_UnknownReason_Scenario2(t *testing.T) {
	// Scenario 2: Using Own Eggs / Did Not Previously Attempt IVF / Unknown Infertility Reason
	// Age: 32, BMI: 22.8
	// All specific factors: FALSE, Unexplained Infertility: FALSE
	// Prior Pregnancies: 1, Prior Live Births: 1
	// Note: Since all specific infertility reasons are FALSE and "Unexplained" is also FALSE,
	// we interpret this as "unknown" reason (reason not yet determined) to satisfy the
	// requirement that at least one reason must be selected.

	weightLbs, heightFt, heightIn := getWeightHeightForBMI(22.8)

	req := CalculateRequest{
		Age:              32,
		WeightLbs:        weightLbs,
		HeightFt:		  heightFt,
		HeightIn:         heightIn,
		PriorIvfCycles:   "no",
		PriorPregnancies: 1,
		PriorBirths:      1,
		Reasons:   []string{"unknown"}, // Reason is unknown (not yet determined)
		EggSource: "own",                // TRUE
	}

	result := Calculate(req)

	// Validate that we got a result
	if result.CumulativeChancePercent != 59.83 {
		t.Errorf("Expected cumulative chance to be 59.83, got %f", result.CumulativeChancePercent)
	}

	t.Logf("Scenario 2 Result: %.2f%% chance of success", result.CumulativeChancePercent)
	t.Logf("BMI: %.2f (Weight: %d lbs, Height: %d in)", 22.8, weightLbs, heightIn)
}

func TestCalculate_OwnEggs_PriorIVF_KnownReason_Scenario3(t *testing.T) {
	// Scenario 3: Using Own Eggs / Previously Attempted IVF / Known Infertility Reason
	// Age: 32, BMI: 22.8
	// Tubal Factor: TRUE, Diminished Ovarian Reserve: TRUE
	// Prior Pregnancies: 1, Prior Live Births: 1

	weightLbs, heightFt, heightIn := getWeightHeightForBMI(22.8)

	req := CalculateRequest{
		Age:              32,
		WeightLbs:        weightLbs,
		HeightFt:		  heightFt,
		HeightIn:         heightIn,
		PriorIvfCycles:   "yes",
		PriorPregnancies: 1,
		PriorBirths:      1,
		Reasons: []string{
			"tubal_factor",
			"diminished_ovarian_reserve",
		},
		EggSource: "own", // TRUE
	}

	result := Calculate(req)

	// Validate that we got a result
	if result.CumulativeChancePercent != 40.89 {
		t.Errorf("Expected cumulative chance to be 40.89, got %f", result.CumulativeChancePercent)
	}

	t.Logf("Scenario 3 Result: %.2f%% chance of success", result.CumulativeChancePercent)
	t.Logf("BMI: %.2f (Weight: %d lbs, Height: %d in)", 22.8, weightLbs, heightIn)
}

func TestCalculate_OwnEggs_PriorIVF_UnknownReason_Scenario4(t *testing.T) {
	// Scenario 4: Using Own Eggs / Previously Attempted IVF / Unknown Infertility Reason
	// Age: 32, BMI: 22.8
	// All specific factors: FALSE, Unexplained Infertility: FALSE
	// Prior Pregnancies: 1, Prior Live Births: 1

	weightLbs, heightFt, heightIn := getWeightHeightForBMI(22.8)

	req := CalculateRequest{
		Age:              32,
		WeightLbs:        weightLbs,
		HeightFt:		  heightFt,
		HeightIn:         heightIn,
		PriorIvfCycles:   "yes",
		PriorPregnancies: 1,
		PriorBirths:      1,
		Reasons: []string{
			"unknown",
		},
		EggSource: "own", // TRUE
	}

	result := Calculate(req)

	// Validate that we got a result
	if result.CumulativeChancePercent != 53.82 {
		t.Errorf("Expected cumulative chance to be 53.82, got %f", result.CumulativeChancePercent)
	}

	t.Logf("Scenario 4 Result: %.2f%% chance of success", result.CumulativeChancePercent)
	t.Logf("BMI: %.2f (Weight: %d lbs, Height: %d in)", 22.8, weightLbs, heightIn)
}

func TestCalculate_DonorEggs_KnownReason_Scenario5(t *testing.T) {
	// Scenario 5: Donor Eggs / Known Infertility Reason
	// Age: 32, BMI: 22.8
	// Uterine Factor: TRUE
	// Prior Pregnancies: 2, Prior Live Births: 1

	weightLbs, heightFt, heightIn := getWeightHeightForBMI(22.8)

	req := CalculateRequest{
		Age:              32,
		WeightLbs:        weightLbs,
		HeightFt:		  heightFt,
		HeightIn:         heightIn,
		PriorIvfCycles:   "",
		PriorPregnancies: 2,
		PriorBirths:      1,
		Reasons: []string{
			"uterine_factor",
		},
		EggSource: "donor", // TRUE
	}

	result := Calculate(req)

	// Validate that we got a result
	if result.CumulativeChancePercent != 55.43 {
		t.Errorf("Expected cumulative chance to be 55.43, got %f", result.CumulativeChancePercent)
	}

	t.Logf("Scenario 5 Result: %.2f%% chance of success", result.CumulativeChancePercent)
	t.Logf("BMI: %.2f (Weight: %d lbs, Height: %d in)", 22.8, weightLbs, heightIn)
}

func TestCalculate_DonorEggs_UnknownReason_Scenario6(t *testing.T) {
	// Scenario 5: Donor Eggs / Unknown Infertility Reason
	// Age: 32, BMI: 22.8
	// Uterine Factor: TRUE
	// Prior Pregnancies: 0, Prior Live Births: 0

	weightLbs, heightFt, heightIn := getWeightHeightForBMI(22.8)

	req := CalculateRequest{
		Age:              32,
		WeightLbs:        weightLbs,
		HeightFt:		  heightFt,
		HeightIn:         heightIn,
		PriorIvfCycles:   "",
		PriorPregnancies: 0,
		PriorBirths:      0,
		Reasons: []string{
			"unknown",
		},
		EggSource: "donor", // TRUE
	}

	result := Calculate(req)

	// Validate that we got a result
	if result.CumulativeChancePercent != 56.8 {
		t.Errorf("Expected cumulative chance to be 56.8, got %f", result.CumulativeChancePercent)
	}

	t.Logf("Scenario 6 Result: %.2f%% chance of success", result.CumulativeChancePercent)
	t.Logf("BMI: %.2f (Weight: %d lbs, Height: %d in)", 22.8, weightLbs, heightIn)
}

// Test that formulas are loaded correctly
func TestFormulaLoading(t *testing.T) {
	if len(formulas) == 0 {
		t.Error("Formulas were not loaded. Expected at least one formula.")
		return
	}

	t.Logf("Loaded %d formulas", len(formulas))

	// Check that we have formulas for different scenarios
	hasOwnEggsNoPrior := false
	hasOwnEggsPrior := false
	hasDonorEggs := false

	for _, formula := range formulas {
		if formula.UsingOwnEggs {
			if formula.AttemptedIVFPreviously != nil && *formula.AttemptedIVFPreviously {
				hasOwnEggsPrior = true
			} else if formula.AttemptedIVFPreviously != nil && !*formula.AttemptedIVFPreviously {
				hasOwnEggsNoPrior = true
			}
		} else {
			hasDonorEggs = true
		}
	}

	t.Logf("Formula coverage: Own eggs no prior IVF: %v, Own eggs prior IVF: %v, Donor eggs: %v",
		hasOwnEggsNoPrior, hasOwnEggsPrior, hasDonorEggs)
}

// Test formula matching logic
func TestFindMatchingFormula1(t *testing.T) {
	// Test scenario 1: Own eggs, no prior IVF, known reason
	req := CalculateRequest{
		EggSource:      "own",
		PriorIvfCycles: "no",
		Reasons:        []string{"endometriosis", "ovulatory_disorder"},
	}

	formula1 := findMatchingFormula(req)
	if formula1 == nil {
		t.Error("Expected to find a matching formula for scenario 1, got nil")
	} else {
		if !formula1.UsingOwnEggs {
			t.Error("Expected formula with UsingOwnEggs=true")
		}
		if formula1.AttemptedIVFPreviously == nil || *formula1.AttemptedIVFPreviously {
			t.Error("Expected formula with AttemptedIVFPreviously=false")
		}
		if !formula1.IsReasonKnown {
			t.Error("Expected formula with IsReasonKnown=true")
		}
		t.Logf("Scenario 1 matched formula: %s", formula1.CDCFormula)
	}
}

func TestFindMatchingFormula2(t *testing.T) {
	// Test scenario 2: Own eggs, no prior IVF, unknown reason
	req := CalculateRequest{
		EggSource:      "own",
		PriorIvfCycles: "no",
		Reasons:        []string{"unknown"},
	}

	formula2 := findMatchingFormula(req)
	if formula2 == nil {
		t.Error("Expected to find a matching formula for scenario 2, got nil")
	} else {
		if !formula2.UsingOwnEggs {
			t.Error("Expected formula with UsingOwnEggs=true")
		}
		if formula2.AttemptedIVFPreviously == nil || *formula2.AttemptedIVFPreviously {
			t.Error("Expected formula with AttemptedIVFPreviously=false")
		}
		if formula2.IsReasonKnown {
			t.Error("Expected formula with IsReasonKnown=false")
		}
		t.Logf("Scenario 2 matched formula: %s", formula2.CDCFormula)
	}
}

func TestFindMatchingFormula3(t *testing.T) {
	// Test scenario 3: Own eggs, prior IVF, known reason
	req := CalculateRequest{
		EggSource:      "own",
		PriorIvfCycles: "yes",
		Reasons:        []string{"tubal_factor", "diminished_ovarian_reserve"},
	}

	formula3 := findMatchingFormula(req)
	if formula3 == nil {
		t.Error("Expected to find a matching formula for scenario 3, got nil")
	} else {
		if !formula3.UsingOwnEggs {
			t.Error("Expected formula with UsingOwnEggs=true")
		}
		if formula3.AttemptedIVFPreviously == nil || !*formula3.AttemptedIVFPreviously {
			t.Error("Expected formula with AttemptedIVFPreviously=true")
		}
		if !formula3.IsReasonKnown {
			t.Error("Expected formula with IsReasonKnown=true")
		}
		t.Logf("Scenario 3 matched formula: %s", formula3.CDCFormula)
	}
}

func TestFindMatchingFormula4(t *testing.T) {
	// Test scenario 4: Own eggs, prior IVF, unknown reason
	req := CalculateRequest{
		EggSource:      "own",
		PriorIvfCycles: "yes",
		Reasons:        []string{"unknown"},
	}

	formula4 := findMatchingFormula(req)
	if formula4 == nil {
		t.Error("Expected to find a matching formula for scenario 4, got nil")
	} else {
		if !formula4.UsingOwnEggs {
			t.Error("Expected formula with UsingOwnEggs=true")
		}
		if formula4.AttemptedIVFPreviously == nil || !*formula4.AttemptedIVFPreviously {
			t.Error("Expected formula with AttemptedIVFPreviously=true")
		}
		if formula4.IsReasonKnown {
			t.Error("Expected formula with IsReasonKnown=false")
		}
		t.Logf("Scenario 4 matched formula: %s", formula4.CDCFormula)
	}
}

func TestFindMatchingFormula5(t *testing.T) {
	// Test scenario 5: Donor eggs, n/a IVF, known reason
	req := CalculateRequest{
		EggSource:      "donor",
		PriorIvfCycles: "yes",
		Reasons:        []string{"uterine_factor"},
	}

	formula5 := findMatchingFormula(req)
	if formula5 == nil {
		t.Error("Expected to find a matching formula for scenario 5, got nil")
	} else {
		if formula5.UsingOwnEggs {
			t.Error("Expected formula with UsingOwnEggs=false")
		}
		if formula5.AttemptedIVFPreviously != nil {
			t.Error("Expected formula with nil AttemptedIVFPreviously")
		}
		if !formula5.IsReasonKnown {
			t.Error("Expected formula with IsReasonKnown=true")
		}
		t.Logf("Scenario 5 matched formula: %s", formula5.CDCFormula)
	}
}

func TestFindMatchingFormula6(t *testing.T) {
	// Test scenario 6: Donor eggs, n/a IVF, unknown reason
	req := CalculateRequest{
		EggSource:      "donor",
		PriorIvfCycles: "yes",
		Reasons:        []string{"unknown"},
	}

	formula6 := findMatchingFormula(req)
	if formula6 == nil {
		t.Error("Expected to find a matching formula for scenario 6, got nil")
	} else {
		if formula6.UsingOwnEggs {
			t.Error("Expected formula with UsingOwnEggs=false")
		}
		if formula6.AttemptedIVFPreviously != nil {
			t.Error("Expected formula with nil AttemptedIVFPreviously")
		}
		if formula6.IsReasonKnown {
			t.Error("Expected formula with IsReasonKnown=false")
		}
		t.Logf("Scenario 6 matched formula: %s", formula6.CDCFormula)
	}
}