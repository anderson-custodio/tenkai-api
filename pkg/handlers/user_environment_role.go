package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/softplan/tenkai-api/pkg/dbms/model"
	"github.com/softplan/tenkai-api/pkg/global"
	"github.com/softplan/tenkai-api/pkg/util"
)

func (appContext *AppContext) createOrUpdateUserEnvironmentRole(w http.ResponseWriter, r *http.Request) {

	principal := util.GetPrincipal(r)
	if principal.Email == "" {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	w.Header().Set(global.ContentType, global.JSONContentType)
	var payload model.UserEnvironmentRole

	if err := util.UnmarshalPayload(r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := appContext.Repositories.UserEnvironmentRoleDAO.CreateOrUpdate(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	appContext.auditCreateOrUpdateUserEnvironmentRole(r.Context(), principal, payload)

	w.WriteHeader(http.StatusCreated)
}

func (appContext *AppContext) auditCreateOrUpdateUserEnvironmentRole(ctx context.Context, principal model.Principal, payload model.UserEnvironmentRole) {
	function := "auditCreateOrUpdateUserEnvironmentRole"
	auditError := "Error on audit - "
	env, err := appContext.Repositories.EnvironmentDAO.GetByID(int(payload.EnvironmentID))
	if err != nil {
		global.Logger.Error(global.AppFields{global.Function: function}, auditError+err.Error())
		return
	}
	roles, err := appContext.Repositories.SecurityOperationDAO.List()
	if err != nil {
		global.Logger.Error(global.AppFields{global.Function: function}, auditError+err.Error())
		return
	}
	user, err := appContext.Repositories.UserDAO.FindByID(strconv.Itoa(int(payload.UserID)))
	if err != nil {
		global.Logger.Error(global.AppFields{global.Function: function}, auditError+err.Error())
		return
	}

	auditValues := make(map[string]string)
	auditValues["userID"] = strconv.Itoa(int(payload.UserID))
	auditValues["user"] = user.Email
	auditValues["environmentID"] = strconv.Itoa(int(payload.EnvironmentID))
	auditValues["environment"] = env.Name
	auditValues["securityOperationID"] = strconv.Itoa(int(payload.SecurityOperationID))
	for _, role := range roles {
		if role.ID == payload.SecurityOperationID {
			auditValues["securityOperation"] = role.Name
		}
	}
	appContext.Auditing.DoAudit(ctx, appContext.Elk, principal.Email, "updateOrCreateEnvironmentRole", auditValues)
}

func (appContext *AppContext) getUserPolicyByEnvironment(w http.ResponseWriter, r *http.Request) {

	var payload model.GetUserPolicyByEnvironmentRequest
	if err := util.UnmarshalPayload(r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user model.User
	var err error
	if user, err = appContext.Repositories.UserDAO.FindByEmail(payload.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := &model.SecurityOperation{}
	if result, err = appContext.Repositories.UserEnvironmentRoleDAO.
		GetRoleByUserAndEnvironment(user, uint(payload.EnvironmentID)); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func (appContext *AppContext) getEnvironmentUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	list, err := appContext.Repositories.UserEnvironmentRoleDAO.GetUsersAndRoleByEnv(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(list)
	w.Header().Set(global.ContentType, global.JSONContentType)
	w.Write(data)
}
