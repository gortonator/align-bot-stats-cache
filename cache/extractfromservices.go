package cache

import (
	"net/http"
	"io/ioutil"
	"github.com/buger/jsonparser"
	"github.com/json-iterator/go"
	"bytes"
	"encoding/json"
	"time"
)

func parseResponseForProperty(res *http.Response, property string) (string, error) {
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	return jsonparser.GetString(body, property)
}

func getAndParse(url, property string) (string, error) {
	timeout := time.Duration(requestTimeOutSeconds * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	res, err := client.Get(url)
	if err != nil {
		return "", err
	}
	return parseResponseForProperty(res, property)
}

func postRequest(jsonBody []byte, url, contentType string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", contentType)
	timeout := time.Duration(requestTimeOutSeconds * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	res, err := client.Do(req)
	return res, err
}

func makeJsonForCampus(campus string) ([]byte, error) {
	reqBodyMap := make(map[string]interface{})
	reqBodyMap[requestCampusProperty] = []string{campus}
	body, err := jsoniter.Marshal(reqBodyMap)
	return body, err
}

func postAndParseTopKArray(url string) ([]string, error) {
	var topKValues []string
	reqBody := []byte(requestTopKValue)
	res, err := postRequest(reqBody, url, requestTypeText)
	if err != nil {
		return topKValues, err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return topKValues, err
	}
	jsonErr := json.Unmarshal(body, &topKValues)
	return topKValues, jsonErr
}

func extractStudentNumbersAll() (string, error) {
	url := servicesServerBase + servicesStudentCountTotalEndpoint
	return getAndParse(url, responseStudentCountProperty)
}

func requestStudentNumberByCampus(campus string) (string, error) {
	url := servicesServerBase + servicesStudentCountEndpoint
	reqBodyJson, err := makeJsonForCampus(campus)
	if err != nil {
		return "", err
	}
	res, err := postRequest(reqBodyJson, url, requestTypeJson)
	if err != nil {
		return "", err
	}
	return parseResponseForProperty(res, responseStudentCountProperty)
}

func extractStudentNumbersSeattle() (string, error) {
	return requestStudentNumberByCampus(requestCampusSeattle)
}

func extractStudentNumbersBoston() (string, error) {
	return requestStudentNumberByCampus(requestCampusBoston)
}

func extractStudentNumbersCharlotte() (string, error) {
	return requestStudentNumberByCampus(requestCampusCharlotte)
}

func extractStudentNumbersSilValley() (string, error) {
	return requestStudentNumberByCampus(requestCampusSiliconValley)
}

func extractNumGraduates() (string, error) {
	url := servicesServerBase + servicesGraduateCountEndpoint
	return getAndParse(url, responseGraduateNumProperty)
}

func extractTopEmployers() ([]string, error) {
	url := servicesServerBase + servicesTopEmpEndpoint
	return postAndParseTopKArray(url)
}

func extractTopBackgrounds() ([]string, error) {
	url := servicesServerBase + servicesTopBgsEndpoint
	return postAndParseTopKArray(url)
}
