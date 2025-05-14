package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/libxml2"
	"unsafe"
)

//go:linkname NewLocale C.xsltNewLocale
func NewLocale(langName *libxml2.Char, lowerFirst c.Int) unsafe.Pointer

//go:linkname FreeLocale C.xsltFreeLocale
func FreeLocale(locale unsafe.Pointer)

//go:linkname Strxfrm C.xsltStrxfrm
func Strxfrm(locale unsafe.Pointer, string *libxml2.Char) *libxml2.Char

//go:linkname FreeLocales C.xsltFreeLocales
func FreeLocales()

type Locale unsafe.Pointer
type LocaleChar libxml2.Char

//go:linkname LocaleStrcmp C.xsltLocaleStrcmp
func LocaleStrcmp(locale unsafe.Pointer, str1 *libxml2.Char, str2 *libxml2.Char) c.Int
