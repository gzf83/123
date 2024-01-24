// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// EthClient is an autogenerated mock type for the EthClient type
type EthClient struct {
	mock.Mock
}

type EthClient_Expecter struct {
	mock *mock.Mock
}

func (_m *EthClient) EXPECT() *EthClient_Expecter {
	return &EthClient_Expecter{mock: &_m.Mock}
}

// BlockByNumber provides a mock function with given fields: ctx, number
func (_m *EthClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	ret := _m.Called(ctx, number)

	if len(ret) == 0 {
		panic("no return value specified for BlockByNumber")
	}

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Block, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Block); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthClient_BlockByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockByNumber'
type EthClient_BlockByNumber_Call struct {
	*mock.Call
}

// BlockByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - number *big.Int
func (_e *EthClient_Expecter) BlockByNumber(ctx interface{}, number interface{}) *EthClient_BlockByNumber_Call {
	return &EthClient_BlockByNumber_Call{Call: _e.mock.On("BlockByNumber", ctx, number)}
}

func (_c *EthClient_BlockByNumber_Call) Run(run func(ctx context.Context, number *big.Int)) *EthClient_BlockByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*big.Int))
	})
	return _c
}

func (_c *EthClient_BlockByNumber_Call) Return(_a0 *types.Block, _a1 error) *EthClient_BlockByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthClient_BlockByNumber_Call) RunAndReturn(run func(context.Context, *big.Int) (*types.Block, error)) *EthClient_BlockByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// CodeAt provides a mock function with given fields: ctx, account, blockNumber
func (_m *EthClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, account, blockNumber)

	if len(ret) == 0 {
		panic("no return value specified for CodeAt")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) ([]byte, error)); ok {
		return rf(ctx, account, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) []byte); ok {
		r0 = rf(ctx, account, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, account, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthClient_CodeAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CodeAt'
type EthClient_CodeAt_Call struct {
	*mock.Call
}

// CodeAt is a helper method to define mock.On call
//   - ctx context.Context
//   - account common.Address
//   - blockNumber *big.Int
func (_e *EthClient_Expecter) CodeAt(ctx interface{}, account interface{}, blockNumber interface{}) *EthClient_CodeAt_Call {
	return &EthClient_CodeAt_Call{Call: _e.mock.On("CodeAt", ctx, account, blockNumber)}
}

func (_c *EthClient_CodeAt_Call) Run(run func(ctx context.Context, account common.Address, blockNumber *big.Int)) *EthClient_CodeAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(*big.Int))
	})
	return _c
}

func (_c *EthClient_CodeAt_Call) Return(_a0 []byte, _a1 error) *EthClient_CodeAt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EthClient_CodeAt_Call) RunAndReturn(run func(context.Context, common.Address, *big.Int) ([]byte, error)) *EthClient_CodeAt_Call {
	_c.Call.Return(run)
	return _c
}

// NewEthClient creates a new instance of EthClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEthClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *EthClient {
	mock := &EthClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
