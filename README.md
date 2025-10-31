# IVF Success Calculator

A full-stack application to calculate the chance of having a live birth using in vitro fertilization (IVF). Built with React + TypeScript (frontend) and Go with Gin (backend). Uses CDC statistical models based on logit regression formulas for accurate IVF success rate calculations.

## Project Structure

```
ivf-calculator/
├── frontend/          # React + TypeScript + Vite + Tailwind
├── backend/           # Go + Gin
└── README.md
```

## Prerequisites

- Node.js 18+ and npm
- Go 1.22+

## Getting Started

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the server:
   ```bash
   go run ./cmd/server
   ```

   The backend will start on `http://localhost:8080` by default.

   You can set a custom port using the `PORT` environment variable:
   ```bash
   PORT=3000 go run ./cmd/server
   ```

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```

   The frontend will start on `http://localhost:5173` by default.

4. Configure API base URL (optional):
   Create a `.env` file in the frontend directory:
   ```
   VITE_API_BASE=http://localhost:8080
   ```

### Running Both Services

Open two terminal windows:

**Terminal 1 (Backend):**
```bash
cd backend
go run ./cmd/server
```

**Terminal 2 (Frontend):**
```bash
cd frontend
npm run dev
```

Then open your browser to `http://localhost:5173`.

## API Endpoints

### `GET /healthz`
Health check endpoint.

**Response:**
```json
{
  "status": "ok"
}
```

### `POST /api/calculate`
Calculate IVF success probability.

**Request Body:**
```json
{
  "age": 34,
  "weightLbs": 150,
  "heightIn": 66,
  "priorIvfCycles": 0,
  "priorPregnancies": 0,
  "priorBirths": 0,
  "reasons": ["male_factor", "unexplained"],
  "eggSource": "own"
}
```

**Response:**
```json
{
  "cumulativeChancePercent": 37.5,
  "notes": [
    "Calculations are based on CDC statistical models for IVF success rates.",
    "Consult a physician for personalized assessment based on your specific medical history.",
    "Actual success rates may vary based on clinic, specific protocols, and individual factors."
  ]
}
```

**Validation Rules:**
- `age`: 20-50
- `weightLbs`: 80-300
- `heightIn`: 55-78
- `priorIvfCycles`: 0-3
- `priorPregnancies`: 0-2
- `priorBirths`: 0-2
- `reasons`: Array of valid reason strings (at least one required)
- `eggSource`: "own" or "donor"

## Development

### Building for Production

**Frontend:**
```bash
cd frontend
npm run build
```

**Backend:**
```bash
cd backend
go build -o server ./cmd/server
```

## Testing

The backend includes comprehensive tests for the calculator using CDC formulas. The test suite covers multiple scenarios and validates formula selection and calculation accuracy.

### Running Tests

**Run all tests:**
```bash
cd backend
go test ./internal/calculator -v
```

**Run specific test scenarios:**
```bash
# Scenario 1: Own eggs, no prior IVF, known reason (endometriosis + ovulatory disorder)
go test ./internal/calculator -v -run TestCalculate_OwnEggs_NoPriorIVF_KnownReason_Scenario1

# Scenario 2: Own eggs, no prior IVF, unknown reason
go test ./internal/calculator -v -run TestCalculate_OwnEggs_NoPriorIVF_UnknownReason_Scenario2

# Scenario 3: Own eggs, prior IVF, known reason (tubal factor + diminished ovarian reserve)
go test ./internal/calculator -v -run TestCalculate_OwnEggs_PriorIVF_KnownReason_Scenario3

# Run comparison test (all scenarios together)
go test ./internal/calculator -v -run TestAllScenarios_Comparison

# Test formula loading
go test ./internal/calculator -v -run TestFormulaLoading

# Test formula matching logic
go test ./internal/calculator -v -run TestFindMatchingFormula
```

### Test Coverage

The test suite includes:

1. **Scenario Tests**: Tests for three specific patient scenarios:
   - Own eggs, no prior IVF attempts, known infertility reasons
   - Own eggs, no prior IVF attempts, unknown infertility reasons
   - Own eggs, prior IVF attempts, known infertility reasons

2. **Formula Matching Tests**: Validates that the correct CDC formula is selected based on:
   - Egg source (own vs donor)
   - Prior IVF attempts
   - Whether infertility reason is known

3. **Formula Loading Tests**: Verifies that formulas are successfully loaded from the CSV file

4. **Result Validation**: Ensures calculated success rates are within expected bounds (0.1% - 95%)

5. **Comparison Tests**: Runs all scenarios together to compare results

Each test validates that:
- The calculation produces a valid result
- The correct formula is selected for the given patient parameters
- Results are within reasonable probability bounds
- The calculation uses actual CDC formula coefficients

## Notes

- The calculation logic uses CDC statistical models based on logit regression formulas. Formulas are loaded from `backend/internal/calculator/ivf_success_formulas.csv` and selected based on patient parameters (egg source, prior IVF attempts, known infertility reasons).
- The calculator considers factors including age, BMI, infertility reasons, prior pregnancies, prior live births, and number of retrievals.
- This tool does not provide medical advice. Always consult with a healthcare provider.

## License

This project is a demonstration/take-home assessment and is not intended for production medical use.


## Non AI-generated notes 

- Used Cursor and gave it the CDC website and CSV file to code a "simple" Go/React app (used plan mode)
   - I *didn't* copy the README instructions directly. This was mostly out of curiosity to see how powerful Cursor is with minimal prompting.
- Removed docker since this is a take-home assignment and doesn't need to be deployed or put in a production environment
- My first task was to try to spin up the app before checking the details of the code
   - Since I'm unfamiliar with Go, I relied on Cursor to resolve errors
   - 1 circular dependency error later and the app is up and running!
- Type error coming from endpoint, shifted gears togit having Cursor write tests given the 3 examples in the README without giving it the expected answers
   - Specific tests are not precise, just vaguely checks if its within "bounds"
- Cursor added an unnecessary "retrievals" (int) field which it reasoned as "The CDC formulas give per-cycle probability. 'Retrievals' computes cumulative probability across multiple cycles using P(at least one success) = 1 - (1 - p)^n."
   - Asked to remove since it is not part of the product brief


