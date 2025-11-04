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

	weightLbs, heightIn := getWeightHeightForBMI(22.8)

	req := CalculateRequest{
		Age:              32,
		WeightLbs:        weightLbs,
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

	weightLbs, heightIn := getWeightHeightForBMI(22.8)

	req := CalculateRequest{
		Age:              32,
		WeightLbs:        weightLbs,
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
		PriorIvfCycles: "no",
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
		PriorIvfCycles: "no",
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
		PriorIvfCycles: "yes",
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

