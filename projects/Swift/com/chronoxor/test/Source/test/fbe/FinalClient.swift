// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.3.0.0


import Foundation

import fbe

// Fast Binary Encoding test final client
open class FinalClient : fbe.ClientProtocol, FinalReceiverListener {
    // Imported senders
    let protoSender: proto.FinalClient

    // Imported receivers
    let protoReceiver: proto.FinalClient? = null

    // Client sender models accessors
    let StructSimpleSenderModel: StructSimpleFinalModel
    let StructOptionalSenderModel: StructOptionalFinalModel
    let StructNestedSenderModel: StructNestedFinalModel
    let StructBytesSenderModel: StructBytesFinalModel
    let StructArraySenderModel: StructArrayFinalModel
    let StructVectorSenderModel: StructVectorFinalModel
    let StructListSenderModel: StructListFinalModel
    let StructSetSenderModel: StructSetFinalModel
    let StructMapSenderModel: StructMapFinalModel
    let StructHashSenderModel: StructHashFinalModel
    let StructHashExSenderModel: StructHashExFinalModel
    let StructEmptySenderModel: StructEmptyFinalModel

    // Client receiver values accessors
    private var StructSimpleReceiverValue: test.StructSimple
    private var StructOptionalReceiverValue: test.StructOptional
    private var StructNestedReceiverValue: test.StructNested
    private var StructBytesReceiverValue: test.StructBytes
    private var StructArrayReceiverValue: test.StructArray
    private var StructVectorReceiverValue: test.StructVector
    private var StructListReceiverValue: test.StructList
    private var StructSetReceiverValue: test.StructSet
    private var StructMapReceiverValue: test.StructMap
    private var StructHashReceiverValue: test.StructHash
    private var StructHashExReceiverValue: test.StructHashEx
    private var StructEmptyReceiverValue: test.StructEmpty

    // Client receiver models accessors
    private let StructSimpleReceiverModel: StructSimpleFinalModel
    private let StructOptionalReceiverModel: StructOptionalFinalModel
    private let StructNestedReceiverModel: StructNestedFinalModel
    private let StructBytesReceiverModel: StructBytesFinalModel
    private let StructArrayReceiverModel: StructArrayFinalModel
    private let StructVectorReceiverModel: StructVectorFinalModel
    private let StructListReceiverModel: StructListFinalModel
    private let StructSetReceiverModel: StructSetFinalModel
    private let StructMapReceiverModel: StructMapFinalModel
    private let StructHashReceiverModel: StructHashFinalModel
    private let StructHashExReceiverModel: StructHashExFinalModel
    private let StructEmptyReceiverModel: StructEmptyFinalModel

    public var sendBuffer: Buffer = Buffer()
    public var receiveBuffer: Buffer = Buffer()
    public var logging: Bool = false
    public var final: Bool = false

    public init() {
        protoSender = proto.FinalClient(sendBuffer: sendBuffer, receiveBuffer: receiveBuffer)
        protoReceiver = proto.FinalClient(sendBuffer: sendBuffer, receiveBuffer: receiveBuffer)
        StructSimpleSenderModel = StructSimpleFinalModel(buffer: sendBuffer)
        StructSimpleReceiverValue = test.StructSimple()
        StructSimpleReceiverModel = StructSimpleFinalModel()
        StructOptionalSenderModel = StructOptionalFinalModel(buffer: sendBuffer)
        StructOptionalReceiverValue = test.StructOptional()
        StructOptionalReceiverModel = StructOptionalFinalModel()
        StructNestedSenderModel = StructNestedFinalModel(buffer: sendBuffer)
        StructNestedReceiverValue = test.StructNested()
        StructNestedReceiverModel = StructNestedFinalModel()
        StructBytesSenderModel = StructBytesFinalModel(buffer: sendBuffer)
        StructBytesReceiverValue = test.StructBytes()
        StructBytesReceiverModel = StructBytesFinalModel()
        StructArraySenderModel = StructArrayFinalModel(buffer: sendBuffer)
        StructArrayReceiverValue = test.StructArray()
        StructArrayReceiverModel = StructArrayFinalModel()
        StructVectorSenderModel = StructVectorFinalModel(buffer: sendBuffer)
        StructVectorReceiverValue = test.StructVector()
        StructVectorReceiverModel = StructVectorFinalModel()
        StructListSenderModel = StructListFinalModel(buffer: sendBuffer)
        StructListReceiverValue = test.StructList()
        StructListReceiverModel = StructListFinalModel()
        StructSetSenderModel = StructSetFinalModel(buffer: sendBuffer)
        StructSetReceiverValue = test.StructSet()
        StructSetReceiverModel = StructSetFinalModel()
        StructMapSenderModel = StructMapFinalModel(buffer: sendBuffer)
        StructMapReceiverValue = test.StructMap()
        StructMapReceiverModel = StructMapFinalModel()
        StructHashSenderModel = StructHashFinalModel(buffer: sendBuffer)
        StructHashReceiverValue = test.StructHash()
        StructHashReceiverModel = StructHashFinalModel()
        StructHashExSenderModel = StructHashExFinalModel(buffer: sendBuffer)
        StructHashExReceiverValue = test.StructHashEx()
        StructHashExReceiverModel = StructHashExFinalModel()
        StructEmptySenderModel = StructEmptyFinalModel(buffer: sendBuffer)
        StructEmptyReceiverValue = test.StructEmpty()
        StructEmptyReceiverModel = StructEmptyFinalModel()
        build(with: true)
    }

    public init(sendBuffer: fbe.Buffer, receiveBuffer: fbe.Buffer) {
        protoSender = proto.FinalClient(sendBuffer: sendBuffer, receiveBuffer: receiveBuffer)
        protoReceiver = proto.FinalClient(sendBuffer: sendBuffer, receiveBuffer: receiveBuffer)
        StructSimpleSenderModel = StructSimpleFinalModel(buffer: sendBuffer)
        StructSimpleReceiverValue = test.StructSimple()
        StructSimpleReceiverModel = StructSimpleFinalModel()
        StructOptionalSenderModel = StructOptionalFinalModel(buffer: sendBuffer)
        StructOptionalReceiverValue = test.StructOptional()
        StructOptionalReceiverModel = StructOptionalFinalModel()
        StructNestedSenderModel = StructNestedFinalModel(buffer: sendBuffer)
        StructNestedReceiverValue = test.StructNested()
        StructNestedReceiverModel = StructNestedFinalModel()
        StructBytesSenderModel = StructBytesFinalModel(buffer: sendBuffer)
        StructBytesReceiverValue = test.StructBytes()
        StructBytesReceiverModel = StructBytesFinalModel()
        StructArraySenderModel = StructArrayFinalModel(buffer: sendBuffer)
        StructArrayReceiverValue = test.StructArray()
        StructArrayReceiverModel = StructArrayFinalModel()
        StructVectorSenderModel = StructVectorFinalModel(buffer: sendBuffer)
        StructVectorReceiverValue = test.StructVector()
        StructVectorReceiverModel = StructVectorFinalModel()
        StructListSenderModel = StructListFinalModel(buffer: sendBuffer)
        StructListReceiverValue = test.StructList()
        StructListReceiverModel = StructListFinalModel()
        StructSetSenderModel = StructSetFinalModel(buffer: sendBuffer)
        StructSetReceiverValue = test.StructSet()
        StructSetReceiverModel = StructSetFinalModel()
        StructMapSenderModel = StructMapFinalModel(buffer: sendBuffer)
        StructMapReceiverValue = test.StructMap()
        StructMapReceiverModel = StructMapFinalModel()
        StructHashSenderModel = StructHashFinalModel(buffer: sendBuffer)
        StructHashReceiverValue = test.StructHash()
        StructHashReceiverModel = StructHashFinalModel()
        StructHashExSenderModel = StructHashExFinalModel(buffer: sendBuffer)
        StructHashExReceiverValue = test.StructHashEx()
        StructHashExReceiverModel = StructHashExFinalModel()
        StructEmptySenderModel = StructEmptyFinalModel(buffer: sendBuffer)
        StructEmptyReceiverValue = test.StructEmpty()
        StructEmptyReceiverModel = StructEmptyFinalModel()
        build(with: sendBuffer, receiveBuffer: receiveBuffer, final: true)
    }

    public func send(obj: Any) throws -> Int {
        switch obj {
            case is test.StructSimple: return try send(value: obj as! test.StructSimple)
            case is test.StructOptional: return try send(value: obj as! test.StructOptional)
            case is test.StructNested: return try send(value: obj as! test.StructNested)
            case is test.StructBytes: return try send(value: obj as! test.StructBytes)
            case is test.StructArray: return try send(value: obj as! test.StructArray)
            case is test.StructVector: return try send(value: obj as! test.StructVector)
            case is test.StructList: return try send(value: obj as! test.StructList)
            case is test.StructSet: return try send(value: obj as! test.StructSet)
            case is test.StructMap: return try send(value: obj as! test.StructMap)
            case is test.StructHash: return try send(value: obj as! test.StructHash)
            case is test.StructHashEx: return try send(value: obj as! test.StructHashEx)
            case is test.StructEmpty: return try send(value: obj as! test.StructEmpty)
            default: break
        }

        // Try to send using imported clients
        var result: Int = 0
        result = protoSender.send(obj: obj)
        if result > 0 { return result }

        return 0
    }

    public func send(value: test.StructSimple) throws -> Int {
        // Serialize the value into the FBE stream
        let serialized = try StructSimpleSenderModel.serialize(value: value)
        assert(serialized > 0, "test.StructSimple serialization failed!")
        assert(StructSimpleSenderModel.verify(), "test.StructSimple validation failed!")

        // Log the value
        if logging {
            let message = value.description
            onSendLog(message: message)
        }

        // Send the serialized value
        return try sendSerialized(serialized: serialized)
    }
    public func send(value: test.StructOptional) throws -> Int {
        // Serialize the value into the FBE stream
        let serialized = try StructOptionalSenderModel.serialize(value: value)
        assert(serialized > 0, "test.StructOptional serialization failed!")
        assert(StructOptionalSenderModel.verify(), "test.StructOptional validation failed!")

        // Log the value
        if logging {
            let message = value.description
            onSendLog(message: message)
        }

        // Send the serialized value
        return try sendSerialized(serialized: serialized)
    }
    public func send(value: test.StructNested) throws -> Int {
        // Serialize the value into the FBE stream
        let serialized = try StructNestedSenderModel.serialize(value: value)
        assert(serialized > 0, "test.StructNested serialization failed!")
        assert(StructNestedSenderModel.verify(), "test.StructNested validation failed!")

        // Log the value
        if logging {
            let message = value.description
            onSendLog(message: message)
        }

        // Send the serialized value
        return try sendSerialized(serialized: serialized)
    }
    public func send(value: test.StructBytes) throws -> Int {
        // Serialize the value into the FBE stream
        let serialized = try StructBytesSenderModel.serialize(value: value)
        assert(serialized > 0, "test.StructBytes serialization failed!")
        assert(StructBytesSenderModel.verify(), "test.StructBytes validation failed!")

        // Log the value
        if logging {
            let message = value.description
            onSendLog(message: message)
        }

        // Send the serialized value
        return try sendSerialized(serialized: serialized)
    }
    public func send(value: test.StructArray) throws -> Int {
        // Serialize the value into the FBE stream
        let serialized = try StructArraySenderModel.serialize(value: value)
        assert(serialized > 0, "test.StructArray serialization failed!")
        assert(StructArraySenderModel.verify(), "test.StructArray validation failed!")

        // Log the value
        if logging {
            let message = value.description
            onSendLog(message: message)
        }

        // Send the serialized value
        return try sendSerialized(serialized: serialized)
    }
    public func send(value: test.StructVector) throws -> Int {
        // Serialize the value into the FBE stream
        let serialized = try StructVectorSenderModel.serialize(value: value)
        assert(serialized > 0, "test.StructVector serialization failed!")
        assert(StructVectorSenderModel.verify(), "test.StructVector validation failed!")

        // Log the value
        if logging {
            let message = value.description
            onSendLog(message: message)
        }

        // Send the serialized value
        return try sendSerialized(serialized: serialized)
    }
    public func send(value: test.StructList) throws -> Int {
        // Serialize the value into the FBE stream
        let serialized = try StructListSenderModel.serialize(value: value)
        assert(serialized > 0, "test.StructList serialization failed!")
        assert(StructListSenderModel.verify(), "test.StructList validation failed!")

        // Log the value
        if logging {
            let message = value.description
            onSendLog(message: message)
        }

        // Send the serialized value
        return try sendSerialized(serialized: serialized)
    }
    public func send(value: test.StructSet) throws -> Int {
        // Serialize the value into the FBE stream
        let serialized = try StructSetSenderModel.serialize(value: value)
        assert(serialized > 0, "test.StructSet serialization failed!")
        assert(StructSetSenderModel.verify(), "test.StructSet validation failed!")

        // Log the value
        if logging {
            let message = value.description
            onSendLog(message: message)
        }

        // Send the serialized value
        return try sendSerialized(serialized: serialized)
    }
    public func send(value: test.StructMap) throws -> Int {
        // Serialize the value into the FBE stream
        let serialized = try StructMapSenderModel.serialize(value: value)
        assert(serialized > 0, "test.StructMap serialization failed!")
        assert(StructMapSenderModel.verify(), "test.StructMap validation failed!")

        // Log the value
        if logging {
            let message = value.description
            onSendLog(message: message)
        }

        // Send the serialized value
        return try sendSerialized(serialized: serialized)
    }
    public func send(value: test.StructHash) throws -> Int {
        // Serialize the value into the FBE stream
        let serialized = try StructHashSenderModel.serialize(value: value)
        assert(serialized > 0, "test.StructHash serialization failed!")
        assert(StructHashSenderModel.verify(), "test.StructHash validation failed!")

        // Log the value
        if logging {
            let message = value.description
            onSendLog(message: message)
        }

        // Send the serialized value
        return try sendSerialized(serialized: serialized)
    }
    public func send(value: test.StructHashEx) throws -> Int {
        // Serialize the value into the FBE stream
        let serialized = try StructHashExSenderModel.serialize(value: value)
        assert(serialized > 0, "test.StructHashEx serialization failed!")
        assert(StructHashExSenderModel.verify(), "test.StructHashEx validation failed!")

        // Log the value
        if logging {
            let message = value.description
            onSendLog(message: message)
        }

        // Send the serialized value
        return try sendSerialized(serialized: serialized)
    }
    public func send(value: test.StructEmpty) throws -> Int {
        // Serialize the value into the FBE stream
        let serialized = try StructEmptySenderModel.serialize(value: value)
        assert(serialized > 0, "test.StructEmpty serialization failed!")
        assert(StructEmptySenderModel.verify(), "test.StructEmpty validation failed!")

        // Log the value
        if logging {
            let message = value.description
            onSendLog(message: message)
        }

        // Send the serialized value
        return try sendSerialized(serialized: serialized)
    }

    // Send message handler
    open func onSend(buffer: Data, offset: Int, size: Int) throws -> Int { throw NSError() }
    open func onReceive(type: Int, buffer: Data, offset: Int, size: Int) -> Bool {
        return onReceiveListener(listener: self, type: type, buffer: buffer, offset: offset, size: size)
    }

    open func onReceiveListener(listener: FinalReceiverListener, type: Int, buffer: Data, offset: Int, size: Int) -> Bool {
        switch type {
        case test.StructSimpleFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            StructSimpleReceiverModel.attach(buffer: buffer, offset: offset)
            assert(StructSimpleReceiverModel.verify(), "test.StructSimple validation failed!")
            let deserialized = StructSimpleReceiverModel.deserialize(value: &StructSimpleReceiverValue)
            assert(deserialized > 0, "test.StructSimple deserialization failed!")

            // Log the value
            if logging {
                let message = StructSimpleReceiverValue.description
                onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: StructSimpleReceiverValue)
            return true
        case test.StructOptionalFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            StructOptionalReceiverModel.attach(buffer: buffer, offset: offset)
            assert(StructOptionalReceiverModel.verify(), "test.StructOptional validation failed!")
            let deserialized = StructOptionalReceiverModel.deserialize(value: &StructOptionalReceiverValue)
            assert(deserialized > 0, "test.StructOptional deserialization failed!")

            // Log the value
            if logging {
                let message = StructOptionalReceiverValue.description
                onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: StructOptionalReceiverValue)
            return true
        case test.StructNestedFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            StructNestedReceiverModel.attach(buffer: buffer, offset: offset)
            assert(StructNestedReceiverModel.verify(), "test.StructNested validation failed!")
            let deserialized = StructNestedReceiverModel.deserialize(value: &StructNestedReceiverValue)
            assert(deserialized > 0, "test.StructNested deserialization failed!")

            // Log the value
            if logging {
                let message = StructNestedReceiverValue.description
                onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: StructNestedReceiverValue)
            return true
        case test.StructBytesFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            StructBytesReceiverModel.attach(buffer: buffer, offset: offset)
            assert(StructBytesReceiverModel.verify(), "test.StructBytes validation failed!")
            let deserialized = StructBytesReceiverModel.deserialize(value: &StructBytesReceiverValue)
            assert(deserialized > 0, "test.StructBytes deserialization failed!")

            // Log the value
            if logging {
                let message = StructBytesReceiverValue.description
                onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: StructBytesReceiverValue)
            return true
        case test.StructArrayFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            StructArrayReceiverModel.attach(buffer: buffer, offset: offset)
            assert(StructArrayReceiverModel.verify(), "test.StructArray validation failed!")
            let deserialized = StructArrayReceiverModel.deserialize(value: &StructArrayReceiverValue)
            assert(deserialized > 0, "test.StructArray deserialization failed!")

            // Log the value
            if logging {
                let message = StructArrayReceiverValue.description
                onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: StructArrayReceiverValue)
            return true
        case test.StructVectorFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            StructVectorReceiverModel.attach(buffer: buffer, offset: offset)
            assert(StructVectorReceiverModel.verify(), "test.StructVector validation failed!")
            let deserialized = StructVectorReceiverModel.deserialize(value: &StructVectorReceiverValue)
            assert(deserialized > 0, "test.StructVector deserialization failed!")

            // Log the value
            if logging {
                let message = StructVectorReceiverValue.description
                onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: StructVectorReceiverValue)
            return true
        case test.StructListFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            StructListReceiverModel.attach(buffer: buffer, offset: offset)
            assert(StructListReceiverModel.verify(), "test.StructList validation failed!")
            let deserialized = StructListReceiverModel.deserialize(value: &StructListReceiverValue)
            assert(deserialized > 0, "test.StructList deserialization failed!")

            // Log the value
            if logging {
                let message = StructListReceiverValue.description
                onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: StructListReceiverValue)
            return true
        case test.StructSetFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            StructSetReceiverModel.attach(buffer: buffer, offset: offset)
            assert(StructSetReceiverModel.verify(), "test.StructSet validation failed!")
            let deserialized = StructSetReceiverModel.deserialize(value: &StructSetReceiverValue)
            assert(deserialized > 0, "test.StructSet deserialization failed!")

            // Log the value
            if logging {
                let message = StructSetReceiverValue.description
                onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: StructSetReceiverValue)
            return true
        case test.StructMapFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            StructMapReceiverModel.attach(buffer: buffer, offset: offset)
            assert(StructMapReceiverModel.verify(), "test.StructMap validation failed!")
            let deserialized = StructMapReceiverModel.deserialize(value: &StructMapReceiverValue)
            assert(deserialized > 0, "test.StructMap deserialization failed!")

            // Log the value
            if logging {
                let message = StructMapReceiverValue.description
                onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: StructMapReceiverValue)
            return true
        case test.StructHashFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            StructHashReceiverModel.attach(buffer: buffer, offset: offset)
            assert(StructHashReceiverModel.verify(), "test.StructHash validation failed!")
            let deserialized = StructHashReceiverModel.deserialize(value: &StructHashReceiverValue)
            assert(deserialized > 0, "test.StructHash deserialization failed!")

            // Log the value
            if logging {
                let message = StructHashReceiverValue.description
                onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: StructHashReceiverValue)
            return true
        case test.StructHashExFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            StructHashExReceiverModel.attach(buffer: buffer, offset: offset)
            assert(StructHashExReceiverModel.verify(), "test.StructHashEx validation failed!")
            let deserialized = StructHashExReceiverModel.deserialize(value: &StructHashExReceiverValue)
            assert(deserialized > 0, "test.StructHashEx deserialization failed!")

            // Log the value
            if logging {
                let message = StructHashExReceiverValue.description
                onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: StructHashExReceiverValue)
            return true
        case test.StructEmptyFinalModel.fbeTypeConst:
            // Deserialize the value from the FBE stream
            StructEmptyReceiverModel.attach(buffer: buffer, offset: offset)
            assert(StructEmptyReceiverModel.verify(), "test.StructEmpty validation failed!")
            let deserialized = StructEmptyReceiverModel.deserialize(value: &StructEmptyReceiverValue)
            assert(deserialized > 0, "test.StructEmpty deserialization failed!")

            // Log the value
            if logging {
                let message = StructEmptyReceiverValue.description
                onReceiveLog(message: message)
            }

            // Call receive handler with deserialized value
            listener.onReceive(value: StructEmptyReceiverValue)
            return true
        default: break
        }

        if let protoReceiver == protoReceiver, protoReceiver.onReceiveListener(listener, type, buffer, offset, size) {
            return true
        }

        return false
    }

    open func onReceive(value: test.StructSimple) { }
    open func onReceive(value: test.StructOptional) { }
    open func onReceive(value: test.StructNested) { }
    open func onReceive(value: test.StructBytes) { }
    open func onReceive(value: test.StructArray) { }
    open func onReceive(value: test.StructVector) { }
    open func onReceive(value: test.StructList) { }
    open func onReceive(value: test.StructSet) { }
    open func onReceive(value: test.StructMap) { }
    open func onReceive(value: test.StructHash) { }
    open func onReceive(value: test.StructHashEx) { }
    open func onReceive(value: test.StructEmpty) { }
}
