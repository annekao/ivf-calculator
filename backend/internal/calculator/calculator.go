package calculator

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// CalculateRequest represents the request body for the calculate endpoint
type CalculateRequest struct {
	Age              int      `json:"age" binding:"required"`
	WeightLbs        int      `json:"weightLbs" binding:"required"`
	HeightIn         int      `json:"heightIn" binding:"required"`
	PriorIvfCycles   int      `json:"priorIvfCycles" binding:"required"`
	PriorPregnancies int      `json:"priorPregnancies" binding:"required"`
	PriorBirths      int      `json:"priorBirths" binding:"required"`
	Reasons          []string `json:"reasons" binding:"required"`
	EggSource        string   `json:"eggSource" binding:"required"`
	Retrievals       int      `json:"retrievals" binding:"required"`
}

// CalculateResponse represents the response from the calculate endpoint
type CalculateResponse struct {
	CumulativeChancePercent float64  `json:"cumulativeChancePercent"`
	Notes                   []string `json:"notes"`
}

// Formula represents a CDC formula with all its coefficients
type Formula struct {
	UsingOwnEggs                bool
	AttemptedIVFPreviously      *bool // nil means N/A (donor eggs)
	IsReasonKnown               bool
	CDCFormula                  string
	Intercept                   float64
	AgeLinearCoeff              float64
	AgePowerCoeff               float64
	AgePowerFactor              float64
	BMILinearCoeff              float64
	BMIPowerCoeff               float64
	BMIPowerFactor              float64
	TubalFactorTrue             float64
	TubalFactorFalse            float64
	MaleFactorTrue              float64
	MaleFactorFalse             float64
	EndometriosisTrue           float64
	EndometriosisFalse          float64
	OvulatoryDisorderTrue       float64
	OvulatoryDisorderFalse      float64
	DiminishedOvarianReserveTrue  float64
	DiminishedOvarianReserveFalse float64
	UterineFactorTrue           float64
	UterineFactorFalse          float64
	OtherReasonTrue             float64
	OtherReasonFalse            float64
	UnexplainedInfertilityTrue  float64
	UnexplainedInfertilityFalse float64
	PriorPregnancies0           float64
	PriorPregnancies1           float64
	PriorPregnancies2Plus       float64
	PriorLiveBirths0            float64
	PriorLiveBirths1            float64
	PriorLiveBirths2Plus        float64
}

var formulas []Formula

// init loads the formulas from CSV on package initialization
func init() {
	if err := loadFormulas(); err != nil {
		// Log error but don't fail - will fall back to mock calculation
		fmt.Printf("Warning: Failed to load formulas: %v\n", err)
	}
}

// loadFormulas reads and parses the CSV file containing IVF success formulas
func loadFormulas() error {
	csvPath := getCSVPath()

	file, err := os.Open(csvPath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	// Read header
	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}

	// Create a map of column indices
	colIndex := make(map[string]int)
	for i, col := range header {
		colIndex[strings.TrimSpace(col)] = i
	}

	formulas = []Formula{}
	
	// Read data rows
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read CSV record: %w", err)
		}

		formula := Formula{}

		// Parse boolean parameters
		formula.UsingOwnEggs = parseBool(record[colIndex["param_using_own_eggs"]])
		
		attemptedIVFStr := record[colIndex["param_attempted_ivf_previously"]]
		if attemptedIVFStr == "N/A" || attemptedIVFStr == "" {
			formula.AttemptedIVFPreviously = nil
		} else {
			val := parseBool(attemptedIVFStr)
			formula.AttemptedIVFPreviously = &val
		}
		
		formula.IsReasonKnown = parseBool(record[colIndex["param_is_reason_for_infertility_known"]])
		formula.CDCFormula = record[colIndex["cdc_formula"]]

		// Parse numeric coefficients
		formula.Intercept = parseFloat(record[colIndex["formula_intercept"]])
		formula.AgeLinearCoeff = parseFloat(record[colIndex["formula_age_linear_coefficient"]])
		formula.AgePowerCoeff = parseFloat(record[colIndex["formula_age_power_coefficient"]])
		formula.AgePowerFactor = parseFloat(record[colIndex["formula_age_power_factor"]])
		formula.BMILinearCoeff = parseFloat(record[colIndex["formula_bmi_linear_coefficient"]])
		formula.BMIPowerCoeff = parseFloat(record[colIndex["formula_bmi_power_coefficient"]])
		formula.BMIPowerFactor = parseFloat(record[colIndex["formula_bmi_power_factor"]])
		formula.TubalFactorTrue = parseFloat(record[colIndex["formula_tubal_factor_true_value"]])
		formula.TubalFactorFalse = parseFloat(record[colIndex["formula_tubal_factor_false_value"]])
		formula.MaleFactorTrue = parseFloat(record[colIndex["formula_male_factor_infertility_true_value"]])
		formula.MaleFactorFalse = parseFloat(record[colIndex["formula_male_factor_infertility_false_value"]])
		formula.EndometriosisTrue = parseFloat(record[colIndex["formula_endometriosis_true_value"]])
		formula.EndometriosisFalse = parseFloat(record[colIndex["formula_endometriosis_false_value"]])
		formula.OvulatoryDisorderTrue = parseFloat(record[colIndex["formula_ovulatory_disorder_true_value"]])
		formula.OvulatoryDisorderFalse = parseFloat(record[colIndex["formula_ovulatory_disorder_false_value"]])
		formula.DiminishedOvarianReserveTrue = parseFloat(record[colIndex["formula_diminished_ovarian_reserve_true_value"]])
		formula.DiminishedOvarianReserveFalse = parseFloat(record[colIndex["formula_diminished_ovarian_reserve_false_value"]])
		formula.UterineFactorTrue = parseFloat(record[colIndex["formula_uterine_factor_true_value"]])
		formula.UterineFactorFalse = parseFloat(record[colIndex["formula_uterine_factor_false_value"]])
		formula.OtherReasonTrue = parseFloat(record[colIndex["formula_other_reason_true_value"]])
		formula.OtherReasonFalse = parseFloat(record[colIndex["formula_other_reason_false_value"]])
		formula.UnexplainedInfertilityTrue = parseFloat(record[colIndex["formula_unexplained_infertility_true_value"]])
		formula.UnexplainedInfertilityFalse = parseFloat(record[colIndex["formula_unexplained_infertility_false_value"]])
		formula.PriorPregnancies0 = parseFloat(record[colIndex["formula_prior_pregnancies_0_value"]])
		formula.PriorPregnancies1 = parseFloat(record[colIndex["formula_prior_pregnancies_1_value"]])
		formula.PriorPregnancies2Plus = parseFloat(record[colIndex["formula_prior_pregnancies_2+_value"]])
		formula.PriorLiveBirths0 = parseFloat(record[colIndex["formula_prior_live_births_0_value"]])
		formula.PriorLiveBirths1 = parseFloat(record[colIndex["formula_prior_live_births_1_value"]])
		formula.PriorLiveBirths2Plus = parseFloat(record[colIndex["formula_prior_live_births_2+_value"]])

		formulas = append(formulas, formula)
	}

	return nil
}

// Helper functions
func parseBool(s string) bool {
	return strings.ToUpper(strings.TrimSpace(s)) == "TRUE"
}

func parseFloat(s string) float64 {
	val, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0.0
	}
	return val
}

// getCSVPath returns the path to the CSV file
func getCSVPath() string {
	// Get the directory of the caller (calculator.go)
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		// Fallback: try relative path
		_, filename, _, ok = runtime.Caller(0)
		if !ok {
			return "ivf_success_formulas.csv"
		}
	}
	return filepath.Join(filepath.Dir(filename), "ivf_success_formulas.csv")
}

// findMatchingFormula selects the appropriate formula based on patient parameters
func findMatchingFormula(req CalculateRequest) *Formula {
	usingOwnEggs := req.EggSource == "own"
	attemptedIVFPreviously := req.PriorIvfCycles > 0
	
	// Determine if reason is known (not unexplained or unknown)
	isReasonKnown := true
	hasUnexplainedOrUnknown := false
	for _, reason := range req.Reasons {
		if reason == "unexplained" || reason == "unknown" {
			hasUnexplainedOrUnknown = true
		}
	}
	// If only unexplained/unknown, then reason is not known
	if hasUnexplainedOrUnknown && len(req.Reasons) == 1 {
		isReasonKnown = false
	} else if hasUnexplainedOrUnknown {
		// Mixed: if it's only unexplained/unknown, treat as unknown, otherwise known
		isReasonKnown = true
	}

	for i := range formulas {
		f := &formulas[i]
		
		// Match using own eggs
		if f.UsingOwnEggs != usingOwnEggs {
			continue
		}
		
		// Match attempted IVF previously (for own eggs only)
		if f.UsingOwnEggs {
			if f.AttemptedIVFPreviously == nil {
				continue
			}
			if *f.AttemptedIVFPreviously != attemptedIVFPreviously {
				continue
			}
		} else {
			// For donor eggs, attemptedIVFPreviously should be nil (N/A)
			if f.AttemptedIVFPreviously != nil {
				continue
			}
		}
		
		// Match reason known status
		if f.IsReasonKnown != isReasonKnown {
			continue
		}
		
		return f
	}
	
	return nil
}

// calculateBMI computes BMI from weight in pounds and height in inches
func calculateBMI(weightLbs, heightIn int) float64 {
	weightKg := float64(weightLbs) * 0.453592
	heightM := float64(heightIn) * 0.0254
	return weightKg / (heightM * heightM)
}

// getPriorPregnanciesValue returns the coefficient for prior pregnancies
func (f *Formula) getPriorPregnanciesValue(count int) float64 {
	if count == 0 {
		return f.PriorPregnancies0
	} else if count == 1 {
		return f.PriorPregnancies1
	}
	return f.PriorPregnancies2Plus
}

// getPriorLiveBirthsValue returns the coefficient for prior live births
func (f *Formula) getPriorLiveBirthsValue(count int) float64 {
	if count == 0 {
		return f.PriorLiveBirths0
	} else if count == 1 {
		return f.PriorLiveBirths1
	}
	return f.PriorLiveBirths2Plus
}

// Calculate performs IVF success rate calculation using CDC formulas
func Calculate(req CalculateRequest) CalculateResponse {
	// If no formulas loaded, fall back to mock calculation
	if len(formulas) == 0 {
		return mockCalculate(req)
	}

	// Find matching formula
	formula := findMatchingFormula(req)
	if formula == nil {
		// If no matching formula, fall back to mock
		return mockCalculate(req)
	}

	// Calculate BMI
	bmi := calculateBMI(req.WeightLbs, req.HeightIn)
	age := float64(req.Age)

	// Calculate logit using the formula
	logit := formula.Intercept

	// Age terms: linear + power terms
	// Typically polynomial regression uses (age - center)^2 for quadratic terms
	agePower := math.Pow(age-formula.AgePowerFactor, 2.0)
	logit += formula.AgeLinearCoeff*age + formula.AgePowerCoeff*agePower

	// BMI terms: linear + power terms
	// BMI power factor is 2.0 for all formulas, so use (bmi - 2)^2
	bmiPower := math.Pow(bmi-formula.BMIPowerFactor, 2.0)
	logit += formula.BMILinearCoeff*bmi + formula.BMIPowerCoeff*bmiPower

	// Infertility factor terms
	hasTubalFactor := contains(req.Reasons, "tubal_factor")
	logit += ternary(hasTubalFactor, formula.TubalFactorTrue, formula.TubalFactorFalse)

	hasMaleFactor := contains(req.Reasons, "male_factor")
	logit += ternary(hasMaleFactor, formula.MaleFactorTrue, formula.MaleFactorFalse)

	hasEndometriosis := contains(req.Reasons, "endometriosis")
	logit += ternary(hasEndometriosis, formula.EndometriosisTrue, formula.EndometriosisFalse)

	hasOvulatoryDisorder := contains(req.Reasons, "ovulatory_disorder")
	logit += ternary(hasOvulatoryDisorder, formula.OvulatoryDisorderTrue, formula.OvulatoryDisorderFalse)

	hasDiminishedOvarianReserve := contains(req.Reasons, "diminished_ovarian_reserve")
	logit += ternary(hasDiminishedOvarianReserve, formula.DiminishedOvarianReserveTrue, formula.DiminishedOvarianReserveFalse)

	hasUterineFactor := contains(req.Reasons, "uterine_factor")
	logit += ternary(hasUterineFactor, formula.UterineFactorTrue, formula.UterineFactorFalse)

	hasOther := contains(req.Reasons, "other")
	logit += ternary(hasOther, formula.OtherReasonTrue, formula.OtherReasonFalse)

	hasUnexplained := contains(req.Reasons, "unexplained") || contains(req.Reasons, "unknown")
	logit += ternary(hasUnexplained, formula.UnexplainedInfertilityTrue, formula.UnexplainedInfertilityFalse)

	// Prior pregnancies and births
	logit += formula.getPriorPregnanciesValue(req.PriorPregnancies)
	logit += formula.getPriorLiveBirthsValue(req.PriorBirths)

	// Convert logit to probability
	probability := 1.0 / (1.0 + math.Exp(-logit))

	// Convert to percentage and apply retrieval multiplier
	// The formula gives per-cycle probability, so for multiple retrievals:
	// P(at least one success) = 1 - (1 - p)^n
	perCycleProb := probability
	cumulativeProb := 1.0 - math.Pow(1.0-perCycleProb, float64(req.Retrievals))
	chancePercent := cumulativeProb * 100.0

	// Clamp between 0.1% and 95%
	if chancePercent < 0.1 {
		chancePercent = 0.1
	}
	if chancePercent > 95.0 {
		chancePercent = 95.0
	}

	return CalculateResponse{
		CumulativeChancePercent: chancePercent,
		Notes: []string{
			"Calculations are based on CDC statistical models for IVF success rates.",
			"Consult a physician for personalized assessment based on your specific medical history.",
			"Actual success rates may vary based on clinic, specific protocols, and individual factors.",
		},
	}
}

// Helper functions
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func ternary(condition bool, trueVal, falseVal float64) float64 {
	if condition {
		return trueVal
	}
	return falseVal
}

// mockCalculate is the fallback calculation if formulas aren't loaded
func mockCalculate(req CalculateRequest) CalculateResponse {
	// Base chance percentage
	baseChance := 45.0

	// Age adjustment: decrease chance as age increases above 30
	agePenalty := 0.0
	if req.Age > 30 {
		agePenalty = float64(req.Age-30) * 1.2
	}

	// Retrieval boost: more retrievals generally increase cumulative chance
	retrievalBoost := float64(req.Retrievals-1) * 5.0

	// Egg source adjustment
	eggSourceAdjustment := 0.0
	if req.EggSource == "donor" {
		eggSourceAdjustment = 10.0 // Donor eggs typically have higher success rates
	}

	// Prior IVF cycles adjustment (negative impact)
	priorCyclePenalty := float64(req.PriorIvfCycles) * 3.0

	// Prior pregnancies adjustment (positive impact)
	pregnancyBoost := float64(req.PriorPregnancies) * 2.0

	// Calculate final chance
	chance := baseChance - agePenalty + retrievalBoost + eggSourceAdjustment - priorCyclePenalty + pregnancyBoost

	// Clamp between 1% and 75%
	if chance < 1.0 {
		chance = 1.0
	}
	if chance > 75.0 {
		chance = 75.0
	}

	return CalculateResponse{
		CumulativeChancePercent: chance,
		Notes: []string{
			"Calculations are illustrative for this demo and based on simplified calculations.",
			"Consult a physician for personalized assessment based on your specific medical history.",
			"Actual success rates may vary based on clinic, specific protocols, and individual factors.",
		},
	}
}
