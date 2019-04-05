// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.3.0.0

@file:Suppress("UnusedImport", "unused")

package test

import java.io.*
import java.lang.*
import java.lang.reflect.*
import java.math.*
import java.nio.charset.*
import java.time.*
import java.util.*

import fbe.*
import proto.*

@Suppress("MemberVisibilityCanBePrivate", "RemoveRedundantCallsOfConversionMethods")
class FlagsSimple : Comparable<FlagsSimple>
{
    companion object
    {
        val FLAG_VALUE_0 = FlagsSimple(FlagsSimpleEnum.FLAG_VALUE_0)
        val FLAG_VALUE_1 = FlagsSimple(FlagsSimpleEnum.FLAG_VALUE_1)
        val FLAG_VALUE_2 = FlagsSimple(FlagsSimpleEnum.FLAG_VALUE_2)
        val FLAG_VALUE_3 = FlagsSimple(FlagsSimpleEnum.FLAG_VALUE_3)
        val FLAG_VALUE_4 = FlagsSimple(FlagsSimpleEnum.FLAG_VALUE_4)
        val FLAG_VALUE_5 = FlagsSimple(FlagsSimpleEnum.FLAG_VALUE_5)

        fun fromSet(set: EnumSet<FlagsSimpleEnum>): FlagsSimple
        {
            @Suppress("CanBeVal")
            var result = 0
            if (set.contains(FLAG_VALUE_0.value!!))
            {
                result = result.toInt() or FLAG_VALUE_0.raw.toInt()
            }
            if (set.contains(FLAG_VALUE_1.value!!))
            {
                result = result.toInt() or FLAG_VALUE_1.raw.toInt()
            }
            if (set.contains(FLAG_VALUE_2.value!!))
            {
                result = result.toInt() or FLAG_VALUE_2.raw.toInt()
            }
            if (set.contains(FLAG_VALUE_3.value!!))
            {
                result = result.toInt() or FLAG_VALUE_3.raw.toInt()
            }
            if (set.contains(FLAG_VALUE_4.value!!))
            {
                result = result.toInt() or FLAG_VALUE_4.raw.toInt()
            }
            if (set.contains(FLAG_VALUE_5.value!!))
            {
                result = result.toInt() or FLAG_VALUE_5.raw.toInt()
            }
            return FlagsSimple(result.toInt())
        }
    }

    var value: FlagsSimpleEnum? = FlagsSimpleEnum.values()[0]
        private set

    var raw: Int = value!!.raw
        private set

    constructor()
    constructor(value: Int) { setEnum(value) }
    constructor(value: FlagsSimpleEnum) { setEnum(value) }
    constructor(value: EnumSet<FlagsSimpleEnum>) { setEnum(value) }
    constructor(value: FlagsSimple) { setEnum(value) }

    fun setDefault() { setEnum(0.toInt()) }

    fun setEnum(value: Int) { this.raw = value; this.value = FlagsSimpleEnum.mapValue(value) }
    fun setEnum(value: FlagsSimpleEnum) { this.value = value; this.raw = value.raw; }
    fun setEnum(value: EnumSet<FlagsSimpleEnum>) { setEnum(FlagsSimple.fromSet(value)) }
    fun setEnum(value: FlagsSimple) { this.value = value.value; this.raw = value.raw }

    fun hasFlags(flags: Int): Boolean = ((raw.toInt() and flags.toInt()) != 0) && ((raw.toInt() and flags.toInt()) == flags.toInt())
    fun hasFlags(flags: FlagsSimpleEnum): Boolean = hasFlags(flags.raw)
    fun hasFlags(flags: FlagsSimple): Boolean = hasFlags(flags.raw)

    fun setFlags(flags: Int): FlagsSimple { setEnum((raw.toInt() or flags.toInt()).toInt()); return this }
    fun setFlags(flags: FlagsSimpleEnum): FlagsSimple { setFlags(flags.raw); return this }
    fun setFlags(flags: FlagsSimple): FlagsSimple { setFlags(flags.raw); return this }

    fun removeFlags(flags: Int): FlagsSimple { setEnum((raw.toInt() and flags.toInt().inv()).toInt()); return this }
    fun removeFlags(flags: FlagsSimpleEnum): FlagsSimple { removeFlags(flags.raw); return this }
    fun removeFlags(flags: FlagsSimple): FlagsSimple { removeFlags(flags.raw); return this }

    val allSet: EnumSet<FlagsSimpleEnum> get() = value!!.allSet
    val noneSet: EnumSet<FlagsSimpleEnum> get() = value!!.noneSet
    val currentSet: EnumSet<FlagsSimpleEnum> get() = value!!.currentSet

    override fun compareTo(other: FlagsSimple): Int
    {
        return (raw - other.raw).toInt()
    }

    override fun equals(other: Any?): Boolean
    {
        if (other == null)
            return false

        if (!FlagsSimple::class.java.isAssignableFrom(other.javaClass))
            return false

        val flg = other as FlagsSimple? ?: return false

        if (raw != flg.raw)
            return false
        return true
    }

    override fun hashCode(): Int
    {
        var hash = 17
        hash = hash * 31 + raw.toInt()
        return hash.toInt()
    }

    override fun toString(): String
    {
        val sb = StringBuilder()
        var first = true
        if (hasFlags(FLAG_VALUE_0.raw))
        {
            sb.append(if (first) "" else "|").append("FLAG_VALUE_0")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        if (hasFlags(FLAG_VALUE_1.raw))
        {
            sb.append(if (first) "" else "|").append("FLAG_VALUE_1")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        if (hasFlags(FLAG_VALUE_2.raw))
        {
            sb.append(if (first) "" else "|").append("FLAG_VALUE_2")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        if (hasFlags(FLAG_VALUE_3.raw))
        {
            sb.append(if (first) "" else "|").append("FLAG_VALUE_3")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        if (hasFlags(FLAG_VALUE_4.raw))
        {
            sb.append(if (first) "" else "|").append("FLAG_VALUE_4")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        if (hasFlags(FLAG_VALUE_5.raw))
        {
            sb.append(if (first) "" else "|").append("FLAG_VALUE_5")
            @Suppress("UNUSED_VALUE")
            first = false
        }
        return sb.toString()
    }
}
