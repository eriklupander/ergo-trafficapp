// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ergo-services/ergo/gen (interfaces: Process)
//
// Generated by this command:
//
//	mockgen -package mock_ergo github.com/ergo-services/ergo/gen Process
//
// Package mock_ergo is a generated GoMock package.
package mock_ergo

import (
	context "context"
	reflect "reflect"
	time "time"

	etf "github.com/ergo-services/ergo/etf"
	gen "github.com/ergo-services/ergo/gen"
	gomock "go.uber.org/mock/gomock"
)

// MockProcess is a mock of Process interface.
type MockProcess struct {
	ctrl     *gomock.Controller
	recorder *MockProcessMockRecorder
}

// MockProcessMockRecorder is the mock recorder for MockProcess.
type MockProcessMockRecorder struct {
	mock *MockProcess
}

// NewMockProcess creates a new mock instance.
func NewMockProcess(ctrl *gomock.Controller) *MockProcess {
	mock := &MockProcess{ctrl: ctrl}
	mock.recorder = &MockProcessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcess) EXPECT() *MockProcessMockRecorder {
	return m.recorder
}

// Aliases mocks base method.
func (m *MockProcess) Aliases() []etf.Alias {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Aliases")
	ret0, _ := ret[0].([]etf.Alias)
	return ret0
}

// Aliases indicates an expected call of Aliases.
func (mr *MockProcessMockRecorder) Aliases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aliases", reflect.TypeOf((*MockProcess)(nil).Aliases))
}

// Behavior mocks base method.
func (m *MockProcess) Behavior() gen.ProcessBehavior {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Behavior")
	ret0, _ := ret[0].(gen.ProcessBehavior)
	return ret0
}

// Behavior indicates an expected call of Behavior.
func (mr *MockProcessMockRecorder) Behavior() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Behavior", reflect.TypeOf((*MockProcess)(nil).Behavior))
}

// CancelSyncRequest mocks base method.
func (m *MockProcess) CancelSyncRequest(arg0 etf.Ref) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelSyncRequest", arg0)
}

// CancelSyncRequest indicates an expected call of CancelSyncRequest.
func (mr *MockProcessMockRecorder) CancelSyncRequest(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelSyncRequest", reflect.TypeOf((*MockProcess)(nil).CancelSyncRequest), arg0)
}

// Children mocks base method.
func (m *MockProcess) Children() ([]etf.Pid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Children")
	ret0, _ := ret[0].([]etf.Pid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Children indicates an expected call of Children.
func (mr *MockProcessMockRecorder) Children() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Children", reflect.TypeOf((*MockProcess)(nil).Children))
}

// Compression mocks base method.
func (m *MockProcess) Compression() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compression")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Compression indicates an expected call of Compression.
func (mr *MockProcessMockRecorder) Compression() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compression", reflect.TypeOf((*MockProcess)(nil).Compression))
}

// CompressionLevel mocks base method.
func (m *MockProcess) CompressionLevel() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompressionLevel")
	ret0, _ := ret[0].(int)
	return ret0
}

// CompressionLevel indicates an expected call of CompressionLevel.
func (mr *MockProcessMockRecorder) CompressionLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompressionLevel", reflect.TypeOf((*MockProcess)(nil).CompressionLevel))
}

// CompressionThreshold mocks base method.
func (m *MockProcess) CompressionThreshold() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompressionThreshold")
	ret0, _ := ret[0].(int)
	return ret0
}

// CompressionThreshold indicates an expected call of CompressionThreshold.
func (mr *MockProcessMockRecorder) CompressionThreshold() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompressionThreshold", reflect.TypeOf((*MockProcess)(nil).CompressionThreshold))
}

// Context mocks base method.
func (m *MockProcess) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockProcessMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockProcess)(nil).Context))
}

// CreateAlias mocks base method.
func (m *MockProcess) CreateAlias() (etf.Alias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAlias")
	ret0, _ := ret[0].(etf.Alias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAlias indicates an expected call of CreateAlias.
func (mr *MockProcessMockRecorder) CreateAlias() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAlias", reflect.TypeOf((*MockProcess)(nil).CreateAlias))
}

// DeleteAlias mocks base method.
func (m *MockProcess) DeleteAlias(arg0 etf.Alias) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAlias", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAlias indicates an expected call of DeleteAlias.
func (mr *MockProcessMockRecorder) DeleteAlias(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAlias", reflect.TypeOf((*MockProcess)(nil).DeleteAlias), arg0)
}

// DemonitorEvent mocks base method.
func (m *MockProcess) DemonitorEvent(arg0 gen.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DemonitorEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DemonitorEvent indicates an expected call of DemonitorEvent.
func (mr *MockProcessMockRecorder) DemonitorEvent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DemonitorEvent", reflect.TypeOf((*MockProcess)(nil).DemonitorEvent), arg0)
}

// DemonitorNode mocks base method.
func (m *MockProcess) DemonitorNode(arg0 etf.Ref) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DemonitorNode", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DemonitorNode indicates an expected call of DemonitorNode.
func (mr *MockProcessMockRecorder) DemonitorNode(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DemonitorNode", reflect.TypeOf((*MockProcess)(nil).DemonitorNode), arg0)
}

// DemonitorProcess mocks base method.
func (m *MockProcess) DemonitorProcess(arg0 etf.Ref) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DemonitorProcess", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DemonitorProcess indicates an expected call of DemonitorProcess.
func (mr *MockProcessMockRecorder) DemonitorProcess(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DemonitorProcess", reflect.TypeOf((*MockProcess)(nil).DemonitorProcess), arg0)
}

// Direct mocks base method.
func (m *MockProcess) Direct(arg0 any) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Direct", arg0)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Direct indicates an expected call of Direct.
func (mr *MockProcessMockRecorder) Direct(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Direct", reflect.TypeOf((*MockProcess)(nil).Direct), arg0)
}

// DirectWithTimeout mocks base method.
func (m *MockProcess) DirectWithTimeout(arg0 any, arg1 int) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DirectWithTimeout", arg0, arg1)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DirectWithTimeout indicates an expected call of DirectWithTimeout.
func (mr *MockProcessMockRecorder) DirectWithTimeout(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DirectWithTimeout", reflect.TypeOf((*MockProcess)(nil).DirectWithTimeout), arg0, arg1)
}

// Env mocks base method.
func (m *MockProcess) Env(arg0 gen.EnvKey) any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Env", arg0)
	ret0, _ := ret[0].(any)
	return ret0
}

// Env indicates an expected call of Env.
func (mr *MockProcessMockRecorder) Env(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Env", reflect.TypeOf((*MockProcess)(nil).Env), arg0)
}

// Exit mocks base method.
func (m *MockProcess) Exit(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exit indicates an expected call of Exit.
func (mr *MockProcessMockRecorder) Exit(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exit", reflect.TypeOf((*MockProcess)(nil).Exit), arg0)
}

// GroupLeader mocks base method.
func (m *MockProcess) GroupLeader() gen.Process {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupLeader")
	ret0, _ := ret[0].(gen.Process)
	return ret0
}

// GroupLeader indicates an expected call of GroupLeader.
func (mr *MockProcessMockRecorder) GroupLeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupLeader", reflect.TypeOf((*MockProcess)(nil).GroupLeader))
}

// Info mocks base method.
func (m *MockProcess) Info() gen.ProcessInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info")
	ret0, _ := ret[0].(gen.ProcessInfo)
	return ret0
}

// Info indicates an expected call of Info.
func (mr *MockProcessMockRecorder) Info() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockProcess)(nil).Info))
}

// IsAlias mocks base method.
func (m *MockProcess) IsAlias(arg0 etf.Alias) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAlias", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAlias indicates an expected call of IsAlias.
func (mr *MockProcessMockRecorder) IsAlias(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAlias", reflect.TypeOf((*MockProcess)(nil).IsAlias), arg0)
}

// IsAlive mocks base method.
func (m *MockProcess) IsAlive() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAlive")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAlive indicates an expected call of IsAlive.
func (mr *MockProcessMockRecorder) IsAlive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAlive", reflect.TypeOf((*MockProcess)(nil).IsAlive))
}

// IsMonitor mocks base method.
func (m *MockProcess) IsMonitor(arg0 etf.Ref) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMonitor", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsMonitor indicates an expected call of IsMonitor.
func (mr *MockProcessMockRecorder) IsMonitor(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMonitor", reflect.TypeOf((*MockProcess)(nil).IsMonitor), arg0)
}

// Kill mocks base method.
func (m *MockProcess) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockProcessMockRecorder) Kill() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockProcess)(nil).Kill))
}

// Link mocks base method.
func (m *MockProcess) Link(arg0 etf.Pid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Link", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Link indicates an expected call of Link.
func (mr *MockProcessMockRecorder) Link(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Link", reflect.TypeOf((*MockProcess)(nil).Link), arg0)
}

// Links mocks base method.
func (m *MockProcess) Links() []etf.Pid {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Links")
	ret0, _ := ret[0].([]etf.Pid)
	return ret0
}

// Links indicates an expected call of Links.
func (mr *MockProcessMockRecorder) Links() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Links", reflect.TypeOf((*MockProcess)(nil).Links))
}

// ListEnv mocks base method.
func (m *MockProcess) ListEnv() map[gen.EnvKey]any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEnv")
	ret0, _ := ret[0].(map[gen.EnvKey]any)
	return ret0
}

// ListEnv indicates an expected call of ListEnv.
func (mr *MockProcessMockRecorder) ListEnv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnv", reflect.TypeOf((*MockProcess)(nil).ListEnv))
}

// MakeRef mocks base method.
func (m *MockProcess) MakeRef() etf.Ref {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRef")
	ret0, _ := ret[0].(etf.Ref)
	return ret0
}

// MakeRef indicates an expected call of MakeRef.
func (mr *MockProcessMockRecorder) MakeRef() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRef", reflect.TypeOf((*MockProcess)(nil).MakeRef))
}

// MonitorEvent mocks base method.
func (m *MockProcess) MonitorEvent(arg0 gen.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonitorEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MonitorEvent indicates an expected call of MonitorEvent.
func (mr *MockProcessMockRecorder) MonitorEvent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonitorEvent", reflect.TypeOf((*MockProcess)(nil).MonitorEvent), arg0)
}

// MonitorNode mocks base method.
func (m *MockProcess) MonitorNode(arg0 string) etf.Ref {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonitorNode", arg0)
	ret0, _ := ret[0].(etf.Ref)
	return ret0
}

// MonitorNode indicates an expected call of MonitorNode.
func (mr *MockProcessMockRecorder) MonitorNode(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonitorNode", reflect.TypeOf((*MockProcess)(nil).MonitorNode), arg0)
}

// MonitorProcess mocks base method.
func (m *MockProcess) MonitorProcess(arg0 any) etf.Ref {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonitorProcess", arg0)
	ret0, _ := ret[0].(etf.Ref)
	return ret0
}

// MonitorProcess indicates an expected call of MonitorProcess.
func (mr *MockProcessMockRecorder) MonitorProcess(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonitorProcess", reflect.TypeOf((*MockProcess)(nil).MonitorProcess), arg0)
}

// MonitoredBy mocks base method.
func (m *MockProcess) MonitoredBy() []etf.Pid {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonitoredBy")
	ret0, _ := ret[0].([]etf.Pid)
	return ret0
}

// MonitoredBy indicates an expected call of MonitoredBy.
func (mr *MockProcessMockRecorder) MonitoredBy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonitoredBy", reflect.TypeOf((*MockProcess)(nil).MonitoredBy))
}

// Monitors mocks base method.
func (m *MockProcess) Monitors() []etf.Pid {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Monitors")
	ret0, _ := ret[0].([]etf.Pid)
	return ret0
}

// Monitors indicates an expected call of Monitors.
func (mr *MockProcessMockRecorder) Monitors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Monitors", reflect.TypeOf((*MockProcess)(nil).Monitors))
}

// MonitorsByName mocks base method.
func (m *MockProcess) MonitorsByName() []gen.ProcessID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonitorsByName")
	ret0, _ := ret[0].([]gen.ProcessID)
	return ret0
}

// MonitorsByName indicates an expected call of MonitorsByName.
func (mr *MockProcessMockRecorder) MonitorsByName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonitorsByName", reflect.TypeOf((*MockProcess)(nil).MonitorsByName))
}

// Name mocks base method.
func (m *MockProcess) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockProcessMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockProcess)(nil).Name))
}

// NodeName mocks base method.
func (m *MockProcess) NodeName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeName")
	ret0, _ := ret[0].(string)
	return ret0
}

// NodeName indicates an expected call of NodeName.
func (mr *MockProcessMockRecorder) NodeName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeName", reflect.TypeOf((*MockProcess)(nil).NodeName))
}

// NodeStop mocks base method.
func (m *MockProcess) NodeStop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NodeStop")
}

// NodeStop indicates an expected call of NodeStop.
func (mr *MockProcessMockRecorder) NodeStop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeStop", reflect.TypeOf((*MockProcess)(nil).NodeStop))
}

// NodeUptime mocks base method.
func (m *MockProcess) NodeUptime() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeUptime")
	ret0, _ := ret[0].(int64)
	return ret0
}

// NodeUptime indicates an expected call of NodeUptime.
func (mr *MockProcessMockRecorder) NodeUptime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeUptime", reflect.TypeOf((*MockProcess)(nil).NodeUptime))
}

// Parent mocks base method.
func (m *MockProcess) Parent() gen.Process {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parent")
	ret0, _ := ret[0].(gen.Process)
	return ret0
}

// Parent indicates an expected call of Parent.
func (mr *MockProcessMockRecorder) Parent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parent", reflect.TypeOf((*MockProcess)(nil).Parent))
}

// ProcessByAlias mocks base method.
func (m *MockProcess) ProcessByAlias(arg0 etf.Alias) gen.Process {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessByAlias", arg0)
	ret0, _ := ret[0].(gen.Process)
	return ret0
}

// ProcessByAlias indicates an expected call of ProcessByAlias.
func (mr *MockProcessMockRecorder) ProcessByAlias(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessByAlias", reflect.TypeOf((*MockProcess)(nil).ProcessByAlias), arg0)
}

// ProcessByName mocks base method.
func (m *MockProcess) ProcessByName(arg0 string) gen.Process {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessByName", arg0)
	ret0, _ := ret[0].(gen.Process)
	return ret0
}

// ProcessByName indicates an expected call of ProcessByName.
func (mr *MockProcessMockRecorder) ProcessByName(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessByName", reflect.TypeOf((*MockProcess)(nil).ProcessByName), arg0)
}

// ProcessByPid mocks base method.
func (m *MockProcess) ProcessByPid(arg0 etf.Pid) gen.Process {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessByPid", arg0)
	ret0, _ := ret[0].(gen.Process)
	return ret0
}

// ProcessByPid indicates an expected call of ProcessByPid.
func (mr *MockProcessMockRecorder) ProcessByPid(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessByPid", reflect.TypeOf((*MockProcess)(nil).ProcessByPid), arg0)
}

// ProcessChannels mocks base method.
func (m *MockProcess) ProcessChannels() gen.ProcessChannels {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessChannels")
	ret0, _ := ret[0].(gen.ProcessChannels)
	return ret0
}

// ProcessChannels indicates an expected call of ProcessChannels.
func (mr *MockProcessMockRecorder) ProcessChannels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessChannels", reflect.TypeOf((*MockProcess)(nil).ProcessChannels))
}

// ProcessInfo mocks base method.
func (m *MockProcess) ProcessInfo(arg0 etf.Pid) (gen.ProcessInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessInfo", arg0)
	ret0, _ := ret[0].(gen.ProcessInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessInfo indicates an expected call of ProcessInfo.
func (mr *MockProcessMockRecorder) ProcessInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessInfo", reflect.TypeOf((*MockProcess)(nil).ProcessInfo), arg0)
}

// ProcessList mocks base method.
func (m *MockProcess) ProcessList() []gen.Process {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessList")
	ret0, _ := ret[0].([]gen.Process)
	return ret0
}

// ProcessList indicates an expected call of ProcessList.
func (mr *MockProcessMockRecorder) ProcessList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessList", reflect.TypeOf((*MockProcess)(nil).ProcessList))
}

// PutSyncReply mocks base method.
func (m *MockProcess) PutSyncReply(arg0 etf.Ref, arg1 etf.Term, arg2 error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutSyncReply", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutSyncReply indicates an expected call of PutSyncReply.
func (mr *MockProcessMockRecorder) PutSyncReply(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutSyncReply", reflect.TypeOf((*MockProcess)(nil).PutSyncReply), arg0, arg1, arg2)
}

// PutSyncRequest mocks base method.
func (m *MockProcess) PutSyncRequest(arg0 etf.Ref) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutSyncRequest", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutSyncRequest indicates an expected call of PutSyncRequest.
func (mr *MockProcessMockRecorder) PutSyncRequest(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutSyncRequest", reflect.TypeOf((*MockProcess)(nil).PutSyncRequest), arg0)
}

// RegisterBehavior mocks base method.
func (m *MockProcess) RegisterBehavior(arg0, arg1 string, arg2 gen.ProcessBehavior, arg3 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterBehavior", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterBehavior indicates an expected call of RegisterBehavior.
func (mr *MockProcessMockRecorder) RegisterBehavior(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterBehavior", reflect.TypeOf((*MockProcess)(nil).RegisterBehavior), arg0, arg1, arg2, arg3)
}

// RegisterEvent mocks base method.
func (m *MockProcess) RegisterEvent(arg0 gen.Event, arg1 ...gen.EventMessage) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterEvent", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterEvent indicates an expected call of RegisterEvent.
func (mr *MockProcessMockRecorder) RegisterEvent(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterEvent", reflect.TypeOf((*MockProcess)(nil).RegisterEvent), varargs...)
}

// RegisterName mocks base method.
func (m *MockProcess) RegisterName(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterName", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterName indicates an expected call of RegisterName.
func (mr *MockProcessMockRecorder) RegisterName(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterName", reflect.TypeOf((*MockProcess)(nil).RegisterName), arg0)
}

// RegisteredBehavior mocks base method.
func (m *MockProcess) RegisteredBehavior(arg0, arg1 string) (gen.RegisteredBehavior, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisteredBehavior", arg0, arg1)
	ret0, _ := ret[0].(gen.RegisteredBehavior)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisteredBehavior indicates an expected call of RegisteredBehavior.
func (mr *MockProcessMockRecorder) RegisteredBehavior(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisteredBehavior", reflect.TypeOf((*MockProcess)(nil).RegisteredBehavior), arg0, arg1)
}

// RegisteredBehaviorGroup mocks base method.
func (m *MockProcess) RegisteredBehaviorGroup(arg0 string) []gen.RegisteredBehavior {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisteredBehaviorGroup", arg0)
	ret0, _ := ret[0].([]gen.RegisteredBehavior)
	return ret0
}

// RegisteredBehaviorGroup indicates an expected call of RegisteredBehaviorGroup.
func (mr *MockProcessMockRecorder) RegisteredBehaviorGroup(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisteredBehaviorGroup", reflect.TypeOf((*MockProcess)(nil).RegisteredBehaviorGroup), arg0)
}

// RemoteSpawn mocks base method.
func (m *MockProcess) RemoteSpawn(arg0, arg1 string, arg2 gen.RemoteSpawnOptions, arg3 ...etf.Term) (etf.Pid, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoteSpawn", varargs...)
	ret0, _ := ret[0].(etf.Pid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteSpawn indicates an expected call of RemoteSpawn.
func (mr *MockProcessMockRecorder) RemoteSpawn(arg0, arg1, arg2 any, arg3 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteSpawn", reflect.TypeOf((*MockProcess)(nil).RemoteSpawn), varargs...)
}

// RemoteSpawnWithTimeout mocks base method.
func (m *MockProcess) RemoteSpawnWithTimeout(arg0 int, arg1, arg2 string, arg3 gen.RemoteSpawnOptions, arg4 ...etf.Term) (etf.Pid, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoteSpawnWithTimeout", varargs...)
	ret0, _ := ret[0].(etf.Pid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteSpawnWithTimeout indicates an expected call of RemoteSpawnWithTimeout.
func (mr *MockProcessMockRecorder) RemoteSpawnWithTimeout(arg0, arg1, arg2, arg3 any, arg4 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteSpawnWithTimeout", reflect.TypeOf((*MockProcess)(nil).RemoteSpawnWithTimeout), varargs...)
}

// Self mocks base method.
func (m *MockProcess) Self() etf.Pid {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Self")
	ret0, _ := ret[0].(etf.Pid)
	return ret0
}

// Self indicates an expected call of Self.
func (mr *MockProcessMockRecorder) Self() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Self", reflect.TypeOf((*MockProcess)(nil).Self))
}

// Send mocks base method.
func (m *MockProcess) Send(arg0 any, arg1 etf.Term) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockProcessMockRecorder) Send(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockProcess)(nil).Send), arg0, arg1)
}

// SendAfter mocks base method.
func (m *MockProcess) SendAfter(arg0 any, arg1 etf.Term, arg2 time.Duration) gen.CancelFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAfter", arg0, arg1, arg2)
	ret0, _ := ret[0].(gen.CancelFunc)
	return ret0
}

// SendAfter indicates an expected call of SendAfter.
func (mr *MockProcessMockRecorder) SendAfter(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAfter", reflect.TypeOf((*MockProcess)(nil).SendAfter), arg0, arg1, arg2)
}

// SendEventMessage mocks base method.
func (m *MockProcess) SendEventMessage(arg0 gen.Event, arg1 gen.EventMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEventMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEventMessage indicates an expected call of SendEventMessage.
func (mr *MockProcessMockRecorder) SendEventMessage(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEventMessage", reflect.TypeOf((*MockProcess)(nil).SendEventMessage), arg0, arg1)
}

// SetCompression mocks base method.
func (m *MockProcess) SetCompression(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCompression", arg0)
}

// SetCompression indicates an expected call of SetCompression.
func (mr *MockProcessMockRecorder) SetCompression(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCompression", reflect.TypeOf((*MockProcess)(nil).SetCompression), arg0)
}

// SetCompressionLevel mocks base method.
func (m *MockProcess) SetCompressionLevel(arg0 int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCompressionLevel", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SetCompressionLevel indicates an expected call of SetCompressionLevel.
func (mr *MockProcessMockRecorder) SetCompressionLevel(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCompressionLevel", reflect.TypeOf((*MockProcess)(nil).SetCompressionLevel), arg0)
}

// SetCompressionThreshold mocks base method.
func (m *MockProcess) SetCompressionThreshold(arg0 int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCompressionThreshold", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SetCompressionThreshold indicates an expected call of SetCompressionThreshold.
func (mr *MockProcessMockRecorder) SetCompressionThreshold(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCompressionThreshold", reflect.TypeOf((*MockProcess)(nil).SetCompressionThreshold), arg0)
}

// SetEnv mocks base method.
func (m *MockProcess) SetEnv(arg0 gen.EnvKey, arg1 any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetEnv", arg0, arg1)
}

// SetEnv indicates an expected call of SetEnv.
func (mr *MockProcessMockRecorder) SetEnv(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEnv", reflect.TypeOf((*MockProcess)(nil).SetEnv), arg0, arg1)
}

// SetTrapExit mocks base method.
func (m *MockProcess) SetTrapExit(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrapExit", arg0)
}

// SetTrapExit indicates an expected call of SetTrapExit.
func (mr *MockProcessMockRecorder) SetTrapExit(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrapExit", reflect.TypeOf((*MockProcess)(nil).SetTrapExit), arg0)
}

// Spawn mocks base method.
func (m *MockProcess) Spawn(arg0 string, arg1 gen.ProcessOptions, arg2 gen.ProcessBehavior, arg3 ...etf.Term) (gen.Process, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Spawn", varargs...)
	ret0, _ := ret[0].(gen.Process)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Spawn indicates an expected call of Spawn.
func (mr *MockProcessMockRecorder) Spawn(arg0, arg1, arg2 any, arg3 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Spawn", reflect.TypeOf((*MockProcess)(nil).Spawn), varargs...)
}

// TrapExit mocks base method.
func (m *MockProcess) TrapExit() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrapExit")
	ret0, _ := ret[0].(bool)
	return ret0
}

// TrapExit indicates an expected call of TrapExit.
func (mr *MockProcessMockRecorder) TrapExit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrapExit", reflect.TypeOf((*MockProcess)(nil).TrapExit))
}

// Unlink mocks base method.
func (m *MockProcess) Unlink(arg0 etf.Pid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlink", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unlink indicates an expected call of Unlink.
func (mr *MockProcessMockRecorder) Unlink(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlink", reflect.TypeOf((*MockProcess)(nil).Unlink), arg0)
}

// UnregisterBehavior mocks base method.
func (m *MockProcess) UnregisterBehavior(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterBehavior", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterBehavior indicates an expected call of UnregisterBehavior.
func (mr *MockProcessMockRecorder) UnregisterBehavior(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterBehavior", reflect.TypeOf((*MockProcess)(nil).UnregisterBehavior), arg0, arg1)
}

// UnregisterEvent mocks base method.
func (m *MockProcess) UnregisterEvent(arg0 gen.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterEvent indicates an expected call of UnregisterEvent.
func (mr *MockProcessMockRecorder) UnregisterEvent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterEvent", reflect.TypeOf((*MockProcess)(nil).UnregisterEvent), arg0)
}

// UnregisterName mocks base method.
func (m *MockProcess) UnregisterName(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterName", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterName indicates an expected call of UnregisterName.
func (mr *MockProcessMockRecorder) UnregisterName(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterName", reflect.TypeOf((*MockProcess)(nil).UnregisterName), arg0)
}

// Wait mocks base method.
func (m *MockProcess) Wait() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Wait")
}

// Wait indicates an expected call of Wait.
func (mr *MockProcessMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockProcess)(nil).Wait))
}

// WaitSyncReply mocks base method.
func (m *MockProcess) WaitSyncReply(arg0 etf.Ref, arg1 int) (etf.Term, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitSyncReply", arg0, arg1)
	ret0, _ := ret[0].(etf.Term)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitSyncReply indicates an expected call of WaitSyncReply.
func (mr *MockProcessMockRecorder) WaitSyncReply(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitSyncReply", reflect.TypeOf((*MockProcess)(nil).WaitSyncReply), arg0, arg1)
}

// WaitWithTimeout mocks base method.
func (m *MockProcess) WaitWithTimeout(arg0 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitWithTimeout", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitWithTimeout indicates an expected call of WaitWithTimeout.
func (mr *MockProcessMockRecorder) WaitWithTimeout(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitWithTimeout", reflect.TypeOf((*MockProcess)(nil).WaitWithTimeout), arg0)
}
