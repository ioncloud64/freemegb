package engine

import (

)

type GPUType struct {
	control		byte
	scrollX		byte
	scrollY		byte
	scanline	byte
	tick		byte
}

var GPU = GPUType {
	control:	0x00,
	scrollX:	0x00,
	scrollY:	0x00,
	scanline:	0x00,
	tick:		0x00,
}