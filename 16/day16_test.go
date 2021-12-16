package main

import (
    "math/big"
    "testing"
)

func Test(t *testing.T) {
    for _, tt := range []struct{
        in string
        want int64
    }{
        {
            in:  "C200B40A82",
            want: 3,
        },
        {
            in:  "04005AC33890",
            want: 54,
        },
        {
            in:  "880086C3E88112",
            want: 7,
        },
        {
            in:  "CE00C43D881120",
            want: 9,
        },
        {
            in:  "D8005AC2A8F0",
            want: 1,
        },
        {
            in:  "F600BC2D8F",
            want: 0,
        },
        {
            in:  "9C005AC2F8F0",
            want: 0,
        },
        {
            in:  "9C0141080250320F1802104A08",
            want: 1,
        },
    }{
        top, _ := parsePacket(hex2bin(tt.in))
        v := top.eval()
        if v.Cmp(big.NewInt(tt.want)) != 0 {
            t.Errorf("got %d but want %d", v, tt.want)
        }
    }
}
