// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.3.0.0

package test

import "errors"
import "../fbe"

// Fast Binary Encoding FlagsTyped field model
type FieldModelFlagsTyped struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int
}

// Create a new FlagsTyped field model
func NewFieldModelFlagsTyped(buffer *fbe.Buffer, offset int) *FieldModelFlagsTyped {
    return &FieldModelFlagsTyped{buffer: buffer, offset: offset}
}

// Get the field size
func (fm *FieldModelFlagsTyped) FBESize() int { return 8 }
// Get the field extra size
func (fm *FieldModelFlagsTyped) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelFlagsTyped) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelFlagsTyped) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelFlagsTyped) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelFlagsTyped) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FieldModelFlagsTyped) Verify() bool { return true }

// Get the value
func (fm *FieldModelFlagsTyped) Get() (*FlagsTyped, error) {
    var value FlagsTyped
    return &value, fm.GetValueDefault(&value, FlagsTyped(0))
}

// Get the value with provided default value
func (fm *FieldModelFlagsTyped) GetDefault(defaults FlagsTyped) (*FlagsTyped, error) {
    var value FlagsTyped
    err := fm.GetValueDefault(&value, defaults)
    return &value, err
}

// Get the value by the given pointer
func (fm *FieldModelFlagsTyped) GetValue(value *FlagsTyped) error {
    return fm.GetValueDefault(value, FlagsTyped(0))
}

// Get the value by the given pointer with provided default value
func (fm *FieldModelFlagsTyped) GetValueDefault(value *FlagsTyped, defaults FlagsTyped) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        *value = defaults
        return nil
    }

    *value = FlagsTyped(fbe.ReadUInt64(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    return nil
}

// Set the value by the given pointer
func (fm *FieldModelFlagsTyped) Set(value *FlagsTyped) error {
    return fm.SetValue(*value)
}

// Set the value
func (fm *FieldModelFlagsTyped) SetValue(value FlagsTyped) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteUInt64(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint64(value))
    return nil
}
