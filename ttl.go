/*
Copyright (c) 2009 The Go Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

As this is fork of the official Go code the same license applies.
Extensions of the original work are copyright (c) 2011 Miek Gieben

This function was copied from https://github.com/miekg/dns/blob/master/scan.go
The original license applies!

*/
package main

func stringToTTL(token string) (uint32, bool) {
	var s, i uint32
	for _, c := range token {
		switch c {
		case 's', 'S':
			s += i
			i = 0
		case 'm', 'M':
			s += i * 60
			i = 0
		case 'h', 'H':
			s += i * 60 * 60
			i = 0
		case 'd', 'D':
			s += i * 60 * 60 * 24
			i = 0
		case 'w', 'W':
			s += i * 60 * 60 * 24 * 7
			i = 0
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			i *= 10
			i += uint32(c) - '0'
		default:
			return 0, false
		}
	}
	return s + i, true
}
