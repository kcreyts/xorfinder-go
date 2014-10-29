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

func XorByReadByte(rdr io.ByteScanner, xorby byte) (byte, error) {
    b, err := rdr.ReadByte()
    return b ^ xorby, err
}

func XorByReader(rdr io.ByteScanner, xorby byte) ([]byte, error) {
    bs := []byte{}

    for b, err := XorByReadByte(rdr, xorby); err != io.EOF; b, err = XorByReadByte(rdr, xorby) {
        if err != nil {
            return nil, err
        }
        bs = append(bs, b)
    }
    return bs, nil
}

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
    base_test_string := "bfraaeui;nebabrieuangf;aonblahugoriagh;hreanoi;arebad adbnaabd acbv abc abab"
    aho := ahocorasick.NewAhoCorasick(sigs)
    aho2 := ahocorasick.NewAhoCorasick(sigs)
    fmt.Printf("%s", "base test string: " + base_test_string + "\n")
    input, _ := XorReader(strings.NewReader(base_test_string))
    for match := range ahocorasick.MatchBytes(input, aho) {
      fmt.Printf("Found match %q at index %d\n", match.Value, match.Index)
    }

    fmt.Println("\n\nand now the proof:\n")
    xord, _ := XorByReader(strings.NewReader(base_test_string), 3)
    fmt.Printf("%s", "xord test string: " + string(xord) + "\n")
    input2, _ := XorReader(strings.NewReader(string(xord)))
    for match := range ahocorasick.MatchBytes(input2, aho2) {
      fmt.Printf("Found match %q at index %d\n", match.Value, match.Index)
    }
}
