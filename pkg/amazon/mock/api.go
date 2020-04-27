// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/amazon/api.go

// Package mock is a generated GoMock package.
package mock

import (
	cloudformation "github.com/awslabs/goformation/v4/cloudformation"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// ClusterExists mocks base method
func (m *MockAPI) ClusterExists(name string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterExists", name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClusterExists indicates an expected call of ClusterExists
func (mr *MockAPIMockRecorder) ClusterExists(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterExists", reflect.TypeOf((*MockAPI)(nil).ClusterExists), name)
}

// CreateCluster mocks base method
func (m *MockAPI) CreateCluster(name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCluster indicates an expected call of CreateCluster
func (mr *MockAPIMockRecorder) CreateCluster(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockAPI)(nil).CreateCluster), name)
}

// DeleteCluster mocks base method
func (m *MockAPI) DeleteCluster(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster
func (mr *MockAPIMockRecorder) DeleteCluster(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockAPI)(nil).DeleteCluster), name)
}

// GetDefaultVPC mocks base method
func (m *MockAPI) GetDefaultVPC() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultVPC")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultVPC indicates an expected call of GetDefaultVPC
func (mr *MockAPIMockRecorder) GetDefaultVPC() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultVPC", reflect.TypeOf((*MockAPI)(nil).GetDefaultVPC))
}

// GetSubNets mocks base method
func (m *MockAPI) GetSubNets(vpcId string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubNets", vpcId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubNets indicates an expected call of GetSubNets
func (mr *MockAPIMockRecorder) GetSubNets(vpcId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubNets", reflect.TypeOf((*MockAPI)(nil).GetSubNets), vpcId)
}

// ListRolesForPolicy mocks base method
func (m *MockAPI) ListRolesForPolicy(policy string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRolesForPolicy", policy)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRolesForPolicy indicates an expected call of ListRolesForPolicy
func (mr *MockAPIMockRecorder) ListRolesForPolicy(policy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRolesForPolicy", reflect.TypeOf((*MockAPI)(nil).ListRolesForPolicy), policy)
}

// GetRoleArn mocks base method
func (m *MockAPI) GetRoleArn(name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleArn", name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleArn indicates an expected call of GetRoleArn
func (mr *MockAPIMockRecorder) GetRoleArn(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleArn", reflect.TypeOf((*MockAPI)(nil).GetRoleArn), name)
}

// StackExists mocks base method
func (m *MockAPI) StackExists(name string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StackExists", name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StackExists indicates an expected call of StackExists
func (mr *MockAPIMockRecorder) StackExists(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StackExists", reflect.TypeOf((*MockAPI)(nil).StackExists), name)
}

// CreateStack mocks base method
func (m *MockAPI) CreateStack(name string, template *cloudformation.Template) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStack", name, template)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStack indicates an expected call of CreateStack
func (mr *MockAPIMockRecorder) CreateStack(name, template interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStack", reflect.TypeOf((*MockAPI)(nil).CreateStack), name, template)
}

// DescribeStackEvents mocks base method
func (m *MockAPI) DescribeStackEvents(stack string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeStackEvents", stack)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeStackEvents indicates an expected call of DescribeStackEvents
func (mr *MockAPIMockRecorder) DescribeStackEvents(stack interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStackEvents", reflect.TypeOf((*MockAPI)(nil).DescribeStackEvents), stack)
}

// DeleteStack mocks base method
func (m *MockAPI) DeleteStack(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStack", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStack indicates an expected call of DeleteStack
func (mr *MockAPIMockRecorder) DeleteStack(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStack", reflect.TypeOf((*MockAPI)(nil).DeleteStack), name)
}
