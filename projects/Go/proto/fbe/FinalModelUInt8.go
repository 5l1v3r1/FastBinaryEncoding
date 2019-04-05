// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.3.0.0

package fbe

import "errors"

// Fast Binary Encoding UInt8 final model
type FinalModelUInt8 struct {
    // Final model buffer
    buffer *Buffer
    // Final model buffer offset
    offset int
}

// Create a new UInt8 final model
func NewFinalModelUInt8(buffer *Buffer, offset int) *FinalModelUInt8 {
    return &FinalModelUInt8{buffer: buffer, offset: offset}
}

// Get the allocation size
func (fm *FinalModelUInt8) FBEAllocationSize(value uint8) int { return fm.FBESize() }

// Get the final size
func (fm *FinalModelUInt8) FBESize() int { return 1 }

// Get the final offset
func (fm *FinalModelUInt8) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelUInt8) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelUInt8) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelUInt8) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FinalModelUInt8) Verify() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return MaxInt
    }

    return fm.FBESize()
}

// Get the value
func (fm *FinalModelUInt8) Get() (uint8, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, 0, errors.New("model is broken")
    }

    return ReadUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelUInt8) Set(value uint8) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    WriteUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
    return fm.FBESize(), nil
}
