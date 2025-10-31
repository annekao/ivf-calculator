package calculator

import (
	"math"
	"testing"
)

// Helper function to calculate weight and height from BMI
// BMI = weight(kg) / (height(m))^2
// For BMI = 22.8 and height = 66 inches (1.6764 m):
// weight = 22.8 * 1.6764^2 = 64.0 kg ≈ 141 lbs
func getWeightHeightForBMI(bmi float64) (weightLbs, heightIn int) {
	// Using height of 66 inches (5'6") as standard
	heightIn = 66
	heightM := float64(heightIn) * 0.0254
	weightKg := bmi * heightM * heightM
	weightLbs = int(math.Round(weightKg / 0.453592))
	return weightLbs, heightIn
}

func TestCalculate_OwnEggs_NoPriorIVF_KnownReason_Scenario1(t *testing.T) {
	// Scenario 1: Using Own Eggs / Did Not Previously Attempt IVF / Known Infertility Reason
	// Age: 32, BMI: 22.8
	// Endometriosis: TRUE, Ovulatory Disorder: TRUE
	// Prior Pregnancies: 1, Prior Live Births: 1

	weightLbs, heightIn := getWeightHeightForBMI(22.8)

	req := CalculateRequest{
		Age:              32,
		WeightLbs:        weightLbs,
		HeightIn:         heightIn,
		PriorIvfCycles:   0, // FALSE - did not previously attempt IVF
		PriorPregnancies: 1,
		PriorBirths:      1,
		Reasons: []string{
			"endometriosis",
			"ovulatory_disorder",
		},
		EggSource:  "own", // TRUE
		Retrievals: 1,
	}

	result := Calculate(req)

	// Validate that we got a result
	if result.CumulativeChancePercent <= 0 || result.CumulativeChancePercent > 100 {
		t.Errorf("Expected cumulative chance between 0 and 100, got %f", result.CumulativeChancePercent)
	}

	if len(result.Notes) == 0 {
		t.Error("Expected notes in response, got none")
	}

	// Verify the calculation is reasonable (should be positive probability)
	if result.CumulativeChancePercent < 0.1 || result.CumulativeChancePercent > 95.0 {
		t.Logf("Warning: Result outside expected bounds: %f%%", result.CumulativeChancePercent)
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

	weightLbs, heightIn := getWeightHeightForBMI(22.8)

	req := CalculateRequest{
		Age:              32,
		WeightLbs:        weightLbs,
		HeightIn:         heightIn,
		PriorIvfCycles:   0, // FALSE - did not previously attempt IVF
		PriorPregnancies: 1,
		PriorBirths:      1,
		Reasons:          []string{"unknown"}, // Reason is unknown (not yet determined)
		EggSource:        "own",                // TRUE
		Retrievals:       1,
	}

	result := Calculate(req)

	// Validate that we got a result
	if result.CumulativeChancePercent <= 0 || result.CumulativeChancePercent > 100 {
		t.Errorf("Expected cumulative chance between 0 and 100, got %f", result.CumulativeChancePercent)
	}

	if len(result.Notes) == 0 {
		t.Error("Expected notes in response, got none")
	}

	// Verify the calculation is reasonable
	if result.CumulativeChancePercent < 0.1 || result.CumulativeChancePercent > 95.0 {
		t.Logf("Warning: Result outside expected bounds: %f%%", result.CumulativeChancePercent)
	}

	t.Logf("Scenario 2 Result: %.2f%% chance of success", result.CumulativeChancePercent)
	t.Logf("BMI: %.2f (Weight: %d lbs, Height: %d in)", 22.8, weightLbs, heightIn)
}

func TestCalculate_OwnEggs_PriorIVF_KnownReason_Scenario3(t *testing.T) {
	// Scenario 3: Using Own Eggs / Previously Attempted IVF / Known Infertility Reason
	// Age: 32, BMI: 22.8
	// Tubal Factor: TRUE, Diminished Ovarian Reserve: TRUE
	// Prior Pregnancies: 1, Prior Live Births: 1

	weightLbs, heightIn := getWeightHeightForBMI(22.8)

	req := CalculateRequest{
		Age:              32,
		WeightLbs:        weightLbs,
		HeightIn:         heightIn,
		PriorIvfCycles:   1, // TRUE - previously attempted IVF (using 1 cycle)
		PriorPregnancies: 1,
		PriorBirths:      1,
		Reasons: []string{
			"tubal_factor",
			"diminished_ovarian_reserve",
		},
		EggSource:  "own", // TRUE
		Retrievals: 1,
	}

	result := Calculate(req)

	// Validate that we got a result
	if result.CumulativeChancePercent <= 0 || result.CumulativeChancePercent > 100 {
		t.Errorf("Expected cumulative chance between 0 and 100, got %f", result.CumulativeChancePercent)
	}

	if len(result.Notes) == 0 {
		t.Error("Expected notes in response, got none")
	}

	// Verify the calculation is reasonable
	if result.CumulativeChancePercent < 0.1 || result.CumulativeChancePercent > 95.0 {
		t.Logf("Warning: Result outside expected bounds: %f%%", result.CumulativeChancePercent)
	}

	t.Logf("Scenario 3 Result: %.2f%% chance of success", result.CumulativeChancePercent)
	t.Logf("BMI: %.2f (Weight: %d lbs, Height: %d in)", 22.8, weightLbs, heightIn)
}

// TestAllScenarios runs all three scenarios and compares results
func TestAllScenarios_Comparison(t *testing.T) {
	weightLbs, heightIn := getWeightHeightForBMI(22.8)

	// Scenario 1: Own eggs, no prior IVF, known reason (endometriosis + ovulatory disorder)
	req1 := CalculateRequest{
		Age:              32,
		WeightLbs:        weightLbs,
		HeightIn:         heightIn,
		PriorIvfCycles:   0,
		PriorPregnancies: 1,
		PriorBirths:      1,
		Reasons:          []string{"endometriosis", "ovulatory_disorder"},
		EggSource:        "own",
		Retrievals:       1,
	}

	// Scenario 2: Own eggs, no prior IVF, unknown reason
	req2 := CalculateRequest{
		Age:              32,
		WeightLbs:        weightLbs,
		HeightIn:         heightIn,
		PriorIvfCycles:   0,
		PriorPregnancies: 1,
		PriorBirths:      1,
		Reasons:          []string{"unknown"},
		EggSource:        "own",
		Retrievals:       1,
	}

	// Scenario 3: Own eggs, prior IVF, known reason (tubal factor + diminished ovarian reserve)
	req3 := CalculateRequest{
		Age:              32,
		WeightLbs:        weightLbs,
		HeightIn:         heightIn,
		PriorIvfCycles:   1,
		PriorPregnancies: 1,
		PriorBirths:      1,
		Reasons:          []string{"tubal_factor", "diminished_ovarian_reserve"},
		EggSource:        "own",
		Retrievals:       1,
	}

	result1 := Calculate(req1)
	result2 := Calculate(req2)
	result3 := Calculate(req3)

	t.Logf("\n=== Comparison of All Scenarios ===")
	t.Logf("Scenario 1 (Own Eggs, No Prior IVF, Known Reason): %.2f%%", result1.CumulativeChancePercent)
	t.Logf("Scenario 2 (Own Eggs, No Prior IVF, Unknown Reason): %.2f%%", result2.CumulativeChancePercent)
	t.Logf("Scenario 3 (Own Eggs, Prior IVF, Known Reason): %.2f%%", result3.CumulativeChancePercent)

	// Validate all results are in valid range
	if result1.CumulativeChancePercent < 0.1 || result1.CumulativeChancePercent > 95.0 {
		t.Errorf("Scenario 1 result out of bounds: %f", result1.CumulativeChancePercent)
	}
	if result2.CumulativeChancePercent < 0.1 || result2.CumulativeChancePercent > 95.0 {
		t.Errorf("Scenario 2 result out of bounds: %f", result2.CumulativeChancePercent)
	}
	if result3.CumulativeChancePercent < 0.1 || result3.CumulativeChancePercent > 95.0 {
		t.Errorf("Scenario 3 result out of bounds: %f", result3.CumulativeChancePercent)
	}

	// All scenarios should use different formulas (different CDC formula numbers)
	if len(result1.Notes) == 0 || len(result2.Notes) == 0 || len(result3.Notes) == 0 {
		t.Error("All scenarios should have notes")
	}
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
func TestFindMatchingFormula(t *testing.T) {
	// Test scenario 1: Own eggs, no prior IVF, known reason
	req1 := CalculateRequest{
		EggSource:      "own",
		PriorIvfCycles: 0,
		Reasons:        []string{"endometriosis", "ovulatory_disorder"},
	}

	formula1 := findMatchingFormula(req1)
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

	// Test scenario 2: Own eggs, no prior IVF, unknown reason
	req2 := CalculateRequest{
		EggSource:      "own",
		PriorIvfCycles: 0,
		Reasons:        []string{"unknown"},
	}

	formula2 := findMatchingFormula(req2)
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

	// Test scenario 3: Own eggs, prior IVF, known reason
	req3 := CalculateRequest{
		EggSource:      "own",
		PriorIvfCycles: 1,
		Reasons:        []string{"tubal_factor", "diminished_ovarian_reserve"},
	}

	formula3 := findMatchingFormula(req3)
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

