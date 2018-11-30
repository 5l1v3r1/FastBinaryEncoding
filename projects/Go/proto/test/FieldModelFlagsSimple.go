// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

package test

import "errors"
import "../fbe"

// Fast Binary Encoding FlagsSimple field model
type FieldModelFlagsSimple struct {
    buffer *fbe.Buffer  // Field model buffer
    offset int          // Field model buffer offset
}

// Create a new field model
func NewFieldModelFlagsSimple(buffer *fbe.Buffer, offset int) *FieldModelFlagsSimple {
    return &FieldModelFlagsSimple{buffer: buffer, offset: offset}
}

// Get the field size
func (fm *FieldModelFlagsSimple) FBESize() int { return 4 }
// Get the field extra size
func (fm *FieldModelFlagsSimple) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelFlagsSimple) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelFlagsSimple) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelFlagsSimple) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelFlagsSimple) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FieldModelFlagsSimple) Verify() bool { return true }

// Get the value
func (fm *FieldModelFlagsSimple) Get() (*FlagsSimple, error) {
    var value FlagsSimple
    return &value, fm.GetValueDefault(&value, FlagsSimple(0))
}

// Get the value with provided default value
func (fm *FieldModelFlagsSimple) GetDefault(defaults FlagsSimple) (*FlagsSimple, error) {
    var value FlagsSimple
    err := fm.GetValueDefault(&value, defaults)
    return &value, err
}

// Get the value by the given pointer
func (fm *FieldModelFlagsSimple) GetValue(value *FlagsSimple) error {
    return fm.GetValueDefault(value, FlagsSimple(0))
}

// Get the value by the given pointer with provided default value
func (fm *FieldModelFlagsSimple) GetValueDefault(value *FlagsSimple, defaults FlagsSimple) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        *value = defaults
        return nil
    }

    *value = FlagsSimple(fbe.ReadInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    return nil
}

// Set the value
func (fm *FieldModelFlagsSimple) Set(value *FlagsSimple) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), int32(*value))
    return nil
}
