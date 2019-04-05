// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.3.0.0

package fbe

import "errors"

// Fast Binary Encoding Double field model
type FieldModelDouble struct {
    // Field model buffer
    buffer *Buffer
    // Field model buffer offset
    offset int
}

// Create a new Double field model
func NewFieldModelDouble(buffer *Buffer, offset int) *FieldModelDouble {
    return &FieldModelDouble{buffer: buffer, offset: offset}
}

// Get the field size
func (fm *FieldModelDouble) FBESize() int { return 8 }
// Get the field extra size
func (fm *FieldModelDouble) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelDouble) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelDouble) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelDouble) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelDouble) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FieldModelDouble) Verify() bool { return true }

// Get the value
func (fm *FieldModelDouble) Get() (float64, error) {
    return fm.GetDefault(0.0)
}

// Get the value with provided default value
func (fm *FieldModelDouble) GetDefault(defaults float64) (float64, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return defaults, nil
    }

    return ReadDouble(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()), nil
}

// Set the value
func (fm *FieldModelDouble) Set(value float64) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    WriteDouble(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
    return nil
}
