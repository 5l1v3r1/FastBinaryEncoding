// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.3.0.0

@file:Suppress("UnusedImport", "unused")

package test.fbe

import java.io.*
import java.lang.*
import java.lang.reflect.*
import java.math.*
import java.nio.charset.*
import java.time.*
import java.util.*

import fbe.*
import test.*

// Fast Binary Encoding StructArray final model
@Suppress("MemberVisibilityCanBePrivate", "RemoveRedundantCallsOfConversionMethods", "ReplaceGetOrSet")
class FinalModelStructArray(buffer: Buffer, offset: Long) : FinalModel(buffer, offset)
{
    val f1: FinalModelArrayByte = FinalModelArrayByte(buffer, 0, 2)
    val f2: FinalModelArrayOptionalByte = FinalModelArrayOptionalByte(buffer, 0, 2)
    val f3: FinalModelArrayBytes = FinalModelArrayBytes(buffer, 0, 2)
    val f4: FinalModelArrayOptionalBytes = FinalModelArrayOptionalBytes(buffer, 0, 2)
    val f5: FinalModelArrayEnumSimple = FinalModelArrayEnumSimple(buffer, 0, 2)
    val f6: FinalModelArrayOptionalEnumSimple = FinalModelArrayOptionalEnumSimple(buffer, 0, 2)
    val f7: FinalModelArrayFlagsSimple = FinalModelArrayFlagsSimple(buffer, 0, 2)
    val f8: FinalModelArrayOptionalFlagsSimple = FinalModelArrayOptionalFlagsSimple(buffer, 0, 2)
    val f9: FinalModelArrayStructSimple = FinalModelArrayStructSimple(buffer, 0, 2)
    val f10: FinalModelArrayOptionalStructSimple = FinalModelArrayOptionalStructSimple(buffer, 0, 2)

    // Get the allocation size
    @Suppress("UNUSED_PARAMETER")
    fun fbeAllocationSize(fbeValue: StructArray): Long = (0
        + f1.fbeAllocationSize(fbeValue.f1)
        + f2.fbeAllocationSize(fbeValue.f2)
        + f3.fbeAllocationSize(fbeValue.f3)
        + f4.fbeAllocationSize(fbeValue.f4)
        + f5.fbeAllocationSize(fbeValue.f5)
        + f6.fbeAllocationSize(fbeValue.f6)
        + f7.fbeAllocationSize(fbeValue.f7)
        + f8.fbeAllocationSize(fbeValue.f8)
        + f9.fbeAllocationSize(fbeValue.f9)
        + f10.fbeAllocationSize(fbeValue.f10)
        )

    // Field type
    var fbeType: Long = fbeTypeConst

    companion object
    {
        const val fbeTypeConst: Long = 125
    }

    // Check if the struct value is valid
    override fun verify(): Long
    {
        _buffer.shift(fbeOffset)
        val fbeResult = verifyFields()
        _buffer.unshift(fbeOffset)
        return fbeResult
    }

    // Check if the struct fields are valid
    fun verifyFields(): Long
    {
        var fbeCurrentOffset = 0L
        @Suppress("VARIABLE_WITH_REDUNDANT_INITIALIZER")
        var fbeFieldSize = 0L

        f1.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f1.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f2.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f2.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f3.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f3.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f4.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f4.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f5.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f5.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f6.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f6.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f7.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f7.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f8.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f8.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f9.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f9.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        f10.fbeOffset = fbeCurrentOffset
        fbeFieldSize = f10.verify()
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE
        fbeCurrentOffset += fbeFieldSize

        return fbeCurrentOffset
    }

    // Get the struct value
    fun get(fbeSize: Size, fbeValue: StructArray = StructArray()): StructArray
    {
        _buffer.shift(fbeOffset)
        fbeSize.value = getFields(fbeValue)
        _buffer.unshift(fbeOffset)
        return fbeValue
    }

    // Get the struct fields values
    @Suppress("UNUSED_PARAMETER")
    fun getFields(fbeValue: StructArray): Long
    {
        var fbeCurrentOffset = 0L
        var fbeCurrentSize = 0L
        val fbeFieldSize = Size(0)

        f1.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f1.get(fbeValue.f1)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f2.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f2.get(fbeValue.f2)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f3.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f3.get(fbeValue.f3)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f4.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f4.get(fbeValue.f4)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f5.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f5.get(fbeValue.f5)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f6.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f6.get(fbeValue.f6)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f7.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f7.get(fbeValue.f7)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f8.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f8.get(fbeValue.f8)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f9.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f9.get(fbeValue.f9)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f10.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f10.get(fbeValue.f10)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        return fbeCurrentSize
    }

    // Set the struct value
    fun set(fbeValue: StructArray): Long
    {
        _buffer.shift(fbeOffset)
        val fbeSize = setFields(fbeValue)
        _buffer.unshift(fbeOffset)
        return fbeSize
    }

    // Set the struct fields values
    @Suppress("UNUSED_PARAMETER")
    fun setFields(fbeValue: StructArray): Long
    {
        var fbeCurrentOffset = 0L
        var fbeCurrentSize = 0L
        val fbeFieldSize = Size(0)

        f1.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f1.set(fbeValue.f1)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f2.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f2.set(fbeValue.f2)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f3.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f3.set(fbeValue.f3)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f4.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f4.set(fbeValue.f4)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f5.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f5.set(fbeValue.f5)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f6.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f6.set(fbeValue.f6)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f7.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f7.set(fbeValue.f7)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f8.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f8.set(fbeValue.f8)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f9.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f9.set(fbeValue.f9)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        f10.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = f10.set(fbeValue.f10)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        return fbeCurrentSize
    }
}
