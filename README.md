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
  "eggSource": "own"
  "priorIvfCycles": "no",
  "priorPregnancies": 0,
  "priorBirths": 0,
  "reasons": ["male_factor_infertility", "unexplained"],
}
```

**Response:**
```json
{
  "cumulativeChancePercent": 37.5
}
```

**Validation Rules:**
- `age`: 20-50
- `weightLbs`: 80-300
- `heightIn`: 55-78
- `eggSource`: "own" or "donor"
- `priorIvfCycles`: "yes" or "no" required when using 'own' eggs
- `priorPregnancies`: 0-2
- `priorBirths`: 0-2
- `reasons`: Array of valid reason strings (at least one required)

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
go test ./internal/validation -v
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

Each test validates that:
- The calculation produces correct result
- The correct formula is selected for the given patient parameters
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
   - Gave me a couple different of stack options, went with what felt the most low weight and familiar
   - After I moved the files, I unfortunately lost the history of this and plan mode doesn't auto-save
- Removed docker since this is a take-home assignment and doesn't need to be deployed or put in a production environment
- My first task was to try to spin up the app before checking the details of the code
   - Since I'm unfamiliar with Go, I relied on Cursor to resolve errors
   - 1 circular dependency error later and the app is up and running!
     <img width="552" height="689" alt="Screenshot 2025-10-31 at 3 36 26 PM" src="https://github.com/user-attachments/assets/415a6665-94ce-45fc-a758-f99f38d37d77" />
- Type error coming from endpoint, shifted gears to having Cursor write tests given the 3 examples in the README without giving it the expected answers
   - Specific tests are not precise, just vaguely checks if its within "bounds"
   - Tests had a lot of fluff and logging that I ended up removing since it's not realistically something I would deploy (i.e. checking if it's between 0 and 100)
   - Removed "Test All Scenarios" as that was redundant
   - Liked how it tested the correct formula was used
   - Generally tests are not asserting, but instead throwing errors - unclear if this is Go convention or not
- Cursor added unnecessary 'Notes' to CalculateResponse
   - Asked to remove since it is not part of the product brief
   - Kept as CalculateResponse as a struct even though it doesn't need to be since there is only one value
- Hit the free tier monthly limit after the above
- Fix BMI and age formula, remove redundant test checks and unnecessary code, remove mock calculator
- Fix front-end/make it look better (i.e. weird dark mode theme that didn't look right, use radio instead of select)
- Change Prior IVF cycles from number to string since product brief says that it can be true/false/nan
- Split out inches into feet and inches
- Nest IVF cycles q under Egg q
- Add # of live births constraint / toggling of available options
- Set up toggling for unexplained/unknown reason
- Removed Retrievals field (initially wasn't sure why it was added in, but then realized that it was part of the CDC calculation after submitting)
   - Out of scope / no explicit instruction on how to calculate it so I didn't want to just go along with whatever Cursor said
- Add more tests since examples given only cover the first 3 formulas
- Lots of edge case handling with what can be empty and acceptable values (pregnancies vs births)
- ChatGPT test for validation/calculate_request.go

Overall most manual coding was spent with edge cases / error-handling and getting the front-end dynamic to my liking.

### Things I would like to add or spend more time on
- front-end / e2e testing
- error-handling testing
- more interesting / less bland UI
- separate Formula loading from the Calculator 
