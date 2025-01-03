// Code generated by MockGen. DO NOT EDIT.
// Source: ../protocol/interface.go
//
// Generated by this command:
//
//	mockgen -source ../protocol/interface.go
//

// Package mock_protocol is a generated GoMock package.
package mock_protocol

import (
	big "math/big"
	reflect "reflect"

	common "github.com/dominant-strategies/go-quai/common"
	types "github.com/dominant-strategies/go-quai/core/types"
	requestManager "github.com/dominant-strategies/go-quai/p2p/node/requestManager"
	metrics "github.com/libp2p/go-libp2p/core/metrics"
	network "github.com/libp2p/go-libp2p/core/network"
	peer "github.com/libp2p/go-libp2p/core/peer"
	gomock "go.uber.org/mock/gomock"
)

// MockQuaiP2PNode is a mock of QuaiP2PNode interface.
type MockQuaiP2PNode struct {
	ctrl     *gomock.Controller
	recorder *MockQuaiP2PNodeMockRecorder
}

// MockQuaiP2PNodeMockRecorder is the mock recorder for MockQuaiP2PNode.
type MockQuaiP2PNodeMockRecorder struct {
	mock *MockQuaiP2PNode
}

// NewMockQuaiP2PNode creates a new mock instance.
func NewMockQuaiP2PNode(ctrl *gomock.Controller) *MockQuaiP2PNode {
	mock := &MockQuaiP2PNode{ctrl: ctrl}
	mock.recorder = &MockQuaiP2PNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuaiP2PNode) EXPECT() *MockQuaiP2PNodeMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockQuaiP2PNode) Connect(arg0 peer.AddrInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockQuaiP2PNodeMockRecorder) Connect(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockQuaiP2PNode)(nil).Connect), arg0)
}

// GetBandwidthCounter mocks base method.
func (m *MockQuaiP2PNode) GetBandwidthCounter() metrics.Reporter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBandwidthCounter")
	ret0, _ := ret[0].(metrics.Reporter)
	return ret0
}

// GetBandwidthCounter indicates an expected call of GetBandwidthCounter.
func (mr *MockQuaiP2PNodeMockRecorder) GetBandwidthCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBandwidthCounter", reflect.TypeOf((*MockQuaiP2PNode)(nil).GetBandwidthCounter))
}

// GetHeight returns the latest accepted height for the given location
func (m *MockQuaiP2PNode) GetHeight(location common.Location) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeight", location)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetBlockHashByNumber mocks base method.
func (m *MockQuaiP2PNode) GetBlockHashByNumber(number *big.Int, location common.Location) *common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockHashByNumber", number, location)
	ret0, _ := ret[0].(*common.Hash)
	return ret0
}

// GetBlockHashByNumber indicates an expected call of GetBlockHashByNumber.
func (mr *MockQuaiP2PNodeMockRecorder) GetBlockHashByNumber(number, location any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockHashByNumber", reflect.TypeOf((*MockQuaiP2PNode)(nil).GetBlockHashByNumber), number, location)
}

// GetRequestManager mocks base method.
func (m *MockQuaiP2PNode) GetRequestManager() requestManager.RequestManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestManager")
	ret0, _ := ret[0].(requestManager.RequestManager)
	return ret0
}

// GetRequestManager indicates an expected call of GetRequestManager.
func (mr *MockQuaiP2PNodeMockRecorder) GetRequestManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestManager", reflect.TypeOf((*MockQuaiP2PNode)(nil).GetRequestManager))
}

// GetStream mocks base method.
func (m *MockQuaiP2PNode) GetStream(arg0 peer.ID) (network.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStream", arg0)
	ret0, _ := ret[0].(network.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStream indicates an expected call of GetStream.
func (mr *MockQuaiP2PNodeMockRecorder) GetStream(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStream", reflect.TypeOf((*MockQuaiP2PNode)(nil).GetStream), arg0)
}

// GetWorkObject mocks base method.
func (m *MockQuaiP2PNode) GetWorkObject(hash common.Hash, location common.Location) *types.WorkObject {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkObject", hash, location)
	ret0, _ := ret[0].(*types.WorkObject)
	return ret0
}

// GetWorkObject indicates an expected call of GetWorkObject.
func (mr *MockQuaiP2PNodeMockRecorder) GetWorkObject(hash, location any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkObject", reflect.TypeOf((*MockQuaiP2PNode)(nil).GetWorkObject), hash, location)
}

// GetWorkObjectsFrom mocks base method.
func (m *MockQuaiP2PNode) GetWorkObjectsFrom(hash common.Hash, location common.Location, count int) []*types.WorkObjectBlockView {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkObjectsFrom", hash, location, count)
	ret0, _ := ret[0].([]*types.WorkObjectBlockView)
	return ret0
}

// GetWorkObjectsFrom indicates an expected call of GetWorkObjectsFrom.
func (mr *MockQuaiP2PNodeMockRecorder) GetWorkObjectsFrom(hash, location, count any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkObjectsFrom", reflect.TypeOf((*MockQuaiP2PNode)(nil).GetWorkObjectsFrom), hash, location, count)
}
