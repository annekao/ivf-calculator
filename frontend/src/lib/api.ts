import type { CalculateRequest, CalculateResponse } from '../types/calculate'

const API_BASE = import.meta.env.VITE_API_BASE || 'http://localhost:8080'

export async function calculate(
  request: CalculateRequest
): Promise<CalculateResponse> {
  const response = await fetch(`${API_BASE}/api/calculate`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(request),
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.details ? JSON.stringify(error.details) : 'Request failed')
  }

  return response.json()
}
