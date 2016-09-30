package gotube

import (
	"os"
	"testing"
)

var applicationID string
var applicationKey string

func TestMain(m *testing.M) {
	applicationID = ""  //os.Getenv("applicationID")
	applicationKey = "" //os.Getenv("applicationKey")
	retCode := m.Run()
	os.Exit(retCode)
}

func TestTFLConstructor_withValidCtor_hasCorrectAPIKeys(t *testing.T) {
	underTest := NewTFL(applicationID, applicationKey)

	if underTest.applicationID != applicationID ||
		underTest.applicationKey != applicationKey {
		t.Fail()
	}
}

func TestTFLConstructor_buildsQueryString_withCorrectParameters(t *testing.T) {
	underTest := NewTFL(applicationID, applicationKey)

	params := map[string]string{"foo": "bar"}
	queryString := underTest.buildURL("/test", &params)
	expected := "https://api.tfl.gov.uk/test?app_id=" + applicationID +
		"&app_key=" + applicationKey + "&foo=bar"

	if queryString != expected {
		t.Fatalf("Expected: %s. Received: %s", expected, queryString)
		t.Fail()
	}
}

func TestTFLConstructor_getLineStatus_returnsNonEmptyArray(t *testing.T) {
	underTest := NewTFL(applicationID, applicationKey)

	results := underTest.GetStatus()

	if len(*results) == 0 {
		t.Fatal("No line status information retrieved")
		t.Fail()
	}
}

// We should always be able to retrieve data for the bakerloo line
func TestTFLConstructor_getLineStatusForSpecificLine_returnsNonEmptyString(t *testing.T) {
	underTest := NewTFL(applicationID, applicationKey)

	results, err := underTest.GetStatusForLine("bakerloo")

	if err != nil || (*results)[0].SeverityDescription == "" {
		t.Fatal("No line status information retrieved")
		t.Fail()
	}
}

func TestTFL_getStopPointId_resolvesCorrectStopPoint(t *testing.T) {
	underTest := NewTFL(applicationID, applicationKey)

	kingsCrossStopPointID := underTest.getStopPointID("Kings")
	expected := "HUBKGX"

	if kingsCrossStopPointID != expected {
		t.Fatalf("Expected: %s. Actual: %s", expected, kingsCrossStopPointID)
		t.Fail()
	}
}
