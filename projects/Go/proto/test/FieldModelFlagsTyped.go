// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

package test

import "errors"
import "../fbe"

// Fast Binary Encoding FlagsTyped field model class
type FieldModelFlagsTyped struct {
    buffer *fbe.Buffer  // Field model buffer
    offset int          // Field model buffer offset
}

// Get the field size
func (fm FieldModelFlagsTyped) FBESize() int { return 8 }
// Get the field extra size
func (fm FieldModelFlagsTyped) FBEExtra() int { return 0 }

// Get the field offset
func (fm FieldModelFlagsTyped) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelFlagsTyped) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelFlagsTyped) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelFlagsTyped) FBEUnshift(size int) { fm.offset -= size }

// Create a new field model
func NewFieldModelFlagsTyped(buffer *fbe.Buffer, offset int) *FieldModelFlagsTyped {
    return &FieldModelFlagsTyped{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FieldModelFlagsTyped) Verify() bool { return true }

// Get the value
func (fm FieldModelFlagsTyped) Get() (FlagsTyped, error) {
    return fm.GetDefault(0)
}

// Get the value with provided default value
func (fm FieldModelFlagsTyped) GetDefault(defaults FlagsTyped) (FlagsTyped, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return FlagsTyped(0), nil
    }

    return FlagsTyped(fbe.ReadUInt64(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), nil
}

// Set the value
func (fm *FieldModelFlagsTyped) Set(value FlagsTyped) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteUInt64(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint64(value))
    return nil
}
