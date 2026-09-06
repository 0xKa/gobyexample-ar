[العربية](README.md) | **English**

# Go by Example (Arabic)

An independent Arabic translation of [Go by Example](https://gobyexample.com), a hands-on introduction to Go using annotated example programs.

> Arabic translation of [Go by Example](https://github.com/mmcgrana/gobyexample). This translation is maintained as a separate website and follows the upstream examples and build toolchain.

Website URL: <https://0xKa.github.io/gobyexample-ar/>

## Project Status

- Translation and review of all 85 examples are complete; see the [Examples Progress Table](docs/PROGRESS.md).
- The Arabic interface supports RTL (Right-to-Left) layout while keeping code blocks in LTR (Left-to-Right).
- CI verifies translation integrity and ensures generated output files match their source code.
- The website deploys automatically from `master` via GitHub Actions once GitHub Pages is enabled with **GitHub Actions** selected as the build source.

## How It Works

Source code and commentary reside in `examples/`. The build tool extracts documentation comments from `.go` and `.sh` files, merges them with HTML templates in `templates/`, and renders a static website into `public/`.

| Path | Purpose |
| --- | --- |
| `examples/` | Code and commentary representing the primary source content |
| `examples.txt` | Ordering of examples and original English names used to derive permalinks |
| `examples.ar.txt` | Arabic titles corresponding to constant example IDs |
| `templates/` | HTML templates, CSS styles, JavaScript, and shared assets |
| `tools/` | Scripts for testing, building, generating, and running a local development server |
| `public/` | The generated static website; **never edit manually** |

## Building & Previewing Locally

Requires the Go version specified in [`go.mod`](go.mod) and a Bash environment. On Windows, use Git Bash or WSL.

To build the site and refresh `public/`:

```console
$ tools/build
```

To verify that generated files match sources bit-for-bit:

```console
$ VERBOSE=1 TESTING=1 tools/build
```

To preview the website locally:

```console
$ tools/serve
```

Then visit `http://127.0.0.1:8000/`.

## Contributing

We welcome typo fixes, terminology standardizations, technical accuracy reviews, and synchronization of changes from the original repository.

Before you start, please read:

- [Contribution Guide](docs/CONTRIBUTING.md)
- [Glossary of Terms](docs/GLOSSARY.md)
- [Examples Progress Table](docs/PROGRESS.md)
- [Upstream Sync Guide](docs/sync.md)

Always modify source files first, run the build tool, and include the generated changes in `public/`. Pull requests modifying only `public/` will not be accepted.

## Upstream Synchronization

This repository is an independent translation and not a replacement for the original project. Bug fixes and new examples are pulled from [`mmcgrana/gobyexample`](https://github.com/mmcgrana/gobyexample), and new commentary is translated before merging into the publishing branch.

New examples and structural changes should be proposed to the original project first, as this translation aims to remain compatible with upstream.

## Attribution & License

Original content by [Mark McGranaghan](https://markmcgranaghan.com) and [Eli Bendersky](https://eli.thegreenplace.net). Source code is available in the [original repository](https://github.com/mmcgrana/gobyexample).

Original work is copyright Mark McGranaghan and released under a [Creative Commons Attribution 3.0 Unported License](https://creativecommons.org/licenses/by/3.0/). This translation preserves original attribution, source links, and license terms.

The Go Gopher logo was designed by [Renée French](https://reneefrench.blogspot.com/) and is available under the same license.

Arabic translation maintained by [0xKa](https://github.com/0xKa).

## Acknowledgments

The original project was inspired by [Docco](http://jashkenas.github.io/docco/), created by [Jeremy Ashkenas](https://github.com/jashkenas).
