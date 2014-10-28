xorfinder-go
============
The net goal of this project is to provide a very fast, very efficient facility for searching streams for xor'd versions of known signatures.

The current incarnation does provide moderately efficient facilities for searching streams for single-byte XOR of known signatures (strings).  


###TO DO###

- make it faster and more efficient
- investigate tradeoffs of algo implementation performance
- read sensibly rather than one byte at a time
- xor larger blocks of bytes at a time
- support longer keylengths by default
  - 2^n byte keys should be pretty easy
- support better mechanism for supplying input
- support better mechanism for supplying rules
