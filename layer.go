package oboci

import (
	"archive/tar"
	"io"

	mtree "github.com/vbatts/go-mtree"
)

// Manifest represents the state of a given rootfs at a particular point in
// time.
type Manifest struct {
	dh *mtree.DirectoryHierarchy
}

func (m *Manifest) WriteTo(w io.Writer) (int64, error) {
	return m.dh.WriteTo(w)
}

// ComputeManifest creates a new manifest for the given path. It should be
// used before applying any operation that will result in the rootfs layer
// being scanned changing.
func ComputeManifest(rootfs string) (*Manifest, error) {
	keywords := append(mtree.DefaultKeywords[:], []string{
		"sha256digest",
		"xattr",
	}...)

	dh, err := mtree.Walk(rootfs, nil, keywords)
	if err != nil {
		return nil, err
	}

	return &Manifest{dh: dh}, nil
}

// Delta represents the set of changes between two states of a rootfs.
type Delta struct {
	Includes  []string
	Whiteouts []string
}

// ComputeDelta computes the delta between a manifest and the current sate of a
// rootfs. This should be done after all operations that result in changes in
// the layer have completed.
func ComputeDelta(manifest *Manifest, rootfs string) (*Delta, error) {
	res, err := mtree.Check(rootfs, manifest.dh)
	if err != nil {
		return nil, err
	}

	var delta Delta
	for _, failure := range res.Failures {
		switch failure.Type() {
		case Present, Modified:
			delta.Includes = append(delta.Includes, failure.Path())
		case Missing:
			delta.Whiteouts = append(delta.Whiteouts, failure.Path())
		default:
			panic("Unexpected failure type!")
		}
	}

	return &delta, nil
}

// Blob generates a tar blob for the delta and writes it to the given writer.
// XXX: Currently we don't use tar-split. This is not good.
func (d *Delta) Blob(w io.Writer) error {
	tw := tar.NewWriter(w)

	for _, include := range d.Includes {
		fi, err := os.Lstat(include)
		if err != nil {
			return err
		}

		var link string
		if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
			link, err = os.Readlink(include)
			if err != nil {
				return err
			}
		}

		hdr := tar.FileInfoHeader(fi, link)
		hdr.Name = include
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		data, err := ioutil.ReadAll(include)
		if err != nil {
			return err
		}

		if n, err := tw.Write(data); err != nil {
			return err
		} else if n != len(data) {
			return fmt.Errorf("delta blob: whole object not writen: %s", include)
		}
	}

	for _, include := range d.Whiteouts {
		if err := tw.WriteHeader(&tar.Header{
			Name: filepath.Join(filepath.Dir(include), ".wh."+filepath.Base(include)),
		}); err != nil {
			return err
		}
	}

	tw.Flush()
	return nil
}
