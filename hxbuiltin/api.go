// Copyright 2013 The gohaxelib Authors listed in the AUTHORS file. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package hxbuiltin // This package contains functions over-loaded by the Go to Haxe transpiler

// Insert the given Haxe code at this point.
// BEWARE! It is very easy to write code that will break the system.
// code string must be a well-formed Haxe statement, probably terminated with a ";"
// ret is a Haxe Dynamic value mapped into Go as a uintptr
func HAXE(code string) (ret uintptr) { return }
