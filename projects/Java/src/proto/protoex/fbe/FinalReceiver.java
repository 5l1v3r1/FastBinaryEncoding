// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.3.0.0

package protoex.fbe;

import java.io.*;
import java.lang.*;
import java.lang.reflect.*;
import java.math.*;
import java.nio.ByteBuffer;
import java.nio.charset.*;
import java.time.*;
import java.util.*;

import fbe.*;
import protoex.*;

// Fast Binary Encoding protoex final receiver
public class FinalReceiver extends fbe.Receiver
{
    // Imported receivers
    public proto.fbe.FinalReceiver protoReceiver;

    // Receiver values accessors
    private final protoex.Order OrderValue;
    private final protoex.Balance BalanceValue;
    private final protoex.Account AccountValue;

    // Receiver models accessors
    private final OrderFinalModel OrderModel;
    private final BalanceFinalModel BalanceModel;
    private final AccountFinalModel AccountModel;

    public FinalReceiver()
    {
        super(true);
        protoReceiver = new proto.fbe.FinalReceiver(getBuffer());
        OrderValue = new protoex.Order();
        OrderModel = new OrderFinalModel();
        BalanceValue = new protoex.Balance();
        BalanceModel = new BalanceFinalModel();
        AccountValue = new protoex.Account();
        AccountModel = new AccountFinalModel();
    }
    public FinalReceiver(Buffer buffer)
    {
        super(buffer, true);
        protoReceiver = new proto.fbe.FinalReceiver(getBuffer());
        OrderValue = new protoex.Order();
        OrderModel = new OrderFinalModel();
        BalanceValue = new protoex.Balance();
        BalanceModel = new BalanceFinalModel();
        AccountValue = new protoex.Account();
        AccountModel = new AccountFinalModel();
    }

    // Receive handlers
    protected void onReceive(protoex.Order value) {}
    protected void onReceive(protoex.Balance value) {}
    protected void onReceive(protoex.Account value) {}

    @Override
    public boolean onReceive(long type, byte[] buffer, long offset, long size)
    {
        switch ((int)type)
        {
            case (int)protoex.fbe.OrderFinalModel.fbeTypeConst:
            {
                // Deserialize the value from the FBE stream
                OrderModel.attach(buffer, offset);
                assert OrderModel.verify() : "protoex.Order validation failed!";
                long deserialized = OrderModel.deserialize(OrderValue);
                assert (deserialized > 0) : "protoex.Order deserialization failed!";

                // Log the value
                if (getLogging())
                {
                    String message = OrderValue.toString();
                    onReceiveLog(message);
                }

                // Call receive handler with deserialized value
                onReceive(OrderValue);
                return true;
            }
            case (int)protoex.fbe.BalanceFinalModel.fbeTypeConst:
            {
                // Deserialize the value from the FBE stream
                BalanceModel.attach(buffer, offset);
                assert BalanceModel.verify() : "protoex.Balance validation failed!";
                long deserialized = BalanceModel.deserialize(BalanceValue);
                assert (deserialized > 0) : "protoex.Balance deserialization failed!";

                // Log the value
                if (getLogging())
                {
                    String message = BalanceValue.toString();
                    onReceiveLog(message);
                }

                // Call receive handler with deserialized value
                onReceive(BalanceValue);
                return true;
            }
            case (int)protoex.fbe.AccountFinalModel.fbeTypeConst:
            {
                // Deserialize the value from the FBE stream
                AccountModel.attach(buffer, offset);
                assert AccountModel.verify() : "protoex.Account validation failed!";
                long deserialized = AccountModel.deserialize(AccountValue);
                assert (deserialized > 0) : "protoex.Account deserialization failed!";

                // Log the value
                if (getLogging())
                {
                    String message = AccountValue.toString();
                    onReceiveLog(message);
                }

                // Call receive handler with deserialized value
                onReceive(AccountValue);
                return true;
            }
        }

        if ((protoReceiver != null) && protoReceiver.onReceive(type, buffer, offset, size))
            return true;

        return false;
    }
}
