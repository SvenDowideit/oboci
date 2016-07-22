## oboci ##

Oboci builds open containers images (it's an acronym). The idea is to replace
`docker build` with a fully OCI-compliant stack (using [ocitools][ocitools] for
the [runtime spec][runtime] generation, [runC][runc] for the executor,
[go-mtree][mtree] for the diff generation and some other tools to generate the
final [OCI image][image]).

[ocitools]: https://github.com/opencontainers/ocitools
[runtime]: https://github.com/opencontainers/runtime-spec
[runc]: https://github.com/opencontainers/runc
[mtree]: https://github.com/vbatts/go-mtree
[image]: https://github.com/opencontainers/image-spec
