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

// Fast Binary Encoding test final receiver
@Suppress("MemberVisibilityCanBePrivate", "PrivatePropertyName", "UNUSED_PARAMETER")
open class FinalReceiver : fbe.Receiver
{
    // Imported receivers
    var protoReceiver: proto.fbe.FinalReceiver? = null

    // Receiver values accessors
    private val StructSimpleValue: test.StructSimple
    private val StructOptionalValue: test.StructOptional
    private val StructNestedValue: test.StructNested
    private val StructBytesValue: test.StructBytes
    private val StructArrayValue: test.StructArray
    private val StructVectorValue: test.StructVector
    private val StructListValue: test.StructList
    private val StructSetValue: test.StructSet
    private val StructMapValue: test.StructMap
    private val StructHashValue: test.StructHash
    private val StructHashExValue: test.StructHashEx
    private val StructEmptyValue: test.StructEmpty

    // Receiver models accessors
    private val StructSimpleModel: StructSimpleFinalModel
    private val StructOptionalModel: StructOptionalFinalModel
    private val StructNestedModel: StructNestedFinalModel
    private val StructBytesModel: StructBytesFinalModel
    private val StructArrayModel: StructArrayFinalModel
    private val StructVectorModel: StructVectorFinalModel
    private val StructListModel: StructListFinalModel
    private val StructSetModel: StructSetFinalModel
    private val StructMapModel: StructMapFinalModel
    private val StructHashModel: StructHashFinalModel
    private val StructHashExModel: StructHashExFinalModel
    private val StructEmptyModel: StructEmptyFinalModel

    constructor() : super(true)
    {
        protoReceiver = proto.fbe.FinalReceiver(buffer)
        StructSimpleValue = test.StructSimple()
        StructSimpleModel = StructSimpleFinalModel()
        StructOptionalValue = test.StructOptional()
        StructOptionalModel = StructOptionalFinalModel()
        StructNestedValue = test.StructNested()
        StructNestedModel = StructNestedFinalModel()
        StructBytesValue = test.StructBytes()
        StructBytesModel = StructBytesFinalModel()
        StructArrayValue = test.StructArray()
        StructArrayModel = StructArrayFinalModel()
        StructVectorValue = test.StructVector()
        StructVectorModel = StructVectorFinalModel()
        StructListValue = test.StructList()
        StructListModel = StructListFinalModel()
        StructSetValue = test.StructSet()
        StructSetModel = StructSetFinalModel()
        StructMapValue = test.StructMap()
        StructMapModel = StructMapFinalModel()
        StructHashValue = test.StructHash()
        StructHashModel = StructHashFinalModel()
        StructHashExValue = test.StructHashEx()
        StructHashExModel = StructHashExFinalModel()
        StructEmptyValue = test.StructEmpty()
        StructEmptyModel = StructEmptyFinalModel()
    }

    constructor(buffer: Buffer) : super(buffer, true)
    {
        protoReceiver = proto.fbe.FinalReceiver(buffer)
        StructSimpleValue = test.StructSimple()
        StructSimpleModel = StructSimpleFinalModel()
        StructOptionalValue = test.StructOptional()
        StructOptionalModel = StructOptionalFinalModel()
        StructNestedValue = test.StructNested()
        StructNestedModel = StructNestedFinalModel()
        StructBytesValue = test.StructBytes()
        StructBytesModel = StructBytesFinalModel()
        StructArrayValue = test.StructArray()
        StructArrayModel = StructArrayFinalModel()
        StructVectorValue = test.StructVector()
        StructVectorModel = StructVectorFinalModel()
        StructListValue = test.StructList()
        StructListModel = StructListFinalModel()
        StructSetValue = test.StructSet()
        StructSetModel = StructSetFinalModel()
        StructMapValue = test.StructMap()
        StructMapModel = StructMapFinalModel()
        StructHashValue = test.StructHash()
        StructHashModel = StructHashFinalModel()
        StructHashExValue = test.StructHashEx()
        StructHashExModel = StructHashExFinalModel()
        StructEmptyValue = test.StructEmpty()
        StructEmptyModel = StructEmptyFinalModel()
    }

    // Receive handlers
    protected open fun onReceive(value: test.StructSimple) {}
    protected open fun onReceive(value: test.StructOptional) {}
    protected open fun onReceive(value: test.StructNested) {}
    protected open fun onReceive(value: test.StructBytes) {}
    protected open fun onReceive(value: test.StructArray) {}
    protected open fun onReceive(value: test.StructVector) {}
    protected open fun onReceive(value: test.StructList) {}
    protected open fun onReceive(value: test.StructSet) {}
    protected open fun onReceive(value: test.StructMap) {}
    protected open fun onReceive(value: test.StructHash) {}
    protected open fun onReceive(value: test.StructHashEx) {}
    protected open fun onReceive(value: test.StructEmpty) {}

    override fun onReceive(type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {
        when (type)
        {
            test.fbe.StructSimpleFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                StructSimpleModel.attach(buffer, offset)
                assert(StructSimpleModel.verify()) { "test.StructSimple validation failed!" }
                val deserialized = StructSimpleModel.deserialize(StructSimpleValue)
                assert(deserialized > 0) { "test.StructSimple deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = StructSimpleValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                onReceive(StructSimpleValue)
                return true
            }
            test.fbe.StructOptionalFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                StructOptionalModel.attach(buffer, offset)
                assert(StructOptionalModel.verify()) { "test.StructOptional validation failed!" }
                val deserialized = StructOptionalModel.deserialize(StructOptionalValue)
                assert(deserialized > 0) { "test.StructOptional deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = StructOptionalValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                onReceive(StructOptionalValue)
                return true
            }
            test.fbe.StructNestedFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                StructNestedModel.attach(buffer, offset)
                assert(StructNestedModel.verify()) { "test.StructNested validation failed!" }
                val deserialized = StructNestedModel.deserialize(StructNestedValue)
                assert(deserialized > 0) { "test.StructNested deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = StructNestedValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                onReceive(StructNestedValue)
                return true
            }
            test.fbe.StructBytesFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                StructBytesModel.attach(buffer, offset)
                assert(StructBytesModel.verify()) { "test.StructBytes validation failed!" }
                val deserialized = StructBytesModel.deserialize(StructBytesValue)
                assert(deserialized > 0) { "test.StructBytes deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = StructBytesValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                onReceive(StructBytesValue)
                return true
            }
            test.fbe.StructArrayFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                StructArrayModel.attach(buffer, offset)
                assert(StructArrayModel.verify()) { "test.StructArray validation failed!" }
                val deserialized = StructArrayModel.deserialize(StructArrayValue)
                assert(deserialized > 0) { "test.StructArray deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = StructArrayValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                onReceive(StructArrayValue)
                return true
            }
            test.fbe.StructVectorFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                StructVectorModel.attach(buffer, offset)
                assert(StructVectorModel.verify()) { "test.StructVector validation failed!" }
                val deserialized = StructVectorModel.deserialize(StructVectorValue)
                assert(deserialized > 0) { "test.StructVector deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = StructVectorValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                onReceive(StructVectorValue)
                return true
            }
            test.fbe.StructListFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                StructListModel.attach(buffer, offset)
                assert(StructListModel.verify()) { "test.StructList validation failed!" }
                val deserialized = StructListModel.deserialize(StructListValue)
                assert(deserialized > 0) { "test.StructList deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = StructListValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                onReceive(StructListValue)
                return true
            }
            test.fbe.StructSetFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                StructSetModel.attach(buffer, offset)
                assert(StructSetModel.verify()) { "test.StructSet validation failed!" }
                val deserialized = StructSetModel.deserialize(StructSetValue)
                assert(deserialized > 0) { "test.StructSet deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = StructSetValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                onReceive(StructSetValue)
                return true
            }
            test.fbe.StructMapFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                StructMapModel.attach(buffer, offset)
                assert(StructMapModel.verify()) { "test.StructMap validation failed!" }
                val deserialized = StructMapModel.deserialize(StructMapValue)
                assert(deserialized > 0) { "test.StructMap deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = StructMapValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                onReceive(StructMapValue)
                return true
            }
            test.fbe.StructHashFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                StructHashModel.attach(buffer, offset)
                assert(StructHashModel.verify()) { "test.StructHash validation failed!" }
                val deserialized = StructHashModel.deserialize(StructHashValue)
                assert(deserialized > 0) { "test.StructHash deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = StructHashValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                onReceive(StructHashValue)
                return true
            }
            test.fbe.StructHashExFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                StructHashExModel.attach(buffer, offset)
                assert(StructHashExModel.verify()) { "test.StructHashEx validation failed!" }
                val deserialized = StructHashExModel.deserialize(StructHashExValue)
                assert(deserialized > 0) { "test.StructHashEx deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = StructHashExValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                onReceive(StructHashExValue)
                return true
            }
            test.fbe.StructEmptyFinalModel.fbeTypeConst ->
            {
                // Deserialize the value from the FBE stream
                StructEmptyModel.attach(buffer, offset)
                assert(StructEmptyModel.verify()) { "test.StructEmpty validation failed!" }
                val deserialized = StructEmptyModel.deserialize(StructEmptyValue)
                assert(deserialized > 0) { "test.StructEmpty deserialization failed!" }

                // Log the value
                if (logging)
                {
                    val message = StructEmptyValue.toString()
                    onReceiveLog(message)
                }

                // Call receive handler with deserialized value
                onReceive(StructEmptyValue)
                return true
            }
        }

        if ((protoReceiver != null) && protoReceiver!!.onReceive(type, buffer, offset, size))
            return true

        return false
    }
}
