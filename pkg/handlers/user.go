package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/softplan/tenkai-api/pkg/constraints"
	"github.com/softplan/tenkai-api/pkg/global"

	"github.com/gorilla/mux"
	"github.com/softplan/tenkai-api/pkg/dbms/model"
	"github.com/softplan/tenkai-api/pkg/util"
)

func (appContext *AppContext) newUser(w http.ResponseWriter, r *http.Request) {

	principal := util.GetPrincipal(r)
	if principal.Email == "" {
		http.Error(w, errors.New("Acccess Defined").Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set(global.ContentType, global.JSONContentType)

	var payload model.User

	if err := util.UnmarshalPayload(r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := appContext.Repositories.UserDAO.CreateUser(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (appContext *AppContext) createOrUpdateUser(w http.ResponseWriter, r *http.Request) {

	principal := util.GetPrincipal(r)
	if principal.Email == "" {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	w.Header().Set(global.ContentType, global.JSONContentType)

	var payload model.User

	if err := util.UnmarshalPayload(r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := appContext.Repositories.UserDAO.CreateOrUpdateUser(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auditValues := make(map[string]string)
	auditValues["userEmail"] = payload.Email
	auditValues["userID"] = strconv.Itoa(int(payload.ID))
	auditValues["environmentsName"] = appContext.envListToStringNames(payload.Environments)
	auditValues["environmentsID"] = envListToStringIDs(payload.Environments)
	auditValues["principal"] = principal.Email

	appContext.Auditing.DoAudit(r.Context(), appContext.Elk, principal.Email, "updateOrCreate", auditValues)

	w.WriteHeader(http.StatusCreated)
}

func envListToStringIDs(list []model.Environment) string {
	str := ""
	for i, env := range list {
		if i == 0 {
			str = strconv.Itoa(int(env.ID))
		} else {
			str = str + ", " + strconv.Itoa(int(env.ID))
		}
	}
	return str
}

func (appContext *AppContext) envListToStringNames(list []model.Environment) string {
	str := ""
	for i, env := range list {
		environment, _ := appContext.Repositories.EnvironmentDAO.GetByID(int(env.ID))
		if i == 0 {
			str = environment.Name
		} else {
			str = str + ", " + environment.Name
		}
	}
	return str
}

func (appContext *AppContext) listUsers(w http.ResponseWriter, r *http.Request) {

	result := &model.UserResult{}
	var err error

	keys := r.URL.Query()
	email := keys.Get("email")

	if result.Users, err = appContext.Repositories.UserDAO.ListAllUsers(email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func (appContext *AppContext) deleteUser(w http.ResponseWriter, r *http.Request) {

	principal := util.GetPrincipal(r)
	if !util.Contains(principal.Roles, constraints.TenkaiAdmin) {
		http.Error(w, errors.New("Acccess Denied").Error(), http.StatusUnauthorized)
	}

	vars := mux.Vars(r)
	sl := vars["id"]
	id, _ := strconv.Atoi(sl)
	w.Header().Set(global.ContentType, global.JSONContentType)
	if err := appContext.Repositories.UserDAO.DeleteUser(id); err != nil {
		log.Println("Error deleting variable: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (appContext *AppContext) getUser(w http.ResponseWriter, r *http.Request) {

	principal := util.GetPrincipal(r)
	if principal.Email == "" {
		http.Error(w, errors.New("Acccess Denied").Error(), http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var user model.User
	var err error
	if util.Contains(principal.Roles, "tenkai-manager") {
		manager, err := appContext.Repositories.UserDAO.FindByEmail(principal.Email)
		if err != nil {
			global.Logger.Info(global.AppFields{global.Function: "getUser"}, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		idAsInt, _ := strconv.Atoi(id)
		user, err = appContext.Repositories.UserDAO.FindByUsersIDFilteredByIntersectionEnv(idAsInt, int(manager.ID))
	} else {
		if user, err = appContext.Repositories.UserDAO.FindByID(id); err != nil {
			global.Logger.Info(global.AppFields{global.Function: "getUser"}, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	data, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	w.Header().Add(global.ContentType, global.JSONContentType)
}
