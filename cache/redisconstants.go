package cache

const GenderRatioMale = "GR_MALE"
const GenderRatioFemale = "GR_FEMALE"
const GenderRatioOther = "GR_OTHER"

const StudentCountTotal = "STU_CNT_TOTAL"
const StudentCountSeattle = "STU_CNT_SEA"
const StudentCountBoston = "STU_CNT_BOS"
const StudentCountCharlotte = "STU_CNT_CLT"
const StudentCountSiliconValley = "STU_CNT_SIL"

const NumberOfGraduatesTotal = "NUM_GRAD_TOTAL"

const ProgramCostTotal = "PROG_COST_TOTAL"
const ProgramCostPerCredit = "PROG_COST_PER_CREDIT"

const TopEmployers = "TOP_EMPLOYERS"

const TopBackgrounds = "TOP_BACKGROUNDS"

const DropOutRate = "DROP_OUT_RATE"

const EmploymentRate = "EMPLOYMENT_RATE"

const SalaryAvg = "SALARY_AVG"

const AcceptanceRates = "ACCEPTANCE_RATE"

const servicesServerBase = "http://asd2.ccs.neu.edu:8080"
const servicesStudentCountTotalEndpoint = "/stats/total-student-count"
const servicesStudentCountEndpoint = "/stats/student-count"
const servicesGraduateCountEndpoint = "/stats/graduates"
const servicesTopEmpEndpoint = "/coops"
const servicesTopBgsEndpoint = "/undergradmajors"

const responseStudentCountProperty = "studentcount"
const responseGraduateNumProperty = "graduates"

const requestTimeOutSeconds = 1
const requestTypeJson = "application/json"
const requestTypeText = "text/plain"
const requestCampusProperty = "campus"
const requestCampusSeattle = "SEATTLE"
const requestCampusBoston = "BOSTON"
const requestCampusSiliconValley = "SILICON_VALLEY"
const requestCampusCharlotte = "CHARLOTTE"
const requestTopKValue = "5"
