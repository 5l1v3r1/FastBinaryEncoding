// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.3.0.0


import fbe

import Foundation

// Fast Binary Encoding com.chronoxor.enums proxy listener
public protocol ProxyListener{
    func onProxy(model: EnumsModel, type: Int, buffer: Data, offset: Int, size: Int)
}

public extension ProxyListener{
    func onProxy(model: EnumsModel, type: Int, buffer: Data, offset: Int, size: Int) { }
}
