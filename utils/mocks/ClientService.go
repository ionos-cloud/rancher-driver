// Code generated by MockGen. DO NOT EDIT.
// Source: utils/client_service.go

// Package mock_utils is a generated GoMock package.
package mock_utils

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ionoscloud "github.com/ionos-cloud/sdk-go/v5"
)

// MockClientService is a mock of ClientService interface.
type MockClientService struct {
	ctrl     *gomock.Controller
	recorder *MockClientServiceMockRecorder
}

// MockClientServiceMockRecorder is the mock recorder for MockClientService.
type MockClientServiceMockRecorder struct {
	mock *MockClientService
}

// NewMockClientService creates a new mock instance.
func NewMockClientService(ctrl *gomock.Controller) *MockClientService {
	mock := &MockClientService{ctrl: ctrl}
	mock.recorder = &MockClientServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientService) EXPECT() *MockClientServiceMockRecorder {
	return m.recorder
}

// CreateAttachNIC mocks base method.
func (m *MockClientService) CreateAttachNIC(datacenterId, serverId, name string, dhcp bool, lanId int32, ips *[]string) (*ionoscloud.Nic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAttachNIC", datacenterId, serverId, name, dhcp, lanId, ips)
	ret0, _ := ret[0].(*ionoscloud.Nic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAttachNIC indicates an expected call of CreateAttachNIC.
func (mr *MockClientServiceMockRecorder) CreateAttachNIC(datacenterId, serverId, name, dhcp, lanId, ips interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAttachNIC", reflect.TypeOf((*MockClientService)(nil).CreateAttachNIC), datacenterId, serverId, name, dhcp, lanId, ips)
}

// CreateAttachVolume mocks base method.
func (m *MockClientService) CreateAttachVolume(datacenterId, serverId, diskType, name, imageAlias, imagePassword, zone, sshKey string, diskSize float32) (*ionoscloud.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAttachVolume", datacenterId, serverId, diskType, name, imageAlias, imagePassword, zone, sshKey, diskSize)
	ret0, _ := ret[0].(*ionoscloud.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAttachVolume indicates an expected call of CreateAttachVolume.
func (mr *MockClientServiceMockRecorder) CreateAttachVolume(datacenterId, serverId, diskType, name, imageAlias, imagePassword, zone, sshKey, diskSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAttachVolume", reflect.TypeOf((*MockClientService)(nil).CreateAttachVolume), datacenterId, serverId, diskType, name, imageAlias, imagePassword, zone, sshKey, diskSize)
}

// CreateDatacenter mocks base method.
func (m *MockClientService) CreateDatacenter(name, location string) (*ionoscloud.Datacenter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDatacenter", name, location)
	ret0, _ := ret[0].(*ionoscloud.Datacenter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDatacenter indicates an expected call of CreateDatacenter.
func (mr *MockClientServiceMockRecorder) CreateDatacenter(name, location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDatacenter", reflect.TypeOf((*MockClientService)(nil).CreateDatacenter), name, location)
}

// CreateIpBlock mocks base method.
func (m *MockClientService) CreateIpBlock(size int32, location string) (*ionoscloud.IpBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIpBlock", size, location)
	ret0, _ := ret[0].(*ionoscloud.IpBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIpBlock indicates an expected call of CreateIpBlock.
func (mr *MockClientServiceMockRecorder) CreateIpBlock(size, location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIpBlock", reflect.TypeOf((*MockClientService)(nil).CreateIpBlock), size, location)
}

// CreateLan mocks base method.
func (m *MockClientService) CreateLan(datacenterId, name string, public bool) (*ionoscloud.LanPost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLan", datacenterId, name, public)
	ret0, _ := ret[0].(*ionoscloud.LanPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLan indicates an expected call of CreateLan.
func (mr *MockClientServiceMockRecorder) CreateLan(datacenterId, name, public interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLan", reflect.TypeOf((*MockClientService)(nil).CreateLan), datacenterId, name, public)
}

// CreateServer mocks base method.
func (m *MockClientService) CreateServer(datacenterId, location, name, cpufamily, zone string, ram, cores int32) (*ionoscloud.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServer", datacenterId, location, name, cpufamily, zone, ram, cores)
	ret0, _ := ret[0].(*ionoscloud.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServer indicates an expected call of CreateServer.
func (mr *MockClientServiceMockRecorder) CreateServer(datacenterId, location, name, cpufamily, zone, ram, cores interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServer", reflect.TypeOf((*MockClientService)(nil).CreateServer), datacenterId, location, name, cpufamily, zone, ram, cores)
}

// GetDatacenter mocks base method.
func (m *MockClientService) GetDatacenter(datacenterId string) (*ionoscloud.Datacenter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatacenter", datacenterId)
	ret0, _ := ret[0].(*ionoscloud.Datacenter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDatacenter indicates an expected call of GetDatacenter.
func (mr *MockClientServiceMockRecorder) GetDatacenter(datacenterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatacenter", reflect.TypeOf((*MockClientService)(nil).GetDatacenter), datacenterId)
}

// GetImages mocks base method.
func (m *MockClientService) GetImages() (ionoscloud.Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImages")
	ret0, _ := ret[0].(ionoscloud.Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImages indicates an expected call of GetImages.
func (mr *MockClientServiceMockRecorder) GetImages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImages", reflect.TypeOf((*MockClientService)(nil).GetImages))
}

// GetIpBlock mocks base method.
func (m *MockClientService) GetIpBlock(ipBlock *ionoscloud.IpBlock) (*[]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIpBlock", ipBlock)
	ret0, _ := ret[0].(*[]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIpBlock indicates an expected call of GetIpBlock.
func (mr *MockClientServiceMockRecorder) GetIpBlock(ipBlock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIpBlock", reflect.TypeOf((*MockClientService)(nil).GetIpBlock), ipBlock)
}

// GetLocationById mocks base method.
func (m *MockClientService) GetLocationById(regionId, locationId string) (*ionoscloud.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocationById", regionId, locationId)
	ret0, _ := ret[0].(*ionoscloud.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocationById indicates an expected call of GetLocationById.
func (mr *MockClientServiceMockRecorder) GetLocationById(regionId, locationId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocationById", reflect.TypeOf((*MockClientService)(nil).GetLocationById), regionId, locationId)
}

// GetServer mocks base method.
func (m *MockClientService) GetServer(datacenterId, serverId string) (*ionoscloud.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServer", datacenterId, serverId)
	ret0, _ := ret[0].(*ionoscloud.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServer indicates an expected call of GetServer.
func (mr *MockClientServiceMockRecorder) GetServer(datacenterId, serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServer", reflect.TypeOf((*MockClientService)(nil).GetServer), datacenterId, serverId)
}

// RemoveDatacenter mocks base method.
func (m *MockClientService) RemoveDatacenter(datacenterId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDatacenter", datacenterId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDatacenter indicates an expected call of RemoveDatacenter.
func (mr *MockClientServiceMockRecorder) RemoveDatacenter(datacenterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDatacenter", reflect.TypeOf((*MockClientService)(nil).RemoveDatacenter), datacenterId)
}

// RemoveIpBlock mocks base method.
func (m *MockClientService) RemoveIpBlock(ipAddress string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveIpBlock", ipAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveIpBlock indicates an expected call of RemoveIpBlock.
func (mr *MockClientServiceMockRecorder) RemoveIpBlock(ipAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIpBlock", reflect.TypeOf((*MockClientService)(nil).RemoveIpBlock), ipAddress)
}

// RemoveLan mocks base method.
func (m *MockClientService) RemoveLan(datacenterId, lanId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLan", datacenterId, lanId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveLan indicates an expected call of RemoveLan.
func (mr *MockClientServiceMockRecorder) RemoveLan(datacenterId, lanId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLan", reflect.TypeOf((*MockClientService)(nil).RemoveLan), datacenterId, lanId)
}

// RemoveNic mocks base method.
func (m *MockClientService) RemoveNic(datacenterId, serverId, nicId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveNic", datacenterId, serverId, nicId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNic indicates an expected call of RemoveNic.
func (mr *MockClientServiceMockRecorder) RemoveNic(datacenterId, serverId, nicId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNic", reflect.TypeOf((*MockClientService)(nil).RemoveNic), datacenterId, serverId, nicId)
}

// RemoveServer mocks base method.
func (m *MockClientService) RemoveServer(datacenterId, serverId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveServer", datacenterId, serverId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveServer indicates an expected call of RemoveServer.
func (mr *MockClientServiceMockRecorder) RemoveServer(datacenterId, serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveServer", reflect.TypeOf((*MockClientService)(nil).RemoveServer), datacenterId, serverId)
}

// RemoveVolume mocks base method.
func (m *MockClientService) RemoveVolume(datacenterId, volumeId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveVolume", datacenterId, volumeId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveVolume indicates an expected call of RemoveVolume.
func (mr *MockClientServiceMockRecorder) RemoveVolume(datacenterId, volumeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVolume", reflect.TypeOf((*MockClientService)(nil).RemoveVolume), datacenterId, volumeId)
}

// RestartServer mocks base method.
func (m *MockClientService) RestartServer(datacenterId, serverId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestartServer", datacenterId, serverId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestartServer indicates an expected call of RestartServer.
func (mr *MockClientServiceMockRecorder) RestartServer(datacenterId, serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestartServer", reflect.TypeOf((*MockClientService)(nil).RestartServer), datacenterId, serverId)
}

// StartServer mocks base method.
func (m *MockClientService) StartServer(datacenterId, serverId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartServer", datacenterId, serverId)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartServer indicates an expected call of StartServer.
func (mr *MockClientServiceMockRecorder) StartServer(datacenterId, serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartServer", reflect.TypeOf((*MockClientService)(nil).StartServer), datacenterId, serverId)
}

// StopServer mocks base method.
func (m *MockClientService) StopServer(datacenterId, serverId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopServer", datacenterId, serverId)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopServer indicates an expected call of StopServer.
func (mr *MockClientServiceMockRecorder) StopServer(datacenterId, serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopServer", reflect.TypeOf((*MockClientService)(nil).StopServer), datacenterId, serverId)
}
