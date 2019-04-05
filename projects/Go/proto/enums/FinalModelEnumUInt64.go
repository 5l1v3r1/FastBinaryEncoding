// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.3.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumUInt64 final model
type FinalModelEnumUInt64 struct {
    // Final model buffer
    buffer *fbe.Buffer
    // Final model buffer offset
    offset int
}

// Create a new EnumUInt64 final model
func NewFinalModelEnumUInt64(buffer *fbe.Buffer, offset int) *FinalModelEnumUInt64 {
    return &FinalModelEnumUInt64{buffer: buffer, offset: offset}
}

// Get the allocation size
func (fm *FinalModelEnumUInt64) FBEAllocationSize(value *EnumUInt64) int { return fm.FBESize() }

// Get the final size
func (fm *FinalModelEnumUInt64) FBESize() int { return 8 }

// Get the final offset
func (fm *FinalModelEnumUInt64) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelEnumUInt64) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelEnumUInt64) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelEnumUInt64) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FinalModelEnumUInt64) Verify() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return fbe.MaxInt
    }

    return fm.FBESize()
}

// Get the value
func (fm *FinalModelEnumUInt64) Get() (*EnumUInt64, int, error) {
    var value EnumUInt64
    size, err := fm.GetValueDefault(&value, EnumUInt64(0))
    return &value, size, err
}

// Get the value with provided default value
func (fm *FinalModelEnumUInt64) GetDefault(defaults EnumUInt64) (*EnumUInt64, int, error) {
    var value EnumUInt64
    size, err := fm.GetValueDefault(&value, defaults)
    return &value, size, err
}

// Get the value by the given pointer
func (fm *FinalModelEnumUInt64) GetValue(value *EnumUInt64) (int, error) {
    return fm.GetValueDefault(value, EnumUInt64(0))
}

// Get the value by the given pointer with provided default value
func (fm *FinalModelEnumUInt64) GetValueDefault(value *EnumUInt64, defaults EnumUInt64) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        *value = defaults
        return 0, errors.New("model is broken")
    }

    *value = EnumUInt64(fbe.ReadUInt64(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    return fm.FBESize(), nil
}

// Set the value by the given pointer
func (fm *FinalModelEnumUInt64) Set(value *EnumUInt64) (int, error) {
    return fm.SetValue(*value)
}

// Set the value
func (fm *FinalModelEnumUInt64) SetValue(value EnumUInt64) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteUInt64(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint64(value))
    return fm.FBESize(), nil
}
