package oboci

import (
	spec "github.com/opencontainers/image-spec/specs-go/v1"
)

type InstructionType string

const (
	Both InstructionType = "both"
	Data InstructionType = "data"
	Meta InstructionType = "meta"
)

type Instruction interface {
	// Type returns what type of instruction this is.
	Type() InstructionType

	// Apply mutates an image's data (rootfs) according the instruction to be
	// executed. This may involve running a container.
	Apply(image *spec.Image, rootfs string)

	// Mutate mutates an image's metadata (RunConfig) according to the
	// instruction being executed. Always run this *after* Apply().
	Mutate(image *spec.Image, rootfs string)
}
