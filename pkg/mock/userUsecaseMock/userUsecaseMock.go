// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interface/user.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	domain "golang_project_ecommerce/pkg/domain"
	utils "golang_project_ecommerce/pkg/utils"
	req "golang_project_ecommerce/pkg/utils/req"
	res "golang_project_ecommerce/pkg/utils/res"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserUsecase is a mock of UserUsecase interface.
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase.
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance.
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserUsecase) AddAddress(c context.Context, address domain.Address) (res.AddressResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", c, address)
	ret0, _ := ret[0].(res.AddressResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserUsecaseMockRecorder) AddAddress(c, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserUsecase)(nil).AddAddress), c, address)
}

// DeleteAddress mocks base method.
func (m *MockUserUsecase) DeleteAddress(c context.Context, addressid uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", c, addressid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockUserUsecaseMockRecorder) DeleteAddress(c, addressid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockUserUsecase)(nil).DeleteAddress), c, addressid)
}

// EditAddress mocks base method.
func (m *MockUserUsecase) EditAddress(c context.Context, address domain.Address) (res.AddressResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditAddress", c, address)
	ret0, _ := ret[0].(res.AddressResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditAddress indicates an expected call of EditAddress.
func (mr *MockUserUsecaseMockRecorder) EditAddress(c, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditAddress", reflect.TypeOf((*MockUserUsecase)(nil).EditAddress), c, address)
}

// EditUserDetails mocks base method.
func (m *MockUserUsecase) EditUserDetails(c context.Context, id int, user req.Usereditreq) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditUserDetails", c, id, user)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditUserDetails indicates an expected call of EditUserDetails.
func (mr *MockUserUsecaseMockRecorder) EditUserDetails(c, id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditUserDetails", reflect.TypeOf((*MockUserUsecase)(nil).EditUserDetails), c, id, user)
}

// FindUserById mocks base method.
func (m *MockUserUsecase) FindUserById(c context.Context, id int) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", c, id)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockUserUsecaseMockRecorder) FindUserById(c, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockUserUsecase)(nil).FindUserById), c, id)
}

// ForgotPassword mocks base method.
func (m *MockUserUsecase) ForgotPassword(c context.Context, usrphn, newpass string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotPassword", c, usrphn, newpass)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForgotPassword indicates an expected call of ForgotPassword.
func (mr *MockUserUsecaseMockRecorder) ForgotPassword(c, usrphn, newpass interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotPassword", reflect.TypeOf((*MockUserUsecase)(nil).ForgotPassword), c, usrphn, newpass)
}

// ListAddresses mocks base method.
func (m *MockUserUsecase) ListAddresses(c context.Context, pagination utils.Pagination, id uint) ([]res.AddressResponse, utils.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAddresses", c, pagination, id)
	ret0, _ := ret[0].([]res.AddressResponse)
	ret1, _ := ret[1].(utils.Metadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAddresses indicates an expected call of ListAddresses.
func (mr *MockUserUsecaseMockRecorder) ListAddresses(c, pagination, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAddresses", reflect.TypeOf((*MockUserUsecase)(nil).ListAddresses), c, pagination, id)
}

// Login mocks base method.
func (m *MockUserUsecase) Login(ctx context.Context, user domain.Users) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, user)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserUsecaseMockRecorder) Login(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserUsecase)(nil).Login), ctx, user)
}

// Register mocks base method.
func (m *MockUserUsecase) Register(ctx context.Context, user domain.Users) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, user)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockUserUsecaseMockRecorder) Register(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserUsecase)(nil).Register), ctx, user)
}

// SendOtpForgotPass mocks base method.
func (m *MockUserUsecase) SendOtpForgotPass(c context.Context, phn string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendOtpForgotPass", c, phn)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendOtpForgotPass indicates an expected call of SendOtpForgotPass.
func (mr *MockUserUsecaseMockRecorder) SendOtpForgotPass(c, phn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendOtpForgotPass", reflect.TypeOf((*MockUserUsecase)(nil).SendOtpForgotPass), c, phn)
}

// SendOtpPhn mocks base method.
func (m *MockUserUsecase) SendOtpPhn(c context.Context, phn domain.Users) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendOtpPhn", c, phn)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendOtpPhn indicates an expected call of SendOtpPhn.
func (mr *MockUserUsecaseMockRecorder) SendOtpPhn(c, phn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendOtpPhn", reflect.TypeOf((*MockUserUsecase)(nil).SendOtpPhn), c, phn)
}

// UpdateStatus mocks base method.
func (m *MockUserUsecase) UpdateStatus(c context.Context, user domain.Users) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", c, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockUserUsecaseMockRecorder) UpdateStatus(c, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockUserUsecase)(nil).UpdateStatus), c, user)
}

// VerifyOtp mocks base method.
func (m *MockUserUsecase) VerifyOtp(c context.Context, phn, otp string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyOtp", c, phn, otp)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyOtp indicates an expected call of VerifyOtp.
func (mr *MockUserUsecaseMockRecorder) VerifyOtp(c, phn, otp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyOtp", reflect.TypeOf((*MockUserUsecase)(nil).VerifyOtp), c, phn, otp)
}

// VerifyOtpForgotpass mocks base method.
func (m *MockUserUsecase) VerifyOtpForgotpass(c context.Context, phn, otp string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyOtpForgotpass", c, phn, otp)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyOtpForgotpass indicates an expected call of VerifyOtpForgotpass.
func (mr *MockUserUsecaseMockRecorder) VerifyOtpForgotpass(c, phn, otp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyOtpForgotpass", reflect.TypeOf((*MockUserUsecase)(nil).VerifyOtpForgotpass), c, phn, otp)
}