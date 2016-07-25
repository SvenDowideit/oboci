package oboci

// TODO: This should probably live in a builder/ package so we can support
//       multiple Dockerfile-like formats. Oh, and the wrapping of the
//       Dockerfile parser is ... iffy.

import (
	"io"
	"path/filepath"

	"github.com/docker/docker/builder/dockerfile/parser"
)

const dockerfileName = "Dockerfile"

type steps struct {
	instructions []Instruction
}

func dockerfileParse(path string) ([]Instructions, error) {
	r, err := os.Open(filepath.Join(path, dockerfileName))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	// Get the AST from Docker.
	ast, err := parser.Parse(r)
	if err != nil {
		return nil, err
	}

	// TODO: Wrap these with dispatches.
	return ast.Children, nil
}
