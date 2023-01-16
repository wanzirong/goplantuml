package parser

import (
	"strings"

	"golang.org/x/tools/go/packages"
)

var standardPackages = make(map[string]struct{})

func init() {
	pkgs, err := packages.Load(nil, "std")
	if err != nil {
		panic(err)
	}

	for _, p := range pkgs {
		standardPackages[p.PkgPath] = struct{}{}
	}
}

func isStandardPackage(fType string) bool {
	packageName := fType
	res := strings.SplitN(fType, ".", 2)
	if len(res) > 1 {
		packageName = res[0]
	}
	_, ok := standardPackages[packageName]
	return ok
}
