package handlers

import (
	"net/http"

	"ivf-calculator-backend/internal/calculator"
	"ivf-calculator-backend/internal/validation"

	"github.com/gin-gonic/gin"
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

// PostCalculate handles POST /api/calculate requests
func PostCalculate(c *gin.Context) {
	var req CalculateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request format",
			"details": err.Error(),
		})
		return
	}

	// Validate the request
	if errors := validation.ValidateCalculateRequest(req); len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errors,
		})
		return
	}

	// Calculate the result
	result := calculator.Calculate(req)

	c.JSON(http.StatusOK, result)
}
