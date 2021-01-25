// Copyright 2021 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by mockery v2.2.2. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/mendersoftware/deviceconfig/model"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// DecommissionDevice provides a mock function with given fields: ctx, devID
func (_m *App) DecommissionDevice(ctx context.Context, devID uuid.UUID) error {
	ret := _m.Called(ctx, devID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, devID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HealthCheck provides a mock function with given fields: ctx
func (_m *App) HealthCheck(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProvisionDevice provides a mock function with given fields: ctx, dev
func (_m *App) ProvisionDevice(ctx context.Context, dev model.NewDevice) error {
	ret := _m.Called(ctx, dev)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.NewDevice) error); ok {
		r0 = rf(ctx, dev)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProvisionTenant provides a mock function with given fields: ctx, tenant
func (_m *App) ProvisionTenant(ctx context.Context, tenant model.NewTenant) error {
	ret := _m.Called(ctx, tenant)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.NewTenant) error); ok {
		r0 = rf(ctx, tenant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
