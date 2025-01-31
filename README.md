# Silent Map Access in Go

This example demonstrates a potential pitfall in Go: accessing a map with a non-existent key does not cause a panic; instead, it returns the zero value for the map's value type.  This can lead to unexpected behavior if not handled carefully.

The `bug.go` file contains code that illustrates this issue.  The `bugSolution.go` file provides a solution.

## Bug

The primary issue is the silent retrieval of a zero value instead of a clear indication of a missing key. This behavior can make debugging harder since no error is explicitly raised. The code will compile and run without a crash but may produce incorrect results.

## Solution

The recommended approach is to explicitly check for the key's existence before accessing its value, using the `ok` idiom.