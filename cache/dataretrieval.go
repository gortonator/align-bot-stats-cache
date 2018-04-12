package cache

import (
	"log"
	"github.com/rapidclock/align-bot-stats-cache/models"
)

func GetGenderRatio() *models.GenderRatio {
	curRatio := new(models.GenderRatio)
	maleRatio, err := Get(GenderRatioMale)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	femaleRatio, err := Get(GenderRatioFemale)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	curRatio.Male = maleRatio
	curRatio.Female = femaleRatio
	return curRatio
}

func GetStudentCounts(location string) *models.StudentCount {
	counts := make(models.StudentCount)
	switch location {
	case "seattle":
		counts["Seattle"] = getSeattleCount()
	case "boston":
		counts["Boston"] = getBostonCount()
	case "charlotte":
		counts["Charlotte"] = getCharlotteCount()
	case "siliconValley":
		counts["SiliconValley"] = getSiliconValleyCount()
	default:
		counts["Seattle"] = getSeattleCount()
		counts["Boston"] = getBostonCount()
		counts["Charlotte"] = getCharlotteCount()
		counts["SiliconValley"] = getSiliconValleyCount()
	}
	return &counts
}

func getSeattleCount() string {
	count, err := Get(StudentCountSeattle)
	handleError(err)
	return count
}

func getBostonCount() string {
	count, err := Get(StudentCountBoston)
	handleError(err)
	return count
}

func getCharlotteCount() string {
	count, err := Get(StudentCountCharlotte)
	handleError(err)
	return count
}

func getSiliconValleyCount() string {
	count, err := Get(StudentCountSiliconValley)
	handleError(err)
	return count
}

func GetGraduatesCount() *models.GraduatesCount {
	gradNum := new(models.GraduatesCount)
	totNum, err := Get(NumberOfGraduatesTotal)
	handleError(err)
	gradNum.TotalGraduates = totNum
	return gradNum
}

func GetTotalCost() *models.TotalProgramCost {
	progCost := new(models.TotalProgramCost)
	cost, err := Get(ProgramCostTotal)
	handleError(err)
	progCost.TotalCost = cost
	return progCost
}

func GetPerCreditCost() *models.PerCreditProgramCost {
	perCredCost := new(models.PerCreditProgramCost)
	cost, err := Get(ProgramCostPerCredit)
	handleError(err)
	perCredCost.PerCreditCost = cost
	return perCredCost
}

func GetTopEmployers(k int) *models.TopEmployers {
	topEmpData := GetFromUnorderedSet(TopEmployers)
	topEmp := new(models.TopEmployers)
	topEmp.Employers = topEmpData[:k]
	return topEmp
}

func GetTopBackgrounds(k int) *models.TopBackgrounds {
	topBgData := GetFromUnorderedSet(TopBackgrounds)
	topBg := new(models.TopBackgrounds)
	topBg.Backgrounds = topBgData[:k]
	return topBg
}

func GetDropOutRate() *models.DropOutRate {
	doRate := new(models.DropOutRate)
	rate, err := Get(DropOutRate)
	handleError(err)
	doRate.Rate = rate
	return doRate
}

func GetEmploymentRate() *models.EmploymentRate {
	emplRate := new(models.EmploymentRate)
	rate, err := Get(EmploymentRate)
	handleError(err)
	emplRate.Rate = rate
	return emplRate
}

func GetAverageSalary() *models.AverageSalary {
	avgSal := new(models.AverageSalary)
	sal, err := Get(SalaryAvg)
	handleError(err)
	avgSal.AvgSalary = sal
	return avgSal
}

func GetAcceptanceRate() *models.AcceptanceRate {
	accRate := new(models.AcceptanceRate)
	rate, err := Get(AcceptanceRates)
	handleError(err)
	accRate.Rate = rate
	return accRate
}
