package cache

import (
	"net/http"
	"io/ioutil"
	"github.com/buger/jsonparser"
	"github.com/json-iterator/go"
	"bytes"
	"encoding/json"
	"time"
	"crypto/tls"
)

var customHttpsClient = &http.Client{
	Timeout: time.Duration(requestTimeOutSeconds * time.Second),
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

func parseResponseForProperty(res *http.Response, property string) (string, error) {
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	handleError(err)
	return jsonparser.GetString(body, property)
}

func getAndParse(url, property string) (string, error) {
	req, _ := http.NewRequest("GET", url, nil)
	res, err := customHttpsClient.Do(req)
	handleError(err)
	return parseResponseForProperty(res, property)
}

func postRequest(jsonBody []byte, url, contentType string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	handleError(err)
	req.Header.Set("Content-Type", contentType)
	return customHttpsClient.Do(req)
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
	handleError(err)

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	handleError(err)

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
	handleError(err)

	res, err := postRequest(reqBodyJson, url, requestTypeJson)
	handleError(err)

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
