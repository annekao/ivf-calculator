import { useState } from 'react'
import type { CalculateFormData, CalculateReason, EggSource, FormErrors } from '../types/calculate'

interface CalculatorFormProps {
  onSubmit: (data: CalculateFormData) => void
  isSubmitting?: boolean
}

export default function CalculatorForm({ onSubmit, isSubmitting = false }: CalculatorFormProps) {
  const [formData, setFormData] = useState<CalculateFormData>({
    age: '',
    weightLbs: '',
    heightIn: '',
    priorIvfCycles: '',
    priorPregnancies: '',
    priorBirths: '',
    reasons: [],
    eggSource: 'own',
  })

  const [errors, setErrors] = useState<FormErrors>({})

  const validateField = (name: keyof CalculateFormData, value: string | number | CalculateReason[] | EggSource): string | undefined => {
    if (name === 'age') {
      const age = Number(value)
      if (isNaN(age) || age < 20 || age > 50) {
        return 'Age must be between 20 and 50'
      }
    }
    if (name === 'weightLbs') {
      const weight = Number(value)
      if (isNaN(weight) || weight < 80 || weight > 300) {
        return 'Weight must be between 80 and 300 lbs'
      }
    }
    if (name === 'heightIn') {
      const height = Number(value)
      if (isNaN(height) || height < 55 || height > 78) {
        return 'Height must be between 55 and 78 inches'
      }
    }
    if (name === 'priorIvfCycles') {
      const cycles = Number(value)
      if (isNaN(cycles) || cycles < 0 || cycles > 3) {
        return 'Prior IVF cycles must be 0-3'
      }
    }
    if (name === 'priorPregnancies') {
      const pregnancies = Number(value)
      if (isNaN(pregnancies) || pregnancies < 0 || pregnancies > 2) {
        return 'Prior pregnancies must be 0, 1, or 2+'
      }
    }
    if (name === 'priorBirths') {
      const births = Number(value)
      if (isNaN(births) || births < 0 || births > 2) {
        return 'Prior births must be 0, 1, or 2+'
      }
    }
    if (name === 'reasons') {
      if (!Array.isArray(value) || value.length === 0) {
        return 'Please select at least one reason'
      }
    }
    return undefined
  }

  const handleChange = (name: keyof CalculateFormData, value: string | number | CalculateReason[] | EggSource) => {
    setFormData((prev) => ({ ...prev, [name]: value }))
    const error = validateField(name, value)
    setErrors((prev) => ({ ...prev, [name]: error }))
  }

  const handleReasonToggle = (reason: CalculateReason) => {
    const newReasons = formData.reasons.includes(reason)
      ? formData.reasons.filter((r) => r !== reason)
      : [...formData.reasons, reason]
    handleChange('reasons', newReasons)
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    
    // Validate all fields
    const newErrors: FormErrors = {}
    Object.keys(formData).forEach((key) => {
      const error = validateField(key as keyof CalculateFormData, formData[key as keyof CalculateFormData])
      if (error) {
        newErrors[key as keyof FormErrors] = error
      }
    })

    if (Object.keys(newErrors).length > 0) {
      setErrors(newErrors)
      return
    }

    onSubmit(formData)
  }

  const reasons: { value: CalculateReason; label: string }[] = [
    { value: 'male_factor', label: 'Male factor infertility' },
    { value: 'endometriosis', label: 'Endometriosis' },
    { value: 'tubal_factor', label: 'Tubal factor' },
    { value: 'ovulatory_disorder', label: 'Ovulatory disorder (including PCOS)' },
    { value: 'diminished_ovarian_reserve', label: 'Diminished ovarian reserve' },
    { value: 'uterine_factor', label: 'Uterine factor' },
    { value: 'other', label: 'Other reason' },
    { value: 'unexplained', label: 'Unexplained (Idiopathic) infertility' },
    { value: 'unknown', label: "I don't know/no reason" },
  ]

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div>
        <label htmlFor="age" className="block text-sm font-medium text-gray-700 mb-1">
          How old are you?
        </label>
        <input
          type="number"
          id="age"
          min="20"
          max="50"
          value={formData.age}
          onChange={(e) => handleChange('age', e.target.value)}
          className={`w-full px-3 py-2 border rounded-md ${errors.age ? 'border-red-500' : 'border-gray-300'}`}
        />
        {errors.age && <p className="mt-1 text-sm text-red-600">{errors.age}</p>}
      </div>

      <div>
        <label htmlFor="weightLbs" className="block text-sm font-medium text-gray-700 mb-1">
          How much do you weigh? (lbs)
        </label>
        <input
          type="number"
          id="weightLbs"
          min="80"
          max="300"
          value={formData.weightLbs}
          onChange={(e) => handleChange('weightLbs', e.target.value)}
          className={`w-full px-3 py-2 border rounded-md ${errors.weightLbs ? 'border-red-500' : 'border-gray-300'}`}
        />
        {errors.weightLbs && <p className="mt-1 text-sm text-red-600">{errors.weightLbs}</p>}
      </div>

      <div>
        <label htmlFor="heightIn" className="block text-sm font-medium text-gray-700 mb-1">
          How tall are you? (inches)
        </label>
        <input
          type="number"
          id="heightIn"
          min="55"
          max="78"
          value={formData.heightIn}
          onChange={(e) => handleChange('heightIn', e.target.value)}
          className={`w-full px-3 py-2 border rounded-md ${errors.heightIn ? 'border-red-500' : 'border-gray-300'}`}
        />
        {errors.heightIn && <p className="mt-1 text-sm text-red-600">{errors.heightIn}</p>}
      </div>

      <div>
        <label htmlFor="priorIvfCycles" className="block text-sm font-medium text-gray-700 mb-1">
          How many times have you used IVF in the past?
        </label>
        <select
          id="priorIvfCycles"
          value={formData.priorIvfCycles}
          onChange={(e) => handleChange('priorIvfCycles', e.target.value)}
          className={`w-full px-3 py-2 border rounded-md ${errors.priorIvfCycles ? 'border-red-500' : 'border-gray-300'}`}
        >
          <option value="">-- select an option --</option>
          <option value="0">I've never used IVF</option>
          <option value="1">1</option>
          <option value="2">2</option>
          <option value="3">3 or more</option>
        </select>
        {errors.priorIvfCycles && <p className="mt-1 text-sm text-red-600">{errors.priorIvfCycles}</p>}
      </div>

      <div>
        <label htmlFor="priorPregnancies" className="block text-sm font-medium text-gray-700 mb-1">
          How many prior pregnancies have you had?
        </label>
        <select
          id="priorPregnancies"
          value={formData.priorPregnancies}
          onChange={(e) => handleChange('priorPregnancies', e.target.value)}
          className={`w-full px-3 py-2 border rounded-md ${errors.priorPregnancies ? 'border-red-500' : 'border-gray-300'}`}
        >
          <option value="">-- select an option --</option>
          <option value="0">None</option>
          <option value="1">1</option>
          <option value="2">2 or more</option>
        </select>
        {errors.priorPregnancies && <p className="mt-1 text-sm text-red-600">{errors.priorPregnancies}</p>}
      </div>

      <div>
        <label htmlFor="priorBirths" className="block text-sm font-medium text-gray-700 mb-1">
          How many prior births have you had?
        </label>
        <select
          id="priorBirths"
          value={formData.priorBirths}
          onChange={(e) => handleChange('priorBirths', e.target.value)}
          className={`w-full px-3 py-2 border rounded-md ${errors.priorBirths ? 'border-red-500' : 'border-gray-300'}`}
        >
          <option value="">-- select an option --</option>
          <option value="0">None</option>
          <option value="1">1</option>
          <option value="2">2 or more</option>
        </select>
        {errors.priorBirths && <p className="mt-1 text-sm text-red-600">{errors.priorBirths}</p>}
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-2">
          What is the reason you are using IVF? (select all that apply)
        </label>
        <div className="space-y-2">
          {reasons.map((reason) => (
            <label key={reason.value} className="flex items-center">
              <input
                type="checkbox"
                checked={formData.reasons.includes(reason.value)}
                onChange={() => handleReasonToggle(reason.value)}
                className="mr-2"
              />
              <span className="text-sm">{reason.label}</span>
            </label>
          ))}
        </div>
        {errors.reasons && <p className="mt-1 text-sm text-red-600">{errors.reasons}</p>}
      </div>

      <div>
        <label htmlFor="eggSource" className="block text-sm font-medium text-gray-700 mb-1">
          Do you plan to use your own eggs or donor eggs?
        </label>
        <select
          id="eggSource"
          value={formData.eggSource}
          onChange={(e) => handleChange('eggSource', e.target.value as EggSource)}
          className={`w-full px-3 py-2 border rounded-md ${errors.eggSource ? 'border-red-500' : 'border-gray-300'}`}
        >
          <option value="own">My own eggs</option>
          <option value="donor">Donor eggs</option>
        </select>
        {errors.eggSource && <p className="mt-1 text-sm text-red-600">{errors.eggSource}</p>}
      </div>

      <button
        type="submit"
        disabled={isSubmitting}
        className="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        {isSubmitting ? 'Calculating...' : 'Calculate Success'}
      </button>
    </form>
  )
}
