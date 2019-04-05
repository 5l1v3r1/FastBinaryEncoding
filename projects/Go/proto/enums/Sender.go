// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.3.0.0

package enums

import "errors"
import "../fbe"

// Workaround for Go unused imports issue
var _ = fbe.Version

// Fast Binary Encoding enums sender
type Sender struct {
    *fbe.Sender
    enumsModel *EnumsModel
}

// Create a new enums sender with an empty buffer
func NewSender() *Sender {
    return NewSenderWithBuffer(fbe.NewEmptyBuffer())
}

// Create a new enums sender with the given buffer
func NewSenderWithBuffer(buffer *fbe.Buffer) *Sender {
    return &Sender{
        fbe.NewSender(buffer, false),
        NewEnumsModel(buffer),
    }
}

// Sender models accessors

func (s *Sender) EnumsModel() *EnumsModel { return s.enumsModel }

// Send methods

func (s *Sender) Send(value interface{}) (int, error) {
    switch value := value.(type) {
    case *Enums:
        return s.SendEnums(value)
    }
    return 0, nil
}

func (s *Sender) SendEnums(value *Enums) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.enumsModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("enums.Enums serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.enumsModel.Verify() {
        return 0, errors.New("enums.Enums validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}
