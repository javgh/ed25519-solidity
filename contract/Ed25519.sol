pragma solidity ^0.5.3;

// Based on https://ed25519.cr.yp.to/python/ed25519.py

contract Ed25519 {
    uint constant q = 2 ** 255 - 19;
    uint constant d = 37095705934669439343138083508754565189542113879843219016388785533085940283555;
                      // = -(121665/121666)
    uint constant Bx = 15112221349535400772501151409588531511454012693041857206046113283949847762202;
    uint constant By = 46316835694926478169428394003475163141307993866256225615783033603165251855960;

    function inv(uint _x) internal returns (uint o) {
        // use bigModExp precompile
        assembly {
            let p := mload(0x40)
            mstore(p, 0x20)
            mstore(add(p, 0x20), 0x20)
            mstore(add(p, 0x40), 0x20)
            mstore(add(p, 0x60), _x)
            mstore(add(p, 0x80), 0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffeb) // q - 2
            mstore(add(p, 0xa0), 0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed) // q
            if iszero(call(not(0), 0x05, 0, p, 0xc0, o, 0x20)) {
                revert(0, 0)
            }
        }
    }

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
        (, uint y) = scalarmult(2 ** 80 - 1);
        return y;
        //return expmod(2, (q - 1) / 4, q);
    }
}
