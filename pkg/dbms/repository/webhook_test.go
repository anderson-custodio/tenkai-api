package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/softplan/tenkai-api/pkg/dbms/model"
	"github.com/stretchr/testify/assert"
)

func getWebHook() model.WebHook {
	var item model.WebHook
	item.Name = "Product Deploy"
	item.Type = "HOOK_DEPLOY_PRODUCT"
	item.URL = "http://example.com"
	item.EnvironmentID = 999
	return item
}

func beforeWebHookTest(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, WebHookDAOImpl, model.WebHook) {
	db, mock, err := sqlmock.New()
	gormDB, err := gorm.Open("postgres", db)

	assert.Nil(t, err)

	dao := WebHookDAOImpl{}
	dao.Db = gormDB

	mock.MatchExpectationsInOrder(false)

	item := getWebHook()

	return gormDB, mock, dao, item
}

func TestCreateWebHook(t *testing.T) {
	gormDB, mock, dao, item := beforeWebHookTest(t)
	defer gormDB.Close()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery(`INSERT INTO "web_hooks" *`).
		WithArgs(AnyTime{}, AnyTime{}, nil, item.Name, item.Type, item.URL, item.EnvironmentID).
		WillReturnRows(rows)

	result, e := dao.CreateWebHook(item)
	assert.Nil(t, e)
	assert.Equal(t, 1, result)

	mock.ExpectationsWereMet()
}

func TestCreateWebHook_Error(t *testing.T) {
	gormDB, mock, dao, item := beforeWebHookTest(t)
	defer gormDB.Close()

	mock.ExpectQuery(`INSERT INTO "web_hooks"`).
		WithArgs(AnyTime{}, AnyTime{}, nil, item.Name, item.Type, item.URL, item.EnvironmentID).
		WillReturnError(errors.New("some error"))

	_, e := dao.CreateWebHook(item)
	assert.Error(t, e)

	mock.ExpectationsWereMet()
}

func TestEditWebHook(t *testing.T) {
	gormDB, mock, dao, item := beforeWebHookTest(t)
	defer gormDB.Close()

	item.ID = 999

	mock.ExpectExec(`UPDATE "web_hooks" SET (.*) WHERE (.*)`).
		WithArgs(AnyTime{}, nil, item.Name, item.Type, item.URL, item.EnvironmentID, item.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	e := dao.EditWebHook(item)
	assert.Nil(t, e)

	mock.ExpectationsWereMet()
}

func TestDeleteWebHook(t *testing.T) {
	gormDB, mock, dao, item := beforeWebHookTest(t)
	defer gormDB.Close()

	item.ID = 999
	mock.ExpectExec(`DELETE FROM "web_hooks" WHERE (.*)`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := dao.DeleteWebHook(999)
	assert.Nil(t, err)

	mock.ExpectationsWereMet()
}

func TestListWebHooks(t *testing.T) {
	gormDB, mock, dao, item := beforeWebHookTest(t)
	defer gormDB.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(999, item.Name)

	item.ID = 999
	mock.ExpectQuery(`SELECT (.+) FROM "web_hooks" WHERE (.+)`).
		WillReturnRows(rows)

	result, err := dao.ListWebHooks()
	assert.Nil(t, err)
	assert.NotNil(t, result)

	mock.ExpectationsWereMet()
}

func TestListWebHooks_NotFound(t *testing.T) {
	gormDB, mock, dao, item := beforeWebHookTest(t)
	defer gormDB.Close()

	item.ID = 999
	mock.ExpectQuery(`SELECT (.+) FROM "web_hooks" WHERE (.+)`).
		WillReturnError(gorm.ErrRecordNotFound)

	result, err := dao.ListWebHooks()
	assert.Nil(t, err)
	assert.Empty(t, result)

	mock.ExpectationsWereMet()
}

func TestListWebHooks_Error(t *testing.T) {
	gormDB, mock, dao, item := beforeWebHookTest(t)
	defer gormDB.Close()

	item.ID = 999
	mock.ExpectQuery(`SELECT (.+) FROM "web_hooks" WHERE (.+)`).
		WillReturnError(errors.New("mock error"))

	_, err := dao.ListWebHooks()
	assert.Error(t, err)

	mock.ExpectationsWereMet()
}
