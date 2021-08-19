package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	mockAud "github.com/softplan/tenkai-api/pkg/audit/mocks"
	"github.com/softplan/tenkai-api/pkg/dbms/model"
	"github.com/softplan/tenkai-api/pkg/dbms/repository/mocks"
	mockRepo "github.com/softplan/tenkai-api/pkg/dbms/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserPolicyByEnvironment(t *testing.T) {
	appContext := AppContext{}

	p := getUserPolicyByEnv()

	user := mockUser()
	mockUserDao := &mockRepo.UserDAOInterface{}
	mockUserDao.On("FindByEmail", mock.Anything).Return(user, nil)

	result := mockSecurityOperations()
	mockUserEnvRoleDao := &mockRepo.UserEnvironmentRoleDAOInterface{}
	mockUserEnvRoleDao.On("GetRoleByUserAndEnvironment", user, uint(p.EnvironmentID)).
		Return(&result, nil)

	appContext.Repositories.UserDAO = mockUserDao
	appContext.Repositories.UserEnvironmentRoleDAO = mockUserEnvRoleDao

	req, err := http.NewRequest("POST", "/getUserPolicyByEnvironment", payload(p))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appContext.getUserPolicyByEnvironment)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Response should be Ok")

	response := string(rr.Body.Bytes())
	assert.Contains(t, response, `{"ID":999,`)
	assert.Contains(t, response, `"name":"ONLY_DEPLOY",`)
	assert.Contains(t, response, `"policies":["ACTION_DEPLOY"]}`)
}

func TestGetUserPolicyByEnvironment_UnmarshalError(t *testing.T) {
	appContext := AppContext{}
	rr := testUnmarshalPayloadError(t, "/getUserPolicyByEnvironment", appContext.getUserPolicyByEnvironment)
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response should be 500.")
}

func TestGetUserPolicyByEnvironment_UserError(t *testing.T) {
	appContext := AppContext{}

	p := getUserPolicyByEnv()
	user := mockUser()

	mockUserDao := &mockRepo.UserDAOInterface{}
	mockUserDao.On("FindByEmail", mock.Anything).Return(user, errors.New("some error"))

	appContext.Repositories.UserDAO = mockUserDao

	req, err := http.NewRequest("POST", "/getUserPolicyByEnvironment", payload(p))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appContext.getUserPolicyByEnvironment)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response should be 500")
}

func TestGetUserPolicyByEnvironment_RoleError(t *testing.T) {
	appContext := AppContext{}

	p := getUserPolicyByEnv()
	user := mockUser()

	mockUserDao := &mockRepo.UserDAOInterface{}
	mockUserDao.On("FindByEmail", mock.Anything).Return(user, nil)

	mockUserEnvRoleDao := &mockRepo.UserEnvironmentRoleDAOInterface{}
	mockUserEnvRoleDao.On("GetRoleByUserAndEnvironment", user, uint(p.EnvironmentID)).
		Return(nil, errors.New("some error"))

	appContext.Repositories.UserDAO = mockUserDao
	appContext.Repositories.UserEnvironmentRoleDAO = mockUserEnvRoleDao

	req, err := http.NewRequest("POST", "/getUserPolicyByEnvironment", payload(p))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appContext.getUserPolicyByEnvironment)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response should be 500")
}

func mockRepositoriesForAudit(appContext *AppContext) {
	mockEnv := mocks.EnvironmentDAOInterface{}
	mockEnv.On("GetByID", mock.Anything).Return(&model.Environment{Name: "xpto"}, nil)

	mockSecOp := mocks.SecurityOperationDAOInterface{}
	mockSecOp.On("List").Return([]model.SecurityOperation{}, nil)

	mockUser := mocks.UserDAOInterface{}
	mockUser.On("FindByID", mock.Anything).Return(model.User{Email: "xpto@mail.com"}, nil)

	appContext.Repositories.EnvironmentDAO = &mockEnv
	appContext.Repositories.SecurityOperationDAO = &mockSecOp
	appContext.Repositories.UserDAO = &mockUser

	mockAudit := &mockAud.AuditingInterface{}
	mockAudit.On("DoAudit", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	appContext.Auditing = mockAudit
}

func TestCreateOrUpdateUserEnvironmentRole(t *testing.T) {
	appContext := AppContext{}

	p := mockUserEnvRole()

	mockUserEnvRoleDao := &mockRepo.UserEnvironmentRoleDAOInterface{}
	mockUserEnvRoleDao.On("CreateOrUpdate", p).Return(nil)
	appContext.Repositories.UserEnvironmentRoleDAO = mockUserEnvRoleDao

	mockRepositoriesForAudit(&appContext)

	req, err := http.NewRequest("POST", "/createOrUpdateUserEnvironmentRole", payload(p))
	assert.NoError(t, err)
	assert.NotNil(t, req)
	mockPrincipal(req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appContext.createOrUpdateUserEnvironmentRole)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Response should be Created")
}

func TestCreateOrUpdateUserEnvironmentRole_UnmarshalError(t *testing.T) {
	appContext := AppContext{}
	rr := testUnmarshalPayloadError(t, "/createOrUpdateUserEnvironmentRole", appContext.createOrUpdateUserEnvironmentRole)
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response should be 500.")
}

func TestCreateOrUpdateUserEnvironmentRoleError(t *testing.T) {
	appContext := AppContext{}

	p := mockUserEnvRole()

	mockUserEnvRoleDao := &mockRepo.UserEnvironmentRoleDAOInterface{}
	mockUserEnvRoleDao.On("CreateOrUpdate", p).Return(errors.New("some error"))
	appContext.Repositories.UserEnvironmentRoleDAO = mockUserEnvRoleDao

	req, err := http.NewRequest("POST", "/createOrUpdateUserEnvironmentRole", payload(p))
	assert.NoError(t, err)
	assert.NotNil(t, req)
	mockPrincipal(req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(appContext.createOrUpdateUserEnvironmentRole)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response should be 500")
}

func TestGetEnvironmentUsersDBError(t *testing.T) {
	appContext := AppContext{}

	mockUserEnvRoleDao := &mockRepo.UserEnvironmentRoleDAOInterface{}
	mockUserEnvRoleDao.On("GetUsersAndRoleByEnv", 999).Return(nil, errors.New("some database error"))
	appContext.Repositories.UserEnvironmentRoleDAO = mockUserEnvRoleDao

	req, err := http.NewRequest("GET", "/environments/999/users", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/environments/{id}/users", appContext.getEnvironmentUsers).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestGetEnvironmentUsersOK(t *testing.T) {
	appContext := AppContext{}

	returnable := []model.UserEnvRole{
		{
			User:        "user",
			Environment: "env",
			Role:        "role",
		},
	}

	mockUserEnvRoleDao := &mockRepo.UserEnvironmentRoleDAOInterface{}
	mockUserEnvRoleDao.On("GetUsersAndRoleByEnv", 999).Return(returnable, nil)
	appContext.Repositories.UserEnvironmentRoleDAO = mockUserEnvRoleDao

	req, err := http.NewRequest("GET", "/environments/999/users", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/environments/{id}/users", appContext.getEnvironmentUsers).Methods("GET")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
