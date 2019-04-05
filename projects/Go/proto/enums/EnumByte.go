// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.3.0.0

package enums

import "../fbe"

// Workaround for Go unused imports issue
var _ = fbe.Version

// EnumByte enum key
type EnumByteKey byte

// Convert EnumByte enum key to string
func (k EnumByteKey) String() string {
    return EnumByte(k).String()
}

// EnumByte enum
type EnumByte byte

// EnumByte enum values
//noinspection GoSnakeCaseUsage
const (
    EnumByte_ENUM_VALUE_0 = EnumByte(0 + 0)
    EnumByte_ENUM_VALUE_1 = EnumByte(0 + 0)
    EnumByte_ENUM_VALUE_2 = EnumByte(0 + 1)
    EnumByte_ENUM_VALUE_3 = EnumByte(254 + 0)
    EnumByte_ENUM_VALUE_4 = EnumByte(254 + 1)
    EnumByte_ENUM_VALUE_5 = EnumByte(EnumByte_ENUM_VALUE_3)
)

// Create a new EnumByte enum
func NewEnumByte() *EnumByte {
    return new(EnumByte)
}

// Create a new EnumByte enum from the given value
func NewEnumByteFromValue(value byte) *EnumByte {
    result := EnumByte(value)
    return &result
}

// Get the enum key
func (e EnumByte) Key() EnumByteKey {
    return EnumByteKey(e)
}

// Convert enum to optional
func (e *EnumByte) Optional() *EnumByte {
    return e
}

// Convert enum to string
func (e EnumByte) String() string {
    if e == EnumByte_ENUM_VALUE_0 {
        return "ENUM_VALUE_0"
    }
    if e == EnumByte_ENUM_VALUE_1 {
        return "ENUM_VALUE_1"
    }
    if e == EnumByte_ENUM_VALUE_2 {
        return "ENUM_VALUE_2"
    }
    if e == EnumByte_ENUM_VALUE_3 {
        return "ENUM_VALUE_3"
    }
    if e == EnumByte_ENUM_VALUE_4 {
        return "ENUM_VALUE_4"
    }
    if e == EnumByte_ENUM_VALUE_5 {
        return "ENUM_VALUE_5"
    }
    return "<unknown>"
}

// Convert enum to JSON
func (e EnumByte) MarshalJSON() ([]byte, error) {
    value := byte(e)
    return fbe.Json.Marshal(&value)
}

// Convert JSON to enum
func (e *EnumByte) UnmarshalJSON(buffer []byte) error {
    var result byte
    err := fbe.Json.Unmarshal(buffer, &result)
    if err != nil {
        return err
    }
    *e = EnumByte(result)
    return nil
}
