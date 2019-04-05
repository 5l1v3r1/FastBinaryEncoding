// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.3.0.0

package proto

import "errors"
import "../fbe"

// Workaround for Go unused imports issue
var _ = fbe.Version

// Receive Order interface
type OnReceiveOrder interface {
    OnReceiveOrder(value *Order)
}

// Receive Order function
type OnReceiveOrderFunc func(value *Order)
func (f OnReceiveOrderFunc) OnReceiveOrder(value *Order) {
    f(value)
}

// Receive Balance interface
type OnReceiveBalance interface {
    OnReceiveBalance(value *Balance)
}

// Receive Balance function
type OnReceiveBalanceFunc func(value *Balance)
func (f OnReceiveBalanceFunc) OnReceiveBalance(value *Balance) {
    f(value)
}

// Receive Account interface
type OnReceiveAccount interface {
    OnReceiveAccount(value *Account)
}

// Receive Account function
type OnReceiveAccountFunc func(value *Account)
func (f OnReceiveAccountFunc) OnReceiveAccount(value *Account) {
    f(value)
}

// Fast Binary Encoding proto receiver
type Receiver struct {
    *fbe.Receiver
    orderValue *Order
    orderModel *OrderModel
    balanceValue *Balance
    balanceModel *BalanceModel
    accountValue *Account
    accountModel *AccountModel

    // Receive Order handler
    HandlerOnReceiveOrder OnReceiveOrder
    // Receive Balance handler
    HandlerOnReceiveBalance OnReceiveBalance
    // Receive Account handler
    HandlerOnReceiveAccount OnReceiveAccount
}

// Create a new proto receiver with an empty buffer
func NewReceiver() *Receiver {
    return NewReceiverWithBuffer(fbe.NewEmptyBuffer())
}

// Create a new proto receiver with the given buffer
func NewReceiverWithBuffer(buffer *fbe.Buffer) *Receiver {
    receiver := &Receiver{
        fbe.NewReceiver(buffer, false),
        NewOrder(),
        NewOrderModel(buffer),
        NewBalance(),
        NewBalanceModel(buffer),
        NewAccount(),
        NewAccountModel(buffer),
        nil,
        nil,
        nil,
    }
    receiver.SetupHandlerOnReceive(receiver)
    receiver.SetupHandlerOnReceiveOrderFunc(func(value *Order) {})
    receiver.SetupHandlerOnReceiveBalanceFunc(func(value *Balance) {})
    receiver.SetupHandlerOnReceiveAccountFunc(func(value *Account) {})
    return receiver
}

// Setup handlers
func (r *Receiver) SetupHandlers(handlers interface{}) {
    r.Receiver.SetupHandlers(handlers)
    if handler, ok := handlers.(OnReceiveOrder); ok {
        r.SetupHandlerOnReceiveOrder(handler)
    }
    if handler, ok := handlers.(OnReceiveBalance); ok {
        r.SetupHandlerOnReceiveBalance(handler)
    }
    if handler, ok := handlers.(OnReceiveAccount); ok {
        r.SetupHandlerOnReceiveAccount(handler)
    }
}

// Setup receive Order handler
func (r *Receiver) SetupHandlerOnReceiveOrder(handler OnReceiveOrder) { r.HandlerOnReceiveOrder = handler }
// Setup receive Order handler function
func (r *Receiver) SetupHandlerOnReceiveOrderFunc(function func(value *Order)) { r.HandlerOnReceiveOrder = OnReceiveOrderFunc(function) }
// Setup receive Balance handler
func (r *Receiver) SetupHandlerOnReceiveBalance(handler OnReceiveBalance) { r.HandlerOnReceiveBalance = handler }
// Setup receive Balance handler function
func (r *Receiver) SetupHandlerOnReceiveBalanceFunc(function func(value *Balance)) { r.HandlerOnReceiveBalance = OnReceiveBalanceFunc(function) }
// Setup receive Account handler
func (r *Receiver) SetupHandlerOnReceiveAccount(handler OnReceiveAccount) { r.HandlerOnReceiveAccount = handler }
// Setup receive Account handler function
func (r *Receiver) SetupHandlerOnReceiveAccountFunc(function func(value *Account)) { r.HandlerOnReceiveAccount = OnReceiveAccountFunc(function) }

// Receive message handler
func (r *Receiver) OnReceive(fbeType int, buffer []byte) (bool, error) {
    switch fbeType {
    case r.orderModel.FBEType():
        // Deserialize the value from the FBE stream
        r.orderModel.Buffer().Attach(buffer)
        if !r.orderModel.Verify() {
            return false, errors.New("proto.Order validation failed")
        }
        deserialized, err := r.orderModel.DeserializeValue(r.orderValue)
        if deserialized <= 0 {
            return false, errors.New("proto.Order deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.orderValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveOrder.OnReceiveOrder(r.orderValue)
        return true, nil
    case r.balanceModel.FBEType():
        // Deserialize the value from the FBE stream
        r.balanceModel.Buffer().Attach(buffer)
        if !r.balanceModel.Verify() {
            return false, errors.New("proto.Balance validation failed")
        }
        deserialized, err := r.balanceModel.DeserializeValue(r.balanceValue)
        if deserialized <= 0 {
            return false, errors.New("proto.Balance deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.balanceValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveBalance.OnReceiveBalance(r.balanceValue)
        return true, nil
    case r.accountModel.FBEType():
        // Deserialize the value from the FBE stream
        r.accountModel.Buffer().Attach(buffer)
        if !r.accountModel.Verify() {
            return false, errors.New("proto.Account validation failed")
        }
        deserialized, err := r.accountModel.DeserializeValue(r.accountValue)
        if deserialized <= 0 {
            return false, errors.New("proto.Account deserialization failed")
        }
        if err != nil {
            return false, err
        }

        // Log the value
        if r.Logging() {
            message := r.accountValue.String()
            r.HandlerOnReceiveLog.OnReceiveLog(message)
        }

        // Call receive handler with deserialized value
        r.HandlerOnReceiveAccount.OnReceiveAccount(r.accountValue)
        return true, nil
    }

    return false, nil
}
