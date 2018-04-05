package routehandlers

type RedisAppHandler struct {
	statsHandler *StatsHandler
}

type StatsHandler struct {
	genderRoleHandler         *GenderRoleHandler
	studentCountHandler       *StudentCountHandler
	graduatesCountHandler     *GraduatesCountHandler
	programCostHandler        *ProgramCostHandler
	employersHandler          *EmployersHandler
	studentBackgroundsHandler *StudentBackgroundsHandler
	dropOutRateHandler        *DropOutRateHandler
	studentEmpRateHandler     *StudentEmploymentRateHandler
	salaryHandler             *SalaryHandler
	acceptanceRateHandler     *AcceptanceRateHandler
}

type GenderRoleHandler struct{}
type StudentCountHandler struct{}
type GraduatesCountHandler struct{}
type ProgramCostHandler struct {
	totalCostHandler *TotalCostHandler
	perCreditHandler *PerCreditCostHandler
}
type EmployersHandler struct {
	topEmployersHandler *TopEmployersHandler
}
type StudentBackgroundsHandler struct {
	topBackgroundsHandler *TopBackgroundsHandler
}
type DropOutRateHandler struct{}
type StudentEmploymentRateHandler struct{}
type SalaryHandler struct {
	avgSalaryHandler *AvgSalaryHandler
}
type AcceptanceRateHandler struct{}

type TotalCostHandler struct{}
type PerCreditCostHandler struct{}
type TopEmployersHandler struct{}
type TopBackgroundsHandler struct{}
type AvgSalaryHandler struct{}

func NewRedisAppHandler() *RedisAppHandler {
	newApp := new(RedisAppHandler)
	newApp.statsHandler = newStatsHandler()
	return newApp
}

func newStatsHandler() *StatsHandler {
	handler := new(StatsHandler)
	handler.genderRoleHandler = newGenderRoleHandler()
	handler.studentCountHandler = newStudentCountHandler()
	handler.graduatesCountHandler = newGraduatesCountHandler()
	handler.programCostHandler = newProgramCostHandler()
	handler.employersHandler = newEmployersHandler()
	handler.studentBackgroundsHandler = newStudentBackgroundsHandler()
	handler.dropOutRateHandler = newDropOutRateHandler()
	handler.studentEmpRateHandler = newStudentEmploymentRateHandler()
	handler.salaryHandler = newSalaryHandler()
	handler.acceptanceRateHandler = newAcceptanceRateHandler()
	return handler
}

func newGenderRoleHandler() *GenderRoleHandler {
	handler := new(GenderRoleHandler)
	return handler
}

func newStudentCountHandler() *StudentCountHandler {
	handler := new(StudentCountHandler)
	return handler
}

func newGraduatesCountHandler() *GraduatesCountHandler {
	handler := new(GraduatesCountHandler)
	return handler
}

func newProgramCostHandler() *ProgramCostHandler {
	handler := new(ProgramCostHandler)
	handler.totalCostHandler = newTotalCostHandler()
	handler.perCreditHandler = newPerCreditCostHandler()
	return handler
}

func newTotalCostHandler() *TotalCostHandler {
	handler := new(TotalCostHandler)
	return handler
}

func newPerCreditCostHandler() *PerCreditCostHandler {
	handler := new(PerCreditCostHandler)
	return handler
}

func newEmployersHandler() *EmployersHandler {
	handler := new(EmployersHandler)
	handler.topEmployersHandler = newTopEmployersHandler()
	return handler
}

func newTopEmployersHandler() *TopEmployersHandler {
	handler := new(TopEmployersHandler)
	return handler
}

func newStudentBackgroundsHandler() *StudentBackgroundsHandler {
	handler := new(StudentBackgroundsHandler)
	handler.topBackgroundsHandler = newTopBackgroundsHandler()
	return handler
}

func newTopBackgroundsHandler() *TopBackgroundsHandler {
	handler := new(TopBackgroundsHandler)
	return handler
}

func newDropOutRateHandler() *DropOutRateHandler {
	handler := new(DropOutRateHandler)
	return handler
}

func newStudentEmploymentRateHandler() *StudentEmploymentRateHandler {
	handler := new(StudentEmploymentRateHandler)
	return handler
}

func newSalaryHandler() *SalaryHandler {
	handler := new(SalaryHandler)
	handler.avgSalaryHandler = newAvgSalaryHandler()
	return handler
}

func newAvgSalaryHandler() *AvgSalaryHandler {
	handler := new(AvgSalaryHandler)
	return handler
}

func newAcceptanceRateHandler() *AcceptanceRateHandler {
	handler := new(AcceptanceRateHandler)
	return handler
}
