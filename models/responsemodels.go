package models

type Employer struct {
	Name string
}

type Background struct {
	Name string
}

type GenderRatio struct {
	Male   string
	Female string
}

type StudentCount map[string]interface{}

type GraduatesCount struct {
	TotalGraduates string
}

type TotalProgramCost struct {
	TotalCost string
}

type PerCreditProgramCost struct {
	PerCreditCost string
}

type TopEmployers struct {
	Employers []string
}

type TopBackgrounds struct {
	Backgrounds []string
}

type DropOutRate struct {
	Rate string
}

type EmploymentRate struct {
	Rate string
}

type AverageSalary struct {
	AvgSalary string
}

type AcceptanceRate struct {
	Rate string
}