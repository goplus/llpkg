package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/libxml2"
	_ "unsafe"
)

type X_xsltSecurityPrefs struct {
	Unused [8]uint8
}
type SecurityPrefs X_xsltSecurityPrefs
type SecurityPrefsPtr *SecurityPrefs
type SecurityOption c.Int

const (
	SECPREFREADFILE        SecurityOption = 1
	SECPREFWRITEFILE       SecurityOption = 2
	SECPREFCREATEDIRECTORY SecurityOption = 3
	SECPREFREADNETWORK     SecurityOption = 4
	SECPREFWRITENETWORK    SecurityOption = 5
)

// llgo:type C
type SecurityCheck func(SecurityPrefsPtr, TransformContextPtr, *int8) c.Int

/*
 * Module interfaces
 */
//go:linkname NewSecurityPrefs C.xsltNewSecurityPrefs
func NewSecurityPrefs() SecurityPrefsPtr

//go:linkname FreeSecurityPrefs C.xsltFreeSecurityPrefs
func FreeSecurityPrefs(sec SecurityPrefsPtr)

//go:linkname SetSecurityPrefs C.xsltSetSecurityPrefs
func SetSecurityPrefs(sec SecurityPrefsPtr, option SecurityOption, func_ SecurityCheck) c.Int

//go:linkname GetSecurityPrefs C.xsltGetSecurityPrefs
func GetSecurityPrefs(sec SecurityPrefsPtr, option SecurityOption) SecurityCheck

//go:linkname SetDefaultSecurityPrefs C.xsltSetDefaultSecurityPrefs
func SetDefaultSecurityPrefs(sec SecurityPrefsPtr)

//go:linkname GetDefaultSecurityPrefs C.xsltGetDefaultSecurityPrefs
func GetDefaultSecurityPrefs() SecurityPrefsPtr

//go:linkname SetCtxtSecurityPrefs C.xsltSetCtxtSecurityPrefs
func SetCtxtSecurityPrefs(sec SecurityPrefsPtr, ctxt TransformContextPtr) c.Int

//go:linkname SecurityAllow C.xsltSecurityAllow
func SecurityAllow(sec SecurityPrefsPtr, ctxt TransformContextPtr, value *int8) c.Int

//go:linkname SecurityForbid C.xsltSecurityForbid
func SecurityForbid(sec SecurityPrefsPtr, ctxt TransformContextPtr, value *int8) c.Int

/*
 * internal interfaces
 */
//go:linkname CheckWrite C.xsltCheckWrite
func CheckWrite(sec SecurityPrefsPtr, ctxt TransformContextPtr, URL *libxml2.Char) c.Int

//go:linkname CheckRead C.xsltCheckRead
func CheckRead(sec SecurityPrefsPtr, ctxt TransformContextPtr, URL *libxml2.Char) c.Int
