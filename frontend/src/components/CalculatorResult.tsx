import type { CalculateResponse } from '../types/calculate'

interface CalculatorResultProps {
  result: CalculateResponse | null
  isLoading?: boolean
  error?: string | null
  retrievals: number
  onRetrievalsChange: (retrievals: number) => void
}

export default function CalculatorResult({
  result,
  isLoading = false,
  error = null,
  retrievals,
  onRetrievalsChange,
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
      <h2 className="mb-4 text-lg font-medium">Cumulative Chance of Live Birth*</h2>
      <p className="mb-2 text-xs text-gray-500">* After {retrievals} retrieval{retrievals > 1 ? 's' : ''} and all transfers within 12 months</p>

      <div className="mb-6">
        <div className="flex items-baseline gap-2">
          <span className="text-4xl font-bold text-blue-600">
            {result.cumulativeChancePercent.toFixed(1)}%
          </span>
        </div>
      </div>

      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">
          Explore by retrieval:
        </label>
        <div className="flex gap-2">
          {[1, 2, 3].map((num) => (
            <button
              key={num}
              onClick={() => onRetrievalsChange(num)}
              className={`px-4 py-2 rounded-md text-sm font-medium ${
                retrievals === num
                  ? 'bg-blue-600 text-white'
                  : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
              }`}
            >
              {num} Retrieval{num > 1 ? 's' : ''}
            </button>
          ))}
        </div>
      </div>

      {result.notes && result.notes.length > 0 && (
        <div className="mt-4 space-y-2">
          {result.notes.map((note, index) => (
            <p key={index} className="text-sm text-gray-600">
              {note}
            </p>
          ))}
        </div>
      )}

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
