pragma solidity ^0.5.3;

// Based on https://ed25519.cr.yp.to/python/ed25519.py

contract Ed25519 {
    uint constant b = 256;
    uint constant q = 2 ** 255 - 19;
    uint constant d = 37095705934669439343138083508754565189542113879843219016388785533085940283555;
                      // = -(121665/121666)
    uint constant I = 19681161376707505956807079304988542015446066515923890162744021073123829784752;
    uint constant Bx = 15112221349535400772501151409588531511454012693041857206046113283949847762202;
    uint constant By = 46316835694926478169428394003475163141307993866256225615783033603165251855960;

    //function expmod(uint _b, uint _e, uint _m) internal pure returns (uint) {
    //    uint[] memory es = new uint[](512);
    //    uint numberOfEs = 0;

    //    while (_e > 0) {
    //        es[numberOfEs] = _e;
    //        _e /= 2;
    //        numberOfEs += 1;
    //    }

    //    uint stepResult = 1;
    //    while (numberOfEs > 0) {
    //        _e = es[numberOfEs - 1];
    //        uint t = mulmod(stepResult, stepResult, _m);
    //        if (_e & 1 == 1) { t = mulmod(t, _b, _m); }
    //        stepResult = t;
    //        numberOfEs -= 1;
    //    }

    //    return stepResult;
    //}

    //function expmod(uint _b, uint _e, uint _m) internal pure returns (uint) {
    //    uint result = 1;
    //    while (_e > 0) {
    //        if (_e & 1 == 1) { result = mulmod(result, _b, _m); }
    //        _e = _e >> 1;
    //        _b = mulmod(_b, _b, _m);
    //    }
    //    return result;
    //}

    function expmod(uint _b, uint _e, uint _m) internal returns (uint o) {
        // use bigModExp precompile
        assembly {
            let p := mload(0x40)
            mstore(p, 0x20)
            mstore(add(p, 0x20), 0x20)
            mstore(add(p, 0x40), 0x20)
            mstore(add(p, 0x60), _b)
            mstore(add(p, 0x80), _e)
            mstore(add(p, 0xa0), _m)
            if iszero(call(not(0), 0x05, 0, p, 0xc0, o, 0x20)) {
                revert(0, 0)
            }
        }
    }

    function inv(uint _x) internal returns (uint) {
        return expmod(_x, q - 2, q);
    }

    //function xrecover(uint _y) internal view returns (uint) {
    //    uint xx = mulmod((mulmod(_y, _y, q) - 1),
    //                     inv(mulmod(d, mulmod(_y, _y, q), q) + 1),
    //                     q);
    //    uint x = expmod(xx, (q + 3) / 8, q);
    //    if (addmod(mulmod(x, x, q), q - xx, q) != 0) { x = mulmod(x, I, q); }
    //    if (x % 2 != 0) { x = q - x; }
    //    return x;
    //}

    function edwards(uint _Px, uint _Py,
                     uint _Qx, uint _Qy) internal returns (uint, uint) {
        uint x = mulmod(addmod(mulmod(_Px, _Qy, q), mulmod(_Qx, _Py, q), q),
                        inv(addmod(1, mulmod(d, mulmod(
                            mulmod(_Px, _Qx, q), mulmod(_Py, _Qy, q), q), q), q)), q);
        uint y = mulmod(addmod(mulmod(_Py, _Qy, q), mulmod(_Px, _Qx, q), q),
                        inv(addmod(1, q - mulmod(d, mulmod(
                            mulmod(_Px, _Qx, q), mulmod(_Py, _Qy, q), q), q), q)), q);
        return (x, y);
    }

    //function scalarmult(uint _Px, uint _Py,
    //                    uint _e) internal pure returns (uint, uint) {
    //    uint[] memory es = new uint[](1024);
    //    uint numberOfEs = 0;

    //    while (_e > 0) {
    //        es[numberOfEs] = _e;
    //        _e /= 2;
    //        numberOfEs += 1;
    //    }

    //    uint stepResultX = 0;
    //    uint stepResultY = 1;
    //    while (numberOfEs > 0) {
    //        _e = es[numberOfEs - 1];
    //        (uint Qx, uint Qy) = edwards(stepResultX, stepResultY,
    //                                     stepResultX, stepResultY);
    //        if (_e & 1 == 1) { (Qx, Qy) = edwards(Qx, Qy, _Px, _Py); }
    //        stepResultX = Qx;
    //        stepResultY = Qy;
    //        numberOfEs -= 1;
    //    }

    //    return (stepResultX, stepResultY);
    //}

    function scalarmult(uint _e) internal returns (uint, uint) {
        uint Px = Bx;
        uint Py = By;
        uint x = 0;
        uint y = 1;
        while (_e > 0) {
            if (_e & 1 == 1) { (x, y) = edwards(x, y, Px, Py); }
            _e = _e >> 1;
            (Px, Py) = edwards(Px, Py, Px, Py);
        }
        return (x, y);
    }

    function debug(uint _s) public returns (uint) {
        (, uint y) = scalarmult(_s);
        return y;
        //return mulmod(_s, inv(_s), q);
    }

    function gasEstimation() public returns (uint) {
        (, uint y) = scalarmult(2 ** 220);
        return y;
        //return expmod(2, (q - 1) / 4, q);
    }
}
