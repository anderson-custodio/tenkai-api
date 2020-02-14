// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	model "github.com/softplan/tenkai-api/pkg/dbms/model"
	mock "github.com/stretchr/testify/mock"
)

// CompareEnvsQueryDAOInterface is an autogenerated mock type for the CompareEnvsQueryDAOInterface type
type CompareEnvsQueryDAOInterface struct {
	mock.Mock
}

// DeleteCompareEnvQuery provides a mock function with given fields: id
func (_m *CompareEnvsQueryDAOInterface) DeleteCompareEnvQuery(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByUser provides a mock function with given fields: userID
func (_m *CompareEnvsQueryDAOInterface) GetByUser(userID int) ([]model.CompareEnvsQuery, error) {
	ret := _m.Called(userID)

	var r0 []model.CompareEnvsQuery
	if rf, ok := ret.Get(0).(func(int) []model.CompareEnvsQuery); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.CompareEnvsQuery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveCompareEnvsQuery provides a mock function with given fields: env
func (_m *CompareEnvsQueryDAOInterface) SaveCompareEnvsQuery(env model.CompareEnvsQuery) (int, error) {
	ret := _m.Called(env)

	var r0 int
	if rf, ok := ret.Get(0).(func(model.CompareEnvsQuery) int); ok {
		r0 = rf(env)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.CompareEnvsQuery) error); ok {
		r1 = rf(env)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}