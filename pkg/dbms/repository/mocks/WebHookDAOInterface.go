// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	model "github.com/softplan/tenkai-api/pkg/dbms/model"
	mock "github.com/stretchr/testify/mock"
)

// WebHookDAOInterface is an autogenerated mock type for the WebHookDAOInterface type
type WebHookDAOInterface struct {
	mock.Mock
}

// CreateWebHook provides a mock function with given fields: e
func (_m *WebHookDAOInterface) CreateWebHook(e model.WebHook) (int, error) {
	ret := _m.Called(e)

	var r0 int
	if rf, ok := ret.Get(0).(func(model.WebHook) int); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.WebHook) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWebHook provides a mock function with given fields: id
func (_m *WebHookDAOInterface) DeleteWebHook(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditWebHook provides a mock function with given fields: e
func (_m *WebHookDAOInterface) EditWebHook(e model.WebHook) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.WebHook) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListWebHooks provides a mock function with given fields:
func (_m *WebHookDAOInterface) ListWebHooks() ([]model.WebHook, error) {
	ret := _m.Called()

	var r0 []model.WebHook
	if rf, ok := ret.Get(0).(func() []model.WebHook); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.WebHook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
