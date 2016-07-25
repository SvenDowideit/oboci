package oboci

import (
	"crypto/sha256"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cyphar/oboci/archive"
	spec "github.com/opencontainers/image-spec/specs-go/v1"
)

func Build(ctx string) (*spec.Image, error) {
	var img spec.Image

	// 0. Set up a temporary directory for build context.
	dir, err := ioutil.TempDir("", "oboci-build")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	// XXX: We should be able to do something nicer here.
	rootfs := filepath.Join(dir, "rootfs")

	// Copy the context before we do anything else.
	if err := archive.CopyWithTar(ctx, dir); err != nil {
		return err
	}

	// TODO: Make the format of the build script modular. Instruction is a good start.
	// 1. Open the Dockerfile, parse it.
	//     -- Deal with FROM? [skopeo]

	steps, err := dockerfileParse(dir)
	if err != nil {
		return err
	}

	// 2. For each step:
	for idx, step := range steps.Instructions {
		// We only generate all of the things necessary to mutate the data if
		// it's a data mutation step. This is purely to save cycles.
		if step.Type() == "data" {
			var blob bytes.Buffer

			// 3. Generate the manifest of the rootfs.
			mani, err := ComputeManifest(rootfs)
			if err != nil {
				return err
			}

			// 4. Run a container [if applicable].
			if err := step.Apply(&image, rootfs); err != nil {
				return err
			}

			// 5. Get a diff for the rootfs, generate the delta blob.
			delta, err := ComputeDelta(mani, rootfs)
			if err != nil {
				return err
			}

			hash := sha256.New()
			if err := delta.Blob(io.MultiWriter(&blob, hash)); err != nil {
				return err
			}
			digest := hash.Sum(nil)

			// 6. Add the delta blob to the image.

		}

		// 7. Generate the new RunConfig.

		// 8. Save it all and update references.
	}

	// 8. Package the final image as a single blob.
	return nil
}
