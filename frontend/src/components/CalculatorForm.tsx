import { useState } from 'react'
import { eggSources, priorIvfCyclesOptions } from '../types/calculate'
import type { CalculateFormData, CalculateReason, EggSource, FormErrors, PriorIvfCyclesOption } from '../types/calculate'

interface CalculatorFormProps {
  onSubmit: (data: CalculateFormData) => void
  isSubmitting?: boolean

}

export default function CalculatorForm({ onSubmit, isSubmitting = false }: CalculatorFormProps) {
  const [formData, setFormData] = useState<CalculateFormData>({
    age: '',
    weightLbs: '',
    heightFt: '',
    heightIn: '',
    priorIvfCycles: '',
    priorPregnancies: '',
    priorBirths: '',
    reasons: [],
    eggSource: '',
  })

  const [errors, setErrors] = useState<FormErrors>({})

  const validateField = (name: keyof CalculateFormData, value: string | CalculateReason[] | EggSource): string | undefined => {
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
    if (name === 'heightFt') {
      const heightFt = Number(value)
      if (isNaN(heightFt) || heightFt < 4 || heightFt > 6) {
        return 'Height feet must be between 4 and 7 feet'
      }
    }
    if (name === 'heightIn') {
      const heightIn = Number(value)
      if (isNaN(heightIn) || heightIn < 0 || heightIn > 11) {
        return 'Height inches must be between 0 and 12 inches'
      }
    }
    if (name === 'eggSource') {
      if (value === '' || !eggSources.includes(value as EggSource) ) {
        return 'Please select "My own eggs" or "Donor eggs"'
      }
    }
    if (name === 'priorIvfCycles') {
      if (!priorIvfCyclesOptions.includes(value as PriorIvfCyclesOption)) {
        return 'Please select "Yes" or "No"'
      }
    }
    if (name === 'priorPregnancies') {
      const pregnancies = Number(value)
      if (isNaN(pregnancies) || pregnancies < 0 || pregnancies > 2) {
        return 'Prior pregnancies must be 2+, 1, or None'
      }
    }
    if (name === 'priorBirths') {
      const births = Number(value)
      if (isNaN(births) || births < 0 || births > 2) {
        return 'Prior births must be 2+, 1, or None'
      }
    }
    if (name === 'reasons') {
      if (!Array.isArray(value) || value.length === 0) {
        return 'Please select at least one reason'
      }
    }
    return undefined
  }

  const handleChange = (name: keyof CalculateFormData, value: string | CalculateReason[] | EggSource) => {
    setFormData((prev) => ({ ...prev, [name]: value }))
    const error = validateField(name, value)
    setErrors((prev) => ({ ...prev, [name]: error }))
  }

  const handleReasonToggle = (reason: CalculateReason, or?: boolean) => {
    let newReasons
    if (or) {
      newReasons = [reason]
    } else {
      newReasons = formData.reasons.includes(reason)
      ? formData.reasons.filter((r) => r !== reason )
      : [...formData.reasons, reason]
      newReasons = newReasons.filter((r) => !['unexplained', 'unknown'].includes(r))
    }
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

  const reasons: { value: CalculateReason; label: string, or?: boolean }[] = [
    { value: 'male_factor_infertility', label: 'Male factor infertility' },
    { value: 'endometriosis', label: 'Endometriosis' },
    { value: 'tubal_factor', label: 'Tubal factor' },
    { value: 'ovulatory_disorder', label: 'Ovulatory disorder (including PCOS)' },
    { value: 'diminished_ovarian_reserve', label: 'Diminished ovarian reserve' },
    { value: 'uterine_factor', label: 'Uterine factor' },
    { value: 'other', label: 'Other reason' },
    { value: 'unexplained', label: 'Unexplained (Idiopathic) infertility', or: true },
    { value: 'unknown', label: "I don't know/no reason", or: true },
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
        <label className="block text-sm font-medium text-gray-700 mb-1">How tall are you?</label>
        <div className="flex space-x-4">
          <div className="flex-1">
            <input
              type="number"
              id="heightFeet"
              min="4"
              max="6"
              placeholder="Feet"
              value={formData.heightFt}
              onChange={(e) => handleChange('heightFt', e.target.value)}
              className={`w-full px-3 py-2 border rounded-md ${errors.heightFt ? 'border-red-500' : 'border-gray-300'}`}
            />
          </div>

          <div className="flex-1">
            <input
              type="number"
              id="heightInches"
              min="0"
              max="11"
              placeholder="Inches"
              value={formData.heightIn}
              onChange={(e) => handleChange('heightIn', e.target.value)}
              className={`w-full px-3 py-2 border rounded-md ${errors.heightIn ? 'border-red-500' : 'border-gray-300'}`}
            />
          </div>
        </div>
        {errors.heightFt && <p className="mt-1 text-sm text-red-600">{errors.heightFt}</p>}
        {errors.heightIn && <p className="mt-1 text-sm text-red-600">{errors.heightIn}</p>}
      </div>

      <div>
        <label htmlFor="eggSource" className="block text-sm font-medium text-gray-700 mb-1">
          Do you plan to use your own eggs or donor eggs?
        </label>
        <div className="space-x-5 flex">
          <label className="flex items-center">
            <input type="radio" id="eggSourceOwn" name="eggSource" value="own" checked={formData.eggSource === 'own'} onChange={(e) => handleChange('eggSource', e.target.value)} className="mr-2" />
            <span className="text-sm">My own eggs</span>
          </label>
          <label className="flex items-center">
            <input type="radio" id="eggSourceDonor" name="eggSource" value="donor" checked={formData.eggSource === 'donor'} onChange={(e) => handleChange('eggSource', e.target.value)} className="mr-2" />
            <span className="text-sm">Donor eggs</span>
          </label>
        </div>
        {errors.eggSource && <p className="mt-1 text-sm text-red-600">{errors.eggSource}</p>}
      </div>

      {formData.eggSource === 'own' && (
        <div>
          <label htmlFor="priorIvfCycles" className="block text-sm font-medium text-gray-700 mb-1">
            Have you used IVF in the past?
          </label>
          <div className="space-x-5 flex">
            <label className="flex items-center">
              <input type="radio" id="priorIvfCyclesYes" name="priorIvfCycles" value="yes" checked={formData.priorIvfCycles === 'yes'} onChange={(e) => handleChange('priorIvfCycles', e.target.value)} className="mr-2" />
              <span className="text-sm">Yes</span>
            </label>
            <label className="flex items-center">
              <input type="radio" id="priorIvfCyclesNo" name="priorIvfCycles" value="no" checked={formData.priorIvfCycles === 'no'} onChange={(e) => handleChange('priorIvfCycles', e.target.value)} className="mr-2" />
              <span className="text-sm">No, I've never used IVF</span>
            </label>
          </div>
          {errors.priorIvfCycles && <p className="mt-1 text-sm text-red-600">{errors.priorIvfCycles}</p>}
        </div>
      )}

      <div>
        <label htmlFor="priorPregnancies" className="block text-sm font-medium text-gray-700 mb-1">
          How many prior pregnancies have you had?
        </label>
        <div className="space-x-5 flex">
          <label className="flex items-center">
            <input type="radio" id="priorPregnancies2" name="priorPregnancies" value="2" checked={formData.priorPregnancies === '2'} onChange={(e) => handleChange('priorPregnancies', e.target.value)} className="mr-2" />
            <span className="text-sm">2 or more</span>
          </label>
          <label className="flex items-center">
            <input type="radio" id="priorPregnancies1" name="priorPregnancies" value="1" checked={formData.priorPregnancies === '1'} onChange={(e) => handleChange('priorPregnancies', e.target.value)} className="mr-2" />
            <span className="text-sm">1</span>
          </label>
          <label className="flex items-center">
            <input type="radio" id="priorPregnancies0" name="priorPregnancies" value="0" checked={formData.priorPregnancies === '0'} onChange={(e) => handleChange('priorPregnancies', e.target.value)} className="mr-2" />
            <span className="text-sm">None</span>
          </label>
        </div>
        {errors.priorPregnancies && <p className="mt-1 text-sm text-red-600">{errors.priorPregnancies}</p>}
      </div>

      <div>

        <label htmlFor="priorBirths" className="block text-sm font-medium text-gray-700 mb-1">
          How many prior births have you had?
        </label>
        <div className="space-x-5 flex">
          <label className="flex items-center">
            <input type="radio" id="priorBirths2" name="priorBirths" value="2" checked={formData.priorBirths === '2'} onChange={(e) => handleChange('priorBirths', e.target.value)} className="mr-2" />
            <span className="text-sm">2 or more</span>
          </label>
          <label className="flex items-center">
            <input type="radio" id="priorBirths1" name="priorBirths" value="1" checked={formData.priorBirths === '1'} onChange={(e) => handleChange('priorBirths', e.target.value)} className="mr-2" />
            <span className="text-sm">1</span>
          </label>
          <label className="flex items-center">
            <input type="radio" id="priorBirths0" name="priorBirths" value="0" checked={formData.priorBirths === '0'}  onChange={(e) => handleChange('priorBirths', e.target.value)} className="mr-2" />
            <span className="text-sm">None</span>
          </label>
        </div>
        {errors.priorBirths && <p className="mt-1 text-sm text-red-600">{errors.priorBirths}</p>}
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-2">
          What is the reason you are using IVF? (select all that apply)
        </label>
        <div className="space-y-2">
          {reasons.map((reason) => (
            <div>
              {reason.or && <div className="mb-2">(or)</div>}
              <label key={reason.value} className="flex items-center">
                <input
                  type="checkbox"
                  checked={formData.reasons.includes(reason.value)}
                  onChange={() => handleReasonToggle(reason.value, reason.or)}
                  className="mr-2"
                />
                <span className="text-sm">{reason.label}</span>
              </label>
              </div>
          ))}
        </div>
        {errors.reasons && <p className="mt-1 text-sm text-red-600">{errors.reasons}</p>}
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
