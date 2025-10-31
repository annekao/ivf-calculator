import type { CalculateResponse } from '../types/calculate'

interface CalculatorResultProps {
  result: CalculateResponse | null
  isLoading?: boolean
  error?: string | null
}

export default function CalculatorResult({
  result,
  isLoading = false,
  error = null,
}: CalculatorResultProps) {
  if (isLoading) {
    return (
      <div className="rounded-lg border p-6">
        <h2 className="mb-4 text-lg font-medium">Calculating...</h2>
        <p className="text-sm text-gray-600">Please wait while we calculate your result.</p>
      </div>
    )
  }

  if (error) {
    return (
      <div className="rounded-lg border border-red-200 bg-red-50 p-6">
        <h2 className="mb-2 text-lg font-medium text-red-800">Error</h2>
        <p className="text-sm text-red-600">{error}</p>
      </div>
    )
  }

  if (!result) {
    return (
      <div className="rounded-lg border p-6">
        <h2 className="mb-4 text-lg font-medium">Cumulative Chance of Live Birth</h2>
        <p className="text-sm text-gray-600">
          Enter your information and click "Calculate Success" to see your result.
        </p>
      </div>
    )
  }

  return (
    <div className="rounded-lg border p-6">
      <h2 className="mb-4 text-lg font-medium">Chance of Live Birth</h2>

      <div className="mb-6">
        <div className="flex items-baseline gap-2">
          <span className="text-4xl font-bold text-blue-600">
            {result.cumulativeChancePercent.toFixed(1)}%
          </span>
        </div>
      </div>

      <div className="mt-4 border-t pt-4">
        <p className="text-xs text-gray-500">
          <strong>Disclaimer:</strong> The information you enter is not stored and is only used to calculate your chances of success. 
          The IVF Success Calculator does not provide medical advice, diagnosis, or treatment. These calculations may not reflect your 
          actual chances of success during ART treatment and are only being provided for informational purposes. Please see your 
          doctor or healthcare provider for a personalized treatment plan.
        </p>
      </div>
    </div>
  )
}
