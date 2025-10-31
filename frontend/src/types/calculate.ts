export type CalculateReason =
  | 'male_factor'
  | 'endometriosis'
  | 'tubal_factor'
  | 'ovulatory_disorder'
  | 'diminished_ovarian_reserve'
  | 'uterine_factor'
  | 'other'
  | 'unexplained'
  | 'unknown'

export type EggSource = 'own' | 'donor'

export interface CalculateRequest {
  age: number
  weightLbs: number
  heightIn: number
  priorIvfCycles: number
  priorPregnancies: number
  priorBirths: number
  reasons: CalculateReason[]
  eggSource: EggSource
}

export interface CalculateResponse {
  cumulativeChancePercent: number
  notes: string[]
}

export interface CalculateFormData {
  age: string
  weightLbs: string
  heightIn: string
  priorIvfCycles: string
  priorPregnancies: string
  priorBirths: string
  reasons: CalculateReason[]
  eggSource: EggSource
}

export interface FormErrors {
  age?: string
  weightLbs?: string
  heightIn?: string
  priorIvfCycles?: string
  priorPregnancies?: string
  priorBirths?: string
  reasons?: string
  eggSource?: string
}
