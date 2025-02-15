// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andys920605/meme-coin/pkg/database (interfaces: TransactionManager)
//
// Generated by this command:
//
//	mockgen -destination=../../internal/mock/itransaction_manager_mock_database.go -package=mock github.com/andys920605/meme-coin/pkg/database TransactionManager
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	database "github.com/andys920605/meme-coin/pkg/database"
	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockTransactionManager is a mock of TransactionManager interface.
type MockTransactionManager struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionManagerMockRecorder
	isgomock struct{}
}

// MockTransactionManagerMockRecorder is the mock recorder for MockTransactionManager.
type MockTransactionManagerMockRecorder struct {
	mock *MockTransactionManager
}

// NewMockTransactionManager creates a new mock instance.
func NewMockTransactionManager(ctrl *gomock.Controller) *MockTransactionManager {
	mock := &MockTransactionManager{ctrl: ctrl}
	mock.recorder = &MockTransactionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionManager) EXPECT() *MockTransactionManagerMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockTransactionManager) Execute(ctx context.Context, fn database.TxFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockTransactionManagerMockRecorder) Execute(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockTransactionManager)(nil).Execute), ctx, fn)
}

// GetTransaction mocks base method.
func (m *MockTransactionManager) GetTransaction(ctx context.Context) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", ctx)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// GetTransaction indicates an expected call of GetTransaction.
func (mr *MockTransactionManagerMockRecorder) GetTransaction(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockTransactionManager)(nil).GetTransaction), ctx)
}
