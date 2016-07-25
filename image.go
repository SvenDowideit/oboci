package oboci

import (
	"io/ioutil"

	"github.com/containers/image/transports"
	"github.com/containers/image/types"
)

type image struct {
	path string
	ref  types.ImageReference
}

// NewImage creates a new temporary directory for use as the new image.
func NewImage(name string) (*image, error) {
	path, err := ioutil.TempDir("", name)
	if err != nil {
		return nil, err
	}

	ref, err := transports.ParseReference("dir:" + path)
	if err != nil {
		return nil, err
	}

	return &image{
		path: path,
		ref:  ref,
	}, nil
}
