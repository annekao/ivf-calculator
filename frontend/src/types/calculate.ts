export type CalculateReason =
  | 'male_factor_infertility'
  | 'endometriosis'
  | 'tubal_factor'
  | 'ovulatory_disorder'
  | 'diminished_ovarian_reserve'
  | 'uterine_factor'
  | 'other'
  | 'unexplained'
  | 'unknown'

export const eggSources = ['own', 'donor', ''] as const
export type EggSource = typeof eggSources[number]

export const priorIvfCyclesOptions = ["yes", "no", ''] as const
export type PriorIvfCyclesOption = typeof priorIvfCyclesOptions[number]

export interface CalculateRequest {
  age: number
  weightLbs: number
  heightFt: number
  heightIn: number
  eggSource: EggSource
  priorIvfCycles: PriorIvfCyclesOption
  priorPregnancies: number
  priorBirths: number
  reasons: CalculateReason[]
}

export interface CalculateResponse {
  cumulativeChancePercent: number
}

export interface CalculateFormData {
  age: string
  weightLbs: string
  heightFt: string
  heightIn: string
  eggSource: EggSource
  priorIvfCycles: PriorIvfCyclesOption
  priorPregnancies: string
  priorBirths: string
  reasons: CalculateReason[]
}

export interface FormErrors {
  age?: string
  weightLbs?: string
  heightFt?: string
  heightIn?: string
  eggSource?: string
  priorIvfCycles?: string
  priorPregnancies?: string
  priorBirths?: string
  reasons?: string
}
