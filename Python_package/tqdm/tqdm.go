package tqdm

import (
	"github.com/goplus/lib/py"
	_ "unsafe"
)

const LLGoPackage = "py.tqdm"
//
//     Registers the given `tqdm` instance with
//     `pandas.core.groupby.DataFrameGroupBy.progress_apply`.
//
//
//go:linkname TqdmPandas py.tqdm_pandas
func TqdmPandas(tclass *py.Object) *py.Object
//
//     Parameters (internal use only)
//     ---------
//     fp  : file-like object for tqdm
//     argv  : list (default: sys.argv[1:])
//
//
//go:linkname Main py.main
func Main(fp *py.Object, argv *py.Object) *py.Object
// Shortcut for `tqdm.gui.tqdm(range(*args), **kwargs)`.
//
//go:linkname Tgrange py.tgrange
func Tgrange(__llgo_va_list ...interface{}) *py.Object
// Shortcut for tqdm(range(*args), **kwargs).
//
//go:linkname Trange py.trange
func Trange(__llgo_va_list ...interface{}) *py.Object
// See tqdm.notebook.tqdm for full documentation
//
//go:linkname TqdmNotebook py.tqdm_notebook
func TqdmNotebook(__llgo_va_list ...interface{}) *py.Object
// Shortcut for `tqdm.notebook.tqdm(range(*args), **kwargs)`.
//
//go:linkname Tnrange py.tnrange
func Tnrange(__llgo_va_list ...interface{}) *py.Object
