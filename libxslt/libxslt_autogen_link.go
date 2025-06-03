package libxslt

import (
	_ "github.com/goplus/lib/c"
	_ "github.com/goplus/lib/c/os"
	_ "github.com/goplus/llpkg/libxml2"
)

const LLGoPackage string = "link: $(pkg-config --libs libxslt);"
