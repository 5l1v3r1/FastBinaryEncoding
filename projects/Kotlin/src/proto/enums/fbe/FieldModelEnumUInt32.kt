// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding

package enums.fbe;

import java.io.*;
import java.lang.*;
import java.lang.reflect.*;
import java.math.*;
import java.nio.charset.*;
import java.time.*;
import java.util.*;

import fbe.*;
import enums.*;

// Fast Binary Encoding EnumUInt32 field model class
class FieldModelEnumUInt32(buffer: Buffer, offset: Long) : FieldModel(buffer, offset)
{
    // Field size
    override val FBESize: Long = 4

    // Get the value
    fun get(defaults: EnumUInt32 = EnumUInt32()): EnumUInt32
    {
        if (_buffer.offset + FBEOffset + FBESize > _buffer.size)
            return defaults

        return EnumUInt32(readUInt32(FBEOffset))
    }

    // Set the value
    fun set(value: EnumUInt32)
    {
        assert(_buffer.offset + FBEOffset + FBESize <= _buffer.size) { "Model is broken!" }
        if (_buffer.offset + FBEOffset + FBESize > _buffer.size)
            return

        write(FBEOffset, value.raw)
    }
}