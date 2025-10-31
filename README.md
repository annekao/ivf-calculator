# IVF Success Calculator

A full-stack application to calculate the chance of having a live birth using in vitro fertilization (IVF). Built with React + TypeScript (frontend) and Go with Gin (backend).

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
  "eggSource": "own",
  "retrievals": 1
}
```

**Response:**
```json
{
  "cumulativeChancePercent": 37.5,
  "notes": [
    "Calculations are illustrative for this demo and based on simplified calculations.",
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
- `retrievals`: 1-3

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

## Notes

- The calculation logic in the backend is a simplified mock for demonstration purposes.
- In a production implementation, you would integrate with actual CDC data and statistical models.
- This tool does not provide medical advice. Always consult with a healthcare provider.

## License

This project is a demonstration/take-home assessment and is not intended for production medical use.
