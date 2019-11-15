package engine

import ()

/*
	System Structure
	================
	CPU Structure
	---> Instructions Array
	---> Registers Structure
	---> MMU Structure
	GPU Structure
	---> Graphics Processing
	---> Screen Pipeline
*/
type SystemType struct {
	CPU CPUType
}

/*
	System Structure
	================
	CPU Structure
	---> Instructions Array
	---> Registers Structure
	---> MMU Structure
	GPU Structure
	---> Graphics Processing
	---> Screen Pipeline
*/
var System = SystemType {
	CPU: CPU,
}
