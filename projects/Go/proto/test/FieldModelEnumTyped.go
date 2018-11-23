// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

package test

import "errors"
import "../fbe"

// Fast Binary Encoding EnumTyped field model class
type FieldModelEnumTyped struct {
    buffer *fbe.Buffer  // Field model buffer
    offset int          // Field model buffer offset
}

// Get the field size
func (fm FieldModelEnumTyped) FBESize() int { return 1 }
// Get the field extra size
func (fm FieldModelEnumTyped) FBEExtra() int { return 0 }

// Get the field offset
func (fm FieldModelEnumTyped) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelEnumTyped) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelEnumTyped) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelEnumTyped) FBEUnshift(size int) { fm.offset -= size }

// Create a new field model
func NewFieldModelEnumTyped(buffer *fbe.Buffer, offset int) *FieldModelEnumTyped {
    return &FieldModelEnumTyped{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FieldModelEnumTyped) Verify() bool { return true }

// Get the value
func (fm FieldModelEnumTyped) Get() (EnumTyped, error) {
    return fm.GetDefault(0)
}

// Get the value with provided default value
func (fm FieldModelEnumTyped) GetDefault(defaults EnumTyped) (EnumTyped, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return EnumTyped(0), nil
    }

    return EnumTyped(fbe.ReadUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), nil
}

// Set the value
func (fm *FieldModelEnumTyped) Set(value EnumTyped) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint8(value))
    return nil
}
