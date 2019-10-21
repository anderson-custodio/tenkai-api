// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import bytes "bytes"
import helmapi "github.com/softplan/tenkai-api/pkg/service/helm"
import mock "github.com/stretchr/testify/mock"
import model "github.com/softplan/tenkai-api/pkg/dbms/model"
import sync "sync"

// HelmServiceInterface is an autogenerated mock type for the HelmServiceInterface type
type HelmServiceInterface struct {
	mock.Mock
}

// AddRepository provides a mock function with given fields: repo
func (_m *HelmServiceInterface) AddRepository(repo model.Repository) error {
	ret := _m.Called(repo)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Repository) error); ok {
		r0 = rf(repo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteHelmRelease provides a mock function with given fields: kubeconfig, releaseName, purge
func (_m *HelmServiceInterface) DeleteHelmRelease(kubeconfig string, releaseName string, purge bool) error {
	ret := _m.Called(kubeconfig, releaseName, purge)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool) error); ok {
		r0 = rf(kubeconfig, releaseName, purge)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePod provides a mock function with given fields: kubeconfig, podName, namespace
func (_m *HelmServiceInterface) DeletePod(kubeconfig string, podName string, namespace string) error {
	ret := _m.Called(kubeconfig, podName, namespace)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(kubeconfig, podName, namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureSettings provides a mock function with given fields: kubeconfig
func (_m *HelmServiceInterface) EnsureSettings(kubeconfig string) {
	_m.Called(kubeconfig)
}

// Get provides a mock function with given fields: kubeconfig, releaseName, revision
func (_m *HelmServiceInterface) Get(kubeconfig string, releaseName string, revision int) (string, error) {
	ret := _m.Called(kubeconfig, releaseName, revision)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, int) string); ok {
		r0 = rf(kubeconfig, releaseName, revision)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int) error); ok {
		r1 = rf(kubeconfig, releaseName, revision)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeployment provides a mock function with given fields: chartName, version
func (_m *HelmServiceInterface) GetDeployment(chartName string, version string) ([]byte, error) {
	ret := _m.Called(chartName, version)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(chartName, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(chartName, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHelmReleaseHistory provides a mock function with given fields: kubeconfig, releaseName
func (_m *HelmServiceInterface) GetHelmReleaseHistory(kubeconfig string, releaseName string) (helmapi.ReleaseHistory, error) {
	ret := _m.Called(kubeconfig, releaseName)

	var r0 helmapi.ReleaseHistory
	if rf, ok := ret.Get(0).(func(string, string) helmapi.ReleaseHistory); ok {
		r0 = rf(kubeconfig, releaseName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(helmapi.ReleaseHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(kubeconfig, releaseName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPods provides a mock function with given fields: kubeconfig, namespace
func (_m *HelmServiceInterface) GetPods(kubeconfig string, namespace string) ([]model.Pod, error) {
	ret := _m.Called(kubeconfig, namespace)

	var r0 []model.Pod
	if rf, ok := ret.Get(0).(func(string, string) []model.Pod); ok {
		r0 = rf(kubeconfig, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(kubeconfig, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReleaseHistory provides a mock function with given fields: kubeconfig, releaseName
func (_m *HelmServiceInterface) GetReleaseHistory(kubeconfig string, releaseName string) (bool, error) {
	ret := _m.Called(kubeconfig, releaseName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(kubeconfig, releaseName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(kubeconfig, releaseName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepositories provides a mock function with given fields:
func (_m *HelmServiceInterface) GetRepositories() ([]model.Repository, error) {
	ret := _m.Called()

	var r0 []model.Repository
	if rf, ok := ret.Get(0).(func() []model.Repository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Repository)
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

// GetServices provides a mock function with given fields: kubeconfig, namespace
func (_m *HelmServiceInterface) GetServices(kubeconfig string, namespace string) ([]model.Service, error) {
	ret := _m.Called(kubeconfig, namespace)

	var r0 []model.Service
	if rf, ok := ret.Get(0).(func(string, string) []model.Service); ok {
		r0 = rf(kubeconfig, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(kubeconfig, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTemplate provides a mock function with given fields: mutex, chartName, version, kind
func (_m *HelmServiceInterface) GetTemplate(mutex *sync.Mutex, chartName string, version string, kind string) ([]byte, error) {
	ret := _m.Called(mutex, chartName, version, kind)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*sync.Mutex, string, string, string) []byte); ok {
		r0 = rf(mutex, chartName, version, kind)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sync.Mutex, string, string, string) error); ok {
		r1 = rf(mutex, chartName, version, kind)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetValues provides a mock function with given fields: chartName, version
func (_m *HelmServiceInterface) GetValues(chartName string, version string) ([]byte, error) {
	ret := _m.Called(chartName, version)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(chartName, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(chartName, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitializeHelm provides a mock function with given fields:
func (_m *HelmServiceInterface) InitializeHelm() {
	_m.Called()
}

// IsThereAnyPodWithThisVersion provides a mock function with given fields: kubeconfig, namespace, releaseName, tag
func (_m *HelmServiceInterface) IsThereAnyPodWithThisVersion(kubeconfig string, namespace string, releaseName string, tag string) (bool, error) {
	ret := _m.Called(kubeconfig, namespace, releaseName, tag)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string, string) bool); ok {
		r0 = rf(kubeconfig, namespace, releaseName, tag)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(kubeconfig, namespace, releaseName, tag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHelmDeployments provides a mock function with given fields: kubeconfig, namespace
func (_m *HelmServiceInterface) ListHelmDeployments(kubeconfig string, namespace string) (*helmapi.HelmListResult, error) {
	ret := _m.Called(kubeconfig, namespace)

	var r0 *helmapi.HelmListResult
	if rf, ok := ret.Get(0).(func(string, string) *helmapi.HelmListResult); ok {
		r0 = rf(kubeconfig, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*helmapi.HelmListResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(kubeconfig, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveRepository provides a mock function with given fields: name
func (_m *HelmServiceInterface) RemoveRepository(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RepoUpdate provides a mock function with given fields:
func (_m *HelmServiceInterface) RepoUpdate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RollbackRelease provides a mock function with given fields: kubeconfig, releaseName, revision
func (_m *HelmServiceInterface) RollbackRelease(kubeconfig string, releaseName string, revision int) error {
	ret := _m.Called(kubeconfig, releaseName, revision)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int) error); ok {
		r0 = rf(kubeconfig, releaseName, revision)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchCharts provides a mock function with given fields: searchTerms, allVersions
func (_m *HelmServiceInterface) SearchCharts(searchTerms []string, allVersions bool) *[]model.SearchResult {
	ret := _m.Called(searchTerms, allVersions)

	var r0 *[]model.SearchResult
	if rf, ok := ret.Get(0).(func([]string, bool) *[]model.SearchResult); ok {
		r0 = rf(searchTerms, allVersions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]model.SearchResult)
		}
	}

	return r0
}

// Upgrade provides a mock function with given fields: upgradeRequest, out
func (_m *HelmServiceInterface) Upgrade(upgradeRequest helmapi.UpgradeRequest, out *bytes.Buffer) error {
	ret := _m.Called(upgradeRequest, out)

	var r0 error
	if rf, ok := ret.Get(0).(func(helmapi.UpgradeRequest, *bytes.Buffer) error); ok {
		r0 = rf(upgradeRequest, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
