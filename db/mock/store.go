// Code generated by MockGen. DO NOT EDIT.
// Source: db/sqlc/store.go
//
// Generated by this command:
//
//	mockgen -source=db/sqlc/store.go -destination=db/mock/store.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	db "github.com/jpmoraess/moraes-bank/db/sqlc"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddAccountBalance mocks base method.
func (m *MockStore) AddAccountBalance(ctx context.Context, arg db.AddAccountBalanceParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccountBalance", ctx, arg)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccountBalance indicates an expected call of AddAccountBalance.
func (mr *MockStoreMockRecorder) AddAccountBalance(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccountBalance", reflect.TypeOf((*MockStore)(nil).AddAccountBalance), ctx, arg)
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, arg)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), ctx, arg)
}

// CreateEntry mocks base method.
func (m *MockStore) CreateEntry(ctx context.Context, arg db.CreateEntryParams) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", ctx, arg)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry.
func (mr *MockStoreMockRecorder) CreateEntry(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockStore)(nil).CreateEntry), ctx, arg)
}

// CreateTransfer mocks base method.
func (m *MockStore) CreateTransfer(ctx context.Context, arg db.CreateTransferParams) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", ctx, arg)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockStoreMockRecorder) CreateTransfer(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockStore)(nil).CreateTransfer), ctx, arg)
}

// GetAccount mocks base method.
func (m *MockStore) GetAccount(ctx context.Context, id int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, id)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockStoreMockRecorder) GetAccount(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockStore)(nil).GetAccount), ctx, id)
}

// GetAccountForUpdate mocks base method.
func (m *MockStore) GetAccountForUpdate(ctx context.Context, id int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountForUpdate", ctx, id)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountForUpdate indicates an expected call of GetAccountForUpdate.
func (mr *MockStoreMockRecorder) GetAccountForUpdate(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountForUpdate", reflect.TypeOf((*MockStore)(nil).GetAccountForUpdate), ctx, id)
}

// GetAccounts mocks base method.
func (m *MockStore) GetAccounts(ctx context.Context, arg db.GetAccountsParams) ([]db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", ctx, arg)
	ret0, _ := ret[0].([]db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockStoreMockRecorder) GetAccounts(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockStore)(nil).GetAccounts), ctx, arg)
}

// GetEntries mocks base method.
func (m *MockStore) GetEntries(ctx context.Context, arg db.GetEntriesParams) ([]db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntries", ctx, arg)
	ret0, _ := ret[0].([]db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntries indicates an expected call of GetEntries.
func (mr *MockStoreMockRecorder) GetEntries(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntries", reflect.TypeOf((*MockStore)(nil).GetEntries), ctx, arg)
}

// GetEntriesByAccount mocks base method.
func (m *MockStore) GetEntriesByAccount(ctx context.Context, accountID int64) ([]db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntriesByAccount", ctx, accountID)
	ret0, _ := ret[0].([]db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntriesByAccount indicates an expected call of GetEntriesByAccount.
func (mr *MockStoreMockRecorder) GetEntriesByAccount(ctx, accountID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntriesByAccount", reflect.TypeOf((*MockStore)(nil).GetEntriesByAccount), ctx, accountID)
}

// GetEntry mocks base method.
func (m *MockStore) GetEntry(ctx context.Context, id int64) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntry", ctx, id)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntry indicates an expected call of GetEntry.
func (mr *MockStoreMockRecorder) GetEntry(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntry", reflect.TypeOf((*MockStore)(nil).GetEntry), ctx, id)
}

// GetTransfer mocks base method.
func (m *MockStore) GetTransfer(ctx context.Context, id int64) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfer", ctx, id)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfer indicates an expected call of GetTransfer.
func (mr *MockStoreMockRecorder) GetTransfer(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfer", reflect.TypeOf((*MockStore)(nil).GetTransfer), ctx, id)
}

// GetTransfers mocks base method.
func (m *MockStore) GetTransfers(ctx context.Context, arg db.GetTransfersParams) ([]db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfers", ctx, arg)
	ret0, _ := ret[0].([]db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfers indicates an expected call of GetTransfers.
func (mr *MockStoreMockRecorder) GetTransfers(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfers", reflect.TypeOf((*MockStore)(nil).GetTransfers), ctx, arg)
}

// GetTransfersFromAccount mocks base method.
func (m *MockStore) GetTransfersFromAccount(ctx context.Context, fromAccountID int64) ([]db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfersFromAccount", ctx, fromAccountID)
	ret0, _ := ret[0].([]db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfersFromAccount indicates an expected call of GetTransfersFromAccount.
func (mr *MockStoreMockRecorder) GetTransfersFromAccount(ctx, fromAccountID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfersFromAccount", reflect.TypeOf((*MockStore)(nil).GetTransfersFromAccount), ctx, fromAccountID)
}

// GetTransfersToAccount mocks base method.
func (m *MockStore) GetTransfersToAccount(ctx context.Context, toAccountID int64) ([]db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfersToAccount", ctx, toAccountID)
	ret0, _ := ret[0].([]db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfersToAccount indicates an expected call of GetTransfersToAccount.
func (mr *MockStoreMockRecorder) GetTransfersToAccount(ctx, toAccountID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfersToAccount", reflect.TypeOf((*MockStore)(nil).GetTransfersToAccount), ctx, toAccountID)
}

// TransferTx mocks base method.
func (m *MockStore) TransferTx(ctx context.Context, arg db.TransferTxParams) (db.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", ctx, arg)
	ret0, _ := ret[0].(db.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockStoreMockRecorder) TransferTx(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockStore)(nil).TransferTx), ctx, arg)
}

// UpdateAccount mocks base method.
func (m *MockStore) UpdateAccount(ctx context.Context, arg db.UpdateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", ctx, arg)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockStoreMockRecorder) UpdateAccount(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockStore)(nil).UpdateAccount), ctx, arg)
}