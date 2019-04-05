// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.3.0.0

package fbe

import "errors"

// Fast Binary Encoding Bool final model
type FinalModelBool struct {
    // Final model buffer
    buffer *Buffer
    // Final model buffer offset
    offset int
}

// Create a new Bool final model
func NewFinalModelBool(buffer *Buffer, offset int) *FinalModelBool {
    return &FinalModelBool{buffer: buffer, offset: offset}
}

// Get the allocation size
func (fm *FinalModelBool) FBEAllocationSize(value bool) int { return fm.FBESize() }

// Get the final size
func (fm *FinalModelBool) FBESize() int { return 1 }

// Get the final offset
func (fm *FinalModelBool) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelBool) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelBool) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelBool) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FinalModelBool) Verify() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return MaxInt
    }

    return fm.FBESize()
}

// Get the value
func (fm *FinalModelBool) Get() (bool, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0, errors.New("model is broken")
    }

    return ReadBool(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelBool) Set(value bool) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    WriteBool(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
    return fm.FBESize(), nil
}
