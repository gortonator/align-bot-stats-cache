package routehandlers

import (
	"net/http"
	"github.com/rapidclock/align-bot-stats-cache/pathparser"
	"strconv"
)

func (h *RedisAppHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if head == "stats" {
		res.Header().Set("Content-Type", "application/json")
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
	case "employment-rate":
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
	if len(head) > 1 || len(req.URL.Path) > 1 {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	h.processAllCount(res)
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
		h.processSeattleCount(res)
	case "boston":
		h.processBostonCount(res)
	case "charlotte":
		h.processCharlotteCount(res)
	case "svalley":
		h.processSiliconValleyCount(res)
	case "":
		h.processAllCount(res)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

func (h *GraduatesCountHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if len(head) > 0 || len(req.URL.Path) > 1 {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	h.processTotalCount(res)
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
		h.processTotalCost(res)
	case "per-credit":
		h.processPerCreditCost(res)
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
	h.processTopK(k, res)
}

func (h *DropOutRateHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if len(req.URL.Path) > 1 {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	h.processDropOutRate(res)
}

func (h *StudentEmploymentRateHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if len(req.URL.Path) > 1 {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	h.processEmploymentRate(res)
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
	h.processAverageSalary(res)
}

func (h *AcceptanceRateHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_, req.URL.Path = pathparser.ShiftPath(req.URL.Path)
	if len(req.URL.Path) > 1 {
		http.Error(res, "Invalid Path", http.StatusBadRequest)
		return
	}
	h.processAcceptanceRate(res)
}
