// Code generated by MockGen. DO NOT EDIT.
// Source: git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/websocket (interfaces: Upgrader,Connection,Authorizer,Hub)

// Package mock_websocket is a generated GoMock package.
package mock_websocket

import (
	context "context"
	net "net"
	http "net/http"
	reflect "reflect"
	time "time"

	websocket "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/websocket"
	gomock "github.com/golang/mock/gomock"
)

// MockUpgrader is a mock of Upgrader interface.
type MockUpgrader struct {
	ctrl     *gomock.Controller
	recorder *MockUpgraderMockRecorder
}

// MockUpgraderMockRecorder is the mock recorder for MockUpgrader.
type MockUpgraderMockRecorder struct {
	mock *MockUpgrader
}

// NewMockUpgrader creates a new mock instance.
func NewMockUpgrader(ctrl *gomock.Controller) *MockUpgrader {
	mock := &MockUpgrader{ctrl: ctrl}
	mock.recorder = &MockUpgraderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgrader) EXPECT() *MockUpgraderMockRecorder {
	return m.recorder
}

// Upgrade mocks base method.
func (m *MockUpgrader) Upgrade(arg0 http.ResponseWriter, arg1 *http.Request, arg2 http.Header) (websocket.Connection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", arg0, arg1, arg2)
	ret0, _ := ret[0].(websocket.Connection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockUpgraderMockRecorder) Upgrade(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockUpgrader)(nil).Upgrade), arg0, arg1, arg2)
}

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockConnection) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close))
}

// ReadJSON mocks base method.
func (m *MockConnection) ReadJSON(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadJSON", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadJSON indicates an expected call of ReadJSON.
func (mr *MockConnectionMockRecorder) ReadJSON(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadJSON", reflect.TypeOf((*MockConnection)(nil).ReadJSON), arg0)
}

// RemoteAddr mocks base method.
func (m *MockConnection) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockConnectionMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockConnection)(nil).RemoteAddr))
}

// SetPongHandler mocks base method.
func (m *MockConnection) SetPongHandler(arg0 func(string) error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPongHandler", arg0)
}

// SetPongHandler indicates an expected call of SetPongHandler.
func (mr *MockConnectionMockRecorder) SetPongHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPongHandler", reflect.TypeOf((*MockConnection)(nil).SetPongHandler), arg0)
}

// SetReadDeadline mocks base method.
func (m *MockConnection) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline.
func (mr *MockConnectionMockRecorder) SetReadDeadline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockConnection)(nil).SetReadDeadline), arg0)
}

// WriteControl mocks base method.
func (m *MockConnection) WriteControl(arg0 int, arg1 []byte, arg2 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteControl", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteControl indicates an expected call of WriteControl.
func (mr *MockConnectionMockRecorder) WriteControl(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteControl", reflect.TypeOf((*MockConnection)(nil).WriteControl), arg0, arg1, arg2)
}

// WriteJSON mocks base method.
func (m *MockConnection) WriteJSON(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteJSON", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteJSON indicates an expected call of WriteJSON.
func (mr *MockConnectionMockRecorder) WriteJSON(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteJSON", reflect.TypeOf((*MockConnection)(nil).WriteJSON), arg0)
}

// MockAuthorizer is a mock of Authorizer interface.
type MockAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizerMockRecorder
}

// MockAuthorizerMockRecorder is the mock recorder for MockAuthorizer.
type MockAuthorizerMockRecorder struct {
	mock *MockAuthorizer
}

// NewMockAuthorizer creates a new mock instance.
func NewMockAuthorizer(ctrl *gomock.Controller) *MockAuthorizer {
	mock := &MockAuthorizer{ctrl: ctrl}
	mock.recorder = &MockAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizer) EXPECT() *MockAuthorizerMockRecorder {
	return m.recorder
}

// AddRoom mocks base method.
func (m *MockAuthorizer) AddRoom(arg0 string, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoom", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRoom indicates an expected call of AddRoom.
func (mr *MockAuthorizerMockRecorder) AddRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoom", reflect.TypeOf((*MockAuthorizer)(nil).AddRoom), arg0, arg1)
}

// Authenticate mocks base method.
func (m *MockAuthorizer) Authenticate(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockAuthorizerMockRecorder) Authenticate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthorizer)(nil).Authenticate), arg0, arg1)
}

// Authorize mocks base method.
func (m *MockAuthorizer) Authorize(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorize", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authorize indicates an expected call of Authorize.
func (mr *MockAuthorizerMockRecorder) Authorize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthorizer)(nil).Authorize), arg0, arg1)
}

// RemoveRoom mocks base method.
func (m *MockAuthorizer) RemoveRoom(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRoom", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRoom indicates an expected call of RemoveRoom.
func (mr *MockAuthorizerMockRecorder) RemoveRoom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRoom", reflect.TypeOf((*MockAuthorizer)(nil).RemoveRoom), arg0)
}

// MockHub is a mock of Hub interface.
type MockHub struct {
	ctrl     *gomock.Controller
	recorder *MockHubMockRecorder
}

// MockHubMockRecorder is the mock recorder for MockHub.
type MockHubMockRecorder struct {
	mock *MockHub
}

// NewMockHub creates a new mock instance.
func NewMockHub(ctrl *gomock.Controller) *MockHub {
	mock := &MockHub{ctrl: ctrl}
	mock.recorder = &MockHubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHub) EXPECT() *MockHubMockRecorder {
	return m.recorder
}

// CloseRoom mocks base method.
func (m *MockHub) CloseRoom(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseRoom", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseRoom indicates an expected call of CloseRoom.
func (mr *MockHubMockRecorder) CloseRoom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseRoom", reflect.TypeOf((*MockHub)(nil).CloseRoom), arg0)
}

// CloseRoomAndNotify mocks base method.
func (m *MockHub) CloseRoomAndNotify(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseRoomAndNotify", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseRoomAndNotify indicates an expected call of CloseRoomAndNotify.
func (mr *MockHubMockRecorder) CloseRoomAndNotify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseRoomAndNotify", reflect.TypeOf((*MockHub)(nil).CloseRoomAndNotify), arg0)
}

// Connect mocks base method.
func (m *MockHub) Connect(arg0 http.ResponseWriter, arg1 *http.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockHubMockRecorder) Connect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockHub)(nil).Connect), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockHub) GetUsers() map[string][]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].(map[string][]string)
	return ret0
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockHubMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockHub)(nil).GetUsers))
}

// RegisterRoom mocks base method.
func (m *MockHub) RegisterRoom(arg0 string, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterRoom", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterRoom indicates an expected call of RegisterRoom.
func (mr *MockHubMockRecorder) RegisterRoom(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRoom", reflect.TypeOf((*MockHub)(nil).RegisterRoom), varargs...)
}

// RoomHasConnection mocks base method.
func (m *MockHub) RoomHasConnection(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoomHasConnection", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// RoomHasConnection indicates an expected call of RoomHasConnection.
func (mr *MockHubMockRecorder) RoomHasConnection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoomHasConnection", reflect.TypeOf((*MockHub)(nil).RoomHasConnection), arg0)
}

// Send mocks base method.
func (m *MockHub) Send(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", arg0, arg1)
}

// Send indicates an expected call of Send.
func (mr *MockHubMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockHub)(nil).Send), arg0, arg1)
}

// Start mocks base method.
func (m *MockHub) Start(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", arg0)
}

// Start indicates an expected call of Start.
func (mr *MockHubMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockHub)(nil).Start), arg0)
}
