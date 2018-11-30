// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

package test

import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// Fast Binary Encoding StructHashEx final model
type FinalModelStructHashEx struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset

    F1 *FinalModelMapStructSimpleStructNested
    F2 *FinalModelMapStructSimpleOptionalStructNested
}

// Create a new StructHashEx final model
func NewFinalModelStructHashEx(buffer *fbe.Buffer, offset int) *FinalModelStructHashEx {
    fbeResult := FinalModelStructHashEx{buffer: buffer, offset: offset}
    fbeResult.F1 = NewFinalModelMapStructSimpleStructNested(buffer, 0)
    fbeResult.F2 = NewFinalModelMapStructSimpleOptionalStructNested(buffer, 0)
    return &fbeResult
}

// Get the allocation size
func (fm *FinalModelStructHashEx) FBEAllocationSize(fbeValue *StructHashEx) int {
    fbeResult := 0 +
        fm.F1.FBEAllocationSize(fbeValue.F1) +
        fm.F2.FBEAllocationSize(fbeValue.F2) +
        0
    return fbeResult
}

// Get the final size
func (fm *FinalModelStructHashEx) FBESize() int { return 0 }

// Get the final extra size
func (fm *FinalModelStructHashEx) FBEExtra() int { return 0 }

// Get the final type
func (fm *FinalModelStructHashEx) FBEType() int { return 142 }

// Get the final offset
func (fm *FinalModelStructHashEx) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelStructHashEx) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelStructHashEx) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelStructHashEx) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FinalModelStructHashEx) Verify() int {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult := fm.VerifyFields()
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult
}

// Check if the struct fields are valid
func (fm *FinalModelStructHashEx) VerifyFields() int {
    fbeCurrentOffset := 0
    fbeFieldSize := 0


    fm.F1.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize := fm.F1.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F2.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize := fm.F2.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    return fbeCurrentOffset
}

// Get the struct value
func (fm *FinalModelStructHashEx) Get() (*StructHashEx, int, error) {
    fbeResult := NewStructHashEx()
    fbeSize, err := fm.GetValue(fbeResult)
    return fbeResult, fbeSize, err
}

// Get the struct value by the given pointer
func (fm *FinalModelStructHashEx) GetValue(fbeValue *StructHashEx) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeSize, err := fm.GetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeSize, err
}

// Get the struct fields values
func (fm *FinalModelStructHashEx) GetFields(fbeValue *StructHashEx) (int, error) {
    var err error = nil
    fbeCurrentOffset := 0
    fbeCurrentSize := 0
    fbeFieldSize := 0

    fm.F1.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1.GetValue(fbeValue.F1); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F2.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F2.GetValue(fbeValue.F2); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, nil
}

// Set the struct value
func (fm *FinalModelStructHashEx) Set(fbeValue *StructHashEx) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult, err := fm.SetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult, err
}

// Set the struct fields values
func (fm *FinalModelStructHashEx) SetFields(fbeValue *StructHashEx) (int, error) {
    var err error = nil
    fbeCurrentOffset := 0
    fbeCurrentSize := 0
    fbeFieldSize := 0

    fm.F1.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1.Set(fbeValue.F1); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F2.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F2.Set(fbeValue.F2); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, nil
}
