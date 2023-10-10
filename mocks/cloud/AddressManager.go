// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	compute "google.golang.org/api/compute/v1"
)

// AddressManager is an autogenerated mock type for the AddressManager type
type AddressManager struct {
	mock.Mock
}

type AddressManager_Expecter struct {
	mock *mock.Mock
}

func (_m *AddressManager) EXPECT() *AddressManager_Expecter {
	return &AddressManager_Expecter{mock: &_m.Mock}
}

// AddAccessConfig provides a mock function with given fields: project, zone, instance, networkInterface, accessconfig
func (_m *AddressManager) AddAccessConfig(project string, zone string, instance string, networkInterface string, accessconfig *compute.AccessConfig) (*compute.Operation, error) {
	ret := _m.Called(project, zone, instance, networkInterface, accessconfig)

	var r0 *compute.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, *compute.AccessConfig) (*compute.Operation, error)); ok {
		return rf(project, zone, instance, networkInterface, accessconfig)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, string, *compute.AccessConfig) *compute.Operation); ok {
		r0 = rf(project, zone, instance, networkInterface, accessconfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, string, *compute.AccessConfig) error); ok {
		r1 = rf(project, zone, instance, networkInterface, accessconfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddressManager_AddAccessConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddAccessConfig'
type AddressManager_AddAccessConfig_Call struct {
	*mock.Call
}

// AddAccessConfig is a helper method to define mock.On call
//   - project string
//   - zone string
//   - instance string
//   - networkInterface string
//   - accessconfig *compute.AccessConfig
func (_e *AddressManager_Expecter) AddAccessConfig(project interface{}, zone interface{}, instance interface{}, networkInterface interface{}, accessconfig interface{}) *AddressManager_AddAccessConfig_Call {
	return &AddressManager_AddAccessConfig_Call{Call: _e.mock.On("AddAccessConfig", project, zone, instance, networkInterface, accessconfig)}
}

func (_c *AddressManager_AddAccessConfig_Call) Run(run func(project string, zone string, instance string, networkInterface string, accessconfig *compute.AccessConfig)) *AddressManager_AddAccessConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(string), args[4].(*compute.AccessConfig))
	})
	return _c
}

func (_c *AddressManager_AddAccessConfig_Call) Return(_a0 *compute.Operation, _a1 error) *AddressManager_AddAccessConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AddressManager_AddAccessConfig_Call) RunAndReturn(run func(string, string, string, string, *compute.AccessConfig) (*compute.Operation, error)) *AddressManager_AddAccessConfig_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAccessConfig provides a mock function with given fields: project, zone, instance, accessConfig, networkInterface
func (_m *AddressManager) DeleteAccessConfig(project string, zone string, instance string, accessConfig string, networkInterface string) (*compute.Operation, error) {
	ret := _m.Called(project, zone, instance, accessConfig, networkInterface)

	var r0 *compute.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) (*compute.Operation, error)); ok {
		return rf(project, zone, instance, accessConfig, networkInterface)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) *compute.Operation); ok {
		r0 = rf(project, zone, instance, accessConfig, networkInterface)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, string, string) error); ok {
		r1 = rf(project, zone, instance, accessConfig, networkInterface)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddressManager_DeleteAccessConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAccessConfig'
type AddressManager_DeleteAccessConfig_Call struct {
	*mock.Call
}

// DeleteAccessConfig is a helper method to define mock.On call
//   - project string
//   - zone string
//   - instance string
//   - accessConfig string
//   - networkInterface string
func (_e *AddressManager_Expecter) DeleteAccessConfig(project interface{}, zone interface{}, instance interface{}, accessConfig interface{}, networkInterface interface{}) *AddressManager_DeleteAccessConfig_Call {
	return &AddressManager_DeleteAccessConfig_Call{Call: _e.mock.On("DeleteAccessConfig", project, zone, instance, accessConfig, networkInterface)}
}

func (_c *AddressManager_DeleteAccessConfig_Call) Run(run func(project string, zone string, instance string, accessConfig string, networkInterface string)) *AddressManager_DeleteAccessConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *AddressManager_DeleteAccessConfig_Call) Return(_a0 *compute.Operation, _a1 error) *AddressManager_DeleteAccessConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AddressManager_DeleteAccessConfig_Call) RunAndReturn(run func(string, string, string, string, string) (*compute.Operation, error)) *AddressManager_DeleteAccessConfig_Call {
	_c.Call.Return(run)
	return _c
}

// NewAddressManager creates a new instance of AddressManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAddressManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *AddressManager {
	mock := &AddressManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
