package main

import "code.google.com/p/ahocorasick"
import "io"
import "fmt"
import "strings"

// net goal is to provide:
// a very fast, very efficient, facility for searching
// streams for xor'd versions of known signatures.
//
// Today, it does provide moderately fast and efficient
// facilities for searching streams for single-byte XOR
// of known signatures.
// 
// TODO
//
// make it fast and efficient
//  -(investigate tradeoffs of algo implementation performance) 
//  -read sensibly rather than one byte at a time
//  -xor larger blocks of bytes at a time
// support longer keylengths by default
//  - 2^n byte keys should be pretty easy
// support better mechanism for supplying input
// support better mechanism for supplying rules

func XorReadByte(rdr io.ByteScanner) (byte, error) {
    c, err := rdr.ReadByte()
    d, err := rdr.ReadByte()
    rdr.UnreadByte()
    return c ^ d, err
}

func XorReader(rdr io.ByteScanner) ([]byte, error) {
    cs := []byte{}

    for c, err := XorReadByte(rdr); err != io.EOF; c, err = XorReadByte(rdr) {
        if err != nil {
            return nil, err
        }
        cs = append(cs, c)
    }
    return cs, nil
}

func PrepSigs(inp []string) ([]string) {
    outp := []string{}
    for i := 0; i<len(inp); i++  {
        s,_ := XorReader(strings.NewReader(inp[i]))
        outp = append(outp,string(s))
    }
    return outp
}

func main() {
    sigs := []string{"abc", "abd", "abab", "blah"}
    sigs = PrepSigs(sigs)
    aho := ahocorasick.NewAhoCorasick(sigs)
    input, _ := XorReader(strings.NewReader("bfraaeui;nebabrieuangf;aonblahugoriagh;hreanoi;arebad adbnaabd acbv abc abab"))
    for match := range ahocorasick.MatchBytes(input, aho) {
      fmt.Printf("Found match %q at index %d\n", match.Value, match.Index)
    }
}
