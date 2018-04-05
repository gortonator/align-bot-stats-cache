package routehandlers

import (
	"net/http"
	"strconv"
	"encoding/json"
	"../models"
	"fmt"
	"io/ioutil"
	"../pathparser"
)

func (h *RedisAppHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if head == "stats" {
		h.statsHandler.ServeHTTP(res, req)
	} else {
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

func (h *StatsHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	switch head {
	case "gender-ratio":
		h.genderRoleHandler.ServeHTTP(res, req)
	case "student-numbers":
		h.studentCountHandler.ServeHTTP(res, req)
	case "num-graduates":
		h.graduatesCountHandler.ServeHTTP(res, req)
	case "cost":
		h.programCostHandler.ServeHTTP(res, req)
	case "employers":
		h.employersHandler.ServeHTTP(res, req)
	case "backgrounds":
		h.studentBackgroundsHandler.ServeHTTP(res, req)
	case "drop-out-rate":
		h.dropOutRateHandler.ServeHTTP(res, req)
	case "employment-rate-out-rate":
		h.studentEmpRateHandler.ServeHTTP(res, req)
	case "salary":
		h.salaryHandler.ServeHTTP(res, req)
	case "acceptance-rate":
		h.acceptanceRateHandler.ServeHTTP(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

func (h *GenderRoleHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if len(req.URL.Path) > 1 {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	switch head {
	case "male":
		h.processMaleCount()
	case "female":
		h.processFemaleCount()
	case "others":
		h.processOtherCount()
	case "":
		h.processAllCount()
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

func (h *StudentCountHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if len(req.URL.Path) > 1 {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	switch head {
	case "seattle":
		h.processSeattleCount()
	case "boston":
		h.processBostonCount()
	case "charlotte":
		h.processCharlotteCount()
	case "svalley":
		h.processSiliconValleyCount()
	case "":
		h.processTotalCount()
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

func (h *GraduatesCountHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if len(head) > 0 || len(req.URL.Path) > 0 {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	h.processTotalCount()
}

func (h *ProgramCostHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if len(req.URL.Path) > 1 {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	switch head {
	case "total":
		h.processTotalCost()
	case "per-credit":
		h.processPerCreditCost()
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}
func (h *EmployersHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	switch head {
	case "top":
		h.topEmployersHandler.ServeHTTP(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

func (h *TopEmployersHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	k, err := strconv.Atoi(req.FormValue("k"))
	if len(req.URL.Path) > 1 || err != nil {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	h.processTopK(k, res)
}

func (h *StudentBackgroundsHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	switch head {
	case "top":
		h.topBackgroundsHandler.ServeHTTP(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

func (h *TopBackgroundsHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	k, err := strconv.Atoi(req.FormValue("k"))
	if len(req.URL.Path) > 1 || err != nil {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	h.processTopK(k)
}

func (h *DropOutRateHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if len(req.URL.Path) > 1 {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	h.processDropOutRate()
}

func (h *StudentEmploymentRateHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if len(req.URL.Path) > 1 {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	h.processEmploymentRate()
}

func (h *SalaryHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if head == "avg" {
		h.avgSalaryHandler.ServeHTTP(res, req)
	} else {
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

func (h *AvgSalaryHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if len(req.URL.Path) > 1 {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	h.processAverageSalary()
}

func (h *AcceptanceRateHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if len(req.URL.Path) > 1 {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	h.processAcceptanceRate()
}

func (h *GenderRoleHandler) processMaleCount() {

}
func (h *GenderRoleHandler) processFemaleCount() {}
func (h *GenderRoleHandler) processOtherCount()  {}
func (h *GenderRoleHandler) processAllCount()    {}

func (h *StudentCountHandler) processTotalCount()         {}
func (h *StudentCountHandler) processBostonCount()        {}
func (h *StudentCountHandler) processSeattleCount()       {}
func (h *StudentCountHandler) processCharlotteCount()     {}
func (h *StudentCountHandler) processSiliconValleyCount() {}

func (h *GraduatesCountHandler) processTotalCount() {}

func (h *ProgramCostHandler) processTotalCost()     {}
func (h *ProgramCostHandler) processPerCreditCost() {}

func (h *TopEmployersHandler) processTopK(k int, res http.ResponseWriter) {
	if k == 0 {
		k = 1
	}
	if k > 5 {
		k = 5
	}
	resp, err := http.Get("http://localhost:10000/top5Emp")
	fmt.Println(resp)
	if err != nil {
		http.Error(res, "Unavailable Resource", http.StatusNotFound)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result []models.Employer
	json.Unmarshal(body, &result)
	fmt.Println(result)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(result)
}

func (h *TopBackgroundsHandler) processTopK(k int) {}

func (h *DropOutRateHandler) processDropOutRate() {}

func (h *StudentEmploymentRateHandler) processEmploymentRate() {}

func (h *AvgSalaryHandler) processAverageSalary() {}

func (h *AcceptanceRateHandler) processAcceptanceRate() {}
