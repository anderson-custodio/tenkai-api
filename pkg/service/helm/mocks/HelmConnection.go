// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import helm "k8s.io/helm/pkg/helm"

import kube "k8s.io/helm/pkg/kube"
import kubernetes "k8s.io/client-go/kubernetes"
import mock "github.com/stretchr/testify/mock"
import rest "k8s.io/client-go/rest"

// HelmConnection is an autogenerated mock type for the HelmConnection type
type HelmConnection struct {
	mock.Mock
}

// ConfigForContext provides a mock function with given fields: context, kubeconfig
func (_m *HelmConnection) ConfigForContext(context string, kubeconfig string) (*rest.Config, error) {
	ret := _m.Called(context, kubeconfig)

	var r0 *rest.Config
	if rf, ok := ret.Get(0).(func(string, string) *rest.Config); ok {
		r0 = rf(context, kubeconfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Config)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(context, kubeconfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKubeClient provides a mock function with given fields: context, kubeconfig
func (_m *HelmConnection) GetKubeClient(context string, kubeconfig string) (*rest.Config, kubernetes.Interface, error) {
	ret := _m.Called(context, kubeconfig)

	var r0 *rest.Config
	if rf, ok := ret.Get(0).(func(string, string) *rest.Config); ok {
		r0 = rf(context, kubeconfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Config)
		}
	}

	var r1 kubernetes.Interface
	if rf, ok := ret.Get(1).(func(string, string) kubernetes.Interface); ok {
		r1 = rf(context, kubeconfig)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(kubernetes.Interface)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(context, kubeconfig)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewClient provides a mock function with given fields: tillerHost
func (_m *HelmConnection) NewClient(tillerHost string) helm.Interface {
	ret := _m.Called(tillerHost)

	var r0 helm.Interface
	if rf, ok := ret.Get(0).(func(string) helm.Interface); ok {
		r0 = rf(tillerHost)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(helm.Interface)
		}
	}

	return r0
}

// SetupConnection provides a mock function with given fields: kubeConfig
func (_m *HelmConnection) SetupConnection(kubeConfig string) (string, *kube.Tunnel, error) {
	ret := _m.Called(kubeConfig)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(kubeConfig)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *kube.Tunnel
	if rf, ok := ret.Get(1).(func(string) *kube.Tunnel); ok {
		r1 = rf(kubeConfig)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*kube.Tunnel)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(kubeConfig)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Teardown provides a mock function with given fields: tillerTunnel
func (_m *HelmConnection) Teardown(tillerTunnel *kube.Tunnel) {
	_m.Called(tillerTunnel)
}
