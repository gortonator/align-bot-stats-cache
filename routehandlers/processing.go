package routehandlers

import (
	"net/http"
	"encoding/json"
	"github.com/rapidclock/align-bot-stats-cache/cache"
	"log"
	"github.com/json-iterator/go"
	"github.com/rapidclock/align-bot-stats-cache/models"
)

func marshallToJson(marshaler json.Marshaler) []byte {
	body, err := marshaler.MarshalJSON()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return body
}

func jsoniterMarshalling(count *models.StudentCount) []byte {
	body, err := jsoniter.Marshal(count)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func (h *GenderRoleHandler) processAllCount(w http.ResponseWriter) {
	curRatio := cache.GetGenderRatio()
	body, err := curRatio.MarshalJSON()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	w.Write(body)
}

func (h *StudentCountHandler) processAllCount(w http.ResponseWriter) {
	allStudentCount := cache.GetStudentCounts("all")
	w.Write(jsoniterMarshalling(allStudentCount))
}
func (h *StudentCountHandler) processBostonCount(w http.ResponseWriter) {
	bostonStudentCount := cache.GetStudentCounts("boston")
	w.Write(jsoniterMarshalling(bostonStudentCount))
}
func (h *StudentCountHandler) processSeattleCount(w http.ResponseWriter) {
	seattleStudentCount := cache.GetStudentCounts("seattle")
	w.Write(jsoniterMarshalling(seattleStudentCount))
}
func (h *StudentCountHandler) processCharlotteCount(w http.ResponseWriter) {
	charlotteStudentCount := cache.GetStudentCounts("charlotte")
	w.Write(jsoniterMarshalling(charlotteStudentCount))
}
func (h *StudentCountHandler) processSiliconValleyCount(w http.ResponseWriter) {
	siValStudentCount := cache.GetStudentCounts("siliconValley")
	w.Write(jsoniterMarshalling(siValStudentCount))
}

func (h *GraduatesCountHandler) processTotalCount(w http.ResponseWriter) {
	gradCount := cache.GetGraduatesCount()
	w.Write(marshallToJson(gradCount))
}

func (h *ProgramCostHandler) processTotalCost(w http.ResponseWriter) {
	totalCost := cache.GetTotalCost()
	w.Write(marshallToJson(totalCost))
}

func (h *ProgramCostHandler) processPerCreditCost(w http.ResponseWriter) {
	perCreditCost := cache.GetPerCreditCost()
	w.Write(marshallToJson(perCreditCost))
}

func (h *TopEmployersHandler) processTopK(k int, w http.ResponseWriter) {
	topEmp := cache.GetTopEmployers()
	body := marshallToJson(topEmp)
	w.Write(body)
}

func (h *TopBackgroundsHandler) processTopK(k int, w http.ResponseWriter) {
	topBgs := cache.GetTopBackgrounds()
	body := marshallToJson(topBgs)
	w.Write(body)
}

func (h *DropOutRateHandler) processDropOutRate(w http.ResponseWriter) {
	dropOutRate := cache.GetDropOutRate()
	w.Write(marshallToJson(dropOutRate))
}

func (h *StudentEmploymentRateHandler) processEmploymentRate(w http.ResponseWriter) {
	stuEmpRate := cache.GetEmploymentRate()
	w.Write(marshallToJson(stuEmpRate))
}

func (h *AvgSalaryHandler) processAverageSalary(w http.ResponseWriter) {
	avgSalary := cache.GetAverageSalary()
	w.Write(marshallToJson(avgSalary))
}

func (h *AcceptanceRateHandler) processAcceptanceRate(w http.ResponseWriter) {
	accRate := cache.GetAcceptanceRate()
	w.Write(marshallToJson(accRate))
}
