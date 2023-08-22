//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloudstack/APIDiscoveryService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAPIDiscoveryServiceIface is a mock of APIDiscoveryServiceIface interface.
type MockAPIDiscoveryServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockAPIDiscoveryServiceIfaceMockRecorder
}

// MockAPIDiscoveryServiceIfaceMockRecorder is the mock recorder for MockAPIDiscoveryServiceIface.
type MockAPIDiscoveryServiceIfaceMockRecorder struct {
	mock *MockAPIDiscoveryServiceIface
}

// NewMockAPIDiscoveryServiceIface creates a new mock instance.
func NewMockAPIDiscoveryServiceIface(ctrl *gomock.Controller) *MockAPIDiscoveryServiceIface {
	mock := &MockAPIDiscoveryServiceIface{ctrl: ctrl}
	mock.recorder = &MockAPIDiscoveryServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIDiscoveryServiceIface) EXPECT() *MockAPIDiscoveryServiceIfaceMockRecorder {
	return m.recorder
}

// ListApis mocks base method.
func (m *MockAPIDiscoveryServiceIface) ListApis(p *ListApisParams) (*ListApisResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApis", p)
	ret0, _ := ret[0].(*ListApisResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApis indicates an expected call of ListApis.
func (mr *MockAPIDiscoveryServiceIfaceMockRecorder) ListApis(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApis", reflect.TypeOf((*MockAPIDiscoveryServiceIface)(nil).ListApis), p)
}

// NewListApisParams mocks base method.
func (m *MockAPIDiscoveryServiceIface) NewListApisParams() *ListApisParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListApisParams")
	ret0, _ := ret[0].(*ListApisParams)
	return ret0
}

// NewListApisParams indicates an expected call of NewListApisParams.
func (mr *MockAPIDiscoveryServiceIfaceMockRecorder) NewListApisParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListApisParams", reflect.TypeOf((*MockAPIDiscoveryServiceIface)(nil).NewListApisParams))
}
