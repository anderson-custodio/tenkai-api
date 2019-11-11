package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/softplan/tenkai-api/pkg/constraints"
	"github.com/softplan/tenkai-api/pkg/dbms/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeployTrafficRule(t *testing.T) {

	var payload model.TrafficRequest
	payload.EnvironmentID = 999
	payload.ServiceName = "myservicename-master"
	payload.Domain = "mydomain"
	payload.ContextPath = "/abc"
	payload.HeaderName = "alfa"
	payload.HeaderValue = "beta"

	payS, _ := json.Marshal(payload)

	appContext := AppContext{}

	mockConvention := mockConventionInterface(&appContext)
	mockHelmSvc := mockHelmSvcWithLotOfThings(&appContext)

	envDAO := mockGetByID(&appContext)

	appContext.Repositories.EnvironmentDAO = envDAO
	appContext.HelmServiceAPI = mockHelmSvc

	req, err := http.NewRequest("POST", "/deployTrafficRule", bytes.NewBuffer(payS))
	if err != nil {
		t.Fatal(err)
	}

	mockPrincipal(req, []string{constraints.TenkaiAdmin})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appContext.deployTrafficRule)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	mockConvention.AssertNumberOfCalls(t, "GetKubeConfigFileName", 1)
	envDAO.AssertNumberOfCalls(t, "GetByID", 1)
	mockHelmSvc.AssertNumberOfCalls(t, "Upgrade", 1)

}