package gotube

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/FireEater64/go-tube/types"
)

const (
	baseURL          string = "https://api.tfl.gov.uk"
	stopPointSuffix  string = "/Stoppoint/Search/"
	lineStatusSuffix string = "Line/Mode/tube,dlr,overground,tflrail/Status/"
)

// TFL is the main class used to access the TFL unified API
type TFL struct {
	applicationID  string // The TFL applicationID
	applicationKey string // The TFL applicationKey
	baseURL        url.URL
}

// NewTFL return a pointer to a TFL API object, using the given applicationID
// and applicationKey
func NewTFL(applicationID string, applicationKey string) *TFL {
	var baseURLToAdd *url.URL
	baseURLToAdd, _ = baseURLToAdd.Parse(baseURL)

	toReturn := TFL{
		applicationID:  applicationID,
		applicationKey: applicationKey,
		baseURL:        *baseURLToAdd}

	return &toReturn
}

func (tfl *TFL) GetStatus() *[]types.LineStatusResponse {
	response := []types.LineStatusResponse{}
	params := map[string]string{}
	url := tfl.buildURL(lineStatusSuffix, &params)
	tfl.getJSONResponse(url, &response)

	return &response
}

func (tfl *TFL) GetStatusForLine(givenLineId string) (string, error) {
	lineStatuses := tfl.GetStatus() // No method of being efficient in the API

	for _, status := range *lineStatuses {
		if status.ID == givenLineId && len(status.Statuses) != 0 {
			return status.Statuses[0].SeverityDescription, nil // TODO: Should probably check all status items?
		}
	}

	return "", errors.New("Could not find specified line in line status response")
}

// *****************************************************************************
// HELPER FUNCTIONS
// *****************************************************************************

func (tfl *TFL) getStopPointID(given string) string {
	response := types.StoppointSearchResponse{}
	params := map[string]string{"modes": "tube"}
	url := tfl.buildURL(stopPointSuffix+given, &params)
	tfl.getJSONResponse(url, &response)

	if response.TotalMatches == 0 {
		// We couldn't find the queried station
		fmt.Printf("Could not find queried station: %s\n", given)
		return "" // TODO: Log
	}

	return response.Matches[0].ID // Assume the biggest hit
}

func (tfl *TFL) getJSONResponse(url string, resultToFill interface{}) {
	resp, getErr := http.Get(url)
	if getErr != nil {
		panic(getErr) // TODO: Logging
	}

	defer resp.Body.Close()
	jsonResponse, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		panic(readErr) // TODO: Logging
	}

	unmarshalErr := json.Unmarshal(jsonResponse, &resultToFill)

	if unmarshalErr != nil {
		panic(unmarshalErr) // TODO: Logging
	}
}

func (tfl *TFL) buildURL(suffix string, params *map[string]string) string {
	toReturn := tfl.baseURL
	toReturn.Path += suffix

	parameters := url.Values{}

	// Add the given key-value pairs (if any)
	for key, value := range *params {
		parameters.Add(key, value)
	}

	parameters.Add("app_id", tfl.applicationID)
	parameters.Add("app_key", tfl.applicationKey)
	toReturn.RawQuery = parameters.Encode()

	return toReturn.String()
}
