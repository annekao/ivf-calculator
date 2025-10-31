import { useState } from 'react'
import CalculatorForm from './components/CalculatorForm'
import CalculatorResult from './components/CalculatorResult'
import { calculate } from './lib/api'
import type { CalculateFormData, CalculateRequest, CalculateResponse } from './types/calculate'

export default function App() {
  const [result, setResult] = useState<CalculateResponse | null>(null)
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const handleSubmit = async (formData: CalculateFormData) => {
    setIsLoading(true)
    setError(null)

    try {
      const request: CalculateRequest = {
        age: Number(formData.age),
        weightLbs: Number(formData.weightLbs),
        heightIn: Number(formData.heightIn),
        priorIvfCycles: Number(formData.priorIvfCycles),
        priorPregnancies: Number(formData.priorPregnancies),
        priorBirths: Number(formData.priorBirths),
        reasons: formData.reasons,
        eggSource: formData.eggSource,
      }

      const response = await calculate(request)
      setResult(response)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred')
      setResult(null)
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <header className="border-b bg-white shadow-sm">
        <div className="mx-auto max-w-5xl px-4 py-6">
          <h1 className="text-2xl font-semibold text-gray-900">
            IVF Success Calculator
          </h1>
          <p className="mt-2 text-sm text-gray-600">
            Calculate your chance of having a live birth using in vitro fertilization (IVF).
          </p>
        </div>
      </header>
      <main className="mx-auto max-w-5xl px-4 py-8">
        <div className="grid gap-8 md:grid-cols-2">
          <section className="rounded-lg border bg-white p-6 shadow-sm">
            <h2 className="mb-4 text-lg font-medium text-gray-900">
              Background and History
            </h2>
            <CalculatorForm onSubmit={handleSubmit} isSubmitting={isLoading} />
          </section>
          <section className="rounded-lg border bg-white p-6 shadow-sm">
            <CalculatorResult
              result={result}
              isLoading={isLoading}
              error={error}
            />
          </section>
        </div>
        <div className="mt-6 rounded-lg border border-gray-200 bg-white p-4">
          <p className="text-xs text-gray-500">
            <strong>Disclaimer:</strong> The information you enter is not stored and is only used to calculate your chances of success. 
            The IVF Success Calculator does not provide medical advice, diagnosis, or treatment. These calculations may not reflect your 
            actual chances of success during ART treatment and are only being provided for informational purposes. Calculations are less 
            reliable at certain ranges and values of age, weight, height, and previous pregnancy and ART experiences. Please see your 
            doctor or healthcare provider for a personalized treatment plan.
          </p>
        </div>
      </main>
    </div>
  )
}