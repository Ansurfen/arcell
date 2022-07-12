package build

const (
	// flag type
	BOOL = iota
	STRING
	INT
	// script fit in os
	BATCH = "bat"
	SHELL = "sh"
)

type flagType = int