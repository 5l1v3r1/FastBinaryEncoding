// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.3.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumInt16 field model
type FieldModelEnumInt16 struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int
}

// Create a new EnumInt16 field model
func NewFieldModelEnumInt16(buffer *fbe.Buffer, offset int) *FieldModelEnumInt16 {
    return &FieldModelEnumInt16{buffer: buffer, offset: offset}
}

// Get the field size
func (fm *FieldModelEnumInt16) FBESize() int { return 2 }
// Get the field extra size
func (fm *FieldModelEnumInt16) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelEnumInt16) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelEnumInt16) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelEnumInt16) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelEnumInt16) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FieldModelEnumInt16) Verify() bool { return true }

// Get the value
func (fm *FieldModelEnumInt16) Get() (*EnumInt16, error) {
    var value EnumInt16
    return &value, fm.GetValueDefault(&value, EnumInt16(0))
}

// Get the value with provided default value
func (fm *FieldModelEnumInt16) GetDefault(defaults EnumInt16) (*EnumInt16, error) {
    var value EnumInt16
    err := fm.GetValueDefault(&value, defaults)
    return &value, err
}

// Get the value by the given pointer
func (fm *FieldModelEnumInt16) GetValue(value *EnumInt16) error {
    return fm.GetValueDefault(value, EnumInt16(0))
}

// Get the value by the given pointer with provided default value
func (fm *FieldModelEnumInt16) GetValueDefault(value *EnumInt16, defaults EnumInt16) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        *value = defaults
        return nil
    }

    *value = EnumInt16(fbe.ReadInt16(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    return nil
}

// Set the value by the given pointer
func (fm *FieldModelEnumInt16) Set(value *EnumInt16) error {
    return fm.SetValue(*value)
}

// Set the value
func (fm *FieldModelEnumInt16) SetValue(value EnumInt16) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteInt16(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), int16(value))
    return nil
}
