package core

import ()

// GPUType is the structure to define what's inside a GPU <br>
//  GPU Structure
//  ================
//  ================
type GPUType struct {
	control  byte
	scrollX  byte
	scrollY  byte
	scanline byte
	tick     byte
}

// GPU is the exported object used in the system <br>
// GPU is exported to become a shared variable in the System object
var GPU = GPUType{
	control:  0x00,
	scrollX:  0x00,
	scrollY:  0x00,
	scanline: 0x00,
	tick:     0x00,
}
