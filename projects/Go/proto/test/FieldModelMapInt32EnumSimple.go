// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.3.0.0

package test

import "errors"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// Fast Binary Encoding Int32->EnumSimple map field model
type FieldModelMapInt32EnumSimple struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int

    // Map key field model
    modelKey *fbe.FieldModelInt32
    // Map value field model
    modelValue *FieldModelEnumSimple
}

// Create a new Int32->EnumSimple map field model
func NewFieldModelMapInt32EnumSimple(buffer *fbe.Buffer, offset int) *FieldModelMapInt32EnumSimple {
    fbeResult := FieldModelMapInt32EnumSimple{buffer: buffer, offset: offset}
    fbeResult.modelKey = fbe.NewFieldModelInt32(buffer, offset)
    fbeResult.modelValue = NewFieldModelEnumSimple(buffer, offset)
    return &fbeResult
}

// Get the field size
func (fm *FieldModelMapInt32EnumSimple) FBESize() int { return 4 }

// Get the field extra size
func (fm *FieldModelMapInt32EnumSimple) FBEExtra() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0
    }

    fbeMapOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeMapOffset == 0) || ((fm.buffer.Offset() + fbeMapOffset + 4) > fm.buffer.Size()) {
        return 0
    }

    fbeMapSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeMapOffset))

    fbeResult := 0
    fm.modelKey.SetFBEOffset(fbeMapOffset + 4)
    fm.modelValue.SetFBEOffset(fbeMapOffset + 4 + fm.modelKey.FBESize())
    for i := fbeMapSize; i > 0; i-- {
        fbeResult += fm.modelKey.FBESize() + fm.modelKey.FBEExtra()
        fm.modelKey.FBEShift(fm.modelKey.FBESize() + fm.modelValue.FBESize())

        fbeResult += fm.modelValue.FBESize() + fm.modelValue.FBEExtra()
        fm.modelValue.FBEShift(fm.modelKey.FBESize() + fm.modelValue.FBESize())
    }
    return fbeResult
}

// Get the field offset
func (fm *FieldModelMapInt32EnumSimple) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelMapInt32EnumSimple) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelMapInt32EnumSimple) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelMapInt32EnumSimple) FBEUnshift(size int) { fm.offset -= size }

// Get the map offset
func (fm *FieldModelMapInt32EnumSimple) Offset() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0
    }

    fbeMapOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    return fbeMapOffset
}

// Get the map size
func (fm *FieldModelMapInt32EnumSimple) Size() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0
    }

    fbeMapOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeMapOffset == 0) || ((fm.buffer.Offset() + fbeMapOffset + 4) > fm.buffer.Size()) {
        return 0
    }

    fbeMapSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeMapOffset))
    return fbeMapSize
}

// Map index operator
func (fm *FieldModelMapInt32EnumSimple) GetItem(index int) (*fbe.FieldModelInt32, *FieldModelEnumSimple, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return nil, nil, errors.New("model is broken")
    }

    fbeMapOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeMapOffset == 0) || ((fm.buffer.Offset() + fbeMapOffset + 4) > fm.buffer.Size()) {
        return nil, nil, errors.New("model is broken")
    }

    fbeMapSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeMapOffset))
    if index >= fbeMapSize {
        return nil, nil, errors.New("index is out of bounds")
    }

    fm.modelKey.SetFBEOffset(fbeMapOffset + 4)
    fm.modelValue.SetFBEOffset(fbeMapOffset + 4 + fm.modelKey.FBESize())
    fm.modelKey.FBEShift(index * (fm.modelKey.FBESize() + fm.modelValue.FBESize()))
    fm.modelValue.FBEShift(index * (fm.modelKey.FBESize() + fm.modelValue.FBESize()))
    return fm.modelKey, fm.modelValue, nil
}

// Resize the map and get its first model
func (fm *FieldModelMapInt32EnumSimple) Resize(size int) (*fbe.FieldModelInt32, *FieldModelEnumSimple, error) {
    fbeMapSize := size * (fm.modelKey.FBESize() + fm.modelValue.FBESize())
    fbeMapOffset := fm.buffer.Allocate(4 + fbeMapSize) - fm.buffer.Offset()
    if (fbeMapOffset == 0) || ((fm.buffer.Offset() + fbeMapOffset + 4) > fm.buffer.Size()) {
        return nil, nil, errors.New("model is broken")
    }

    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint32(fbeMapOffset))
    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeMapOffset, uint32(size))
    fbe.WriteCount(fm.buffer.Data(), fm.buffer.Offset() + fbeMapOffset + 4, 0, fbeMapSize)

    fm.modelKey.SetFBEOffset(fbeMapOffset + 4)
    fm.modelValue.SetFBEOffset(fbeMapOffset + 4 + fm.modelKey.FBESize())
    return fm.modelKey, fm.modelValue, nil
}

// Check if the map is valid
func (fm *FieldModelMapInt32EnumSimple) Verify() bool {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return true
    }

    fbeMapOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if fbeMapOffset == 0 {
        return true
    }

    if (fm.buffer.Offset() + fbeMapOffset + 4) > fm.buffer.Size() {
        return false
    }

    fbeMapSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeMapOffset))

    fm.modelKey.SetFBEOffset(fbeMapOffset + 4)
    fm.modelValue.SetFBEOffset(fbeMapOffset + 4 + fm.modelKey.FBESize())
    for i := fbeMapSize; i > 0; i-- {
        if !fm.modelKey.Verify() {
            return false
        }
        fm.modelKey.FBEShift(fm.modelKey.FBESize() + fm.modelValue.FBESize())
        if !fm.modelValue.Verify() {
            return false
        }
        fm.modelValue.FBEShift(fm.modelKey.FBESize() + fm.modelValue.FBESize())
    }

    return true
}

// Get the map
func (fm *FieldModelMapInt32EnumSimple) Get() (map[int32]EnumSimple, error) {
    values := make(map[int32]EnumSimple)

    fbeMapSize := fm.Size()
    if fbeMapSize == 0 {
        return values, nil
    }

    fbeModelKey, fbeModelValue, err := fm.GetItem(0)
    if err != nil {
        return values, err
    }

    for i := fbeMapSize; i > 0; i-- {
        key, err := fbeModelKey.Get()
        if err != nil {
            return values, err
        }
        value, err := fbeModelValue.Get()
        if err != nil {
            return values, err
        }
        values[key] = *value
        fbeModelKey.FBEShift(fbeModelKey.FBESize() + fbeModelValue.FBESize())
        fbeModelValue.FBEShift(fbeModelKey.FBESize() + fbeModelValue.FBESize())
    }

    return values, nil
}

// Set the map
func (fm *FieldModelMapInt32EnumSimple) Set(values map[int32]EnumSimple) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbeModelKey, fbeModelValue, err := fm.Resize(len(values))
    if err != nil {
        return err
    }

    for key, value := range values {
        err := fbeModelKey.Set(key)
        if err != nil {
            return err
        }
        fbeModelKey.FBEShift(fbeModelKey.FBESize() + fbeModelValue.FBESize())
        err = fbeModelValue.Set(&value)
        if err != nil {
            return err
        }
        fbeModelValue.FBEShift(fbeModelKey.FBESize() + fbeModelValue.FBESize())
    }

    return nil
}
