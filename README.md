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

### License ###

`oboci` is licensed under the ([GPLv3 or later][GPL-3.0]). This may change in
the future.

```
oboci: Builds Open Container Images
Copyright (C) 2016 SUSE Linux GmbH

oboci is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

oboci is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with oboci.  If not, see <http://www.gnu.org/licenses/>.
```

[GPL-3.0]: https://www.gnu.org/licenses/gpl-3.0.en.html
