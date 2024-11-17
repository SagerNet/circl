// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

//go:build (amd64 && !purego) || (arm64 && !purego)
// +build amd64,!purego arm64,!purego

package p503

import (
	. "github.com/sagernet/circl/dh/sidh/internal/common"
)

// If choice = 0, leave x unchanged. If choice = 1, sets x to y.
// If choice is neither 0 nor 1 then behaviour is undefined.
// This function executes in constant time.
//
//go:noescape
func cmovP503(x, y *Fp, choice uint8)

// If choice = 0, leave x,y unchanged. If choice = 1, set x,y = y,x.
// If choice is neither 0 nor 1 then behaviour is undefined.
// This function executes in constant time.
//
//go:noescape
func cswapP503(x, y *Fp, choice uint8)

// Compute z = x + y (mod p).
//
//go:noescape
func addP503(z, x, y *Fp)

// Compute z = x - y (mod p).
//
//go:noescape
func subP503(z, x, y *Fp)

// Compute z = x + y, without reducing mod p.
//
//go:noescape
func adlP503(z, x, y *FpX2)

// Compute z = x - y, without reducing mod p.
//
//go:noescape
func sulP503(z, x, y *FpX2)

// Reduce a field element in [0, 2*p) to one in [0,p).
//
//go:noescape
func modP503(x *Fp)

// Computes z = x * y.
//
//go:noescape
func mulP503(z *FpX2, x, y *Fp)

// Computes the Montgomery reduction z = x R^{-1} (mod 2*p). On return value
// of x may be changed. z=x not allowed.
//
//go:noescape
func rdcP503(z *Fp, x *FpX2)
