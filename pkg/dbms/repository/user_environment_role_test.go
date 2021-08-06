package repository

import (
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestGetUsersAndRoleByEnvError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)

	gormDB, err := gorm.Open("postgres", db)

	userEnv := UserEnvironmentRoleDAOImpl{}
	userEnv.Db = gormDB
	defer gormDB.Close()
	id := 999
	mock.ExpectQuery(`.*`).
		WithArgs(id).WillReturnError(errors.New("some error"))

	_, err = userEnv.GetUsersAndRoleByEnv(id)
	assert.Error(t, err)
}

func TestGetUsersAndRoleByEnvOK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)

	gormDB, err := gorm.Open("postgres", db)
	defer gormDB.Close()
	userEnv := UserEnvironmentRoleDAOImpl{}
	userEnv.Db = gormDB

	rows := sqlmock.NewRows([]string{"email", "name", "name"}).AddRow("mymail", "envName", "role")

	id := 999
	sql := fmt.Sprintf(`select
			distinct u.email,
			e."name",
			so."name"
		from
			user_environment_roles uer
		join environments e on
			e.id = uer.environment_id
		join users u on
			u.id = uer.user_id
		join security_operations so on
			so.id = uer.security_operation_id
		where
			e.id = %d
		`, id)
	mock.ExpectQuery(sql).
		WillReturnRows(rows)

	data, err := userEnv.GetUsersAndRoleByEnv(id)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(data))
}
