// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.3.0.0


import com.chronoxor.protoex.fbe

// Fast Binary Encoding com.chronoxor.protoex proxy
@Suppress("MemberVisibilityCanBePrivate", "PrivatePropertyName", "UNUSED_PARAMETER")
open class Proxy : com.chronoxor.fbe.Receiver, ProxyListener
{
    // Imported proxy
    var protoProxy: com.chronoxor.proto.fbe.Proxy? = null

    // Proxy models accessors
    private val OrderModel: OrderModel
    private val BalanceModel: BalanceModel
    private val AccountModel: AccountModel

    constructor() : super(false)
    {
        protoProxy = com.chronoxor.proto.fbe.Proxy(buffer)
        OrderModel = OrderModel()
        BalanceModel = BalanceModel()
        AccountModel = AccountModel()
    }

    constructor(buffer: com.chronoxor.fbe.Buffer) : super(buffer, false)
    {
        protoProxy = com.chronoxor.proto.fbe.Proxy(buffer)
        OrderModel = OrderModel()
        BalanceModel = BalanceModel()
        AccountModel = AccountModel()
    }

    override fun onReceive(type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {
        return onReceiveListener(this, type, buffer, offset, size)
    }

    open fun onReceiveListener(listener: ProxyListener, type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {
        when (type)
        {
            com.chronoxor.protoex.fbe.OrderModel.fbeTypeConst ->
            {
                // Attach the FBE stream to the proxy model
                OrderModel.attach((buffer: buffer, offset: offset)
                assertionFailure(OrderModel.verify()) { "com.chronoxor.protoex.Order validation failed!" }

                val fbeBegin = OrderModel.model.getBegin()
                if (fbeBegin == 0L)
                    return false
                // Call proxy handler
                listener.onProxy(OrderModel, type, buffer, offset, size)
                OrderModel.model.getEnd(fbeBegin)
                return true
            }
            com.chronoxor.protoex.fbe.BalanceModel.fbeTypeConst ->
            {
                // Attach the FBE stream to the proxy model
                BalanceModel.attach((buffer: buffer, offset: offset)
                assertionFailure(BalanceModel.verify()) { "com.chronoxor.protoex.Balance validation failed!" }

                val fbeBegin = BalanceModel.model.getBegin()
                if (fbeBegin == 0L)
                    return false
                // Call proxy handler
                listener.onProxy(BalanceModel, type, buffer, offset, size)
                BalanceModel.model.getEnd(fbeBegin)
                return true
            }
            com.chronoxor.protoex.fbe.AccountModel.fbeTypeConst ->
            {
                // Attach the FBE stream to the proxy model
                AccountModel.attach((buffer: buffer, offset: offset)
                assertionFailure(AccountModel.verify()) { "com.chronoxor.protoex.Account validation failed!" }

                val fbeBegin = AccountModel.model.getBegin()
                if (fbeBegin == 0L)
                    return false
                // Call proxy handler
                listener.onProxy(AccountModel, type, buffer, offset, size)
                AccountModel.model.getEnd(fbeBegin)
                return true
            }
        }

        if ((protoProxy != null) && protoProxy!!.onReceiveListener(listener, type, buffer, offset, size))
            return true

        return false
    }
}
