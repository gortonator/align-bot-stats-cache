package cache

func InitializeWithStats() {
	initGenderRatio()
	initStudentNumbers()
	initNumOfGraduates()
	initProgramCost()
	initTopEmployers()
	initTopBackgrounds()
	initDropOutRate()
	initEmploymentRate()
	initAvgSalary()
	initAcceptanceRates()
}

func initGenderRatio() {
	var maleRatio = "55"
	var femaleRatio = "45"

	Post(GenderRatioMale, maleRatio)
	Post(GenderRatioFemale, femaleRatio)
}

func initStudentNumbers() {
	seaCount := "300"
	bosCount := "50"
	ctlCount := "25"
	silCount := "43"
	totalCount := "418"

	Post(StudentCountTotal, totalCount)
	Post(StudentCountSeattle, seaCount)
	Post(StudentCountBoston, bosCount)
	Post(StudentCountCharlotte, ctlCount)
	Post(StudentCountSiliconValley, silCount)
}

func initNumOfGraduates() {
	totalNumGrads := "175"

	Post(NumberOfGraduatesTotal, totalNumGrads)
}

func initProgramCost() {
	totalCost := "72000"
	perCreditCost := "1650"

	Post(ProgramCostTotal, totalCost)
	Post(ProgramCostPerCredit, perCreditCost)
}

func initTopEmployers() {
	topEmplrs := []string{"Amazon", "Google", "Facebook", "Zillow", "Apple"}

	AddToUnorderedSet(TopEmployers, topEmplrs...)
}

func initTopBackgrounds() {
	topBgs := []string{"Criminal Justice", "Arts", "Math", "Chemistry", "Physics"}

	AddToUnorderedSet(TopBackgrounds, topBgs...)
}

func initDropOutRate() {
	rate := "4"

	Post(DropOutRate, rate)
}

func initEmploymentRate() {
	rate := "94"

	Post(EmploymentRate, rate)
}

func initAvgSalary() {
	avgSal := "5000"

	Post(SalaryAvg, avgSal)
}

func initAcceptanceRates() {
	rate := "23"

	Post(AcceptanceRates, rate)
}
