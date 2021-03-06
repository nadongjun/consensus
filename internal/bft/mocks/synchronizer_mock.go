// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	types "github.com/SmartBFT-Go/consensus/pkg/types"
	mock "github.com/stretchr/testify/mock"
)

// SynchronizerMock is an autogenerated mock type for the SynchronizerMock type
type SynchronizerMock struct {
	mock.Mock
}

// Sync provides a mock function with given fields:
func (_m *SynchronizerMock) Sync() types.Decision {
	ret := _m.Called()

	var r0 types.Decision
	if rf, ok := ret.Get(0).(func() types.Decision); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.Decision)
	}

	return r0
}
