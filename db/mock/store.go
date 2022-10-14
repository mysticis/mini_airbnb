// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mysticis/airbnb_mini/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	sqlc "github.com/mysticis/airbnb_mini/db/sqlc"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAdmin mocks base method
func (m *MockStore) CreateAdmin(arg0 context.Context, arg1 sqlc.CreateAdminParams) (sqlc.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdmin", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAdmin indicates an expected call of CreateAdmin
func (mr *MockStoreMockRecorder) CreateAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdmin", reflect.TypeOf((*MockStore)(nil).CreateAdmin), arg0, arg1)
}

// CreateLandlord mocks base method
func (m *MockStore) CreateLandlord(arg0 context.Context, arg1 sqlc.CreateLandlordParams) (sqlc.Landlord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLandlord", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Landlord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLandlord indicates an expected call of CreateLandlord
func (mr *MockStoreMockRecorder) CreateLandlord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLandlord", reflect.TypeOf((*MockStore)(nil).CreateLandlord), arg0, arg1)
}

// CreateReservation mocks base method
func (m *MockStore) CreateReservation(arg0 context.Context, arg1 sqlc.CreateReservationParams) (sqlc.Reservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReservation", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Reservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReservation indicates an expected call of CreateReservation
func (mr *MockStoreMockRecorder) CreateReservation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReservation", reflect.TypeOf((*MockStore)(nil).CreateReservation), arg0, arg1)
}

// CreateReview mocks base method
func (m *MockStore) CreateReview(arg0 context.Context, arg1 sqlc.CreateReviewParams) (sqlc.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReview", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReview indicates an expected call of CreateReview
func (mr *MockStoreMockRecorder) CreateReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReview", reflect.TypeOf((*MockStore)(nil).CreateReview), arg0, arg1)
}

// CreateRoom mocks base method
func (m *MockStore) CreateRoom(arg0 context.Context, arg1 sqlc.CreateRoomParams) (sqlc.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRoom indicates an expected call of CreateRoom
func (mr *MockStoreMockRecorder) CreateRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockStore)(nil).CreateRoom), arg0, arg1)
}

// CreateTenant mocks base method
func (m *MockStore) CreateTenant(arg0 context.Context, arg1 sqlc.CreateTenantParams) (sqlc.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTenant", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTenant indicates an expected call of CreateTenant
func (mr *MockStoreMockRecorder) CreateTenant(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTenant", reflect.TypeOf((*MockStore)(nil).CreateTenant), arg0, arg1)
}

// DeleteReservation mocks base method
func (m *MockStore) DeleteReservation(arg0 context.Context, arg1 sqlc.DeleteReservationParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReservation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReservation indicates an expected call of DeleteReservation
func (mr *MockStoreMockRecorder) DeleteReservation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReservation", reflect.TypeOf((*MockStore)(nil).DeleteReservation), arg0, arg1)
}

// DeleteReview mocks base method
func (m *MockStore) DeleteReview(arg0 context.Context, arg1 sqlc.DeleteReviewParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReview", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReview indicates an expected call of DeleteReview
func (mr *MockStoreMockRecorder) DeleteReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReview", reflect.TypeOf((*MockStore)(nil).DeleteReview), arg0, arg1)
}

// DeleteRoom mocks base method
func (m *MockStore) DeleteRoom(arg0 context.Context, arg1 sqlc.DeleteRoomParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoom", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoom indicates an expected call of DeleteRoom
func (mr *MockStoreMockRecorder) DeleteRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoom", reflect.TypeOf((*MockStore)(nil).DeleteRoom), arg0, arg1)
}

// DeleteTenant mocks base method
func (m *MockStore) DeleteTenant(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTenant", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTenant indicates an expected call of DeleteTenant
func (mr *MockStoreMockRecorder) DeleteTenant(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTenant", reflect.TypeOf((*MockStore)(nil).DeleteTenant), arg0, arg1)
}

// GetLandlord mocks base method
func (m *MockStore) GetLandlord(arg0 context.Context, arg1 int32) (sqlc.Landlord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLandlord", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Landlord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLandlord indicates an expected call of GetLandlord
func (mr *MockStoreMockRecorder) GetLandlord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLandlord", reflect.TypeOf((*MockStore)(nil).GetLandlord), arg0, arg1)
}

// GetReservations mocks base method
func (m *MockStore) GetReservations(arg0 context.Context, arg1 int32) ([]sqlc.Reservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReservations", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.Reservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReservations indicates an expected call of GetReservations
func (mr *MockStoreMockRecorder) GetReservations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservations", reflect.TypeOf((*MockStore)(nil).GetReservations), arg0, arg1)
}

// GetRoomByOwner mocks base method
func (m *MockStore) GetRoomByOwner(arg0 context.Context, arg1 sqlc.GetRoomByOwnerParams) (sqlc.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomByOwner", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomByOwner indicates an expected call of GetRoomByOwner
func (mr *MockStoreMockRecorder) GetRoomByOwner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomByOwner", reflect.TypeOf((*MockStore)(nil).GetRoomByOwner), arg0, arg1)
}

// GetRoomReview mocks base method
func (m *MockStore) GetRoomReview(arg0 context.Context, arg1 sqlc.GetRoomReviewParams) (sqlc.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomReview", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomReview indicates an expected call of GetRoomReview
func (mr *MockStoreMockRecorder) GetRoomReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomReview", reflect.TypeOf((*MockStore)(nil).GetRoomReview), arg0, arg1)
}

// GetTenant mocks base method
func (m *MockStore) GetTenant(arg0 context.Context, arg1 int32) (sqlc.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTenant", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTenant indicates an expected call of GetTenant
func (mr *MockStoreMockRecorder) GetTenant(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTenant", reflect.TypeOf((*MockStore)(nil).GetTenant), arg0, arg1)
}

// ListAllRooms mocks base method
func (m *MockStore) ListAllRooms(arg0 context.Context, arg1 sqlc.ListAllRoomsParams) ([]sqlc.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllRooms", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllRooms indicates an expected call of ListAllRooms
func (mr *MockStoreMockRecorder) ListAllRooms(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllRooms", reflect.TypeOf((*MockStore)(nil).ListAllRooms), arg0, arg1)
}

// ListLandlords mocks base method
func (m *MockStore) ListLandlords(arg0 context.Context) ([]sqlc.Landlord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLandlords", arg0)
	ret0, _ := ret[0].([]sqlc.Landlord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLandlords indicates an expected call of ListLandlords
func (mr *MockStoreMockRecorder) ListLandlords(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLandlords", reflect.TypeOf((*MockStore)(nil).ListLandlords), arg0)
}

// ListReservations mocks base method
func (m *MockStore) ListReservations(arg0 context.Context) ([]sqlc.Reservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReservations", arg0)
	ret0, _ := ret[0].([]sqlc.Reservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReservations indicates an expected call of ListReservations
func (mr *MockStoreMockRecorder) ListReservations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReservations", reflect.TypeOf((*MockStore)(nil).ListReservations), arg0)
}

// ListRoomReviews mocks base method
func (m *MockStore) ListRoomReviews(arg0 context.Context, arg1 int32) ([]sqlc.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoomReviews", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoomReviews indicates an expected call of ListRoomReviews
func (mr *MockStoreMockRecorder) ListRoomReviews(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoomReviews", reflect.TypeOf((*MockStore)(nil).ListRoomReviews), arg0, arg1)
}

// ListRoomsByOwner mocks base method
func (m *MockStore) ListRoomsByOwner(arg0 context.Context, arg1 int32) ([]sqlc.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoomsByOwner", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoomsByOwner indicates an expected call of ListRoomsByOwner
func (mr *MockStoreMockRecorder) ListRoomsByOwner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoomsByOwner", reflect.TypeOf((*MockStore)(nil).ListRoomsByOwner), arg0, arg1)
}

// ListTenants mocks base method
func (m *MockStore) ListTenants(arg0 context.Context) ([]sqlc.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTenants", arg0)
	ret0, _ := ret[0].([]sqlc.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTenants indicates an expected call of ListTenants
func (mr *MockStoreMockRecorder) ListTenants(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTenants", reflect.TypeOf((*MockStore)(nil).ListTenants), arg0)
}

// UpdateReservation mocks base method
func (m *MockStore) UpdateReservation(arg0 context.Context, arg1 sqlc.UpdateReservationParams) (sqlc.Reservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReservation", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Reservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateReservation indicates an expected call of UpdateReservation
func (mr *MockStoreMockRecorder) UpdateReservation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReservation", reflect.TypeOf((*MockStore)(nil).UpdateReservation), arg0, arg1)
}

// UpdateReview mocks base method
func (m *MockStore) UpdateReview(arg0 context.Context, arg1 sqlc.UpdateReviewParams) (sqlc.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReview", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateReview indicates an expected call of UpdateReview
func (mr *MockStoreMockRecorder) UpdateReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReview", reflect.TypeOf((*MockStore)(nil).UpdateReview), arg0, arg1)
}

// UpdateRoom mocks base method
func (m *MockStore) UpdateRoom(arg0 context.Context, arg1 sqlc.UpdateRoomParams) (sqlc.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoom", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRoom indicates an expected call of UpdateRoom
func (mr *MockStoreMockRecorder) UpdateRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoom", reflect.TypeOf((*MockStore)(nil).UpdateRoom), arg0, arg1)
}