// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.3.0.0

package com.chronoxor.test.fbe;

import java.io.*;
import java.lang.*;
import java.lang.reflect.*;
import java.math.*;
import java.nio.ByteBuffer;
import java.nio.charset.*;
import java.time.*;
import java.util.*;

import com.chronoxor.fbe.*;
import com.chronoxor.test.*;

// Fast Binary Encoding com.chronoxor.test sender
public class Sender extends com.chronoxor.fbe.Sender
{
    // Imported senders
    public final com.chronoxor.proto.fbe.Sender protoSender;

    // Sender models accessors
    public final StructSimpleModel StructSimpleModel;
    public final StructOptionalModel StructOptionalModel;
    public final StructNestedModel StructNestedModel;
    public final StructBytesModel StructBytesModel;
    public final StructArrayModel StructArrayModel;
    public final StructVectorModel StructVectorModel;
    public final StructListModel StructListModel;
    public final StructSetModel StructSetModel;
    public final StructMapModel StructMapModel;
    public final StructHashModel StructHashModel;
    public final StructHashExModel StructHashExModel;
    public final StructEmptyModel StructEmptyModel;

    public Sender()
    {
        super(false);
        protoSender = new com.chronoxor.proto.fbe.Sender(getBuffer());
        StructSimpleModel = new StructSimpleModel(getBuffer());
        StructOptionalModel = new StructOptionalModel(getBuffer());
        StructNestedModel = new StructNestedModel(getBuffer());
        StructBytesModel = new StructBytesModel(getBuffer());
        StructArrayModel = new StructArrayModel(getBuffer());
        StructVectorModel = new StructVectorModel(getBuffer());
        StructListModel = new StructListModel(getBuffer());
        StructSetModel = new StructSetModel(getBuffer());
        StructMapModel = new StructMapModel(getBuffer());
        StructHashModel = new StructHashModel(getBuffer());
        StructHashExModel = new StructHashExModel(getBuffer());
        StructEmptyModel = new StructEmptyModel(getBuffer());
    }
    public Sender(Buffer buffer)
    {
        super(buffer, false);
        protoSender = new com.chronoxor.proto.fbe.Sender(getBuffer());
        StructSimpleModel = new StructSimpleModel(getBuffer());
        StructOptionalModel = new StructOptionalModel(getBuffer());
        StructNestedModel = new StructNestedModel(getBuffer());
        StructBytesModel = new StructBytesModel(getBuffer());
        StructArrayModel = new StructArrayModel(getBuffer());
        StructVectorModel = new StructVectorModel(getBuffer());
        StructListModel = new StructListModel(getBuffer());
        StructSetModel = new StructSetModel(getBuffer());
        StructMapModel = new StructMapModel(getBuffer());
        StructHashModel = new StructHashModel(getBuffer());
        StructHashExModel = new StructHashExModel(getBuffer());
        StructEmptyModel = new StructEmptyModel(getBuffer());
    }

    public long send(com.chronoxor.test.StructSimple value)
    {
        // Serialize the value into the FBE stream
        long serialized = StructSimpleModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.test.StructSimple serialization failed!";
        assert StructSimpleModel.verify() : "com.chronoxor.test.StructSimple validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }
    public long send(com.chronoxor.test.StructOptional value)
    {
        // Serialize the value into the FBE stream
        long serialized = StructOptionalModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.test.StructOptional serialization failed!";
        assert StructOptionalModel.verify() : "com.chronoxor.test.StructOptional validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }
    public long send(com.chronoxor.test.StructNested value)
    {
        // Serialize the value into the FBE stream
        long serialized = StructNestedModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.test.StructNested serialization failed!";
        assert StructNestedModel.verify() : "com.chronoxor.test.StructNested validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }
    public long send(com.chronoxor.test.StructBytes value)
    {
        // Serialize the value into the FBE stream
        long serialized = StructBytesModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.test.StructBytes serialization failed!";
        assert StructBytesModel.verify() : "com.chronoxor.test.StructBytes validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }
    public long send(com.chronoxor.test.StructArray value)
    {
        // Serialize the value into the FBE stream
        long serialized = StructArrayModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.test.StructArray serialization failed!";
        assert StructArrayModel.verify() : "com.chronoxor.test.StructArray validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }
    public long send(com.chronoxor.test.StructVector value)
    {
        // Serialize the value into the FBE stream
        long serialized = StructVectorModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.test.StructVector serialization failed!";
        assert StructVectorModel.verify() : "com.chronoxor.test.StructVector validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }
    public long send(com.chronoxor.test.StructList value)
    {
        // Serialize the value into the FBE stream
        long serialized = StructListModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.test.StructList serialization failed!";
        assert StructListModel.verify() : "com.chronoxor.test.StructList validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }
    public long send(com.chronoxor.test.StructSet value)
    {
        // Serialize the value into the FBE stream
        long serialized = StructSetModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.test.StructSet serialization failed!";
        assert StructSetModel.verify() : "com.chronoxor.test.StructSet validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }
    public long send(com.chronoxor.test.StructMap value)
    {
        // Serialize the value into the FBE stream
        long serialized = StructMapModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.test.StructMap serialization failed!";
        assert StructMapModel.verify() : "com.chronoxor.test.StructMap validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }
    public long send(com.chronoxor.test.StructHash value)
    {
        // Serialize the value into the FBE stream
        long serialized = StructHashModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.test.StructHash serialization failed!";
        assert StructHashModel.verify() : "com.chronoxor.test.StructHash validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }
    public long send(com.chronoxor.test.StructHashEx value)
    {
        // Serialize the value into the FBE stream
        long serialized = StructHashExModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.test.StructHashEx serialization failed!";
        assert StructHashExModel.verify() : "com.chronoxor.test.StructHashEx validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }
    public long send(com.chronoxor.test.StructEmpty value)
    {
        // Serialize the value into the FBE stream
        long serialized = StructEmptyModel.serialize(value);
        assert (serialized > 0) : "com.chronoxor.test.StructEmpty serialization failed!";
        assert StructEmptyModel.verify() : "com.chronoxor.test.StructEmpty validation failed!";

        // Log the value
        if (getLogging())
        {
            String message = value.toString();
            onSendLog(message);
        }

        // Send the serialized value
        return sendSerialized(serialized);
    }

    // Send message handler
    @Override
    protected long onSend(byte[] buffer, long offset, long size) { throw new UnsupportedOperationException("com.chronoxor.test.fbe.Sender.onSend() not implemented!"); }
}