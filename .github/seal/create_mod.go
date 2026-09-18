// Command create_mod produces the module source zip that proxy.golang.org
// serves for this module, using golang.org/x/mod/zip.CreateFromDir — Go's own
// module-zip implementation, which applies the proxy's exact inclusion rules
// (nested modules, vendor/ and .git/ are excluded automatically).
//
// It lives under .github/ for two reasons: the go tool ignores directories
// whose names begin with a dot, so `go build ./...` and `go test ./...` never
// see it; and the zip build stages the tree with .github excluded, so this
// helper is never packaged into the artifact it produces.
//
// Usage: create_mod <module-path> <version> <source-dir> <output-zip>
package main

import (
	"log"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatal("usage: create_mod <module-path> <version> <source-dir> <output-zip>")
	}
	m := module.Version{Path: os.Args[1], Version: os.Args[2]}
	f, err := os.Create(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := zip.CreateFromDir(f, m, os.Args[3]); err != nil {
		log.Fatal(err)
	}
	log.Printf("created module zip: %s", os.Args[4])
}
