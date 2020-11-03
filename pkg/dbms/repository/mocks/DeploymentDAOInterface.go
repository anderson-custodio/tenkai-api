// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	model "github.com/softplan/tenkai-api/pkg/dbms/model"
	mock "github.com/stretchr/testify/mock"
)

// DeploymentDAOInterface is an autogenerated mock type for the DeploymentDAOInterface type
type DeploymentDAOInterface struct {
	mock.Mock
}

// CreateDeployment provides a mock function with given fields: deployment
func (_m *DeploymentDAOInterface) CreateDeployment(deployment model.Deployment) (int, error) {
	ret := _m.Called(deployment)

	var r0 int
	if rf, ok := ret.Get(0).(func(model.Deployment) int); ok {
		r0 = rf(deployment)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Deployment) error); ok {
		r1 = rf(deployment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditDeployment provides a mock function with given fields: deployment
func (_m *DeploymentDAOInterface) EditDeployment(deployment model.Deployment) error {
	ret := _m.Called(deployment)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Deployment) error); ok {
		r0 = rf(deployment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDeploymentByID provides a mock function with given fields: id
func (_m *DeploymentDAOInterface) GetDeploymentByID(id int) (model.Deployment, error) {
	ret := _m.Called(id)

	var r0 model.Deployment
	if rf, ok := ret.Get(0).(func(int) model.Deployment); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Deployment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDeploymentByUserID provides a mock function with given fields: userID
func (_m *DeploymentDAOInterface) ListDeploymentByUserID(userID int) ([]model.Deployment, error) {
	ret := _m.Called(userID)

	var r0 []model.Deployment
	if rf, ok := ret.Get(0).(func(int) []model.Deployment); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Deployment)
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
