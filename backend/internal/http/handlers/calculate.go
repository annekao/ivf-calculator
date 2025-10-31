package handlers

import (
	"net/http"

	"ivf-calculator-backend/internal/calculator"
	"ivf-calculator-backend/internal/validation"

	"github.com/gin-gonic/gin"
)

// CalculateRequest and CalculateResponse are now defined in the calculator package
// Import them for convenience
type CalculateRequest = calculator.CalculateRequest
type CalculateResponse = calculator.CalculateResponse

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
