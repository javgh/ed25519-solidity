package main

import (
    "crypto/rand"
    "encoding/hex"
    "flag"
    "fmt"

    "golang.org/x/crypto/curve25519"
)

func main() {
    var n = flag.Int("n", 10, "number of test cases to generate")
    flag.Parse()

    var in, dst [32]byte
    for i := 0; i < *n; i++ {
        rand.Read(in[:])
        curve25519.ScalarBaseMult(&dst, &in)
        fmt.Printf("%s:%s\n", hex.EncodeToString(in[:]), hex.EncodeToString(dst[:]))
    }
}
