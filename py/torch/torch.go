package torch

import (
	"github.com/goplus/lib/py"
	_ "unsafe"
)

const LLGoPackage = "py.torch"
// None
//
//go:linkname Classproperty py.classproperty
func Classproperty(func_ *py.Object) *py.Object
// None
//
//go:linkname GetFilePath py.get_file_path
func GetFilePath(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname PrepareMultiprocessingEnvironment py.prepare_multiprocessing_environment
func PrepareMultiprocessingEnvironment(path *py.Object) *py.Object
//
// get_num_threads() -> int
//
// Returns the number of threads used for parallelizing CPU operations
//
//
//go:linkname GetNumThreads py.get_num_threads
func GetNumThreads() *py.Object
//
// set_num_threads(int)
//
// Sets the number of threads used for intraop parallelism on CPU.
//
// .. warning::
//     To ensure that the correct number of threads is used, set_num_threads
//     must be called before running eager, JIT or autograd code.
//
//
//go:linkname SetNumThreads py.set_num_threads
func SetNumThreads(int *py.Object) *py.Object
//
// get_num_interop_threads() -> int
//
// Returns the number of threads used for inter-op parallelism on CPU
// (e.g. in JIT interpreter)
//
//
//go:linkname GetNumInteropThreads py.get_num_interop_threads
func GetNumInteropThreads() *py.Object
//
// set_num_interop_threads(int)
//
// Sets the number of threads used for interop parallelism
// (e.g. in JIT interpreter) on CPU.
//
// .. warning::
//     Can only be called once and before any inter-op parallel work
//     is started (e.g. JIT execution).
//
//
//go:linkname SetNumInteropThreads py.set_num_interop_threads
func SetNumInteropThreads(int *py.Object) *py.Object
//
// set_flush_denormal(mode) -> bool
//
// Disables denormal floating numbers on CPU.
//
// Returns ``True`` if your system supports flushing denormal numbers and it
// successfully configures flush denormal mode.  :meth:`~torch.set_flush_denormal`
// is only supported on x86 architectures supporting SSE3.
//
// Args:
//     mode (bool): Controls whether to enable flush denormal mode or not
//
// Example::
//
//     >>> torch.set_flush_denormal(True)
//     True
//     >>> torch.tensor([1e-323], dtype=torch.float64)
//     tensor([ 0.], dtype=torch.float64)
//     >>> torch.set_flush_denormal(False)
//     True
//     >>> torch.tensor([1e-323], dtype=torch.float64)
//     tensor(9.88131e-324 *
//            [ 1.0000], dtype=torch.float64)
//
//
//go:linkname SetFlushDenormal py.set_flush_denormal
func SetFlushDenormal(mode *py.Object) *py.Object
//
// get_default_dtype() -> torch.dtype
//
// Get the current default floating point :class:`torch.dtype`.
//
// Example::
//
//     >>> torch.get_default_dtype()  # initial default for floating point is torch.float32
//     torch.float32
//     >>> torch.set_default_dtype(torch.float64)
//     >>> torch.get_default_dtype()  # default is now changed to torch.float64
//     torch.float64
//     >>> torch.set_default_tensor_type(torch.FloatTensor)  # setting tensor type also affects this
//     >>> torch.get_default_dtype()  # changed to torch.float32, the dtype for torch.FloatTensor
//     torch.float32
//
//
//
//go:linkname GetDefaultDtype py.get_default_dtype
func GetDefaultDtype() *py.Object
//
// is_grad_enabled() -> (bool)
//
// Returns True if grad mode is currently enabled.
//
//
//go:linkname IsGradEnabled py.is_grad_enabled
func IsGradEnabled() *py.Object
//
// is_inference_mode_enabled() -> (bool)
//
// Returns True if inference mode is currently enabled.
//
//
//go:linkname IsInferenceModeEnabled py.is_inference_mode_enabled
func IsInferenceModeEnabled() *py.Object
// None
//
//go:linkname SetAutocastEnabled py.set_autocast_enabled
func SetAutocastEnabled(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname IsAutocastEnabled py.is_autocast_enabled
func IsAutocastEnabled(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname ClearAutocastCache py.clear_autocast_cache
func ClearAutocastCache(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname SetAutocastCpuEnabled py.set_autocast_cpu_enabled
func SetAutocastCpuEnabled(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname IsAutocastCpuEnabled py.is_autocast_cpu_enabled
func IsAutocastCpuEnabled(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname SetAutocastCpuDtype py.set_autocast_cpu_dtype
func SetAutocastCpuDtype(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname GetAutocastCpuDtype py.get_autocast_cpu_dtype
func GetAutocastCpuDtype(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname SetAutocastGpuDtype py.set_autocast_gpu_dtype
func SetAutocastGpuDtype(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname GetAutocastGpuDtype py.get_autocast_gpu_dtype
func GetAutocastGpuDtype(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname SetAutocastXlaEnabled py.set_autocast_xla_enabled
func SetAutocastXlaEnabled(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname IsAutocastXlaEnabled py.is_autocast_xla_enabled
func IsAutocastXlaEnabled(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname SetAutocastXlaDtype py.set_autocast_xla_dtype
func SetAutocastXlaDtype(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname GetAutocastXlaDtype py.get_autocast_xla_dtype
func GetAutocastXlaDtype(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname SetAutocastIpuEnabled py.set_autocast_ipu_enabled
func SetAutocastIpuEnabled(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname IsAutocastIpuEnabled py.is_autocast_ipu_enabled
func IsAutocastIpuEnabled(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname SetAutocastIpuDtype py.set_autocast_ipu_dtype
func SetAutocastIpuDtype(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname GetAutocastIpuDtype py.get_autocast_ipu_dtype
func GetAutocastIpuDtype(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname AutocastIncrementNesting py.autocast_increment_nesting
func AutocastIncrementNesting(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname AutocastDecrementNesting py.autocast_decrement_nesting
func AutocastDecrementNesting(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname IsAutocastCacheEnabled py.is_autocast_cache_enabled
func IsAutocastCacheEnabled(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname SetAutocastCacheEnabled py.set_autocast_cache_enabled
func SetAutocastCacheEnabled(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname SetAnomalyEnabled py.set_anomaly_enabled
func SetAnomalyEnabled(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname IsAnomalyEnabled py.is_anomaly_enabled
func IsAnomalyEnabled(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname IsAnomalyCheckNanEnabled py.is_anomaly_check_nan_enabled
func IsAnomalyCheckNanEnabled(__llgo_va_list ...interface{}) *py.Object
// parse_ir(input: str, parse_tensor_constants: bool = False) -> torch::jit::Graph
//
//
//go:linkname ParseIr py.parse_ir
func ParseIr(input *py.Object, parseTensorConstants *py.Object) *py.Object
// parse_schema(arg0: str) -> c10::FunctionSchema
//
//
//go:linkname ParseSchema py.parse_schema
func ParseSchema(arg0 *py.Object) *py.Object
// unify_type_list(arg0: List[c10::Type]) -> c10::Type
//
//
//go:linkname UnifyTypeList py.unify_type_list
func UnifyTypeList(arg0 *py.Object) *py.Object
// fork(*args, **kwargs) -> torch._C.Future
//
//
//go:linkname Fork py.fork
func Fork(__llgo_va_list ...interface{}) *py.Object
// wait(arg0: torch._C.Future) -> object
//
//
//go:linkname Wait py.wait
func Wait(arg0 *py.Object) *py.Object
// parse_type_comment(arg0: str) -> torch._C._jit_tree_views.Decl
//
//
//go:linkname ParseTypeComment py.parse_type_comment
func ParseTypeComment(arg0 *py.Object) *py.Object
// merge_type_from_type_comment(arg0: torch._C._jit_tree_views.Decl, arg1: torch._C._jit_tree_views.Decl, arg2: bool) -> torch._C._jit_tree_views.Decl
//
//
//go:linkname MergeTypeFromTypeComment py.merge_type_from_type_comment
func MergeTypeFromTypeComment(arg0 *py.Object, arg1 *py.Object, arg2 *py.Object) *py.Object
// import_ir_module(arg0: torch._C.CompilationUnit, arg1: str, arg2: object, arg3: dict, arg4: bool) -> torch._C.ScriptModule
//
//
//go:linkname ImportIrModule py.import_ir_module
func ImportIrModule(arg0 *py.Object, arg1 *py.Object, arg2 *py.Object, arg3 *py.Object, arg4 *py.Object) *py.Object
// import_ir_module_from_buffer(arg0: torch._C.CompilationUnit, arg1: str, arg2: object, arg3: dict, arg4: bool) -> torch._C.ScriptModule
//
//
//go:linkname ImportIrModuleFromBuffer py.import_ir_module_from_buffer
func ImportIrModuleFromBuffer(arg0 *py.Object, arg1 *py.Object, arg2 *py.Object, arg3 *py.Object, arg4 *py.Object) *py.Object
// vitals_enabled() -> bool
//
//
//go:linkname VitalsEnabled py.vitals_enabled
func VitalsEnabled() *py.Object
// set_vital(arg0: str, arg1: str, arg2: str) -> bool
//
//
//go:linkname SetVital py.set_vital
func SetVital(arg0 *py.Object, arg1 *py.Object, arg2 *py.Object) *py.Object
// read_vitals() -> str
//
//
//go:linkname ReadVitals py.read_vitals
func ReadVitals() *py.Object
// init_num_threads() -> None
//
//
// init_num_threads()
//
// Initializes the number of parallel threads used on the current thread.
//
// Call this whenever a new thread is created in order to propagate values from
// :func:`torch.set_num_threads` onto the new thread.
//
//
//
//go:linkname InitNumThreads py.init_num_threads
func InitNumThreads() *py.Object
//  SymInt-aware utility for logical negation.
//
//     Args:
//         a (SymBool or bool): Object to negate
//
//
//go:linkname SymNot py.sym_not
func SymNot(a *py.Object) *py.Object
//  SymInt-aware utility for float casting.
//
//     Args:
//         a (SymInt, SymFloat, or object): Object to cast
//
//
//go:linkname SymFloat py.sym_float
func SymFloat(a *py.Object) *py.Object
//  SymInt-aware utility for int casting.
//
//     Args:
//         a (SymInt, SymFloat, or object): Object to cast
//
//
//go:linkname SymInt py.sym_int
func SymInt(a *py.Object) *py.Object
//  SymInt-aware utility for max().
//
//go:linkname SymMax py.sym_max
func SymMax(a *py.Object, b *py.Object) *py.Object
//  SymInt-aware utility for max().
//
//go:linkname SymMin py.sym_min
func SymMin(a *py.Object, b *py.Object) *py.Object
// None
//
//go:linkname SymSqrt py.sym_sqrt
func SymSqrt(a *py.Object) *py.Object
// None
//
//go:linkname SymIte py.sym_ite
func SymIte(b *py.Object, t *py.Object, f *py.Object) *py.Object
//
// zeros_like(input, *, dtype=None, layout=None, device=None, requires_grad=False, memory_format=torch.preserve_format) -> Tensor
//
// Returns a tensor filled with the scalar value `0`, with the same size as
// :attr:`input`. ``torch.zeros_like(input)`` is equivalent to
// ``torch.zeros(input.size(), dtype=input.dtype, layout=input.layout, device=input.device)``.
//
// .. warning::
//     As of 0.4, this function does not support an :attr:`out` keyword. As an alternative,
//     the old ``torch.zeros_like(input, out=output)`` is equivalent to
//     ``torch.zeros(input.size(), out=output)``.
//
// Args:
//     input (Tensor): the size of :attr:`input` will determine size of the output tensor.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned Tensor.
//         Default: if ``None``, defaults to the dtype of :attr:`input`.
//     layout (:class:`torch.layout`, optional): the desired layout of returned tensor.
//         Default: if ``None``, defaults to the layout of :attr:`input`.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, defaults to the device of :attr:`input`.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     memory_format (:class:`torch.memory_format`, optional): the desired memory format of
//         returned Tensor. Default: ``torch.preserve_format``.
//
// Example::
//
//     >>> input = torch.empty(2, 3)
//     >>> torch.zeros_like(input)
//     tensor([[ 0.,  0.,  0.],
//             [ 0.,  0.,  0.]])
//
//
//go:linkname Obj py.obj
func Obj(__llgo_va_list ...interface{}) *py.Object
// wait(arg0: torch._C.Future) -> object
//
//
//go:linkname Candidate py.candidate
func Candidate(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Typename py.typename
func Typename(o *py.Object) *py.Object
// Returns True if `obj` is a PyTorch tensor.
//
//     Note that this function is simply doing ``isinstance(obj, Tensor)``.
//     Using that ``isinstance`` check is better for typechecking with mypy,
//     and more explicit - so it's recommended to use that instead of
//     ``is_tensor``.
//
//     Args:
//         obj (Object): Object to test
//     Example::
//
//         >>> x = torch.tensor([1, 2, 3])
//         >>> torch.is_tensor(x)
//         True
//
//
//
//go:linkname IsTensor py.is_tensor
func IsTensor(obj *py.Object) *py.Object
// Returns True if `obj` is a PyTorch storage object.
//
//     Args:
//         obj (Object): Object to test
//
//
//go:linkname IsStorage py.is_storage
func IsStorage(obj *py.Object) *py.Object
// Sets the default ``torch.Tensor`` to be allocated on ``device``.  This
//     does not affect factory function calls which are called with an explicit
//     ``device`` argument.  Factory calls will be performed as if they
//     were passed ``device`` as an argument.
//
//     To only temporarily change the default device instead of setting it
//     globally, use ``with torch.device(device):`` instead.
//
//     The default device is initially ``cpu``.  If you set the default tensor
//     device to another device (e.g., ``cuda``) without a device index, tensors
//     will be allocated on whatever the current device for the device type,
//     even after :func:`torch.cuda.set_device` is called.
//
//     .. warning::
//
//         This function imposes a slight performance cost on every Python
//         call to the torch API (not just factory functions).  If this
//         is causing problems for you, please comment on
//         https://github.com/pytorch/pytorch/issues/92701
//
//     .. note::
//
//         This doesn't affect functions that create tensors that share the same memory as the input, like:
//         :func:`torch.from_numpy` and :func:`torch.frombuffer`
//
//     Args:
//         device (device or string): the device to set as default
//
//     Example::
//
//         >>> # xdoctest: +SKIP("requires cuda, changes global state")
//         >>> torch.tensor([1.2, 3]).device
//         device(type='cpu')
//         >>> torch.set_default_device('cuda')  # current device is 0
//         >>> torch.tensor([1.2, 3]).device
//         device(type='cuda', index=0)
//         >>> torch.set_default_device('cuda:1')
//         >>> torch.tensor([1.2, 3]).device
//         device(type='cuda', index=1)
//
//
//
//go:linkname SetDefaultDevice py.set_default_device
func SetDefaultDevice(device *py.Object) *py.Object
// Sets the default ``torch.Tensor`` type to floating point tensor type
//     ``t``. This type will also be used as default floating point type for
//     type inference in :func:`torch.tensor`.
//
//     The default floating point tensor type is initially ``torch.FloatTensor``.
//
//     Args:
//         t (type or string): the floating point tensor type or its name
//
//     Example::
//
//         >>> # xdoctest: +SKIP("Other tests may have changed the default type. Can we reset it?")
//         >>> torch.tensor([1.2, 3]).dtype    # initial default for floating point is torch.float32
//         torch.float32
//         >>> torch.set_default_tensor_type(torch.DoubleTensor)
//         >>> torch.tensor([1.2, 3]).dtype    # a new floating point tensor
//         torch.float64
//
//
//
//go:linkname SetDefaultTensorType py.set_default_tensor_type
func SetDefaultTensorType(t *py.Object) *py.Object
//
//
//     Sets the default floating point dtype to :attr:`d`. Supports torch.float32
//     and torch.float64 as inputs. Other dtypes may be accepted without complaint
//     but are not supported and are unlikely to work as expected.
//
//     When PyTorch is initialized its default floating point dtype is torch.float32,
//     and the intent of set_default_dtype(torch.float64) is to facilitate NumPy-like
//     type inference. The default floating point dtype is used to:
//
//     1. Implicitly determine the default complex dtype. When the default floating point
//        type is float32 the default complex dtype is complex64, and when the default
//        floating point type is float64 the default complex type is complex128.
//     2. Infer the dtype for tensors constructed using Python floats or complex Python
//        numbers. See examples below.
//     3. Determine the result of type promotion between bool and integer tensors and
//        Python floats and complex Python numbers.
//
//     Args:
//         d (:class:`torch.dtype`): the floating point dtype to make the default.
//                                   Either torch.float32 or torch.float64.
//
//     Example:
//         >>> # xdoctest: +SKIP("Other tests may have changed the default type. Can we reset it?")
//         >>> # initial default for floating point is torch.float32
//         >>> # Python floats are interpreted as float32
//         >>> torch.tensor([1.2, 3]).dtype
//         torch.float32
//         >>> # initial default for floating point is torch.complex64
//         >>> # Complex Python numbers are interpreted as complex64
//         >>> torch.tensor([1.2, 3j]).dtype
//         torch.complex64
//
//         >>> torch.set_default_dtype(torch.float64)
//
//         >>> # Python floats are now interpreted as float64
//         >>> torch.tensor([1.2, 3]).dtype    # a new floating point tensor
//         torch.float64
//         >>> # Complex Python numbers are now interpreted as complex128
//         >>> torch.tensor([1.2, 3j]).dtype   # a new complex tensor
//         torch.complex128
//
//
//
//go:linkname SetDefaultDtype py.set_default_dtype
func SetDefaultDtype(d *py.Object) *py.Object
//  Sets whether PyTorch operations must use "deterministic"
//     algorithms. That is, algorithms which, given the same input, and when
//     run on the same software and hardware, always produce the same output.
//     When enabled, operations will use deterministic algorithms when available,
//     and if only nondeterministic algorithms are available they will throw a
//     :class:`RuntimeError` when called.
//
//     .. note:: This setting alone is not always enough to make an application
//         reproducible. Refer to :ref:`reproducibility` for more information.
//
//     .. note:: :func:`torch.set_deterministic_debug_mode` offers an alternative
//         interface for this feature.
//
//     The following normally-nondeterministic operations will act
//     deterministically when ``mode=True``:
//
//         * :class:`torch.nn.Conv1d` when called on CUDA tensor
//         * :class:`torch.nn.Conv2d` when called on CUDA tensor
//         * :class:`torch.nn.Conv3d` when called on CUDA tensor
//         * :class:`torch.nn.ConvTranspose1d` when called on CUDA tensor
//         * :class:`torch.nn.ConvTranspose2d` when called on CUDA tensor
//         * :class:`torch.nn.ConvTranspose3d` when called on CUDA tensor
//         * :class:`torch.nn.ReplicationPad2d` when attempting to differentiate a CUDA tensor
//         * :func:`torch.bmm` when called on sparse-dense CUDA tensors
//         * :func:`torch.Tensor.__getitem__` when attempting to differentiate a CPU tensor
//           and the index is a list of tensors
//         * :func:`torch.Tensor.index_put` with ``accumulate=False``
//         * :func:`torch.Tensor.index_put` with ``accumulate=True`` when called on a CPU
//           tensor
//         * :func:`torch.Tensor.put_` with ``accumulate=True`` when called on a CPU
//           tensor
//         * :func:`torch.Tensor.scatter_add_` when called on a CUDA tensor
//         * :func:`torch.gather` when called on a CUDA tensor that requires grad
//         * :func:`torch.index_add` when called on CUDA tensor
//         * :func:`torch.index_select` when attempting to differentiate a CUDA tensor
//         * :func:`torch.repeat_interleave` when attempting to differentiate a CUDA tensor
//         * :func:`torch.Tensor.index_copy` when called on a CPU or CUDA tensor
//         * :func:`torch.Tensor.scatter` when `src` type is Tensor and called on CUDA tensor
//         * :func:`torch.Tensor.scatter_reduce` when ``reduce='sum'`` or ``reduce='mean'`` and called on CUDA tensor
//
//     The following normally-nondeterministic operations will throw a
//     :class:`RuntimeError` when ``mode=True``:
//
//         * :class:`torch.nn.AvgPool3d` when attempting to differentiate a CUDA tensor
//         * :class:`torch.nn.AdaptiveAvgPool2d` when attempting to differentiate a CUDA tensor
//         * :class:`torch.nn.AdaptiveAvgPool3d` when attempting to differentiate a CUDA tensor
//         * :class:`torch.nn.MaxPool3d` when attempting to differentiate a CUDA tensor
//         * :class:`torch.nn.AdaptiveMaxPool2d` when attempting to differentiate a CUDA tensor
//         * :class:`torch.nn.FractionalMaxPool2d` when attempting to differentiate a CUDA tensor
//         * :class:`torch.nn.FractionalMaxPool3d` when attempting to differentiate a CUDA tensor
//         * :class:`torch.nn.MaxUnpool1d`
//         * :class:`torch.nn.MaxUnpool2d`
//         * :class:`torch.nn.MaxUnpool3d`
//         * :func:`torch.nn.functional.interpolate` when attempting to differentiate a CUDA tensor
//           and one of the following modes is used:
//
//           - ``linear``
//           - ``bilinear``
//           - ``bicubic``
//           - ``trilinear``
//
//         * :class:`torch.nn.ReflectionPad1d` when attempting to differentiate a CUDA tensor
//         * :class:`torch.nn.ReflectionPad2d` when attempting to differentiate a CUDA tensor
//         * :class:`torch.nn.ReflectionPad3d` when attempting to differentiate a CUDA tensor
//         * :class:`torch.nn.ReplicationPad1d` when attempting to differentiate a CUDA tensor
//         * :class:`torch.nn.ReplicationPad3d` when attempting to differentiate a CUDA tensor
//         * :class:`torch.nn.NLLLoss` when called on a CUDA tensor
//         * :class:`torch.nn.CTCLoss` when attempting to differentiate a CUDA tensor
//         * :class:`torch.nn.EmbeddingBag` when attempting to differentiate a CUDA tensor when
//           ``mode='max'``
//         * :func:`torch.Tensor.put_` when ``accumulate=False``
//         * :func:`torch.Tensor.put_` when ``accumulate=True`` and called on a CUDA tensor
//         * :func:`torch.histc` when called on a CUDA tensor
//         * :func:`torch.bincount` when called on a CUDA tensor and ``weights``
//           tensor is given
//         * :func:`torch.kthvalue` with called on a CUDA tensor
//         * :func:`torch.median` with indices output when called on a CUDA tensor
//         * :func:`torch.nn.functional.grid_sample` when attempting to differentiate a CUDA tensor
//         * :func:`torch.cumsum` when called on a CUDA tensor when dtype is floating point or complex
//         * :func:`torch.Tensor.scatter_reduce` when ``reduce='prod'`` and called on CUDA tensor
//         * :func:`torch.Tensor.resize_` when called with a quantized tensor
//
//     In addition, several operations fill uninitialized memory when this setting
//     is turned on and when
//     :attr:`torch.utils.deterministic.fill_uninitialized_memory` is turned on.
//     See the documentation for that attribute for more information.
//
//     A handful of CUDA operations are nondeterministic if the CUDA version is
//     10.2 or greater, unless the environment variable ``CUBLAS_WORKSPACE_CONFIG=:4096:8``
//     or ``CUBLAS_WORKSPACE_CONFIG=:16:8`` is set. See the CUDA documentation for more
//     details: `<https://docs.nvidia.com/cuda/cublas/index.html#cublasApi_reproducibility>`_
//     If one of these environment variable configurations is not set, a :class:`RuntimeError`
//     will be raised from these operations when called with CUDA tensors:
//
//         * :func:`torch.mm`
//         * :func:`torch.mv`
//         * :func:`torch.bmm`
//
//     Note that deterministic operations tend to have worse performance than
//     nondeterministic operations.
//
//     .. note::
//
//         This flag does not detect or prevent nondeterministic behavior caused
//         by calling an inplace operation on a tensor with an internal memory
//         overlap or by giving such a tensor as the :attr:`out` argument for an
//         operation. In these cases, multiple writes of different data may target
//         a single memory location, and the order of writes is not guaranteed.
//
//     Args:
//         mode (:class:`bool`): If True, makes potentially nondeterministic
//             operations switch to a deterministic algorithm or throw a runtime
//             error. If False, allows nondeterministic operations.
//
//     Keyword args:
//         warn_only (:class:`bool`, optional): If True, operations that do not
//             have a deterministic implementation will throw a warning instead of
//             an error. Default: ``False``
//
//     Example::
//
//         >>> # xdoctest: +SKIP
//         >>> torch.use_deterministic_algorithms(True)
//
//         # Forward mode nondeterministic error
//         >>> torch.randn(10, device='cuda').kthvalue(1)
//         ...
//         RuntimeError: kthvalue CUDA does not have a deterministic implementation...
//
//         # Backward mode nondeterministic error
//         >>> torch.nn.AvgPool3d(1)(torch.randn(3, 4, 5, 6, requires_grad=True).cuda()).sum().backward()
//         ...
//         RuntimeError: avg_pool3d_backward_cuda does not have a deterministic implementation...
//
//
//go:linkname UseDeterministicAlgorithms py.use_deterministic_algorithms
func UseDeterministicAlgorithms(mode *py.Object) *py.Object
// Returns True if the global deterministic flag is turned on. Refer to
//     :func:`torch.use_deterministic_algorithms` documentation for more details.
//
//
//go:linkname AreDeterministicAlgorithmsEnabled py.are_deterministic_algorithms_enabled
func AreDeterministicAlgorithmsEnabled() *py.Object
// Returns True if the global deterministic flag is set to warn only.
//     Refer to :func:`torch.use_deterministic_algorithms` documentation for more
//     details.
//
//
//go:linkname IsDeterministicAlgorithmsWarnOnlyEnabled py.is_deterministic_algorithms_warn_only_enabled
func IsDeterministicAlgorithmsWarnOnlyEnabled() *py.Object
// Sets the debug mode for deterministic operations.
//
//     .. note:: This is an alternative interface for
//         :func:`torch.use_deterministic_algorithms`. Refer to that function's
//         documentation for details about affected operations.
//
//     Args:
//         debug_mode(str or int): If "default" or 0, don't error or warn on
//             nondeterministic operations. If "warn" or 1, warn on
//             nondeterministic operations. If "error" or 2, error on
//             nondeterministic operations.
//
//
//go:linkname SetDeterministicDebugMode py.set_deterministic_debug_mode
func SetDeterministicDebugMode(debugMode *py.Object) *py.Object
// Returns the current value of the debug mode for deterministic
//     operations. Refer to :func:`torch.set_deterministic_debug_mode`
//     documentation for more details.
//
//
//go:linkname GetDeterministicDebugMode py.get_deterministic_debug_mode
func GetDeterministicDebugMode() *py.Object
// Returns the current value of float32 matrix multiplication precision. Refer to
//     :func:`torch.set_float32_matmul_precision` documentation for more details.
//
//
//go:linkname GetFloat32MatmulPrecision py.get_float32_matmul_precision
func GetFloat32MatmulPrecision() *py.Object
// Sets the internal precision of float32 matrix multiplications.
//
//     Running float32 matrix multiplications in lower precision may significantly increase
//     performance, and in some programs the loss of precision has a negligible impact.
//
//     Supports three settings:
//
//         * "highest", float32 matrix multiplications use the float32 datatype (24 mantissa
//           bits) for internal computations.
//         * "high", float32 matrix multiplications either use the TensorFloat32 datatype (10
//           mantissa bits) or treat each float32 number as the sum of two bfloat16 numbers
//           (approximately 16 mantissa bits), if the appropriate fast matrix multiplication
//           algorithms are available.  Otherwise float32 matrix multiplications are computed
//           as if the precision is "highest".  See below for more information on the bfloat16
//           approach.
//         * "medium", float32 matrix multiplications use the bfloat16 datatype (8 mantissa
//           bits) for internal computations, if a fast matrix multiplication algorithm
//           using that datatype internally is available. Otherwise float32
//           matrix multiplications are computed as if the precision is "high".
//
//     When using "high" precision, float32 multiplications may use a bfloat16-based algorithm
//     that is more complicated than simply truncating to some smaller number mantissa bits
//     (e.g. 10 for TensorFloat32, 8 for bfloat16).  Refer to [Henry2019]_ for a complete
//     description of this algorithm.  To briefly explain here, the first step is to realize
//     that we can perfectly encode a single float32 number as the sum of three bfloat16
//     numbers (because float32 has 24 mantissa bits while bfloat16 has 8, and both have the
//     same number of exponent bits).  This means that the product of two float32 numbers can
//     be exactly given by the sum of nine products of bfloat16 numbers.  We can then trade
//     accuracy for speed by dropping some of these products.  The "high" precision algorithm
//     specifically keeps only the three most significant products, which conveniently excludes
//     all of the products involving the last 8 mantissa bits of either input.  This means that
//     we can represent our inputs as the sum of two bfloat16 numbers rather than three.
//     Because bfloat16 fused-multiply-add (FMA) instructions are typically >10x faster than
//     float32 ones, it's faster to do three multiplications and 2 additions with bfloat16
//     precision than it is to do a single multiplication with float32 precision.
//
//     .. [Henry2019] http://arxiv.org/abs/1904.06376
//
//     .. note::
//
//         This does not change the output dtype of float32 matrix multiplications,
//         it controls how the internal computation of the matrix multiplication is performed.
//
//     .. note::
//
//         This does not change the precision of convolution operations. Other flags,
//         like `torch.backends.cudnn.allow_tf32`, may control the precision of convolution
//         operations.
//
//     .. note::
//
//         This flag currently only affects one native device type: CUDA.
//         If "high" or "medium" are set then the TensorFloat32 datatype will be used
//         when computing float32 matrix multiplications, equivalent to setting
//         `torch.backends.cuda.matmul.allow_tf32 = True`. When "highest" (the default)
//         is set then the float32 datatype is used for internal computations, equivalent
//         to setting `torch.backends.cuda.matmul.allow_tf32 = False`.
//
//     Args:
//         precision(str): can be set to "highest" (default), "high", or "medium" (see above).
//
//
//
//go:linkname SetFloat32MatmulPrecision py.set_float32_matmul_precision
func SetFloat32MatmulPrecision(precision *py.Object) *py.Object
// When this flag is False (default) then some PyTorch warnings may only
//     appear once per process. This helps avoid excessive warning information.
//     Setting it to True causes these warnings to always appear, which may be
//     helpful when debugging.
//
//     Args:
//         b (:class:`bool`): If True, force warnings to always be emitted
//                            If False, set to the default behaviour
//
//
//go:linkname SetWarnAlways py.set_warn_always
func SetWarnAlways(b *py.Object) *py.Object
// Returns True if the global warn_always flag is turned on. Refer to
//     :func:`torch.set_warn_always` documentation for more details.
//
//
//go:linkname IsWarnAlwaysEnabled py.is_warn_always_enabled
func IsWarnAlwaysEnabled() *py.Object
// Sets the random number generator state.
//
//     .. note: This function only works for CPU. For CUDA, please use
//              torch.manual_seed(seed), which works for both CPU and CUDA.
//
//     Args:
//         new_state (torch.ByteTensor): The desired state
//
//
//go:linkname SetRngState py.set_rng_state
func SetRngState(newState *py.Object) *py.Object
// Returns the random number generator state as a `torch.ByteTensor`.
//
//go:linkname GetRngState py.get_rng_state
func GetRngState() *py.Object
// Sets the seed for generating random numbers. Returns a
//     `torch.Generator` object.
//
//     Args:
//         seed (int): The desired seed. Value must be within the inclusive range
//             `[-0x8000_0000_0000_0000, 0xffff_ffff_ffff_ffff]`. Otherwise, a RuntimeError
//             is raised. Negative inputs are remapped to positive values with the formula
//             `0xffff_ffff_ffff_ffff + seed`.
//
//
//go:linkname ManualSeed py.manual_seed
func ManualSeed(seed *py.Object) *py.Object
// Returns the initial seed for generating random numbers as a
//     Python `long`.
//
//
//go:linkname InitialSeed py.initial_seed
func InitialSeed() *py.Object
// Sets the seed for generating random numbers to a non-deterministic
//     random number. Returns a 64 bit number used to seed the RNG.
//
//
//go:linkname Seed py.seed
func Seed() *py.Object
// save(obj, f, pickle_module=pickle, pickle_protocol=DEFAULT_PROTOCOL, _use_new_zipfile_serialization=True)
//
//     Saves an object to a disk file.
//
//     See also: :ref:`saving-loading-tensors`
//
//     Args:
//         obj: saved object
//         f: a file-like object (has to implement write and flush) or a string or
//            os.PathLike object containing a file name
//         pickle_module: module used for pickling metadata and objects
//         pickle_protocol: can be specified to override the default protocol
//
//     .. note::
//         A common PyTorch convention is to save tensors using .pt file extension.
//
//     .. note::
//         PyTorch preserves storage sharing across serialization. See
//         :ref:`preserve-storage-sharing` for more details.
//
//     .. note::
//         The 1.6 release of PyTorch switched ``torch.save`` to use a new
//         zipfile-based file format. ``torch.load`` still retains the ability to
//         load files in the old format. If for any reason you want ``torch.save``
//         to use the old format, pass the kwarg ``_use_new_zipfile_serialization=False``.
//
//     Example:
//         >>> # xdoctest: +SKIP("makes cwd dirty")
//         >>> # Save to file
//         >>> x = torch.tensor([0, 1, 2, 3, 4])
//         >>> torch.save(x, 'tensor.pt')
//         >>> # Save to io.BytesIO buffer
//         >>> buffer = io.BytesIO()
//         >>> torch.save(x, buffer)
//
//
//go:linkname Save py.save
func Save(obj *py.Object, f *py.Object, pickleModule *py.Object, pickleProtocol *py.Object, UseNewZipfileSerialization *py.Object, DisableByteorderRecord *py.Object) *py.Object
// load(f, map_location=None, pickle_module=pickle, *, weights_only=False, mmap=None, **pickle_load_args)
//
//     Loads an object saved with :func:`torch.save` from a file.
//
//     :func:`torch.load` uses Python's unpickling facilities but treats storages,
//     which underlie tensors, specially. They are first deserialized on the
//     CPU and are then moved to the device they were saved from. If this fails
//     (e.g. because the run time system doesn't have certain devices), an exception
//     is raised. However, storages can be dynamically remapped to an alternative
//     set of devices using the :attr:`map_location` argument.
//
//     If :attr:`map_location` is a callable, it will be called once for each serialized
//     storage with two arguments: storage and location. The storage argument
//     will be the initial deserialization of the storage, residing on the CPU.
//     Each serialized storage has a location tag associated with it which
//     identifies the device it was saved from, and this tag is the second
//     argument passed to :attr:`map_location`. The builtin location tags are ``'cpu'``
//     for CPU tensors and ``'cuda:device_id'`` (e.g. ``'cuda:2'``) for CUDA tensors.
//     :attr:`map_location` should return either ``None`` or a storage. If
//     :attr:`map_location` returns a storage, it will be used as the final deserialized
//     object, already moved to the right device. Otherwise, :func:`torch.load` will
//     fall back to the default behavior, as if :attr:`map_location` wasn't specified.
//
//     If :attr:`map_location` is a :class:`torch.device` object or a string containing
//     a device tag, it indicates the location where all tensors should be loaded.
//
//     Otherwise, if :attr:`map_location` is a dict, it will be used to remap location tags
//     appearing in the file (keys), to ones that specify where to put the
//     storages (values).
//
//     User extensions can register their own location tags and tagging and
//     deserialization methods using :func:`torch.serialization.register_package`.
//
//     Args:
//         f: a file-like object (has to implement :meth:`read`, :meth:`readline`, :meth:`tell`, and :meth:`seek`),
//             or a string or os.PathLike object containing a file name
//         map_location: a function, :class:`torch.device`, string or a dict specifying how to remap storage
//             locations
//         pickle_module: module used for unpickling metadata and objects (has to
//             match the :attr:`pickle_module` used to serialize file)
//         weights_only: Indicates whether unpickler should be restricted to
//             loading only tensors, primitive types and dictionaries
//         mmap: Indicates whether the file should be mmaped rather than loading all the storages into memory.
//             Typically, tensor storages in the file will first be moved from disk to CPU memory, after which they
//             are moved to the location that they were tagged with when saving, or specified by ``map_location``. This
//             second step is a no-op if the final location is CPU. When the ``mmap`` flag is set, instead of copying the
//             tensor storages from disk to CPU memory in the first step, ``f`` is mmaped.
//         pickle_load_args: (Python 3 only) optional keyword arguments passed over to
//             :func:`pickle_module.load` and :func:`pickle_module.Unpickler`, e.g.,
//             :attr:`errors=...`.
//
//     .. warning::
//         :func:`torch.load()` unless `weights_only` parameter is set to `True`,
//         uses ``pickle`` module implicitly, which is known to be insecure.
//         It is possible to construct malicious pickle data which will execute arbitrary code
//         during unpickling. Never load data that could have come from an untrusted
//         source in an unsafe mode, or that could have been tampered with. **Only load data you trust**.
//
//     .. note::
//         When you call :func:`torch.load()` on a file which contains GPU tensors, those tensors
//         will be loaded to GPU by default. You can call ``torch.load(.., map_location='cpu')``
//         and then :meth:`load_state_dict` to avoid GPU RAM surge when loading a model checkpoint.
//
//     .. note::
//         By default, we decode byte strings as ``utf-8``.  This is to avoid a common error
//         case ``UnicodeDecodeError: 'ascii' codec can't decode byte 0x...``
//         when loading files saved by Python 2 in Python 3.  If this default
//         is incorrect, you may use an extra :attr:`encoding` keyword argument to specify how
//         these objects should be loaded, e.g., :attr:`encoding='latin1'` decodes them
//         to strings using ``latin1`` encoding, and :attr:`encoding='bytes'` keeps them
//         as byte arrays which can be decoded later with ``byte_array.decode(...)``.
//
//     Example:
//         >>> # xdoctest: +SKIP("undefined filepaths")
//         >>> torch.load('tensors.pt', weights_only=True)
//         # Load all tensors onto the CPU
//         >>> torch.load('tensors.pt', map_location=torch.device('cpu'), weights_only=True)
//         # Load all tensors onto the CPU, using a function
//         >>> torch.load('tensors.pt', map_location=lambda storage, loc: storage, weights_only=True)
//         # Load all tensors onto GPU 1
//         >>> torch.load('tensors.pt', map_location=lambda storage, loc: storage.cuda(1), weights_only=True)
//         # Map tensors from GPU 1 to GPU 0
//         >>> torch.load('tensors.pt', map_location={'cuda:1': 'cuda:0'}, weights_only=True)
//         # Load tensor from io.BytesIO object
//         # Loading from a buffer setting weights_only=False, warning this can be unsafe
//         >>> with open('tensor.pt', 'rb') as f:
//         ...     buffer = io.BytesIO(f.read())
//         >>> torch.load(buffer, weights_only=False)
//         # Load a module with 'ascii' encoding for unpickling
//         # Loading from a module setting weights_only=False, warning this can be unsafe
//         >>> torch.load('module.pt', encoding='ascii', weights_only=False)
//
//
//go:linkname Load py.load
func Load(f *py.Object, mapLocation *py.Object, pickleModule *py.Object) *py.Object
// Set options for printing. Items shamelessly taken from NumPy
//
//     Args:
//         precision: Number of digits of precision for floating point output
//             (default = 4).
//         threshold: Total number of array elements which trigger summarization
//             rather than full `repr` (default = 1000).
//         edgeitems: Number of array items in summary at beginning and end of
//             each dimension (default = 3).
//         linewidth: The number of characters per line for the purpose of
//             inserting line breaks (default = 80). Thresholded matrices will
//             ignore this parameter.
//         profile: Sane defaults for pretty printing. Can override with any of
//             the above options. (any one of `default`, `short`, `full`)
//         sci_mode: Enable (True) or disable (False) scientific notation. If
//             None (default) is specified, the value is defined by
//             `torch._tensor_str._Formatter`. This value is automatically chosen
//             by the framework.
//
//     Example::
//
//         >>> # Limit the precision of elements
//         >>> torch.set_printoptions(precision=2)
//         >>> torch.tensor([1.12345])
//         tensor([1.12])
//         >>> # Limit the number of elements shown
//         >>> torch.set_printoptions(threshold=5)
//         >>> torch.arange(10)
//         tensor([0, 1, 2, ..., 7, 8, 9])
//         >>> # Restore defaults
//         >>> torch.set_printoptions(profile='default')
//         >>> torch.tensor([1.12345])
//         tensor([1.1235])
//         >>> torch.arange(10)
//         tensor([0, 1, 2, 3, 4, 5, 6, 7, 8, 9])
//
//
//
//go:linkname SetPrintoptions py.set_printoptions
func SetPrintoptions(precision *py.Object, threshold *py.Object, edgeitems *py.Object, linewidth *py.Object, profile *py.Object, sciMode *py.Object) *py.Object
//
// abs(input, *, out=None) -> Tensor
//
// Computes the absolute value of each element in :attr:`input`.
//
// .. math::
//     \text{out}_{i} = |\text{input}_{i}|
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.abs(torch.tensor([-1, -2, 3]))
//     tensor([ 1,  2,  3])
//
//
//go:linkname Abs py.abs
func Abs(input *py.Object) *py.Object
// None
//
//go:linkname Abs_ py.abs_
func Abs_(__llgo_va_list ...interface{}) *py.Object
//
// absolute(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.abs`
//
//
//go:linkname Absolute py.absolute
func Absolute(input *py.Object) *py.Object
//
// acos(input, *, out=None) -> Tensor
//
// Computes the inverse cosine of each element in :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \cos^{-1}(\text{input}_{i})
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([ 0.3348, -0.5889,  0.2005, -0.1584])
//     >>> torch.acos(a)
//     tensor([ 1.2294,  2.2004,  1.3690,  1.7298])
//
//
//go:linkname Acos py.acos
func Acos(input *py.Object) *py.Object
// None
//
//go:linkname Acos_ py.acos_
func Acos_(__llgo_va_list ...interface{}) *py.Object
//
// acosh(input, *, out=None) -> Tensor
//
// Returns a new tensor with the inverse hyperbolic cosine of the elements of :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \cosh^{-1}(\text{input}_{i})
//
// Note:
//     The domain of the inverse hyperbolic cosine is `[1, inf)` and values outside this range
//     will be mapped to ``NaN``, except for `+ INF` for which the output is mapped to `+ INF`.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword arguments:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4).uniform_(1, 2)
//     >>> a
//     tensor([ 1.3192, 1.9915, 1.9674, 1.7151 ])
//     >>> torch.acosh(a)
//     tensor([ 0.7791, 1.3120, 1.2979, 1.1341 ])
//
//
//go:linkname Acosh py.acosh
func Acosh(input *py.Object) *py.Object
// None
//
//go:linkname Acosh_ py.acosh_
func Acosh_(__llgo_va_list ...interface{}) *py.Object
//
// adaptive_avg_pool1d(input, output_size) -> Tensor
//
// Applies a 1D adaptive average pooling over an input signal composed of
// several input planes.
//
// See :class:`~torch.nn.AdaptiveAvgPool1d` for details and output shape.
//
// Args:
//     output_size: the target output size (single integer)
//
//
//go:linkname AdaptiveAvgPool1d py.adaptive_avg_pool1d
func AdaptiveAvgPool1d(input *py.Object, outputSize *py.Object) *py.Object
// None
//
//go:linkname AdaptiveMaxPool1d py.adaptive_max_pool1d
func AdaptiveMaxPool1d(__llgo_va_list ...interface{}) *py.Object
//
// add(input, other, *, alpha=1, out=None) -> Tensor
//
// Adds :attr:`other`, scaled by :attr:`alpha`, to :attr:`input`.
//
// .. math::
//     \text{{out}}_i = \text{{input}}_i + \text{{alpha}} \times \text{{other}}_i
//
//
// Supports :ref:`broadcasting to a common shape <broadcasting-semantics>`,
// :ref:`type promotion <type-promotion-doc>`, and integer, float, and complex inputs.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor or Number): the tensor or number to add to :attr:`input`.
//
// Keyword arguments:
//     alpha (Number): the multiplier for :attr:`other`.
//     out (Tensor, optional): the output tensor.
//
// Examples::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([ 0.0202,  1.0985,  1.3506, -0.6056])
//     >>> torch.add(a, 20)
//     tensor([ 20.0202,  21.0985,  21.3506,  19.3944])
//
//     >>> b = torch.randn(4)
//     >>> b
//     tensor([-0.9732, -0.3497,  0.6245,  0.4022])
//     >>> c = torch.randn(4, 1)
//     >>> c
//     tensor([[ 0.3743],
//             [-1.7724],
//             [-0.5811],
//             [-0.8017]])
//     >>> torch.add(b, c, alpha=10)
//     tensor([[  2.7695,   3.3930,   4.3672,   4.1450],
//             [-18.6971, -18.0736, -17.0994, -17.3216],
//             [ -6.7845,  -6.1610,  -5.1868,  -5.4090],
//             [ -8.9902,  -8.3667,  -7.3925,  -7.6147]])
//
//
//go:linkname Add py.add
func Add(input *py.Object, other *py.Object) *py.Object
//
// addbmm(input, batch1, batch2, *, beta=1, alpha=1, out=None) -> Tensor
//
// Performs a batch matrix-matrix product of matrices stored
// in :attr:`batch1` and :attr:`batch2`,
// with a reduced add step (all matrix multiplications get accumulated
// along the first dimension).
// :attr:`input` is added to the final result.
//
// :attr:`batch1` and :attr:`batch2` must be 3-D tensors each containing the
// same number of matrices.
//
// If :attr:`batch1` is a :math:`(b \times n \times m)` tensor, :attr:`batch2` is a
// :math:`(b \times m \times p)` tensor, :attr:`input` must be
// :ref:`broadcastable <broadcasting-semantics>` with a :math:`(n \times p)` tensor
// and :attr:`out` will be a :math:`(n \times p)` tensor.
//
// .. math::
//     out = \beta\ \text{input} + \alpha\ (\sum_{i=0}^{b-1} \text{batch1}_i \mathbin{@} \text{batch2}_i)
//
// If :attr:`beta` is 0, then :attr:`input` will be ignored, and `nan` and `inf` in
// it will not be propagated.
//
// For inputs of type `FloatTensor` or `DoubleTensor`, arguments :attr:`beta` and :attr:`alpha`
// must be real numbers, otherwise they should be integers.
//
// This operator supports :ref:`TensorFloat32<tf32_on_ampere>`.
//
// On certain ROCm devices, when using float16 inputs this module will use :ref:`different precision<fp16_on_mi200>` for backward.
//
// Args:
//     batch1 (Tensor): the first batch of matrices to be multiplied
//     batch2 (Tensor): the second batch of matrices to be multiplied
//
// Keyword args:
//     beta (Number, optional): multiplier for :attr:`input` (:math:`\beta`)
//     input (Tensor): matrix to be added
//     alpha (Number, optional): multiplier for `batch1 @ batch2` (:math:`\alpha`)
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> M = torch.randn(3, 5)
//     >>> batch1 = torch.randn(10, 3, 4)
//     >>> batch2 = torch.randn(10, 4, 5)
//     >>> torch.addbmm(M, batch1, batch2)
//     tensor([[  6.6311,   0.0503,   6.9768, -12.0362,  -2.1653],
//             [ -4.8185,  -1.4255,  -6.6760,   8.9453,   2.5743],
//             [ -3.8202,   4.3691,   1.0943,  -1.1109,   5.4730]])
//
//
//go:linkname Addbmm py.addbmm
func Addbmm(input *py.Object, batch1 *py.Object, batch2 *py.Object) *py.Object
//
// addcdiv(input, tensor1, tensor2, *, value=1, out=None) -> Tensor
//
// Performs the element-wise division of :attr:`tensor1` by :attr:`tensor2`,
// multiplies the result by the scalar :attr:`value` and adds it to :attr:`input`.
//
// .. warning::
//     Integer division with addcdiv is no longer supported, and in a future
//     release addcdiv will perform a true division of tensor1 and tensor2.
//     The historic addcdiv behavior can be implemented as
//     (input + value * torch.trunc(tensor1 / tensor2)).to(input.dtype)
//     for integer inputs and as (input + value * tensor1 / tensor2) for float inputs.
//     The future addcdiv behavior is just the latter implementation:
//     (input + value * tensor1 / tensor2), for all dtypes.
//
// .. math::
//     \text{out}_i = \text{input}_i + \text{value} \times \frac{\text{tensor1}_i}{\text{tensor2}_i}
//
//
// The shapes of :attr:`input`, :attr:`tensor1`, and :attr:`tensor2` must be
// :ref:`broadcastable <broadcasting-semantics>`.
//
// For inputs of type `FloatTensor` or `DoubleTensor`, :attr:`value` must be
// a real number, otherwise an integer.
//
// Args:
//     input (Tensor): the tensor to be added
//     tensor1 (Tensor): the numerator tensor
//     tensor2 (Tensor): the denominator tensor
//
// Keyword args:
//     value (Number, optional): multiplier for :math:`\text{tensor1} / \text{tensor2}`
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> t = torch.randn(1, 3)
//     >>> t1 = torch.randn(3, 1)
//     >>> t2 = torch.randn(1, 3)
//     >>> torch.addcdiv(t, t1, t2, value=0.1)
//     tensor([[-0.2312, -3.6496,  0.1312],
//             [-1.0428,  3.4292, -0.1030],
//             [-0.5369, -0.9829,  0.0430]])
//
//
//go:linkname Addcdiv py.addcdiv
func Addcdiv(input *py.Object, tensor1 *py.Object, tensor2 *py.Object) *py.Object
//
// addcmul(input, tensor1, tensor2, *, value=1, out=None) -> Tensor
//
// Performs the element-wise multiplication of :attr:`tensor1`
// by :attr:`tensor2`, multiplies the result by the scalar :attr:`value`
// and adds it to :attr:`input`.
//
// .. math::
//     \text{out}_i = \text{input}_i + \text{value} \times \text{tensor1}_i \times \text{tensor2}_i
//
// The shapes of :attr:`tensor`, :attr:`tensor1`, and :attr:`tensor2` must be
// :ref:`broadcastable <broadcasting-semantics>`.
//
// For inputs of type `FloatTensor` or `DoubleTensor`, :attr:`value` must be
// a real number, otherwise an integer.
//
// Args:
//     input (Tensor): the tensor to be added
//     tensor1 (Tensor): the tensor to be multiplied
//     tensor2 (Tensor): the tensor to be multiplied
//
// Keyword args:
//     value (Number, optional): multiplier for :math:`tensor1 .* tensor2`
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> t = torch.randn(1, 3)
//     >>> t1 = torch.randn(3, 1)
//     >>> t2 = torch.randn(1, 3)
//     >>> torch.addcmul(t, t1, t2, value=0.1)
//     tensor([[-0.8635, -0.6391,  1.6174],
//             [-0.7617, -0.5879,  1.7388],
//             [-0.8353, -0.6249,  1.6511]])
//
//
//go:linkname Addcmul py.addcmul
func Addcmul(input *py.Object, tensor1 *py.Object, tensor2 *py.Object) *py.Object
//
// addmm(input, mat1, mat2, *, beta=1, alpha=1, out=None) -> Tensor
//
// Performs a matrix multiplication of the matrices :attr:`mat1` and :attr:`mat2`.
// The matrix :attr:`input` is added to the final result.
//
// If :attr:`mat1` is a :math:`(n \times m)` tensor, :attr:`mat2` is a
// :math:`(m \times p)` tensor, then :attr:`input` must be
// :ref:`broadcastable <broadcasting-semantics>` with a :math:`(n \times p)` tensor
// and :attr:`out` will be a :math:`(n \times p)` tensor.
//
// :attr:`alpha` and :attr:`beta` are scaling factors on matrix-vector product between
// :attr:`mat1` and :attr:`mat2` and the added matrix :attr:`input` respectively.
//
// .. math::
//     \text{out} = \beta\ \text{input} + \alpha\ (\text{mat1}_i \mathbin{@} \text{mat2}_i)
//
// If :attr:`beta` is 0, then :attr:`input` will be ignored, and `nan` and `inf` in
// it will not be propagated.
//
// For inputs of type `FloatTensor` or `DoubleTensor`, arguments :attr:`beta` and
// :attr:`alpha` must be real numbers, otherwise they should be integers.
//
// This operation has support for arguments with :ref:`sparse layouts<sparse-docs>`. If
// :attr:`input` is sparse the result will have the same layout and if :attr:`out`
// is provided it must have the same layout as :attr:`input`.
//
//
// .. warning::
//     Sparse support is a beta feature and some layout(s)/dtype/device combinations may not be supported,
//     or may not have autograd support. If you notice missing functionality please
//     open a feature request.
//
// This operator supports :ref:`TensorFloat32<tf32_on_ampere>`.
//
// On certain ROCm devices, when using float16 inputs this module will use :ref:`different precision<fp16_on_mi200>` for backward.
//
// Args:
//     input (Tensor): matrix to be added
//     mat1 (Tensor): the first matrix to be matrix multiplied
//     mat2 (Tensor): the second matrix to be matrix multiplied
//
// Keyword args:
//     beta (Number, optional): multiplier for :attr:`input` (:math:`\beta`)
//     alpha (Number, optional): multiplier for :math:`mat1 @ mat2` (:math:`\alpha`)
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> M = torch.randn(2, 3)
//     >>> mat1 = torch.randn(2, 3)
//     >>> mat2 = torch.randn(3, 3)
//     >>> torch.addmm(M, mat1, mat2)
//     tensor([[-4.8716,  1.4671, -1.3746],
//             [ 0.7573, -3.9555, -2.8681]])
//
//
//go:linkname Addmm py.addmm
func Addmm(input *py.Object, mat1 *py.Object, mat2 *py.Object) *py.Object
//
// addmv(input, mat, vec, *, beta=1, alpha=1, out=None) -> Tensor
//
// Performs a matrix-vector product of the matrix :attr:`mat` and
// the vector :attr:`vec`.
// The vector :attr:`input` is added to the final result.
//
// If :attr:`mat` is a :math:`(n \times m)` tensor, :attr:`vec` is a 1-D tensor of
// size `m`, then :attr:`input` must be
// :ref:`broadcastable <broadcasting-semantics>` with a 1-D tensor of size `n` and
// :attr:`out` will be 1-D tensor of size `n`.
//
// :attr:`alpha` and :attr:`beta` are scaling factors on matrix-vector product between
// :attr:`mat` and :attr:`vec` and the added tensor :attr:`input` respectively.
//
// .. math::
//     \text{out} = \beta\ \text{input} + \alpha\ (\text{mat} \mathbin{@} \text{vec})
//
// If :attr:`beta` is 0, then :attr:`input` will be ignored, and `nan` and `inf` in
// it will not be propagated.
//
// For inputs of type `FloatTensor` or `DoubleTensor`, arguments :attr:`beta` and
// :attr:`alpha` must be real numbers, otherwise they should be integers.
//
// Args:
//     input (Tensor): vector to be added
//     mat (Tensor): matrix to be matrix multiplied
//     vec (Tensor): vector to be matrix multiplied
//
// Keyword args:
//     beta (Number, optional): multiplier for :attr:`input` (:math:`\beta`)
//     alpha (Number, optional): multiplier for :math:`mat @ vec` (:math:`\alpha`)
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> M = torch.randn(2)
//     >>> mat = torch.randn(2, 3)
//     >>> vec = torch.randn(3)
//     >>> torch.addmv(M, mat, vec)
//     tensor([-0.3768, -5.5565])
//
//
//go:linkname Addmv py.addmv
func Addmv(input *py.Object, mat *py.Object, vec *py.Object) *py.Object
// None
//
//go:linkname Addmv_ py.addmv_
func Addmv_(__llgo_va_list ...interface{}) *py.Object
//
// addr(input, vec1, vec2, *, beta=1, alpha=1, out=None) -> Tensor
//
// Performs the outer-product of vectors :attr:`vec1` and :attr:`vec2`
// and adds it to the matrix :attr:`input`.
//
// Optional values :attr:`beta` and :attr:`alpha` are scaling factors on the
// outer product between :attr:`vec1` and :attr:`vec2` and the added matrix
// :attr:`input` respectively.
//
// .. math::
//     \text{out} = \beta\ \text{input} + \alpha\ (\text{vec1} \otimes \text{vec2})
//
// If :attr:`beta` is 0, then :attr:`input` will be ignored, and `nan` and `inf` in
// it will not be propagated.
//
// If :attr:`vec1` is a vector of size `n` and :attr:`vec2` is a vector
// of size `m`, then :attr:`input` must be
// :ref:`broadcastable <broadcasting-semantics>` with a matrix of size
// :math:`(n \times m)` and :attr:`out` will be a matrix of size
// :math:`(n \times m)`.
//
// Args:
//     input (Tensor): matrix to be added
//     vec1 (Tensor): the first vector of the outer product
//     vec2 (Tensor): the second vector of the outer product
//
// Keyword args:
//     beta (Number, optional): multiplier for :attr:`input` (:math:`\beta`)
//     alpha (Number, optional): multiplier for :math:`\text{vec1} \otimes \text{vec2}` (:math:`\alpha`)
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> vec1 = torch.arange(1., 4.)
//     >>> vec2 = torch.arange(1., 3.)
//     >>> M = torch.zeros(3, 2)
//     >>> torch.addr(M, vec1, vec2)
//     tensor([[ 1.,  2.],
//             [ 2.,  4.],
//             [ 3.,  6.]])
//
//
//go:linkname Addr py.addr
func Addr(input *py.Object, vec1 *py.Object, vec2 *py.Object) *py.Object
//
// adjoint(Tensor) -> Tensor
// Returns a view of the tensor conjugated and with the last two dimensions transposed.
//
// ``x.adjoint()`` is equivalent to ``x.transpose(-2, -1).conj()`` for complex tensors and
// to ``x.transpose(-2, -1)`` for real tensors.
//
// Example::
//     >>> x = torch.arange(4, dtype=torch.float)
//     >>> A = torch.complex(x, x).reshape(2, 2)
//     >>> A
//     tensor([[0.+0.j, 1.+1.j],
//             [2.+2.j, 3.+3.j]])
//     >>> A.adjoint()
//     tensor([[0.-0.j, 2.-2.j],
//             [1.-1.j, 3.-3.j]])
//     >>> (A.adjoint() == A.mH).all()
//     tensor(True)
//
//
//go:linkname Adjoint py.adjoint
func Adjoint(Tensor *py.Object) *py.Object
// None
//
//go:linkname AffineGridGenerator py.affine_grid_generator
func AffineGridGenerator(__llgo_va_list ...interface{}) *py.Object
//
// Performs the same operation as :func:`torch.alias`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname AliasCopy py.alias_copy
func AliasCopy(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname AlignTensors py.align_tensors
func AlignTensors(__llgo_va_list ...interface{}) *py.Object
//
// all(input) -> Tensor
//
// Tests if all elements in :attr:`input` evaluate to `True`.
//
// .. note:: This function matches the behaviour of NumPy in returning
//           output of dtype `bool` for all supported dtypes except `uint8`.
//           For `uint8` the dtype of output is `uint8` itself.
//
// Example::
//
//     >>> a = torch.rand(1, 2).bool()
//     >>> a
//     tensor([[False, True]], dtype=torch.bool)
//     >>> torch.all(a)
//     tensor(False, dtype=torch.bool)
//     >>> a = torch.arange(0, 3)
//     >>> a
//     tensor([0, 1, 2])
//     >>> torch.all(a)
//     tensor(False)
//
// .. function:: all(input, dim, keepdim=False, *, out=None) -> Tensor
//    :noindex:
//
// For each row of :attr:`input` in the given dimension :attr:`dim`,
// returns `True` if all elements in the row evaluate to `True` and `False` otherwise.
//
//
// If :attr:`keepdim` is ``True``, the output tensor is of the same size
// as :attr:`input` except in the dimension(s) :attr:`dim` where it is of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in the
// output tensor having 1 (or ``len(dim)``) fewer dimension(s).
//
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int or tuple of ints): the dimension or dimensions to reduce.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.rand(4, 2).bool()
//     >>> a
//     tensor([[True, True],
//             [True, False],
//             [True, True],
//             [True, True]], dtype=torch.bool)
//     >>> torch.all(a, dim=1)
//     tensor([ True, False,  True,  True], dtype=torch.bool)
//     >>> torch.all(a, dim=0)
//     tensor([ True, False], dtype=torch.bool)
//
//
//go:linkname All py.all
func All(input *py.Object) *py.Object
//
// allclose(input, other, rtol=1e-05, atol=1e-08, equal_nan=False) -> bool
//
// This function checks if :attr:`input` and :attr:`other` satisfy the condition:
//
// .. math::
//     \lvert \text{input} - \text{other} \rvert \leq \texttt{atol} + \texttt{rtol} \times \lvert \text{other} \rvert
//
// elementwise, for all elements of :attr:`input` and :attr:`other`. The behaviour of this function is analogous to
// `numpy.allclose <https://docs.scipy.org/doc/numpy/reference/generated/numpy.allclose.html>`_
//
// Args:
//     input (Tensor): first tensor to compare
//     other (Tensor): second tensor to compare
//     atol (float, optional): absolute tolerance. Default: 1e-08
//     rtol (float, optional): relative tolerance. Default: 1e-05
//     equal_nan (bool, optional): if ``True``, then two ``NaN`` s will be considered equal. Default: ``False``
//
// Example::
//
//     >>> torch.allclose(torch.tensor([10000., 1e-07]), torch.tensor([10000.1, 1e-08]))
//     False
//     >>> torch.allclose(torch.tensor([10000., 1e-08]), torch.tensor([10000.1, 1e-09]))
//     True
//     >>> torch.allclose(torch.tensor([1.0, float('nan')]), torch.tensor([1.0, float('nan')]))
//     False
//     >>> torch.allclose(torch.tensor([1.0, float('nan')]), torch.tensor([1.0, float('nan')]), equal_nan=True)
//     True
//
//
//go:linkname Allclose py.allclose
func Allclose(input *py.Object, other *py.Object, rtol *py.Object, atol *py.Object, equalNan *py.Object) *py.Object
// None
//
//go:linkname AlphaDropout py.alpha_dropout
func AlphaDropout(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname AlphaDropout_ py.alpha_dropout_
func AlphaDropout_(__llgo_va_list ...interface{}) *py.Object
//
// amax(input, dim, keepdim=False, *, out=None) -> Tensor
//
// Returns the maximum value of each slice of the :attr:`input` tensor in the given
// dimension(s) :attr:`dim`.
//
// .. note::
//     The difference between ``max``/``min`` and ``amax``/``amin`` is:
//         - ``amax``/``amin`` supports reducing on multiple dimensions,
//         - ``amax``/``amin`` does not return indices,
//         - ``amax``/``amin`` evenly distributes gradient between equal values,
//           while ``max(dim)``/``min(dim)`` propagates gradient only to a single
//           index in the source tensor.
//
//
// If :attr:`keepdim` is ``True``, the output tensor is of the same size
// as :attr:`input` except in the dimension(s) :attr:`dim` where it is of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in the
// output tensor having 1 (or ``len(dim)``) fewer dimension(s).
//
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int or tuple of ints): the dimension or dimensions to reduce.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//   out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4, 4)
//     >>> a
//     tensor([[ 0.8177,  1.4878, -0.2491,  0.9130],
//             [-0.7158,  1.1775,  2.0992,  0.4817],
//             [-0.0053,  0.0164, -1.3738, -0.0507],
//             [ 1.9700,  1.1106, -1.0318, -1.0816]])
//     >>> torch.amax(a, 1)
//     tensor([1.4878, 2.0992, 0.0164, 1.9700])
//
//
//go:linkname Amax py.amax
func Amax(input *py.Object, dim *py.Object, keepdim *py.Object) *py.Object
//
// amin(input, dim, keepdim=False, *, out=None) -> Tensor
//
// Returns the minimum value of each slice of the :attr:`input` tensor in the given
// dimension(s) :attr:`dim`.
//
// .. note::
//     The difference between ``max``/``min`` and ``amax``/``amin`` is:
//         - ``amax``/``amin`` supports reducing on multiple dimensions,
//         - ``amax``/``amin`` does not return indices,
//         - ``amax``/``amin`` evenly distributes gradient between equal values,
//           while ``max(dim)``/``min(dim)`` propagates gradient only to a single
//           index in the source tensor.
//
//
// If :attr:`keepdim` is ``True``, the output tensor is of the same size
// as :attr:`input` except in the dimension(s) :attr:`dim` where it is of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in the
// output tensor having 1 (or ``len(dim)``) fewer dimension(s).
//
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int or tuple of ints): the dimension or dimensions to reduce.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//   out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4, 4)
//     >>> a
//     tensor([[ 0.6451, -0.4866,  0.2987, -1.3312],
//             [-0.5744,  1.2980,  1.8397, -0.2713],
//             [ 0.9128,  0.9214, -1.7268, -0.2995],
//             [ 0.9023,  0.4853,  0.9075, -1.6165]])
//     >>> torch.amin(a, 1)
//     tensor([-1.3312, -0.5744, -1.7268, -1.6165])
//
//
//go:linkname Amin py.amin
func Amin(input *py.Object, dim *py.Object, keepdim *py.Object) *py.Object
//
// aminmax(input, *, dim=None, keepdim=False, out=None) -> (Tensor min, Tensor max)
//
// Computes the minimum and maximum values of the :attr:`input` tensor.
//
// Args:
//     input (Tensor):
//         The input tensor
//
// Keyword Args:
//     dim (Optional[int]):
//         The dimension along which to compute the values. If `None`,
//         computes the values over the entire :attr:`input` tensor.
//         Default is `None`.
//     keepdim (bool):
//         If `True`, the reduced dimensions will be kept in the output
//         tensor as dimensions with size 1 for broadcasting, otherwise
//         they will be removed, as if calling (:func:`torch.squeeze`).
//         Default is `False`.
//     out (Optional[Tuple[Tensor, Tensor]]):
//         Optional tensors on which to write the result. Must have the same
//         shape and dtype as the expected output.
//         Default is `None`.
//
// Returns:
//     A named tuple `(min, max)` containing the minimum and maximum values.
//
// Raises:
//     RuntimeError
//         If any of the dimensions to compute the values over has size 0.
//
// .. note::
//     NaN values are propagated to the output if at least one value is NaN.
//
// .. seealso::
//     :func:`torch.amin` computes just the minimum value
//     :func:`torch.amax` computes just the maximum value
//
// Example::
//
//     >>> torch.aminmax(torch.tensor([1, -3, 5]))
//     torch.return_types.aminmax(
//     min=tensor(-3),
//     max=tensor(5))
//
//     >>> # aminmax propagates NaNs
//     >>> torch.aminmax(torch.tensor([1, -3, 5, torch.nan]))
//     torch.return_types.aminmax(
//     min=tensor(nan),
//     max=tensor(nan))
//
//     >>> t = torch.arange(10).view(2, 5)
//     >>> t
//     tensor([[0, 1, 2, 3, 4],
//             [5, 6, 7, 8, 9]])
//     >>> t.aminmax(dim=0, keepdim=True)
//     torch.return_types.aminmax(
//     min=tensor([[0, 1, 2, 3, 4]]),
//     max=tensor([[5, 6, 7, 8, 9]]))
//
//
//go:linkname Aminmax py.aminmax
func Aminmax(input *py.Object) *py.Object
//
// angle(input, *, out=None) -> Tensor
//
// Computes the element-wise angle (in radians) of the given :attr:`input` tensor.
//
// .. math::
//     \text{out}_{i} = angle(\text{input}_{i})
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// .. note:: Starting in PyTorch 1.8, angle returns pi for negative real numbers,
//           zero for non-negative real numbers, and propagates NaNs. Previously
//           the function would return zero for all real numbers and not propagate
//           floating-point NaNs.
//
// Example::
//
//     >>> torch.angle(torch.tensor([-1 + 1j, -2 + 2j, 3 - 3j]))*180/3.14159
//     tensor([ 135.,  135,  -45])
//
//
//go:linkname Angle py.angle
func Angle(input *py.Object) *py.Object
//
// any(input) -> Tensor
//
// Tests if any element in :attr:`input` evaluates to `True`.
//
// .. note:: This function matches the behaviour of NumPy in returning
//           output of dtype `bool` for all supported dtypes except `uint8`.
//           For `uint8` the dtype of output is `uint8` itself.
//
// Example::
//
//     >>> a = torch.rand(1, 2).bool()
//     >>> a
//     tensor([[False, True]], dtype=torch.bool)
//     >>> torch.any(a)
//     tensor(True, dtype=torch.bool)
//     >>> a = torch.arange(0, 3)
//     >>> a
//     tensor([0, 1, 2])
//     >>> torch.any(a)
//     tensor(True)
//
// .. function:: any(input, dim, keepdim=False, *, out=None) -> Tensor
//    :noindex:
//
// For each row of :attr:`input` in the given dimension :attr:`dim`,
// returns `True` if any element in the row evaluate to `True` and `False` otherwise.
//
//
// If :attr:`keepdim` is ``True``, the output tensor is of the same size
// as :attr:`input` except in the dimension(s) :attr:`dim` where it is of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in the
// output tensor having 1 (or ``len(dim)``) fewer dimension(s).
//
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int or tuple of ints): the dimension or dimensions to reduce.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4, 2) < 0
//     >>> a
//     tensor([[ True,  True],
//             [False,  True],
//             [ True,  True],
//             [False, False]])
//     >>> torch.any(a, 1)
//     tensor([ True,  True,  True, False])
//     >>> torch.any(a, 0)
//     tensor([True, True])
//
//
//go:linkname Any py.any
func Any(input *py.Object) *py.Object
//
// arange(start=0, end, step=1, *, out=None, dtype=None, layout=torch.strided, device=None, requires_grad=False) -> Tensor
//
// Returns a 1-D tensor of size :math:`\left\lceil \frac{\text{end} - \text{start}}{\text{step}} \right\rceil`
// with values from the interval ``[start, end)`` taken with common difference
// :attr:`step` beginning from `start`.
//
// Note that non-integer :attr:`step` is subject to floating point rounding errors when
// comparing against :attr:`end`; to avoid inconsistency, we advise subtracting a small epsilon from :attr:`end`
// in such cases.
//
// .. math::
//     \text{out}_{{i+1}} = \text{out}_{i} + \text{step}
//
// Args:
//     start (Number): the starting value for the set of points. Default: ``0``.
//     end (Number): the ending value for the set of points
//     step (Number): the gap between each pair of adjacent points. Default: ``1``.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`). If `dtype` is not given, infer the data type from the other input
//         arguments. If any of `start`, `end`, or `stop` are floating-point, the
//         `dtype` is inferred to be the default dtype, see
//         :meth:`~torch.get_default_dtype`. Otherwise, the `dtype` is inferred to
//         be `torch.int64`.
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
// Example::
//
//     >>> torch.arange(5)
//     tensor([ 0,  1,  2,  3,  4])
//     >>> torch.arange(1, 4)
//     tensor([ 1,  2,  3])
//     >>> torch.arange(1, 2.5, 0.5)
//     tensor([ 1.0000,  1.5000,  2.0000])
//
//
//go:linkname Arange py.arange
func Arange(start *py.Object, end *py.Object, step *py.Object) *py.Object
//
// arccos(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.acos`.
//
//
//go:linkname Arccos py.arccos
func Arccos(input *py.Object) *py.Object
// None
//
//go:linkname Arccos_ py.arccos_
func Arccos_(__llgo_va_list ...interface{}) *py.Object
//
// arccosh(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.acosh`.
//
//
//go:linkname Arccosh py.arccosh
func Arccosh(input *py.Object) *py.Object
// None
//
//go:linkname Arccosh_ py.arccosh_
func Arccosh_(__llgo_va_list ...interface{}) *py.Object
//
// arcsin(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.asin`.
//
//
//go:linkname Arcsin py.arcsin
func Arcsin(input *py.Object) *py.Object
// None
//
//go:linkname Arcsin_ py.arcsin_
func Arcsin_(__llgo_va_list ...interface{}) *py.Object
//
// arcsinh(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.asinh`.
//
//
//go:linkname Arcsinh py.arcsinh
func Arcsinh(input *py.Object) *py.Object
// None
//
//go:linkname Arcsinh_ py.arcsinh_
func Arcsinh_(__llgo_va_list ...interface{}) *py.Object
//
// arctan(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.atan`.
//
//
//go:linkname Arctan py.arctan
func Arctan(input *py.Object) *py.Object
//
// arctan2(input, other, *, out=None) -> Tensor
// Alias for :func:`torch.atan2`.
//
//
//go:linkname Arctan2 py.arctan2
func Arctan2(input *py.Object, other *py.Object) *py.Object
// None
//
//go:linkname Arctan_ py.arctan_
func Arctan_(__llgo_va_list ...interface{}) *py.Object
//
// arctanh(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.atanh`.
//
//
//go:linkname Arctanh py.arctanh
func Arctanh(input *py.Object) *py.Object
// None
//
//go:linkname Arctanh_ py.arctanh_
func Arctanh_(__llgo_va_list ...interface{}) *py.Object
//
// argmax(input) -> LongTensor
//
// Returns the indices of the maximum value of all elements in the :attr:`input` tensor.
//
// This is the second value returned by :meth:`torch.max`. See its
// documentation for the exact semantics of this method.
//
// .. note:: If there are multiple maximal values then the indices of the first maximal value are returned.
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> a = torch.randn(4, 4)
//     >>> a
//     tensor([[ 1.3398,  0.2663, -0.2686,  0.2450],
//             [-0.7401, -0.8805, -0.3402, -1.1936],
//             [ 0.4907, -1.3948, -1.0691, -0.3132],
//             [-1.6092,  0.5419, -0.2993,  0.3195]])
//     >>> torch.argmax(a)
//     tensor(0)
//
// .. function:: argmax(input, dim, keepdim=False) -> LongTensor
//    :noindex:
//
// Returns the indices of the maximum values of a tensor across a dimension.
//
// This is the second value returned by :meth:`torch.max`. See its
// documentation for the exact semantics of this method.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int): the dimension to reduce. If ``None``, the argmax of the flattened input is returned.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Example::
//
//     >>> a = torch.randn(4, 4)
//     >>> a
//     tensor([[ 1.3398,  0.2663, -0.2686,  0.2450],
//             [-0.7401, -0.8805, -0.3402, -1.1936],
//             [ 0.4907, -1.3948, -1.0691, -0.3132],
//             [-1.6092,  0.5419, -0.2993,  0.3195]])
//     >>> torch.argmax(a, dim=1)
//     tensor([ 0,  2,  0,  1])
//
//
//go:linkname Argmax py.argmax
func Argmax(input *py.Object) *py.Object
//
// argmin(input, dim=None, keepdim=False) -> LongTensor
//
// Returns the indices of the minimum value(s) of the flattened tensor or along a dimension
//
// This is the second value returned by :meth:`torch.min`. See its
// documentation for the exact semantics of this method.
//
// .. note:: If there are multiple minimal values then the indices of the first minimal value are returned.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int): the dimension to reduce. If ``None``, the argmin of the flattened input is returned.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Example::
//
//     >>> a = torch.randn(4, 4)
//     >>> a
//     tensor([[ 0.1139,  0.2254, -0.1381,  0.3687],
//             [ 1.0100, -1.1975, -0.0102, -0.4732],
//             [-0.9240,  0.1207, -0.7506, -1.0213],
//             [ 1.7809, -1.2960,  0.9384,  0.1438]])
//     >>> torch.argmin(a)
//     tensor(13)
//     >>> torch.argmin(a, dim=1)
//     tensor([ 2,  1,  3,  1])
//     >>> torch.argmin(a, dim=1, keepdim=True)
//     tensor([[2],
//             [1],
//             [3],
//             [1]])
//
//
//go:linkname Argmin py.argmin
func Argmin(input *py.Object, dim *py.Object, keepdim *py.Object) *py.Object
//
// argsort(input, dim=-1, descending=False, stable=False) -> Tensor
//
// Returns the indices that sort a tensor along a given dimension in ascending
// order by value.
//
// This is the second value returned by :meth:`torch.sort`.  See its documentation
// for the exact semantics of this method.
//
// If :attr:`stable` is ``True`` then the sorting routine becomes stable, preserving
// the order of equivalent elements. If ``False``, the relative order of values
// which compare equal is not guaranteed. ``True`` is slower.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int, optional): the dimension to sort along
//     descending (bool, optional): controls the sorting order (ascending or descending)
//     stable (bool, optional): controls the relative order of equivalent elements
//
// Example::
//
//     >>> a = torch.randn(4, 4)
//     >>> a
//     tensor([[ 0.0785,  1.5267, -0.8521,  0.4065],
//             [ 0.1598,  0.0788, -0.0745, -1.2700],
//             [ 1.2208,  1.0722, -0.7064,  1.2564],
//             [ 0.0669, -0.2318, -0.8229, -0.9280]])
//
//
//     >>> torch.argsort(a, dim=1)
//     tensor([[2, 0, 3, 1],
//             [3, 2, 1, 0],
//             [2, 1, 0, 3],
//             [3, 2, 1, 0]])
//
//
//go:linkname Argsort py.argsort
func Argsort(input *py.Object, dim *py.Object, descending *py.Object, stable *py.Object) *py.Object
//
// argwhere(input) -> Tensor
//
// Returns a tensor containing the indices of all non-zero elements of
// :attr:`input`.  Each row in the result contains the indices of a non-zero
// element in :attr:`input`. The result is sorted lexicographically, with
// the last index changing the fastest (C-style).
//
// If :attr:`input` has :math:`n` dimensions, then the resulting indices tensor
// :attr:`out` is of size :math:`(z \times n)`, where :math:`z` is the total number of
// non-zero elements in the :attr:`input` tensor.
//
// .. note::
//     This function is similar to NumPy's `argwhere`.
//
//     When :attr:`input` is on CUDA, this function causes host-device synchronization.
//
// Args:
//     {input}
//
// Example::
//
//     >>> t = torch.tensor([1, 0, 1])
//     >>> torch.argwhere(t)
//     tensor([[0],
//             [2]])
//     >>> t = torch.tensor([[1, 0, 1], [0, 1, 1]])
//     >>> torch.argwhere(t)
//     tensor([[0, 0],
//             [0, 2],
//             [1, 1],
//             [1, 2]])
//
//
//go:linkname Argwhere py.argwhere
func Argwhere(input *py.Object) *py.Object
//
// as_strided(input, size, stride, storage_offset=None) -> Tensor
//
// Create a view of an existing `torch.Tensor` :attr:`input` with specified
// :attr:`size`, :attr:`stride` and :attr:`storage_offset`.
//
// .. warning::
//     Prefer using other view functions, like :meth:`torch.Tensor.expand`,
//     to setting a view's strides manually with `as_strided`, as this
//     function's behavior depends on the implementation of a tensor's storage.
//     The constructed view of the storage must only refer to elements within
//     the storage or a runtime error will be thrown, and if the view is
//     "overlapped" (with multiple indices referring to the same element in
//     memory) its behavior is undefined.
//
// Args:
//     input (Tensor): the input tensor.
//     size (tuple or ints): the shape of the output tensor
//     stride (tuple or ints): the stride of the output tensor
//     storage_offset (int, optional): the offset in the underlying storage of the output tensor.
//         If ``None``, the storage_offset of the output tensor will match the input tensor.
//
// Example::
//
//     >>> x = torch.randn(3, 3)
//     >>> x
//     tensor([[ 0.9039,  0.6291,  1.0795],
//             [ 0.1586,  2.1939, -0.4900],
//             [-0.1909, -0.7503,  1.9355]])
//     >>> t = torch.as_strided(x, (2, 2), (1, 2))
//     >>> t
//     tensor([[0.9039, 1.0795],
//             [0.6291, 0.1586]])
//     >>> t = torch.as_strided(x, (2, 2), (1, 2), 1)
//     tensor([[0.6291, 0.1586],
//             [1.0795, 2.1939]])
//
//
//go:linkname AsStrided py.as_strided
func AsStrided(input *py.Object, size *py.Object, stride *py.Object, storageOffset *py.Object) *py.Object
// None
//
//go:linkname AsStrided_ py.as_strided_
func AsStrided_(__llgo_va_list ...interface{}) *py.Object
//
// Performs the same operation as :func:`torch.as_strided`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname AsStridedCopy py.as_strided_copy
func AsStridedCopy(__llgo_va_list ...interface{}) *py.Object
//
// as_strided_scatter(input, src, size, stride, storage_offset=None) -> Tensor
//
// Embeds the values of the :attr:`src` tensor into :attr:`input` along
// the elements corresponding to the result of calling
// input.as_strided(size, stride, storage_offset).
//
// This function returns a tensor with fresh storage; it does not
// return a view.
//
// Args:
//     input (Tensor): the input tensor.
//     size (tuple or ints): the shape of the output tensor
//     stride (tuple or ints): the stride of the output tensor
//     storage_offset (int, optional): the offset in the underlying storage of the output tensor
//
// .. note::
//
//     :attr:`src` must be of the proper size in order to be embedded
//     into :attr:`input`. Specifically, it should have the same shape as
//     `torch.as_strided(input, size, stride, storage_offset)`
//
// Example::
//
//     >>> a = torch.arange(4).reshape(2, 2) + 1
//     >>> a
//     tensor([[1, 2],
//             [3, 4]])
//     >>> b = torch.zeros(3, 3)
//     >>> b
//     tensor([[0., 0., 0.],
//             [0., 0., 0.],
//             [0., 0., 0.]])
//     >>> torch.as_strided_scatter(b, a, (2, 2), (1, 2))
//     tensor([[1., 3., 2.],
//             [4., 0., 0.],
//             [0., 0., 0.]])
//
//
//
//go:linkname AsStridedScatter py.as_strided_scatter
func AsStridedScatter(input *py.Object, src *py.Object, size *py.Object, stride *py.Object, storageOffset *py.Object) *py.Object
//
// as_tensor(data, dtype=None, device=None) -> Tensor
//
// Converts :attr:`data` into a tensor, sharing data and preserving autograd
// history if possible.
//
// If :attr:`data` is already a tensor with the requested dtype and device
// then :attr:`data` itself is returned, but if :attr:`data` is a
// tensor with a different dtype or device then it's copied as if using
// `data.to(dtype=dtype, device=device)`.
//
// If :attr:`data` is a NumPy array (an ndarray) with the same dtype and device then a
// tensor is constructed using :func:`torch.from_numpy`.
//
// .. seealso::
//
//     :func:`torch.tensor` never shares its data and creates a new "leaf tensor" (see :doc:`/notes/autograd`).
//
//
// Args:
//     data (array_like): Initial data for the tensor. Can be a list, tuple,
//         NumPy ``ndarray``, scalar, and other types.
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, infers data type from :attr:`data`.
//     device (:class:`torch.device`, optional): the device of the constructed tensor. If None and data is a tensor
//         then the device of data is used. If None and data is not a tensor then
//         the result tensor is constructed on the current device.
//
//
// Example::
//
//     >>> a = numpy.array([1, 2, 3])
//     >>> t = torch.as_tensor(a)
//     >>> t
//     tensor([ 1,  2,  3])
//     >>> t[0] = -1
//     >>> a
//     array([-1,  2,  3])
//
//     >>> a = numpy.array([1, 2, 3])
//     >>> t = torch.as_tensor(a, device=torch.device('cuda'))
//     >>> t
//     tensor([ 1,  2,  3])
//     >>> t[0] = -1
//     >>> a
//     array([1,  2,  3])
//
//
//go:linkname AsTensor py.as_tensor
func AsTensor(data *py.Object, dtype *py.Object, device *py.Object) *py.Object
//
// asarray(obj, *, dtype=None, device=None, copy=None, requires_grad=False) -> Tensor
//
// Converts :attr:`obj` to a tensor.
//
// :attr:`obj` can be one of:
//
// 1. a tensor
// 2. a NumPy array or a NumPy scalar
// 3. a DLPack capsule
// 4. an object that implements Python's buffer protocol
// 5. a scalar
// 6. a sequence of scalars
//
// When :attr:`obj` is a tensor, NumPy array, or DLPack capsule the returned tensor will,
// by default, not require a gradient, have the same datatype as :attr:`obj`, be on the
// same device, and share memory with it. These properties can be controlled with the
// :attr:`dtype`, :attr:`device`, :attr:`copy`, and :attr:`requires_grad` keyword arguments.
// If the returned tensor is of a different datatype, on a different device, or a copy is
// requested then it will not share its memory with :attr:`obj`. If :attr:`requires_grad`
// is ``True`` then the returned tensor will require a gradient, and if :attr:`obj` is
// also a tensor with an autograd history then the returned tensor will have the same history.
//
// When :attr:`obj` is not a tensor, NumPy array, or DLPack capsule but implements Python's
// buffer protocol then the buffer is interpreted as an array of bytes grouped according to
// the size of the datatype passed to the :attr:`dtype` keyword argument. (If no datatype is
// passed then the default floating point datatype is used, instead.) The returned tensor
// will have the specified datatype (or default floating point datatype if none is specified)
// and, by default, be on the CPU device and share memory with the buffer.
//
// When :attr:`obj` is a NumPy scalar, the returned tensor will be a 0-dimensional tensor on
// the CPU and that doesn't share its memory (i.e. ``copy=True``). By default datatype will
// be the PyTorch datatype corresponding to the NumPy's scalar's datatype.
//
// When :attr:`obj` is none of the above but a scalar, or a sequence of scalars then the
// returned tensor will, by default, infer its datatype from the scalar values, be on the
// current default device, and not share its memory.
//
// .. seealso::
//
//     :func:`torch.tensor` creates a tensor that always copies the data from the input object.
//     :func:`torch.from_numpy` creates a tensor that always shares memory from NumPy arrays.
//     :func:`torch.frombuffer` creates a tensor that always shares memory from objects that
//     implement the buffer protocol.
//     :func:`torch.from_dlpack` creates a tensor that always shares memory from
//     DLPack capsules.
//
// Args:
//     obj (object): a tensor, NumPy array, DLPack Capsule, object that implements Python's
//            buffer protocol, scalar, or sequence of scalars.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the datatype of the returned tensor.
//            Default: ``None``, which causes the datatype of the returned tensor to be
//            inferred from :attr:`obj`.
//     copy (bool, optional): controls whether the returned tensor shares memory with :attr:`obj`.
//            Default: ``None``, which causes the returned tensor to share memory with :attr:`obj`
//            whenever possible. If ``True`` then the returned tensor does not share its memory.
//            If ``False`` then the returned tensor shares its memory with :attr:`obj` and an
//            error is thrown if it cannot.
//     device (:class:`torch.device`, optional): the device of the returned tensor.
//            Default: ``None``, which causes the device of :attr:`obj` to be used. Or, if
//            :attr:`obj` is a Python sequence, the current default device will be used.
//     requires_grad (bool, optional): whether the returned tensor requires grad.
//            Default: ``False``, which causes the returned tensor not to require a gradient.
//            If ``True``, then the returned tensor will require a gradient, and if :attr:`obj`
//            is also a tensor with an autograd history then the returned tensor will have
//            the same history.
//
// Example::
//
//     >>> a = torch.tensor([1, 2, 3])
//     >>> # Shares memory with tensor 'a'
//     >>> b = torch.asarray(a)
//     >>> a.data_ptr() == b.data_ptr()
//     True
//     >>> # Forces memory copy
//     >>> c = torch.asarray(a, copy=True)
//     >>> a.data_ptr() == c.data_ptr()
//     False
//
//     >>> a = torch.tensor([1., 2., 3.], requires_grad=True)
//     >>> b = a + 2
//     >>> b
//     tensor([3., 4., 5.], grad_fn=<AddBackward0>)
//     >>> # Shares memory with tensor 'b', with no grad
//     >>> c = torch.asarray(b)
//     >>> c
//     tensor([3., 4., 5.])
//     >>> # Shares memory with tensor 'b', retaining autograd history
//     >>> d = torch.asarray(b, requires_grad=True)
//     >>> d
//     tensor([3., 4., 5.], grad_fn=<AddBackward0>)
//
//     >>> array = numpy.array([1, 2, 3])
//     >>> # Shares memory with array 'array'
//     >>> t1 = torch.asarray(array)
//     >>> array.__array_interface__['data'][0] == t1.data_ptr()
//     True
//     >>> # Copies memory due to dtype mismatch
//     >>> t2 = torch.asarray(array, dtype=torch.float32)
//     >>> array.__array_interface__['data'][0] == t2.data_ptr()
//     False
//
//     >>> scalar = numpy.float64(0.5)
//     >>> torch.asarray(scalar)
//     tensor(0.5000, dtype=torch.float64)
//
//
//go:linkname Asarray py.asarray
func Asarray(obj *py.Object) *py.Object
//
// asin(input, *, out=None) -> Tensor
//
// Returns a new tensor with the arcsine of the elements of :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \sin^{-1}(\text{input}_{i})
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([-0.5962,  1.4985, -0.4396,  1.4525])
//     >>> torch.asin(a)
//     tensor([-0.6387,     nan, -0.4552,     nan])
//
//
//go:linkname Asin py.asin
func Asin(input *py.Object) *py.Object
// None
//
//go:linkname Asin_ py.asin_
func Asin_(__llgo_va_list ...interface{}) *py.Object
//
// asinh(input, *, out=None) -> Tensor
//
// Returns a new tensor with the inverse hyperbolic sine of the elements of :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \sinh^{-1}(\text{input}_{i})
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword arguments:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([ 0.1606, -1.4267, -1.0899, -1.0250 ])
//     >>> torch.asinh(a)
//     tensor([ 0.1599, -1.1534, -0.9435, -0.8990 ])
//
//
//go:linkname Asinh py.asinh
func Asinh(input *py.Object) *py.Object
// None
//
//go:linkname Asinh_ py.asinh_
func Asinh_(__llgo_va_list ...interface{}) *py.Object
//
// atan(input, *, out=None) -> Tensor
//
// Returns a new tensor with the arctangent of the elements of :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \tan^{-1}(\text{input}_{i})
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([ 0.2341,  0.2539, -0.6256, -0.6448])
//     >>> torch.atan(a)
//     tensor([ 0.2299,  0.2487, -0.5591, -0.5727])
//
//
//go:linkname Atan py.atan
func Atan(input *py.Object) *py.Object
//
// atan2(input, other, *, out=None) -> Tensor
//
// Element-wise arctangent of :math:`\text{input}_{i} / \text{other}_{i}`
// with consideration of the quadrant. Returns a new tensor with the signed angles
// in radians between vector :math:`(\text{other}_{i}, \text{input}_{i})`
// and vector :math:`(1, 0)`. (Note that :math:`\text{other}_{i}`, the second
// parameter, is the x-coordinate, while :math:`\text{input}_{i}`, the first
// parameter, is the y-coordinate.)
//
// The shapes of ``input`` and ``other`` must be
// :ref:`broadcastable <broadcasting-semantics>`.
//
// Args:
//     input (Tensor): the first input tensor
//     other (Tensor): the second input tensor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([ 0.9041,  0.0196, -0.3108, -2.4423])
//     >>> torch.atan2(a, torch.randn(4))
//     tensor([ 0.9833,  0.0811, -1.9743, -1.4151])
//
//
//go:linkname Atan2 py.atan2
func Atan2(input *py.Object, other *py.Object) *py.Object
// None
//
//go:linkname Atan_ py.atan_
func Atan_(__llgo_va_list ...interface{}) *py.Object
//
// atanh(input, *, out=None) -> Tensor
//
// Returns a new tensor with the inverse hyperbolic tangent of the elements of :attr:`input`.
//
// Note:
//     The domain of the inverse hyperbolic tangent is `(-1, 1)` and values outside this range
//     will be mapped to ``NaN``, except for the values `1` and `-1` for which the output is
//     mapped to `+/-INF` respectively.
//
// .. math::
//     \text{out}_{i} = \tanh^{-1}(\text{input}_{i})
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword arguments:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4).uniform_(-1, 1)
//     >>> a
//     tensor([ -0.9385, 0.2968, -0.8591, -0.1871 ])
//     >>> torch.atanh(a)
//     tensor([ -1.7253, 0.3060, -1.2899, -0.1893 ])
//
//
//go:linkname Atanh py.atanh
func Atanh(input *py.Object) *py.Object
// None
//
//go:linkname Atanh_ py.atanh_
func Atanh_(__llgo_va_list ...interface{}) *py.Object
//
//     Returns a 1-dimensional view of each input tensor with zero dimensions.
//     Input tensors with one or more dimensions are returned as-is.
//
//     Args:
//         input (Tensor or list of Tensors)
//
//     Returns:
//         output (Tensor or tuple of Tensors)
//
//     Example::
//
//         >>> x = torch.arange(2)
//         >>> x
//         tensor([0, 1])
//         >>> torch.atleast_1d(x)
//         tensor([0, 1])
//         >>> x = torch.tensor(1.)
//         >>> x
//         tensor(1.)
//         >>> torch.atleast_1d(x)
//         tensor([1.])
//         >>> x = torch.tensor(0.5)
//         >>> y = torch.tensor(1.)
//         >>> torch.atleast_1d((x, y))
//         (tensor([0.5000]), tensor([1.]))
//
//
//go:linkname Atleast1d py.atleast_1d
func Atleast1d(__llgo_va_list ...interface{}) *py.Object
//
//     Returns a 2-dimensional view of each input tensor with zero dimensions.
//     Input tensors with two or more dimensions are returned as-is.
//
//     Args:
//         input (Tensor or list of Tensors)
//
//     Returns:
//         output (Tensor or tuple of Tensors)
//
//     Example::
//
//         >>> x = torch.tensor(1.)
//         >>> x
//         tensor(1.)
//         >>> torch.atleast_2d(x)
//         tensor([[1.]])
//         >>> x = torch.arange(4).view(2, 2)
//         >>> x
//         tensor([[0, 1],
//                 [2, 3]])
//         >>> torch.atleast_2d(x)
//         tensor([[0, 1],
//                 [2, 3]])
//         >>> x = torch.tensor(0.5)
//         >>> y = torch.tensor(1.)
//         >>> torch.atleast_2d((x, y))
//         (tensor([[0.5000]]), tensor([[1.]]))
//
//
//go:linkname Atleast2d py.atleast_2d
func Atleast2d(__llgo_va_list ...interface{}) *py.Object
//
//     Returns a 3-dimensional view of each input tensor with zero dimensions.
//     Input tensors with three or more dimensions are returned as-is.
//
//     Args:
//         input (Tensor or list of Tensors)
//
//     Returns:
//         output (Tensor or tuple of Tensors)
//
//     Example:
//
//         >>> x = torch.tensor(0.5)
//         >>> x
//         tensor(0.5000)
//         >>> torch.atleast_3d(x)
//         tensor([[[0.5000]]])
//         >>> y = torch.arange(4).view(2, 2)
//         >>> y
//         tensor([[0, 1],
//                 [2, 3]])
//         >>> torch.atleast_3d(y)
//         tensor([[[0],
//                  [1]],
//                 <BLANKLINE>
//                 [[2],
//                  [3]]])
//         >>> x = torch.tensor(1).view(1, 1, 1)
//         >>> x
//         tensor([[[1]]])
//         >>> torch.atleast_3d(x)
//         tensor([[[1]]])
//         >>> x = torch.tensor(0.5)
//         >>> y = torch.tensor(1.)
//         >>> torch.atleast_3d((x, y))
//         (tensor([[[0.5000]]]), tensor([[[1.]]]))
//
//
//go:linkname Atleast3d py.atleast_3d
func Atleast3d(__llgo_va_list ...interface{}) *py.Object
//
// avg_pool1d(input, kernel_size, stride=None, padding=0, ceil_mode=False, count_include_pad=True) -> Tensor
//
// Applies a 1D average pooling over an input signal composed of several
// input planes.
//
// See :class:`~torch.nn.AvgPool1d` for details and output shape.
//
// Args:
//     input: input tensor of shape :math:`(\text{minibatch} , \text{in\_channels} , iW)`
//     kernel_size: the size of the window. Can be a single number or a
//       tuple `(kW,)`
//     stride: the stride of the window. Can be a single number or a tuple
//       `(sW,)`. Default: :attr:`kernel_size`
//     padding: implicit zero paddings on both sides of the input. Can be a
//       single number or a tuple `(padW,)`. Default: 0
//     ceil_mode: when True, will use `ceil` instead of `floor` to compute the
//         output shape. Default: ``False``
//     count_include_pad: when True, will include the zero-padding in the
//         averaging calculation. Default: ``True``
//
// Examples::
//
//     >>> # pool of square window of size=3, stride=2
//     >>> input = torch.tensor([[[1, 2, 3, 4, 5, 6, 7]]], dtype=torch.float32)
//     >>> F.avg_pool1d(input, kernel_size=3, stride=2)
//     tensor([[[ 2.,  4.,  6.]]])
//
//
//
//go:linkname AvgPool1d py.avg_pool1d
func AvgPool1d(input *py.Object, kernelSize *py.Object, stride *py.Object, padding *py.Object, ceilMode *py.Object, countIncludePad *py.Object) *py.Object
//
// baddbmm(input, batch1, batch2, *, beta=1, alpha=1, out=None) -> Tensor
//
// Performs a batch matrix-matrix product of matrices in :attr:`batch1`
// and :attr:`batch2`.
// :attr:`input` is added to the final result.
//
// :attr:`batch1` and :attr:`batch2` must be 3-D tensors each containing the same
// number of matrices.
//
// If :attr:`batch1` is a :math:`(b \times n \times m)` tensor, :attr:`batch2` is a
// :math:`(b \times m \times p)` tensor, then :attr:`input` must be
// :ref:`broadcastable <broadcasting-semantics>` with a
// :math:`(b \times n \times p)` tensor and :attr:`out` will be a
// :math:`(b \times n \times p)` tensor. Both :attr:`alpha` and :attr:`beta` mean the
// same as the scaling factors used in :meth:`torch.addbmm`.
//
// .. math::
//     \text{out}_i = \beta\ \text{input}_i + \alpha\ (\text{batch1}_i \mathbin{@} \text{batch2}_i)
//
// If :attr:`beta` is 0, then :attr:`input` will be ignored, and `nan` and `inf` in
// it will not be propagated.
//
// For inputs of type `FloatTensor` or `DoubleTensor`, arguments :attr:`beta` and
// :attr:`alpha` must be real numbers, otherwise they should be integers.
//
// This operator supports :ref:`TensorFloat32<tf32_on_ampere>`.
//
// On certain ROCm devices, when using float16 inputs this module will use :ref:`different precision<fp16_on_mi200>` for backward.
//
// Args:
//     input (Tensor): the tensor to be added
//     batch1 (Tensor): the first batch of matrices to be multiplied
//     batch2 (Tensor): the second batch of matrices to be multiplied
//
// Keyword args:
//     beta (Number, optional): multiplier for :attr:`input` (:math:`\beta`)
//     alpha (Number, optional): multiplier for :math:`\text{batch1} \mathbin{@} \text{batch2}` (:math:`\alpha`)
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> M = torch.randn(10, 3, 5)
//     >>> batch1 = torch.randn(10, 3, 4)
//     >>> batch2 = torch.randn(10, 4, 5)
//     >>> torch.baddbmm(M, batch1, batch2).size()
//     torch.Size([10, 3, 5])
//
//
//go:linkname Baddbmm py.baddbmm
func Baddbmm(input *py.Object, batch1 *py.Object, batch2 *py.Object) *py.Object
//
// bartlett_window(window_length, periodic=True, *, dtype=None, layout=torch.strided, device=None, requires_grad=False) -> Tensor
//
// Bartlett window function.
//
// .. math::
//     w[n] = 1 - \left| \frac{2n}{N-1} - 1 \right| = \begin{cases}
//         \frac{2n}{N - 1} & \text{if } 0 \leq n \leq \frac{N - 1}{2} \\
//         2 - \frac{2n}{N - 1} & \text{if } \frac{N - 1}{2} < n < N \\
//     \end{cases},
//
// where :math:`N` is the full window size.
//
// The input :attr:`window_length` is a positive integer controlling the
// returned window size. :attr:`periodic` flag determines whether the returned
// window trims off the last duplicate value from the symmetric window and is
// ready to be used as a periodic window with functions like
// :meth:`torch.stft`. Therefore, if :attr:`periodic` is true, the :math:`N` in
// above formula is in fact :math:`\text{window\_length} + 1`. Also, we always have
// ``torch.bartlett_window(L, periodic=True)`` equal to
// ``torch.bartlett_window(L + 1, periodic=False)[:-1])``.
//
// .. note::
//     If :attr:`window_length` :math:`=1`, the returned window contains a single value 1.
//
// Arguments:
//     window_length (int): the size of returned window
//     periodic (bool, optional): If True, returns a window to be used as periodic
//         function. If False, return a symmetric window.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`). Only floating point types are supported.
//     layout (:class:`torch.layout`, optional): the desired layout of returned window tensor. Only
//           ``torch.strided`` (dense layout) is supported.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
// Returns:
//     Tensor: A 1-D tensor of size :math:`(\text{window\_length},)` containing the window
//
//
//
//go:linkname BartlettWindow py.bartlett_window
func BartlettWindow(windowLength *py.Object, periodic *py.Object) *py.Object
// None
//
//go:linkname BatchNorm py.batch_norm
func BatchNorm(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname BatchNormBackwardElemt py.batch_norm_backward_elemt
func BatchNormBackwardElemt(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname BatchNormBackwardReduce py.batch_norm_backward_reduce
func BatchNormBackwardReduce(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname BatchNormElemt py.batch_norm_elemt
func BatchNormElemt(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname BatchNormGatherStats py.batch_norm_gather_stats
func BatchNormGatherStats(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname BatchNormGatherStatsWithCounts py.batch_norm_gather_stats_with_counts
func BatchNormGatherStatsWithCounts(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname BatchNormStats py.batch_norm_stats
func BatchNormStats(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname BatchNormUpdateStats py.batch_norm_update_stats
func BatchNormUpdateStats(__llgo_va_list ...interface{}) *py.Object
//
// bernoulli(input, *, generator=None, out=None) -> Tensor
//
// Draws binary random numbers (0 or 1) from a Bernoulli distribution.
//
// The :attr:`input` tensor should be a tensor containing probabilities
// to be used for drawing the binary random number.
// Hence, all values in :attr:`input` have to be in the range:
// :math:`0 \leq \text{input}_i \leq 1`.
//
// The :math:`\text{i}^{th}` element of the output tensor will draw a
// value :math:`1` according to the :math:`\text{i}^{th}` probability value given
// in :attr:`input`.
//
// .. math::
//     \text{out}_{i} \sim \mathrm{Bernoulli}(p = \text{input}_{i})
//
// The returned :attr:`out` tensor only has values 0 or 1 and is of the same
// shape as :attr:`input`.
//
// :attr:`out` can have integral ``dtype``, but :attr:`input` must have floating
// point ``dtype``.
//
// Args:
//     input (Tensor): the input tensor of probability values for the Bernoulli distribution
//
// Keyword args:
//     generator (:class:`torch.Generator`, optional): a pseudorandom number generator for sampling
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.empty(3, 3).uniform_(0, 1)  # generate a uniform random matrix with range [0, 1]
//     >>> a
//     tensor([[ 0.1737,  0.0950,  0.3609],
//             [ 0.7148,  0.0289,  0.2676],
//             [ 0.9456,  0.8937,  0.7202]])
//     >>> torch.bernoulli(a)
//     tensor([[ 1.,  0.,  0.],
//             [ 0.,  0.,  0.],
//             [ 1.,  1.,  1.]])
//
//     >>> a = torch.ones(3, 3) # probability of drawing "1" is 1
//     >>> torch.bernoulli(a)
//     tensor([[ 1.,  1.,  1.],
//             [ 1.,  1.,  1.],
//             [ 1.,  1.,  1.]])
//     >>> a = torch.zeros(3, 3) # probability of drawing "1" is 0
//     >>> torch.bernoulli(a)
//     tensor([[ 0.,  0.,  0.],
//             [ 0.,  0.,  0.],
//             [ 0.,  0.,  0.]])
//
//
//go:linkname Bernoulli py.bernoulli
func Bernoulli(input *py.Object) *py.Object
//
// bilinear(input1, input2, weight, bias=None) -> Tensor
//
// Applies a bilinear transformation to the incoming data:
// :math:`y = x_1^T A x_2 + b`
//
// Shape:
//
//     - input1: :math:`(N, *, H_{in1})` where :math:`H_{in1}=\text{in1\_features}`
//       and :math:`*` means any number of additional dimensions.
//       All but the last dimension of the inputs should be the same.
//     - input2: :math:`(N, *, H_{in2})` where :math:`H_{in2}=\text{in2\_features}`
//     - weight: :math:`(\text{out\_features}, \text{in1\_features},
//       \text{in2\_features})`
//     - bias: :math:`(\text{out\_features})`
//     - output: :math:`(N, *, H_{out})` where :math:`H_{out}=\text{out\_features}`
//       and all but the last dimension are the same shape as the input.
//
//
//go:linkname Bilinear py.bilinear
func Bilinear(input1 *py.Object, input2 *py.Object, weight *py.Object, bias *py.Object) *py.Object
// None
//
//go:linkname BinaryCrossEntropyWithLogits py.binary_cross_entropy_with_logits
func BinaryCrossEntropyWithLogits(__llgo_va_list ...interface{}) *py.Object
//
// bincount(input, weights=None, minlength=0) -> Tensor
//
// Count the frequency of each value in an array of non-negative ints.
//
// The number of bins (size 1) is one larger than the largest value in
// :attr:`input` unless :attr:`input` is empty, in which case the result is a
// tensor of size 0. If :attr:`minlength` is specified, the number of bins is at least
// :attr:`minlength` and if :attr:`input` is empty, then the result is tensor of size
// :attr:`minlength` filled with zeros. If ``n`` is the value at position ``i``,
// ``out[n] += weights[i]`` if :attr:`weights` is specified else
// ``out[n] += 1``.
//
// Note:
//     This operation may produce nondeterministic gradients when given tensors on a CUDA device. See :doc:`/notes/randomness` for more information.
//
// Arguments:
//     input (Tensor): 1-d int tensor
//     weights (Tensor): optional, weight for each value in the input tensor.
//         Should be of same size as input tensor.
//     minlength (int): optional, minimum number of bins. Should be non-negative.
//
// Returns:
//     output (Tensor): a tensor of shape ``Size([max(input) + 1])`` if
//     :attr:`input` is non-empty, else ``Size(0)``
//
// Example::
//
//     >>> input = torch.randint(0, 8, (5,), dtype=torch.int64)
//     >>> weights = torch.linspace(0, 1, steps=5)
//     >>> input, weights
//     (tensor([4, 3, 6, 3, 4]),
//      tensor([ 0.0000,  0.2500,  0.5000,  0.7500,  1.0000])
//
//     >>> torch.bincount(input)
//     tensor([0, 0, 0, 2, 2, 0, 1])
//
//     >>> input.bincount(weights)
//     tensor([0.0000, 0.0000, 0.0000, 1.0000, 1.0000, 0.0000, 0.5000])
//
//
//go:linkname Bincount py.bincount
func Bincount(input *py.Object, weights *py.Object, minlength *py.Object) *py.Object
// None
//
//go:linkname Binomial py.binomial
func Binomial(__llgo_va_list ...interface{}) *py.Object
//
// bitwise_and(input, other, *, out=None) -> Tensor
//
// Computes the bitwise AND of :attr:`input` and :attr:`other`. The input tensor must be of
// integral or Boolean types. For bool tensors, it computes the logical AND.
//
// Args:
//     input: the first input tensor
//     other: the second input tensor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.bitwise_and(torch.tensor([-1, -2, 3], dtype=torch.int8), torch.tensor([1, 0, 3], dtype=torch.int8))
//     tensor([1, 0,  3], dtype=torch.int8)
//     >>> torch.bitwise_and(torch.tensor([True, True, False]), torch.tensor([False, True, False]))
//     tensor([ False, True, False])
//
//
//go:linkname BitwiseAnd py.bitwise_and
func BitwiseAnd(input *py.Object, other *py.Object) *py.Object
//
// bitwise_left_shift(input, other, *, out=None) -> Tensor
//
// Computes the left arithmetic shift of :attr:`input` by :attr:`other` bits.
// The input tensor must be of integral type. This operator supports
// :ref:`broadcasting to a common shape <broadcasting-semantics>` and
// :ref:`type promotion <type-promotion-doc>`.
//
// The operation applied is:
//
// .. math::
//     \text{out}_i = \text{input}_i << \text{other}_i
//
// Args:
//     input (Tensor or Scalar): the first input tensor
//     other (Tensor or Scalar): the second input tensor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.bitwise_left_shift(torch.tensor([-1, -2, 3], dtype=torch.int8), torch.tensor([1, 0, 3], dtype=torch.int8))
//     tensor([-2, -2, 24], dtype=torch.int8)
//
//
//go:linkname BitwiseLeftShift py.bitwise_left_shift
func BitwiseLeftShift(input *py.Object, other *py.Object) *py.Object
//
// bitwise_not(input, *, out=None) -> Tensor
//
// Computes the bitwise NOT of the given input tensor. The input tensor must be of
// integral or Boolean types. For bool tensors, it computes the logical NOT.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.bitwise_not(torch.tensor([-1, -2, 3], dtype=torch.int8))
//     tensor([ 0,  1, -4], dtype=torch.int8)
//
//
//go:linkname BitwiseNot py.bitwise_not
func BitwiseNot(input *py.Object) *py.Object
//
// bitwise_or(input, other, *, out=None) -> Tensor
//
// Computes the bitwise OR of :attr:`input` and :attr:`other`. The input tensor must be of
// integral or Boolean types. For bool tensors, it computes the logical OR.
//
// Args:
//     input: the first input tensor
//     other: the second input tensor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.bitwise_or(torch.tensor([-1, -2, 3], dtype=torch.int8), torch.tensor([1, 0, 3], dtype=torch.int8))
//     tensor([-1, -2,  3], dtype=torch.int8)
//     >>> torch.bitwise_or(torch.tensor([True, True, False]), torch.tensor([False, True, False]))
//     tensor([ True, True, False])
//
//
//go:linkname BitwiseOr py.bitwise_or
func BitwiseOr(input *py.Object, other *py.Object) *py.Object
//
// bitwise_right_shift(input, other, *, out=None) -> Tensor
//
// Computes the right arithmetic shift of :attr:`input` by :attr:`other` bits.
// The input tensor must be of integral type. This operator supports
// :ref:`broadcasting to a common shape <broadcasting-semantics>` and
// :ref:`type promotion <type-promotion-doc>`.
// In any case, if the value of the right operand is negative or is greater
// or equal to the number of bits in the promoted left operand, the behavior is undefined.
//
// The operation applied is:
//
// .. math::
//     \text{out}_i = \text{input}_i >> \text{other}_i
//
// Args:
//     input (Tensor or Scalar): the first input tensor
//     other (Tensor or Scalar): the second input tensor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.bitwise_right_shift(torch.tensor([-2, -7, 31], dtype=torch.int8), torch.tensor([1, 0, 3], dtype=torch.int8))
//     tensor([-1, -7,  3], dtype=torch.int8)
//
//
//go:linkname BitwiseRightShift py.bitwise_right_shift
func BitwiseRightShift(input *py.Object, other *py.Object) *py.Object
//
// bitwise_xor(input, other, *, out=None) -> Tensor
//
// Computes the bitwise XOR of :attr:`input` and :attr:`other`. The input tensor must be of
// integral or Boolean types. For bool tensors, it computes the logical XOR.
//
// Args:
//     input: the first input tensor
//     other: the second input tensor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.bitwise_xor(torch.tensor([-1, -2, 3], dtype=torch.int8), torch.tensor([1, 0, 3], dtype=torch.int8))
//     tensor([-2, -2,  0], dtype=torch.int8)
//     >>> torch.bitwise_xor(torch.tensor([True, True, False]), torch.tensor([False, True, False]))
//     tensor([ True, False, False])
//
//
//go:linkname BitwiseXor py.bitwise_xor
func BitwiseXor(input *py.Object, other *py.Object) *py.Object
//
// blackman_window(window_length, periodic=True, *, dtype=None, layout=torch.strided, device=None, requires_grad=False) -> Tensor
//
// Blackman window function.
//
// .. math::
//     w[n] = 0.42 - 0.5 \cos \left( \frac{2 \pi n}{N - 1} \right) + 0.08 \cos \left( \frac{4 \pi n}{N - 1} \right)
//
// where :math:`N` is the full window size.
//
// The input :attr:`window_length` is a positive integer controlling the
// returned window size. :attr:`periodic` flag determines whether the returned
// window trims off the last duplicate value from the symmetric window and is
// ready to be used as a periodic window with functions like
// :meth:`torch.stft`. Therefore, if :attr:`periodic` is true, the :math:`N` in
// above formula is in fact :math:`\text{window\_length} + 1`. Also, we always have
// ``torch.blackman_window(L, periodic=True)`` equal to
// ``torch.blackman_window(L + 1, periodic=False)[:-1])``.
//
// .. note::
//     If :attr:`window_length` :math:`=1`, the returned window contains a single value 1.
//
// Arguments:
//     window_length (int): the size of returned window
//     periodic (bool, optional): If True, returns a window to be used as periodic
//         function. If False, return a symmetric window.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`). Only floating point types are supported.
//     layout (:class:`torch.layout`, optional): the desired layout of returned window tensor. Only
//           ``torch.strided`` (dense layout) is supported.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
// Returns:
//     Tensor: A 1-D tensor of size :math:`(\text{window\_length},)` containing the window
//
//
//
//go:linkname BlackmanWindow py.blackman_window
func BlackmanWindow(windowLength *py.Object, periodic *py.Object) *py.Object
// Create a block diagonal matrix from provided tensors.
//
//     Args:
//         *tensors: One or more tensors with 0, 1, or 2 dimensions.
//
//     Returns:
//         Tensor: A 2 dimensional tensor with all the input tensors arranged in
//         order such that their upper left and lower right corners are
//         diagonally adjacent. All other elements are set to 0.
//
//     Example::
//
//         >>> import torch
//         >>> A = torch.tensor([[0, 1], [1, 0]])
//         >>> B = torch.tensor([[3, 4, 5], [6, 7, 8]])
//         >>> C = torch.tensor(7)
//         >>> D = torch.tensor([1, 2, 3])
//         >>> E = torch.tensor([[4], [5], [6]])
//         >>> torch.block_diag(A, B, C, D, E)
//         tensor([[0, 1, 0, 0, 0, 0, 0, 0, 0, 0],
//                 [1, 0, 0, 0, 0, 0, 0, 0, 0, 0],
//                 [0, 0, 3, 4, 5, 0, 0, 0, 0, 0],
//                 [0, 0, 6, 7, 8, 0, 0, 0, 0, 0],
//                 [0, 0, 0, 0, 0, 7, 0, 0, 0, 0],
//                 [0, 0, 0, 0, 0, 0, 1, 2, 3, 0],
//                 [0, 0, 0, 0, 0, 0, 0, 0, 0, 4],
//                 [0, 0, 0, 0, 0, 0, 0, 0, 0, 5],
//                 [0, 0, 0, 0, 0, 0, 0, 0, 0, 6]])
//
//
//go:linkname BlockDiag py.block_diag
func BlockDiag(__llgo_va_list ...interface{}) *py.Object
//
// bmm(input, mat2, *, out=None) -> Tensor
//
// Performs a batch matrix-matrix product of matrices stored in :attr:`input`
// and :attr:`mat2`.
//
// :attr:`input` and :attr:`mat2` must be 3-D tensors each containing
// the same number of matrices.
//
// If :attr:`input` is a :math:`(b \times n \times m)` tensor, :attr:`mat2` is a
// :math:`(b \times m \times p)` tensor, :attr:`out` will be a
// :math:`(b \times n \times p)` tensor.
//
// .. math::
//     \text{out}_i = \text{input}_i \mathbin{@} \text{mat2}_i
//
// This operator supports :ref:`TensorFloat32<tf32_on_ampere>`.
//
// On certain ROCm devices, when using float16 inputs this module will use :ref:`different precision<fp16_on_mi200>` for backward.
//
// .. note:: This function does not :ref:`broadcast <broadcasting-semantics>`.
//           For broadcasting matrix products, see :func:`torch.matmul`.
//
// Args:
//     input (Tensor): the first batch of matrices to be multiplied
//     mat2 (Tensor): the second batch of matrices to be multiplied
//
// Keyword Args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> input = torch.randn(10, 3, 4)
//     >>> mat2 = torch.randn(10, 4, 5)
//     >>> res = torch.bmm(input, mat2)
//     >>> res.size()
//     torch.Size([10, 3, 5])
//
//
//go:linkname Bmm py.bmm
func Bmm(input *py.Object, mat2 *py.Object) *py.Object
// broadcast_tensors(*tensors) -> List of Tensors
//
//     Broadcasts the given tensors according to :ref:`broadcasting-semantics`.
//
//     Args:
//         *tensors: any number of tensors of the same type
//
//     .. warning::
//
//         More than one element of a broadcasted tensor may refer to a single
//         memory location. As a result, in-place operations (especially ones that
//         are vectorized) may result in incorrect behavior. If you need to write
//         to the tensors, please clone them first.
//
//     Example::
//
//         >>> x = torch.arange(3).view(1, 3)
//         >>> y = torch.arange(2).view(2, 1)
//         >>> a, b = torch.broadcast_tensors(x, y)
//         >>> a.size()
//         torch.Size([2, 3])
//         >>> a
//         tensor([[0, 1, 2],
//                 [0, 1, 2]])
//
//
//go:linkname BroadcastTensors py.broadcast_tensors
func BroadcastTensors(__llgo_va_list ...interface{}) *py.Object
//
// broadcast_to(input, shape) -> Tensor
//
// Broadcasts :attr:`input` to the shape :attr:`\shape`.
// Equivalent to calling ``input.expand(shape)``. See :meth:`~Tensor.expand` for details.
//
// Args:
//     input (Tensor): the input tensor.
//     shape (list, tuple, or :class:`torch.Size`): the new shape.
//
// Example::
//
//     >>> x = torch.tensor([1, 2, 3])
//     >>> torch.broadcast_to(x, (3, 3))
//     tensor([[1, 2, 3],
//             [1, 2, 3],
//             [1, 2, 3]])
//
//
//go:linkname BroadcastTo py.broadcast_to
func BroadcastTo(input *py.Object, shape *py.Object) *py.Object
//
// bucketize(input, boundaries, *, out_int32=False, right=False, out=None) -> Tensor
//
// Returns the indices of the buckets to which each value in the :attr:`input` belongs, where the
// boundaries of the buckets are set by :attr:`boundaries`. Return a new tensor with the same size
// as :attr:`input`. If :attr:`right` is False (default), then the left boundary is open. Note that
// this behavior is opposite the behavior of
// `numpy.digitize <https://docs.scipy.org/doc/numpy/reference/generated/numpy.digitize.html>`_.
// More formally, the returned index satisfies the following rules:
//
// .. list-table::
//    :widths: 15 85
//    :header-rows: 1
//
//    * - :attr:`right`
//      - *returned index satisfies*
//    * - False
//      - ``boundaries[i-1] < input[m][n]...[l][x] <= boundaries[i]``
//    * - True
//      - ``boundaries[i-1] <= input[m][n]...[l][x] < boundaries[i]``
//
// Args:
//     input (Tensor or Scalar): N-D tensor or a Scalar containing the search value(s).
//     boundaries (Tensor): 1-D tensor, must contain a strictly increasing sequence, or the return value is undefined.
//
// Keyword args:
//     out_int32 (bool, optional): indicate the output data type. torch.int32 if True, torch.int64 otherwise.
//                                 Default value is False, i.e. default output data type is torch.int64.
//     right (bool, optional): if False, return the first suitable location that is found. If True, return the
//                             last such index. If no suitable index found, return 0 for non-numerical value
//                             (eg. nan, inf) or the size of :attr:`boundaries` (one pass the last index).
//                             In other words, if False, gets the lower bound index for each value in :attr:`input`
//                             from :attr:`boundaries`. If True, gets the upper bound index instead.
//                             Default value is False.
//     out (Tensor, optional): the output tensor, must be the same size as :attr:`input` if provided.
//
//
// Example::
//
//     >>> boundaries = torch.tensor([1, 3, 5, 7, 9])
//     >>> boundaries
//     tensor([1, 3, 5, 7, 9])
//     >>> v = torch.tensor([[3, 6, 9], [3, 6, 9]])
//     >>> v
//     tensor([[3, 6, 9],
//             [3, 6, 9]])
//     >>> torch.bucketize(v, boundaries)
//     tensor([[1, 3, 4],
//             [1, 3, 4]])
//     >>> torch.bucketize(v, boundaries, right=True)
//     tensor([[2, 3, 5],
//             [2, 3, 5]])
//
//
//go:linkname Bucketize py.bucketize
func Bucketize(input *py.Object, boundaries *py.Object) *py.Object
//
// can_cast(from, to) -> bool
//
// Determines if a type conversion is allowed under PyTorch casting rules
// described in the type promotion :ref:`documentation <type-promotion-doc>`.
//
// Args:
//     from (dtype): The original :class:`torch.dtype`.
//     to (dtype): The target :class:`torch.dtype`.
//
// Example::
//
//     >>> torch.can_cast(torch.double, torch.float)
//     True
//     >>> torch.can_cast(torch.float, torch.int)
//     False
//
//
//go:linkname CanCast py.can_cast
func CanCast(from *py.Object, to *py.Object) *py.Object
// Do cartesian product of the given sequence of tensors. The behavior is similar to
//     python's `itertools.product`.
//
//     Args:
//         *tensors: any number of 1 dimensional tensors.
//
//     Returns:
//         Tensor: A tensor equivalent to converting all the input tensors into lists,
//         do `itertools.product` on these lists, and finally convert the resulting list
//         into tensor.
//
//     Example::
//
//         >>> import itertools
//         >>> a = [1, 2, 3]
//         >>> b = [4, 5]
//         >>> list(itertools.product(a, b))
//         [(1, 4), (1, 5), (2, 4), (2, 5), (3, 4), (3, 5)]
//         >>> tensor_a = torch.tensor(a)
//         >>> tensor_b = torch.tensor(b)
//         >>> torch.cartesian_prod(tensor_a, tensor_b)
//         tensor([[1, 4],
//                 [1, 5],
//                 [2, 4],
//                 [2, 5],
//                 [3, 4],
//                 [3, 5]])
//
//
//go:linkname CartesianProd py.cartesian_prod
func CartesianProd(__llgo_va_list ...interface{}) *py.Object
//
// cat(tensors, dim=0, *, out=None) -> Tensor
//
// Concatenates the given sequence of :attr:`seq` tensors in the given dimension.
// All tensors must either have the same shape (except in the concatenating
// dimension) or be empty.
//
// :func:`torch.cat` can be seen as an inverse operation for :func:`torch.split`
// and :func:`torch.chunk`.
//
// :func:`torch.cat` can be best understood via examples.
//
// .. seealso::
//
//     :func:`torch.stack` concatenates the given sequence along a new dimension.
//
// Args:
//     tensors (sequence of Tensors): any python sequence of tensors of the same type.
//         Non-empty tensors provided must have the same shape, except in the
//         cat dimension.
//     dim (int, optional): the dimension over which the tensors are concatenated
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> x = torch.randn(2, 3)
//     >>> x
//     tensor([[ 0.6580, -1.0969, -0.4614],
//             [-0.1034, -0.5790,  0.1497]])
//     >>> torch.cat((x, x, x), 0)
//     tensor([[ 0.6580, -1.0969, -0.4614],
//             [-0.1034, -0.5790,  0.1497],
//             [ 0.6580, -1.0969, -0.4614],
//             [-0.1034, -0.5790,  0.1497],
//             [ 0.6580, -1.0969, -0.4614],
//             [-0.1034, -0.5790,  0.1497]])
//     >>> torch.cat((x, x, x), 1)
//     tensor([[ 0.6580, -1.0969, -0.4614,  0.6580, -1.0969, -0.4614,  0.6580,
//              -1.0969, -0.4614],
//             [-0.1034, -0.5790,  0.1497, -0.1034, -0.5790,  0.1497, -0.1034,
//              -0.5790,  0.1497]])
//
//
//go:linkname Cat py.cat
func Cat(tensors *py.Object, dim *py.Object) *py.Object
// None
//
//go:linkname CcolIndicesCopy py.ccol_indices_copy
func CcolIndicesCopy(__llgo_va_list ...interface{}) *py.Object
// Computes batched the p-norm distance between each pair of the two collections of row vectors.
//
//     Args:
//         x1 (Tensor): input tensor of shape :math:`B \times P \times M`.
//         x2 (Tensor): input tensor of shape :math:`B \times R \times M`.
//         p: p value for the p-norm distance to calculate between each vector pair
//             :math:`\in [0, \infty]`.
//         compute_mode:
//             'use_mm_for_euclid_dist_if_necessary' - will use matrix multiplication approach to calculate
//             euclidean distance (p = 2) if P > 25 or R > 25
//             'use_mm_for_euclid_dist' - will always use matrix multiplication approach to calculate
//             euclidean distance (p = 2)
//             'donot_use_mm_for_euclid_dist' - will never use matrix multiplication approach to calculate
//             euclidean distance (p = 2)
//             Default: use_mm_for_euclid_dist_if_necessary.
//
//     If x1 has shape :math:`B \times P \times M` and x2 has shape :math:`B \times R \times M` then the
//     output will have shape :math:`B \times P \times R`.
//
//     This function is equivalent to `scipy.spatial.distance.cdist(input,'minkowski', p=p)`
//     if :math:`p \in (0, \infty)`. When :math:`p = 0` it is equivalent to
//     `scipy.spatial.distance.cdist(input, 'hamming') * M`. When :math:`p = \infty`, the closest
//     scipy function is `scipy.spatial.distance.cdist(xn, lambda x, y: np.abs(x - y).max())`.
//
//     Example:
//
//         >>> a = torch.tensor([[0.9041,  0.0196], [-0.3108, -2.4423], [-0.4821,  1.059]])
//         >>> a
//         tensor([[ 0.9041,  0.0196],
//                 [-0.3108, -2.4423],
//                 [-0.4821,  1.0590]])
//         >>> b = torch.tensor([[-2.1763, -0.4713], [-0.6986,  1.3702]])
//         >>> b
//         tensor([[-2.1763, -0.4713],
//                 [-0.6986,  1.3702]])
//         >>> torch.cdist(a, b, p=2)
//         tensor([[3.1193, 2.0959],
//                 [2.7138, 3.8322],
//                 [2.2830, 0.3791]])
//
//
//go:linkname Cdist py.cdist
func Cdist(x1 *py.Object, x2 *py.Object, p *py.Object, computeMode *py.Object) *py.Object
//
// ceil(input, *, out=None) -> Tensor
//
// Returns a new tensor with the ceil of the elements of :attr:`input`,
// the smallest integer greater than or equal to each element.
//
// For integer inputs, follows the array-api convention of returning a
// copy of the input tensor.
//
// .. math::
//     \text{out}_{i} = \left\lceil \text{input}_{i} \right\rceil
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([-0.6341, -1.4208, -1.0900,  0.5826])
//     >>> torch.ceil(a)
//     tensor([-0., -1., -1.,  1.])
//
//
//go:linkname Ceil py.ceil
func Ceil(input *py.Object) *py.Object
// None
//
//go:linkname Ceil_ py.ceil_
func Ceil_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Celu py.celu
func Celu(__llgo_va_list ...interface{}) *py.Object
//
// celu_(input, alpha=1.) -> Tensor
//
// In-place version of :func:`~celu`.
//
//
//go:linkname Celu_ py.celu_
func Celu_(input *py.Object, alpha *py.Object) *py.Object
// Returns the matrix product of the :math:`N` 2-D tensors. This product is efficiently computed
//     using the matrix chain order algorithm which selects the order in which incurs the lowest cost in terms
//     of arithmetic operations (`[CLRS]`_). Note that since this is a function to compute the product, :math:`N`
//     needs to be greater than or equal to 2; if equal to 2 then a trivial matrix-matrix product is returned.
//     If :math:`N` is 1, then this is a no-op - the original matrix is returned as is.
//
//     .. warning::
//
//         :func:`torch.chain_matmul` is deprecated and will be removed in a future PyTorch release.
//         Use :func:`torch.linalg.multi_dot` instead, which accepts a list of two or more tensors
//         rather than multiple arguments.
//
//     Args:
//         matrices (Tensors...): a sequence of 2 or more 2-D tensors whose product is to be determined.
//         out (Tensor, optional): the output tensor. Ignored if :attr:`out` = ``None``.
//
//     Returns:
//         Tensor: if the :math:`i^{th}` tensor was of dimensions :math:`p_{i} \times p_{i + 1}`, then the product
//         would be of dimensions :math:`p_{1} \times p_{N + 1}`.
//
//     Example::
//
//         >>> # xdoctest: +SKIP
//         >>> # xdoctest: +IGNORE_WANT("non-deterministic")
//         >>> a = torch.randn(3, 4)
//         >>> b = torch.randn(4, 5)
//         >>> c = torch.randn(5, 6)
//         >>> d = torch.randn(6, 7)
//         >>> # will raise a deprecation warning
//         >>> torch.chain_matmul(a, b, c, d)
//         tensor([[ -2.3375,  -3.9790,  -4.1119,  -6.6577,   9.5609, -11.5095,  -3.2614],
//                 [ 21.4038,   3.3378,  -8.4982,  -5.2457, -10.2561,  -2.4684,   2.7163],
//                 [ -0.9647,  -5.8917,  -2.3213,  -5.2284,  12.8615, -12.2816,  -2.5095]])
//
//     .. _`[CLRS]`: https://mitpress.mit.edu/books/introduction-algorithms-third-edition
//
//
//go:linkname ChainMatmul py.chain_matmul
func ChainMatmul(__llgo_va_list ...interface{}) *py.Object
//
// channel_shuffle(input, groups) -> Tensor
//
// Divide the channels in a tensor of shape :math:`(*, C , H, W)`
// into g groups and rearrange them as :math:`(*, C \frac g, g, H, W)`,
// while keeping the original tensor shape.
//
// See :class:`~torch.nn.ChannelShuffle` for details.
//
// Args:
//     input (Tensor): the input tensor
//     groups (int): number of groups to divide channels in and rearrange.
//
// Examples::
//
//     >>> input = torch.randn(1, 4, 2, 2)
//     >>> print(input)
//     [[[[1, 2],
//        [3, 4]],
//       [[5, 6],
//        [7, 8]],
//       [[9, 10],
//        [11, 12]],
//       [[13, 14],
//        [15, 16]],
//      ]]
//     >>> output = torch.nn.functional.channel_shuffle(input, 2)
//     >>> print(output)
//     [[[[1, 2],
//        [3, 4]],
//       [[9, 10],
//        [11, 12]],
//       [[5, 6],
//        [7, 8]],
//       [[13, 14],
//        [15, 16]],
//      ]]
//
//
//go:linkname ChannelShuffle py.channel_shuffle
func ChannelShuffle(input *py.Object, groups *py.Object) *py.Object
//
// cholesky(input, upper=False, *, out=None) -> Tensor
//
// Computes the Cholesky decomposition of a symmetric positive-definite
// matrix :math:`A` or for batches of symmetric positive-definite matrices.
//
// If :attr:`upper` is ``True``, the returned matrix ``U`` is upper-triangular, and
// the decomposition has the form:
//
// .. math::
//
//   A = U^TU
//
// If :attr:`upper` is ``False``, the returned matrix ``L`` is lower-triangular, and
// the decomposition has the form:
//
// .. math::
//
//     A = LL^T
//
// If :attr:`upper` is ``True``, and :math:`A` is a batch of symmetric positive-definite
// matrices, then the returned tensor will be composed of upper-triangular Cholesky factors
// of each of the individual matrices. Similarly, when :attr:`upper` is ``False``, the returned
// tensor will be composed of lower-triangular Cholesky factors of each of the individual
// matrices.
//
// .. warning::
//
//     :func:`torch.cholesky` is deprecated in favor of :func:`torch.linalg.cholesky`
//     and will be removed in a future PyTorch release.
//
//     ``L = torch.cholesky(A)`` should be replaced with
//
//     .. code:: python
//
//         L = torch.linalg.cholesky(A)
//
//     ``U = torch.cholesky(A, upper=True)`` should be replaced with
//
//     .. code:: python
//
//         U = torch.linalg.cholesky(A).mH
//
//     This transform will produce equivalent results for all valid (symmetric positive definite) inputs.
//
// Args:
//     input (Tensor): the input tensor :math:`A` of size :math:`(*, n, n)` where `*` is zero or more
//                 batch dimensions consisting of symmetric positive-definite matrices.
//     upper (bool, optional): flag that indicates whether to return a
//                             upper or lower triangular matrix. Default: ``False``
//
// Keyword args:
//     out (Tensor, optional): the output matrix
//
// Example::
//
//     >>> a = torch.randn(3, 3)
//     >>> a = a @ a.mT + 1e-3 # make symmetric positive-definite
//     >>> l = torch.cholesky(a)
//     >>> a
//     tensor([[ 2.4112, -0.7486,  1.4551],
//             [-0.7486,  1.3544,  0.1294],
//             [ 1.4551,  0.1294,  1.6724]])
//     >>> l
//     tensor([[ 1.5528,  0.0000,  0.0000],
//             [-0.4821,  1.0592,  0.0000],
//             [ 0.9371,  0.5487,  0.7023]])
//     >>> l @ l.mT
//     tensor([[ 2.4112, -0.7486,  1.4551],
//             [-0.7486,  1.3544,  0.1294],
//             [ 1.4551,  0.1294,  1.6724]])
//     >>> a = torch.randn(3, 2, 2) # Example for batched input
//     >>> a = a @ a.mT + 1e-03 # make symmetric positive-definite
//     >>> l = torch.cholesky(a)
//     >>> z = l @ l.mT
//     >>> torch.dist(z, a)
//     tensor(2.3842e-07)
//
//
//go:linkname Cholesky py.cholesky
func Cholesky(input *py.Object, upper *py.Object) *py.Object
//
// cholesky_inverse(L, upper=False, *, out=None) -> Tensor
//
// Computes the inverse of a complex Hermitian or real symmetric
// positive-definite matrix given its Cholesky decomposition.
//
// Let :math:`A` be a complex Hermitian or real symmetric positive-definite matrix,
// and :math:`L` its Cholesky decomposition such that:
//
// .. math::
//
//     A = LL^{\text{H}}
//
// where :math:`L^{\text{H}}` is the conjugate transpose when :math:`L` is complex,
// and the transpose when :math:`L` is real-valued.
//
// Computes the inverse matrix :math:`A^{-1}`.
//
// Supports input of float, double, cfloat and cdouble dtypes.
// Also supports batches of matrices, and if :math:`A` is a batch of matrices
// then the output has the same batch dimensions.
//
// Args:
//     L (Tensor): tensor of shape `(*, n, n)` where `*` is zero or more batch dimensions
//         consisting of lower or upper triangular Cholesky decompositions of
//         symmetric or Hermitian positive-definite matrices.
//     upper (bool, optional): flag that indicates whether :math:`L` is lower triangular
//         or upper triangular. Default: ``False``
//
// Keyword args:
//     out (Tensor, optional): output tensor. Ignored if `None`. Default: `None`.
//
// Example::
//
//     >>> A = torch.randn(3, 3)
//     >>> A = A @ A.T + torch.eye(3) * 1e-3 # Creates a symmetric positive-definite matrix
//     >>> L = torch.linalg.cholesky(A) # Extract Cholesky decomposition
//     >>> torch.cholesky_inverse(L)
//     tensor([[ 1.9314,  1.2251, -0.0889],
//             [ 1.2251,  2.4439,  0.2122],
//             [-0.0889,  0.2122,  0.1412]])
//     >>> A.inverse()
//     tensor([[ 1.9314,  1.2251, -0.0889],
//             [ 1.2251,  2.4439,  0.2122],
//             [-0.0889,  0.2122,  0.1412]])
//
//     >>> A = torch.randn(3, 2, 2, dtype=torch.complex64)
//     >>> A = A @ A.mH + torch.eye(2) * 1e-3 # Batch of Hermitian positive-definite matrices
//     >>> L = torch.linalg.cholesky(A)
//     >>> torch.dist(torch.inverse(A), torch.cholesky_inverse(L))
//     tensor(5.6358e-7)
//
//
//go:linkname CholeskyInverse py.cholesky_inverse
func CholeskyInverse(L *py.Object, upper *py.Object) *py.Object
//
// cholesky_solve(B, L, upper=False, *, out=None) -> Tensor
//
// Computes the solution of a system of linear equations with complex Hermitian
// or real symmetric positive-definite lhs given its Cholesky decomposition.
//
// Let :math:`A` be a complex Hermitian or real symmetric positive-definite matrix,
// and :math:`L` its Cholesky decomposition such that:
//
// .. math::
//
//     A = LL^{\text{H}}
//
// where :math:`L^{\text{H}}` is the conjugate transpose when :math:`L` is complex,
// and the transpose when :math:`L` is real-valued.
//
// Returns the solution :math:`X` of the following linear system:
//
// .. math::
//
//     AX = B
//
// Supports inputs of float, double, cfloat and cdouble dtypes.
// Also supports batches of matrices, and if :math:`A` or :math:`B` is a batch of matrices
// then the output has the same batch dimensions.
//
// Args:
//     B (Tensor): right-hand side tensor of shape `(*, n, k)`
//         where :math:`*` is zero or more batch dimensions
//     L (Tensor): tensor of shape `(*, n, n)` where `*` is zero or more batch dimensions
//         consisting of lower or upper triangular Cholesky decompositions of
//         symmetric or Hermitian positive-definite matrices.
//     upper (bool, optional): flag that indicates whether :math:`L` is lower triangular
//         or upper triangular. Default: ``False``.
//
// Keyword args:
//     out (Tensor, optional): output tensor. Ignored if `None`. Default: `None`.
//
// Example::
//
//     >>> A = torch.randn(3, 3)
//     >>> A = A @ A.T + torch.eye(3) * 1e-3 # Creates a symmetric positive-definite matrix
//     >>> L = torch.linalg.cholesky(A) # Extract Cholesky decomposition
//     >>> B = torch.randn(3, 2)
//     >>> torch.cholesky_solve(B, L)
//     tensor([[ -8.1625,  19.6097],
//             [ -5.8398,  14.2387],
//             [ -4.3771,  10.4173]])
//     >>> A.inverse() @  B
//     tensor([[ -8.1626,  19.6097],
//             [ -5.8398,  14.2387],
//             [ -4.3771,  10.4173]])
//
//     >>> A = torch.randn(3, 2, 2, dtype=torch.complex64)
//     >>> A = A @ A.mH + torch.eye(2) * 1e-3 # Batch of Hermitian positive-definite matrices
//     >>> L = torch.linalg.cholesky(A)
//     >>> B = torch.randn(2, 1, dtype=torch.complex64)
//     >>> X = torch.cholesky_solve(B, L)
//     >>> torch.dist(X, A.inverse() @ B)
//     tensor(1.6881e-5)
//
//
//go:linkname CholeskySolve py.cholesky_solve
func CholeskySolve(B *py.Object, L *py.Object, upper *py.Object) *py.Object
// None
//
//go:linkname ChooseQparamsOptimized py.choose_qparams_optimized
func ChooseQparamsOptimized(__llgo_va_list ...interface{}) *py.Object
//
// chunk(input, chunks, dim=0) -> List of Tensors
//
// Attempts to split a tensor into the specified number of chunks. Each chunk is a view of
// the input tensor.
//
//
// .. note::
//
//     This function may return fewer than the specified number of chunks!
//
// .. seealso::
//
//     :func:`torch.tensor_split` a function that always returns exactly the specified number of chunks
//
// If the tensor size along the given dimension :attr:`dim` is divisible by :attr:`chunks`,
// all returned chunks will be the same size.
// If the tensor size along the given dimension :attr:`dim` is not divisible by :attr:`chunks`,
// all returned chunks will be the same size, except the last one.
// If such division is not possible, this function may return fewer
// than the specified number of chunks.
//
// Arguments:
//     input (Tensor): the tensor to split
//     chunks (int): number of chunks to return
//     dim (int): dimension along which to split the tensor
//
// Example:
//     >>> torch.arange(11).chunk(6)
//     (tensor([0, 1]),
//      tensor([2, 3]),
//      tensor([4, 5]),
//      tensor([6, 7]),
//      tensor([8, 9]),
//      tensor([10]))
//     >>> torch.arange(12).chunk(6)
//     (tensor([0, 1]),
//      tensor([2, 3]),
//      tensor([4, 5]),
//      tensor([6, 7]),
//      tensor([8, 9]),
//      tensor([10, 11]))
//     >>> torch.arange(13).chunk(6)
//     (tensor([0, 1, 2]),
//      tensor([3, 4, 5]),
//      tensor([6, 7, 8]),
//      tensor([ 9, 10, 11]),
//      tensor([12]))
//
//
//go:linkname Chunk py.chunk
func Chunk(input *py.Object, chunks *py.Object, dim *py.Object) *py.Object
//
// clamp(input, min=None, max=None, *, out=None) -> Tensor
//
// Clamps all elements in :attr:`input` into the range `[` :attr:`min`, :attr:`max` `]`.
// Letting min_value and max_value be :attr:`min` and :attr:`max`, respectively, this returns:
//
// .. math::
//     y_i = \min(\max(x_i, \text{min\_value}_i), \text{max\_value}_i)
//
// If :attr:`min` is ``None``, there is no lower bound.
// Or, if :attr:`max` is ``None`` there is no upper bound.
//
//
// .. note::
//     If :attr:`min` is greater than :attr:`max` :func:`torch.clamp(..., min, max) <torch.clamp>`
//     sets all elements in :attr:`input` to the value of :attr:`max`.
//
// Args:
//     input (Tensor): the input tensor.
//     min (Number or Tensor, optional): lower-bound of the range to be clamped to
//     max (Number or Tensor, optional): upper-bound of the range to be clamped to
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([-1.7120,  0.1734, -0.0478, -0.0922])
//     >>> torch.clamp(a, min=-0.5, max=0.5)
//     tensor([-0.5000,  0.1734, -0.0478, -0.0922])
//
//     >>> min = torch.linspace(-1, 1, steps=4)
//     >>> torch.clamp(a, min=min)
//     tensor([-1.0000,  0.1734,  0.3333,  1.0000])
//
//
//
//go:linkname Clamp py.clamp
func Clamp(input *py.Object, min *py.Object, max *py.Object) *py.Object
// None
//
//go:linkname Clamp_ py.clamp_
func Clamp_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname ClampMax py.clamp_max
func ClampMax(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname ClampMax_ py.clamp_max_
func ClampMax_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname ClampMin py.clamp_min
func ClampMin(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname ClampMin_ py.clamp_min_
func ClampMin_(__llgo_va_list ...interface{}) *py.Object
//
// clip(input, min=None, max=None, *, out=None) -> Tensor
//
// Alias for :func:`torch.clamp`.
//
//
//go:linkname Clip py.clip
func Clip(input *py.Object, min *py.Object, max *py.Object) *py.Object
// None
//
//go:linkname Clip_ py.clip_
func Clip_(__llgo_va_list ...interface{}) *py.Object
//
// clone(input, *, memory_format=torch.preserve_format) -> Tensor
//
// Returns a copy of :attr:`input`.
//
// .. note::
//
//     This function is differentiable, so gradients will flow back from the
//     result of this operation to :attr:`input`. To create a tensor without an
//     autograd relationship to :attr:`input` see :meth:`~Tensor.detach`.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     memory_format (:class:`torch.memory_format`, optional): the desired memory format of
//         returned tensor. Default: ``torch.preserve_format``.
//
//
//go:linkname Clone py.clone
func Clone(input *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.col_indices`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname ColIndicesCopy py.col_indices_copy
func ColIndicesCopy(__llgo_va_list ...interface{}) *py.Object
//
// column_stack(tensors, *, out=None) -> Tensor
//
// Creates a new tensor by horizontally stacking the tensors in :attr:`tensors`.
//
// Equivalent to ``torch.hstack(tensors)``, except each zero or one dimensional tensor ``t``
// in :attr:`tensors` is first reshaped into a ``(t.numel(), 1)`` column before being stacked horizontally.
//
// Args:
//     tensors (sequence of Tensors): sequence of tensors to concatenate
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([1, 2, 3])
//     >>> b = torch.tensor([4, 5, 6])
//     >>> torch.column_stack((a, b))
//     tensor([[1, 4],
//         [2, 5],
//         [3, 6]])
//     >>> a = torch.arange(5)
//     >>> b = torch.arange(10).reshape(5, 2)
//     >>> torch.column_stack((a, b, b))
//     tensor([[0, 0, 1, 0, 1],
//             [1, 2, 3, 2, 3],
//             [2, 4, 5, 4, 5],
//             [3, 6, 7, 6, 7],
//             [4, 8, 9, 8, 9]])
//
//
//
//go:linkname ColumnStack py.column_stack
func ColumnStack(tensors *py.Object) *py.Object
//
// combinations(input, r=2, with_replacement=False) -> seq
//
// Compute combinations of length :math:`r` of the given tensor. The behavior is similar to
// python's `itertools.combinations` when `with_replacement` is set to `False`, and
// `itertools.combinations_with_replacement` when `with_replacement` is set to `True`.
//
// Arguments:
//     input (Tensor): 1D vector.
//     r (int, optional): number of elements to combine
//     with_replacement (bool, optional): whether to allow duplication in combination
//
// Returns:
//     Tensor: A tensor equivalent to converting all the input tensors into lists, do
//     `itertools.combinations` or `itertools.combinations_with_replacement` on these
//     lists, and finally convert the resulting list into tensor.
//
// Example::
//
//     >>> a = [1, 2, 3]
//     >>> list(itertools.combinations(a, r=2))
//     [(1, 2), (1, 3), (2, 3)]
//     >>> list(itertools.combinations(a, r=3))
//     [(1, 2, 3)]
//     >>> list(itertools.combinations_with_replacement(a, r=2))
//     [(1, 1), (1, 2), (1, 3), (2, 2), (2, 3), (3, 3)]
//     >>> tensor_a = torch.tensor(a)
//     >>> torch.combinations(tensor_a)
//     tensor([[1, 2],
//             [1, 3],
//             [2, 3]])
//     >>> torch.combinations(tensor_a, r=3)
//     tensor([[1, 2, 3]])
//     >>> torch.combinations(tensor_a, with_replacement=True)
//     tensor([[1, 1],
//             [1, 2],
//             [1, 3],
//             [2, 2],
//             [2, 3],
//             [3, 3]])
//
//
//
//go:linkname Combinations py.combinations
func Combinations(input *py.Object, r *py.Object, withReplacement *py.Object) *py.Object
//
// complex(real, imag, *, out=None) -> Tensor
//
// Constructs a complex tensor with its real part equal to :attr:`real` and its
// imaginary part equal to :attr:`imag`.
//
// Args:
//     real (Tensor): The real part of the complex tensor. Must be half, float or double.
//     imag (Tensor): The imaginary part of the complex tensor. Must be same dtype
//         as :attr:`real`.
//
// Keyword args:
//     out (Tensor): If the inputs are ``torch.float32``, must be
//         ``torch.complex64``. If the inputs are ``torch.float64``, must be
//         ``torch.complex128``.
//
// Example::
//
//     >>> real = torch.tensor([1, 2], dtype=torch.float32)
//     >>> imag = torch.tensor([3, 4], dtype=torch.float32)
//     >>> z = torch.complex(real, imag)
//     >>> z
//     tensor([(1.+3.j), (2.+4.j)])
//     >>> z.dtype
//     torch.complex64
//
//
//
//go:linkname Complex py.complex
func Complex(real *py.Object, imag *py.Object) *py.Object
//
// concat(tensors, dim=0, *, out=None) -> Tensor
//
// Alias of :func:`torch.cat`.
//
//
//go:linkname Concat py.concat
func Concat(tensors *py.Object, dim *py.Object) *py.Object
//
// concatenate(tensors, axis=0, out=None) -> Tensor
//
// Alias of :func:`torch.cat`.
//
//
//go:linkname Concatenate py.concatenate
func Concatenate(tensors *py.Object, axis *py.Object, out *py.Object) *py.Object
//
// conj(input) -> Tensor
//
// Returns a view of :attr:`input` with a flipped conjugate bit. If :attr:`input` has a non-complex dtype,
// this function just returns :attr:`input`.
//
// .. note::
//     :func:`torch.conj` performs a lazy conjugation, but the actual conjugated tensor can be materialized
//     at any time using :func:`torch.resolve_conj`.
//
// .. warning:: In the future, :func:`torch.conj` may return a non-writeable view for an :attr:`input` of
//              non-complex dtype. It's recommended that programs not modify the tensor returned by :func:`torch.conj_physical`
//              when :attr:`input` is of non-complex dtype to be compatible with this change.
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> x = torch.tensor([-1 + 1j, -2 + 2j, 3 - 3j])
//     >>> x.is_conj()
//     False
//     >>> y = torch.conj(x)
//     >>> y.is_conj()
//     True
//
//
//go:linkname Conj py.conj
func Conj(input *py.Object) *py.Object
//
// conj_physical(input, *, out=None) -> Tensor
//
// Computes the element-wise conjugate of the given :attr:`input` tensor.
// If :attr:`input` has a non-complex dtype, this function just returns :attr:`input`.
//
// .. note::
//    This performs the conjugate operation regardless of the fact conjugate bit is set or not.
//
// .. warning:: In the future, :func:`torch.conj_physical` may return a non-writeable view for an :attr:`input` of
//              non-complex dtype. It's recommended that programs not modify the tensor returned by :func:`torch.conj_physical`
//              when :attr:`input` is of non-complex dtype to be compatible with this change.
//
// .. math::
//     \text{out}_{i} = conj(\text{input}_{i})
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.conj_physical(torch.tensor([-1 + 1j, -2 + 2j, 3 - 3j]))
//     tensor([-1 - 1j, -2 - 2j, 3 + 3j])
//
//
//go:linkname ConjPhysical py.conj_physical
func ConjPhysical(input *py.Object) *py.Object
// None
//
//go:linkname ConjPhysical_ py.conj_physical_
func ConjPhysical_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname ConstantPadNd py.constant_pad_nd
func ConstantPadNd(__llgo_va_list ...interface{}) *py.Object
//
// conv1d(input, weight, bias=None, stride=1, padding=0, dilation=1, groups=1) -> Tensor
//
// Applies a 1D convolution over an input signal composed of several input
// planes.
//
// This operator supports :ref:`TensorFloat32<tf32_on_ampere>`.
//
// See :class:`~torch.nn.Conv1d` for details and output shape.
//
// Note:
//     In some circumstances when given tensors on a CUDA device and using CuDNN, this operator may select a nondeterministic algorithm to increase performance. If this is undesirable, you can try to make the operation deterministic (potentially at a performance cost) by setting ``torch.backends.cudnn.deterministic = True``. See :doc:`/notes/randomness` for more information.
//
// Note:
//     This operator supports complex data types i.e. ``complex32, complex64, complex128``.
//
//
// Args:
//     input: input tensor of shape :math:`(\text{minibatch} , \text{in\_channels} , iW)`
//     weight: filters of shape :math:`(\text{out\_channels} , \frac{\text{in\_channels}}{\text{groups}} , kW)`
//     bias: optional bias of shape :math:`(\text{out\_channels})`. Default: ``None``
//     stride: the stride of the convolving kernel. Can be a single number or
//       a one-element tuple `(sW,)`. Default: 1
//     padding: implicit paddings on both sides of the input. Can be a string {'valid', 'same'},
//       single number or a one-element tuple `(padW,)`. Default: 0
//       ``padding='valid'`` is the same as no padding. ``padding='same'`` pads
//       the input so the output has the same shape as the input. However, this mode
//       doesn't support any stride values other than 1.
//
//       .. warning::
//           For ``padding='same'``, if the ``weight`` is even-length and
//           ``dilation`` is odd in any dimension, a full :func:`pad` operation
//           may be needed internally. Lowering performance.
//     dilation: the spacing between kernel elements. Can be a single number or
//       a one-element tuple `(dW,)`. Default: 1
//     groups: split input into groups, :math:`\text{in\_channels}` should be divisible by
//       the number of groups. Default: 1
//
// Examples::
//
//     >>> inputs = torch.randn(33, 16, 30)
//     >>> filters = torch.randn(20, 16, 5)
//     >>> F.conv1d(inputs, filters)
//
//
//go:linkname Conv1d py.conv1d
func Conv1d(input *py.Object, weight *py.Object, bias *py.Object, stride *py.Object, padding *py.Object, dilation *py.Object, groups *py.Object) *py.Object
//
// conv2d(input, weight, bias=None, stride=1, padding=0, dilation=1, groups=1) -> Tensor
//
// Applies a 2D convolution over an input image composed of several input
// planes.
//
// This operator supports :ref:`TensorFloat32<tf32_on_ampere>`.
//
// See :class:`~torch.nn.Conv2d` for details and output shape.
//
// Note:
//     In some circumstances when given tensors on a CUDA device and using CuDNN, this operator may select a nondeterministic algorithm to increase performance. If this is undesirable, you can try to make the operation deterministic (potentially at a performance cost) by setting ``torch.backends.cudnn.deterministic = True``. See :doc:`/notes/randomness` for more information.
//
// Note:
//     This operator supports complex data types i.e. ``complex32, complex64, complex128``.
//
//
// Args:
//     input: input tensor of shape :math:`(\text{minibatch} , \text{in\_channels} , iH , iW)`
//     weight: filters of shape :math:`(\text{out\_channels} , \frac{\text{in\_channels}}{\text{groups}} , kH , kW)`
//     bias: optional bias tensor of shape :math:`(\text{out\_channels})`. Default: ``None``
//     stride: the stride of the convolving kernel. Can be a single number or a
//       tuple `(sH, sW)`. Default: 1
//     padding: implicit paddings on both sides of the input. Can be a string {'valid', 'same'},
//       single number or a tuple `(padH, padW)`. Default: 0
//       ``padding='valid'`` is the same as no padding. ``padding='same'`` pads
//       the input so the output has the same shape as the input. However, this mode
//       doesn't support any stride values other than 1.
//
//       .. warning::
//           For ``padding='same'``, if the ``weight`` is even-length and
//           ``dilation`` is odd in any dimension, a full :func:`pad` operation
//           may be needed internally. Lowering performance.
//
//     dilation: the spacing between kernel elements. Can be a single number or
//       a tuple `(dH, dW)`. Default: 1
//     groups: split input into groups, both :math:`\text{in\_channels}` and :math:`\text{out\_channels}`
//       should be divisible by the number of groups. Default: 1
//
// Examples::
//
//     >>> # With square kernels and equal stride
//     >>> filters = torch.randn(8, 4, 3, 3)
//     >>> inputs = torch.randn(1, 4, 5, 5)
//     >>> F.conv2d(inputs, filters, padding=1)
//
//
//go:linkname Conv2d py.conv2d
func Conv2d(input *py.Object, weight *py.Object, bias *py.Object, stride *py.Object, padding *py.Object, dilation *py.Object, groups *py.Object) *py.Object
//
// conv3d(input, weight, bias=None, stride=1, padding=0, dilation=1, groups=1) -> Tensor
//
// Applies a 3D convolution over an input image composed of several input
// planes.
//
// This operator supports :ref:`TensorFloat32<tf32_on_ampere>`.
//
// See :class:`~torch.nn.Conv3d` for details and output shape.
//
// Note:
//     In some circumstances when given tensors on a CUDA device and using CuDNN, this operator may select a nondeterministic algorithm to increase performance. If this is undesirable, you can try to make the operation deterministic (potentially at a performance cost) by setting ``torch.backends.cudnn.deterministic = True``. See :doc:`/notes/randomness` for more information.
//
// Note:
//     This operator supports complex data types i.e. ``complex32, complex64, complex128``.
//
//
// Args:
//     input: input tensor of shape :math:`(\text{minibatch} , \text{in\_channels} , iT , iH , iW)`
//     weight: filters of shape :math:`(\text{out\_channels} , \frac{\text{in\_channels}}{\text{groups}} , kT , kH , kW)`
//     bias: optional bias tensor of shape :math:`(\text{out\_channels})`. Default: None
//     stride: the stride of the convolving kernel. Can be a single number or a
//       tuple `(sT, sH, sW)`. Default: 1
//     padding: implicit paddings on both sides of the input. Can be a string {'valid', 'same'},
//       single number or a tuple `(padT, padH, padW)`. Default: 0
//       ``padding='valid'`` is the same as no padding. ``padding='same'`` pads
//       the input so the output has the same shape as the input. However, this mode
//       doesn't support any stride values other than 1.
//
//       .. warning::
//           For ``padding='same'``, if the ``weight`` is even-length and
//           ``dilation`` is odd in any dimension, a full :func:`pad` operation
//           may be needed internally. Lowering performance.
//
//     dilation: the spacing between kernel elements. Can be a single number or
//       a tuple `(dT, dH, dW)`. Default: 1
//     groups: split input into groups, :math:`\text{in\_channels}` should be divisible by
//       the number of groups. Default: 1
//
// Examples::
//
//     >>> filters = torch.randn(33, 16, 3, 3, 3)
//     >>> inputs = torch.randn(20, 16, 50, 10, 20)
//     >>> F.conv3d(inputs, filters)
//
//
//go:linkname Conv3d py.conv3d
func Conv3d(input *py.Object, weight *py.Object, bias *py.Object, stride *py.Object, padding *py.Object, dilation *py.Object, groups *py.Object) *py.Object
//
// Applies a 1-dimensional sequence convolution over an input sequence.
// Input and output dimensions are (Time, Batch, Channels) - hence TBC.
//
// Args:
//     input: input tensor of shape :math:`(\text{sequence length} \times batch \times \text{in\_channels})`
//     weight: filter of shape (:math:`\text{kernel width} \times \text{in\_channels} \times \text{out\_channels}`)
//     bias: bias of shape (:math:`\text{out\_channels}`)
//     pad: number of timesteps to pad. Default: 0
//
//
//go:linkname ConvTbc py.conv_tbc
func ConvTbc(__llgo_va_list ...interface{}) *py.Object
//
// conv_transpose1d(input, weight, bias=None, stride=1, padding=0, output_padding=0, groups=1, dilation=1) -> Tensor
//
// Applies a 1D transposed convolution operator over an input signal
// composed of several input planes, sometimes also called "deconvolution".
//
// This operator supports :ref:`TensorFloat32<tf32_on_ampere>`.
//
// See :class:`~torch.nn.ConvTranspose1d` for details and output shape.
//
// Note:
//     In some circumstances when given tensors on a CUDA device and using CuDNN, this operator may select a nondeterministic algorithm to increase performance. If this is undesirable, you can try to make the operation deterministic (potentially at a performance cost) by setting ``torch.backends.cudnn.deterministic = True``. See :doc:`/notes/randomness` for more information.
//
//
// Args:
//     input: input tensor of shape :math:`(\text{minibatch} , \text{in\_channels} , iW)`
//     weight: filters of shape :math:`(\text{in\_channels} , \frac{\text{out\_channels}}{\text{groups}} , kW)`
//     bias: optional bias of shape :math:`(\text{out\_channels})`. Default: None
//     stride: the stride of the convolving kernel. Can be a single number or a
//       tuple ``(sW,)``. Default: 1
//     padding: ``dilation * (kernel_size - 1) - padding`` zero-padding will be added to both
//       sides of each dimension in the input. Can be a single number or a tuple
//       ``(padW,)``. Default: 0
//     output_padding: additional size added to one side of each dimension in the
//       output shape. Can be a single number or a tuple ``(out_padW)``. Default: 0
//     groups: split input into groups, :math:`\text{in\_channels}` should be divisible by the
//       number of groups. Default: 1
//     dilation: the spacing between kernel elements. Can be a single number or
//       a tuple ``(dW,)``. Default: 1
//
// Examples::
//
//     >>> inputs = torch.randn(20, 16, 50)
//     >>> weights = torch.randn(16, 33, 5)
//     >>> F.conv_transpose1d(inputs, weights)
//
//
//go:linkname ConvTranspose1d py.conv_transpose1d
func ConvTranspose1d(input *py.Object, weight *py.Object, bias *py.Object, stride *py.Object, padding *py.Object, outputPadding *py.Object, groups *py.Object, dilation *py.Object) *py.Object
//
// conv_transpose2d(input, weight, bias=None, stride=1, padding=0, output_padding=0, groups=1, dilation=1) -> Tensor
//
// Applies a 2D transposed convolution operator over an input image
// composed of several input planes, sometimes also called "deconvolution".
//
// This operator supports :ref:`TensorFloat32<tf32_on_ampere>`.
//
// See :class:`~torch.nn.ConvTranspose2d` for details and output shape.
//
// Note:
//     In some circumstances when given tensors on a CUDA device and using CuDNN, this operator may select a nondeterministic algorithm to increase performance. If this is undesirable, you can try to make the operation deterministic (potentially at a performance cost) by setting ``torch.backends.cudnn.deterministic = True``. See :doc:`/notes/randomness` for more information.
//
//
// Args:
//     input: input tensor of shape :math:`(\text{minibatch} , \text{in\_channels} , iH , iW)`
//     weight: filters of shape :math:`(\text{in\_channels} , \frac{\text{out\_channels}}{\text{groups}} , kH , kW)`
//     bias: optional bias of shape :math:`(\text{out\_channels})`. Default: None
//     stride: the stride of the convolving kernel. Can be a single number or a
//       tuple ``(sH, sW)``. Default: 1
//     padding: ``dilation * (kernel_size - 1) - padding`` zero-padding will be added to both
//       sides of each dimension in the input. Can be a single number or a tuple
//       ``(padH, padW)``. Default: 0
//     output_padding: additional size added to one side of each dimension in the
//       output shape. Can be a single number or a tuple ``(out_padH, out_padW)``.
//       Default: 0
//     groups: split input into groups, :math:`\text{in\_channels}` should be divisible by the
//       number of groups. Default: 1
//     dilation: the spacing between kernel elements. Can be a single number or
//       a tuple ``(dH, dW)``. Default: 1
//
// Examples::
//
//     >>> # With square kernels and equal stride
//     >>> inputs = torch.randn(1, 4, 5, 5)
//     >>> weights = torch.randn(4, 8, 3, 3)
//     >>> F.conv_transpose2d(inputs, weights, padding=1)
//
//
//go:linkname ConvTranspose2d py.conv_transpose2d
func ConvTranspose2d(input *py.Object, weight *py.Object, bias *py.Object, stride *py.Object, padding *py.Object, outputPadding *py.Object, groups *py.Object, dilation *py.Object) *py.Object
//
// conv_transpose3d(input, weight, bias=None, stride=1, padding=0, output_padding=0, groups=1, dilation=1) -> Tensor
//
// Applies a 3D transposed convolution operator over an input image
// composed of several input planes, sometimes also called "deconvolution"
//
// This operator supports :ref:`TensorFloat32<tf32_on_ampere>`.
//
// See :class:`~torch.nn.ConvTranspose3d` for details and output shape.
//
// Note:
//     In some circumstances when given tensors on a CUDA device and using CuDNN, this operator may select a nondeterministic algorithm to increase performance. If this is undesirable, you can try to make the operation deterministic (potentially at a performance cost) by setting ``torch.backends.cudnn.deterministic = True``. See :doc:`/notes/randomness` for more information.
//
//
// Args:
//     input: input tensor of shape :math:`(\text{minibatch} , \text{in\_channels} , iT , iH , iW)`
//     weight: filters of shape :math:`(\text{in\_channels} , \frac{\text{out\_channels}}{\text{groups}} , kT , kH , kW)`
//     bias: optional bias of shape :math:`(\text{out\_channels})`. Default: None
//     stride: the stride of the convolving kernel. Can be a single number or a
//       tuple ``(sT, sH, sW)``. Default: 1
//     padding: ``dilation * (kernel_size - 1) - padding`` zero-padding will be added to both
//       sides of each dimension in the input. Can be a single number or a tuple
//       ``(padT, padH, padW)``. Default: 0
//     output_padding: additional size added to one side of each dimension in the
//       output shape. Can be a single number or a tuple
//       ``(out_padT, out_padH, out_padW)``. Default: 0
//     groups: split input into groups, :math:`\text{in\_channels}` should be divisible by the
//       number of groups. Default: 1
//     dilation: the spacing between kernel elements. Can be a single number or
//       a tuple `(dT, dH, dW)`. Default: 1
//
// Examples::
//
//     >>> inputs = torch.randn(20, 16, 50, 10, 20)
//     >>> weights = torch.randn(16, 33, 3, 3, 3)
//     >>> F.conv_transpose3d(inputs, weights)
//
//
//go:linkname ConvTranspose3d py.conv_transpose3d
func ConvTranspose3d(input *py.Object, weight *py.Object, bias *py.Object, stride *py.Object, padding *py.Object, outputPadding *py.Object, groups *py.Object, dilation *py.Object) *py.Object
// None
//
//go:linkname Convolution py.convolution
func Convolution(__llgo_va_list ...interface{}) *py.Object
//
// copysign(input, other, *, out=None) -> Tensor
//
// Create a new floating-point tensor with the magnitude of :attr:`input` and the sign of :attr:`other`, elementwise.
//
// .. math::
//     \text{out}_{i} = \begin{cases}
//         -|\text{input}_{i}| & \text{if } \text{other}_{i} \leq -0.0 \\
//          |\text{input}_{i}| & \text{if } \text{other}_{i} \geq 0.0 \\
//     \end{cases}
//
//
// Supports :ref:`broadcasting to a common shape <broadcasting-semantics>`,
// and integer and float inputs.
//
// Args:
//     input (Tensor): magnitudes.
//     other (Tensor or Number): contains value(s) whose signbit(s) are
//         applied to the magnitudes in :attr:`input`.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(5)
//     >>> a
//     tensor([-1.2557, -0.0026, -0.5387,  0.4740, -0.9244])
//     >>> torch.copysign(a, 1)
//     tensor([1.2557, 0.0026, 0.5387, 0.4740, 0.9244])
//     >>> a = torch.randn(4, 4)
//     >>> a
//     tensor([[ 0.7079,  0.2778, -1.0249,  0.5719],
//             [-0.0059, -0.2600, -0.4475, -1.3948],
//             [ 0.3667, -0.9567, -2.5757, -0.1751],
//             [ 0.2046, -0.0742,  0.2998, -0.1054]])
//     >>> b = torch.randn(4)
//     tensor([ 0.2373,  0.3120,  0.3190, -1.1128])
//     >>> torch.copysign(a, b)
//     tensor([[ 0.7079,  0.2778,  1.0249, -0.5719],
//             [ 0.0059,  0.2600,  0.4475, -1.3948],
//             [ 0.3667,  0.9567,  2.5757, -0.1751],
//             [ 0.2046,  0.0742,  0.2998, -0.1054]])
//     >>> a = torch.tensor([1.])
//     >>> b = torch.tensor([-0.])
//     >>> torch.copysign(a, b)
//     tensor([-1.])
//
// .. note::
//     copysign handles signed zeros. If the other argument has a negative zero (-0),
//     the corresponding output value will be negative.
//
//
//
//go:linkname Copysign py.copysign
func Copysign(input *py.Object, other *py.Object) *py.Object
//
// corrcoef(input) -> Tensor
//
// Estimates the Pearson product-moment correlation coefficient matrix of the variables given by the :attr:`input` matrix,
// where rows are the variables and columns are the observations.
//
// .. note::
//
//     The correlation coefficient matrix R is computed using the covariance matrix C as given by
//     :math:`R_{ij} = \frac{ C_{ij} } { \sqrt{ C_{ii} * C_{jj} } }`
//
// .. note::
//
//     Due to floating point rounding, the resulting array may not be Hermitian and its diagonal elements may not be 1.
//     The real and imaginary values are clipped to the interval [-1, 1] in an attempt to improve this situation.
//
// Args:
//     input (Tensor): A 2D matrix containing multiple variables and observations, or a
//         Scalar or 1D vector representing a single variable.
//
// Returns:
//     (Tensor) The correlation coefficient matrix of the variables.
//
// .. seealso::
//
//         :func:`torch.cov` covariance matrix.
//
// Example::
//
//     >>> x = torch.tensor([[0, 1, 2], [2, 1, 0]])
//     >>> torch.corrcoef(x)
//     tensor([[ 1., -1.],
//             [-1.,  1.]])
//     >>> x = torch.randn(2, 4)
//     >>> x
//     tensor([[-0.2678, -0.0908, -0.3766,  0.2780],
//             [-0.5812,  0.1535,  0.2387,  0.2350]])
//     >>> torch.corrcoef(x)
//     tensor([[1.0000, 0.3582],
//             [0.3582, 1.0000]])
//     >>> torch.corrcoef(x[0])
//     tensor(1.)
//
//
//go:linkname Corrcoef py.corrcoef
func Corrcoef(input *py.Object) *py.Object
//
// cos(input, *, out=None) -> Tensor
//
// Returns a new tensor with the cosine  of the elements of :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \cos(\text{input}_{i})
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([ 1.4309,  1.2706, -0.8562,  0.9796])
//     >>> torch.cos(a)
//     tensor([ 0.1395,  0.2957,  0.6553,  0.5574])
//
//
//go:linkname Cos py.cos
func Cos(input *py.Object) *py.Object
// None
//
//go:linkname Cos_ py.cos_
func Cos_(__llgo_va_list ...interface{}) *py.Object
//
// cosh(input, *, out=None) -> Tensor
//
// Returns a new tensor with the hyperbolic cosine  of the elements of
// :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \cosh(\text{input}_{i})
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([ 0.1632,  1.1835, -0.6979, -0.7325])
//     >>> torch.cosh(a)
//     tensor([ 1.0133,  1.7860,  1.2536,  1.2805])
//
// .. note::
//    When :attr:`input` is on the CPU, the implementation of torch.cosh may use
//    the Sleef library, which rounds very large results to infinity or negative
//    infinity. See `here <https://sleef.org/purec.xhtml>`_ for details.
//
//
//go:linkname Cosh py.cosh
func Cosh(input *py.Object) *py.Object
// None
//
//go:linkname Cosh_ py.cosh_
func Cosh_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname CosineEmbeddingLoss py.cosine_embedding_loss
func CosineEmbeddingLoss(__llgo_va_list ...interface{}) *py.Object
//
// cosine_similarity(x1, x2, dim=1, eps=1e-8) -> Tensor
//
// Returns cosine similarity between ``x1`` and ``x2``, computed along dim. ``x1`` and ``x2`` must be broadcastable
// to a common shape. ``dim`` refers to the dimension in this common shape. Dimension ``dim`` of the output is
// squeezed (see :func:`torch.squeeze`), resulting in the
// output tensor having 1 fewer dimension.
//
// .. math ::
//     \text{similarity} = \dfrac{x_1 \cdot x_2}{\max(\Vert x_1 \Vert _2, \epsilon) \cdot \max(\Vert x_2 \Vert _2, \epsilon)}
//
// Supports :ref:`type promotion <type-promotion-doc>`.
//
// Args:
//     x1 (Tensor): First input.
//     x2 (Tensor): Second input.
//     dim (int, optional): Dimension along which cosine similarity is computed. Default: 1
//     eps (float, optional): Small value to avoid division by zero.
//         Default: 1e-8
//
// Example::
//
//     >>> input1 = torch.randn(100, 128)
//     >>> input2 = torch.randn(100, 128)
//     >>> output = F.cosine_similarity(input1, input2)
//     >>> print(output)
//
//
//go:linkname CosineSimilarity py.cosine_similarity
func CosineSimilarity(x1 *py.Object, x2 *py.Object, dim *py.Object, eps *py.Object) *py.Object
//
// count_nonzero(input, dim=None) -> Tensor
//
// Counts the number of non-zero values in the tensor :attr:`input` along the given :attr:`dim`.
// If no dim is specified then all non-zeros in the tensor are counted.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int or tuple of ints, optional): Dim or tuple of dims along which to count non-zeros.
//
// Example::
//
//     >>> x = torch.zeros(3,3)
//     >>> x[torch.randn(3,3) > 0.5] = 1
//     >>> x
//     tensor([[0., 1., 1.],
//             [0., 0., 0.],
//             [0., 0., 1.]])
//     >>> torch.count_nonzero(x)
//     tensor(3)
//     >>> torch.count_nonzero(x, dim=0)
//     tensor([0, 1, 2])
//
//
//go:linkname CountNonzero py.count_nonzero
func CountNonzero(input *py.Object, dim *py.Object) *py.Object
//
// cov(input, *, correction=1, fweights=None, aweights=None) -> Tensor
//
// Estimates the covariance matrix of the variables given by the :attr:`input` matrix, where rows are
// the variables and columns are the observations.
//
// A covariance matrix is a square matrix giving the covariance of each pair of variables. The diagonal contains
// the variance of each variable (covariance of a variable with itself). By definition, if :attr:`input` represents
// a single variable (Scalar or 1D) then its variance is returned.
//
// The sample covariance of the variables :math:`x` and :math:`y` is given by:
//
// .. math::
//     \text{cov}(x,y) = \frac{\sum^{N}_{i = 1}(x_{i} - \bar{x})(y_{i} - \bar{y})}{\max(0,~N~-~\delta N)}
//
// where :math:`\bar{x}` and :math:`\bar{y}` are the simple means of the :math:`x` and :math:`y` respectively, and
// :math:`\delta N` is the :attr:`correction`.
//
// If :attr:`fweights` and/or :attr:`aweights` are provided, the weighted covariance
// is calculated, which is given by:
//
// .. math::
//     \text{cov}_w(x,y) = \frac{\sum^{N}_{i = 1}w_i(x_{i} - \mu_x^*)(y_{i} - \mu_y^*)}
//     {\max(0,~\sum^{N}_{i = 1}w_i~-~\frac{\sum^{N}_{i = 1}w_ia_i}{\sum^{N}_{i = 1}w_i}~\delta N)}
//
// where :math:`w` denotes :attr:`fweights` or :attr:`aweights` (``f`` and ``a`` for brevity) based on whichever is
// provided, or :math:`w = f \times a` if both are provided, and
// :math:`\mu_x^* = \frac{\sum^{N}_{i = 1}w_ix_{i} }{\sum^{N}_{i = 1}w_i}` is the weighted mean of the variable. If not
// provided, ``f`` and/or ``a`` can be seen as a :math:`\mathbb{1}` vector of appropriate size.
//
// Args:
//     input (Tensor): A 2D matrix containing multiple variables and observations, or a
//         Scalar or 1D vector representing a single variable.
//
// Keyword Args:
//     correction (int, optional): difference between the sample size and sample degrees of freedom.
//         Defaults to Bessel's correction, ``correction = 1`` which returns the unbiased estimate,
//         even if both :attr:`fweights` and :attr:`aweights` are specified. ``correction = 0``
//         will return the simple average. Defaults to ``1``.
//     fweights (tensor, optional): A Scalar or 1D tensor of observation vector frequencies representing the number of
//         times each observation should be repeated. Its numel must equal the number of columns of :attr:`input`.
//         Must have integral dtype. Ignored if ``None``. Defaults to ``None``.
//     aweights (tensor, optional): A Scalar or 1D array of observation vector weights.
//         These relative weights are typically large for observations considered “important” and smaller for
//         observations considered less “important”. Its numel must equal the number of columns of :attr:`input`.
//         Must have floating point dtype. Ignored if ``None``. Defaults to ``None``.
//
// Returns:
//     (Tensor) The covariance matrix of the variables.
//
// .. seealso::
//
//         :func:`torch.corrcoef` normalized covariance matrix.
//
// Example::
//     >>> x = torch.tensor([[0, 2], [1, 1], [2, 0]]).T
//     >>> x
//     tensor([[0, 1, 2],
//             [2, 1, 0]])
//     >>> torch.cov(x)
//     tensor([[ 1., -1.],
//             [-1.,  1.]])
//     >>> torch.cov(x, correction=0)
//     tensor([[ 0.6667, -0.6667],
//             [-0.6667,  0.6667]])
//     >>> fw = torch.randint(1, 10, (3,))
//     >>> fw
//     tensor([1, 6, 9])
//     >>> aw = torch.rand(3)
//     >>> aw
//     tensor([0.4282, 0.0255, 0.4144])
//     >>> torch.cov(x, fweights=fw, aweights=aw)
//     tensor([[ 0.4169, -0.4169],
//             [-0.4169,  0.4169]])
//
//
//go:linkname Cov py.cov
func Cov(input *py.Object) *py.Object
//
// cross(input, other, dim=None, *, out=None) -> Tensor
//
//
// Returns the cross product of vectors in dimension :attr:`dim` of :attr:`input`
// and :attr:`other`.
//
// Supports input of float, double, cfloat and cdouble dtypes. Also supports batches
// of vectors, for which it computes the product along the dimension :attr:`dim`.
// In this case, the output has the same batch dimensions as the inputs.
//
// .. warning::
//     If :attr:`dim` is not given, it defaults to the first dimension found
//     with the size 3. Note that this might be unexpected.
//
//     This behavior is deprecated and will be changed to match that of :func:`torch.linalg.cross`
//     in a future release.
//
// .. seealso::
//         :func:`torch.linalg.cross` which has dim=-1 as default.
//
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor): the second input tensor
//     dim  (int, optional): the dimension to take the cross-product in.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4, 3)
//     >>> a
//     tensor([[-0.3956,  1.1455,  1.6895],
//             [-0.5849,  1.3672,  0.3599],
//             [-1.1626,  0.7180, -0.0521],
//             [-0.1339,  0.9902, -2.0225]])
//     >>> b = torch.randn(4, 3)
//     >>> b
//     tensor([[-0.0257, -1.4725, -1.2251],
//             [-1.1479, -0.7005, -1.9757],
//             [-1.3904,  0.3726, -1.1836],
//             [-0.9688, -0.7153,  0.2159]])
//     >>> torch.cross(a, b, dim=1)
//     tensor([[ 1.0844, -0.5281,  0.6120],
//             [-2.4490, -1.5687,  1.9792],
//             [-0.8304, -1.3037,  0.5650],
//             [-1.2329,  1.9883,  1.0551]])
//     >>> torch.cross(a, b)
//     tensor([[ 1.0844, -0.5281,  0.6120],
//             [-2.4490, -1.5687,  1.9792],
//             [-0.8304, -1.3037,  0.5650],
//             [-1.2329,  1.9883,  1.0551]])
//
//
//go:linkname Cross py.cross
func Cross(input *py.Object, other *py.Object, dim *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.crow_indices`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname CrowIndicesCopy py.crow_indices_copy
func CrowIndicesCopy(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname CtcLoss py.ctc_loss
func CtcLoss(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname CudnnAffineGridGenerator py.cudnn_affine_grid_generator
func CudnnAffineGridGenerator(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname CudnnBatchNorm py.cudnn_batch_norm
func CudnnBatchNorm(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname CudnnConvolution py.cudnn_convolution
func CudnnConvolution(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname CudnnConvolutionAddRelu py.cudnn_convolution_add_relu
func CudnnConvolutionAddRelu(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname CudnnConvolutionRelu py.cudnn_convolution_relu
func CudnnConvolutionRelu(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname CudnnConvolutionTranspose py.cudnn_convolution_transpose
func CudnnConvolutionTranspose(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname CudnnGridSampler py.cudnn_grid_sampler
func CudnnGridSampler(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname CudnnIsAcceptable py.cudnn_is_acceptable
func CudnnIsAcceptable(__llgo_va_list ...interface{}) *py.Object
//
// cummax(input, dim, *, out=None) -> (Tensor, LongTensor)
// Returns a namedtuple ``(values, indices)`` where ``values`` is the cumulative maximum of
// elements of :attr:`input` in the dimension :attr:`dim`. And ``indices`` is the index
// location of each maximum value found in the dimension :attr:`dim`.
//
// .. math::
//     y_i = max(x_1, x_2, x_3, \dots, x_i)
//
// Args:
//     input (Tensor): the input tensor.
//     dim  (int): the dimension to do the operation over
//
// Keyword args:
//     out (tuple, optional): the result tuple of two output tensors (values, indices)
//
// Example::
//
//     >>> a = torch.randn(10)
//     >>> a
//     tensor([-0.3449, -1.5447,  0.0685, -1.5104, -1.1706,  0.2259,  1.4696, -1.3284,
//          1.9946, -0.8209])
//     >>> torch.cummax(a, dim=0)
//     torch.return_types.cummax(
//         values=tensor([-0.3449, -0.3449,  0.0685,  0.0685,  0.0685,  0.2259,  1.4696,  1.4696,
//          1.9946,  1.9946]),
//         indices=tensor([0, 0, 2, 2, 2, 5, 6, 6, 8, 8]))
//
//
//go:linkname Cummax py.cummax
func Cummax(input *py.Object, dim *py.Object) *py.Object
//
// cummin(input, dim, *, out=None) -> (Tensor, LongTensor)
// Returns a namedtuple ``(values, indices)`` where ``values`` is the cumulative minimum of
// elements of :attr:`input` in the dimension :attr:`dim`. And ``indices`` is the index
// location of each maximum value found in the dimension :attr:`dim`.
//
// .. math::
//     y_i = min(x_1, x_2, x_3, \dots, x_i)
//
// Args:
//     input (Tensor): the input tensor.
//     dim  (int): the dimension to do the operation over
//
// Keyword args:
//     out (tuple, optional): the result tuple of two output tensors (values, indices)
//
// Example::
//
//     >>> a = torch.randn(10)
//     >>> a
//     tensor([-0.2284, -0.6628,  0.0975,  0.2680, -1.3298, -0.4220, -0.3885,  1.1762,
//          0.9165,  1.6684])
//     >>> torch.cummin(a, dim=0)
//     torch.return_types.cummin(
//         values=tensor([-0.2284, -0.6628, -0.6628, -0.6628, -1.3298, -1.3298, -1.3298, -1.3298,
//         -1.3298, -1.3298]),
//         indices=tensor([0, 1, 1, 1, 4, 4, 4, 4, 4, 4]))
//
//
//go:linkname Cummin py.cummin
func Cummin(input *py.Object, dim *py.Object) *py.Object
//
// cumprod(input, dim, *, dtype=None, out=None) -> Tensor
//
// Returns the cumulative product of elements of :attr:`input` in the dimension
// :attr:`dim`.
//
// For example, if :attr:`input` is a vector of size N, the result will also be
// a vector of size N, with elements.
//
// .. math::
//     y_i = x_1 \times x_2\times x_3\times \dots \times x_i
//
// Args:
//     input (Tensor): the input tensor.
//     dim  (int): the dimension to do the operation over
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         If specified, the input tensor is casted to :attr:`dtype` before the operation
//         is performed. This is useful for preventing data type overflows. Default: None.
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(10)
//     >>> a
//     tensor([ 0.6001,  0.2069, -0.1919,  0.9792,  0.6727,  1.0062,  0.4126,
//             -0.2129, -0.4206,  0.1968])
//     >>> torch.cumprod(a, dim=0)
//     tensor([ 0.6001,  0.1241, -0.0238, -0.0233, -0.0157, -0.0158, -0.0065,
//              0.0014, -0.0006, -0.0001])
//
//     >>> a[5] = 0.0
//     >>> torch.cumprod(a, dim=0)
//     tensor([ 0.6001,  0.1241, -0.0238, -0.0233, -0.0157, -0.0000, -0.0000,
//              0.0000, -0.0000, -0.0000])
//
//
//go:linkname Cumprod py.cumprod
func Cumprod(input *py.Object, dim *py.Object) *py.Object
//
// cumsum(input, dim, *, dtype=None, out=None) -> Tensor
//
// Returns the cumulative sum of elements of :attr:`input` in the dimension
// :attr:`dim`.
//
// For example, if :attr:`input` is a vector of size N, the result will also be
// a vector of size N, with elements.
//
// .. math::
//     y_i = x_1 + x_2 + x_3 + \dots + x_i
//
// Args:
//     input (Tensor): the input tensor.
//     dim  (int): the dimension to do the operation over
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         If specified, the input tensor is casted to :attr:`dtype` before the operation
//         is performed. This is useful for preventing data type overflows. Default: None.
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(10)
//     >>> a
//     tensor([-0.8286, -0.4890,  0.5155,  0.8443,  0.1865, -0.1752, -2.0595,
//              0.1850, -1.1571, -0.4243])
//     >>> torch.cumsum(a, dim=0)
//     tensor([-0.8286, -1.3175, -0.8020,  0.0423,  0.2289,  0.0537, -2.0058,
//             -1.8209, -2.9780, -3.4022])
//
//
//go:linkname Cumsum py.cumsum
func Cumsum(input *py.Object, dim *py.Object) *py.Object
//
// cumulative_trapezoid(y, x=None, *, dx=None, dim=-1) -> Tensor
//
// Cumulatively computes the `trapezoidal rule <https://en.wikipedia.org/wiki/Trapezoidal_rule>`_
// along :attr:`dim`. By default the spacing between elements is assumed to be 1, but
// :attr:`dx` can be used to specify a different constant spacing, and :attr:`x` can be
// used to specify arbitrary spacing along :attr:`dim`.
//
// For more details, please read :func:`torch.trapezoid`. The difference between :func:`torch.trapezoid`
// and this function is that, :func:`torch.trapezoid` returns a value for each integration,
// where as this function returns a cumulative value for every spacing within the integration. This
// is analogous to how `.sum` returns a value and `.cumsum` returns a cumulative sum.
//
// Arguments:
//     y (Tensor): Values to use when computing the trapezoidal rule.
//     x (Tensor): If specified, defines spacing between values as specified above.
//
// Keyword arguments:
//     dx (float): constant spacing between values. If neither :attr:`x` or :attr:`dx`
//         are specified then this defaults to 1. Effectively multiplies the result by its value.
//     dim (int): The dimension along which to compute the trapezoidal rule.
//         The last (inner-most) dimension by default.
//
// Examples::
//
//     >>> # Cumulatively computes the trapezoidal rule in 1D, spacing is implicitly 1.
//     >>> y = torch.tensor([1, 5, 10])
//     >>> torch.cumulative_trapezoid(y)
//     tensor([3., 10.5])
//
//     >>> # Computes the same trapezoidal rule directly up to each element to verify
//     >>> (1 + 5) / 2
//     3.0
//     >>> (1 + 10 + 10) / 2
//     10.5
//
//     >>> # Cumulatively computes the trapezoidal rule in 1D with constant spacing of 2
//     >>> # NOTE: the result is the same as before, but multiplied by 2
//     >>> torch.cumulative_trapezoid(y, dx=2)
//     tensor([6., 21.])
//
//     >>> # Cumulatively computes the trapezoidal rule in 1D with arbitrary spacing
//     >>> x = torch.tensor([1, 3, 6])
//     >>> torch.cumulative_trapezoid(y, x)
//     tensor([6., 28.5])
//
//     >>> # Computes the same trapezoidal rule directly up to each element to verify
//     >>> ((3 - 1) * (1 + 5)) / 2
//     6.0
//     >>> ((3 - 1) * (1 + 5) + (6 - 3) * (5 + 10)) / 2
//     28.5
//
//     >>> # Cumulatively computes the trapezoidal rule for each row of a 3x3 matrix
//     >>> y = torch.arange(9).reshape(3, 3)
//     tensor([[0, 1, 2],
//             [3, 4, 5],
//             [6, 7, 8]])
//     >>> torch.cumulative_trapezoid(y)
//     tensor([[ 0.5,  2.],
//             [ 3.5,  8.],
//             [ 6.5, 14.]])
//
//     >>> # Cumulatively computes the trapezoidal rule for each column of the matrix
//     >>> torch.cumulative_trapezoid(y, dim=0)
//     tensor([[ 1.5,  2.5,  3.5],
//             [ 6.0,  8.0, 10.0]])
//
//     >>> # Cumulatively computes the trapezoidal rule for each row of a 3x3 ones matrix
//     >>> #   with the same arbitrary spacing
//     >>> y = torch.ones(3, 3)
//     >>> x = torch.tensor([1, 3, 6])
//     >>> torch.cumulative_trapezoid(y, x)
//     tensor([[2., 5.],
//             [2., 5.],
//             [2., 5.]])
//
//     >>> # Cumulatively computes the trapezoidal rule for each row of a 3x3 ones matrix
//     >>> #   with different arbitrary spacing per row
//     >>> y = torch.ones(3, 3)
//     >>> x = torch.tensor([[1, 2, 3], [1, 3, 5], [1, 4, 7]])
//     >>> torch.cumulative_trapezoid(y, x)
//     tensor([[1., 2.],
//             [2., 4.],
//             [3., 6.]])
//
//
//go:linkname CumulativeTrapezoid py.cumulative_trapezoid
func CumulativeTrapezoid(y *py.Object, x *py.Object) *py.Object
//
// deg2rad(input, *, out=None) -> Tensor
//
// Returns a new tensor with each of the elements of :attr:`input`
// converted from angles in degrees to radians.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword arguments:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([[180.0, -180.0], [360.0, -360.0], [90.0, -90.0]])
//     >>> torch.deg2rad(a)
//     tensor([[ 3.1416, -3.1416],
//             [ 6.2832, -6.2832],
//             [ 1.5708, -1.5708]])
//
//
//
//go:linkname Deg2rad py.deg2rad
func Deg2rad(input *py.Object) *py.Object
// None
//
//go:linkname Deg2rad_ py.deg2rad_
func Deg2rad_(__llgo_va_list ...interface{}) *py.Object
//
// dequantize(tensor) -> Tensor
//
// Returns an fp32 Tensor by dequantizing a quantized Tensor
//
// Args:
//     tensor (Tensor): A quantized Tensor
//
// .. function:: dequantize(tensors) -> sequence of Tensors
//    :noindex:
//
// Given a list of quantized Tensors, dequantize them and return a list of fp32 Tensors
//
// Args:
//      tensors (sequence of Tensors): A list of quantized Tensors
//
//
//go:linkname Dequantize py.dequantize
func Dequantize(tensor *py.Object) *py.Object
//
// det(input) -> Tensor
//
// Alias for :func:`torch.linalg.det`
//
//
//go:linkname Det py.det
func Det(input *py.Object) *py.Object
// None
//
//go:linkname Detach py.detach
func Detach(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Detach_ py.detach_
func Detach_(__llgo_va_list ...interface{}) *py.Object
//
// Performs the same operation as :func:`torch.detach`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname DetachCopy py.detach_copy
func DetachCopy(__llgo_va_list ...interface{}) *py.Object
//
// diag(input, diagonal=0, *, out=None) -> Tensor
//
// - If :attr:`input` is a vector (1-D tensor), then returns a 2-D square tensor
//   with the elements of :attr:`input` as the diagonal.
// - If :attr:`input` is a matrix (2-D tensor), then returns a 1-D tensor with
//   the diagonal elements of :attr:`input`.
//
// The argument :attr:`diagonal` controls which diagonal to consider:
//
// - If :attr:`diagonal` = 0, it is the main diagonal.
// - If :attr:`diagonal` > 0, it is above the main diagonal.
// - If :attr:`diagonal` < 0, it is below the main diagonal.
//
// Args:
//     input (Tensor): the input tensor.
//     diagonal (int, optional): the diagonal to consider
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// .. seealso::
//
//         :func:`torch.diagonal` always returns the diagonal of its input.
//
//         :func:`torch.diagflat` always constructs a tensor with diagonal elements
//         specified by the input.
//
// Examples:
//
// Get the square matrix where the input vector is the diagonal::
//
//     >>> a = torch.randn(3)
//     >>> a
//     tensor([ 0.5950,-0.0872, 2.3298])
//     >>> torch.diag(a)
//     tensor([[ 0.5950, 0.0000, 0.0000],
//             [ 0.0000,-0.0872, 0.0000],
//             [ 0.0000, 0.0000, 2.3298]])
//     >>> torch.diag(a, 1)
//     tensor([[ 0.0000, 0.5950, 0.0000, 0.0000],
//             [ 0.0000, 0.0000,-0.0872, 0.0000],
//             [ 0.0000, 0.0000, 0.0000, 2.3298],
//             [ 0.0000, 0.0000, 0.0000, 0.0000]])
//
// Get the k-th diagonal of a given matrix::
//
//     >>> a = torch.randn(3, 3)
//     >>> a
//     tensor([[-0.4264, 0.0255,-0.1064],
//             [ 0.8795,-0.2429, 0.1374],
//             [ 0.1029,-0.6482,-1.6300]])
//     >>> torch.diag(a, 0)
//     tensor([-0.4264,-0.2429,-1.6300])
//     >>> torch.diag(a, 1)
//     tensor([ 0.0255, 0.1374])
//
//
//go:linkname Diag py.diag
func Diag(input *py.Object, diagonal *py.Object) *py.Object
//
// diag_embed(input, offset=0, dim1=-2, dim2=-1) -> Tensor
//
// Creates a tensor whose diagonals of certain 2D planes (specified by
// :attr:`dim1` and :attr:`dim2`) are filled by :attr:`input`.
// To facilitate creating batched diagonal matrices, the 2D planes formed by
// the last two dimensions of the returned tensor are chosen by default.
//
// The argument :attr:`offset` controls which diagonal to consider:
//
// - If :attr:`offset` = 0, it is the main diagonal.
// - If :attr:`offset` > 0, it is above the main diagonal.
// - If :attr:`offset` < 0, it is below the main diagonal.
//
// The size of the new matrix will be calculated to make the specified diagonal
// of the size of the last input dimension.
// Note that for :attr:`offset` other than :math:`0`, the order of :attr:`dim1`
// and :attr:`dim2` matters. Exchanging them is equivalent to changing the
// sign of :attr:`offset`.
//
// Applying :meth:`torch.diagonal` to the output of this function with
// the same arguments yields a matrix identical to input. However,
// :meth:`torch.diagonal` has different default dimensions, so those
// need to be explicitly specified.
//
// Args:
//     input (Tensor): the input tensor. Must be at least 1-dimensional.
//     offset (int, optional): which diagonal to consider. Default: 0
//         (main diagonal).
//     dim1 (int, optional): first dimension with respect to which to
//         take diagonal. Default: -2.
//     dim2 (int, optional): second dimension with respect to which to
//         take diagonal. Default: -1.
//
// Example::
//
//     >>> a = torch.randn(2, 3)
//     >>> torch.diag_embed(a)
//     tensor([[[ 1.5410,  0.0000,  0.0000],
//              [ 0.0000, -0.2934,  0.0000],
//              [ 0.0000,  0.0000, -2.1788]],
//
//             [[ 0.5684,  0.0000,  0.0000],
//              [ 0.0000, -1.0845,  0.0000],
//              [ 0.0000,  0.0000, -1.3986]]])
//
//     >>> torch.diag_embed(a, offset=1, dim1=0, dim2=2)
//     tensor([[[ 0.0000,  1.5410,  0.0000,  0.0000],
//              [ 0.0000,  0.5684,  0.0000,  0.0000]],
//
//             [[ 0.0000,  0.0000, -0.2934,  0.0000],
//              [ 0.0000,  0.0000, -1.0845,  0.0000]],
//
//             [[ 0.0000,  0.0000,  0.0000, -2.1788],
//              [ 0.0000,  0.0000,  0.0000, -1.3986]],
//
//             [[ 0.0000,  0.0000,  0.0000,  0.0000],
//              [ 0.0000,  0.0000,  0.0000,  0.0000]]])
//
//
//go:linkname DiagEmbed py.diag_embed
func DiagEmbed(input *py.Object, offset *py.Object, dim1 *py.Object, dim2 *py.Object) *py.Object
//
// diagflat(input, offset=0) -> Tensor
//
// - If :attr:`input` is a vector (1-D tensor), then returns a 2-D square tensor
//   with the elements of :attr:`input` as the diagonal.
// - If :attr:`input` is a tensor with more than one dimension, then returns a
//   2-D tensor with diagonal elements equal to a flattened :attr:`input`.
//
// The argument :attr:`offset` controls which diagonal to consider:
//
// - If :attr:`offset` = 0, it is the main diagonal.
// - If :attr:`offset` > 0, it is above the main diagonal.
// - If :attr:`offset` < 0, it is below the main diagonal.
//
// Args:
//     input (Tensor): the input tensor.
//     offset (int, optional): the diagonal to consider. Default: 0 (main
//         diagonal).
//
// Examples::
//
//     >>> a = torch.randn(3)
//     >>> a
//     tensor([-0.2956, -0.9068,  0.1695])
//     >>> torch.diagflat(a)
//     tensor([[-0.2956,  0.0000,  0.0000],
//             [ 0.0000, -0.9068,  0.0000],
//             [ 0.0000,  0.0000,  0.1695]])
//     >>> torch.diagflat(a, 1)
//     tensor([[ 0.0000, -0.2956,  0.0000,  0.0000],
//             [ 0.0000,  0.0000, -0.9068,  0.0000],
//             [ 0.0000,  0.0000,  0.0000,  0.1695],
//             [ 0.0000,  0.0000,  0.0000,  0.0000]])
//
//     >>> a = torch.randn(2, 2)
//     >>> a
//     tensor([[ 0.2094, -0.3018],
//             [-0.1516,  1.9342]])
//     >>> torch.diagflat(a)
//     tensor([[ 0.2094,  0.0000,  0.0000,  0.0000],
//             [ 0.0000, -0.3018,  0.0000,  0.0000],
//             [ 0.0000,  0.0000, -0.1516,  0.0000],
//             [ 0.0000,  0.0000,  0.0000,  1.9342]])
//
//
//go:linkname Diagflat py.diagflat
func Diagflat(input *py.Object, offset *py.Object) *py.Object
//
// diagonal(input, offset=0, dim1=0, dim2=1) -> Tensor
//
// Returns a partial view of :attr:`input` with the its diagonal elements
// with respect to :attr:`dim1` and :attr:`dim2` appended as a dimension
// at the end of the shape.
//
// The argument :attr:`offset` controls which diagonal to consider:
//
// - If :attr:`offset` = 0, it is the main diagonal.
// - If :attr:`offset` > 0, it is above the main diagonal.
// - If :attr:`offset` < 0, it is below the main diagonal.
//
// Applying :meth:`torch.diag_embed` to the output of this function with
// the same arguments yields a diagonal matrix with the diagonal entries
// of the input. However, :meth:`torch.diag_embed` has different default
// dimensions, so those need to be explicitly specified.
//
// Args:
//     input (Tensor): the input tensor. Must be at least 2-dimensional.
//     offset (int, optional): which diagonal to consider. Default: 0
//         (main diagonal).
//     dim1 (int, optional): first dimension with respect to which to
//         take diagonal. Default: 0.
//     dim2 (int, optional): second dimension with respect to which to
//         take diagonal. Default: 1.
//
// .. note::  To take a batch diagonal, pass in dim1=-2, dim2=-1.
//
// Examples::
//
//     >>> a = torch.randn(3, 3)
//     >>> a
//     tensor([[-1.0854,  1.1431, -0.1752],
//             [ 0.8536, -0.0905,  0.0360],
//             [ 0.6927, -0.3735, -0.4945]])
//
//
//     >>> torch.diagonal(a, 0)
//     tensor([-1.0854, -0.0905, -0.4945])
//
//
//     >>> torch.diagonal(a, 1)
//     tensor([ 1.1431,  0.0360])
//
//
//     >>> x = torch.randn(2, 5, 4, 2)
//     >>> torch.diagonal(x, offset=-1, dim1=1, dim2=2)
//     tensor([[[-1.2631,  0.3755, -1.5977, -1.8172],
//              [-1.1065,  1.0401, -0.2235, -0.7938]],
//
//             [[-1.7325, -0.3081,  0.6166,  0.2335],
//              [ 1.0500,  0.7336, -0.3836, -1.1015]]])
//
//
//go:linkname Diagonal py.diagonal
func Diagonal(input *py.Object, offset *py.Object, dim1 *py.Object, dim2 *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.diagonal`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname DiagonalCopy py.diagonal_copy
func DiagonalCopy(__llgo_va_list ...interface{}) *py.Object
//
// diagonal_scatter(input, src, offset=0, dim1=0, dim2=1) -> Tensor
//
// Embeds the values of the :attr:`src` tensor into :attr:`input` along
// the diagonal elements of :attr:`input`, with respect to :attr:`dim1`
// and :attr:`dim2`.
//
// This function returns a tensor with fresh storage; it does not
// return a view.
//
// The argument :attr:`offset` controls which diagonal to consider:
//
// - If :attr:`offset` = 0, it is the main diagonal.
// - If :attr:`offset` > 0, it is above the main diagonal.
// - If :attr:`offset` < 0, it is below the main diagonal.
//
// Args:
//     input (Tensor): the input tensor. Must be at least 2-dimensional.
//     src (Tensor): the tensor to embed into :attr:`input`.
//     offset (int, optional): which diagonal to consider. Default: 0
//         (main diagonal).
//     dim1 (int, optional): first dimension with respect to which to
//         take diagonal. Default: 0.
//     dim2 (int, optional): second dimension with respect to which to
//         take diagonal. Default: 1.
//
// .. note::
//
//     :attr:`src` must be of the proper size in order to be embedded
//     into :attr:`input`. Specifically, it should have the same shape as
//     ``torch.diagonal(input, offset, dim1, dim2)``
//
// Examples::
//
//     >>> a = torch.zeros(3, 3)
//     >>> a
//     tensor([[0., 0., 0.],
//             [0., 0., 0.],
//             [0., 0., 0.]])
//
//     >>> torch.diagonal_scatter(a, torch.ones(3), 0)
//     tensor([[1., 0., 0.],
//             [0., 1., 0.],
//             [0., 0., 1.]])
//
//     >>> torch.diagonal_scatter(a, torch.ones(2), 1)
//     tensor([[0., 1., 0.],
//             [0., 0., 1.],
//             [0., 0., 0.]])
//
//
//go:linkname DiagonalScatter py.diagonal_scatter
func DiagonalScatter(input *py.Object, src *py.Object, offset *py.Object, dim1 *py.Object, dim2 *py.Object) *py.Object
//
// diff(input, n=1, dim=-1, prepend=None, append=None) -> Tensor
//
// Computes the n-th forward difference along the given dimension.
//
// The first-order differences are given by `out[i] = input[i + 1] - input[i]`. Higher-order
// differences are calculated by using :func:`torch.diff` recursively.
//
// Args:
//     input (Tensor): the tensor to compute the differences on
//     n (int, optional): the number of times to recursively compute the difference
//     dim (int, optional): the dimension to compute the difference along.
//         Default is the last dimension.
//     prepend, append (Tensor, optional): values to prepend or append to
//         :attr:`input` along :attr:`dim` before computing the difference.
//         Their dimensions must be equivalent to that of input, and their shapes
//         must match input's shape except on :attr:`dim`.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([1, 3, 2])
//     >>> torch.diff(a)
//     tensor([ 2, -1])
//     >>> b = torch.tensor([4, 5])
//     >>> torch.diff(a, append=b)
//     tensor([ 2, -1,  2,  1])
//     >>> c = torch.tensor([[1, 2, 3], [3, 4, 5]])
//     >>> torch.diff(c, dim=0)
//     tensor([[2, 2, 2]])
//     >>> torch.diff(c, dim=1)
//     tensor([[1, 1],
//             [1, 1]])
//
//
//go:linkname Diff py.diff
func Diff(input *py.Object, n *py.Object, dim *py.Object, prepend *py.Object, append *py.Object) *py.Object
//
// digamma(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.digamma`.
//
//
//go:linkname Digamma py.digamma
func Digamma(input *py.Object) *py.Object
//
// dist(input, other, p=2) -> Tensor
//
// Returns the p-norm of (:attr:`input` - :attr:`other`)
//
// The shapes of :attr:`input` and :attr:`other` must be
// :ref:`broadcastable <broadcasting-semantics>`.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor): the Right-hand-side input tensor
//     p (float, optional): the norm to be computed
//
// Example::
//
//     >>> x = torch.randn(4)
//     >>> x
//     tensor([-1.5393, -0.8675,  0.5916,  1.6321])
//     >>> y = torch.randn(4)
//     >>> y
//     tensor([ 0.0967, -1.0511,  0.6295,  0.8360])
//     >>> torch.dist(x, y, 3.5)
//     tensor(1.6727)
//     >>> torch.dist(x, y, 3)
//     tensor(1.6973)
//     >>> torch.dist(x, y, 0)
//     tensor(4.)
//     >>> torch.dist(x, y, 1)
//     tensor(2.6537)
//
//
//go:linkname Dist py.dist
func Dist(input *py.Object, other *py.Object, p *py.Object) *py.Object
//
// div(input, other, *, rounding_mode=None, out=None) -> Tensor
//
// Divides each element of the input ``input`` by the corresponding element of
// :attr:`other`.
//
// .. math::
//     \text{out}_i = \frac{\text{input}_i}{\text{other}_i}
//
// .. note::
//     By default, this performs a "true" division like Python 3.
//     See the :attr:`rounding_mode` argument for floor division.
//
// Supports :ref:`broadcasting to a common shape <broadcasting-semantics>`,
// :ref:`type promotion <type-promotion-doc>`, and integer, float, and complex inputs.
// Always promotes integer types to the default scalar type.
//
// Args:
//     input (Tensor): the dividend
//     other (Tensor or Number): the divisor
//
// Keyword args:
//     rounding_mode (str, optional): Type of rounding applied to the result:
//
//         * None - default behavior. Performs no rounding and, if both :attr:`input` and
//           :attr:`other` are integer types, promotes the inputs to the default scalar type.
//           Equivalent to true division in Python (the ``/`` operator) and NumPy's ``np.true_divide``.
//         * ``"trunc"`` - rounds the results of the division towards zero.
//           Equivalent to C-style integer division.
//         * ``"floor"`` - rounds the results of the division down.
//           Equivalent to floor division in Python (the ``//`` operator) and NumPy's ``np.floor_divide``.
//
//     out (Tensor, optional): the output tensor.
//
// Examples::
//
//     >>> x = torch.tensor([ 0.3810,  1.2774, -0.2972, -0.3719,  0.4637])
//     >>> torch.div(x, 0.5)
//     tensor([ 0.7620,  2.5548, -0.5944, -0.7438,  0.9274])
//
//     >>> a = torch.tensor([[-0.3711, -1.9353, -0.4605, -0.2917],
//     ...                   [ 0.1815, -1.0111,  0.9805, -1.5923],
//     ...                   [ 0.1062,  1.4581,  0.7759, -1.2344],
//     ...                   [-0.1830, -0.0313,  1.1908, -1.4757]])
//     >>> b = torch.tensor([ 0.8032,  0.2930, -0.8113, -0.2308])
//     >>> torch.div(a, b)
//     tensor([[-0.4620, -6.6051,  0.5676,  1.2639],
//             [ 0.2260, -3.4509, -1.2086,  6.8990],
//             [ 0.1322,  4.9764, -0.9564,  5.3484],
//             [-0.2278, -0.1068, -1.4678,  6.3938]])
//
//     >>> torch.div(a, b, rounding_mode='trunc')
//     tensor([[-0., -6.,  0.,  1.],
//             [ 0., -3., -1.,  6.],
//             [ 0.,  4., -0.,  5.],
//             [-0., -0., -1.,  6.]])
//
//     >>> torch.div(a, b, rounding_mode='floor')
//     tensor([[-1., -7.,  0.,  1.],
//             [ 0., -4., -2.,  6.],
//             [ 0.,  4., -1.,  5.],
//             [-1., -1., -2.,  6.]])
//
//
//
//go:linkname Div py.div
func Div(input *py.Object, other *py.Object) *py.Object
//
// divide(input, other, *, rounding_mode=None, out=None) -> Tensor
//
// Alias for :func:`torch.div`.
//
//
//go:linkname Divide py.divide
func Divide(input *py.Object, other *py.Object) *py.Object
//
// dot(input, other, *, out=None) -> Tensor
//
// Computes the dot product of two 1D tensors.
//
// .. note::
//
//     Unlike NumPy's dot, torch.dot intentionally only supports computing the dot product
//     of two 1D tensors with the same number of elements.
//
// Args:
//     input (Tensor): first tensor in the dot product, must be 1D.
//     other (Tensor): second tensor in the dot product, must be 1D.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.dot(torch.tensor([2, 3]), torch.tensor([2, 1]))
//     tensor(7)
//
//
//go:linkname Dot py.dot
func Dot(input *py.Object, other *py.Object) *py.Object
// None
//
//go:linkname Dropout py.dropout
func Dropout(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Dropout_ py.dropout_
func Dropout_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Dsmm py.dsmm
func Dsmm(__llgo_va_list ...interface{}) *py.Object
//
// dsplit(input, indices_or_sections) -> List of Tensors
//
// Splits :attr:`input`, a tensor with three or more dimensions, into multiple tensors
// depthwise according to :attr:`indices_or_sections`. Each split is a view of
// :attr:`input`.
//
// This is equivalent to calling torch.tensor_split(input, indices_or_sections, dim=2)
// (the split dimension is 2), except that if :attr:`indices_or_sections` is an integer
// it must evenly divide the split dimension or a runtime error will be thrown.
//
// This function is based on NumPy's :func:`numpy.dsplit`.
//
// Args:
//     input (Tensor): tensor to split.
//     indices_or_sections (int or list or tuple of ints): See argument in :func:`torch.tensor_split`.
//
// Example::
//     >>> t = torch.arange(16.0).reshape(2, 2, 4)
//     >>> t
//     tensor([[[ 0.,  1.,  2.,  3.],
//              [ 4.,  5.,  6.,  7.]],
//             [[ 8.,  9., 10., 11.],
//              [12., 13., 14., 15.]]])
//     >>> torch.dsplit(t, 2)
//     (tensor([[[ 0.,  1.],
//             [ 4.,  5.]],
//            [[ 8.,  9.],
//             [12., 13.]]]),
//      tensor([[[ 2.,  3.],
//               [ 6.,  7.]],
//              [[10., 11.],
//               [14., 15.]]]))
//
//     >>> torch.dsplit(t, [3, 6])
//     (tensor([[[ 0.,  1.,  2.],
//               [ 4.,  5.,  6.]],
//              [[ 8.,  9., 10.],
//               [12., 13., 14.]]]),
//      tensor([[[ 3.],
//               [ 7.]],
//              [[11.],
//               [15.]]]),
//      tensor([], size=(2, 2, 0)))
//
//
//
//go:linkname Dsplit py.dsplit
func Dsplit(input *py.Object, indicesOrSections *py.Object) *py.Object
//
// dstack(tensors, *, out=None) -> Tensor
//
// Stack tensors in sequence depthwise (along third axis).
//
// This is equivalent to concatenation along the third axis after 1-D and 2-D tensors have been reshaped by :func:`torch.atleast_3d`.
//
// Args:
//     tensors (sequence of Tensors): sequence of tensors to concatenate
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([1, 2, 3])
//     >>> b = torch.tensor([4, 5, 6])
//     >>> torch.dstack((a,b))
//     tensor([[[1, 4],
//              [2, 5],
//              [3, 6]]])
//     >>> a = torch.tensor([[1],[2],[3]])
//     >>> b = torch.tensor([[4],[5],[6]])
//     >>> torch.dstack((a,b))
//     tensor([[[1, 4]],
//             [[2, 5]],
//             [[3, 6]]])
//
//
//
//
//go:linkname Dstack py.dstack
func Dstack(tensors *py.Object) *py.Object
// einsum(equation, *operands) -> Tensor
//
//     Sums the product of the elements of the input :attr:`operands` along dimensions specified using a notation
//     based on the Einstein summation convention.
//
//     Einsum allows computing many common multi-dimensional linear algebraic array operations by representing them
//     in a short-hand format based on the Einstein summation convention, given by :attr:`equation`. The details of
//     this format are described below, but the general idea is to label every dimension of the input :attr:`operands`
//     with some subscript and define which subscripts are part of the output. The output is then computed by summing
//     the product of the elements of the :attr:`operands` along the dimensions whose subscripts are not part of the
//     output. For example, matrix multiplication can be computed using einsum as `torch.einsum("ij,jk->ik", A, B)`.
//     Here, j is the summation subscript and i and k the output subscripts (see section below for more details on why).
//
//     Equation:
//
//         The :attr:`equation` string specifies the subscripts (letters in `[a-zA-Z]`) for each dimension of
//         the input :attr:`operands` in the same order as the dimensions, separating subscripts for each operand by a
//         comma (','), e.g. `'ij,jk'` specify subscripts for two 2D operands. The dimensions labeled with the same subscript
//         must be broadcastable, that is, their size must either match or be `1`. The exception is if a subscript is
//         repeated for the same input operand, in which case the dimensions labeled with this subscript for this operand
//         must match in size and the operand will be replaced by its diagonal along these dimensions. The subscripts that
//         appear exactly once in the :attr:`equation` will be part of the output, sorted in increasing alphabetical order.
//         The output is computed by multiplying the input :attr:`operands` element-wise, with their dimensions aligned based
//         on the subscripts, and then summing out the dimensions whose subscripts are not part of the output.
//
//         Optionally, the output subscripts can be explicitly defined by adding an arrow ('->') at the end of the equation
//         followed by the subscripts for the output. For instance, the following equation computes the transpose of a
//         matrix multiplication: 'ij,jk->ki'. The output subscripts must appear at least once for some input operand and
//         at most once for the output.
//
//         Ellipsis ('...') can be used in place of subscripts to broadcast the dimensions covered by the ellipsis.
//         Each input operand may contain at most one ellipsis which will cover the dimensions not covered by subscripts,
//         e.g. for an input operand with 5 dimensions, the ellipsis in the equation `'ab...c'` cover the third and fourth
//         dimensions. The ellipsis does not need to cover the same number of dimensions across the :attr:`operands` but the
//         'shape' of the ellipsis (the size of the dimensions covered by them) must broadcast together. If the output is not
//         explicitly defined with the arrow ('->') notation, the ellipsis will come first in the output (left-most dimensions),
//         before the subscript labels that appear exactly once for the input operands. e.g. the following equation implements
//         batch matrix multiplication `'...ij,...jk'`.
//
//         A few final notes: the equation may contain whitespaces between the different elements (subscripts, ellipsis,
//         arrow and comma) but something like `'. . .'` is not valid. An empty string `''` is valid for scalar operands.
//
//     .. note::
//
//         ``torch.einsum`` handles ellipsis ('...') differently from NumPy in that it allows dimensions
//         covered by the ellipsis to be summed over, that is, ellipsis are not required to be part of the output.
//
//     .. note::
//
//         This function uses opt_einsum (https://optimized-einsum.readthedocs.io/en/stable/) to speed up computation or to
//         consume less memory by optimizing contraction order. This optimization occurs when there are at least three
//         inputs, since the order does not matter otherwise. Note that finding _the_ optimal path is an NP-hard problem,
//         thus, opt_einsum relies on different heuristics to achieve near-optimal results. If opt_einsum is not available,
//         the default order is to contract from left to right.
//
//         To bypass this default behavior, add the following line to disable the usage of opt_einsum and skip path
//         calculation: `torch.backends.opt_einsum.enabled = False`
//
//         To specify which strategy you'd like for opt_einsum to compute the contraction path, add the following line:
//         `torch.backends.opt_einsum.strategy = 'auto'`. The default strategy is 'auto', and we also support 'greedy' and
//         'optimal'. Disclaimer that the runtime of 'optimal' is factorial in the number of inputs! See more details in
//         the opt_einsum documentation (https://optimized-einsum.readthedocs.io/en/stable/path_finding.html).
//
//     .. note::
//
//         As of PyTorch 1.10 :func:`torch.einsum` also supports the sublist format (see examples below). In this format,
//         subscripts for each operand are specified by sublists, list of integers in the range [0, 52). These sublists
//         follow their operands, and an extra sublist can appear at the end of the input to specify the output's
//         subscripts., e.g. `torch.einsum(op1, sublist1, op2, sublist2, ..., [subslist_out])`. Python's `Ellipsis` object
//         may be provided in a sublist to enable broadcasting as described in the Equation section above.
//
//     Args:
//         equation (str): The subscripts for the Einstein summation.
//         operands (List[Tensor]): The tensors to compute the Einstein summation of.
//
//     Examples::
//
//         >>> # xdoctest: +IGNORE_WANT("non-deterministic")
//         >>> # trace
//         >>> torch.einsum('ii', torch.randn(4, 4))
//         tensor(-1.2104)
//
//         >>> # xdoctest: +IGNORE_WANT("non-deterministic")
//         >>> # diagonal
//         >>> torch.einsum('ii->i', torch.randn(4, 4))
//         tensor([-0.1034,  0.7952, -0.2433,  0.4545])
//
//         >>> # xdoctest: +IGNORE_WANT("non-deterministic")
//         >>> # outer product
//         >>> x = torch.randn(5)
//         >>> y = torch.randn(4)
//         >>> torch.einsum('i,j->ij', x, y)
//         tensor([[ 0.1156, -0.2897, -0.3918,  0.4963],
//                 [-0.3744,  0.9381,  1.2685, -1.6070],
//                 [ 0.7208, -1.8058, -2.4419,  3.0936],
//                 [ 0.1713, -0.4291, -0.5802,  0.7350],
//                 [ 0.5704, -1.4290, -1.9323,  2.4480]])
//
//         >>> # xdoctest: +IGNORE_WANT("non-deterministic")
//         >>> # batch matrix multiplication
//         >>> As = torch.randn(3, 2, 5)
//         >>> Bs = torch.randn(3, 5, 4)
//         >>> torch.einsum('bij,bjk->bik', As, Bs)
//         tensor([[[-1.0564, -1.5904,  3.2023,  3.1271],
//                 [-1.6706, -0.8097, -0.8025, -2.1183]],
//
//                 [[ 4.2239,  0.3107, -0.5756, -0.2354],
//                 [-1.4558, -0.3460,  1.5087, -0.8530]],
//
//                 [[ 2.8153,  1.8787, -4.3839, -1.2112],
//                 [ 0.3728, -2.1131,  0.0921,  0.8305]]])
//
//         >>> # xdoctest: +IGNORE_WANT("non-deterministic")
//         >>> # with sublist format and ellipsis
//         >>> torch.einsum(As, [..., 0, 1], Bs, [..., 1, 2], [..., 0, 2])
//         tensor([[[-1.0564, -1.5904,  3.2023,  3.1271],
//                 [-1.6706, -0.8097, -0.8025, -2.1183]],
//
//                 [[ 4.2239,  0.3107, -0.5756, -0.2354],
//                 [-1.4558, -0.3460,  1.5087, -0.8530]],
//
//                 [[ 2.8153,  1.8787, -4.3839, -1.2112],
//                 [ 0.3728, -2.1131,  0.0921,  0.8305]]])
//
//         >>> # batch permute
//         >>> A = torch.randn(2, 3, 4, 5)
//         >>> torch.einsum('...ij->...ji', A).shape
//         torch.Size([2, 3, 5, 4])
//
//         >>> # equivalent to torch.nn.functional.bilinear
//         >>> A = torch.randn(3, 5, 4)
//         >>> l = torch.randn(2, 5)
//         >>> r = torch.randn(2, 4)
//         >>> torch.einsum('bn,anm,bm->ba', l, A, r)
//         tensor([[-0.3430, -5.2405,  0.4494],
//                 [ 0.3311,  5.5201, -3.0356]])
//
//
//go:linkname Einsum py.einsum
func Einsum(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Embedding py.embedding
func Embedding(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname EmbeddingBag py.embedding_bag
func EmbeddingBag(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname EmbeddingRenorm_ py.embedding_renorm_
func EmbeddingRenorm_(__llgo_va_list ...interface{}) *py.Object
//
// empty(*size, *, out=None, dtype=None, layout=torch.strided, device=None, requires_grad=False, pin_memory=False, memory_format=torch.contiguous_format) -> Tensor
//
// Returns a tensor filled with uninitialized data. The shape of the tensor is
// defined by the variable argument :attr:`size`.
//
// .. note::
//     If :func:`torch.use_deterministic_algorithms()` and
//     :attr:`torch.utils.deterministic.fill_uninitialized_memory` are both set to
//     ``True``, the output tensor is initialized to prevent any possible
//     nondeterministic behavior from using the data as an input to an operation.
//     Floating point and complex tensors are filled with NaN, and integer tensors
//     are filled with the maximum value.
//
// Args:
//     size (int...): a sequence of integers defining the shape of the output tensor.
//         Can be a variable number of arguments or a collection like a list or tuple.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`).
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     pin_memory (bool, optional): If set, returned tensor would be allocated in
//         the pinned memory. Works only for CPU tensors. Default: ``False``.
//     memory_format (:class:`torch.memory_format`, optional): the desired memory format of
//         returned Tensor. Default: ``torch.contiguous_format``.
//
// Example::
//
//     >>> torch.empty((2,3), dtype=torch.int64)
//     tensor([[ 9.4064e+13,  2.8000e+01,  9.3493e+13],
//             [ 7.5751e+18,  7.1428e+18,  7.5955e+18]])
//
//
//go:linkname Empty py.empty
func Empty(__llgo_va_list ...interface{}) *py.Object
//
// empty_like(input, *, dtype=None, layout=None, device=None, requires_grad=False, memory_format=torch.preserve_format) -> Tensor
//
// Returns an uninitialized tensor with the same size as :attr:`input`.
// ``torch.empty_like(input)`` is equivalent to
// ``torch.empty(input.size(), dtype=input.dtype, layout=input.layout, device=input.device)``.
//
// .. note::
//     If :func:`torch.use_deterministic_algorithms()` and
//     :attr:`torch.utils.deterministic.fill_uninitialized_memory` are both set to
//     ``True``, the output tensor is initialized to prevent any possible
//     nondeterministic behavior from using the data as an input to an operation.
//     Floating point and complex tensors are filled with NaN, and integer tensors
//     are filled with the maximum value.
//
// Args:
//     input (Tensor): the size of :attr:`input` will determine size of the output tensor.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned Tensor.
//         Default: if ``None``, defaults to the dtype of :attr:`input`.
//     layout (:class:`torch.layout`, optional): the desired layout of returned tensor.
//         Default: if ``None``, defaults to the layout of :attr:`input`.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, defaults to the device of :attr:`input`.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     memory_format (:class:`torch.memory_format`, optional): the desired memory format of
//         returned Tensor. Default: ``torch.preserve_format``.
//
// Example::
//
//     >>> a=torch.empty((2,3), dtype=torch.int32, device = 'cuda')
//     >>> torch.empty_like(a)
//     tensor([[0, 0, 0],
//             [0, 0, 0]], device='cuda:0', dtype=torch.int32)
//
//
//go:linkname EmptyLike py.empty_like
func EmptyLike(input *py.Object) *py.Object
//
// empty_permuted(size, physical_layout, *, dtype=None, layout=None, device=None, requires_grad=False, pin_memory=False) -> Tensor
//
// Creates an uninitialized, non-overlapping and dense tensor with the
// specified :attr:`size`, with :attr:`physical_layout` specifying how the
// dimensions are physically laid out in memory (each logical dimension is listed
// from outermost to innermost).  :attr:`physical_layout` is a generalization
// of NCHW/NHWC notation: if each dimension is assigned a number according to
// what order they occur in size (N=0, C=1, H=2, W=3), then NCHW is ``(0, 1, 2, 3)``
// while NHWC is ``(0, 2, 3, 1)``.  Equivalently, the strides of the output
// tensor ``t`` are such that ``t.stride(physical_layout[i]) == contiguous_strides[i]``
// (notably, this function is *not* equivalent to ``torch.empty(size).permute(physical_layout)``).
//
// Unlike :func:`torch.empty_strided`, this is guaranteed to produce a dense
// tensor with no overlaps.  If possible, prefer using this function over
// :func:`torch.empty_strided` or manual use of :func:`torch.as_strided`.
//
// .. note::
//     If :func:`torch.use_deterministic_algorithms()` and
//     :attr:`torch.utils.deterministic.fill_uninitialized_memory` are both set to
//     ``True``, the output tensor is initialized to prevent any possible
//     nondeterministic behavior from using the data as an input to an operation.
//     Floating point and complex tensors are filled with NaN, and integer tensors
//     are filled with the maximum value.
//
// Args:
//     size (tuple of int): the shape of the output tensor
//     physical_layout (tuple of int): the ordering of dimensions physically in memory
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`).
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     pin_memory (bool, optional): If set, returned tensor would be allocated in
//         the pinned memory. Works only for CPU tensors. Default: ``False``.
//
// Examples:
//
//     >>> torch.empty((2, 3, 5, 7)).stride()
//     (105, 35, 7, 1)
//     >>> torch.empty_permuted((2, 3, 5, 7), (0, 1, 2, 3)).stride()
//     (105, 35, 7, 1)
//     >>> torch.empty((2, 3, 5, 7), memory_format=torch.channels_last).stride()
//     (105, 1, 21, 3)
//     >>> torch.empty_permuted((2, 3, 5, 7), (0, 2, 3, 1)).stride()
//     (105, 1, 21, 3)
//     >>> torch.empty_permuted((2, 3, 5, 7), (0, 2, 3, 1)).dim_order()
//     (0, 2, 3, 1)
//
//
//go:linkname EmptyPermuted py.empty_permuted
func EmptyPermuted(size *py.Object, physicalLayout *py.Object) *py.Object
// None
//
//go:linkname EmptyQuantized py.empty_quantized
func EmptyQuantized(__llgo_va_list ...interface{}) *py.Object
//
// empty_strided(size, stride, *, dtype=None, layout=None, device=None, requires_grad=False, pin_memory=False) -> Tensor
//
// Creates a tensor with the specified :attr:`size` and :attr:`stride` and filled with undefined data.
//
// .. warning::
//     If the constructed tensor is "overlapped" (with multiple indices referring to the same element
//     in memory) its behavior is undefined.
//
// .. note::
//     If :func:`torch.use_deterministic_algorithms()` and
//     :attr:`torch.utils.deterministic.fill_uninitialized_memory` are both set to
//     ``True``, the output tensor is initialized to prevent any possible
//     nondeterministic behavior from using the data as an input to an operation.
//     Floating point and complex tensors are filled with NaN, and integer tensors
//     are filled with the maximum value.
//
// Args:
//     size (tuple of int): the shape of the output tensor
//     stride (tuple of int): the strides of the output tensor
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`).
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     pin_memory (bool, optional): If set, returned tensor would be allocated in
//         the pinned memory. Works only for CPU tensors. Default: ``False``.
//
// Example::
//
//     >>> a = torch.empty_strided((2, 3), (1, 2))
//     >>> a
//     tensor([[8.9683e-44, 4.4842e-44, 5.1239e+07],
//             [0.0000e+00, 0.0000e+00, 3.0705e-41]])
//     >>> a.stride()
//     (1, 2)
//     >>> a.size()
//     torch.Size([2, 3])
//
//
//go:linkname EmptyStrided py.empty_strided
func EmptyStrided(size *py.Object, stride *py.Object) *py.Object
//
// eq(input, other, *, out=None) -> Tensor
//
// Computes element-wise equality
//
// The second argument can be a number or a tensor whose shape is
// :ref:`broadcastable <broadcasting-semantics>` with the first argument.
//
// Args:
//     input (Tensor): the tensor to compare
//     other (Tensor or float): the tensor or value to compare
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Returns:
//     A boolean tensor that is True where :attr:`input` is equal to :attr:`other` and False elsewhere
//
// Example::
//
//     >>> torch.eq(torch.tensor([[1, 2], [3, 4]]), torch.tensor([[1, 1], [4, 4]]))
//     tensor([[ True, False],
//             [False, True]])
//
//
//go:linkname Eq py.eq
func Eq(input *py.Object, other *py.Object) *py.Object
//
// equal(input, other) -> bool
//
// ``True`` if two tensors have the same size and elements, ``False`` otherwise.
//
// Example::
//
//     >>> torch.equal(torch.tensor([1, 2]), torch.tensor([1, 2]))
//     True
//
//
//go:linkname Equal py.equal
func Equal(input *py.Object, other *py.Object) *py.Object
//
// erf(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.erf`.
//
//
//go:linkname Erf py.erf
func Erf(input *py.Object) *py.Object
// None
//
//go:linkname Erf_ py.erf_
func Erf_(__llgo_va_list ...interface{}) *py.Object
//
// erfc(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.erfc`.
//
//
//go:linkname Erfc py.erfc
func Erfc(input *py.Object) *py.Object
// None
//
//go:linkname Erfc_ py.erfc_
func Erfc_(__llgo_va_list ...interface{}) *py.Object
//
// erfinv(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.erfinv`.
//
//
//go:linkname Erfinv py.erfinv
func Erfinv(input *py.Object) *py.Object
//
// exp(input, *, out=None) -> Tensor
//
// Returns a new tensor with the exponential of the elements
// of the input tensor :attr:`input`.
//
// .. math::
//     y_{i} = e^{x_{i}}
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.exp(torch.tensor([0, math.log(2.)]))
//     tensor([ 1.,  2.])
//
//
//go:linkname Exp py.exp
func Exp(input *py.Object) *py.Object
//
// exp2(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.exp2`.
//
//
//go:linkname Exp2 py.exp2
func Exp2(input *py.Object) *py.Object
// None
//
//go:linkname Exp2_ py.exp2_
func Exp2_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Exp_ py.exp_
func Exp_(__llgo_va_list ...interface{}) *py.Object
//
// Performs the same operation as :func:`torch.expand`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname ExpandCopy py.expand_copy
func ExpandCopy(__llgo_va_list ...interface{}) *py.Object
//
// expm1(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.expm1`.
//
//
//go:linkname Expm1 py.expm1
func Expm1(input *py.Object) *py.Object
// None
//
//go:linkname Expm1_ py.expm1_
func Expm1_(__llgo_va_list ...interface{}) *py.Object
//
// eye(n, m=None, *, out=None, dtype=None, layout=torch.strided, device=None, requires_grad=False) -> Tensor
//
// Returns a 2-D tensor with ones on the diagonal and zeros elsewhere.
//
// Args:
//     n (int): the number of rows
//     m (int, optional): the number of columns with default being :attr:`n`
//
// Keyword arguments:
//     out (Tensor, optional): the output tensor.
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`).
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
// Returns:
//     Tensor: A 2-D tensor with ones on the diagonal and zeros elsewhere
//
// Example::
//
//     >>> torch.eye(3)
//     tensor([[ 1.,  0.,  0.],
//             [ 0.,  1.,  0.],
//             [ 0.,  0.,  1.]])
//
//
//go:linkname Eye py.eye
func Eye(n *py.Object, m *py.Object) *py.Object
//
// fake_quantize_per_channel_affine(input, scale, zero_point, axis, quant_min, quant_max) -> Tensor
//
// Returns a new tensor with the data in :attr:`input` fake quantized per channel using :attr:`scale`,
// :attr:`zero_point`, :attr:`quant_min` and :attr:`quant_max`, across the channel specified by :attr:`axis`.
//
// .. math::
//     \text{output} = (
//         min(
//             \text{quant\_max},
//             max(
//                 \text{quant\_min},
//                 \text{std::nearby\_int}(\text{input} / \text{scale}) + \text{zero\_point}
//             )
//         ) - \text{zero\_point}
//     ) \times \text{scale}
//
// Args:
//     input (Tensor): the input value(s), in ``torch.float32``
//     scale (Tensor): quantization scale, per channel in ``torch.float32``
//     zero_point (Tensor): quantization zero_point, per channel in ``torch.int32`` or ``torch.half`` or ``torch.float32``
//     axis (int32): channel axis
//     quant_min (int64): lower bound of the quantized domain
//     quant_max (int64): upper bound of the quantized domain
//
// Returns:
//     Tensor: A newly fake_quantized per channel ``torch.float32`` tensor
//
// Example::
//
//     >>> x = torch.randn(2, 2, 2)
//     >>> x
//     tensor([[[-0.2525, -0.0466],
//              [ 0.3491, -0.2168]],
//
//             [[-0.5906,  1.6258],
//              [ 0.6444, -0.0542]]])
//     >>> scales = (torch.randn(2) + 1) * 0.05
//     >>> scales
//     tensor([0.0475, 0.0486])
//     >>> zero_points = torch.zeros(2).to(torch.int32)
//     >>> zero_points
//     tensor([0, 0])
//     >>> torch.fake_quantize_per_channel_affine(x, scales, zero_points, 1, 0, 255)
//     tensor([[[0.0000, 0.0000],
//              [0.3405, 0.0000]],
//
//             [[0.0000, 1.6134],
//             [0.6323, 0.0000]]])
//
//
//go:linkname FakeQuantizePerChannelAffine py.fake_quantize_per_channel_affine
func FakeQuantizePerChannelAffine(input *py.Object, scale *py.Object, zeroPoint *py.Object, axis *py.Object, quantMin *py.Object, quantMax *py.Object) *py.Object
//
// fake_quantize_per_tensor_affine(input, scale, zero_point, quant_min, quant_max) -> Tensor
//
// Returns a new tensor with the data in :attr:`input` fake quantized using :attr:`scale`,
// :attr:`zero_point`, :attr:`quant_min` and :attr:`quant_max`.
//
// .. math::
//     \text{output} = (
//         min(
//             \text{quant\_max},
//             max(
//                 \text{quant\_min},
//                 \text{std::nearby\_int}(\text{input} / \text{scale}) + \text{zero\_point}
//             )
//         ) - \text{zero\_point}
//     ) \times \text{scale}
//
// Args:
//     input (Tensor): the input value(s), ``torch.float32`` tensor
//     scale (double scalar or ``float32`` Tensor): quantization scale
//     zero_point (int64 scalar or ``int32`` Tensor): quantization zero_point
//     quant_min (int64): lower bound of the quantized domain
//     quant_max (int64): upper bound of the quantized domain
//
// Returns:
//     Tensor: A newly fake_quantized ``torch.float32`` tensor
//
// Example::
//
//     >>> x = torch.randn(4)
//     >>> x
//     tensor([ 0.0552,  0.9730,  0.3973, -1.0780])
//     >>> torch.fake_quantize_per_tensor_affine(x, 0.1, 0, 0, 255)
//     tensor([0.1000, 1.0000, 0.4000, 0.0000])
//     >>> torch.fake_quantize_per_tensor_affine(x, torch.tensor(0.1), torch.tensor(0), 0, 255)
//     tensor([0.1000, 1.0000, 0.4000, 0.0000])
//
//
//go:linkname FakeQuantizePerTensorAffine py.fake_quantize_per_tensor_affine
func FakeQuantizePerTensorAffine(input *py.Object, scale *py.Object, zeroPoint *py.Object, quantMin *py.Object, quantMax *py.Object) *py.Object
// None
//
//go:linkname FbgemmLinearFp16Weight py.fbgemm_linear_fp16_weight
func FbgemmLinearFp16Weight(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname FbgemmLinearFp16WeightFp32Activation py.fbgemm_linear_fp16_weight_fp32_activation
func FbgemmLinearFp16WeightFp32Activation(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname FbgemmLinearInt8Weight py.fbgemm_linear_int8_weight
func FbgemmLinearInt8Weight(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname FbgemmLinearInt8WeightFp32Activation py.fbgemm_linear_int8_weight_fp32_activation
func FbgemmLinearInt8WeightFp32Activation(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname FbgemmLinearQuantizeWeight py.fbgemm_linear_quantize_weight
func FbgemmLinearQuantizeWeight(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname FbgemmPackGemmMatrixFp16 py.fbgemm_pack_gemm_matrix_fp16
func FbgemmPackGemmMatrixFp16(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname FbgemmPackQuantizedMatrix py.fbgemm_pack_quantized_matrix
func FbgemmPackQuantizedMatrix(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname FeatureAlphaDropout py.feature_alpha_dropout
func FeatureAlphaDropout(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname FeatureAlphaDropout_ py.feature_alpha_dropout_
func FeatureAlphaDropout_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname FeatureDropout py.feature_dropout
func FeatureDropout(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname FeatureDropout_ py.feature_dropout_
func FeatureDropout_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Fill py.fill
func Fill(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Fill_ py.fill_
func Fill_(__llgo_va_list ...interface{}) *py.Object
//
// fix(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.trunc`
//
//
//go:linkname Fix py.fix
func Fix(input *py.Object) *py.Object
// None
//
//go:linkname Fix_ py.fix_
func Fix_(__llgo_va_list ...interface{}) *py.Object
//
// flatten(input, start_dim=0, end_dim=-1) -> Tensor
//
// Flattens :attr:`input` by reshaping it into a one-dimensional tensor. If :attr:`start_dim` or :attr:`end_dim`
// are passed, only dimensions starting with :attr:`start_dim` and ending with :attr:`end_dim` are flattened.
// The order of elements in :attr:`input` is unchanged.
//
// Unlike NumPy's flatten, which always copies input's data, this function may return the original object, a view,
// or copy. If no dimensions are flattened, then the original object :attr:`input` is returned. Otherwise, if input can
// be viewed as the flattened shape, then that view is returned. Finally, only if the input cannot be viewed as the
// flattened shape is input's data copied. See :meth:`torch.Tensor.view` for details on when a view will be returned.
//
// .. note::
//     Flattening a zero-dimensional tensor will return a one-dimensional view.
//
// Args:
//     input (Tensor): the input tensor.
//     start_dim (int): the first dim to flatten
//     end_dim (int): the last dim to flatten
//
// Example::
//
//     >>> t = torch.tensor([[[1, 2],
//     ...                    [3, 4]],
//     ...                   [[5, 6],
//     ...                    [7, 8]]])
//     >>> torch.flatten(t)
//     tensor([1, 2, 3, 4, 5, 6, 7, 8])
//     >>> torch.flatten(t, start_dim=1)
//     tensor([[1, 2, 3, 4],
//             [5, 6, 7, 8]])
//
//
//go:linkname Flatten py.flatten
func Flatten(input *py.Object, startDim *py.Object, endDim *py.Object) *py.Object
//
// flip(input, dims) -> Tensor
//
// Reverse the order of an n-D tensor along given axis in dims.
//
// .. note::
//     `torch.flip` makes a copy of :attr:`input`'s data. This is different from NumPy's `np.flip`,
//     which returns a view in constant time. Since copying a tensor's data is more work than viewing that data,
//     `torch.flip` is expected to be slower than `np.flip`.
//
// Args:
//     input (Tensor): the input tensor.
//     dims (a list or tuple): axis to flip on
//
// Example::
//
//     >>> x = torch.arange(8).view(2, 2, 2)
//     >>> x
//     tensor([[[ 0,  1],
//              [ 2,  3]],
//
//             [[ 4,  5],
//              [ 6,  7]]])
//     >>> torch.flip(x, [0, 1])
//     tensor([[[ 6,  7],
//              [ 4,  5]],
//
//             [[ 2,  3],
//              [ 0,  1]]])
//
//
//go:linkname Flip py.flip
func Flip(input *py.Object, dims *py.Object) *py.Object
//
// fliplr(input) -> Tensor
//
// Flip tensor in the left/right direction, returning a new tensor.
//
// Flip the entries in each row in the left/right direction.
// Columns are preserved, but appear in a different order than before.
//
// Note:
//     Requires the tensor to be at least 2-D.
//
// .. note::
//     `torch.fliplr` makes a copy of :attr:`input`'s data. This is different from NumPy's `np.fliplr`,
//     which returns a view in constant time. Since copying a tensor's data is more work than viewing that data,
//     `torch.fliplr` is expected to be slower than `np.fliplr`.
//
// Args:
//     input (Tensor): Must be at least 2-dimensional.
//
// Example::
//
//     >>> x = torch.arange(4).view(2, 2)
//     >>> x
//     tensor([[0, 1],
//             [2, 3]])
//     >>> torch.fliplr(x)
//     tensor([[1, 0],
//             [3, 2]])
//
//
//go:linkname Fliplr py.fliplr
func Fliplr(input *py.Object) *py.Object
//
// flipud(input) -> Tensor
//
// Flip tensor in the up/down direction, returning a new tensor.
//
// Flip the entries in each column in the up/down direction.
// Rows are preserved, but appear in a different order than before.
//
// Note:
//     Requires the tensor to be at least 1-D.
//
// .. note::
//     `torch.flipud` makes a copy of :attr:`input`'s data. This is different from NumPy's `np.flipud`,
//     which returns a view in constant time. Since copying a tensor's data is more work than viewing that data,
//     `torch.flipud` is expected to be slower than `np.flipud`.
//
// Args:
//     input (Tensor): Must be at least 1-dimensional.
//
// Example::
//
//     >>> x = torch.arange(4).view(2, 2)
//     >>> x
//     tensor([[0, 1],
//             [2, 3]])
//     >>> torch.flipud(x)
//     tensor([[2, 3],
//             [0, 1]])
//
//
//go:linkname Flipud py.flipud
func Flipud(input *py.Object) *py.Object
//
// float_power(input, exponent, *, out=None) -> Tensor
//
// Raises :attr:`input` to the power of :attr:`exponent`, elementwise, in double precision.
// If neither input is complex returns a ``torch.float64`` tensor,
// and if one or more inputs is complex returns a ``torch.complex128`` tensor.
//
// .. note::
//     This function always computes in double precision, unlike :func:`torch.pow`,
//     which implements more typical :ref:`type promotion <type-promotion-doc>`.
//     This is useful when the computation needs to be performed in a wider or more precise dtype,
//     or the results of the computation may contain fractional values not representable in the input dtypes,
//     like when an integer base is raised to a negative integer exponent.
//
// Args:
//     input (Tensor or Number): the base value(s)
//     exponent (Tensor or Number): the exponent value(s)
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randint(10, (4,))
//     >>> a
//     tensor([6, 4, 7, 1])
//     >>> torch.float_power(a, 2)
//     tensor([36., 16., 49.,  1.], dtype=torch.float64)
//
//     >>> a = torch.arange(1, 5)
//     >>> a
//     tensor([ 1,  2,  3,  4])
//     >>> exp = torch.tensor([2, -3, 4, -5])
//     >>> exp
//     tensor([ 2, -3,  4, -5])
//     >>> torch.float_power(a, exp)
//     tensor([1.0000e+00, 1.2500e-01, 8.1000e+01, 9.7656e-04], dtype=torch.float64)
//
//
//go:linkname FloatPower py.float_power
func FloatPower(input *py.Object, exponent *py.Object) *py.Object
//
// floor(input, *, out=None) -> Tensor
//
// Returns a new tensor with the floor of the elements of :attr:`input`,
// the largest integer less than or equal to each element.
//
// For integer inputs, follows the array-api convention of returning a
// copy of the input tensor.
//
// .. math::
//     \text{out}_{i} = \left\lfloor \text{input}_{i} \right\rfloor
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([-0.8166,  1.5308, -0.2530, -0.2091])
//     >>> torch.floor(a)
//     tensor([-1.,  1., -1., -1.])
//
//
//go:linkname Floor py.floor
func Floor(input *py.Object) *py.Object
// None
//
//go:linkname Floor_ py.floor_
func Floor_(__llgo_va_list ...interface{}) *py.Object
//
// floor_divide(input, other, *, out=None) -> Tensor
//
// .. note::
//
//     Before PyTorch 1.13 :func:`torch.floor_divide` incorrectly performed
//     truncation division. To restore the previous behavior use
//     :func:`torch.div` with ``rounding_mode='trunc'``.
//
// Computes :attr:`input` divided by :attr:`other`, elementwise, and floors
// the result.
//
// .. math::
//     \text{{out}}_i = \text{floor} \left( \frac{{\text{{input}}_i}}{{\text{{other}}_i}} \right)
//
//
//
// Supports broadcasting to a common shape, type promotion, and integer and float inputs.
//
// Args:
//     input (Tensor or Number): the dividend
//     other (Tensor or Number): the divisor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([4.0, 3.0])
//     >>> b = torch.tensor([2.0, 2.0])
//     >>> torch.floor_divide(a, b)
//     tensor([2.0, 1.0])
//     >>> torch.floor_divide(a, 1.4)
//     tensor([2.0, 2.0])
//
//
//go:linkname FloorDivide py.floor_divide
func FloorDivide(input *py.Object, other *py.Object) *py.Object
//
// fmax(input, other, *, out=None) -> Tensor
//
// Computes the element-wise maximum of :attr:`input` and :attr:`other`.
//
// This is like :func:`torch.maximum` except it handles NaNs differently:
// if exactly one of the two elements being compared is a NaN then the non-NaN element is taken as the maximum.
// Only if both elements are NaN is NaN propagated.
//
// This function is a wrapper around C++'s ``std::fmax`` and is similar to NumPy's ``fmax`` function.
//
// Supports :ref:`broadcasting to a common shape <broadcasting-semantics>`,
// :ref:`type promotion <type-promotion-doc>`, and integer and floating-point inputs.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor): the second input tensor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([9.7, float('nan'), 3.1, float('nan')])
//     >>> b = torch.tensor([-2.2, 0.5, float('nan'), float('nan')])
//     >>> torch.fmax(a, b)
//     tensor([9.7000, 0.5000, 3.1000,    nan])
//
//
//go:linkname Fmax py.fmax
func Fmax(input *py.Object, other *py.Object) *py.Object
//
// fmin(input, other, *, out=None) -> Tensor
//
// Computes the element-wise minimum of :attr:`input` and :attr:`other`.
//
// This is like :func:`torch.minimum` except it handles NaNs differently:
// if exactly one of the two elements being compared is a NaN then the non-NaN element is taken as the minimum.
// Only if both elements are NaN is NaN propagated.
//
// This function is a wrapper around C++'s ``std::fmin`` and is similar to NumPy's ``fmin`` function.
//
// Supports :ref:`broadcasting to a common shape <broadcasting-semantics>`,
// :ref:`type promotion <type-promotion-doc>`, and integer and floating-point inputs.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor): the second input tensor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([2.2, float('nan'), 2.1, float('nan')])
//     >>> b = torch.tensor([-9.3, 0.1, float('nan'), float('nan')])
//     >>> torch.fmin(a, b)
//     tensor([-9.3000, 0.1000, 2.1000,    nan])
//
//
//go:linkname Fmin py.fmin
func Fmin(input *py.Object, other *py.Object) *py.Object
//
// fmod(input, other, *, out=None) -> Tensor
//
// Applies C++'s `std::fmod <https://en.cppreference.com/w/cpp/numeric/math/fmod>`_ entrywise.
// The result has the same sign as the dividend :attr:`input` and its absolute value
// is less than that of :attr:`other`.
//
// This function may be defined in terms of :func:`torch.div` as
//
// .. code:: python
//
//     torch.fmod(a, b) == a - a.div(b, rounding_mode="trunc") * b
//
// Supports :ref:`broadcasting to a common shape <broadcasting-semantics>`,
// :ref:`type promotion <type-promotion-doc>`, and integer and float inputs.
//
// .. note::
//
//     When the divisor is zero, returns ``NaN`` for floating point dtypes
//     on both CPU and GPU; raises ``RuntimeError`` for integer division by
//     zero on CPU; Integer division by zero on GPU may return any value.
//
// .. note::
//
//    Complex inputs are not supported. In some cases, it is not mathematically
//    possible to satisfy the definition of a modulo operation with complex numbers.
//
// .. seealso::
//
//     :func:`torch.remainder` which implements Python's modulus operator.
//     This one is defined using division rounding down the result.
//
// Args:
//     input (Tensor): the dividend
//     other (Tensor or Scalar): the divisor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.fmod(torch.tensor([-3., -2, -1, 1, 2, 3]), 2)
//     tensor([-1., -0., -1.,  1.,  0.,  1.])
//     >>> torch.fmod(torch.tensor([1, 2, 3, 4, 5]), -1.5)
//     tensor([1.0000, 0.5000, 0.0000, 1.0000, 0.5000])
//
//
//
//go:linkname Fmod py.fmod
func Fmod(input *py.Object, other *py.Object) *py.Object
//
// frac(input, *, out=None) -> Tensor
//
// Computes the fractional portion of each element in :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \text{input}_{i} - \left\lfloor |\text{input}_{i}| \right\rfloor * \operatorname{sgn}(\text{input}_{i})
//
// Example::
//
//     >>> torch.frac(torch.tensor([1, 2.5, -3.2]))
//     tensor([ 0.0000,  0.5000, -0.2000])
//
//
//go:linkname Frac py.frac
func Frac(input *py.Object) *py.Object
// None
//
//go:linkname Frac_ py.frac_
func Frac_(__llgo_va_list ...interface{}) *py.Object
//
// frexp(input, *, out=None) -> (Tensor mantissa, Tensor exponent)
//
// Decomposes :attr:`input` into mantissa and exponent tensors
// such that :math:`\text{input} = \text{mantissa} \times 2^{\text{exponent}}`.
//
// The range of mantissa is the open interval (-1, 1).
//
// Supports float inputs.
//
// Args:
//     input (Tensor): the input tensor
//
//
// Keyword args:
//     out (tuple, optional): the output tensors
//
// Example::
//
//     >>> x = torch.arange(9.)
//     >>> mantissa, exponent = torch.frexp(x)
//     >>> mantissa
//     tensor([0.0000, 0.5000, 0.5000, 0.7500, 0.5000, 0.6250, 0.7500, 0.8750, 0.5000])
//     >>> exponent
//     tensor([0, 1, 2, 2, 3, 3, 3, 3, 4], dtype=torch.int32)
//     >>> torch.ldexp(mantissa, exponent)
//     tensor([0., 1., 2., 3., 4., 5., 6., 7., 8.])
//
//
//go:linkname Frexp py.frexp
func Frexp(input *py.Object) *py.Object
// None
//
//go:linkname FrobeniusNorm py.frobenius_norm
func FrobeniusNorm(__llgo_va_list ...interface{}) *py.Object
//
// from_file(filename, shared=None, size=0, *, dtype=None, layout=None, device=None, pin_memory=False)
//
// Creates a CPU tensor with a storage backed by a memory-mapped file.
//
// If ``shared`` is True, then memory is shared between processes. All changes are written to the file.
// If ``shared`` is False, then changes to the tensor do not affect the file.
//
// ``size`` is the number of elements in the Tensor. If ``shared`` is ``False``, then the file must contain
// at least ``size * sizeof(dtype)`` bytes. If ``shared`` is ``True`` the file will be created if needed.
//
// .. note::
//     Only CPU tensors can be mapped to files.
//
// .. note::
//     For now, tensors with storages backed by a memory-mapped file cannot be created in pinned memory.
//
//
// Args:
//     filename (str): file name to map
//     shared (bool): whether to share memory (whether ``MAP_SHARED`` or ``MAP_PRIVATE`` is passed to the
//                     underlying `mmap(2) call <https://man7.org/linux/man-pages/man2/mmap.2.html>`_)
//     size (int): number of elements in the tensor
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`).
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     pin_memory (bool, optional): If set, returned tensor would be allocated in
//         the pinned memory. Works only for CPU tensors. Default: ``False``.
//
// Example::
//     >>> t = torch.randn(2, 5, dtype=torch.float64)
//     >>> t.numpy().tofile('storage.pt')
//     >>> t_mapped = torch.from_file('storage.pt', shared=False, size=10, dtype=torch.float64)
//
//
//go:linkname FromFile py.from_file
func FromFile(filename *py.Object, shared *py.Object, size *py.Object) *py.Object
//
// from_numpy(ndarray) -> Tensor
//
// Creates a :class:`Tensor` from a :class:`numpy.ndarray`.
//
// The returned tensor and :attr:`ndarray` share the same memory. Modifications to
// the tensor will be reflected in the :attr:`ndarray` and vice versa. The returned
// tensor is not resizable.
//
// It currently accepts :attr:`ndarray` with dtypes of ``numpy.float64``,
// ``numpy.float32``, ``numpy.float16``, ``numpy.complex64``, ``numpy.complex128``,
// ``numpy.int64``, ``numpy.int32``, ``numpy.int16``, ``numpy.int8``, ``numpy.uint8``,
// and ``bool``.
//
// .. warning::
//     Writing to a tensor created from a read-only NumPy array is not supported and will result in undefined behavior.
//
// Example::
//
//     >>> a = numpy.array([1, 2, 3])
//     >>> t = torch.from_numpy(a)
//     >>> t
//     tensor([ 1,  2,  3])
//     >>> t[0] = -1
//     >>> a
//     array([-1,  2,  3])
//
//
//go:linkname FromNumpy py.from_numpy
func FromNumpy(ndarray *py.Object) *py.Object
//
// frombuffer(buffer, *, dtype, count=-1, offset=0, requires_grad=False) -> Tensor
//
// Creates a 1-dimensional :class:`Tensor` from an object that implements
// the Python buffer protocol.
//
// Skips the first :attr:`offset` bytes in the buffer, and interprets the rest of
// the raw bytes as a 1-dimensional tensor of type :attr:`dtype` with :attr:`count`
// elements.
//
// Note that either of the following must be true:
//
// 1. :attr:`count` is a positive non-zero number, and the total number of bytes
// in the buffer is less than :attr:`offset` plus :attr:`count` times the size
// (in bytes) of :attr:`dtype`.
//
// 2. :attr:`count` is negative, and the length (number of bytes) of the buffer
// subtracted by the :attr:`offset` is a multiple of the size (in bytes) of
// :attr:`dtype`.
//
// The returned tensor and buffer share the same memory. Modifications to
// the tensor will be reflected in the buffer and vice versa. The returned
// tensor is not resizable.
//
// .. note::
//     This function increments the reference count for the object that
//     owns the shared memory. Therefore, such memory will not be deallocated
//     before the returned tensor goes out of scope.
//
// .. warning::
//     This function's behavior is undefined when passed an object implementing
//     the buffer protocol whose data is not on the CPU. Doing so is likely to
//     cause a segmentation fault.
//
// .. warning::
//     This function does not try to infer the :attr:`dtype` (hence, it is not
//     optional). Passing a different :attr:`dtype` than its source may result
//     in unexpected behavior.
//
// Args:
//     buffer (object): a Python object that exposes the buffer interface.
//
// Keyword args:
//     dtype (:class:`torch.dtype`): the desired data type of returned tensor.
//     count (int, optional): the number of desired elements to be read.
//         If negative, all the elements (until the end of the buffer) will be
//         read. Default: -1.
//     offset (int, optional): the number of bytes to skip at the start of
//         the buffer. Default: 0.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
// Example::
//
//     >>> import array
//     >>> a = array.array('i', [1, 2, 3])
//     >>> t = torch.frombuffer(a, dtype=torch.int32)
//     >>> t
//     tensor([ 1,  2,  3])
//     >>> t[0] = -1
//     >>> a
//     array([-1,  2,  3])
//
//     >>> # Interprets the signed char bytes as 32-bit integers.
//     >>> # Each 4 signed char elements will be interpreted as
//     >>> # 1 signed 32-bit integer.
//     >>> import array
//     >>> a = array.array('b', [-1, 0, 0, 0])
//     >>> torch.frombuffer(a, dtype=torch.int32)
//     tensor([255], dtype=torch.int32)
//
//
//go:linkname Frombuffer py.frombuffer
func Frombuffer(buffer *py.Object) *py.Object
//
// full(size, fill_value, *, out=None, dtype=None, layout=torch.strided, device=None, requires_grad=False) -> Tensor
//
// Creates a tensor of size :attr:`size` filled with :attr:`fill_value`. The
// tensor's dtype is inferred from :attr:`fill_value`.
//
// Args:
//     size (int...): a list, tuple, or :class:`torch.Size` of integers defining the
//         shape of the output tensor.
//     fill_value (Scalar): the value to fill the output tensor with.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`).
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
// Example::
//
//     >>> torch.full((2, 3), 3.141592)
//     tensor([[ 3.1416,  3.1416,  3.1416],
//             [ 3.1416,  3.1416,  3.1416]])
//
//
//go:linkname Full py.full
func Full(size *py.Object, fillValue *py.Object) *py.Object
//
// full_like(input, fill_value, \*, dtype=None, layout=torch.strided, device=None, requires_grad=False, memory_format=torch.preserve_format) -> Tensor
//
// Returns a tensor with the same size as :attr:`input` filled with :attr:`fill_value`.
// ``torch.full_like(input, fill_value)`` is equivalent to
// ``torch.full(input.size(), fill_value, dtype=input.dtype, layout=input.layout, device=input.device)``.
//
// Args:
//     input (Tensor): the size of :attr:`input` will determine size of the output tensor.
//     fill_value: the number to fill the output tensor with.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned Tensor.
//         Default: if ``None``, defaults to the dtype of :attr:`input`.
//     layout (:class:`torch.layout`, optional): the desired layout of returned tensor.
//         Default: if ``None``, defaults to the layout of :attr:`input`.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, defaults to the device of :attr:`input`.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     memory_format (:class:`torch.memory_format`, optional): the desired memory format of
//         returned Tensor. Default: ``torch.preserve_format``.
//
//
//go:linkname FullLike py.full_like
func FullLike(input *py.Object, fillValue *py.Object) *py.Object
// None
//
//go:linkname FusedMovingAvgObsFakeQuant py.fused_moving_avg_obs_fake_quant
func FusedMovingAvgObsFakeQuant(__llgo_va_list ...interface{}) *py.Object
//
// gather(input, dim, index, *, sparse_grad=False, out=None) -> Tensor
//
// Gathers values along an axis specified by `dim`.
//
// For a 3-D tensor the output is specified by::
//
//     out[i][j][k] = input[index[i][j][k]][j][k]  # if dim == 0
//     out[i][j][k] = input[i][index[i][j][k]][k]  # if dim == 1
//     out[i][j][k] = input[i][j][index[i][j][k]]  # if dim == 2
//
// :attr:`input` and :attr:`index` must have the same number of dimensions.
// It is also required that ``index.size(d) <= input.size(d)`` for all
// dimensions ``d != dim``.  :attr:`out` will have the same shape as :attr:`index`.
// Note that ``input`` and ``index`` do not broadcast against each other.
//
// Args:
//     input (Tensor): the source tensor
//     dim (int): the axis along which to index
//     index (LongTensor): the indices of elements to gather
//
// Keyword arguments:
//     sparse_grad (bool, optional): If ``True``, gradient w.r.t. :attr:`input` will be a sparse tensor.
//     out (Tensor, optional): the destination tensor
//
// Example::
//
//     >>> t = torch.tensor([[1, 2], [3, 4]])
//     >>> torch.gather(t, 1, torch.tensor([[0, 0], [1, 0]]))
//     tensor([[ 1,  1],
//             [ 4,  3]])
//
//
//go:linkname Gather py.gather
func Gather(input *py.Object, dim *py.Object, index *py.Object) *py.Object
//
// gcd(input, other, *, out=None) -> Tensor
//
// Computes the element-wise greatest common divisor (GCD) of :attr:`input` and :attr:`other`.
//
// Both :attr:`input` and :attr:`other` must have integer types.
//
// .. note::
//     This defines :math:`gcd(0, 0) = 0`.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor): the second input tensor
//
// Keyword arguments:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([5, 10, 15])
//     >>> b = torch.tensor([3, 4, 5])
//     >>> torch.gcd(a, b)
//     tensor([1, 2, 5])
//     >>> c = torch.tensor([3])
//     >>> torch.gcd(a, c)
//     tensor([1, 1, 3])
//
//
//go:linkname Gcd py.gcd
func Gcd(input *py.Object, other *py.Object) *py.Object
// None
//
//go:linkname Gcd_ py.gcd_
func Gcd_(__llgo_va_list ...interface{}) *py.Object
//
// ge(input, other, *, out=None) -> Tensor
//
// Computes :math:`\text{input} \geq \text{other}` element-wise.
//
//
// The second argument can be a number or a tensor whose shape is
// :ref:`broadcastable <broadcasting-semantics>` with the first argument.
//
// Args:
//     input (Tensor): the tensor to compare
//     other (Tensor or float): the tensor or value to compare
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Returns:
//     A boolean tensor that is True where :attr:`input` is greater than or equal to :attr:`other` and False elsewhere
//
// Example::
//
//     >>> torch.ge(torch.tensor([[1, 2], [3, 4]]), torch.tensor([[1, 1], [4, 4]]))
//     tensor([[True, True], [False, True]])
//
//
//go:linkname Ge py.ge
func Ge(input *py.Object, other *py.Object) *py.Object
//
// geqrf(input, *, out=None) -> (Tensor, Tensor)
//
// This is a low-level function for calling LAPACK's geqrf directly. This function
// returns a namedtuple (a, tau) as defined in `LAPACK documentation for geqrf`_ .
//
// Computes a QR decomposition of :attr:`input`.
// Both `Q` and `R` matrices are stored in the same output tensor `a`.
// The elements of `R` are stored on and above the diagonal.
// Elementary reflectors (or Householder vectors) implicitly defining matrix `Q`
// are stored below the diagonal.
// The results of this function can be used together with :func:`torch.linalg.householder_product`
// to obtain the `Q` matrix or
// with :func:`torch.ormqr`, which uses an implicit representation of the `Q` matrix,
// for an efficient matrix-matrix multiplication.
//
// See `LAPACK documentation for geqrf`_ for further details.
//
// .. note::
//     See also :func:`torch.linalg.qr`, which computes Q and R matrices, and :func:`torch.linalg.lstsq`
//     with the ``driver="gels"`` option for a function that can solve matrix equations using a QR decomposition.
//
// Args:
//     input (Tensor): the input matrix
//
// Keyword args:
//     out (tuple, optional): the output tuple of (Tensor, Tensor). Ignored if `None`. Default: `None`.
//
// .. _LAPACK documentation for geqrf:
//     http://www.netlib.org/lapack/explore-html/df/dc5/group__variants_g_ecomputational_ga3766ea903391b5cf9008132f7440ec7b.html
//
//
//
//go:linkname Geqrf py.geqrf
func Geqrf(input *py.Object) *py.Object
//
// ger(input, vec2, *, out=None) -> Tensor
//
// Alias of :func:`torch.outer`.
//
// .. warning::
//     This function is deprecated and will be removed in a future PyTorch release.
//     Use :func:`torch.outer` instead.
//
//
//go:linkname Ger py.ger
func Ger(input *py.Object, vec2 *py.Object) *py.Object
// None
//
//go:linkname GetDevice py.get_device
func GetDevice(__llgo_va_list ...interface{}) *py.Object
//
// gradient(input, *, spacing=1, dim=None, edge_order=1) -> List of Tensors
//
// Estimates the gradient of a function :math:`g : \mathbb{R}^n \rightarrow \mathbb{R}` in
// one or more dimensions using the `second-order accurate central differences method
// <https://www.ams.org/journals/mcom/1988-51-184/S0025-5718-1988-0935077-0/S0025-5718-1988-0935077-0.pdf>`_ and
// either first or second order estimates at the boundaries.
//
// The gradient of :math:`g` is estimated using samples. By default, when :attr:`spacing` is not
// specified, the samples are entirely described by :attr:`input`, and the mapping of input coordinates
// to an output is the same as the tensor's mapping of indices to values. For example, for a three-dimensional
// :attr:`input` the function described is :math:`g : \mathbb{R}^3 \rightarrow \mathbb{R}`, and
// :math:`g(1, 2, 3)\ == input[1, 2, 3]`.
//
// When :attr:`spacing` is specified, it modifies the relationship between :attr:`input` and input coordinates.
// This is detailed in the "Keyword Arguments" section below.
//
// The gradient is estimated by estimating each partial derivative of :math:`g` independently. This estimation is
// accurate if :math:`g` is in :math:`C^3` (it has at least 3 continuous derivatives), and the estimation can be
// improved by providing closer samples. Mathematically, the value at each interior point of a partial derivative
// is estimated using `Taylor’s theorem with remainder <https://en.wikipedia.org/wiki/Taylor%27s_theorem>`_.
// Letting :math:`x` be an interior point with :math:`x-h_l` and :math:`x+h_r` be points neighboring
// it to the left and right respectively, :math:`f(x+h_r)` and :math:`f(x-h_l)` can be estimated using:
//
// .. math::
//     \begin{aligned}
//         f(x+h_r) = f(x) + h_r f'(x) + {h_r}^2  \frac{f''(x)}{2} + {h_r}^3 \frac{f'''(\xi_1)}{6}, \xi_1 \in (x, x+h_r) \\
//         f(x-h_l) = f(x) - h_l f'(x) + {h_l}^2  \frac{f''(x)}{2} - {h_l}^3 \frac{f'''(\xi_2)}{6}, \xi_2 \in (x, x-h_l) \\
//     \end{aligned}
//
// Using the fact that :math:`f \in C^3` and solving the linear system, we derive:
//
// .. math::
//     f'(x) \approx \frac{ {h_l}^2 f(x+h_r) - {h_r}^2 f(x-h_l)
//           + ({h_r}^2-{h_l}^2 ) f(x) }{ {h_r} {h_l}^2 + {h_r}^2 {h_l} }
//
// .. note::
//     We estimate the gradient of functions in complex domain
//     :math:`g : \mathbb{C}^n \rightarrow \mathbb{C}` in the same way.
//
// The value of each partial derivative at the boundary points is computed differently. See edge_order below.
//
// Args:
//     input (``Tensor``): the tensor that represents the values of the function
//
// Keyword args:
//     spacing (``scalar``, ``list of scalar``, ``list of Tensor``, optional): :attr:`spacing` can be used to modify
//         how the :attr:`input` tensor's indices relate to sample coordinates. If :attr:`spacing` is a scalar then
//         the indices are multiplied by the scalar to produce the coordinates. For example, if :attr:`spacing=2` the
//         indices (1, 2, 3) become coordinates (2, 4, 6). If :attr:`spacing` is a list of scalars then the corresponding
//         indices are multiplied. For example, if :attr:`spacing=(2, -1, 3)` the indices (1, 2, 3) become coordinates (2, -2, 9).
//         Finally, if :attr:`spacing` is a list of one-dimensional tensors then each tensor specifies the coordinates for
//         the corresponding dimension. For example, if the indices are (1, 2, 3) and the tensors are (t0, t1, t2), then
//         the coordinates are (t0[1], t1[2], t2[3])
//
//     dim (``int``, ``list of int``, optional): the dimension or dimensions to approximate the gradient over.  By default
//         the partial  gradient in every dimension is computed. Note that when :attr:`dim` is  specified the elements of
//         the :attr:`spacing` argument must correspond with the specified dims."
//
//     edge_order (``int``, optional): 1 or 2, for `first-order
//         <https://www.ams.org/journals/mcom/1988-51-184/S0025-5718-1988-0935077-0/S0025-5718-1988-0935077-0.pdf>`_ or
//         `second-order <https://www.ams.org/journals/mcom/1988-51-184/S0025-5718-1988-0935077-0/S0025-5718-1988-0935077-0.pdf>`_
//         estimation of the boundary ("edge") values, respectively.
//
// Examples::
//
//     >>> # Estimates the gradient of f(x)=x^2 at points [-2, -1, 2, 4]
//     >>> coordinates = (torch.tensor([-2., -1., 1., 4.]),)
//     >>> values = torch.tensor([4., 1., 1., 16.], )
//     >>> torch.gradient(values, spacing = coordinates)
//     (tensor([-3., -2., 2., 5.]),)
//
//     >>> # Estimates the gradient of the R^2 -> R function whose samples are
//     >>> # described by the tensor t. Implicit coordinates are [0, 1] for the outermost
//     >>> # dimension and [0, 1, 2, 3] for the innermost dimension, and function estimates
//     >>> # partial derivative for both dimensions.
//     >>> t = torch.tensor([[1, 2, 4, 8], [10, 20, 40, 80]])
//     >>> torch.gradient(t)
//     (tensor([[ 9., 18., 36., 72.],
//              [ 9., 18., 36., 72.]]),
//      tensor([[ 1.0000, 1.5000, 3.0000, 4.0000],
//              [10.0000, 15.0000, 30.0000, 40.0000]]))
//
//     >>> # A scalar value for spacing modifies the relationship between tensor indices
//     >>> # and input coordinates by multiplying the indices to find the
//     >>> # coordinates. For example, below the indices of the innermost
//     >>> # 0, 1, 2, 3 translate to coordinates of [0, 2, 4, 6], and the indices of
//     >>> # the outermost dimension 0, 1 translate to coordinates of [0, 2].
//     >>> torch.gradient(t, spacing = 2.0) # dim = None (implicitly [0, 1])
//     (tensor([[ 4.5000, 9.0000, 18.0000, 36.0000],
//               [ 4.5000, 9.0000, 18.0000, 36.0000]]),
//      tensor([[ 0.5000, 0.7500, 1.5000, 2.0000],
//               [ 5.0000, 7.5000, 15.0000, 20.0000]]))
//     >>> # doubling the spacing between samples halves the estimated partial gradients.
//
//     >>>
//     >>> # Estimates only the partial derivative for dimension 1
//     >>> torch.gradient(t, dim = 1) # spacing = None (implicitly 1.)
//     (tensor([[ 1.0000, 1.5000, 3.0000, 4.0000],
//              [10.0000, 15.0000, 30.0000, 40.0000]]),)
//
//     >>> # When spacing is a list of scalars, the relationship between the tensor
//     >>> # indices and input coordinates changes based on dimension.
//     >>> # For example, below, the indices of the innermost dimension 0, 1, 2, 3 translate
//     >>> # to coordinates of [0, 3, 6, 9], and the indices of the outermost dimension
//     >>> # 0, 1 translate to coordinates of [0, 2].
//     >>> torch.gradient(t, spacing = [3., 2.])
//     (tensor([[ 4.5000, 9.0000, 18.0000, 36.0000],
//              [ 4.5000, 9.0000, 18.0000, 36.0000]]),
//      tensor([[ 0.3333, 0.5000, 1.0000, 1.3333],
//              [ 3.3333, 5.0000, 10.0000, 13.3333]]))
//
//     >>> # The following example is a replication of the previous one with explicit
//     >>> # coordinates.
//     >>> coords = (torch.tensor([0, 2]), torch.tensor([0, 3, 6, 9]))
//     >>> torch.gradient(t, spacing = coords)
//     (tensor([[ 4.5000, 9.0000, 18.0000, 36.0000],
//              [ 4.5000, 9.0000, 18.0000, 36.0000]]),
//      tensor([[ 0.3333, 0.5000, 1.0000, 1.3333],
//              [ 3.3333, 5.0000, 10.0000, 13.3333]]))
//
//
//
//go:linkname Gradient py.gradient
func Gradient(input *py.Object) *py.Object
//
// greater(input, other, *, out=None) -> Tensor
//
// Alias for :func:`torch.gt`.
//
//
//go:linkname Greater py.greater
func Greater(input *py.Object, other *py.Object) *py.Object
//
// greater_equal(input, other, *, out=None) -> Tensor
//
// Alias for :func:`torch.ge`.
//
//
//go:linkname GreaterEqual py.greater_equal
func GreaterEqual(input *py.Object, other *py.Object) *py.Object
// None
//
//go:linkname GridSampler py.grid_sampler
func GridSampler(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname GridSampler2d py.grid_sampler_2d
func GridSampler2d(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname GridSampler3d py.grid_sampler_3d
func GridSampler3d(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname GroupNorm py.group_norm
func GroupNorm(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Gru py.gru
func Gru(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname GruCell py.gru_cell
func GruCell(__llgo_va_list ...interface{}) *py.Object
//
// gt(input, other, *, out=None) -> Tensor
//
// Computes :math:`\text{input} > \text{other}` element-wise.
//
//
// The second argument can be a number or a tensor whose shape is
// :ref:`broadcastable <broadcasting-semantics>` with the first argument.
//
// Args:
//     input (Tensor): the tensor to compare
//     other (Tensor or float): the tensor or value to compare
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Returns:
//     A boolean tensor that is True where :attr:`input` is greater than :attr:`other` and False elsewhere
//
// Example::
//
//     >>> torch.gt(torch.tensor([[1, 2], [3, 4]]), torch.tensor([[1, 1], [4, 4]]))
//     tensor([[False, True], [False, False]])
//
//
//go:linkname Gt py.gt
func Gt(input *py.Object, other *py.Object) *py.Object
//
// hamming_window(window_length, periodic=True, alpha=0.54, beta=0.46, *, dtype=None, layout=torch.strided, device=None, requires_grad=False) -> Tensor
//
// Hamming window function.
//
// .. math::
//     w[n] = \alpha - \beta\ \cos \left( \frac{2 \pi n}{N - 1} \right),
//
// where :math:`N` is the full window size.
//
// The input :attr:`window_length` is a positive integer controlling the
// returned window size. :attr:`periodic` flag determines whether the returned
// window trims off the last duplicate value from the symmetric window and is
// ready to be used as a periodic window with functions like
// :meth:`torch.stft`. Therefore, if :attr:`periodic` is true, the :math:`N` in
// above formula is in fact :math:`\text{window\_length} + 1`. Also, we always have
// ``torch.hamming_window(L, periodic=True)`` equal to
// ``torch.hamming_window(L + 1, periodic=False)[:-1])``.
//
// .. note::
//     If :attr:`window_length` :math:`=1`, the returned window contains a single value 1.
//
// .. note::
//     This is a generalized version of :meth:`torch.hann_window`.
//
// Arguments:
//     window_length (int): the size of returned window
//     periodic (bool, optional): If True, returns a window to be used as periodic
//         function. If False, return a symmetric window.
//     alpha (float, optional): The coefficient :math:`\alpha` in the equation above
//     beta (float, optional): The coefficient :math:`\beta` in the equation above
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`). Only floating point types are supported.
//     layout (:class:`torch.layout`, optional): the desired layout of returned window tensor. Only
//           ``torch.strided`` (dense layout) is supported.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
// Returns:
//     Tensor: A 1-D tensor of size :math:`(\text{window\_length},)` containing the window.
//
//
//
//go:linkname HammingWindow py.hamming_window
func HammingWindow(windowLength *py.Object, periodic *py.Object, alpha *py.Object, beta *py.Object) *py.Object
//
// hann_window(window_length, periodic=True, *, dtype=None, layout=torch.strided, device=None, requires_grad=False) -> Tensor
//
// Hann window function.
//
// .. math::
//     w[n] = \frac{1}{2}\ \left[1 - \cos \left( \frac{2 \pi n}{N - 1} \right)\right] =
//             \sin^2 \left( \frac{\pi n}{N - 1} \right),
//
// where :math:`N` is the full window size.
//
// The input :attr:`window_length` is a positive integer controlling the
// returned window size. :attr:`periodic` flag determines whether the returned
// window trims off the last duplicate value from the symmetric window and is
// ready to be used as a periodic window with functions like
// :meth:`torch.stft`. Therefore, if :attr:`periodic` is true, the :math:`N` in
// above formula is in fact :math:`\text{window\_length} + 1`. Also, we always have
// ``torch.hann_window(L, periodic=True)`` equal to
// ``torch.hann_window(L + 1, periodic=False)[:-1])``.
//
// .. note::
//     If :attr:`window_length` :math:`=1`, the returned window contains a single value 1.
//
// Arguments:
//     window_length (int): the size of returned window
//     periodic (bool, optional): If True, returns a window to be used as periodic
//         function. If False, return a symmetric window.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`). Only floating point types are supported.
//     layout (:class:`torch.layout`, optional): the desired layout of returned window tensor. Only
//           ``torch.strided`` (dense layout) is supported.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
// Returns:
//     Tensor: A 1-D tensor of size :math:`(\text{window\_length},)` containing the window
//
//
//
//go:linkname HannWindow py.hann_window
func HannWindow(windowLength *py.Object, periodic *py.Object) *py.Object
//
// hardshrink(input, lambd=0.5) -> Tensor
//
// Applies the hard shrinkage function element-wise
//
// See :class:`~torch.nn.Hardshrink` for more details.
//
//
//go:linkname Hardshrink py.hardshrink
func Hardshrink(input *py.Object, lambd *py.Object) *py.Object
//
// heaviside(input, values, *, out=None) -> Tensor
//
// Computes the Heaviside step function for each element in :attr:`input`.
// The Heaviside step function is defined as:
//
// .. math::
//     \text{{heaviside}}(input, values) = \begin{cases}
//         0, & \text{if input < 0}\\
//         values, & \text{if input == 0}\\
//         1, & \text{if input > 0}
//     \end{cases}
//
//
// Args:
//     input (Tensor): the input tensor.
//     values (Tensor): The values to use where :attr:`input` is zero.
//
// Keyword arguments:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> input = torch.tensor([-1.5, 0, 2.0])
//     >>> values = torch.tensor([0.5])
//     >>> torch.heaviside(input, values)
//     tensor([0.0000, 0.5000, 1.0000])
//     >>> values = torch.tensor([1.2, -2.0, 3.5])
//     >>> torch.heaviside(input, values)
//     tensor([0., -2., 1.])
//
//
//
//go:linkname Heaviside py.heaviside
func Heaviside(input *py.Object, values *py.Object) *py.Object
// None
//
//go:linkname HingeEmbeddingLoss py.hinge_embedding_loss
func HingeEmbeddingLoss(__llgo_va_list ...interface{}) *py.Object
//
// histc(input, bins=100, min=0, max=0, *, out=None) -> Tensor
//
// Computes the histogram of a tensor.
//
// The elements are sorted into equal width bins between :attr:`min` and
// :attr:`max`. If :attr:`min` and :attr:`max` are both zero, the minimum and
// maximum values of the data are used.
//
// Elements lower than min and higher than max and ``NaN`` elements are ignored.
//
// Args:
//     input (Tensor): the input tensor.
//     bins (int): number of histogram bins
//     min (Scalar): lower end of the range (inclusive)
//     max (Scalar): upper end of the range (inclusive)
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Returns:
//     Tensor: Histogram represented as a tensor
//
// Example::
//
//     >>> torch.histc(torch.tensor([1., 2, 1]), bins=4, min=0, max=3)
//     tensor([ 0.,  2.,  1.,  0.])
//
//
//go:linkname Histc py.histc
func Histc(input *py.Object, bins *py.Object, min *py.Object, max *py.Object) *py.Object
//
// histogram(input, bins, *, range=None, weight=None, density=False, out=None) -> (Tensor, Tensor)
//
// Computes a histogram of the values in a tensor.
//
// :attr:`bins` can be an integer or a 1D tensor.
//
// If :attr:`bins` is an int, it specifies the number of equal-width bins.
// By default, the lower and upper range of the bins is determined by the
// minimum and maximum elements of the input tensor. The :attr:`range`
// argument can be provided to specify a range for the bins.
//
// If :attr:`bins` is a 1D tensor, it specifies the sequence of bin edges
// including the rightmost edge. It should contain at least 2 elements
// and its elements should be increasing.
//
// Args:
//     input (Tensor): the input tensor.
//     bins: int or 1D Tensor. If int, defines the number of equal-width bins. If tensor,
//           defines the sequence of bin edges including the rightmost edge.
//
// Keyword args:
//     range (tuple of float): Defines the range of the bins.
//     weight (Tensor): If provided, weight should have the same shape as input. Each value in
//                      input contributes its associated weight towards its bin's result.
//     density (bool): If False, the result will contain the count (or total weight) in each bin.
//                     If True, the result is the value of the probability density function over the bins,
//                     normalized such that the integral over the range of the bins is 1.
//     out (Tensor, optional): the output tensor. (tuple, optional): The result tuple of two output tensors (hist, bin_edges).
//
// Returns:
//     hist (Tensor): 1D Tensor containing the values of the histogram.
//     bin_edges(Tensor): 1D Tensor containing the edges of the histogram bins.
//
// Example::
//
//     >>> torch.histogram(torch.tensor([1., 2, 1]), bins=4, range=(0., 3.), weight=torch.tensor([1., 2., 4.]))
//     (tensor([ 0.,  5.,  2.,  0.]), tensor([0., 0.75, 1.5, 2.25, 3.]))
//     >>> torch.histogram(torch.tensor([1., 2, 1]), bins=4, range=(0., 3.), weight=torch.tensor([1., 2., 4.]), density=True)
//     (tensor([ 0.,  0.9524,  0.3810,  0.]), tensor([0., 0.75, 1.5, 2.25, 3.]))
//
//
//go:linkname Histogram py.histogram
func Histogram(input *py.Object, bins *py.Object) *py.Object
//
// histogramdd(input, bins, *, range=None, weight=None, density=False, out=None) -> (Tensor, Tensor[])
//
// Computes a multi-dimensional histogram of the values in a tensor.
//
// Interprets the elements of an input tensor whose innermost dimension has size N
// as a collection of N-dimensional points. Maps each of the points into a set of
// N-dimensional bins and returns the number of points (or total weight) in each bin.
//
// :attr:`input` must be a tensor with at least 2 dimensions.
// If input has shape (M, N), each of its M rows defines a point in N-dimensional space.
// If input has three or more dimensions, all but the last dimension are flattened.
//
// Each dimension is independently associated with its own strictly increasing sequence
// of bin edges. Bin edges may be specified explicitly by passing a sequence of 1D
// tensors. Alternatively, bin edges may be constructed automatically by passing a
// sequence of integers specifying the number of equal-width bins in each dimension.
//
// For each N-dimensional point in input:
//     - Each of its coordinates is binned independently among the bin edges
//         corresponding to its dimension
//     - Binning results are combined to identify the N-dimensional bin (if any)
//         into which the point falls
//     - If the point falls into a bin, the bin's count (or total weight) is incremented
//     - Points which do not fall into any bin do not contribute to the output
//
// :attr:`bins` can be a sequence of N 1D tensors, a sequence of N ints, or a single int.
//
// If :attr:`bins` is a sequence of N 1D tensors, it explicitly specifies the N sequences
// of bin edges. Each 1D tensor should contain a strictly increasing sequence with at
// least one element. A sequence of K bin edges defines K-1 bins, explicitly specifying
// the left and right edges of all bins. Every bin is exclusive of its left edge. Only
// the rightmost bin is inclusive of its right edge.
//
// If :attr:`bins` is a sequence of N ints, it specifies the number of equal-width bins
// in each dimension. By default, the leftmost and rightmost bin edges in each dimension
// are determined by the minimum and maximum elements of the input tensor in the
// corresponding dimension. The :attr:`range` argument can be provided to manually
// specify the leftmost and rightmost bin edges in each dimension.
//
// If :attr:`bins` is an int, it specifies the number of equal-width bins for all dimensions.
//
// .. note::
//     See also :func:`torch.histogram`, which specifically computes 1D histograms.
//     While :func:`torch.histogramdd` infers the dimensionality of its bins and
//     binned values from the shape of :attr:`input`, :func:`torch.histogram`
//     accepts and flattens :attr:`input` of any shape.
//
// Args:
//     input (Tensor): the input tensor.
//     bins: Tensor[], int[], or int.
//             If Tensor[], defines the sequences of bin edges.
//             If int[], defines the number of equal-width bins in each dimension.
//             If int, defines the number of equal-width bins for all dimensions.
// Keyword args:
//     range (sequence of float): Defines the leftmost and rightmost bin edges
//                                 in each dimension.
//     weight (Tensor): By default, each value in the input has weight 1. If a weight
//                         tensor is passed, each N-dimensional coordinate in input
//                         contributes its associated weight towards its bin's result.
//                         The weight tensor should have the same shape as the :attr:`input`
//                         tensor excluding its innermost dimension N.
//     density (bool): If False (default), the result will contain the count (or total weight)
//                     in each bin. If True, each count (weight) is divided by the total count
//                     (total weight), then divided by the volume of its associated bin.
// Returns:
//     hist (Tensor): N-dimensional Tensor containing the values of the histogram.
//     bin_edges(Tensor[]): sequence of N 1D Tensors containing the bin edges.
//
// Example::
//     >>> torch.histogramdd(torch.tensor([[0., 1.], [1., 0.], [2., 0.], [2., 2.]]), bins=[3, 3],
//     ...                   weight=torch.tensor([1., 2., 4., 8.]))
//         torch.return_types.histogramdd(
//             hist=tensor([[0., 1., 0.],
//                          [2., 0., 0.],
//                          [4., 0., 8.]]),
//             bin_edges=(tensor([0.0000, 0.6667, 1.3333, 2.0000]),
//                        tensor([0.0000, 0.6667, 1.3333, 2.0000])))
//
//     >>> torch.histogramdd(torch.tensor([[0., 0.], [1., 1.], [2., 2.]]), bins=[2, 2],
//     ...                   range=[0., 1., 0., 1.], density=True)
//         torch.return_types.histogramdd(
//            hist=tensor([[2., 0.],
//                         [0., 2.]]),
//            bin_edges=(tensor([0.0000, 0.5000, 1.0000]),
//                       tensor([0.0000, 0.5000, 1.0000])))
//
//
//
//go:linkname Histogramdd py.histogramdd
func Histogramdd(input *py.Object, bins *py.Object) *py.Object
// None
//
//go:linkname Hsmm py.hsmm
func Hsmm(__llgo_va_list ...interface{}) *py.Object
//
// hsplit(input, indices_or_sections) -> List of Tensors
//
// Splits :attr:`input`, a tensor with one or more dimensions, into multiple tensors
// horizontally according to :attr:`indices_or_sections`. Each split is a view of
// :attr:`input`.
//
// If :attr:`input` is one dimensional this is equivalent to calling
// torch.tensor_split(input, indices_or_sections, dim=0) (the split dimension is
// zero), and if :attr:`input` has two or more dimensions it's equivalent to calling
// torch.tensor_split(input, indices_or_sections, dim=1) (the split dimension is 1),
// except that if :attr:`indices_or_sections` is an integer it must evenly divide
// the split dimension or a runtime error will be thrown.
//
// This function is based on NumPy's :func:`numpy.hsplit`.
//
// Args:
//     input (Tensor): tensor to split.
//     indices_or_sections (int or list or tuple of ints): See argument in :func:`torch.tensor_split`.
//
// Example::
//     >>> t = torch.arange(16.0).reshape(4,4)
//     >>> t
//     tensor([[ 0.,  1.,  2.,  3.],
//             [ 4.,  5.,  6.,  7.],
//             [ 8.,  9., 10., 11.],
//             [12., 13., 14., 15.]])
//     >>> torch.hsplit(t, 2)
//     (tensor([[ 0.,  1.],
//              [ 4.,  5.],
//              [ 8.,  9.],
//              [12., 13.]]),
//      tensor([[ 2.,  3.],
//              [ 6.,  7.],
//              [10., 11.],
//              [14., 15.]]))
//     >>> torch.hsplit(t, [3, 6])
//     (tensor([[ 0.,  1.,  2.],
//              [ 4.,  5.,  6.],
//              [ 8.,  9., 10.],
//              [12., 13., 14.]]),
//      tensor([[ 3.],
//              [ 7.],
//              [11.],
//              [15.]]),
//      tensor([], size=(4, 0)))
//
//
//
//go:linkname Hsplit py.hsplit
func Hsplit(input *py.Object, indicesOrSections *py.Object) *py.Object
//
// hspmm(mat1, mat2, *, out=None) -> Tensor
//
// Performs a matrix multiplication of a :ref:`sparse COO matrix
// <sparse-coo-docs>` :attr:`mat1` and a strided matrix :attr:`mat2`. The
// result is a (1 + 1)-dimensional :ref:`hybrid COO matrix
// <sparse-hybrid-coo-docs>`.
//
// Args:
//     mat1 (Tensor): the first sparse matrix to be matrix multiplied
//     mat2 (Tensor): the second strided matrix to be matrix multiplied
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
//
//go:linkname Hspmm py.hspmm
func Hspmm(mat1 *py.Object, mat2 *py.Object) *py.Object
//
// hstack(tensors, *, out=None) -> Tensor
//
// Stack tensors in sequence horizontally (column wise).
//
// This is equivalent to concatenation along the first axis for 1-D tensors, and along the second axis for all other tensors.
//
// Args:
//     tensors (sequence of Tensors): sequence of tensors to concatenate
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([1, 2, 3])
//     >>> b = torch.tensor([4, 5, 6])
//     >>> torch.hstack((a,b))
//     tensor([1, 2, 3, 4, 5, 6])
//     >>> a = torch.tensor([[1],[2],[3]])
//     >>> b = torch.tensor([[4],[5],[6]])
//     >>> torch.hstack((a,b))
//     tensor([[1, 4],
//             [2, 5],
//             [3, 6]])
//
//
//
//go:linkname Hstack py.hstack
func Hstack(tensors *py.Object) *py.Object
//
// hypot(input, other, *, out=None) -> Tensor
//
// Given the legs of a right triangle, return its hypotenuse.
//
// .. math::
//     \text{out}_{i} = \sqrt{\text{input}_{i}^{2} + \text{other}_{i}^{2}}
//
// The shapes of ``input`` and ``other`` must be
// :ref:`broadcastable <broadcasting-semantics>`.
//
// Args:
//     input (Tensor): the first input tensor
//     other (Tensor): the second input tensor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.hypot(torch.tensor([4.0]), torch.tensor([3.0, 4.0, 5.0]))
//     tensor([5.0000, 5.6569, 6.4031])
//
//
//
//go:linkname Hypot py.hypot
func Hypot(input *py.Object, other *py.Object) *py.Object
//
// i0(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.i0`.
//
//
//go:linkname I0 py.i0
func I0(input *py.Object) *py.Object
// None
//
//go:linkname I0_ py.i0_
func I0_(__llgo_va_list ...interface{}) *py.Object
//
// igamma(input, other, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.gammainc`.
//
//
//go:linkname Igamma py.igamma
func Igamma(input *py.Object, other *py.Object) *py.Object
//
// igammac(input, other, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.gammaincc`.
//
//
//go:linkname Igammac py.igammac
func Igammac(input *py.Object, other *py.Object) *py.Object
//
// imag(input) -> Tensor
//
// Returns a new tensor containing imaginary values of the :attr:`self` tensor.
// The returned tensor and :attr:`self` share the same underlying storage.
//
// .. warning::
//     :func:`imag` is only supported for tensors with complex dtypes.
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> x=torch.randn(4, dtype=torch.cfloat)
//     >>> x
//     tensor([(0.3100+0.3553j), (-0.5445-0.7896j), (-1.6492-0.0633j), (-0.0638-0.8119j)])
//     >>> x.imag
//     tensor([ 0.3553, -0.7896, -0.0633, -0.8119])
//
//
//
//go:linkname Imag py.imag
func Imag(input *py.Object) *py.Object
//
// index_add(input, dim, index, source, *, alpha=1, out=None) -> Tensor
//
// See :meth:`~Tensor.index_add_` for function description.
//
//
//go:linkname IndexAdd py.index_add
func IndexAdd(input *py.Object, dim *py.Object, index *py.Object, source *py.Object) *py.Object
//
// index_copy(input, dim, index, source, *, out=None) -> Tensor
//
// See :meth:`~Tensor.index_add_` for function description.
//
//
//go:linkname IndexCopy py.index_copy
func IndexCopy(input *py.Object, dim *py.Object, index *py.Object, source *py.Object) *py.Object
// None
//
//go:linkname IndexFill py.index_fill
func IndexFill(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname IndexPut py.index_put
func IndexPut(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname IndexPut_ py.index_put_
func IndexPut_(__llgo_va_list ...interface{}) *py.Object
//
// index_reduce(input, dim, index, source, reduce, *, include_self=True, out=None) -> Tensor
//
// See :meth:`~Tensor.index_reduce_` for function description.
//
//
//go:linkname IndexReduce py.index_reduce
func IndexReduce(input *py.Object, dim *py.Object, index *py.Object, source *py.Object, reduce *py.Object) *py.Object
//
// index_select(input, dim, index, *, out=None) -> Tensor
//
// Returns a new tensor which indexes the :attr:`input` tensor along dimension
// :attr:`dim` using the entries in :attr:`index` which is a `LongTensor`.
//
// The returned tensor has the same number of dimensions as the original tensor
// (:attr:`input`).  The :attr:`dim`\ th dimension has the same size as the length
// of :attr:`index`; other dimensions have the same size as in the original tensor.
//
// .. note:: The returned tensor does **not** use the same storage as the original
//           tensor.  If :attr:`out` has a different shape than expected, we
//           silently change it to the correct shape, reallocating the underlying
//           storage if necessary.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int): the dimension in which we index
//     index (IntTensor or LongTensor): the 1-D tensor containing the indices to index
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> x = torch.randn(3, 4)
//     >>> x
//     tensor([[ 0.1427,  0.0231, -0.5414, -1.0009],
//             [-0.4664,  0.2647, -0.1228, -1.1068],
//             [-1.1734, -0.6571,  0.7230, -0.6004]])
//     >>> indices = torch.tensor([0, 2])
//     >>> torch.index_select(x, 0, indices)
//     tensor([[ 0.1427,  0.0231, -0.5414, -1.0009],
//             [-1.1734, -0.6571,  0.7230, -0.6004]])
//     >>> torch.index_select(x, 1, indices)
//     tensor([[ 0.1427, -0.5414],
//             [-0.4664, -0.1228],
//             [-1.1734,  0.7230]])
//
//
//go:linkname IndexSelect py.index_select
func IndexSelect(input *py.Object, dim *py.Object, index *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.indices`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname IndicesCopy py.indices_copy
func IndicesCopy(__llgo_va_list ...interface{}) *py.Object
//
// inner(input, other, *, out=None) -> Tensor
//
// Computes the dot product for 1D tensors. For higher dimensions, sums the product
// of elements from :attr:`input` and :attr:`other` along their last dimension.
//
// .. note::
//
//     If either :attr:`input` or :attr:`other` is a scalar, the result is equivalent
//     to `torch.mul(input, other)`.
//
//     If both :attr:`input` and :attr:`other` are non-scalars, the size of their last
//     dimension must match and the result is equivalent to `torch.tensordot(input,
//     other, dims=([-1], [-1]))`
//
// Args:
//     input (Tensor): First input tensor
//     other (Tensor): Second input tensor
//
// Keyword args:
//     out (Tensor, optional): Optional output tensor to write result into. The output
//                             shape is `input.shape[:-1] + other.shape[:-1]`.
//
// Example::
//
//     # Dot product
//     >>> torch.inner(torch.tensor([1, 2, 3]), torch.tensor([0, 2, 1]))
//     tensor(7)
//
//     # Multidimensional input tensors
//     >>> a = torch.randn(2, 3)
//     >>> a
//     tensor([[0.8173, 1.0874, 1.1784],
//             [0.3279, 0.1234, 2.7894]])
//     >>> b = torch.randn(2, 4, 3)
//     >>> b
//     tensor([[[-0.4682, -0.7159,  0.1506],
//             [ 0.4034, -0.3657,  1.0387],
//             [ 0.9892, -0.6684,  0.1774],
//             [ 0.9482,  1.3261,  0.3917]],
//
//             [[ 0.4537,  0.7493,  1.1724],
//             [ 0.2291,  0.5749, -0.2267],
//             [-0.7920,  0.3607, -0.3701],
//             [ 1.3666, -0.5850, -1.7242]]])
//     >>> torch.inner(a, b)
//     tensor([[[-0.9837,  1.1560,  0.2907,  2.6785],
//             [ 2.5671,  0.5452, -0.6912, -1.5509]],
//
//             [[ 0.1782,  2.9843,  0.7366,  1.5672],
//             [ 3.5115, -0.4864, -1.2476, -4.4337]]])
//
//     # Scalar input
//     >>> torch.inner(a, torch.tensor(2))
//     tensor([[1.6347, 2.1748, 2.3567],
//             [0.6558, 0.2469, 5.5787]])
//
//
//go:linkname Inner py.inner
func Inner(input *py.Object, other *py.Object) *py.Object
// None
//
//go:linkname InstanceNorm py.instance_norm
func InstanceNorm(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname IntRepr py.int_repr
func IntRepr(__llgo_va_list ...interface{}) *py.Object
//
// inverse(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.linalg.inv`
//
//
//go:linkname Inverse py.inverse
func Inverse(input *py.Object) *py.Object
//
// is_complex(input) -> (bool)
//
// Returns True if the data type of :attr:`input` is a complex data type i.e.,
// one of ``torch.complex64``, and ``torch.complex128``.
//
// Args:
//     input (Tensor): the input tensor.
//
//
//go:linkname IsComplex py.is_complex
func IsComplex(input *py.Object) *py.Object
//
// is_conj(input) -> (bool)
//
// Returns True if the :attr:`input` is a conjugated tensor, i.e. its conjugate bit is set to `True`.
//
// Args:
//     input (Tensor): the input tensor.
//
//
//go:linkname IsConj py.is_conj
func IsConj(input *py.Object) *py.Object
// None
//
//go:linkname IsDistributed py.is_distributed
func IsDistributed(__llgo_va_list ...interface{}) *py.Object
//
// is_floating_point(input) -> (bool)
//
// Returns True if the data type of :attr:`input` is a floating point data type i.e.,
// one of ``torch.float64``, ``torch.float32``, ``torch.float16``, and ``torch.bfloat16``.
//
// Args:
//     input (Tensor): the input tensor.
//
//
//go:linkname IsFloatingPoint py.is_floating_point
func IsFloatingPoint(input *py.Object) *py.Object
//
// is_inference(input) -> (bool)
//
// Returns True if :attr:`input` is an inference tensor.
//
// A non-view tensor is an inference tensor if and only if it was
// allocated during inference mode. A view tensor is an inference
// tensor if and only if the tensor it is a view of is an inference tensor.
//
// For details on inference mode please see
// `Inference Mode <https://pytorch.org/cppdocs/notes/inference_mode.html>`_.
//
// Args:
//     input (Tensor): the input tensor.
//
//
//go:linkname IsInference py.is_inference
func IsInference(input *py.Object) *py.Object
// None
//
//go:linkname IsNeg py.is_neg
func IsNeg(__llgo_va_list ...interface{}) *py.Object
//
// is_nonzero(input) -> (bool)
//
// Returns True if the :attr:`input` is a single element tensor which is not equal to zero
// after type conversions.
// i.e. not equal to ``torch.tensor([0.])`` or ``torch.tensor([0])`` or
// ``torch.tensor([False])``.
// Throws a ``RuntimeError`` if ``torch.numel() != 1`` (even in case
// of sparse tensors).
//
// Args:
//     input (Tensor): the input tensor.
//
// Examples::
//
//     >>> torch.is_nonzero(torch.tensor([0.]))
//     False
//     >>> torch.is_nonzero(torch.tensor([1.5]))
//     True
//     >>> torch.is_nonzero(torch.tensor([False]))
//     False
//     >>> torch.is_nonzero(torch.tensor([3]))
//     True
//     >>> torch.is_nonzero(torch.tensor([1, 3, 5]))
//     Traceback (most recent call last):
//     ...
//     RuntimeError: bool value of Tensor with more than one value is ambiguous
//     >>> torch.is_nonzero(torch.tensor([]))
//     Traceback (most recent call last):
//     ...
//     RuntimeError: bool value of Tensor with no values is ambiguous
//
//
//go:linkname IsNonzero py.is_nonzero
func IsNonzero(input *py.Object) *py.Object
// None
//
//go:linkname IsSameSize py.is_same_size
func IsSameSize(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname IsSigned py.is_signed
func IsSigned(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname IsVulkanAvailable py.is_vulkan_available
func IsVulkanAvailable(__llgo_va_list ...interface{}) *py.Object
//
// isclose(input, other, rtol=1e-05, atol=1e-08, equal_nan=False) -> Tensor
//
// Returns a new tensor with boolean elements representing if each element of
// :attr:`input` is "close" to the corresponding element of :attr:`other`.
// Closeness is defined as:
//
// .. math::
//     \lvert \text{input} - \text{other} \rvert \leq \texttt{atol} + \texttt{rtol} \times \lvert \text{other} \rvert
//
//
// where :attr:`input` and :attr:`other` are finite. Where :attr:`input`
// and/or :attr:`other` are nonfinite they are close if and only if
// they are equal, with NaNs being considered equal to each other when
// :attr:`equal_nan` is True.
//
// Args:
//     input (Tensor): first tensor to compare
//     other (Tensor): second tensor to compare
//     atol (float, optional): absolute tolerance. Default: 1e-08
//     rtol (float, optional): relative tolerance. Default: 1e-05
//     equal_nan (bool, optional): if ``True``, then two ``NaN`` s will be considered equal. Default: ``False``
//
// Examples::
//
//     >>> torch.isclose(torch.tensor((1., 2, 3)), torch.tensor((1 + 1e-10, 3, 4)))
//     tensor([ True, False, False])
//     >>> torch.isclose(torch.tensor((float('inf'), 4)), torch.tensor((float('inf'), 6)), rtol=.5)
//     tensor([True, True])
//
//
//go:linkname Isclose py.isclose
func Isclose(input *py.Object, other *py.Object, rtol *py.Object, atol *py.Object, equalNan *py.Object) *py.Object
//
// isfinite(input) -> Tensor
//
// Returns a new tensor with boolean elements representing if each element is `finite` or not.
//
// Real values are finite when they are not NaN, negative infinity, or infinity.
// Complex values are finite when both their real and imaginary parts are finite.
//
// Args:
//     input (Tensor): the input tensor.
//
// Returns:
//     A boolean tensor that is True where :attr:`input` is finite and False elsewhere
//
// Example::
//
//     >>> torch.isfinite(torch.tensor([1, float('inf'), 2, float('-inf'), float('nan')]))
//     tensor([True,  False,  True,  False,  False])
//
//
//go:linkname Isfinite py.isfinite
func Isfinite(input *py.Object) *py.Object
//
// isin(elements, test_elements, *, assume_unique=False, invert=False) -> Tensor
//
// Tests if each element of :attr:`elements` is in :attr:`test_elements`. Returns
// a boolean tensor of the same shape as :attr:`elements` that is True for elements
// in :attr:`test_elements` and False otherwise.
//
// .. note::
//     One of :attr:`elements` or :attr:`test_elements` can be a scalar, but not both.
//
// Args:
//     elements (Tensor or Scalar): Input elements
//     test_elements (Tensor or Scalar): Values against which to test for each input element
//     assume_unique (bool, optional): If True, assumes both :attr:`elements` and
//         :attr:`test_elements` contain unique elements, which can speed up the
//         calculation. Default: False
//     invert (bool, optional): If True, inverts the boolean return tensor, resulting in True
//         values for elements *not* in :attr:`test_elements`. Default: False
//
// Returns:
//     A boolean tensor of the same shape as :attr:`elements` that is True for elements in
//     :attr:`test_elements` and False otherwise
//
// Example:
//     >>> torch.isin(torch.tensor([[1, 2], [3, 4]]), torch.tensor([2, 3]))
//     tensor([[False,  True],
//             [ True, False]])
//
//
//go:linkname Isin py.isin
func Isin(elements *py.Object, testElements *py.Object) *py.Object
//
// isinf(input) -> Tensor
//
// Tests if each element of :attr:`input` is infinite
// (positive or negative infinity) or not.
//
// .. note::
//     Complex values are infinite when their real or imaginary part is
//     infinite.
//
// Args:
//     input (Tensor): the input tensor.
//
// Returns:
//     A boolean tensor that is True where :attr:`input` is infinite and False elsewhere
//
// Example::
//
//     >>> torch.isinf(torch.tensor([1, float('inf'), 2, float('-inf'), float('nan')]))
//     tensor([False,  True,  False,  True,  False])
//
//
//go:linkname Isinf py.isinf
func Isinf(input *py.Object) *py.Object
//
// isnan(input) -> Tensor
//
// Returns a new tensor with boolean elements representing if each element of :attr:`input`
// is NaN or not. Complex values are considered NaN when either their real
// and/or imaginary part is NaN.
//
// Arguments:
//     input (Tensor): the input tensor.
//
// Returns:
//     A boolean tensor that is True where :attr:`input` is NaN and False elsewhere
//
// Example::
//
//     >>> torch.isnan(torch.tensor([1, float('nan'), 2]))
//     tensor([False, True, False])
//
//
//go:linkname Isnan py.isnan
func Isnan(input *py.Object) *py.Object
//
// isneginf(input, *, out=None) -> Tensor
// Tests if each element of :attr:`input` is negative infinity or not.
//
// Args:
//   input (Tensor): the input tensor.
//
// Keyword args:
//   out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([-float('inf'), float('inf'), 1.2])
//     >>> torch.isneginf(a)
//     tensor([ True, False, False])
//
//
//go:linkname Isneginf py.isneginf
func Isneginf(input *py.Object) *py.Object
//
// isposinf(input, *, out=None) -> Tensor
// Tests if each element of :attr:`input` is positive infinity or not.
//
// Args:
//   input (Tensor): the input tensor.
//
// Keyword args:
//   out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([-float('inf'), float('inf'), 1.2])
//     >>> torch.isposinf(a)
//     tensor([False,  True, False])
//
//
//go:linkname Isposinf py.isposinf
func Isposinf(input *py.Object) *py.Object
//
// isreal(input) -> Tensor
//
// Returns a new tensor with boolean elements representing if each element of :attr:`input` is real-valued or not.
// All real-valued types are considered real. Complex values are considered real when their imaginary part is 0.
//
// Arguments:
//     input (Tensor): the input tensor.
//
// Returns:
//     A boolean tensor that is True where :attr:`input` is real and False elsewhere
//
// Example::
//
//     >>> torch.isreal(torch.tensor([1, 1+1j, 2+0j]))
//     tensor([True, False, True])
//
//
//go:linkname Isreal py.isreal
func Isreal(input *py.Object) *py.Object
// istft(input, n_fft, hop_length=None, win_length=None, window=None, center=True, normalized=False, onesided=None, length=None, return_complex=False) -> Tensor:
//
// Inverse short time Fourier Transform. This is expected to be the inverse of :func:`~torch.stft`.
//
// .. warning::
//     From version 2.1, a warning will be provided if a :attr:`window` is
//     not specified. In a future release, this attribute will be required.
//     Please provide the same window used in the stft call.
//
// It has the same parameters (+ additional optional parameter of :attr:`length`) and it should return the
// least squares estimation of the original signal. The algorithm will check using the NOLA condition (
// nonzero overlap).
//
// Important consideration in the parameters :attr:`window` and :attr:`center` so that the envelop
// created by the summation of all the windows is never zero at certain point in time. Specifically,
// :math:`\sum_{t=-\infty}^{\infty} |w|^2[n-t\times hop\_length] \cancel{=} 0`.
//
// Since :func:`~torch.stft` discards elements at the end of the signal if they do not fit in a frame,
// ``istft`` may return a shorter signal than the original signal (can occur if :attr:`center` is False
// since the signal isn't padded). If `length` is given in the arguments and is longer than expected,
// ``istft`` will pad zeros to the end of the returned signal.
//
// If :attr:`center` is ``True``, then there will be padding e.g. ``'constant'``, ``'reflect'``, etc.
// Left padding can be trimmed off exactly because they can be calculated but right padding cannot be
// calculated without additional information.
//
// Example: Suppose the last window is:
// ``[17, 18, 0, 0, 0]`` vs ``[18, 0, 0, 0, 0]``
//
// The :attr:`n_fft`, :attr:`hop_length`, :attr:`win_length` are all the same which prevents the calculation
// of right padding. These additional values could be zeros or a reflection of the signal so providing
// :attr:`length` could be useful. If :attr:`length` is ``None`` then padding will be aggressively removed
// (some loss of signal).
//
// [1] D. W. Griffin and J. S. Lim, "Signal estimation from modified short-time Fourier transform,"
// IEEE Trans. ASSP, vol.32, no.2, pp.236-243, Apr. 1984.
//
// Args:
//     input (Tensor): The input tensor. Expected to be in the format of :func:`~torch.stft`,
//         output. That is a complex tensor of shape `(B?, N, T)` where
//
//         - `B?` is an optional batch dimension
//         - `N` is the number of frequency samples, `(n_fft // 2) + 1`
//           for onesided input, or otherwise `n_fft`.
//         - `T` is the number of frames, `1 + length // hop_length` for centered stft,
//           or `1 + (length - n_fft) // hop_length` otherwise.
//
//         .. versionchanged:: 2.0
//             Real datatype inputs are no longer supported. Input must now have a
//             complex datatype, as returned by ``stft(..., return_complex=True)``.
//     n_fft (int): Size of Fourier transform
//     hop_length (Optional[int]): The distance between neighboring sliding window frames.
//         (Default: ``n_fft // 4``)
//     win_length (Optional[int]): The size of window frame and STFT filter. (Default: ``n_fft``)
//     window (Optional[torch.Tensor]): The optional window function.
//         Shape must be 1d and `<= n_fft`
//         (Default: ``torch.ones(win_length)``)
//     center (bool): Whether :attr:`input` was padded on both sides so that the :math:`t`-th frame is
//         centered at time :math:`t \times \text{hop\_length}`.
//         (Default: ``True``)
//     normalized (bool): Whether the STFT was normalized. (Default: ``False``)
//     onesided (Optional[bool]): Whether the STFT was onesided.
//         (Default: ``True`` if `n_fft != fft_size` in the input size)
//     length (Optional[int]): The amount to trim the signal by (i.e. the
//         original signal length). Defaults to `(T - 1) * hop_length` for
//         centered stft, or `n_fft + (T - 1) * hop_length` otherwise, where `T`
//         is the number of input frames.
//     return_complex (Optional[bool]):
//         Whether the output should be complex, or if the input should be
//         assumed to derive from a real signal and window.
//         Note that this is incompatible with ``onesided=True``.
//         (Default: ``False``)
//
// Returns:
//     Tensor: Least squares estimation of the original signal of shape `(B?, length)` where
//         `B?` is an optional batch dimension from the input tensor.
//
//
//go:linkname Istft py.istft
func Istft(input *py.Object, nFft *py.Object, hopLength *py.Object, winLength *py.Object, window *py.Object, center *py.Object, normalized *py.Object, onesided *py.Object, length *py.Object, returnComplex *py.Object) *py.Object
//
// kaiser_window(window_length, periodic=True, beta=12.0, *, dtype=None, layout=torch.strided, device=None, requires_grad=False) -> Tensor
//
// Computes the Kaiser window with window length :attr:`window_length` and shape parameter :attr:`beta`.
//
// Let I_0 be the zeroth order modified Bessel function of the first kind (see :func:`torch.i0`) and
// ``N = L - 1`` if :attr:`periodic` is False and ``L`` if :attr:`periodic` is True,
// where ``L`` is the :attr:`window_length`. This function computes:
//
// .. math::
//     out_i = I_0 \left( \beta \sqrt{1 - \left( {\frac{i - N/2}{N/2}} \right) ^2 } \right) / I_0( \beta )
//
// Calling ``torch.kaiser_window(L, B, periodic=True)`` is equivalent to calling
// ``torch.kaiser_window(L + 1, B, periodic=False)[:-1])``.
// The :attr:`periodic` argument is intended as a helpful shorthand
// to produce a periodic window as input to functions like :func:`torch.stft`.
//
// .. note::
//     If :attr:`window_length` is one, then the returned window is a single element tensor containing a one.
//
//
// Args:
//     window_length (int): length of the window.
//     periodic (bool, optional): If True, returns a periodic window suitable for use in spectral analysis.
//         If False, returns a symmetric window suitable for use in filter design.
//     beta (float, optional): shape parameter for the window.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`).
//     layout (:class:`torch.layout`, optional): the desired layout of returned window tensor. Only
//           ``torch.strided`` (dense layout) is supported.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
//
//
//go:linkname KaiserWindow py.kaiser_window
func KaiserWindow(windowLength *py.Object, periodic *py.Object, beta *py.Object) *py.Object
// None
//
//go:linkname KlDiv py.kl_div
func KlDiv(__llgo_va_list ...interface{}) *py.Object
//
// kron(input, other, *, out=None) -> Tensor
//
// Computes the Kronecker product, denoted by :math:`\otimes`, of :attr:`input` and :attr:`other`.
//
// If :attr:`input` is a :math:`(a_0 \times a_1 \times \dots \times a_n)` tensor and :attr:`other` is a
// :math:`(b_0 \times b_1 \times \dots \times b_n)` tensor, the result will be a
// :math:`(a_0*b_0 \times a_1*b_1 \times \dots \times a_n*b_n)` tensor with the following entries:
//
// .. math::
//     (\text{input} \otimes \text{other})_{k_0, k_1, \dots, k_n} =
//         \text{input}_{i_0, i_1, \dots, i_n} * \text{other}_{j_0, j_1, \dots, j_n},
//
// where :math:`k_t = i_t * b_t + j_t` for :math:`0 \leq t \leq n`.
// If one tensor has fewer dimensions than the other it is unsqueezed until it has the same number of dimensions.
//
// Supports real-valued and complex-valued inputs.
//
// .. note::
//     This function generalizes the typical definition of the Kronecker product for two matrices to two tensors,
//     as described above. When :attr:`input` is a :math:`(m \times n)` matrix and :attr:`other` is a
//     :math:`(p \times q)` matrix, the result will be a :math:`(p*m \times q*n)` block matrix:
//
//     .. math::
//         \mathbf{A} \otimes \mathbf{B}=\begin{bmatrix}
//         a_{11} \mathbf{B} & \cdots & a_{1 n} \mathbf{B} \\
//         \vdots & \ddots & \vdots \\
//         a_{m 1} \mathbf{B} & \cdots & a_{m n} \mathbf{B} \end{bmatrix}
//
//     where :attr:`input` is :math:`\mathbf{A}` and :attr:`other` is :math:`\mathbf{B}`.
//
// Arguments:
//     input (Tensor)
//     other (Tensor)
//
// Keyword args:
//     out (Tensor, optional): The output tensor. Ignored if ``None``. Default: ``None``
//
// Examples::
//
//     >>> mat1 = torch.eye(2)
//     >>> mat2 = torch.ones(2, 2)
//     >>> torch.kron(mat1, mat2)
//     tensor([[1., 1., 0., 0.],
//             [1., 1., 0., 0.],
//             [0., 0., 1., 1.],
//             [0., 0., 1., 1.]])
//
//     >>> mat1 = torch.eye(2)
//     >>> mat2 = torch.arange(1, 5).reshape(2, 2)
//     >>> torch.kron(mat1, mat2)
//     tensor([[1., 2., 0., 0.],
//             [3., 4., 0., 0.],
//             [0., 0., 1., 2.],
//             [0., 0., 3., 4.]])
//
//
//go:linkname Kron py.kron
func Kron(input *py.Object, other *py.Object) *py.Object
//
// kthvalue(input, k, dim=None, keepdim=False, *, out=None) -> (Tensor, LongTensor)
//
// Returns a namedtuple ``(values, indices)`` where ``values`` is the :attr:`k` th
// smallest element of each row of the :attr:`input` tensor in the given dimension
// :attr:`dim`. And ``indices`` is the index location of each element found.
//
// If :attr:`dim` is not given, the last dimension of the `input` is chosen.
//
// If :attr:`keepdim` is ``True``, both the :attr:`values` and :attr:`indices` tensors
// are the same size as :attr:`input`, except in the dimension :attr:`dim` where
// they are of size 1. Otherwise, :attr:`dim` is squeezed
// (see :func:`torch.squeeze`), resulting in both the :attr:`values` and
// :attr:`indices` tensors having 1 fewer dimension than the :attr:`input` tensor.
//
// .. note::
//     When :attr:`input` is a CUDA tensor and there are multiple valid
//     :attr:`k` th values, this function may nondeterministically return
//     :attr:`indices` for any of them.
//
// Args:
//     input (Tensor): the input tensor.
//     k (int): k for the k-th smallest element
//     dim (int, optional): the dimension to find the kth value along
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//     out (tuple, optional): the output tuple of (Tensor, LongTensor)
//                            can be optionally given to be used as output buffers
//
// Example::
//
//     >>> x = torch.arange(1., 6.)
//     >>> x
//     tensor([ 1.,  2.,  3.,  4.,  5.])
//     >>> torch.kthvalue(x, 4)
//     torch.return_types.kthvalue(values=tensor(4.), indices=tensor(3))
//
//     >>> x=torch.arange(1.,7.).resize_(2,3)
//     >>> x
//     tensor([[ 1.,  2.,  3.],
//             [ 4.,  5.,  6.]])
//     >>> torch.kthvalue(x, 2, 0, True)
//     torch.return_types.kthvalue(values=tensor([[4., 5., 6.]]), indices=tensor([[1, 1, 1]]))
//
//
//go:linkname Kthvalue py.kthvalue
func Kthvalue(input *py.Object, k *py.Object, dim *py.Object, keepdim *py.Object) *py.Object
// None
//
//go:linkname LayerNorm py.layer_norm
func LayerNorm(__llgo_va_list ...interface{}) *py.Object
//
// lcm(input, other, *, out=None) -> Tensor
//
// Computes the element-wise least common multiple (LCM) of :attr:`input` and :attr:`other`.
//
// Both :attr:`input` and :attr:`other` must have integer types.
//
// .. note::
//     This defines :math:`lcm(0, 0) = 0` and :math:`lcm(0, a) = 0`.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor): the second input tensor
//
// Keyword arguments:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([5, 10, 15])
//     >>> b = torch.tensor([3, 4, 5])
//     >>> torch.lcm(a, b)
//     tensor([15, 20, 15])
//     >>> c = torch.tensor([3])
//     >>> torch.lcm(a, c)
//     tensor([15, 30, 15])
//
//
//go:linkname Lcm py.lcm
func Lcm(input *py.Object, other *py.Object) *py.Object
// None
//
//go:linkname Lcm_ py.lcm_
func Lcm_(__llgo_va_list ...interface{}) *py.Object
//
// ldexp(input, other, *, out=None) -> Tensor
//
// Multiplies :attr:`input` by 2 ** :attr:`other`.
//
// .. math::
//     \text{{out}}_i = \text{{input}}_i * 2^\text{{other}}_i
//
//
// Typically this function is used to construct floating point numbers by multiplying
// mantissas in :attr:`input` with integral powers of two created from the exponents
// in :attr:`other`.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor): a tensor of exponents, typically integers.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.ldexp(torch.tensor([1.]), torch.tensor([1]))
//     tensor([2.])
//     >>> torch.ldexp(torch.tensor([1.0]), torch.tensor([1, 2, 3, 4]))
//     tensor([ 2.,  4.,  8., 16.])
//
//
//
//
//go:linkname Ldexp py.ldexp
func Ldexp(input *py.Object, other *py.Object) *py.Object
// None
//
//go:linkname Ldexp_ py.ldexp_
func Ldexp_(__llgo_va_list ...interface{}) *py.Object
//
// le(input, other, *, out=None) -> Tensor
//
// Computes :math:`\text{input} \leq \text{other}` element-wise.
//
//
// The second argument can be a number or a tensor whose shape is
// :ref:`broadcastable <broadcasting-semantics>` with the first argument.
//
// Args:
//     input (Tensor): the tensor to compare
//     other (Tensor or Scalar): the tensor or value to compare
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Returns:
//     A boolean tensor that is True where :attr:`input` is less than or equal to
//     :attr:`other` and False elsewhere
//
// Example::
//
//     >>> torch.le(torch.tensor([[1, 2], [3, 4]]), torch.tensor([[1, 1], [4, 4]]))
//     tensor([[True, False], [True, True]])
//
//
//go:linkname Le py.le
func Le(input *py.Object, other *py.Object) *py.Object
//
// lerp(input, end, weight, *, out=None)
//
// Does a linear interpolation of two tensors :attr:`start` (given by :attr:`input`) and :attr:`end` based
// on a scalar or tensor :attr:`weight` and returns the resulting :attr:`out` tensor.
//
// .. math::
//     \text{out}_i = \text{start}_i + \text{weight}_i \times (\text{end}_i - \text{start}_i)
//
// The shapes of :attr:`start` and :attr:`end` must be
// :ref:`broadcastable <broadcasting-semantics>`. If :attr:`weight` is a tensor, then
// the shapes of :attr:`weight`, :attr:`start`, and :attr:`end` must be :ref:`broadcastable <broadcasting-semantics>`.
//
// Args:
//     input (Tensor): the tensor with the starting points
//     end (Tensor): the tensor with the ending points
//     weight (float or tensor): the weight for the interpolation formula
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> start = torch.arange(1., 5.)
//     >>> end = torch.empty(4).fill_(10)
//     >>> start
//     tensor([ 1.,  2.,  3.,  4.])
//     >>> end
//     tensor([ 10.,  10.,  10.,  10.])
//     >>> torch.lerp(start, end, 0.5)
//     tensor([ 5.5000,  6.0000,  6.5000,  7.0000])
//     >>> torch.lerp(start, end, torch.full_like(start, 0.5))
//     tensor([ 5.5000,  6.0000,  6.5000,  7.0000])
//
//
//go:linkname Lerp py.lerp
func Lerp(input *py.Object, end *py.Object, weight *py.Object) *py.Object
//
// less(input, other, *, out=None) -> Tensor
//
// Alias for :func:`torch.lt`.
//
//
//go:linkname Less py.less
func Less(input *py.Object, other *py.Object) *py.Object
//
// less_equal(input, other, *, out=None) -> Tensor
//
// Alias for :func:`torch.le`.
//
//
//go:linkname LessEqual py.less_equal
func LessEqual(input *py.Object, other *py.Object) *py.Object
//
// lgamma(input, *, out=None) -> Tensor
//
// Computes the natural logarithm of the absolute value of the gamma function on :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \ln |\Gamma(\text{input}_{i})|
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.arange(0.5, 2, 0.5)
//     >>> torch.lgamma(a)
//     tensor([ 0.5724,  0.0000, -0.1208])
//
//
//go:linkname Lgamma py.lgamma
func Lgamma(input *py.Object) *py.Object
//
// linspace(start, end, steps, *, out=None, dtype=None, layout=torch.strided, device=None, requires_grad=False) -> Tensor
//
// Creates a one-dimensional tensor of size :attr:`steps` whose values are evenly
// spaced from :attr:`start` to :attr:`end`, inclusive. That is, the value are:
//
// .. math::
//     (\text{start},
//     \text{start} + \frac{\text{end} - \text{start}}{\text{steps} - 1},
//     \ldots,
//     \text{start} + (\text{steps} - 2) * \frac{\text{end} - \text{start}}{\text{steps} - 1},
//     \text{end})
//
//
// From PyTorch 1.11 linspace requires the steps argument. Use steps=100 to restore the previous behavior.
//
// Args:
//     start (float or Tensor): the starting value for the set of points. If `Tensor`, it must be 0-dimensional
//     end (float or Tensor): the ending value for the set of points. If `Tensor`, it must be 0-dimensional
//     steps (int): size of the constructed tensor
//
// Keyword arguments:
//     out (Tensor, optional): the output tensor.
//     dtype (torch.dtype, optional): the data type to perform the computation in.
//         Default: if None, uses the global default dtype (see torch.get_default_dtype())
//         when both :attr:`start` and :attr:`end` are real,
//         and corresponding complex dtype when either is complex.
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
//
// Example::
//
//     >>> torch.linspace(3, 10, steps=5)
//     tensor([  3.0000,   4.7500,   6.5000,   8.2500,  10.0000])
//     >>> torch.linspace(-10, 10, steps=5)
//     tensor([-10.,  -5.,   0.,   5.,  10.])
//     >>> torch.linspace(start=-10, end=10, steps=5)
//     tensor([-10.,  -5.,   0.,   5.,  10.])
//     >>> torch.linspace(start=-10, end=10, steps=1)
//     tensor([-10.])
//
//
//go:linkname Linspace py.linspace
func Linspace(start *py.Object, end *py.Object, steps *py.Object) *py.Object
//
// log(input, *, out=None) -> Tensor
//
// Returns a new tensor with the natural logarithm of the elements
// of :attr:`input`.
//
// .. math::
//     y_{i} = \log_{e} (x_{i})
//
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.rand(5) * 5
//     >>> a
//     tensor([4.7767, 4.3234, 1.2156, 0.2411, 4.5739])
//     >>> torch.log(a)
//     tensor([ 1.5637,  1.4640,  0.1952, -1.4226,  1.5204])
//
//
//go:linkname Log py.log
func Log(input *py.Object) *py.Object
//
// log10(input, *, out=None) -> Tensor
//
// Returns a new tensor with the logarithm to the base 10 of the elements
// of :attr:`input`.
//
// .. math::
//     y_{i} = \log_{10} (x_{i})
//
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.rand(5)
//     >>> a
//     tensor([ 0.5224,  0.9354,  0.7257,  0.1301,  0.2251])
//
//
//     >>> torch.log10(a)
//     tensor([-0.2820, -0.0290, -0.1392, -0.8857, -0.6476])
//
//
//
//go:linkname Log10 py.log10
func Log10(input *py.Object) *py.Object
// None
//
//go:linkname Log10_ py.log10_
func Log10_(__llgo_va_list ...interface{}) *py.Object
//
// log1p(input, *, out=None) -> Tensor
//
// Returns a new tensor with the natural logarithm of (1 + :attr:`input`).
//
// .. math::
//     y_i = \log_{e} (x_i + 1)
//
// .. note:: This function is more accurate than :func:`torch.log` for small
//           values of :attr:`input`
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(5)
//     >>> a
//     tensor([-1.0090, -0.9923,  1.0249, -0.5372,  0.2492])
//     >>> torch.log1p(a)
//     tensor([    nan, -4.8653,  0.7055, -0.7705,  0.2225])
//
//
//go:linkname Log1p py.log1p
func Log1p(input *py.Object) *py.Object
// None
//
//go:linkname Log1p_ py.log1p_
func Log1p_(__llgo_va_list ...interface{}) *py.Object
//
// log2(input, *, out=None) -> Tensor
//
// Returns a new tensor with the logarithm to the base 2 of the elements
// of :attr:`input`.
//
// .. math::
//     y_{i} = \log_{2} (x_{i})
//
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.rand(5)
//     >>> a
//     tensor([ 0.8419,  0.8003,  0.9971,  0.5287,  0.0490])
//
//
//     >>> torch.log2(a)
//     tensor([-0.2483, -0.3213, -0.0042, -0.9196, -4.3504])
//
//
//
//go:linkname Log2 py.log2
func Log2(input *py.Object) *py.Object
// None
//
//go:linkname Log2_ py.log2_
func Log2_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Log_ py.log_
func Log_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname LogSoftmax py.log_softmax
func LogSoftmax(__llgo_va_list ...interface{}) *py.Object
//
// logaddexp(input, other, *, out=None) -> Tensor
//
// Logarithm of the sum of exponentiations of the inputs.
//
// Calculates pointwise :math:`\log\left(e^x + e^y\right)`. This function is useful
// in statistics where the calculated probabilities of events may be so small as to
// exceed the range of normal floating point numbers. In such cases the logarithm
// of the calculated probability is stored. This function allows adding
// probabilities stored in such a fashion.
//
// This op should be disambiguated with :func:`torch.logsumexp` which performs a
// reduction on a single tensor.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor): the second input tensor
//
// Keyword arguments:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.logaddexp(torch.tensor([-1.0]), torch.tensor([-1.0, -2, -3]))
//     tensor([-0.3069, -0.6867, -0.8731])
//     >>> torch.logaddexp(torch.tensor([-100.0, -200, -300]), torch.tensor([-1.0, -2, -3]))
//     tensor([-1., -2., -3.])
//     >>> torch.logaddexp(torch.tensor([1.0, 2000, 30000]), torch.tensor([-1.0, -2, -3]))
//     tensor([1.1269e+00, 2.0000e+03, 3.0000e+04])
//
//
//go:linkname Logaddexp py.logaddexp
func Logaddexp(input *py.Object, other *py.Object) *py.Object
//
// logaddexp2(input, other, *, out=None) -> Tensor
//
// Logarithm of the sum of exponentiations of the inputs in base-2.
//
// Calculates pointwise :math:`\log_2\left(2^x + 2^y\right)`. See
// :func:`torch.logaddexp` for more details.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor): the second input tensor
//
// Keyword arguments:
//     out (Tensor, optional): the output tensor.
//
//
//go:linkname Logaddexp2 py.logaddexp2
func Logaddexp2(input *py.Object, other *py.Object) *py.Object
//
// logcumsumexp(input, dim, *, out=None) -> Tensor
// Returns the logarithm of the cumulative summation of the exponentiation of
// elements of :attr:`input` in the dimension :attr:`dim`.
//
// For summation index :math:`j` given by `dim` and other indices :math:`i`, the result is
//
//     .. math::
//         \text{logcumsumexp}(x)_{ij} = \log \sum\limits_{j=0}^{i} \exp(x_{ij})
//
// Args:
//     input (Tensor): the input tensor.
//     dim  (int): the dimension to do the operation over
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(10)
//     >>> torch.logcumsumexp(a, dim=0)
//     tensor([-0.42296738, -0.04462666,  0.86278635,  0.94622083,  1.05277811,
//              1.39202815,  1.83525007,  1.84492621,  2.06084887,  2.06844475]))
//
//
//go:linkname Logcumsumexp py.logcumsumexp
func Logcumsumexp(input *py.Object, dim *py.Object) *py.Object
//
// logdet(input) -> Tensor
//
// Calculates log determinant of a square matrix or batches of square matrices.
//
// It returns ``-inf`` if the input has a determinant of zero, and ``NaN`` if it has
// a negative determinant.
//
// .. note::
//     Backward through :meth:`logdet` internally uses SVD results when :attr:`input`
//     is not invertible. In this case, double backward through :meth:`logdet` will
//     be unstable in when :attr:`input` doesn't have distinct singular values. See
//     :func:`torch.linalg.svd` for details.
//
// .. seealso::
//
//         :func:`torch.linalg.slogdet` computes the sign (resp. angle) and natural logarithm of the
//         absolute value of the determinant of real-valued (resp. complex) square matrices.
//
// Arguments:
//     input (Tensor): the input tensor of size ``(*, n, n)`` where ``*`` is zero or more
//                 batch dimensions.
//
// Example::
//
//     >>> A = torch.randn(3, 3)
//     >>> torch.det(A)
//     tensor(0.2611)
//     >>> torch.logdet(A)
//     tensor(-1.3430)
//     >>> A
//     tensor([[[ 0.9254, -0.6213],
//              [-0.5787,  1.6843]],
//
//             [[ 0.3242, -0.9665],
//              [ 0.4539, -0.0887]],
//
//             [[ 1.1336, -0.4025],
//              [-0.7089,  0.9032]]])
//     >>> A.det()
//     tensor([1.1990, 0.4099, 0.7386])
//     >>> A.det().log()
//     tensor([ 0.1815, -0.8917, -0.3031])
//
//
//go:linkname Logdet py.logdet
func Logdet(input *py.Object) *py.Object
//
// logical_and(input, other, *, out=None) -> Tensor
//
// Computes the element-wise logical AND of the given input tensors. Zeros are treated as ``False`` and nonzeros are
// treated as ``True``.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor): the tensor to compute AND with
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.logical_and(torch.tensor([True, False, True]), torch.tensor([True, False, False]))
//     tensor([ True, False, False])
//     >>> a = torch.tensor([0, 1, 10, 0], dtype=torch.int8)
//     >>> b = torch.tensor([4, 0, 1, 0], dtype=torch.int8)
//     >>> torch.logical_and(a, b)
//     tensor([False, False,  True, False])
//     >>> torch.logical_and(a.double(), b.double())
//     tensor([False, False,  True, False])
//     >>> torch.logical_and(a.double(), b)
//     tensor([False, False,  True, False])
//     >>> torch.logical_and(a, b, out=torch.empty(4, dtype=torch.bool))
//     tensor([False, False,  True, False])
//
//
//go:linkname LogicalAnd py.logical_and
func LogicalAnd(input *py.Object, other *py.Object) *py.Object
//
// logical_not(input, *, out=None) -> Tensor
//
// Computes the element-wise logical NOT of the given input tensor. If not specified, the output tensor will have the bool
// dtype. If the input tensor is not a bool tensor, zeros are treated as ``False`` and non-zeros are treated as ``True``.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.logical_not(torch.tensor([True, False]))
//     tensor([False,  True])
//     >>> torch.logical_not(torch.tensor([0, 1, -10], dtype=torch.int8))
//     tensor([ True, False, False])
//     >>> torch.logical_not(torch.tensor([0., 1.5, -10.], dtype=torch.double))
//     tensor([ True, False, False])
//     >>> torch.logical_not(torch.tensor([0., 1., -10.], dtype=torch.double), out=torch.empty(3, dtype=torch.int16))
//     tensor([1, 0, 0], dtype=torch.int16)
//
//
//go:linkname LogicalNot py.logical_not
func LogicalNot(input *py.Object) *py.Object
//
// logical_or(input, other, *, out=None) -> Tensor
//
// Computes the element-wise logical OR of the given input tensors. Zeros are treated as ``False`` and nonzeros are
// treated as ``True``.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor): the tensor to compute OR with
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.logical_or(torch.tensor([True, False, True]), torch.tensor([True, False, False]))
//     tensor([ True, False,  True])
//     >>> a = torch.tensor([0, 1, 10, 0], dtype=torch.int8)
//     >>> b = torch.tensor([4, 0, 1, 0], dtype=torch.int8)
//     >>> torch.logical_or(a, b)
//     tensor([ True,  True,  True, False])
//     >>> torch.logical_or(a.double(), b.double())
//     tensor([ True,  True,  True, False])
//     >>> torch.logical_or(a.double(), b)
//     tensor([ True,  True,  True, False])
//     >>> torch.logical_or(a, b, out=torch.empty(4, dtype=torch.bool))
//     tensor([ True,  True,  True, False])
//
//
//go:linkname LogicalOr py.logical_or
func LogicalOr(input *py.Object, other *py.Object) *py.Object
//
// logical_xor(input, other, *, out=None) -> Tensor
//
// Computes the element-wise logical XOR of the given input tensors. Zeros are treated as ``False`` and nonzeros are
// treated as ``True``.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor): the tensor to compute XOR with
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.logical_xor(torch.tensor([True, False, True]), torch.tensor([True, False, False]))
//     tensor([False, False,  True])
//     >>> a = torch.tensor([0, 1, 10, 0], dtype=torch.int8)
//     >>> b = torch.tensor([4, 0, 1, 0], dtype=torch.int8)
//     >>> torch.logical_xor(a, b)
//     tensor([ True,  True, False, False])
//     >>> torch.logical_xor(a.double(), b.double())
//     tensor([ True,  True, False, False])
//     >>> torch.logical_xor(a.double(), b)
//     tensor([ True,  True, False, False])
//     >>> torch.logical_xor(a, b, out=torch.empty(4, dtype=torch.bool))
//     tensor([ True,  True, False, False])
//
//
//go:linkname LogicalXor py.logical_xor
func LogicalXor(input *py.Object, other *py.Object) *py.Object
//
// logit(input, eps=None, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.logit`.
//
//
//go:linkname Logit py.logit
func Logit(input *py.Object, eps *py.Object) *py.Object
// None
//
//go:linkname Logit_ py.logit_
func Logit_(__llgo_va_list ...interface{}) *py.Object
//
// logspace(start, end, steps, base=10.0, *,          out=None, dtype=None, layout=torch.strided, device=None, requires_grad=False) -> Tensor
//
//
// Creates a one-dimensional tensor of size :attr:`steps` whose values are evenly
// spaced from :math:`{{\text{{base}}}}^{{\text{{start}}}}` to
// :math:`{{\text{{base}}}}^{{\text{{end}}}}`, inclusive, on a logarithmic scale
// with base :attr:`base`. That is, the values are:
//
// .. math::
//     (\text{base}^{\text{start}},
//     \text{base}^{(\text{start} + \frac{\text{end} - \text{start}}{ \text{steps} - 1})},
//     \ldots,
//     \text{base}^{(\text{start} + (\text{steps} - 2) * \frac{\text{end} - \text{start}}{ \text{steps} - 1})},
//     \text{base}^{\text{end}})
//
//
//
// From PyTorch 1.11 logspace requires the steps argument. Use steps=100 to restore the previous behavior.
//
// Args:
//     start (float or Tensor): the starting value for the set of points. If `Tensor`, it must be 0-dimensional
//     end (float or Tensor): the ending value for the set of points. If `Tensor`, it must be 0-dimensional
//     steps (int): size of the constructed tensor
//     base (float, optional): base of the logarithm function. Default: ``10.0``.
//
// Keyword arguments:
//     out (Tensor, optional): the output tensor.
//     dtype (torch.dtype, optional): the data type to perform the computation in.
//         Default: if None, uses the global default dtype (see torch.get_default_dtype())
//         when both :attr:`start` and :attr:`end` are real,
//         and corresponding complex dtype when either is complex.
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
// Example::
//
//     >>> torch.logspace(start=-10, end=10, steps=5)
//     tensor([ 1.0000e-10,  1.0000e-05,  1.0000e+00,  1.0000e+05,  1.0000e+10])
//     >>> torch.logspace(start=0.1, end=1.0, steps=5)
//     tensor([  1.2589,   2.1135,   3.5481,   5.9566,  10.0000])
//     >>> torch.logspace(start=0.1, end=1.0, steps=1)
//     tensor([1.2589])
//     >>> torch.logspace(start=2, end=2, steps=1, base=2)
//     tensor([4.0])
//
//
//go:linkname Logspace py.logspace
func Logspace(start *py.Object, end *py.Object, steps *py.Object, base *py.Object) *py.Object
//
// logsumexp(input, dim, keepdim=False, *, out=None)
//
// Returns the log of summed exponentials of each row of the :attr:`input`
// tensor in the given dimension :attr:`dim`. The computation is numerically
// stabilized.
//
// For summation index :math:`j` given by `dim` and other indices :math:`i`, the result is
//
//     .. math::
//         \text{logsumexp}(x)_{i} = \log \sum_j \exp(x_{ij})
//
//
// If :attr:`keepdim` is ``True``, the output tensor is of the same size
// as :attr:`input` except in the dimension(s) :attr:`dim` where it is of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in the
// output tensor having 1 (or ``len(dim)``) fewer dimension(s).
//
//
// Args:
//     input (Tensor): the input tensor.
//
//     dim (int or tuple of ints, optional): the dimension or dimensions to reduce.
//         If ``None``, all dimensions are reduced.
//
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(3, 3)
//     >>> torch.logsumexp(a, 1)
//     tensor([1.4907, 1.0593, 1.5696])
//     >>> torch.dist(torch.logsumexp(a, 1), torch.log(torch.sum(torch.exp(a), 1)))
//     tensor(1.6859e-07)
//
//
//go:linkname Logsumexp py.logsumexp
func Logsumexp(input *py.Object, dim *py.Object, keepdim *py.Object) *py.Object
// None
//
//go:linkname Lstm py.lstm
func Lstm(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname LstmCell py.lstm_cell
func LstmCell(__llgo_va_list ...interface{}) *py.Object
//
// lt(input, other, *, out=None) -> Tensor
//
// Computes :math:`\text{input} < \text{other}` element-wise.
//
//
// The second argument can be a number or a tensor whose shape is
// :ref:`broadcastable <broadcasting-semantics>` with the first argument.
//
// Args:
//     input (Tensor): the tensor to compare
//     other (Tensor or float): the tensor or value to compare
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Returns:
//     A boolean tensor that is True where :attr:`input` is less than :attr:`other` and False elsewhere
//
// Example::
//
//     >>> torch.lt(torch.tensor([[1, 2], [3, 4]]), torch.tensor([[1, 1], [4, 4]]))
//     tensor([[False, False], [True, False]])
//
//
//go:linkname Lt py.lt
func Lt(input *py.Object, other *py.Object) *py.Object
//
// lu_solve(b, LU_data, LU_pivots, *, out=None) -> Tensor
//
// Returns the LU solve of the linear system :math:`Ax = b` using the partially pivoted
// LU factorization of A from :func:`~linalg.lu_factor`.
//
// This function supports ``float``, ``double``, ``cfloat`` and ``cdouble`` dtypes for :attr:`input`.
//
// .. warning::
//
//     :func:`torch.lu_solve` is deprecated in favor of :func:`torch.linalg.lu_solve`.
//     :func:`torch.lu_solve` will be removed in a future PyTorch release.
//     ``X = torch.lu_solve(B, LU, pivots)`` should be replaced with
//
//     .. code:: python
//
//         X = linalg.lu_solve(LU, pivots, B)
//
// Arguments:
//     b (Tensor): the RHS tensor of size :math:`(*, m, k)`, where :math:`*`
//                 is zero or more batch dimensions.
//     LU_data (Tensor): the pivoted LU factorization of A from :meth:`~linalg.lu_factor` of size :math:`(*, m, m)`,
//                        where :math:`*` is zero or more batch dimensions.
//     LU_pivots (IntTensor): the pivots of the LU factorization from :meth:`~linalg.lu_factor` of size :math:`(*, m)`,
//                            where :math:`*` is zero or more batch dimensions.
//                            The batch dimensions of :attr:`LU_pivots` must be equal to the batch dimensions of
//                            :attr:`LU_data`.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> A = torch.randn(2, 3, 3)
//     >>> b = torch.randn(2, 3, 1)
//     >>> LU, pivots = torch.linalg.lu_factor(A)
//     >>> x = torch.lu_solve(b, LU, pivots)
//     >>> torch.dist(A @ x, b)
//     tensor(1.00000e-07 *
//            2.8312)
//
//
//go:linkname LuSolve py.lu_solve
func LuSolve(b *py.Object, LUData *py.Object, LUPivots *py.Object) *py.Object
//
// lu_unpack(LU_data, LU_pivots, unpack_data=True, unpack_pivots=True, *, out=None) -> (Tensor, Tensor, Tensor)
//
// Unpacks the LU decomposition returned by :func:`~linalg.lu_factor` into the `P, L, U` matrices.
//
// .. seealso::
//
//     :func:`~linalg.lu` returns the matrices from the LU decomposition. Its gradient formula is more efficient
//     than that of doing :func:`~linalg.lu_factor` followed by :func:`~linalg.lu_unpack`.
//
// Args:
//     LU_data (Tensor): the packed LU factorization data
//     LU_pivots (Tensor): the packed LU factorization pivots
//     unpack_data (bool): flag indicating if the data should be unpacked.
//                         If ``False``, then the returned ``L`` and ``U`` are empty tensors.
//                         Default: ``True``
//     unpack_pivots (bool): flag indicating if the pivots should be unpacked into a permutation matrix ``P``.
//                           If ``False``, then the returned ``P`` is  an empty tensor.
//                           Default: ``True``
//
// Keyword args:
//     out (tuple, optional): output tuple of three tensors. Ignored if `None`.
//
// Returns:
//     A namedtuple ``(P, L, U)``
//
// Examples::
//
//     >>> A = torch.randn(2, 3, 3)
//     >>> LU, pivots = torch.linalg.lu_factor(A)
//     >>> P, L, U = torch.lu_unpack(LU, pivots)
//     >>> # We can recover A from the factorization
//     >>> A_ = P @ L @ U
//     >>> torch.allclose(A, A_)
//     True
//
//     >>> # LU factorization of a rectangular matrix:
//     >>> A = torch.randn(2, 3, 2)
//     >>> LU, pivots = torch.linalg.lu_factor(A)
//     >>> P, L, U = torch.lu_unpack(LU, pivots)
//     >>> # P, L, U are the same as returned by linalg.lu
//     >>> P_, L_, U_ = torch.linalg.lu(A)
//     >>> torch.allclose(P, P_) and torch.allclose(L, L_) and torch.allclose(U, U_)
//     True
//
//
//
//go:linkname LuUnpack py.lu_unpack
func LuUnpack(LUData *py.Object, LUPivots *py.Object, unpackData *py.Object, unpackPivots *py.Object) *py.Object
// None
//
//go:linkname MarginRankingLoss py.margin_ranking_loss
func MarginRankingLoss(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MaskedFill py.masked_fill
func MaskedFill(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MaskedScatter py.masked_scatter
func MaskedScatter(__llgo_va_list ...interface{}) *py.Object
//
// masked_select(input, mask, *, out=None) -> Tensor
//
// Returns a new 1-D tensor which indexes the :attr:`input` tensor according to
// the boolean mask :attr:`mask` which is a `BoolTensor`.
//
// The shapes of the :attr:`mask` tensor and the :attr:`input` tensor don't need
// to match, but they must be :ref:`broadcastable <broadcasting-semantics>`.
//
// .. note:: The returned tensor does **not** use the same storage
//           as the original tensor
//
// Args:
//     input (Tensor): the input tensor.
//     mask  (BoolTensor): the tensor containing the binary mask to index with
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> x = torch.randn(3, 4)
//     >>> x
//     tensor([[ 0.3552, -2.3825, -0.8297,  0.3477],
//             [-1.2035,  1.2252,  0.5002,  0.6248],
//             [ 0.1307, -2.0608,  0.1244,  2.0139]])
//     >>> mask = x.ge(0.5)
//     >>> mask
//     tensor([[False, False, False, False],
//             [False, True, True, True],
//             [False, False, False, True]])
//     >>> torch.masked_select(x, mask)
//     tensor([ 1.2252,  0.5002,  0.6248,  2.0139])
//
//
//go:linkname MaskedSelect py.masked_select
func MaskedSelect(input *py.Object, mask *py.Object) *py.Object
//
// matmul(input, other, *, out=None) -> Tensor
//
// Matrix product of two tensors.
//
// The behavior depends on the dimensionality of the tensors as follows:
//
// - If both tensors are 1-dimensional, the dot product (scalar) is returned.
// - If both arguments are 2-dimensional, the matrix-matrix product is returned.
// - If the first argument is 1-dimensional and the second argument is 2-dimensional,
//   a 1 is prepended to its dimension for the purpose of the matrix multiply.
//   After the matrix multiply, the prepended dimension is removed.
// - If the first argument is 2-dimensional and the second argument is 1-dimensional,
//   the matrix-vector product is returned.
// - If both arguments are at least 1-dimensional and at least one argument is
//   N-dimensional (where N > 2), then a batched matrix multiply is returned.  If the first
//   argument is 1-dimensional, a 1 is prepended to its dimension for the purpose of the
//   batched matrix multiply and removed after.  If the second argument is 1-dimensional, a
//   1 is appended to its dimension for the purpose of the batched matrix multiple and removed after.
//   The non-matrix (i.e. batch) dimensions are :ref:`broadcasted <broadcasting-semantics>` (and thus
//   must be broadcastable).  For example, if :attr:`input` is a
//   :math:`(j \times 1 \times n \times n)` tensor and :attr:`other` is a :math:`(k \times n \times n)`
//   tensor, :attr:`out` will be a :math:`(j \times k \times n \times n)` tensor.
//
//   Note that the broadcasting logic only looks at the batch dimensions when determining if the inputs
//   are broadcastable, and not the matrix dimensions. For example, if :attr:`input` is a
//   :math:`(j \times 1 \times n \times m)` tensor and :attr:`other` is a :math:`(k \times m \times p)`
//   tensor, these inputs are valid for broadcasting even though the final two dimensions (i.e. the
//   matrix dimensions) are different. :attr:`out` will be a :math:`(j \times k \times n \times p)` tensor.
//
// This operation has support for arguments with :ref:`sparse layouts<sparse-docs>`. In particular the
// matrix-matrix (both arguments 2-dimensional) supports sparse arguments with the same restrictions
// as :func:`torch.mm`
//
//
// .. warning::
//     Sparse support is a beta feature and some layout(s)/dtype/device combinations may not be supported,
//     or may not have autograd support. If you notice missing functionality please
//     open a feature request.
//
// This operator supports :ref:`TensorFloat32<tf32_on_ampere>`.
//
// On certain ROCm devices, when using float16 inputs this module will use :ref:`different precision<fp16_on_mi200>` for backward.
//
// .. note::
//
//     The 1-dimensional dot product version of this function does not support an :attr:`out` parameter.
//
// Arguments:
//     input (Tensor): the first tensor to be multiplied
//     other (Tensor): the second tensor to be multiplied
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> # vector x vector
//     >>> tensor1 = torch.randn(3)
//     >>> tensor2 = torch.randn(3)
//     >>> torch.matmul(tensor1, tensor2).size()
//     torch.Size([])
//     >>> # matrix x vector
//     >>> tensor1 = torch.randn(3, 4)
//     >>> tensor2 = torch.randn(4)
//     >>> torch.matmul(tensor1, tensor2).size()
//     torch.Size([3])
//     >>> # batched matrix x broadcasted vector
//     >>> tensor1 = torch.randn(10, 3, 4)
//     >>> tensor2 = torch.randn(4)
//     >>> torch.matmul(tensor1, tensor2).size()
//     torch.Size([10, 3])
//     >>> # batched matrix x batched matrix
//     >>> tensor1 = torch.randn(10, 3, 4)
//     >>> tensor2 = torch.randn(10, 4, 5)
//     >>> torch.matmul(tensor1, tensor2).size()
//     torch.Size([10, 3, 5])
//     >>> # batched matrix x broadcasted matrix
//     >>> tensor1 = torch.randn(10, 3, 4)
//     >>> tensor2 = torch.randn(4, 5)
//     >>> torch.matmul(tensor1, tensor2).size()
//     torch.Size([10, 3, 5])
//
//
//
//go:linkname Matmul py.matmul
func Matmul(input *py.Object, other *py.Object) *py.Object
//
// matrix_exp(A) -> Tensor
//
// Alias for :func:`torch.linalg.matrix_exp`.
//
//
//go:linkname MatrixExp py.matrix_exp
func MatrixExp(A *py.Object) *py.Object
//
// matrix_power(input, n, *, out=None) -> Tensor
//
// Alias for :func:`torch.linalg.matrix_power`
//
//
//go:linkname MatrixPower py.matrix_power
func MatrixPower(input *py.Object, n *py.Object) *py.Object
//
// max(input) -> Tensor
//
// Returns the maximum value of all elements in the ``input`` tensor.
//
// .. warning::
//     This function produces deterministic (sub)gradients unlike ``max(dim=0)``
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> a = torch.randn(1, 3)
//     >>> a
//     tensor([[ 0.6763,  0.7445, -2.2369]])
//     >>> torch.max(a)
//     tensor(0.7445)
//
// .. function:: max(input, dim, keepdim=False, *, out=None) -> (Tensor, LongTensor)
//    :noindex:
//
// Returns a namedtuple ``(values, indices)`` where ``values`` is the maximum
// value of each row of the :attr:`input` tensor in the given dimension
// :attr:`dim`. And ``indices`` is the index location of each maximum value found
// (argmax).
//
// If ``keepdim`` is ``True``, the output tensors are of the same size
// as ``input`` except in the dimension ``dim`` where they are of size 1.
// Otherwise, ``dim`` is squeezed (see :func:`torch.squeeze`), resulting
// in the output tensors having 1 fewer dimension than ``input``.
//
// .. note:: If there are multiple maximal values in a reduced row then
//           the indices of the first maximal value are returned.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int): the dimension to reduce.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not. Default: ``False``.
//
// Keyword args:
//     out (tuple, optional): the result tuple of two output tensors (max, max_indices)
//
// Example::
//
//     >>> a = torch.randn(4, 4)
//     >>> a
//     tensor([[-1.2360, -0.2942, -0.1222,  0.8475],
//             [ 1.1949, -1.1127, -2.2379, -0.6702],
//             [ 1.5717, -0.9207,  0.1297, -1.8768],
//             [-0.6172,  1.0036, -0.6060, -0.2432]])
//     >>> torch.max(a, 1)
//     torch.return_types.max(values=tensor([0.8475, 1.1949, 1.5717, 1.0036]), indices=tensor([3, 0, 0, 1]))
//
// .. function:: max(input, other, *, out=None) -> Tensor
//    :noindex:
//
// See :func:`torch.maximum`.
//
//
//
//go:linkname Max py.max
func Max(input *py.Object) *py.Object
// None
//
//go:linkname MaxPool1d py.max_pool1d
func MaxPool1d(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MaxPool1dWithIndices py.max_pool1d_with_indices
func MaxPool1dWithIndices(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MaxPool2d py.max_pool2d
func MaxPool2d(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MaxPool3d py.max_pool3d
func MaxPool3d(__llgo_va_list ...interface{}) *py.Object
//
// maximum(input, other, *, out=None) -> Tensor
//
// Computes the element-wise maximum of :attr:`input` and :attr:`other`.
//
// .. note::
//     If one of the elements being compared is a NaN, then that element is returned.
//     :func:`maximum` is not supported for tensors with complex dtypes.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor): the second input tensor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor((1, 2, -1))
//     >>> b = torch.tensor((3, 0, 4))
//     >>> torch.maximum(a, b)
//     tensor([3, 2, 4])
//
//
//go:linkname Maximum py.maximum
func Maximum(input *py.Object, other *py.Object) *py.Object
//
// mean(input, *, dtype=None) -> Tensor
//
// Returns the mean value of all elements in the :attr:`input` tensor.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         If specified, the input tensor is casted to :attr:`dtype` before the operation
//         is performed. This is useful for preventing data type overflows. Default: None.
//
// Example::
//
//     >>> a = torch.randn(1, 3)
//     >>> a
//     tensor([[ 0.2294, -0.5481,  1.3288]])
//     >>> torch.mean(a)
//     tensor(0.3367)
//
// .. function:: mean(input, dim, keepdim=False, *, dtype=None, out=None) -> Tensor
//    :noindex:
//
// Returns the mean value of each row of the :attr:`input` tensor in the given
// dimension :attr:`dim`. If :attr:`dim` is a list of dimensions,
// reduce over all of them.
//
//
// If :attr:`keepdim` is ``True``, the output tensor is of the same size
// as :attr:`input` except in the dimension(s) :attr:`dim` where it is of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in the
// output tensor having 1 (or ``len(dim)``) fewer dimension(s).
//
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int or tuple of ints): the dimension or dimensions to reduce.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         If specified, the input tensor is casted to :attr:`dtype` before the operation
//         is performed. This is useful for preventing data type overflows. Default: None.
//     out (Tensor, optional): the output tensor.
//
// .. seealso::
//
//     :func:`torch.nanmean` computes the mean value of `non-NaN` elements.
//
// Example::
//
//     >>> a = torch.randn(4, 4)
//     >>> a
//     tensor([[-0.3841,  0.6320,  0.4254, -0.7384],
//             [-0.9644,  1.0131, -0.6549, -1.4279],
//             [-0.2951, -1.3350, -0.7694,  0.5600],
//             [ 1.0842, -0.9580,  0.3623,  0.2343]])
//     >>> torch.mean(a, 1)
//     tensor([-0.0163, -0.5085, -0.4599,  0.1807])
//     >>> torch.mean(a, 1, True)
//     tensor([[-0.0163],
//             [-0.5085],
//             [-0.4599],
//             [ 0.1807]])
//
//
//go:linkname Mean py.mean
func Mean(input *py.Object) *py.Object
//
// median(input) -> Tensor
//
// Returns the median of the values in :attr:`input`.
//
// .. note::
//     The median is not unique for :attr:`input` tensors with an even number
//     of elements. In this case the lower of the two medians is returned. To
//     compute the mean of both medians, use :func:`torch.quantile` with ``q=0.5`` instead.
//
// .. warning::
//     This function produces deterministic (sub)gradients unlike ``median(dim=0)``
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> a = torch.randn(1, 3)
//     >>> a
//     tensor([[ 1.5219, -1.5212,  0.2202]])
//     >>> torch.median(a)
//     tensor(0.2202)
//
// .. function:: median(input, dim=-1, keepdim=False, *, out=None) -> (Tensor, LongTensor)
//    :noindex:
//
// Returns a namedtuple ``(values, indices)`` where ``values`` contains the median of each row of :attr:`input`
// in the dimension :attr:`dim`, and ``indices`` contains the index of the median values found in the dimension :attr:`dim`.
//
// By default, :attr:`dim` is the last dimension of the :attr:`input` tensor.
//
// If :attr:`keepdim` is ``True``, the output tensors are of the same size
// as :attr:`input` except in the dimension :attr:`dim` where they are of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in
// the outputs tensor having 1 fewer dimension than :attr:`input`.
//
// .. note::
//     The median is not unique for :attr:`input` tensors with an even number
//     of elements in the dimension :attr:`dim`. In this case the lower of the
//     two medians is returned. To compute the mean of both medians in
//     :attr:`input`, use :func:`torch.quantile` with ``q=0.5`` instead.
//
// .. warning::
//     ``indices`` does not necessarily contain the first occurrence of each
//     median value found, unless it is unique.
//     The exact implementation details are device-specific.
//     Do not expect the same result when run on CPU and GPU in general.
//     For the same reason do not expect the gradients to be deterministic.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int): the dimension to reduce.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//     out ((Tensor, Tensor), optional): The first tensor will be populated with the median values and the second
//                                       tensor, which must have dtype long, with their indices in the dimension
//                                       :attr:`dim` of :attr:`input`.
//
// Example::
//
//     >>> a = torch.randn(4, 5)
//     >>> a
//     tensor([[ 0.2505, -0.3982, -0.9948,  0.3518, -1.3131],
//             [ 0.3180, -0.6993,  1.0436,  0.0438,  0.2270],
//             [-0.2751,  0.7303,  0.2192,  0.3321,  0.2488],
//             [ 1.0778, -1.9510,  0.7048,  0.4742, -0.7125]])
//     >>> torch.median(a, 1)
//     torch.return_types.median(values=tensor([-0.3982,  0.2270,  0.2488,  0.4742]), indices=tensor([1, 4, 4, 3]))
//
//
//go:linkname Median py.median
func Median(input *py.Object) *py.Object
// Creates grids of coordinates specified by the 1D inputs in `attr`:tensors.
//
//         This is helpful when you want to visualize data over some
//         range of inputs. See below for a plotting example.
//
//         Given :math:`N` 1D tensors :math:`T_0 \ldots T_{N-1}` as
//         inputs with corresponding sizes :math:`S_0 \ldots S_{N-1}`,
//         this creates :math:`N` N-dimensional tensors :math:`G_0 \ldots
//         G_{N-1}`, each with shape :math:`(S_0, ..., S_{N-1})` where
//         the output :math:`G_i` is constructed by expanding :math:`T_i`
//         to the result shape.
//
//         .. note::
//             0D inputs are treated equivalently to 1D inputs of a
//             single element.
//
//         .. warning::
//             `torch.meshgrid(*tensors)` currently has the same behavior
//             as calling `numpy.meshgrid(*arrays, indexing='ij')`.
//
//             In the future `torch.meshgrid` will transition to
//             `indexing='xy'` as the default.
//
//             https://github.com/pytorch/pytorch/issues/50276 tracks
//             this issue with the goal of migrating to NumPy's behavior.
//
//         .. seealso::
//
//             :func:`torch.cartesian_prod` has the same effect but it
//             collects the data in a tensor of vectors.
//
//         Args:
//             tensors (list of Tensor): list of scalars or 1 dimensional tensors. Scalars will be
//                 treated as tensors of size :math:`(1,)` automatically
//
//             indexing: (str, optional): the indexing mode, either "xy"
//                 or "ij", defaults to "ij". See warning for future changes.
//
//                 If "xy" is selected, the first dimension corresponds
//                 to the cardinality of the second input and the second
//                 dimension corresponds to the cardinality of the first
//                 input.
//
//                 If "ij" is selected, the dimensions are in the same
//                 order as the cardinality of the inputs.
//
//         Returns:
//             seq (sequence of Tensors): If the input has :math:`N`
//             tensors of size :math:`S_0 \ldots S_{N-1}``, then the
//             output will also have :math:`N` tensors, where each tensor
//             is of shape :math:`(S_0, ..., S_{N-1})`.
//
//         Example::
//
//             >>> x = torch.tensor([1, 2, 3])
//             >>> y = torch.tensor([4, 5, 6])
//
//             Observe the element-wise pairings across the grid, (1, 4),
//             (1, 5), ..., (3, 6). This is the same thing as the
//             cartesian product.
//             >>> grid_x, grid_y = torch.meshgrid(x, y, indexing='ij')
//             >>> grid_x
//             tensor([[1, 1, 1],
//                     [2, 2, 2],
//                     [3, 3, 3]])
//             >>> grid_y
//             tensor([[4, 5, 6],
//                     [4, 5, 6],
//                     [4, 5, 6]])
//
//             This correspondence can be seen when these grids are
//             stacked properly.
//             >>> torch.equal(torch.cat(tuple(torch.dstack([grid_x, grid_y]))),
//             ...             torch.cartesian_prod(x, y))
//             True
//
//             `torch.meshgrid` is commonly used to produce a grid for
//             plotting.
//             >>> # xdoctest: +REQUIRES(module:matplotlib)
//             >>> # xdoctest: +REQUIRES(env:DOCTEST_SHOW)
//             >>> import matplotlib.pyplot as plt
//             >>> xs = torch.linspace(-5, 5, steps=100)
//             >>> ys = torch.linspace(-5, 5, steps=100)
//             >>> x, y = torch.meshgrid(xs, ys, indexing='xy')
//             >>> z = torch.sin(torch.sqrt(x * x + y * y))
//             >>> ax = plt.axes(projection='3d')
//             >>> ax.plot_surface(x.numpy(), y.numpy(), z.numpy())
//             >>> plt.show()
//
//         .. image:: ../_static/img/meshgrid.png
//             :width: 512
//
//
//
//go:linkname Meshgrid py.meshgrid
func Meshgrid(__llgo_va_list ...interface{}) *py.Object
//
// min(input) -> Tensor
//
// Returns the minimum value of all elements in the :attr:`input` tensor.
//
// .. warning::
//     This function produces deterministic (sub)gradients unlike ``min(dim=0)``
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> a = torch.randn(1, 3)
//     >>> a
//     tensor([[ 0.6750,  1.0857,  1.7197]])
//     >>> torch.min(a)
//     tensor(0.6750)
//
// .. function:: min(input, dim, keepdim=False, *, out=None) -> (Tensor, LongTensor)
//    :noindex:
//
// Returns a namedtuple ``(values, indices)`` where ``values`` is the minimum
// value of each row of the :attr:`input` tensor in the given dimension
// :attr:`dim`. And ``indices`` is the index location of each minimum value found
// (argmin).
//
// If :attr:`keepdim` is ``True``, the output tensors are of the same size as
// :attr:`input` except in the dimension :attr:`dim` where they are of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in
// the output tensors having 1 fewer dimension than :attr:`input`.
//
// .. note:: If there are multiple minimal values in a reduced row then
//           the indices of the first minimal value are returned.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int): the dimension to reduce.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//     out (tuple, optional): the tuple of two output tensors (min, min_indices)
//
// Example::
//
//     >>> a = torch.randn(4, 4)
//     >>> a
//     tensor([[-0.6248,  1.1334, -1.1899, -0.2803],
//             [-1.4644, -0.2635, -0.3651,  0.6134],
//             [ 0.2457,  0.0384,  1.0128,  0.7015],
//             [-0.1153,  2.9849,  2.1458,  0.5788]])
//     >>> torch.min(a, 1)
//     torch.return_types.min(values=tensor([-1.1899, -1.4644,  0.0384, -0.1153]), indices=tensor([2, 0, 1, 0]))
//
// .. function:: min(input, other, *, out=None) -> Tensor
//    :noindex:
//
// See :func:`torch.minimum`.
//
//
//go:linkname Min py.min
func Min(input *py.Object) *py.Object
//
// minimum(input, other, *, out=None) -> Tensor
//
// Computes the element-wise minimum of :attr:`input` and :attr:`other`.
//
// .. note::
//     If one of the elements being compared is a NaN, then that element is returned.
//     :func:`minimum` is not supported for tensors with complex dtypes.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor): the second input tensor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor((1, 2, -1))
//     >>> b = torch.tensor((3, 0, 4))
//     >>> torch.minimum(a, b)
//     tensor([1, 0, -1])
//
//
//go:linkname Minimum py.minimum
func Minimum(input *py.Object, other *py.Object) *py.Object
// None
//
//go:linkname MiopenBatchNorm py.miopen_batch_norm
func MiopenBatchNorm(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MiopenConvolution py.miopen_convolution
func MiopenConvolution(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MiopenConvolutionAddRelu py.miopen_convolution_add_relu
func MiopenConvolutionAddRelu(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MiopenConvolutionRelu py.miopen_convolution_relu
func MiopenConvolutionRelu(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MiopenConvolutionTranspose py.miopen_convolution_transpose
func MiopenConvolutionTranspose(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MiopenDepthwiseConvolution py.miopen_depthwise_convolution
func MiopenDepthwiseConvolution(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MiopenRnn py.miopen_rnn
func MiopenRnn(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MkldnnAdaptiveAvgPool2d py.mkldnn_adaptive_avg_pool2d
func MkldnnAdaptiveAvgPool2d(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MkldnnConvolution py.mkldnn_convolution
func MkldnnConvolution(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MkldnnLinearBackwardWeights py.mkldnn_linear_backward_weights
func MkldnnLinearBackwardWeights(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MkldnnMaxPool2d py.mkldnn_max_pool2d
func MkldnnMaxPool2d(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MkldnnMaxPool3d py.mkldnn_max_pool3d
func MkldnnMaxPool3d(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname MkldnnRnnLayer py.mkldnn_rnn_layer
func MkldnnRnnLayer(__llgo_va_list ...interface{}) *py.Object
//
// mm(input, mat2, *, out=None) -> Tensor
//
// Performs a matrix multiplication of the matrices :attr:`input` and :attr:`mat2`.
//
// If :attr:`input` is a :math:`(n \times m)` tensor, :attr:`mat2` is a
// :math:`(m \times p)` tensor, :attr:`out` will be a :math:`(n \times p)` tensor.
//
// .. note:: This function does not :ref:`broadcast <broadcasting-semantics>`.
//           For broadcasting matrix products, see :func:`torch.matmul`.
//
// Supports strided and sparse 2-D tensors as inputs, autograd with
// respect to strided inputs.
//
// This operation has support for arguments with :ref:`sparse layouts<sparse-docs>`.
// If :attr:`out` is provided it's layout will be used. Otherwise, the result
// layout will be deduced from that of :attr:`input`.
//
//
// .. warning::
//     Sparse support is a beta feature and some layout(s)/dtype/device combinations may not be supported,
//     or may not have autograd support. If you notice missing functionality please
//     open a feature request.
//
// This operator supports :ref:`TensorFloat32<tf32_on_ampere>`.
//
// On certain ROCm devices, when using float16 inputs this module will use :ref:`different precision<fp16_on_mi200>` for backward.
//
// Args:
//     input (Tensor): the first matrix to be matrix multiplied
//     mat2 (Tensor): the second matrix to be matrix multiplied
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> mat1 = torch.randn(2, 3)
//     >>> mat2 = torch.randn(3, 3)
//     >>> torch.mm(mat1, mat2)
//     tensor([[ 0.4851,  0.5037, -0.3633],
//             [-0.0760, -3.6705,  2.4784]])
//
//
//go:linkname Mm py.mm
func Mm(input *py.Object, mat2 *py.Object) *py.Object
//
// mode(input, dim=-1, keepdim=False, *, out=None) -> (Tensor, LongTensor)
//
// Returns a namedtuple ``(values, indices)`` where ``values`` is the mode
// value of each row of the :attr:`input` tensor in the given dimension
// :attr:`dim`, i.e. a value which appears most often
// in that row, and ``indices`` is the index location of each mode value found.
//
// By default, :attr:`dim` is the last dimension of the :attr:`input` tensor.
//
// If :attr:`keepdim` is ``True``, the output tensors are of the same size as
// :attr:`input` except in the dimension :attr:`dim` where they are of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting
// in the output tensors having 1 fewer dimension than :attr:`input`.
//
// .. note:: This function is not defined for ``torch.cuda.Tensor`` yet.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int): the dimension to reduce.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//     out (tuple, optional): the result tuple of two output tensors (values, indices)
//
// Example::
//
//     >>> a = torch.randint(10, (5,))
//     >>> a
//     tensor([6, 5, 1, 0, 2])
//     >>> b = a + (torch.randn(50, 1) * 5).long()
//     >>> torch.mode(b, 0)
//     torch.return_types.mode(values=tensor([6, 5, 1, 0, 2]), indices=tensor([2, 2, 2, 2, 2]))
//
//
//go:linkname Mode py.mode
func Mode(input *py.Object, dim *py.Object, keepdim *py.Object) *py.Object
//
// moveaxis(input, source, destination) -> Tensor
//
// Alias for :func:`torch.movedim`.
//
// This function is equivalent to NumPy's moveaxis function.
//
// Examples::
//
//     >>> t = torch.randn(3,2,1)
//     >>> t
//     tensor([[[-0.3362],
//             [-0.8437]],
//
//             [[-0.9627],
//             [ 0.1727]],
//
//             [[ 0.5173],
//             [-0.1398]]])
//     >>> torch.moveaxis(t, 1, 0).shape
//     torch.Size([2, 3, 1])
//     >>> torch.moveaxis(t, 1, 0)
//     tensor([[[-0.3362],
//             [-0.9627],
//             [ 0.5173]],
//
//             [[-0.8437],
//             [ 0.1727],
//             [-0.1398]]])
//     >>> torch.moveaxis(t, (1, 2), (0, 1)).shape
//     torch.Size([2, 1, 3])
//     >>> torch.moveaxis(t, (1, 2), (0, 1))
//     tensor([[[-0.3362, -0.9627,  0.5173]],
//
//             [[-0.8437,  0.1727, -0.1398]]])
//
//
//go:linkname Moveaxis py.moveaxis
func Moveaxis(input *py.Object, source *py.Object, destination *py.Object) *py.Object
//
// movedim(input, source, destination) -> Tensor
//
// Moves the dimension(s) of :attr:`input` at the position(s) in :attr:`source`
// to the position(s) in :attr:`destination`.
//
// Other dimensions of :attr:`input` that are not explicitly moved remain in
// their original order and appear at the positions not specified in :attr:`destination`.
//
// Args:
//     input (Tensor): the input tensor.
//     source (int or tuple of ints): Original positions of the dims to move. These must be unique.
//     destination (int or tuple of ints): Destination positions for each of the original dims. These must also be unique.
//
// Examples::
//
//     >>> t = torch.randn(3,2,1)
//     >>> t
//     tensor([[[-0.3362],
//             [-0.8437]],
//
//             [[-0.9627],
//             [ 0.1727]],
//
//             [[ 0.5173],
//             [-0.1398]]])
//     >>> torch.movedim(t, 1, 0).shape
//     torch.Size([2, 3, 1])
//     >>> torch.movedim(t, 1, 0)
//     tensor([[[-0.3362],
//             [-0.9627],
//             [ 0.5173]],
//
//             [[-0.8437],
//             [ 0.1727],
//             [-0.1398]]])
//     >>> torch.movedim(t, (1, 2), (0, 1)).shape
//     torch.Size([2, 1, 3])
//     >>> torch.movedim(t, (1, 2), (0, 1))
//     tensor([[[-0.3362, -0.9627,  0.5173]],
//
//             [[-0.8437,  0.1727, -0.1398]]])
//
//
//go:linkname Movedim py.movedim
func Movedim(input *py.Object, source *py.Object, destination *py.Object) *py.Object
//
// msort(input, *, out=None) -> Tensor
//
// Sorts the elements of the :attr:`input` tensor along its first dimension
// in ascending order by value.
//
// .. note:: `torch.msort(t)` is equivalent to `torch.sort(t, dim=0)[0]`.
//           See also :func:`torch.sort`.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> t = torch.randn(3, 4)
//     >>> t
//     tensor([[-0.1321,  0.4370, -1.2631, -1.1289],
//             [-2.0527, -1.1250,  0.2275,  0.3077],
//             [-0.0881, -0.1259, -0.5495,  1.0284]])
//     >>> torch.msort(t)
//     tensor([[-2.0527, -1.1250, -1.2631, -1.1289],
//             [-0.1321, -0.1259, -0.5495,  0.3077],
//             [-0.0881,  0.4370,  0.2275,  1.0284]])
//
//
//go:linkname Msort py.msort
func Msort(input *py.Object) *py.Object
//
// mul(input, other, *, out=None) -> Tensor
//
// Multiplies :attr:`input` by :attr:`other`.
//
//
// .. math::
//     \text{out}_i = \text{input}_i \times \text{other}_i
//
//
// Supports :ref:`broadcasting to a common shape <broadcasting-semantics>`,
// :ref:`type promotion <type-promotion-doc>`, and integer, float, and complex inputs.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor or Number) - the tensor or number to multiply input by.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Examples::
//
//     >>> a = torch.randn(3)
//     >>> a
//     tensor([ 0.2015, -0.4255,  2.6087])
//     >>> torch.mul(a, 100)
//     tensor([  20.1494,  -42.5491,  260.8663])
//
//     >>> b = torch.randn(4, 1)
//     >>> b
//     tensor([[ 1.1207],
//             [-0.3137],
//             [ 0.0700],
//             [ 0.8378]])
//     >>> c = torch.randn(1, 4)
//     >>> c
//     tensor([[ 0.5146,  0.1216, -0.5244,  2.2382]])
//     >>> torch.mul(b, c)
//     tensor([[ 0.5767,  0.1363, -0.5877,  2.5083],
//             [-0.1614, -0.0382,  0.1645, -0.7021],
//             [ 0.0360,  0.0085, -0.0367,  0.1567],
//             [ 0.4312,  0.1019, -0.4394,  1.8753]])
//
//
//go:linkname Mul py.mul
func Mul(input *py.Object, other *py.Object) *py.Object
//
// multinomial(input, num_samples, replacement=False, *, generator=None, out=None) -> LongTensor
//
// Returns a tensor where each row contains :attr:`num_samples` indices sampled
// from the multinomial (a stricter definition would be multivariate,
// refer to torch.distributions.multinomial.Multinomial for more details)
// probability distribution located in the corresponding row
// of tensor :attr:`input`.
//
// .. note::
//     The rows of :attr:`input` do not need to sum to one (in which case we use
//     the values as weights), but must be non-negative, finite and have
//     a non-zero sum.
//
// Indices are ordered from left to right according to when each was sampled
// (first samples are placed in first column).
//
// If :attr:`input` is a vector, :attr:`out` is a vector of size :attr:`num_samples`.
//
// If :attr:`input` is a matrix with `m` rows, :attr:`out` is an matrix of shape
// :math:`(m \times \text{num\_samples})`.
//
// If replacement is ``True``, samples are drawn with replacement.
//
// If not, they are drawn without replacement, which means that when a
// sample index is drawn for a row, it cannot be drawn again for that row.
//
// .. note::
//     When drawn without replacement, :attr:`num_samples` must be lower than
//     number of non-zero elements in :attr:`input` (or the min number of non-zero
//     elements in each row of :attr:`input` if it is a matrix).
//
// Args:
//     input (Tensor): the input tensor containing probabilities
//     num_samples (int): number of samples to draw
//     replacement (bool, optional): whether to draw with replacement or not
//
// Keyword args:
//     generator (:class:`torch.Generator`, optional): a pseudorandom number generator for sampling
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> weights = torch.tensor([0, 10, 3, 0], dtype=torch.float) # create a tensor of weights
//     >>> torch.multinomial(weights, 2)
//     tensor([1, 2])
//     >>> torch.multinomial(weights, 4) # ERROR!
//     RuntimeError: invalid argument 2: invalid multinomial distribution (with replacement=False,
//     not enough non-negative category to sample) at ../aten/src/TH/generic/THTensorRandom.cpp:320
//     >>> torch.multinomial(weights, 4, replacement=True)
//     tensor([ 2,  1,  1,  1])
//
//
//go:linkname Multinomial py.multinomial
func Multinomial(input *py.Object, numSamples *py.Object, replacement *py.Object) *py.Object
//
// multiply(input, other, *, out=None)
//
// Alias for :func:`torch.mul`.
//
//
//go:linkname Multiply py.multiply
func Multiply(input *py.Object, other *py.Object) *py.Object
//
// mv(input, vec, *, out=None) -> Tensor
//
// Performs a matrix-vector product of the matrix :attr:`input` and the vector
// :attr:`vec`.
//
// If :attr:`input` is a :math:`(n \times m)` tensor, :attr:`vec` is a 1-D tensor of
// size :math:`m`, :attr:`out` will be 1-D of size :math:`n`.
//
// .. note:: This function does not :ref:`broadcast <broadcasting-semantics>`.
//
// Args:
//     input (Tensor): matrix to be multiplied
//     vec (Tensor): vector to be multiplied
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> mat = torch.randn(2, 3)
//     >>> vec = torch.randn(3)
//     >>> torch.mv(mat, vec)
//     tensor([ 1.0404, -0.6361])
//
//
//go:linkname Mv py.mv
func Mv(input *py.Object, vec *py.Object) *py.Object
//
// mvlgamma(input, p, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.multigammaln`.
//
//
//go:linkname Mvlgamma py.mvlgamma
func Mvlgamma(input *py.Object, p *py.Object) *py.Object
//
// nan_to_num(input, nan=0.0, posinf=None, neginf=None, *, out=None) -> Tensor
//
// Replaces :literal:`NaN`, positive infinity, and negative infinity values in :attr:`input`
// with the values specified by :attr:`nan`, :attr:`posinf`, and :attr:`neginf`, respectively.
// By default, :literal:`NaN`\ s are replaced with zero, positive infinity is replaced with the
// greatest finite value representable by :attr:`input`'s dtype, and negative infinity
// is replaced with the least finite value representable by :attr:`input`'s dtype.
//
// Args:
//     input (Tensor): the input tensor.
//     nan (Number, optional): the value to replace :literal:`NaN`\s with. Default is zero.
//     posinf (Number, optional): if a Number, the value to replace positive infinity values with.
//         If None, positive infinity values are replaced with the greatest finite value representable by :attr:`input`'s dtype.
//         Default is None.
//     neginf (Number, optional): if a Number, the value to replace negative infinity values with.
//         If None, negative infinity values are replaced with the lowest finite value representable by :attr:`input`'s dtype.
//         Default is None.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> x = torch.tensor([float('nan'), float('inf'), -float('inf'), 3.14])
//     >>> torch.nan_to_num(x)
//     tensor([ 0.0000e+00,  3.4028e+38, -3.4028e+38,  3.1400e+00])
//     >>> torch.nan_to_num(x, nan=2.0)
//     tensor([ 2.0000e+00,  3.4028e+38, -3.4028e+38,  3.1400e+00])
//     >>> torch.nan_to_num(x, nan=2.0, posinf=1.0)
//     tensor([ 2.0000e+00,  1.0000e+00, -3.4028e+38,  3.1400e+00])
//
//
//
//go:linkname NanToNum py.nan_to_num
func NanToNum(input *py.Object, nan *py.Object, posinf *py.Object, neginf *py.Object) *py.Object
// None
//
//go:linkname NanToNum_ py.nan_to_num_
func NanToNum_(__llgo_va_list ...interface{}) *py.Object
//
// nanmean(input, dim=None, keepdim=False, *, dtype=None, out=None) -> Tensor
//
// Computes the mean of all `non-NaN` elements along the specified dimensions.
//
// This function is identical to :func:`torch.mean` when there are no `NaN` values
// in the :attr:`input` tensor. In the presence of `NaN`, :func:`torch.mean` will
// propagate the `NaN` to the output whereas :func:`torch.nanmean` will ignore the
// `NaN` values (`torch.nanmean(a)` is equivalent to `torch.mean(a[~a.isnan()])`).
//
//
// If :attr:`keepdim` is ``True``, the output tensor is of the same size
// as :attr:`input` except in the dimension(s) :attr:`dim` where it is of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in the
// output tensor having 1 (or ``len(dim)``) fewer dimension(s).
//
//
// Args:
//     input (Tensor): the input tensor.
//
//     dim (int or tuple of ints, optional): the dimension or dimensions to reduce.
//         If ``None``, all dimensions are reduced.
//
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         If specified, the input tensor is casted to :attr:`dtype` before the operation
//         is performed. This is useful for preventing data type overflows. Default: None.
//     out (Tensor, optional): the output tensor.
//
// .. seealso::
//
//     :func:`torch.mean` computes the mean value, propagating `NaN`.
//
// Example::
//
//     >>> x = torch.tensor([[torch.nan, 1, 2], [1, 2, 3]])
//     >>> x.mean()
//     tensor(nan)
//     >>> x.nanmean()
//     tensor(1.8000)
//     >>> x.mean(dim=0)
//     tensor([   nan, 1.5000, 2.5000])
//     >>> x.nanmean(dim=0)
//     tensor([1.0000, 1.5000, 2.5000])
//
//     # If all elements in the reduced dimensions are NaN then the result is NaN
//     >>> torch.tensor([torch.nan]).nanmean()
//     tensor(nan)
//
//
//go:linkname Nanmean py.nanmean
func Nanmean(input *py.Object, dim *py.Object, keepdim *py.Object) *py.Object
//
// nanmedian(input) -> Tensor
//
// Returns the median of the values in :attr:`input`, ignoring ``NaN`` values.
//
// This function is identical to :func:`torch.median` when there are no ``NaN`` values in :attr:`input`.
// When :attr:`input` has one or more ``NaN`` values, :func:`torch.median` will always return ``NaN``,
// while this function will return the median of the non-``NaN`` elements in :attr:`input`.
// If all the elements in :attr:`input` are ``NaN`` it will also return ``NaN``.
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> a = torch.tensor([1, float('nan'), 3, 2])
//     >>> a.median()
//     tensor(nan)
//     >>> a.nanmedian()
//     tensor(2.)
//
// .. function:: nanmedian(input, dim=-1, keepdim=False, *, out=None) -> (Tensor, LongTensor)
//    :noindex:
//
// Returns a namedtuple ``(values, indices)`` where ``values`` contains the median of each row of :attr:`input`
// in the dimension :attr:`dim`, ignoring ``NaN`` values, and ``indices`` contains the index of the median values
// found in the dimension :attr:`dim`.
//
// This function is identical to :func:`torch.median` when there are no ``NaN`` values in a reduced row. When a reduced row has
// one or more ``NaN`` values, :func:`torch.median` will always reduce it to ``NaN``, while this function will reduce it to the
// median of the non-``NaN`` elements. If all the elements in a reduced row are ``NaN`` then it will be reduced to ``NaN``, too.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int): the dimension to reduce.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//     out ((Tensor, Tensor), optional): The first tensor will be populated with the median values and the second
//                                       tensor, which must have dtype long, with their indices in the dimension
//                                       :attr:`dim` of :attr:`input`.
//
// Example::
//
//     >>> a = torch.tensor([[2, 3, 1], [float('nan'), 1, float('nan')]])
//     >>> a
//     tensor([[2., 3., 1.],
//             [nan, 1., nan]])
//     >>> a.median(0)
//     torch.return_types.median(values=tensor([nan, 1., nan]), indices=tensor([1, 1, 1]))
//     >>> a.nanmedian(0)
//     torch.return_types.nanmedian(values=tensor([2., 1., 1.]), indices=tensor([0, 1, 0]))
//
//
//go:linkname Nanmedian py.nanmedian
func Nanmedian(input *py.Object) *py.Object
//
// nanquantile(input, q, dim=None, keepdim=False, *, interpolation='linear', out=None) -> Tensor
//
// This is a variant of :func:`torch.quantile` that "ignores" ``NaN`` values,
// computing the quantiles :attr:`q` as if ``NaN`` values in :attr:`input` did
// not exist. If all values in a reduced row are ``NaN`` then the quantiles for
// that reduction will be ``NaN``. See the documentation for :func:`torch.quantile`.
//
// Args:
//     input (Tensor): the input tensor.
//     q (float or Tensor): a scalar or 1D tensor of quantile values in the range [0, 1]
//     dim (int): the dimension to reduce.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword arguments:
//     interpolation (str): interpolation method to use when the desired quantile lies between two data points.
//                             Can be ``linear``, ``lower``, ``higher``, ``midpoint`` and ``nearest``.
//                             Default is ``linear``.
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> t = torch.tensor([float('nan'), 1, 2])
//     >>> t.quantile(0.5)
//     tensor(nan)
//     >>> t.nanquantile(0.5)
//     tensor(1.5000)
//     >>> t = torch.tensor([[float('nan'), float('nan')], [1, 2]])
//     >>> t
//     tensor([[nan, nan],
//             [1., 2.]])
//     >>> t.nanquantile(0.5, dim=0)
//     tensor([1., 2.])
//     >>> t.nanquantile(0.5, dim=1)
//     tensor([   nan, 1.5000])
//
//
//go:linkname Nanquantile py.nanquantile
func Nanquantile(input *py.Object, q *py.Object, dim *py.Object, keepdim *py.Object) *py.Object
//
// nansum(input, *, dtype=None) -> Tensor
//
// Returns the sum of all elements, treating Not a Numbers (NaNs) as zero.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         If specified, the input tensor is casted to :attr:`dtype` before the operation
//         is performed. This is useful for preventing data type overflows. Default: None.
//
// Example::
//
//     >>> a = torch.tensor([1., 2., float('nan'), 4.])
//     >>> torch.nansum(a)
//     tensor(7.)
//
// .. function:: nansum(input, dim, keepdim=False, *, dtype=None) -> Tensor
//    :noindex:
//
// Returns the sum of each row of the :attr:`input` tensor in the given
// dimension :attr:`dim`, treating Not a Numbers (NaNs) as zero.
// If :attr:`dim` is a list of dimensions, reduce over all of them.
//
//
// If :attr:`keepdim` is ``True``, the output tensor is of the same size
// as :attr:`input` except in the dimension(s) :attr:`dim` where it is of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in the
// output tensor having 1 (or ``len(dim)``) fewer dimension(s).
//
//
// Args:
//     input (Tensor): the input tensor.
//
//     dim (int or tuple of ints, optional): the dimension or dimensions to reduce.
//         If ``None``, all dimensions are reduced.
//
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         If specified, the input tensor is casted to :attr:`dtype` before the operation
//         is performed. This is useful for preventing data type overflows. Default: None.
//
// Example::
//
//     >>> torch.nansum(torch.tensor([1., float("nan")]))
//     1.0
//     >>> a = torch.tensor([[1, 2], [3., float("nan")]])
//     >>> torch.nansum(a)
//     tensor(6.)
//     >>> torch.nansum(a, dim=0)
//     tensor([4., 2.])
//     >>> torch.nansum(a, dim=1)
//     tensor([3., 3.])
//
//
//go:linkname Nansum py.nansum
func Nansum(input *py.Object) *py.Object
//
// narrow(input, dim, start, length) -> Tensor
//
// Returns a new tensor that is a narrowed version of :attr:`input` tensor. The
// dimension :attr:`dim` is input from :attr:`start` to ``start + length``. The
// returned tensor and :attr:`input` tensor share the same underlying storage.
//
// Args:
//     input (Tensor): the tensor to narrow
//     dim (int): the dimension along which to narrow
//     start (int or Tensor): index of the element to start the narrowed dimension
//         from. Can be negative, which means indexing from the end of `dim`. If
//         `Tensor`, it must be an 0-dim integral `Tensor` (bools not allowed)
//     length (int): length of the narrowed dimension, must be weakly positive
//
// Example::
//
//     >>> x = torch.tensor([[1, 2, 3], [4, 5, 6], [7, 8, 9]])
//     >>> torch.narrow(x, 0, 0, 2)
//     tensor([[ 1,  2,  3],
//             [ 4,  5,  6]])
//     >>> torch.narrow(x, 1, 1, 2)
//     tensor([[ 2,  3],
//             [ 5,  6],
//             [ 8,  9]])
//     >>> torch.narrow(x, -1, torch.tensor(-1), 1)
//     tensor([[3],
//             [6],
//             [9]])
//
//
//go:linkname Narrow py.narrow
func Narrow(input *py.Object, dim *py.Object, start *py.Object, length *py.Object) *py.Object
//
// narrow_copy(input, dim, start, length, *, out=None) -> Tensor
//
// Same as :meth:`Tensor.narrow` except this returns a copy rather
// than shared storage. This is primarily for sparse tensors, which
// do not have a shared-storage narrow method.
//
// Args:
//     input (Tensor): the tensor to narrow
//     dim (int): the dimension along which to narrow
//     start (int): index of the element to start the narrowed dimension from. Can
//         be negative, which means indexing from the end of `dim`
//     length (int): length of the narrowed dimension, must be weakly positive
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> x = torch.tensor([[1, 2, 3], [4, 5, 6], [7, 8, 9]])
//     >>> torch.narrow_copy(x, 0, 0, 2)
//     tensor([[ 1,  2,  3],
//             [ 4,  5,  6]])
//     >>> torch.narrow_copy(x, 1, 1, 2)
//     tensor([[ 2,  3],
//             [ 5,  6],
//             [ 8,  9]])
//     >>> s = torch.arange(16).reshape(2, 2, 2, 2).to_sparse(2)
//     >>> torch.narrow_copy(s, 0, 0, 1)
//     tensor(indices=tensor([[0, 0],
//                            [0, 1]]),
//            values=tensor([[[0, 1],
//                            [2, 3]],
//
//                           [[4, 5],
//                            [6, 7]]]),
//            size=(1, 2, 2, 2), nnz=2, layout=torch.sparse_coo)
//
// .. seealso::
//
//         :func:`torch.narrow` for a non copy variant
//
//
//
//go:linkname NarrowCopy py.narrow_copy
func NarrowCopy(input *py.Object, dim *py.Object, start *py.Object, length *py.Object) *py.Object
// None
//
//go:linkname NativeBatchNorm py.native_batch_norm
func NativeBatchNorm(__llgo_va_list ...interface{}) *py.Object
//
// native_channel_shuffle(input, groups) -> Tensor
//
// Native kernel level implementation of the `channel_shuffle`.
// This function might become private in future releases, use with caution.
//
// Divide the channels in a tensor of shape :math:`(*, C , H, W)`
// into g groups and rearrange them as :math:`(*, C \frac g, g, H, W)`,
// while keeping the original tensor shape.
//
// See :class:`~torch.nn.ChannelShuffle` for details.
//
// Args:
//     input (Tensor): the input tensor
//     groups (int): number of groups to divide channels in and rearrange.
//
// Examples::
//
//     >>> input = torch.randn(1, 4, 2, 2)
//     >>> print(input)
//     [[[[1, 2],
//        [3, 4]],
//       [[5, 6],
//        [7, 8]],
//       [[9, 10],
//        [11, 12]],
//       [[13, 14],
//        [15, 16]],
//      ]]
//     >>> output = torch.nn.functional.native_channel_shuffle(input, 2)
//     >>> print(output)
//     [[[[1, 2],
//        [3, 4]],
//       [[9, 10],
//        [11, 12]],
//       [[5, 6],
//        [7, 8]],
//       [[13, 14],
//        [15, 16]],
//      ]]
//
//
//go:linkname NativeChannelShuffle py.native_channel_shuffle
func NativeChannelShuffle(input *py.Object, groups *py.Object) *py.Object
// None
//
//go:linkname NativeDropout py.native_dropout
func NativeDropout(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname NativeGroupNorm py.native_group_norm
func NativeGroupNorm(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname NativeLayerNorm py.native_layer_norm
func NativeLayerNorm(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname NativeNorm py.native_norm
func NativeNorm(__llgo_va_list ...interface{}) *py.Object
//
// ne(input, other, *, out=None) -> Tensor
//
// Computes :math:`\text{input} \neq \text{other}` element-wise.
//
//
// The second argument can be a number or a tensor whose shape is
// :ref:`broadcastable <broadcasting-semantics>` with the first argument.
//
// Args:
//     input (Tensor): the tensor to compare
//     other (Tensor or float): the tensor or value to compare
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Returns:
//     A boolean tensor that is True where :attr:`input` is not equal to :attr:`other` and False elsewhere
//
// Example::
//
//     >>> torch.ne(torch.tensor([[1, 2], [3, 4]]), torch.tensor([[1, 1], [4, 4]]))
//     tensor([[False, True], [True, False]])
//
//
//go:linkname Ne py.ne
func Ne(input *py.Object, other *py.Object) *py.Object
//
// neg(input, *, out=None) -> Tensor
//
// Returns a new tensor with the negative of the elements of :attr:`input`.
//
// .. math::
//     \text{out} = -1 \times \text{input}
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(5)
//     >>> a
//     tensor([ 0.0090, -0.2262, -0.0682, -0.2866,  0.3940])
//     >>> torch.neg(a)
//     tensor([-0.0090,  0.2262,  0.0682,  0.2866, -0.3940])
//
//
//go:linkname Neg py.neg
func Neg(input *py.Object) *py.Object
// None
//
//go:linkname Neg_ py.neg_
func Neg_(__llgo_va_list ...interface{}) *py.Object
//
// negative(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.neg`
//
//
//go:linkname Negative py.negative
func Negative(input *py.Object) *py.Object
// None
//
//go:linkname Negative_ py.negative_
func Negative_(__llgo_va_list ...interface{}) *py.Object
//
// nextafter(input, other, *, out=None) -> Tensor
//
// Return the next floating-point value after :attr:`input` towards :attr:`other`, elementwise.
//
// The shapes of ``input`` and ``other`` must be
// :ref:`broadcastable <broadcasting-semantics>`.
//
// Args:
//     input (Tensor): the first input tensor
//     other (Tensor): the second input tensor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> eps = torch.finfo(torch.float32).eps
//     >>> torch.nextafter(torch.tensor([1.0, 2.0]), torch.tensor([2.0, 1.0])) == torch.tensor([eps + 1, 2 - eps])
//     tensor([True, True])
//
//
//
//go:linkname Nextafter py.nextafter
func Nextafter(input *py.Object, other *py.Object) *py.Object
//
// nonzero(input, *, out=None, as_tuple=False) -> LongTensor or tuple of LongTensors
//
// .. note::
//     :func:`torch.nonzero(..., as_tuple=False) <torch.nonzero>` (default) returns a
//     2-D tensor where each row is the index for a nonzero value.
//
//     :func:`torch.nonzero(..., as_tuple=True) <torch.nonzero>` returns a tuple of 1-D
//     index tensors, allowing for advanced indexing, so ``x[x.nonzero(as_tuple=True)]``
//     gives all nonzero values of tensor ``x``. Of the returned tuple, each index tensor
//     contains nonzero indices for a certain dimension.
//
//     See below for more details on the two behaviors.
//
//     When :attr:`input` is on CUDA, :func:`torch.nonzero() <torch.nonzero>` causes
//     host-device synchronization.
//
// **When** :attr:`as_tuple` **is** ``False`` **(default)**:
//
// Returns a tensor containing the indices of all non-zero elements of
// :attr:`input`.  Each row in the result contains the indices of a non-zero
// element in :attr:`input`. The result is sorted lexicographically, with
// the last index changing the fastest (C-style).
//
// If :attr:`input` has :math:`n` dimensions, then the resulting indices tensor
// :attr:`out` is of size :math:`(z \times n)`, where :math:`z` is the total number of
// non-zero elements in the :attr:`input` tensor.
//
// **When** :attr:`as_tuple` **is** ``True``:
//
// Returns a tuple of 1-D tensors, one for each dimension in :attr:`input`,
// each containing the indices (in that dimension) of all non-zero elements of
// :attr:`input` .
//
// If :attr:`input` has :math:`n` dimensions, then the resulting tuple contains :math:`n`
// tensors of size :math:`z`, where :math:`z` is the total number of
// non-zero elements in the :attr:`input` tensor.
//
// As a special case, when :attr:`input` has zero dimensions and a nonzero scalar
// value, it is treated as a one-dimensional tensor with one element.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (LongTensor, optional): the output tensor containing indices
//
// Returns:
//     LongTensor or tuple of LongTensor: If :attr:`as_tuple` is ``False``, the output
//     tensor containing indices. If :attr:`as_tuple` is ``True``, one 1-D tensor for
//     each dimension, containing the indices of each nonzero element along that
//     dimension.
//
// Example::
//
//     >>> torch.nonzero(torch.tensor([1, 1, 1, 0, 1]))
//     tensor([[ 0],
//             [ 1],
//             [ 2],
//             [ 4]])
//     >>> torch.nonzero(torch.tensor([[0.6, 0.0, 0.0, 0.0],
//     ...                             [0.0, 0.4, 0.0, 0.0],
//     ...                             [0.0, 0.0, 1.2, 0.0],
//     ...                             [0.0, 0.0, 0.0,-0.4]]))
//     tensor([[ 0,  0],
//             [ 1,  1],
//             [ 2,  2],
//             [ 3,  3]])
//     >>> torch.nonzero(torch.tensor([1, 1, 1, 0, 1]), as_tuple=True)
//     (tensor([0, 1, 2, 4]),)
//     >>> torch.nonzero(torch.tensor([[0.6, 0.0, 0.0, 0.0],
//     ...                             [0.0, 0.4, 0.0, 0.0],
//     ...                             [0.0, 0.0, 1.2, 0.0],
//     ...                             [0.0, 0.0, 0.0,-0.4]]), as_tuple=True)
//     (tensor([0, 1, 2, 3]), tensor([0, 1, 2, 3]))
//     >>> torch.nonzero(torch.tensor(5), as_tuple=True)
//     (tensor([0]),)
//
//
//go:linkname Nonzero py.nonzero
func Nonzero(input *py.Object) *py.Object
// None
//
//go:linkname NonzeroStatic py.nonzero_static
func NonzeroStatic(__llgo_va_list ...interface{}) *py.Object
// Returns the matrix norm or vector norm of a given tensor.
//
//     .. warning::
//
//         torch.norm is deprecated and may be removed in a future PyTorch release.
//         Its documentation and behavior may be incorrect, and it is no longer
//         actively maintained.
//
//         Use :func:`torch.linalg.vector_norm` when computing vector norms and
//         :func:`torch.linalg.matrix_norm` when computing matrix norms.
//         For a function with a similar behavior as this one see :func:`torch.linalg.norm`.
//         Note, however, the signature for these functions is slightly different than the
//         signature for ``torch.norm``.
//
//     Args:
//         input (Tensor): The input tensor. Its data type must be either a floating
//             point or complex type. For complex inputs, the norm is calculated using the
//             absolute value of each element. If the input is complex and neither
//             :attr:`dtype` nor :attr:`out` is specified, the result's data type will
//             be the corresponding floating point type (e.g. float if :attr:`input` is
//             complexfloat).
//
//         p (int, float, inf, -inf, 'fro', 'nuc', optional): the order of norm. Default: ``'fro'``
//             The following norms can be calculated:
//
//             ======  ==============  ==========================
//             ord     matrix norm     vector norm
//             ======  ==============  ==========================
//             'fro'   Frobenius norm  --
//             'nuc'   nuclear norm    --
//             Number  --              sum(abs(x)**ord)**(1./ord)
//             ======  ==============  ==========================
//
//             The vector norm can be calculated across any number of dimensions.
//             The corresponding dimensions of :attr:`input` are flattened into
//             one dimension, and the norm is calculated on the flattened
//             dimension.
//
//             Frobenius norm produces the same result as ``p=2`` in all cases
//             except when :attr:`dim` is a list of three or more dims, in which
//             case Frobenius norm throws an error.
//
//             Nuclear norm can only be calculated across exactly two dimensions.
//
//         dim (int, tuple of ints, list of ints, optional):
//             Specifies which dimension or dimensions of :attr:`input` to
//             calculate the norm across. If :attr:`dim` is ``None``, the norm will
//             be calculated across all dimensions of :attr:`input`. If the norm
//             type indicated by :attr:`p` does not support the specified number of
//             dimensions, an error will occur.
//         keepdim (bool, optional): whether the output tensors have :attr:`dim`
//             retained or not. Ignored if :attr:`dim` = ``None`` and
//             :attr:`out` = ``None``. Default: ``False``
//         out (Tensor, optional): the output tensor. Ignored if
//             :attr:`dim` = ``None`` and :attr:`out` = ``None``.
//         dtype (:class:`torch.dtype`, optional): the desired data type of
//             returned tensor. If specified, the input tensor is casted to
//             :attr:`dtype` while performing the operation. Default: None.
//
//     .. note::
//         Even though ``p='fro'`` supports any number of dimensions, the true
//         mathematical definition of Frobenius norm only applies to tensors with
//         exactly two dimensions. :func:`torch.linalg.matrix_norm` with ``ord='fro'``
//         aligns with the mathematical definition, since it can only be applied across
//         exactly two dimensions.
//
//     Example::
//
//         >>> import torch
//         >>> a = torch.arange(9, dtype= torch.float) - 4
//         >>> b = a.reshape((3, 3))
//         >>> torch.norm(a)
//         tensor(7.7460)
//         >>> torch.norm(b)
//         tensor(7.7460)
//         >>> torch.norm(a, float('inf'))
//         tensor(4.)
//         >>> torch.norm(b, float('inf'))
//         tensor(4.)
//         >>> c = torch.tensor([[ 1, 2, 3], [-1, 1, 4]] , dtype=torch.float)
//         >>> torch.norm(c, dim=0)
//         tensor([1.4142, 2.2361, 5.0000])
//         >>> torch.norm(c, dim=1)
//         tensor([3.7417, 4.2426])
//         >>> torch.norm(c, p=1, dim=1)
//         tensor([6., 6.])
//         >>> d = torch.arange(8, dtype=torch.float).reshape(2, 2, 2)
//         >>> torch.norm(d, dim=(1, 2))
//         tensor([ 3.7417, 11.2250])
//         >>> torch.norm(d[0, :, :]), torch.norm(d[1, :, :])
//         (tensor(3.7417), tensor(11.2250))
//
//
//go:linkname Norm py.norm
func Norm(input *py.Object, p *py.Object, dim *py.Object, keepdim *py.Object, out *py.Object, dtype *py.Object) *py.Object
// None
//
//go:linkname NormExceptDim py.norm_except_dim
func NormExceptDim(__llgo_va_list ...interface{}) *py.Object
//
// normal(mean, std, *, generator=None, out=None) -> Tensor
//
// Returns a tensor of random numbers drawn from separate normal distributions
// whose mean and standard deviation are given.
//
// The :attr:`mean` is a tensor with the mean of
// each output element's normal distribution
//
// The :attr:`std` is a tensor with the standard deviation of
// each output element's normal distribution
//
// The shapes of :attr:`mean` and :attr:`std` don't need to match, but the
// total number of elements in each tensor need to be the same.
//
// .. note:: When the shapes do not match, the shape of :attr:`mean`
//           is used as the shape for the returned output tensor
//
// .. note:: When :attr:`std` is a CUDA tensor, this function synchronizes
//           its device with the CPU.
//
// Args:
//     mean (Tensor): the tensor of per-element means
//     std (Tensor): the tensor of per-element standard deviations
//
// Keyword args:
//     generator (:class:`torch.Generator`, optional): a pseudorandom number generator for sampling
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.normal(mean=torch.arange(1., 11.), std=torch.arange(1, 0, -0.1))
//     tensor([  1.0425,   3.5672,   2.7969,   4.2925,   4.7229,   6.2134,
//               8.0505,   8.1408,   9.0563,  10.0566])
//
// .. function:: normal(mean=0.0, std, *, out=None) -> Tensor
//    :noindex:
//
// Similar to the function above, but the means are shared among all drawn
// elements.
//
// Args:
//     mean (float, optional): the mean for all distributions
//     std (Tensor): the tensor of per-element standard deviations
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.normal(mean=0.5, std=torch.arange(1., 6.))
//     tensor([-1.2793, -1.0732, -2.0687,  5.1177, -1.2303])
//
// .. function:: normal(mean, std=1.0, *, out=None) -> Tensor
//    :noindex:
//
// Similar to the function above, but the standard deviations are shared among
// all drawn elements.
//
// Args:
//     mean (Tensor): the tensor of per-element means
//     std (float, optional): the standard deviation for all distributions
//
// Keyword args:
//     out (Tensor, optional): the output tensor
//
// Example::
//
//     >>> torch.normal(mean=torch.arange(1., 6.))
//     tensor([ 1.1552,  2.6148,  2.6535,  5.8318,  4.2361])
//
// .. function:: normal(mean, std, size, *, out=None) -> Tensor
//    :noindex:
//
// Similar to the function above, but the means and standard deviations are shared
// among all drawn elements. The resulting tensor has size given by :attr:`size`.
//
// Args:
//     mean (float): the mean for all distributions
//     std (float): the standard deviation for all distributions
//     size (int...): a sequence of integers defining the shape of the output tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.normal(2, 3, size=(1, 4))
//     tensor([[-1.3987, -1.9544,  3.6048,  0.7909]])
//
//
//go:linkname Normal py.normal
func Normal(mean *py.Object, std *py.Object) *py.Object
//
// not_equal(input, other, *, out=None) -> Tensor
//
// Alias for :func:`torch.ne`.
//
//
//go:linkname NotEqual py.not_equal
func NotEqual(input *py.Object, other *py.Object) *py.Object
// None
//
//go:linkname NuclearNorm py.nuclear_norm
func NuclearNorm(__llgo_va_list ...interface{}) *py.Object
//
// numel(input) -> int
//
// Returns the total number of elements in the :attr:`input` tensor.
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> a = torch.randn(1, 2, 3, 4, 5)
//     >>> torch.numel(a)
//     120
//     >>> a = torch.zeros(4,4)
//     >>> torch.numel(a)
//     16
//
//
//
//go:linkname Numel py.numel
func Numel(input *py.Object) *py.Object
//
// ones(*size, *, out=None, dtype=None, layout=torch.strided, device=None, requires_grad=False) -> Tensor
//
// Returns a tensor filled with the scalar value `1`, with the shape defined
// by the variable argument :attr:`size`.
//
// Args:
//     size (int...): a sequence of integers defining the shape of the output tensor.
//         Can be a variable number of arguments or a collection like a list or tuple.
//
// Keyword arguments:
//     out (Tensor, optional): the output tensor.
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`).
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
// Example::
//
//     >>> torch.ones(2, 3)
//     tensor([[ 1.,  1.,  1.],
//             [ 1.,  1.,  1.]])
//
//     >>> torch.ones(5)
//     tensor([ 1.,  1.,  1.,  1.,  1.])
//
//
//
//go:linkname Ones py.ones
func Ones(__llgo_va_list ...interface{}) *py.Object
//
// ones_like(input, *, dtype=None, layout=None, device=None, requires_grad=False, memory_format=torch.preserve_format) -> Tensor
//
// Returns a tensor filled with the scalar value `1`, with the same size as
// :attr:`input`. ``torch.ones_like(input)`` is equivalent to
// ``torch.ones(input.size(), dtype=input.dtype, layout=input.layout, device=input.device)``.
//
// .. warning::
//     As of 0.4, this function does not support an :attr:`out` keyword. As an alternative,
//     the old ``torch.ones_like(input, out=output)`` is equivalent to
//     ``torch.ones(input.size(), out=output)``.
//
// Args:
//     input (Tensor): the size of :attr:`input` will determine size of the output tensor.
//
// Keyword arguments:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned Tensor.
//         Default: if ``None``, defaults to the dtype of :attr:`input`.
//     layout (:class:`torch.layout`, optional): the desired layout of returned tensor.
//         Default: if ``None``, defaults to the layout of :attr:`input`.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, defaults to the device of :attr:`input`.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     memory_format (:class:`torch.memory_format`, optional): the desired memory format of
//         returned Tensor. Default: ``torch.preserve_format``.
//
// Example::
//
//     >>> input = torch.empty(2, 3)
//     >>> torch.ones_like(input)
//     tensor([[ 1.,  1.,  1.],
//             [ 1.,  1.,  1.]])
//
//
//go:linkname OnesLike py.ones_like
func OnesLike(input *py.Object) *py.Object
//
// orgqr(input, tau) -> Tensor
//
// Alias for :func:`torch.linalg.householder_product`.
//
//
//go:linkname Orgqr py.orgqr
func Orgqr(input *py.Object, tau *py.Object) *py.Object
//
// ormqr(input, tau, other, left=True, transpose=False, *, out=None) -> Tensor
//
// Computes the matrix-matrix multiplication of a product of Householder matrices with a general matrix.
//
// Multiplies a :math:`m \times n` matrix `C` (given by :attr:`other`) with a matrix `Q`,
// where `Q` is represented using Householder reflectors `(input, tau)`.
// See `Representation of Orthogonal or Unitary Matrices`_ for further details.
//
// If :attr:`left` is `True` then `op(Q)` times `C` is computed, otherwise the result is `C` times `op(Q)`.
// When :attr:`left` is `True`, the implicit matrix `Q` has size :math:`m \times m`.
// It has size :math:`n \times n` otherwise.
// If :attr:`transpose` is `True` then `op` is the conjugate transpose operation, otherwise it's a no-op.
//
// Supports inputs of float, double, cfloat and cdouble dtypes.
// Also supports batched inputs, and, if the input is batched, the output is batched with the same dimensions.
//
// .. seealso::
//         :func:`torch.geqrf` can be used to form the Householder representation `(input, tau)` of matrix `Q`
//         from the QR decomposition.
//
// .. note::
//         This function supports backward but it is only fast when ``(input, tau)`` do not require gradients
//         and/or ``tau.size(-1)`` is very small.
//         ``
//
// Args:
//     input (Tensor): tensor of shape `(*, mn, k)` where `*` is zero or more batch dimensions
//                     and `mn` equals to `m` or `n` depending on the :attr:`left`.
//     tau (Tensor): tensor of shape `(*, min(mn, k))` where `*` is zero or more batch dimensions.
//     other (Tensor): tensor of shape `(*, m, n)` where `*` is zero or more batch dimensions.
//     left (bool): controls the order of multiplication.
//     transpose (bool): controls whether the matrix `Q` is conjugate transposed or not.
//
// Keyword args:
//     out (Tensor, optional): the output Tensor. Ignored if `None`. Default: `None`.
//
// .. _Representation of Orthogonal or Unitary Matrices:
//     https://www.netlib.org/lapack/lug/node128.html
//
//
//go:linkname Ormqr py.ormqr
func Ormqr(input *py.Object, tau *py.Object, other *py.Object, left *py.Object, transpose *py.Object) *py.Object
//
// outer(input, vec2, *, out=None) -> Tensor
//
// Outer product of :attr:`input` and :attr:`vec2`.
// If :attr:`input` is a vector of size :math:`n` and :attr:`vec2` is a vector of
// size :math:`m`, then :attr:`out` must be a matrix of size :math:`(n \times m)`.
//
// .. note:: This function does not :ref:`broadcast <broadcasting-semantics>`.
//
// Args:
//     input (Tensor): 1-D input vector
//     vec2 (Tensor): 1-D input vector
//
// Keyword args:
//     out (Tensor, optional): optional output matrix
//
// Example::
//
//     >>> v1 = torch.arange(1., 5.)
//     >>> v2 = torch.arange(1., 4.)
//     >>> torch.outer(v1, v2)
//     tensor([[  1.,   2.,   3.],
//             [  2.,   4.,   6.],
//             [  3.,   6.,   9.],
//             [  4.,   8.,  12.]])
//
//
//go:linkname Outer py.outer
func Outer(input *py.Object, vec2 *py.Object) *py.Object
//
// pairwise_distance(x1, x2, p=2.0, eps=1e-6, keepdim=False) -> Tensor
//
// See :class:`torch.nn.PairwiseDistance` for details
//
//
//go:linkname PairwiseDistance py.pairwise_distance
func PairwiseDistance(x1 *py.Object, x2 *py.Object, p *py.Object, eps *py.Object, keepdim *py.Object) *py.Object
//
// pdist(input, p=2) -> Tensor
//
// Computes the p-norm distance between every pair of row vectors in the input.
// This is identical to the upper triangular portion, excluding the diagonal, of
// `torch.norm(input[:, None] - input, dim=2, p=p)`. This function will be faster
// if the rows are contiguous.
//
// If input has shape :math:`N \times M` then the output will have shape
// :math:`\frac{1}{2} N (N - 1)`.
//
// This function is equivalent to ``scipy.spatial.distance.pdist(input,
// 'minkowski', p=p)`` if :math:`p \in (0, \infty)`. When :math:`p = 0` it is
// equivalent to ``scipy.spatial.distance.pdist(input, 'hamming') * M``.
// When :math:`p = \infty`, the closest scipy function is
// ``scipy.spatial.distance.pdist(xn, lambda x, y: np.abs(x - y).max())``.
//
// Args:
//     input: input tensor of shape :math:`N \times M`.
//     p: p value for the p-norm distance to calculate between each vector pair
//         :math:`\in [0, \infty]`.
//
//
//go:linkname Pdist py.pdist
func Pdist(input *py.Object, p *py.Object) *py.Object
//
// permute(input, dims) -> Tensor
//
// Returns a view of the original tensor :attr:`input` with its dimensions permuted.
//
// Args:
//     input (Tensor): the input tensor.
//     dims (tuple of int): The desired ordering of dimensions
//
// Example:
//     >>> x = torch.randn(2, 3, 5)
//     >>> x.size()
//     torch.Size([2, 3, 5])
//     >>> torch.permute(x, (2, 0, 1)).size()
//     torch.Size([5, 2, 3])
//
//
//go:linkname Permute py.permute
func Permute(input *py.Object, dims *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.permute`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname PermuteCopy py.permute_copy
func PermuteCopy(__llgo_va_list ...interface{}) *py.Object
//
// pinverse(input, rcond=1e-15) -> Tensor
//
// Alias for :func:`torch.linalg.pinv`
//
//
//go:linkname Pinverse py.pinverse
func Pinverse(input *py.Object, rcond *py.Object) *py.Object
//
// pixel_shuffle(input, upscale_factor) -> Tensor
//
// Rearranges elements in a tensor of shape :math:`(*, C \times r^2, H, W)` to a
// tensor of shape :math:`(*, C, H \times r, W \times r)`, where r is the :attr:`upscale_factor`.
//
// See :class:`~torch.nn.PixelShuffle` for details.
//
// Args:
//     input (Tensor): the input tensor
//     upscale_factor (int): factor to increase spatial resolution by
//
// Examples::
//
//     >>> input = torch.randn(1, 9, 4, 4)
//     >>> output = torch.nn.functional.pixel_shuffle(input, 3)
//     >>> print(output.size())
//     torch.Size([1, 1, 12, 12])
//
//
//go:linkname PixelShuffle py.pixel_shuffle
func PixelShuffle(input *py.Object, upscaleFactor *py.Object) *py.Object
//
// pixel_unshuffle(input, downscale_factor) -> Tensor
//
// Reverses the :class:`~torch.nn.PixelShuffle` operation by rearranging elements in a
// tensor of shape :math:`(*, C, H \times r, W \times r)` to a tensor of shape
// :math:`(*, C \times r^2, H, W)`, where r is the :attr:`downscale_factor`.
//
// See :class:`~torch.nn.PixelUnshuffle` for details.
//
// Args:
//     input (Tensor): the input tensor
//     downscale_factor (int): factor to increase spatial resolution by
//
// Examples::
//
//     >>> input = torch.randn(1, 1, 12, 12)
//     >>> output = torch.nn.functional.pixel_unshuffle(input, 3)
//     >>> print(output.size())
//     torch.Size([1, 9, 4, 4])
//
//
//go:linkname PixelUnshuffle py.pixel_unshuffle
func PixelUnshuffle(input *py.Object, downscaleFactor *py.Object) *py.Object
//
// poisson(input, generator=None) -> Tensor
//
// Returns a tensor of the same size as :attr:`input` with each element
// sampled from a Poisson distribution with rate parameter given by the corresponding
// element in :attr:`input` i.e.,
//
// .. math::
//     \text{out}_i \sim \text{Poisson}(\text{input}_i)
//
// :attr:`input` must be non-negative.
//
// Args:
//     input (Tensor): the input tensor containing the rates of the Poisson distribution
//
// Keyword args:
//     generator (:class:`torch.Generator`, optional): a pseudorandom number generator for sampling
//
// Example::
//
//     >>> rates = torch.rand(4, 4) * 5  # rate parameter between 0 and 5
//     >>> torch.poisson(rates)
//     tensor([[9., 1., 3., 5.],
//             [8., 6., 6., 0.],
//             [0., 4., 5., 3.],
//             [2., 1., 4., 2.]])
//
//
//go:linkname Poisson py.poisson
func Poisson(input *py.Object, generator *py.Object) *py.Object
// None
//
//go:linkname PoissonNllLoss py.poisson_nll_loss
func PoissonNllLoss(__llgo_va_list ...interface{}) *py.Object
//
// polar(abs, angle, *, out=None) -> Tensor
//
// Constructs a complex tensor whose elements are Cartesian coordinates
// corresponding to the polar coordinates with absolute value :attr:`abs` and angle
// :attr:`angle`.
//
// .. math::
//     \text{out} = \text{abs} \cdot \cos(\text{angle}) + \text{abs} \cdot \sin(\text{angle}) \cdot j
//
// .. note::
//     `torch.polar` is similar to
//     `std::polar <https://en.cppreference.com/w/cpp/numeric/complex/polar>`_
//     and does not compute the polar decomposition
//     of a complex tensor like Python's `cmath.polar` and SciPy's `linalg.polar` do.
//     The behavior of this function is undefined if `abs` is negative or NaN, or if `angle` is
//     infinite.
//
//
// Args:
//     abs (Tensor): The absolute value the complex tensor. Must be float or double.
//     angle (Tensor): The angle of the complex tensor. Must be same dtype as
//         :attr:`abs`.
//
// Keyword args:
//     out (Tensor): If the inputs are ``torch.float32``, must be
//         ``torch.complex64``. If the inputs are ``torch.float64``, must be
//         ``torch.complex128``.
//
// Example::
//
//     >>> import numpy as np
//     >>> abs = torch.tensor([1, 2], dtype=torch.float64)
//     >>> angle = torch.tensor([np.pi / 2, 5 * np.pi / 4], dtype=torch.float64)
//     >>> z = torch.polar(abs, angle)
//     >>> z
//     tensor([(0.0000+1.0000j), (-1.4142-1.4142j)], dtype=torch.complex128)
//
//
//go:linkname Polar py.polar
func Polar(abs *py.Object, angle *py.Object) *py.Object
//
// polygamma(n, input, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.polygamma`.
//
//
//go:linkname Polygamma py.polygamma
func Polygamma(n *py.Object, input *py.Object) *py.Object
//
// positive(input) -> Tensor
//
// Returns :attr:`input`.
// Throws a runtime error if :attr:`input` is a bool tensor.
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> t = torch.randn(5)
//     >>> t
//     tensor([ 0.0090, -0.2262, -0.0682, -0.2866,  0.3940])
//     >>> torch.positive(t)
//     tensor([ 0.0090, -0.2262, -0.0682, -0.2866,  0.3940])
//
//
//go:linkname Positive py.positive
func Positive(input *py.Object) *py.Object
//
// pow(input, exponent, *, out=None) -> Tensor
//
// Takes the power of each element in :attr:`input` with :attr:`exponent` and
// returns a tensor with the result.
//
// :attr:`exponent` can be either a single ``float`` number or a `Tensor`
// with the same number of elements as :attr:`input`.
//
// When :attr:`exponent` is a scalar value, the operation applied is:
//
// .. math::
//     \text{out}_i = x_i ^ \text{exponent}
//
// When :attr:`exponent` is a tensor, the operation applied is:
//
// .. math::
//     \text{out}_i = x_i ^ {\text{exponent}_i}
//
// When :attr:`exponent` is a tensor, the shapes of :attr:`input`
// and :attr:`exponent` must be :ref:`broadcastable <broadcasting-semantics>`.
//
// Args:
//     input (Tensor): the input tensor.
//     exponent (float or tensor): the exponent value
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([ 0.4331,  1.2475,  0.6834, -0.2791])
//     >>> torch.pow(a, 2)
//     tensor([ 0.1875,  1.5561,  0.4670,  0.0779])
//     >>> exp = torch.arange(1., 5.)
//
//     >>> a = torch.arange(1., 5.)
//     >>> a
//     tensor([ 1.,  2.,  3.,  4.])
//     >>> exp
//     tensor([ 1.,  2.,  3.,  4.])
//     >>> torch.pow(a, exp)
//     tensor([   1.,    4.,   27.,  256.])
//
// .. function:: pow(self, exponent, *, out=None) -> Tensor
//    :noindex:
//
// :attr:`self` is a scalar ``float`` value, and :attr:`exponent` is a tensor.
// The returned tensor :attr:`out` is of the same shape as :attr:`exponent`
//
// The operation applied is:
//
// .. math::
//     \text{out}_i = \text{self} ^ {\text{exponent}_i}
//
// Args:
//     self (float): the scalar base value for the power operation
//     exponent (Tensor): the exponent tensor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> exp = torch.arange(1., 5.)
//     >>> base = 2
//     >>> torch.pow(base, exp)
//     tensor([  2.,   4.,   8.,  16.])
//
//
//go:linkname Pow py.pow
func Pow(input *py.Object, exponent *py.Object) *py.Object
// prelu(input, weight) -> Tensor
//
// Applies element-wise the function
// :math:`\text{PReLU}(x) = \max(0,x) + \text{weight} * \min(0,x)` where weight is a
// learnable parameter.
//
// .. note::
//     `weight` is expected to be a scalar or 1-D tensor. If `weight` is 1-D,
//     its size must match the number of input channels, determined by
//     `input.size(1)` when `input.dim() >= 2`, otherwise 1.
//     In the 1-D case, note that when `input` has dim > 2, `weight` can be expanded
//     to the shape of `input` in a way that is not possible using normal
//     :ref:`broadcasting semantics<broadcasting-semantics>`.
//
// See :class:`~torch.nn.PReLU` for more details.
//
//
//go:linkname Prelu py.prelu
func Prelu(input *py.Object, weight *py.Object) *py.Object
//
// prod(input, *, dtype=None) -> Tensor
//
// Returns the product of all elements in the :attr:`input` tensor.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         If specified, the input tensor is casted to :attr:`dtype` before the operation
//         is performed. This is useful for preventing data type overflows. Default: None.
//
// Example::
//
//     >>> a = torch.randn(1, 3)
//     >>> a
//     tensor([[-0.8020,  0.5428, -1.5854]])
//     >>> torch.prod(a)
//     tensor(0.6902)
//
// .. function:: prod(input, dim, keepdim=False, *, dtype=None) -> Tensor
//    :noindex:
//
// Returns the product of each row of the :attr:`input` tensor in the given
// dimension :attr:`dim`.
//
// If :attr:`keepdim` is ``True``, the output tensor is of the same size
// as :attr:`input` except in the dimension :attr:`dim` where it is of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in
// the output tensor having 1 fewer dimension than :attr:`input`.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int): the dimension to reduce.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         If specified, the input tensor is casted to :attr:`dtype` before the operation
//         is performed. This is useful for preventing data type overflows. Default: None.
//
// Example::
//
//     >>> a = torch.randn(4, 2)
//     >>> a
//     tensor([[ 0.5261, -0.3837],
//             [ 1.1857, -0.2498],
//             [-1.1646,  0.0705],
//             [ 1.1131, -1.0629]])
//     >>> torch.prod(a, 1)
//     tensor([-0.2018, -0.2962, -0.0821, -1.1831])
//
//
//go:linkname Prod py.prod
func Prod(input *py.Object) *py.Object
//
// promote_types(type1, type2) -> dtype
//
// Returns the :class:`torch.dtype` with the smallest size and scalar kind that is
// not smaller nor of lower kind than either `type1` or `type2`. See type promotion
// :ref:`documentation <type-promotion-doc>` for more information on the type
// promotion logic.
//
// Args:
//     type1 (:class:`torch.dtype`)
//     type2 (:class:`torch.dtype`)
//
// Example::
//
//     >>> torch.promote_types(torch.int32, torch.float32)
//     torch.float32
//     >>> torch.promote_types(torch.uint8, torch.long)
//     torch.long
//
//
//go:linkname PromoteTypes py.promote_types
func PromoteTypes(type1 *py.Object, type2 *py.Object) *py.Object
// None
//
//go:linkname Put py.put
func Put(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname QPerChannelAxis py.q_per_channel_axis
func QPerChannelAxis(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname QPerChannelScales py.q_per_channel_scales
func QPerChannelScales(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname QPerChannelZeroPoints py.q_per_channel_zero_points
func QPerChannelZeroPoints(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname QScale py.q_scale
func QScale(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname QZeroPoint py.q_zero_point
func QZeroPoint(__llgo_va_list ...interface{}) *py.Object
//
// qr(input, some=True, *, out=None) -> (Tensor, Tensor)
//
// Computes the QR decomposition of a matrix or a batch of matrices :attr:`input`,
// and returns a namedtuple (Q, R) of tensors such that :math:`\text{input} = Q R`
// with :math:`Q` being an orthogonal matrix or batch of orthogonal matrices and
// :math:`R` being an upper triangular matrix or batch of upper triangular matrices.
//
// If :attr:`some` is ``True``, then this function returns the thin (reduced) QR factorization.
// Otherwise, if :attr:`some` is ``False``, this function returns the complete QR factorization.
//
// .. warning::
//
//     :func:`torch.qr` is deprecated in favor of :func:`torch.linalg.qr`
//     and will be removed in a future PyTorch release. The boolean parameter :attr:`some` has been
//     replaced with a string parameter :attr:`mode`.
//
//     ``Q, R = torch.qr(A)`` should be replaced with
//
//     .. code:: python
//
//         Q, R = torch.linalg.qr(A)
//
//     ``Q, R = torch.qr(A, some=False)`` should be replaced with
//
//     .. code:: python
//
//         Q, R = torch.linalg.qr(A, mode="complete")
//
// .. warning::
//           If you plan to backpropagate through QR, note that the current backward implementation
//           is only well-defined when the first :math:`\min(input.size(-1), input.size(-2))`
//           columns of :attr:`input` are linearly independent.
//           This behavior will probably change once QR supports pivoting.
//
// .. note:: This function uses LAPACK for CPU inputs and MAGMA for CUDA inputs,
//           and may produce different (valid) decompositions on different device types
//           or different platforms.
//
// Args:
//     input (Tensor): the input tensor of size :math:`(*, m, n)` where `*` is zero or more
//                 batch dimensions consisting of matrices of dimension :math:`m \times n`.
//     some (bool, optional): Set to ``True`` for reduced QR decomposition and ``False`` for
//                 complete QR decomposition. If `k = min(m, n)` then:
//
//                   * ``some=True`` : returns `(Q, R)` with dimensions (m, k), (k, n) (default)
//
//                   * ``'some=False'``: returns `(Q, R)` with dimensions (m, m), (m, n)
//
// Keyword args:
//     out (tuple, optional): tuple of `Q` and `R` tensors.
//                 The dimensions of `Q` and `R` are detailed in the description of :attr:`some` above.
//
// Example::
//
//     >>> a = torch.tensor([[12., -51, 4], [6, 167, -68], [-4, 24, -41]])
//     >>> q, r = torch.qr(a)
//     >>> q
//     tensor([[-0.8571,  0.3943,  0.3314],
//             [-0.4286, -0.9029, -0.0343],
//             [ 0.2857, -0.1714,  0.9429]])
//     >>> r
//     tensor([[ -14.0000,  -21.0000,   14.0000],
//             [   0.0000, -175.0000,   70.0000],
//             [   0.0000,    0.0000,  -35.0000]])
//     >>> torch.mm(q, r).round()
//     tensor([[  12.,  -51.,    4.],
//             [   6.,  167.,  -68.],
//             [  -4.,   24.,  -41.]])
//     >>> torch.mm(q.t(), q).round()
//     tensor([[ 1.,  0.,  0.],
//             [ 0.,  1., -0.],
//             [ 0., -0.,  1.]])
//     >>> a = torch.randn(3, 4, 5)
//     >>> q, r = torch.qr(a, some=False)
//     >>> torch.allclose(torch.matmul(q, r), a)
//     True
//     >>> torch.allclose(torch.matmul(q.mT, q), torch.eye(5))
//     True
//
//
//go:linkname Qr py.qr
func Qr(input *py.Object, some *py.Object) *py.Object
//
// quantile(input, q, dim=None, keepdim=False, *, interpolation='linear', out=None) -> Tensor
//
// Computes the q-th quantiles of each row of the :attr:`input` tensor along the dimension :attr:`dim`.
//
// To compute the quantile, we map q in [0, 1] to the range of indices [0, n] to find the location
// of the quantile in the sorted input. If the quantile lies between two data points ``a < b`` with
// indices ``i`` and ``j`` in the sorted order, result is computed according to the given
// :attr:`interpolation` method as follows:
//
// - ``linear``: ``a + (b - a) * fraction``, where ``fraction`` is the fractional part of the computed quantile index.
// - ``lower``: ``a``.
// - ``higher``: ``b``.
// - ``nearest``: ``a`` or ``b``, whichever's index is closer to the computed quantile index (rounding down for .5 fractions).
// - ``midpoint``: ``(a + b) / 2``.
//
// If :attr:`q` is a 1D tensor, the first dimension of the output represents the quantiles and has size
// equal to the size of :attr:`q`, the remaining dimensions are what remains from the reduction.
//
// .. note::
//     By default :attr:`dim` is ``None`` resulting in the :attr:`input` tensor being flattened before computation.
//
// Args:
//     input (Tensor): the input tensor.
//     q (float or Tensor): a scalar or 1D tensor of values in the range [0, 1].
//     dim (int): the dimension to reduce.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword arguments:
//     interpolation (str): interpolation method to use when the desired quantile lies between two data points.
//                             Can be ``linear``, ``lower``, ``higher``, ``midpoint`` and ``nearest``.
//                             Default is ``linear``.
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(2, 3)
//     >>> a
//     tensor([[ 0.0795, -1.2117,  0.9765],
//             [ 1.1707,  0.6706,  0.4884]])
//     >>> q = torch.tensor([0.25, 0.5, 0.75])
//     >>> torch.quantile(a, q, dim=1, keepdim=True)
//     tensor([[[-0.5661],
//             [ 0.5795]],
//
//             [[ 0.0795],
//             [ 0.6706]],
//
//             [[ 0.5280],
//             [ 0.9206]]])
//     >>> torch.quantile(a, q, dim=1, keepdim=True).shape
//     torch.Size([3, 2, 1])
//     >>> a = torch.arange(4.)
//     >>> a
//     tensor([0., 1., 2., 3.])
//     >>> torch.quantile(a, 0.6, interpolation='linear')
//     tensor(1.8000)
//     >>> torch.quantile(a, 0.6, interpolation='lower')
//     tensor(1.)
//     >>> torch.quantile(a, 0.6, interpolation='higher')
//     tensor(2.)
//     >>> torch.quantile(a, 0.6, interpolation='midpoint')
//     tensor(1.5000)
//     >>> torch.quantile(a, 0.6, interpolation='nearest')
//     tensor(2.)
//     >>> torch.quantile(a, 0.4, interpolation='nearest')
//     tensor(1.)
//
//
//go:linkname Quantile py.quantile
func Quantile(input *py.Object, q *py.Object, dim *py.Object, keepdim *py.Object) *py.Object
//
// quantize_per_channel(input, scales, zero_points, axis, dtype) -> Tensor
//
// Converts a float tensor to a per-channel quantized tensor with given scales and zero points.
//
// Arguments:
//     input (Tensor): float tensor to quantize
//     scales (Tensor): float 1D tensor of scales to use, size should match ``input.size(axis)``
//     zero_points (int): integer 1D tensor of offset to use, size should match ``input.size(axis)``
//     axis (int): dimension on which apply per-channel quantization
//     dtype (:class:`torch.dtype`): the desired data type of returned tensor.
//         Has to be one of the quantized dtypes: ``torch.quint8``, ``torch.qint8``, ``torch.qint32``
//
// Returns:
//     Tensor: A newly quantized tensor
//
// Example::
//
//     >>> x = torch.tensor([[-1.0, 0.0], [1.0, 2.0]])
//     >>> torch.quantize_per_channel(x, torch.tensor([0.1, 0.01]), torch.tensor([10, 0]), 0, torch.quint8)
//     tensor([[-1.,  0.],
//             [ 1.,  2.]], size=(2, 2), dtype=torch.quint8,
//            quantization_scheme=torch.per_channel_affine,
//            scale=tensor([0.1000, 0.0100], dtype=torch.float64),
//            zero_point=tensor([10,  0]), axis=0)
//     >>> torch.quantize_per_channel(x, torch.tensor([0.1, 0.01]), torch.tensor([10, 0]), 0, torch.quint8).int_repr()
//     tensor([[  0,  10],
//             [100, 200]], dtype=torch.uint8)
//
//
//go:linkname QuantizePerChannel py.quantize_per_channel
func QuantizePerChannel(input *py.Object, scales *py.Object, zeroPoints *py.Object, axis *py.Object, dtype *py.Object) *py.Object
//
// quantize_per_tensor(input, scale, zero_point, dtype) -> Tensor
//
// Converts a float tensor to a quantized tensor with given scale and zero point.
//
// Arguments:
//     input (Tensor): float tensor or list of tensors to quantize
//     scale (float or Tensor): scale to apply in quantization formula
//     zero_point (int or Tensor): offset in integer value that maps to float zero
//     dtype (:class:`torch.dtype`): the desired data type of returned tensor.
//         Has to be one of the quantized dtypes: ``torch.quint8``, ``torch.qint8``, ``torch.qint32``
//
// Returns:
//     Tensor: A newly quantized tensor or list of quantized tensors.
//
// Example::
//
//     >>> torch.quantize_per_tensor(torch.tensor([-1.0, 0.0, 1.0, 2.0]), 0.1, 10, torch.quint8)
//     tensor([-1.,  0.,  1.,  2.], size=(4,), dtype=torch.quint8,
//            quantization_scheme=torch.per_tensor_affine, scale=0.1, zero_point=10)
//     >>> torch.quantize_per_tensor(torch.tensor([-1.0, 0.0, 1.0, 2.0]), 0.1, 10, torch.quint8).int_repr()
//     tensor([ 0, 10, 20, 30], dtype=torch.uint8)
//     >>> torch.quantize_per_tensor([torch.tensor([-1.0, 0.0]), torch.tensor([-2.0, 2.0])],
//     >>> torch.tensor([0.1, 0.2]), torch.tensor([10, 20]), torch.quint8)
//     (tensor([-1.,  0.], size=(2,), dtype=torch.quint8,
//         quantization_scheme=torch.per_tensor_affine, scale=0.1, zero_point=10),
//         tensor([-2.,  2.], size=(2,), dtype=torch.quint8,
//         quantization_scheme=torch.per_tensor_affine, scale=0.2, zero_point=20))
//     >>> torch.quantize_per_tensor(torch.tensor([-1.0, 0.0, 1.0, 2.0]), torch.tensor(0.1), torch.tensor(10), torch.quint8)
//     tensor([-1.,  0.,  1.,  2.], size=(4,), dtype=torch.quint8,
//        quantization_scheme=torch.per_tensor_affine, scale=0.10, zero_point=10)
//
//
//go:linkname QuantizePerTensor py.quantize_per_tensor
func QuantizePerTensor(input *py.Object, scale *py.Object, zeroPoint *py.Object, dtype *py.Object) *py.Object
//
// quantize_per_tensor_dynamic(input, dtype, reduce_range) -> Tensor
//
// Converts a float tensor to a quantized tensor with scale and zero_point calculated
// dynamically based on the input.
//
// Arguments:
//     input (Tensor): float tensor or list of tensors to quantize
//     dtype (:class:`torch.dtype`): the desired data type of returned tensor.
//         Has to be one of the quantized dtypes: ``torch.quint8``, ``torch.qint8``
//     reduce_range (bool): a flag to indicate whether to reduce the range of quantized
//     data by 1 bit, it's required to avoid instruction overflow for some hardwares
//
// Returns:
//     Tensor: A newly (dynamically) quantized tensor
//
// Example::
//
//     >>> t = torch.quantize_per_tensor_dynamic(torch.tensor([-1.0, 0.0, 1.0, 2.0]), torch.quint8, False)
//     >>> print(t)
//     tensor([-1.,  0.,  1.,  2.], size=(4,), dtype=torch.quint8,
//            quantization_scheme=torch.per_tensor_affine, scale=0.011764705882352941,
//            zero_point=85)
//     >>> t.int_repr()
//     tensor([  0,  85, 170, 255], dtype=torch.uint8)
//
//
//go:linkname QuantizePerTensorDynamic py.quantize_per_tensor_dynamic
func QuantizePerTensorDynamic(input *py.Object, dtype *py.Object, reduceRange *py.Object) *py.Object
//
// quantized_batch_norm(input, weight=None, bias=None, mean, var, eps, output_scale, output_zero_point) -> Tensor
//
// Applies batch normalization on a 4D (NCHW) quantized tensor.
//
// .. math::
//
//         y = \frac{x - \mathrm{E}[x]}{\sqrt{\mathrm{Var}[x] + \epsilon}} * \gamma + \beta
//
// Arguments:
//     input (Tensor): quantized tensor
//     weight (Tensor): float tensor that corresponds to the gamma, size C
//     bias (Tensor):  float tensor that corresponds to the beta, size C
//     mean (Tensor): float mean value in batch normalization, size C
//     var (Tensor): float tensor for variance, size C
//     eps (float): a value added to the denominator for numerical stability.
//     output_scale (float): output quantized tensor scale
//     output_zero_point (int): output quantized tensor zero_point
//
// Returns:
//     Tensor: A quantized tensor with batch normalization applied.
//
// Example::
//
//     >>> qx = torch.quantize_per_tensor(torch.rand(2, 2, 2, 2), 1.5, 3, torch.quint8)
//     >>> torch.quantized_batch_norm(qx, torch.ones(2), torch.zeros(2), torch.rand(2), torch.rand(2), 0.00001, 0.2, 2)
//     tensor([[[[-0.2000, -0.2000],
//           [ 1.6000, -0.2000]],
//
//          [[-0.4000, -0.4000],
//           [-0.4000,  0.6000]]],
//
//
//         [[[-0.2000, -0.2000],
//           [-0.2000, -0.2000]],
//
//          [[ 0.6000, -0.4000],
//           [ 0.6000, -0.4000]]]], size=(2, 2, 2, 2), dtype=torch.quint8,
//        quantization_scheme=torch.per_tensor_affine, scale=0.2, zero_point=2)
//
//
//go:linkname QuantizedBatchNorm py.quantized_batch_norm
func QuantizedBatchNorm(input *py.Object, weight *py.Object, bias *py.Object, mean *py.Object, var_ *py.Object, eps *py.Object, outputScale *py.Object, outputZeroPoint *py.Object) *py.Object
// None
//
//go:linkname QuantizedGruCell py.quantized_gru_cell
func QuantizedGruCell(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname QuantizedLstmCell py.quantized_lstm_cell
func QuantizedLstmCell(__llgo_va_list ...interface{}) *py.Object
//
// quantized_max_pool1d(input, kernel_size, stride=[], padding=0, dilation=1, ceil_mode=False) -> Tensor
//
// Applies a 1D max pooling over an input quantized tensor composed of several input planes.
//
// Arguments:
//     input (Tensor): quantized tensor
//     kernel_size (list of int): the size of the sliding window
//     stride (``list of int``, optional): the stride of the sliding window
//     padding (``list of int``, optional): padding to be added on both sides, must be >= 0 and <= kernel_size / 2
//     dilation (``list of int``, optional): The stride between elements within a sliding window, must be > 0. Default 1
//     ceil_mode (bool, optional):  If True, will use ceil instead of floor to compute the output shape.
//         Defaults to False.
//
//
// Returns:
//     Tensor: A quantized tensor with max_pool1d applied.
//
// Example::
//
//     >>> qx = torch.quantize_per_tensor(torch.rand(2, 2), 1.5, 3, torch.quint8)
//     >>> torch.quantized_max_pool1d(qx, [2])
//     tensor([[0.0000],
//             [1.5000]], size=(2, 1), dtype=torch.quint8,
//         quantization_scheme=torch.per_tensor_affine, scale=1.5, zero_point=3)
//
//
//go:linkname QuantizedMaxPool1d py.quantized_max_pool1d
func QuantizedMaxPool1d(input *py.Object, kernelSize *py.Object, stride *py.Object, padding *py.Object, dilation *py.Object, ceilMode *py.Object) *py.Object
//
// quantized_max_pool2d(input, kernel_size, stride=[], padding=0, dilation=1, ceil_mode=False) -> Tensor
//
// Applies a 2D max pooling over an input quantized tensor composed of several input planes.
//
// Arguments:
//     input (Tensor): quantized tensor
//     kernel_size (``list of int``): the size of the sliding window
//     stride (``list of int``, optional): the stride of the sliding window
//     padding (``list of int``, optional): padding to be added on both sides, must be >= 0 and <= kernel_size / 2
//     dilation (``list of int``, optional): The stride between elements within a sliding window, must be > 0. Default 1
//     ceil_mode (bool, optional):  If True, will use ceil instead of floor to compute the output shape.
//         Defaults to False.
//
//
// Returns:
//     Tensor: A quantized tensor with max_pool2d applied.
//
// Example::
//
//     >>> qx = torch.quantize_per_tensor(torch.rand(2, 2, 2, 2), 1.5, 3, torch.quint8)
//     >>> torch.quantized_max_pool2d(qx, [2,2])
//     tensor([[[[1.5000]],
//
//             [[1.5000]]],
//
//
//             [[[0.0000]],
//
//             [[0.0000]]]], size=(2, 2, 1, 1), dtype=torch.quint8,
//         quantization_scheme=torch.per_tensor_affine, scale=1.5, zero_point=3)
//
//
//go:linkname QuantizedMaxPool2d py.quantized_max_pool2d
func QuantizedMaxPool2d(input *py.Object, kernelSize *py.Object, stride *py.Object, padding *py.Object, dilation *py.Object, ceilMode *py.Object) *py.Object
// None
//
//go:linkname QuantizedMaxPool3d py.quantized_max_pool3d
func QuantizedMaxPool3d(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname QuantizedRnnReluCell py.quantized_rnn_relu_cell
func QuantizedRnnReluCell(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname QuantizedRnnTanhCell py.quantized_rnn_tanh_cell
func QuantizedRnnTanhCell(__llgo_va_list ...interface{}) *py.Object
//
// rad2deg(input, *, out=None) -> Tensor
//
// Returns a new tensor with each of the elements of :attr:`input`
// converted from angles in radians to degrees.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword arguments:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([[3.142, -3.142], [6.283, -6.283], [1.570, -1.570]])
//     >>> torch.rad2deg(a)
//     tensor([[ 180.0233, -180.0233],
//             [ 359.9894, -359.9894],
//             [  89.9544,  -89.9544]])
//
//
//
//go:linkname Rad2deg py.rad2deg
func Rad2deg(input *py.Object) *py.Object
// None
//
//go:linkname Rad2deg_ py.rad2deg_
func Rad2deg_(__llgo_va_list ...interface{}) *py.Object
//
// rand(*size, *, generator=None, out=None, dtype=None, layout=torch.strided, device=None, requires_grad=False, pin_memory=False) -> Tensor
//
// Returns a tensor filled with random numbers from a uniform distribution
// on the interval :math:`[0, 1)`
//
// The shape of the tensor is defined by the variable argument :attr:`size`.
//
// Args:
//     size (int...): a sequence of integers defining the shape of the output tensor.
//         Can be a variable number of arguments or a collection like a list or tuple.
//
// Keyword args:
//     generator (:class:`torch.Generator`, optional): a pseudorandom number generator for sampling
//     out (Tensor, optional): the output tensor.
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`).
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     pin_memory (bool, optional): If set, returned tensor would be allocated in
//         the pinned memory. Works only for CPU tensors. Default: ``False``.
//
// Example::
//
//     >>> torch.rand(4)
//     tensor([ 0.5204,  0.2503,  0.3525,  0.5673])
//     >>> torch.rand(2, 3)
//     tensor([[ 0.8237,  0.5781,  0.6879],
//             [ 0.3816,  0.7249,  0.0998]])
//
//
//go:linkname Rand py.rand
func Rand(__llgo_va_list ...interface{}) *py.Object
//
// rand_like(input, *, dtype=None, layout=None, device=None, requires_grad=False, memory_format=torch.preserve_format) -> Tensor
//
// Returns a tensor with the same size as :attr:`input` that is filled with
// random numbers from a uniform distribution on the interval :math:`[0, 1)`.
// ``torch.rand_like(input)`` is equivalent to
// ``torch.rand(input.size(), dtype=input.dtype, layout=input.layout, device=input.device)``.
//
// Args:
//     input (Tensor): the size of :attr:`input` will determine size of the output tensor.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned Tensor.
//         Default: if ``None``, defaults to the dtype of :attr:`input`.
//     layout (:class:`torch.layout`, optional): the desired layout of returned tensor.
//         Default: if ``None``, defaults to the layout of :attr:`input`.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, defaults to the device of :attr:`input`.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     memory_format (:class:`torch.memory_format`, optional): the desired memory format of
//         returned Tensor. Default: ``torch.preserve_format``.
//
//
//
//go:linkname RandLike py.rand_like
func RandLike(input *py.Object) *py.Object
//
// randint(low=0, high, size, \*, generator=None, out=None, dtype=None, layout=torch.strided, device=None, requires_grad=False) -> Tensor
//
// Returns a tensor filled with random integers generated uniformly
// between :attr:`low` (inclusive) and :attr:`high` (exclusive).
//
// The shape of the tensor is defined by the variable argument :attr:`size`.
//
// .. note::
//     With the global dtype default (``torch.float32``), this function returns
//     a tensor with dtype ``torch.int64``.
//
// Args:
//     low (int, optional): Lowest integer to be drawn from the distribution. Default: 0.
//     high (int): One above the highest integer to be drawn from the distribution.
//     size (tuple): a tuple defining the shape of the output tensor.
//
// Keyword args:
//     generator (:class:`torch.Generator`, optional): a pseudorandom number generator for sampling
//     out (Tensor, optional): the output tensor.
//     dtype (`torch.dtype`, optional) - the desired data type of returned tensor. Default: if ``None``,
//         this function returns a tensor with dtype ``torch.int64``.
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
// Example::
//
//     >>> torch.randint(3, 5, (3,))
//     tensor([4, 3, 4])
//
//
//     >>> torch.randint(10, (2, 2))
//     tensor([[0, 2],
//             [5, 5]])
//
//
//     >>> torch.randint(3, 10, (2, 2))
//     tensor([[4, 5],
//             [6, 7]])
//
//
//
//
//go:linkname Randint py.randint
func Randint(low *py.Object, high *py.Object, size *py.Object) *py.Object
//
// randint_like(input, low=0, high, \*, dtype=None, layout=torch.strided, device=None, requires_grad=False, memory_format=torch.preserve_format) -> Tensor
//
// Returns a tensor with the same shape as Tensor :attr:`input` filled with
// random integers generated uniformly between :attr:`low` (inclusive) and
// :attr:`high` (exclusive).
//
// .. note:
//     With the global dtype default (``torch.float32``), this function returns
//     a tensor with dtype ``torch.int64``.
//
// Args:
//     input (Tensor): the size of :attr:`input` will determine size of the output tensor.
//     low (int, optional): Lowest integer to be drawn from the distribution. Default: 0.
//     high (int): One above the highest integer to be drawn from the distribution.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned Tensor.
//         Default: if ``None``, defaults to the dtype of :attr:`input`.
//     layout (:class:`torch.layout`, optional): the desired layout of returned tensor.
//         Default: if ``None``, defaults to the layout of :attr:`input`.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, defaults to the device of :attr:`input`.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     memory_format (:class:`torch.memory_format`, optional): the desired memory format of
//         returned Tensor. Default: ``torch.preserve_format``.
//
//
//
//go:linkname RandintLike py.randint_like
func RandintLike(input *py.Object, low *py.Object, high *py.Object) *py.Object
//
// randn(*size, *, generator=None, out=None, dtype=None, layout=torch.strided, device=None, requires_grad=False, pin_memory=False) -> Tensor
//
//
// Returns a tensor filled with random numbers from a normal distribution
// with mean `0` and variance `1` (also called the standard normal
// distribution).
//
// .. math::
//     \text{out}_{i} \sim \mathcal{N}(0, 1)
//
// The shape of the tensor is defined by the variable argument :attr:`size`.
//
// Args:
//     size (int...): a sequence of integers defining the shape of the output tensor.
//         Can be a variable number of arguments or a collection like a list or tuple.
//
// Keyword args:
//     generator (:class:`torch.Generator`, optional): a pseudorandom number generator for sampling
//     out (Tensor, optional): the output tensor.
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`).
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     pin_memory (bool, optional): If set, returned tensor would be allocated in
//         the pinned memory. Works only for CPU tensors. Default: ``False``.
//
// Example::
//
//     >>> torch.randn(4)
//     tensor([-2.1436,  0.9966,  2.3426, -0.6366])
//     >>> torch.randn(2, 3)
//     tensor([[ 1.5954,  2.8929, -1.0923],
//             [ 1.1719, -0.4709, -0.1996]])
//
//
//go:linkname Randn py.randn
func Randn(__llgo_va_list ...interface{}) *py.Object
//
// randn_like(input, *, dtype=None, layout=None, device=None, requires_grad=False, memory_format=torch.preserve_format) -> Tensor
//
// Returns a tensor with the same size as :attr:`input` that is filled with
// random numbers from a normal distribution with mean 0 and variance 1.
// ``torch.randn_like(input)`` is equivalent to
// ``torch.randn(input.size(), dtype=input.dtype, layout=input.layout, device=input.device)``.
//
// Args:
//     input (Tensor): the size of :attr:`input` will determine size of the output tensor.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned Tensor.
//         Default: if ``None``, defaults to the dtype of :attr:`input`.
//     layout (:class:`torch.layout`, optional): the desired layout of returned tensor.
//         Default: if ``None``, defaults to the layout of :attr:`input`.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, defaults to the device of :attr:`input`.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     memory_format (:class:`torch.memory_format`, optional): the desired memory format of
//         returned Tensor. Default: ``torch.preserve_format``.
//
//
//
//go:linkname RandnLike py.randn_like
func RandnLike(input *py.Object) *py.Object
//
// randperm(n, *, generator=None, out=None, dtype=torch.int64,layout=torch.strided, device=None, requires_grad=False, pin_memory=False) -> Tensor
//
// Returns a random permutation of integers from ``0`` to ``n - 1``.
//
// Args:
//     n (int): the upper bound (exclusive)
//
// Keyword args:
//     generator (:class:`torch.Generator`, optional): a pseudorandom number generator for sampling
//     out (Tensor, optional): the output tensor.
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: ``torch.int64``.
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     pin_memory (bool, optional): If set, returned tensor would be allocated in
//         the pinned memory. Works only for CPU tensors. Default: ``False``.
//
// Example::
//
//     >>> torch.randperm(4)
//     tensor([2, 1, 0, 3])
//
//
//go:linkname Randperm py.randperm
func Randperm(n *py.Object) *py.Object
//
// range(start=0, end, step=1, *, out=None, dtype=None, layout=torch.strided, device=None, requires_grad=False) -> Tensor
//
// Returns a 1-D tensor of size :math:`\left\lfloor \frac{\text{end} - \text{start}}{\text{step}} \right\rfloor + 1`
// with values from :attr:`start` to :attr:`end` with step :attr:`step`. Step is
// the gap between two values in the tensor.
//
// .. math::
//     \text{out}_{i+1} = \text{out}_i + \text{step}.
//
// .. warning::
//     This function is deprecated and will be removed in a future release because its behavior is inconsistent with
//     Python's range builtin. Instead, use :func:`torch.arange`, which produces values in [start, end).
//
// Args:
//     start (float): the starting value for the set of points. Default: ``0``.
//     end (float): the ending value for the set of points
//     step (float): the gap between each pair of adjacent points. Default: ``1``.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`). If `dtype` is not given, infer the data type from the other input
//         arguments. If any of `start`, `end`, or `stop` are floating-point, the
//         `dtype` is inferred to be the default dtype, see
//         :meth:`~torch.get_default_dtype`. Otherwise, the `dtype` is inferred to
//         be `torch.int64`.
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
// Example::
//
//     >>> torch.range(1, 4)
//     tensor([ 1.,  2.,  3.,  4.])
//     >>> torch.range(1, 4, 0.5)
//     tensor([ 1.0000,  1.5000,  2.0000,  2.5000,  3.0000,  3.5000,  4.0000])
//
//
//go:linkname Range py.range
func Range(start *py.Object, end *py.Object, step *py.Object) *py.Object
//
// ravel(input) -> Tensor
//
// Return a contiguous flattened tensor. A copy is made only if needed.
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> t = torch.tensor([[[1, 2],
//     ...                    [3, 4]],
//     ...                   [[5, 6],
//     ...                    [7, 8]]])
//     >>> torch.ravel(t)
//     tensor([1, 2, 3, 4, 5, 6, 7, 8])
//
//
//go:linkname Ravel py.ravel
func Ravel(input *py.Object) *py.Object
//
// real(input) -> Tensor
//
// Returns a new tensor containing real values of the :attr:`self` tensor.
// The returned tensor and :attr:`self` share the same underlying storage.
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> x=torch.randn(4, dtype=torch.cfloat)
//     >>> x
//     tensor([(0.3100+0.3553j), (-0.5445-0.7896j), (-1.6492-0.0633j), (-0.0638-0.8119j)])
//     >>> x.real
//     tensor([ 0.3100, -0.5445, -1.6492, -0.0638])
//
//
//
//go:linkname Real py.real
func Real(input *py.Object) *py.Object
//
// reciprocal(input, *, out=None) -> Tensor
//
// Returns a new tensor with the reciprocal of the elements of :attr:`input`
//
// .. math::
//     \text{out}_{i} = \frac{1}{\text{input}_{i}}
//
// .. note::
//     Unlike NumPy's reciprocal, torch.reciprocal supports integral inputs. Integral
//     inputs to reciprocal are automatically :ref:`promoted <type-promotion-doc>` to
//     the default scalar type.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([-0.4595, -2.1219, -1.4314,  0.7298])
//     >>> torch.reciprocal(a)
//     tensor([-2.1763, -0.4713, -0.6986,  1.3702])
//
//
//go:linkname Reciprocal py.reciprocal
func Reciprocal(input *py.Object) *py.Object
// None
//
//go:linkname Reciprocal_ py.reciprocal_
func Reciprocal_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Relu py.relu
func Relu(__llgo_va_list ...interface{}) *py.Object
//
// relu_(input) -> Tensor
//
// In-place version of :func:`~relu`.
//
//
//go:linkname Relu_ py.relu_
func Relu_(input *py.Object) *py.Object
//
// remainder(input, other, *, out=None) -> Tensor
//
// Computes
// `Python's modulus operation <https://docs.python.org/3/reference/expressions.html#binary-arithmetic-operations>`_
// entrywise.  The result has the same sign as the divisor :attr:`other` and its absolute value
// is less than that of :attr:`other`.
//
// It may also be defined in terms of :func:`torch.div` as
//
// .. code:: python
//
//     torch.remainder(a, b) == a - a.div(b, rounding_mode="floor") * b
//
// Supports :ref:`broadcasting to a common shape <broadcasting-semantics>`,
// :ref:`type promotion <type-promotion-doc>`, and integer and float inputs.
//
// .. note::
//     Complex inputs are not supported. In some cases, it is not mathematically
//     possible to satisfy the definition of a modulo operation with complex numbers.
//     See :func:`torch.fmod` for how division by zero is handled.
//
// .. seealso::
//
//     :func:`torch.fmod` which implements C++'s `std::fmod <https://en.cppreference.com/w/cpp/numeric/math/fmod>`_.
//     This one is defined in terms of division rounding towards zero.
//
// Args:
//     input (Tensor or Scalar): the dividend
//     other (Tensor or Scalar): the divisor
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.remainder(torch.tensor([-3., -2, -1, 1, 2, 3]), 2)
//     tensor([ 1.,  0.,  1.,  1.,  0.,  1.])
//     >>> torch.remainder(torch.tensor([1, 2, 3, 4, 5]), -1.5)
//     tensor([ -0.5000, -1.0000,  0.0000, -0.5000, -1.0000 ])
//
//
//go:linkname Remainder py.remainder
func Remainder(input *py.Object, other *py.Object) *py.Object
//
// renorm(input, p, dim, maxnorm, *, out=None) -> Tensor
//
// Returns a tensor where each sub-tensor of :attr:`input` along dimension
// :attr:`dim` is normalized such that the `p`-norm of the sub-tensor is lower
// than the value :attr:`maxnorm`
//
// .. note:: If the norm of a row is lower than `maxnorm`, the row is unchanged
//
// Args:
//     input (Tensor): the input tensor.
//     p (float): the power for the norm computation
//     dim (int): the dimension to slice over to get the sub-tensors
//     maxnorm (float): the maximum norm to keep each sub-tensor under
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> x = torch.ones(3, 3)
//     >>> x[1].fill_(2)
//     tensor([ 2.,  2.,  2.])
//     >>> x[2].fill_(3)
//     tensor([ 3.,  3.,  3.])
//     >>> x
//     tensor([[ 1.,  1.,  1.],
//             [ 2.,  2.,  2.],
//             [ 3.,  3.,  3.]])
//     >>> torch.renorm(x, 1, 0, 5)
//     tensor([[ 1.0000,  1.0000,  1.0000],
//             [ 1.6667,  1.6667,  1.6667],
//             [ 1.6667,  1.6667,  1.6667]])
//
//
//go:linkname Renorm py.renorm
func Renorm(input *py.Object, p *py.Object, dim *py.Object, maxnorm *py.Object) *py.Object
//
// repeat_interleave(input, repeats, dim=None, *, output_size=None) -> Tensor
//
// Repeat elements of a tensor.
//
// .. warning::
//
//     This is different from :meth:`torch.Tensor.repeat` but similar to ``numpy.repeat``.
//
// Args:
//     input (Tensor): the input tensor.
//     repeats (Tensor or int): The number of repetitions for each element.
//         repeats is broadcasted to fit the shape of the given axis.
//     dim (int, optional): The dimension along which to repeat values.
//         By default, use the flattened input array, and return a flat output
//         array.
//
// Keyword args:
//     output_size (int, optional): Total output size for the given axis
//         ( e.g. sum of repeats). If given, it will avoid stream synchronization
//         needed to calculate output shape of the tensor.
//
// Returns:
//     Tensor: Repeated tensor which has the same shape as input, except along the given axis.
//
// Example::
//
//     >>> x = torch.tensor([1, 2, 3])
//     >>> x.repeat_interleave(2)
//     tensor([1, 1, 2, 2, 3, 3])
//     >>> y = torch.tensor([[1, 2], [3, 4]])
//     >>> torch.repeat_interleave(y, 2)
//     tensor([1, 1, 2, 2, 3, 3, 4, 4])
//     >>> torch.repeat_interleave(y, 3, dim=1)
//     tensor([[1, 1, 1, 2, 2, 2],
//             [3, 3, 3, 4, 4, 4]])
//     >>> torch.repeat_interleave(y, torch.tensor([1, 2]), dim=0)
//     tensor([[1, 2],
//             [3, 4],
//             [3, 4]])
//     >>> torch.repeat_interleave(y, torch.tensor([1, 2]), dim=0, output_size=3)
//     tensor([[1, 2],
//             [3, 4],
//             [3, 4]])
//
// If the `repeats` is `tensor([n1, n2, n3, ...])`, then the output will be
// `tensor([0, 0, ..., 1, 1, ..., 2, 2, ..., ...])` where `0` appears `n1` times,
// `1` appears `n2` times, `2` appears `n3` times, etc.
//
// .. function:: repeat_interleave(repeats, *) -> Tensor
//    :noindex:
//
// Repeats 0 repeats[0] times, 1 repeats[1] times, 2 repeats[2] times, etc.
//
// Args:
//     repeats (Tensor): The number of repetitions for each element.
//
// Returns:
//     Tensor: Repeated tensor of size `sum(repeats)`.
//
// Example::
//
//     >>> torch.repeat_interleave(torch.tensor([1, 2, 3]))
//     tensor([0, 1, 1, 2, 2, 2])
//
//
//
//go:linkname RepeatInterleave py.repeat_interleave
func RepeatInterleave(input *py.Object, repeats *py.Object, dim *py.Object) *py.Object
//
// reshape(input, shape) -> Tensor
//
// Returns a tensor with the same data and number of elements as :attr:`input`,
// but with the specified shape. When possible, the returned tensor will be a view
// of :attr:`input`. Otherwise, it will be a copy. Contiguous inputs and inputs
// with compatible strides can be reshaped without copying, but you should not
// depend on the copying vs. viewing behavior.
//
// See :meth:`torch.Tensor.view` on when it is possible to return a view.
//
// A single dimension may be -1, in which case it's inferred from the remaining
// dimensions and the number of elements in :attr:`input`.
//
// Args:
//     input (Tensor): the tensor to be reshaped
//     shape (tuple of int): the new shape
//
// Example::
//
//     >>> a = torch.arange(4.)
//     >>> torch.reshape(a, (2, 2))
//     tensor([[ 0.,  1.],
//             [ 2.,  3.]])
//     >>> b = torch.tensor([[0, 1], [2, 3]])
//     >>> torch.reshape(b, (-1,))
//     tensor([ 0,  1,  2,  3])
//
//
//go:linkname Reshape py.reshape
func Reshape(input *py.Object, shape *py.Object) *py.Object
// None
//
//go:linkname ResizeAs_ py.resize_as_
func ResizeAs_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname ResizeAsSparse_ py.resize_as_sparse_
func ResizeAsSparse_(__llgo_va_list ...interface{}) *py.Object
//
// resolve_conj(input) -> Tensor
//
// Returns a new tensor with materialized conjugation if :attr:`input`'s conjugate bit is set to `True`,
// else returns :attr:`input`. The output tensor will always have its conjugate bit set to `False`.
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> x = torch.tensor([-1 + 1j, -2 + 2j, 3 - 3j])
//     >>> y = x.conj()
//     >>> y.is_conj()
//     True
//     >>> z = y.resolve_conj()
//     >>> z
//     tensor([-1 - 1j, -2 - 2j, 3 + 3j])
//     >>> z.is_conj()
//     False
//
//
//go:linkname ResolveConj py.resolve_conj
func ResolveConj(input *py.Object) *py.Object
//
// resolve_neg(input) -> Tensor
//
// Returns a new tensor with materialized negation if :attr:`input`'s negative bit is set to `True`,
// else returns :attr:`input`. The output tensor will always have its negative bit set to `False`.
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> x = torch.tensor([-1 + 1j, -2 + 2j, 3 - 3j])
//     >>> y = x.conj()
//     >>> z = y.imag
//     >>> z.is_neg()
//     True
//     >>> out = z.resolve_neg()
//     >>> out
//     tensor([-1., -2., 3.])
//     >>> out.is_neg()
//     False
//
//
//go:linkname ResolveNeg py.resolve_neg
func ResolveNeg(input *py.Object) *py.Object
//
// result_type(tensor1, tensor2) -> dtype
//
// Returns the :class:`torch.dtype` that would result from performing an arithmetic
// operation on the provided input tensors. See type promotion :ref:`documentation <type-promotion-doc>`
// for more information on the type promotion logic.
//
// Args:
//     tensor1 (Tensor or Number): an input tensor or number
//     tensor2 (Tensor or Number): an input tensor or number
//
// Example::
//
//     >>> torch.result_type(torch.tensor([1, 2], dtype=torch.int), 1.0)
//     torch.float32
//     >>> torch.result_type(torch.tensor([1, 2], dtype=torch.uint8), torch.tensor(1))
//     torch.uint8
//
//
//go:linkname ResultType py.result_type
func ResultType(tensor1 *py.Object, tensor2 *py.Object) *py.Object
// None
//
//go:linkname RnnRelu py.rnn_relu
func RnnRelu(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname RnnReluCell py.rnn_relu_cell
func RnnReluCell(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname RnnTanh py.rnn_tanh
func RnnTanh(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname RnnTanhCell py.rnn_tanh_cell
func RnnTanhCell(__llgo_va_list ...interface{}) *py.Object
//
// roll(input, shifts, dims=None) -> Tensor
//
// Roll the tensor :attr:`input` along the given dimension(s). Elements that are
// shifted beyond the last position are re-introduced at the first position. If
// :attr:`dims` is `None`, the tensor will be flattened before rolling and then
// restored to the original shape.
//
// Args:
//     input (Tensor): the input tensor.
//     shifts (int or tuple of ints): The number of places by which the elements
//         of the tensor are shifted. If shifts is a tuple, dims must be a tuple of
//         the same size, and each dimension will be rolled by the corresponding
//         value
//     dims (int or tuple of ints): Axis along which to roll
//
// Example::
//
//     >>> x = torch.tensor([1, 2, 3, 4, 5, 6, 7, 8]).view(4, 2)
//     >>> x
//     tensor([[1, 2],
//             [3, 4],
//             [5, 6],
//             [7, 8]])
//     >>> torch.roll(x, 1)
//     tensor([[8, 1],
//             [2, 3],
//             [4, 5],
//             [6, 7]])
//     >>> torch.roll(x, 1, 0)
//     tensor([[7, 8],
//             [1, 2],
//             [3, 4],
//             [5, 6]])
//     >>> torch.roll(x, -1, 0)
//     tensor([[3, 4],
//             [5, 6],
//             [7, 8],
//             [1, 2]])
//     >>> torch.roll(x, shifts=(2, 1), dims=(0, 1))
//     tensor([[6, 5],
//             [8, 7],
//             [2, 1],
//             [4, 3]])
//
//
//go:linkname Roll py.roll
func Roll(input *py.Object, shifts *py.Object, dims *py.Object) *py.Object
//
// rot90(input, k=1, dims=[0,1]) -> Tensor
//
// Rotate an n-D tensor by 90 degrees in the plane specified by dims axis.
// Rotation direction is from the first towards the second axis if k > 0, and from the second towards the first for k < 0.
//
// Args:
//     input (Tensor): the input tensor.
//     k (int): number of times to rotate. Default value is 1
//     dims (a list or tuple): axis to rotate. Default value is [0, 1]
//
// Example::
//
//     >>> x = torch.arange(4).view(2, 2)
//     >>> x
//     tensor([[0, 1],
//             [2, 3]])
//     >>> torch.rot90(x, 1, [0, 1])
//     tensor([[1, 3],
//             [0, 2]])
//
//     >>> x = torch.arange(8).view(2, 2, 2)
//     >>> x
//     tensor([[[0, 1],
//              [2, 3]],
//
//             [[4, 5],
//              [6, 7]]])
//     >>> torch.rot90(x, 1, [1, 2])
//     tensor([[[1, 3],
//              [0, 2]],
//
//             [[5, 7],
//              [4, 6]]])
//
//
//go:linkname Rot90 py.rot90
func Rot90(input *py.Object, k *py.Object, dims *py.Object) *py.Object
//
// round(input, *, decimals=0, out=None) -> Tensor
//
// Rounds elements of :attr:`input` to the nearest integer.
//
// For integer inputs, follows the array-api convention of returning a
// copy of the input tensor.
// The return type of output is same as that of input's dtype.
//
// .. note::
//     This function implements the "round half to even" to
//     break ties when a number is equidistant from two
//     integers (e.g. `round(2.5)` is 2).
//
//     When the :attr:\`decimals\` argument is specified the
//     algorithm used is similar to NumPy's `around`. This
//     algorithm is fast but inexact and it can easily
//     overflow for low precision dtypes.
//     Eg. `round(tensor([10000], dtype=torch.float16), decimals=3)` is `inf`.
//
// .. seealso::
//     :func:`torch.ceil`, which rounds up.
//     :func:`torch.floor`, which rounds down.
//     :func:`torch.trunc`, which rounds towards zero.
//
// Args:
//     input (Tensor): the input tensor.
//     decimals (int): Number of decimal places to round to (default: 0).
//         If decimals is negative, it specifies the number of positions
//         to the left of the decimal point.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> torch.round(torch.tensor((4.7, -2.3, 9.1, -7.7)))
//     tensor([ 5.,  -2.,  9., -8.])
//
//     >>> # Values equidistant from two integers are rounded towards the
//     >>> #   the nearest even value (zero is treated as even)
//     >>> torch.round(torch.tensor([-0.5, 0.5, 1.5, 2.5]))
//     tensor([-0., 0., 2., 2.])
//
//     >>> # A positive decimals argument rounds to the to that decimal place
//     >>> torch.round(torch.tensor([0.1234567]), decimals=3)
//     tensor([0.1230])
//
//     >>> # A negative decimals argument rounds to the left of the decimal
//     >>> torch.round(torch.tensor([1200.1234567]), decimals=-3)
//     tensor([1000.])
//
//
//go:linkname Round py.round
func Round(input *py.Object) *py.Object
// None
//
//go:linkname Round_ py.round_
func Round_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname RowIndicesCopy py.row_indices_copy
func RowIndicesCopy(__llgo_va_list ...interface{}) *py.Object
//
// row_stack(tensors, *, out=None) -> Tensor
//
// Alias of :func:`torch.vstack`.
//
//
//go:linkname RowStack py.row_stack
func RowStack(tensors *py.Object) *py.Object
// None
//
//go:linkname Rrelu py.rrelu
func Rrelu(__llgo_va_list ...interface{}) *py.Object
//
// rrelu_(input, lower=1./8, upper=1./3, training=False) -> Tensor
//
// In-place version of :func:`~rrelu`.
//
//
//go:linkname Rrelu_ py.rrelu_
func Rrelu_(input *py.Object, lower *py.Object, upper *py.Object, training *py.Object) *py.Object
//
// rsqrt(input, *, out=None) -> Tensor
//
// Returns a new tensor with the reciprocal of the square-root of each of
// the elements of :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \frac{1}{\sqrt{\text{input}_{i}}}
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([-0.0370,  0.2970,  1.5420, -0.9105])
//     >>> torch.rsqrt(a)
//     tensor([    nan,  1.8351,  0.8053,     nan])
//
//
//go:linkname Rsqrt py.rsqrt
func Rsqrt(input *py.Object) *py.Object
// None
//
//go:linkname Rsqrt_ py.rsqrt_
func Rsqrt_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Rsub py.rsub
func Rsub(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Saddmm py.saddmm
func Saddmm(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname ScalarTensor py.scalar_tensor
func ScalarTensor(__llgo_va_list ...interface{}) *py.Object
//
// scatter(input, dim, index, src) -> Tensor
//
// Out-of-place version of :meth:`torch.Tensor.scatter_`
//
//
//go:linkname Scatter py.scatter
func Scatter(input *py.Object, dim *py.Object, index *py.Object, src *py.Object) *py.Object
//
// scatter_add(input, dim, index, src) -> Tensor
//
// Out-of-place version of :meth:`torch.Tensor.scatter_add_`
//
//
//go:linkname ScatterAdd py.scatter_add
func ScatterAdd(input *py.Object, dim *py.Object, index *py.Object, src *py.Object) *py.Object
//
// scatter_reduce(input, dim, index, src, reduce, *, include_self=True) -> Tensor
//
// Out-of-place version of :meth:`torch.Tensor.scatter_reduce_`
//
//
//go:linkname ScatterReduce py.scatter_reduce
func ScatterReduce(input *py.Object, dim *py.Object, index *py.Object, src *py.Object, reduce *py.Object) *py.Object
//
// searchsorted(sorted_sequence, values, *, out_int32=False, right=False, side='left', out=None, sorter=None) -> Tensor
//
// Find the indices from the *innermost* dimension of :attr:`sorted_sequence` such that, if the
// corresponding values in :attr:`values` were inserted before the indices, when sorted, the order
// of the corresponding *innermost* dimension within :attr:`sorted_sequence` would be preserved.
// Return a new tensor with the same size as :attr:`values`. More formally,
// the returned index satisfies the following rules:
//
// .. list-table::
//    :widths: 12 10 78
//    :header-rows: 1
//
//    * - :attr:`sorted_sequence`
//      - :attr:`right`
//      - *returned index satisfies*
//    * - 1-D
//      - False
//      - ``sorted_sequence[i-1] < values[m][n]...[l][x] <= sorted_sequence[i]``
//    * - 1-D
//      - True
//      - ``sorted_sequence[i-1] <= values[m][n]...[l][x] < sorted_sequence[i]``
//    * - N-D
//      - False
//      - ``sorted_sequence[m][n]...[l][i-1] < values[m][n]...[l][x] <= sorted_sequence[m][n]...[l][i]``
//    * - N-D
//      - True
//      - ``sorted_sequence[m][n]...[l][i-1] <= values[m][n]...[l][x] < sorted_sequence[m][n]...[l][i]``
//
// Args:
//     sorted_sequence (Tensor): N-D or 1-D tensor, containing monotonically increasing sequence on the *innermost*
//                               dimension unless :attr:`sorter` is provided, in which case the sequence does not
//                               need to be sorted
//     values (Tensor or Scalar): N-D tensor or a Scalar containing the search value(s).
//
// Keyword args:
//     out_int32 (bool, optional): indicate the output data type. torch.int32 if True, torch.int64 otherwise.
//                                 Default value is False, i.e. default output data type is torch.int64.
//     right (bool, optional): if False, return the first suitable location that is found. If True, return the
//                             last such index. If no suitable index found, return 0 for non-numerical value
//                             (eg. nan, inf) or the size of *innermost* dimension within :attr:`sorted_sequence`
//                             (one pass the last index of the *innermost* dimension). In other words, if False,
//                             gets the lower bound index for each value in :attr:`values` on the corresponding
//                             *innermost* dimension of the :attr:`sorted_sequence`. If True, gets the upper
//                             bound index instead. Default value is False. :attr:`side` does the same and is
//                             preferred. It will error if :attr:`side` is set to "left" while this is True.
//     side (str, optional): the same as :attr:`right` but preferred. "left" corresponds to False for :attr:`right`
//                             and "right" corresponds to True for :attr:`right`. It will error if this is set to
//                             "left" while :attr:`right` is True.
//     out (Tensor, optional): the output tensor, must be the same size as :attr:`values` if provided.
//     sorter (LongTensor, optional): if provided, a tensor matching the shape of the unsorted
//                             :attr:`sorted_sequence` containing a sequence of indices that sort it in the
//                             ascending order on the innermost dimension
//
//
// Example::
//
//     >>> sorted_sequence = torch.tensor([[1, 3, 5, 7, 9], [2, 4, 6, 8, 10]])
//     >>> sorted_sequence
//     tensor([[ 1,  3,  5,  7,  9],
//             [ 2,  4,  6,  8, 10]])
//     >>> values = torch.tensor([[3, 6, 9], [3, 6, 9]])
//     >>> values
//     tensor([[3, 6, 9],
//             [3, 6, 9]])
//     >>> torch.searchsorted(sorted_sequence, values)
//     tensor([[1, 3, 4],
//             [1, 2, 4]])
//     >>> torch.searchsorted(sorted_sequence, values, side='right')
//     tensor([[2, 3, 5],
//             [1, 3, 4]])
//
//     >>> sorted_sequence_1d = torch.tensor([1, 3, 5, 7, 9])
//     >>> sorted_sequence_1d
//     tensor([1, 3, 5, 7, 9])
//     >>> torch.searchsorted(sorted_sequence_1d, values)
//     tensor([[1, 3, 4],
//             [1, 3, 4]])
//
//
//go:linkname Searchsorted py.searchsorted
func Searchsorted(sortedSequence *py.Object, values *py.Object) *py.Object
// None
//
//go:linkname SegmentReduce py.segment_reduce
func SegmentReduce(__llgo_va_list ...interface{}) *py.Object
//
// select(input, dim, index) -> Tensor
//
// Slices the :attr:`input` tensor along the selected dimension at the given index.
// This function returns a view of the original tensor with the given dimension removed.
//
// .. note:: If :attr:`input` is a sparse tensor and returning a view of
//           the tensor is not possible, a RuntimeError exception is
//           raised. In this is the case, consider using
//           :func:`torch.select_copy` function.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int): the dimension to slice
//     index (int): the index to select with
//
// .. note::
//
//     :meth:`select` is equivalent to slicing. For example,
//     ``tensor.select(0, index)`` is equivalent to ``tensor[index]`` and
//     ``tensor.select(2, index)`` is equivalent to ``tensor[:,:,index]``.
//
//
//go:linkname Select py.select
func Select(input *py.Object, dim *py.Object, index *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.select`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname SelectCopy py.select_copy
func SelectCopy(__llgo_va_list ...interface{}) *py.Object
//
// select_scatter(input, src, dim, index) -> Tensor
//
// Embeds the values of the :attr:`src` tensor into :attr:`input` at the given index.
// This function returns a tensor with fresh storage; it does not create a view.
//
//
// Args:
//     input (Tensor): the input tensor.
//     src (Tensor): The tensor to embed into :attr:`input`
//     dim (int): the dimension to insert the slice into.
//     index (int): the index to select with
//
// .. note::
//
//     :attr:`src` must be of the proper size in order to be embedded
//     into :attr:`input`. Specifically, it should have the same shape as
//     ``torch.select(input, dim, index)``
//
// Example::
//
//     >>> a = torch.zeros(2, 2)
//     >>> b = torch.ones(2)
//     >>> a.select_scatter(b, 0, 0)
//     tensor([[1., 1.],
//             [0., 0.]])
//
//
//go:linkname SelectScatter py.select_scatter
func SelectScatter(input *py.Object, src *py.Object, dim *py.Object, index *py.Object) *py.Object
// None
//
//go:linkname Selu py.selu
func Selu(__llgo_va_list ...interface{}) *py.Object
//
// selu_(input) -> Tensor
//
// In-place version of :func:`~selu`.
//
//
//go:linkname Selu_ py.selu_
func Selu_(input *py.Object) *py.Object
//
// sgn(input, *, out=None) -> Tensor
//
// This function is an extension of torch.sign() to complex tensors.
// It computes a new tensor whose elements have
// the same angles as the corresponding elements of :attr:`input` and
// absolute values (i.e. magnitudes) of one for complex tensors and
// is equivalent to torch.sign() for non-complex tensors.
//
// .. math::
//     \text{out}_{i} = \begin{cases}
//                     0 & |\text{{input}}_i| == 0 \\
//                     \frac{{\text{{input}}_i}}{|{\text{{input}}_i}|} & \text{otherwise}
//                     \end{cases}
//
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//   out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> t = torch.tensor([3+4j, 7-24j, 0, 1+2j])
//     >>> t.sgn()
//     tensor([0.6000+0.8000j, 0.2800-0.9600j, 0.0000+0.0000j, 0.4472+0.8944j])
//
//
//go:linkname Sgn py.sgn
func Sgn(input *py.Object) *py.Object
//
// sigmoid(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.expit`.
//
//
//go:linkname Sigmoid py.sigmoid
func Sigmoid(input *py.Object) *py.Object
// None
//
//go:linkname Sigmoid_ py.sigmoid_
func Sigmoid_(__llgo_va_list ...interface{}) *py.Object
//
// sign(input, *, out=None) -> Tensor
//
// Returns a new tensor with the signs of the elements of :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \operatorname{sgn}(\text{input}_{i})
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([0.7, -1.2, 0., 2.3])
//     >>> a
//     tensor([ 0.7000, -1.2000,  0.0000,  2.3000])
//     >>> torch.sign(a)
//     tensor([ 1., -1.,  0.,  1.])
//
//
//go:linkname Sign py.sign
func Sign(input *py.Object) *py.Object
//
// signbit(input, *, out=None) -> Tensor
//
// Tests if each element of :attr:`input` has its sign bit set or not.
//
// Args:
//   input (Tensor): the input tensor.
//
// Keyword args:
//   out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([0.7, -1.2, 0., 2.3])
//     >>> torch.signbit(a)
//     tensor([ False, True,  False,  False])
//     >>> a = torch.tensor([-0.0, 0.0])
//     >>> torch.signbit(a)
//     tensor([ True,  False])
//
// .. note::
//     signbit handles signed zeros, so negative zero (-0) returns True.
//
//
//
//go:linkname Signbit py.signbit
func Signbit(input *py.Object) *py.Object
//
// sin(input, *, out=None) -> Tensor
//
// Returns a new tensor with the sine of the elements of :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \sin(\text{input}_{i})
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([-0.5461,  0.1347, -2.7266, -0.2746])
//     >>> torch.sin(a)
//     tensor([-0.5194,  0.1343, -0.4032, -0.2711])
//
//
//go:linkname Sin py.sin
func Sin(input *py.Object) *py.Object
// None
//
//go:linkname Sin_ py.sin_
func Sin_(__llgo_va_list ...interface{}) *py.Object
//
// sinc(input, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.sinc`.
//
//
//go:linkname Sinc py.sinc
func Sinc(input *py.Object) *py.Object
// None
//
//go:linkname Sinc_ py.sinc_
func Sinc_(__llgo_va_list ...interface{}) *py.Object
//
// sinh(input, *, out=None) -> Tensor
//
// Returns a new tensor with the hyperbolic sine of the elements of
// :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \sinh(\text{input}_{i})
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([ 0.5380, -0.8632, -0.1265,  0.9399])
//     >>> torch.sinh(a)
//     tensor([ 0.5644, -0.9744, -0.1268,  1.0845])
//
// .. note::
//    When :attr:`input` is on the CPU, the implementation of torch.sinh may use
//    the Sleef library, which rounds very large results to infinity or negative
//    infinity. See `here <https://sleef.org/purec.xhtml>`_ for details.
//
//
//go:linkname Sinh py.sinh
func Sinh(input *py.Object) *py.Object
// None
//
//go:linkname Sinh_ py.sinh_
func Sinh_(__llgo_va_list ...interface{}) *py.Object
//
// Performs the same operation as :func:`torch.slice`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname SliceCopy py.slice_copy
func SliceCopy(__llgo_va_list ...interface{}) *py.Object
//
// slice_scatter(input, src, dim=0, start=None, end=None, step=1) -> Tensor
//
// Embeds the values of the :attr:`src` tensor into :attr:`input` at the given
// dimension.
// This function returns a tensor with fresh storage; it does not create a view.
//
//
// Args:
//     input (Tensor): the input tensor.
//     src (Tensor): The tensor to embed into :attr:`input`
//     dim (int): the dimension to insert the slice into
//     start (Optional[int]): the start index of where to insert the slice
//     end (Optional[int]): the end index of where to insert the slice
//     step (int): the how many elements to skip in
//
// Example::
//
//     >>> a = torch.zeros(8, 8)
//     >>> b = torch.ones(2, 8)
//     >>> a.slice_scatter(b, start=6)
//     tensor([[0., 0., 0., 0., 0., 0., 0., 0.],
//             [0., 0., 0., 0., 0., 0., 0., 0.],
//             [0., 0., 0., 0., 0., 0., 0., 0.],
//             [0., 0., 0., 0., 0., 0., 0., 0.],
//             [0., 0., 0., 0., 0., 0., 0., 0.],
//             [0., 0., 0., 0., 0., 0., 0., 0.],
//             [1., 1., 1., 1., 1., 1., 1., 1.],
//             [1., 1., 1., 1., 1., 1., 1., 1.]])
//
//     >>> b = torch.ones(8, 2)
//     >>> a.slice_scatter(b, dim=1, start=2, end=6, step=2)
//     tensor([[0., 0., 1., 0., 1., 0., 0., 0.],
//             [0., 0., 1., 0., 1., 0., 0., 0.],
//             [0., 0., 1., 0., 1., 0., 0., 0.],
//             [0., 0., 1., 0., 1., 0., 0., 0.],
//             [0., 0., 1., 0., 1., 0., 0., 0.],
//             [0., 0., 1., 0., 1., 0., 0., 0.],
//             [0., 0., 1., 0., 1., 0., 0., 0.],
//             [0., 0., 1., 0., 1., 0., 0., 0.]])
//
//
//go:linkname SliceScatter py.slice_scatter
func SliceScatter(input *py.Object, src *py.Object, dim *py.Object, start *py.Object, end *py.Object, step *py.Object) *py.Object
//
// slogdet(input) -> (Tensor, Tensor)
//
// Alias for :func:`torch.linalg.slogdet`
//
//
//go:linkname Slogdet py.slogdet
func Slogdet(input *py.Object) *py.Object
//
// smm(input, mat) -> Tensor
//
// Performs a matrix multiplication of the sparse matrix :attr:`input`
// with the dense matrix :attr:`mat`.
//
// Args:
//     input (Tensor): a sparse matrix to be matrix multiplied
//     mat (Tensor): a dense matrix to be matrix multiplied
//
//
//go:linkname Smm py.smm
func Smm(input *py.Object, mat *py.Object) *py.Object
//
// softmax(input, dim, *, dtype=None) -> Tensor
//
// Alias for :func:`torch.nn.functional.softmax`.
//
//
//go:linkname Softmax py.softmax
func Softmax(input *py.Object, dim *py.Object) *py.Object
//
// sort(input, dim=-1, descending=False, stable=False, *, out=None) -> (Tensor, LongTensor)
//
// Sorts the elements of the :attr:`input` tensor along a given dimension
// in ascending order by value.
//
// If :attr:`dim` is not given, the last dimension of the `input` is chosen.
//
// If :attr:`descending` is ``True`` then the elements are sorted in descending
// order by value.
//
// If :attr:`stable` is ``True`` then the sorting routine becomes stable, preserving
// the order of equivalent elements.
//
// A namedtuple of (values, indices) is returned, where the `values` are the
// sorted values and `indices` are the indices of the elements in the original
// `input` tensor.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int, optional): the dimension to sort along
//     descending (bool, optional): controls the sorting order (ascending or descending)
//     stable (bool, optional): makes the sorting routine stable, which guarantees that the order
//        of equivalent elements is preserved.
//
// Keyword args:
//     out (tuple, optional): the output tuple of (`Tensor`, `LongTensor`) that can
//         be optionally given to be used as output buffers
//
// Example::
//
//     >>> x = torch.randn(3, 4)
//     >>> sorted, indices = torch.sort(x)
//     >>> sorted
//     tensor([[-0.2162,  0.0608,  0.6719,  2.3332],
//             [-0.5793,  0.0061,  0.6058,  0.9497],
//             [-0.5071,  0.3343,  0.9553,  1.0960]])
//     >>> indices
//     tensor([[ 1,  0,  2,  3],
//             [ 3,  1,  0,  2],
//             [ 0,  3,  1,  2]])
//
//     >>> sorted, indices = torch.sort(x, 0)
//     >>> sorted
//     tensor([[-0.5071, -0.2162,  0.6719, -0.5793],
//             [ 0.0608,  0.0061,  0.9497,  0.3343],
//             [ 0.6058,  0.9553,  1.0960,  2.3332]])
//     >>> indices
//     tensor([[ 2,  0,  0,  1],
//             [ 0,  1,  1,  2],
//             [ 1,  2,  2,  0]])
//     >>> x = torch.tensor([0, 1] * 9)
//     >>> x.sort()
//     torch.return_types.sort(
//         values=tensor([0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1]),
//         indices=tensor([ 2, 16,  4,  6, 14,  8,  0, 10, 12,  9, 17, 15, 13, 11,  7,  5,  3,  1]))
//     >>> x.sort(stable=True)
//     torch.return_types.sort(
//         values=tensor([0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1]),
//         indices=tensor([ 0,  2,  4,  6,  8, 10, 12, 14, 16,  1,  3,  5,  7,  9, 11, 13, 15, 17]))
//
//
//go:linkname Sort py.sort
func Sort(input *py.Object, dim *py.Object, descending *py.Object, stable *py.Object) *py.Object
// sparse_bsc_tensor(ccol_indices, row_indices, values, size=None, *, dtype=None, device=None, requires_grad=False, check_invariants=None) -> Tensor
//
// Constructs a :ref:`sparse tensor in BSC (Block Compressed Sparse
// Column)) <sparse-bsc-docs>` with specified 2-dimensional blocks at the
// given :attr:`ccol_indices` and :attr:`row_indices`. Sparse matrix
// multiplication operations in BSC format are typically faster than that
// for sparse tensors in COO format. Make you have a look at :ref:`the
// note on the data type of the indices <sparse-bsc-docs>`.
//
// .. note::
//
//    If the ``device`` argument is not specified the device of the given
//    :attr:`values` and indices tensor(s) must match. If, however, the
//    argument is specified the input Tensors will be converted to the
//    given device and in turn determine the device of the constructed
//    sparse tensor.
//
// Args:
//     ccol_indices (array_like): (B+1)-dimensional array of size
//         ``(*batchsize, ncolblocks + 1)``. The last element of each
//         batch is the number of non-zeros. This tensor encodes the
//         index in values and row_indices depending on where the given
//         column starts. Each successive number in the tensor subtracted
//         by the number before it denotes the number of elements in a
//         given column.
//     row_indices (array_like): Row block co-ordinates of each block in
//         values. (B+1)-dimensional tensor with the same length
//         as values.
//     values (array_list): Initial blocks for the tensor. Can be a list,
//         tuple, NumPy ``ndarray``, and other types that
//         represents a (1 + 2 + K)-dimensional tensor where ``K`` is the
//         number of dense dimensions.
//     size (list, tuple, :class:`torch.Size`, optional): Size of the
//         sparse tensor: ``(*batchsize, nrows * blocksize[0], ncols *
//         blocksize[1], *densesize)`` If not provided, the size will be
//         inferred as the minimum size big enough to hold all non-zero
//         blocks.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of
//         returned tensor.  Default: if None, infers data type from
//         :attr:`values`.
//     device (:class:`torch.device`, optional): the desired device of
//         returned tensor.  Default: if None, uses the current device
//         for the default tensor type (see
//         :func:`torch.set_default_tensor_type`). :attr:`device` will be
//         the CPU for CPU tensor types and the current CUDA device for
//         CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     check_invariants (bool, optional): If sparse tensor invariants are checked.
//         Default: as returned by :func:`torch.sparse.check_sparse_tensor_invariants.is_enabled`,
//         initially False.
//
// Example::
//     >>> ccol_indices = [0, 1, 2]
//     >>> row_indices = [0, 1]
//     >>> values = [[[1, 2], [3, 4]], [[5, 6], [7, 8]]]
//     >>> torch.sparse_bsc_tensor(torch.tensor(ccol_indices, dtype=torch.int64),
//     ...                         torch.tensor(row_indices, dtype=torch.int64),
//     ...                         torch.tensor(values), dtype=torch.double)
//     tensor(ccol_indices=tensor([0, 1, 2]),
//            row_indices=tensor([0, 1]),
//            values=tensor([[[1., 2.],
//                            [3., 4.]],
//                           [[5., 6.],
//                            [7., 8.]]]), size=(2, 2), nnz=2, dtype=torch.float64,
//            layout=torch.sparse_bsc)
//
//
//go:linkname SparseBscTensor py.sparse_bsc_tensor
func SparseBscTensor(ccolIndices *py.Object, rowIndices *py.Object, values *py.Object, size *py.Object) *py.Object
// sparse_bsr_tensor(crow_indices, col_indices, values, size=None, *, dtype=None, device=None, requires_grad=False, check_invariants=None) -> Tensor
//
// Constructs a :ref:`sparse tensor in BSR (Block Compressed Sparse Row))
// <sparse-bsr-docs>` with specified 2-dimensional blocks at the given
// :attr:`crow_indices` and :attr:`col_indices`. Sparse matrix
// multiplication operations in BSR format are typically faster than that
// for sparse tensors in COO format. Make you have a look at :ref:`the
// note on the data type of the indices <sparse-bsr-docs>`.
//
// .. note::
//
//    If the ``device`` argument is not specified the device of the given
//    :attr:`values` and indices tensor(s) must match. If, however, the
//    argument is specified the input Tensors will be converted to the
//    given device and in turn determine the device of the constructed
//    sparse tensor.
//
// Args:
//     crow_indices (array_like): (B+1)-dimensional array of size
//         ``(*batchsize, nrowblocks + 1)``.  The last element of each
//         batch is the number of non-zeros. This tensor encodes the
//         block index in values and col_indices depending on where the
//         given row block starts. Each successive number in the tensor
//         subtracted by the number before it denotes the number of
//         blocks in a given row.
//     col_indices (array_like): Column block co-ordinates of each block
//         in values. (B+1)-dimensional tensor with the same length as
//         values.
//     values (array_list): Initial values for the tensor. Can be a list,
//         tuple, NumPy ``ndarray``, scalar, and other types that
//         represents a (1 + 2 + K)-dimensional tensor where ``K`` is the
//         number of dense dimensions.
//     size (list, tuple, :class:`torch.Size`, optional): Size of the
//         sparse tensor: ``(*batchsize, nrows * blocksize[0], ncols *
//         blocksize[1], *densesize)`` where ``blocksize ==
//         values.shape[1:3]``. If not provided, the size will be
//         inferred as the minimum size big enough to hold all non-zero
//         blocks.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of
//         returned tensor.  Default: if None, infers data type from
//         :attr:`values`.
//     device (:class:`torch.device`, optional): the desired device of
//         returned tensor.  Default: if None, uses the current device
//         for the default tensor type (see
//         :func:`torch.set_default_tensor_type`). :attr:`device` will be
//         the CPU for CPU tensor types and the current CUDA device for
//         CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     check_invariants (bool, optional): If sparse tensor invariants are checked.
//         Default: as returned by :func:`torch.sparse.check_sparse_tensor_invariants.is_enabled`,
//         initially False.
//
// Example::
//     >>> crow_indices = [0, 1, 2]
//     >>> col_indices = [0, 1]
//     >>> values = [[[1, 2], [3, 4]], [[5, 6], [7, 8]]]
//     >>> torch.sparse_bsr_tensor(torch.tensor(crow_indices, dtype=torch.int64),
//     ...                         torch.tensor(col_indices, dtype=torch.int64),
//     ...                         torch.tensor(values), dtype=torch.double)
//     tensor(crow_indices=tensor([0, 1, 2]),
//            col_indices=tensor([0, 1]),
//            values=tensor([[[1., 2.],
//                            [3., 4.]],
//                           [[5., 6.],
//                            [7., 8.]]]), size=(2, 2), nnz=2, dtype=torch.float64,
//            layout=torch.sparse_bsr)
//
//
//go:linkname SparseBsrTensor py.sparse_bsr_tensor
func SparseBsrTensor(crowIndices *py.Object, colIndices *py.Object, values *py.Object, size *py.Object) *py.Object
// sparse_compressed_tensor(compressed_indices, plain_indices, values, size=None, *, dtype=None, layout=None, device=None, requires_grad=False, check_invariants=None) -> Tensor
//
// Constructs a :ref:`sparse tensor in Compressed Sparse format - CSR,
// CSC, BSR, or BSC - <sparse-compressed-docs>` with specified values at
// the given :attr:`compressed_indices` and :attr:`plain_indices`. Sparse
// matrix multiplication operations in Compressed Sparse format are
// typically faster than that for sparse tensors in COO format. Make you
// have a look at :ref:`the note on the data type of the indices
// <sparse-compressed-docs>`.
//
// .. note::
//
//    If the ``device`` argument is not specified the device of the given
//    :attr:`values` and indices tensor(s) must match. If, however, the
//    argument is specified the input Tensors will be converted to the
//    given device and in turn determine the device of the constructed
//    sparse tensor.
//
// Args:
//     compressed_indices (array_like): (B+1)-dimensional array of size
//         ``(*batchsize, compressed_dim_size + 1)``.  The last element of
//         each batch is the number of non-zero elements or blocks. This
//         tensor encodes the index in ``values`` and ``plain_indices``
//         depending on where the given compressed dimension (row or
//         column) starts. Each successive number in the tensor
//         subtracted by the number before it denotes the number of
//         elements or blocks in a given compressed dimension.
//     plain_indices (array_like): Plain dimension (column or row)
//         co-ordinates of each element or block in values. (B+1)-dimensional
//         tensor with the same length as values.
//
//     values (array_list): Initial values for the tensor. Can be a list,
//         tuple, NumPy ``ndarray``, scalar, and other types.  that
//         represents a (1+K)-dimensional (for CSR and CSC layouts) or
//         (1+2+K)-dimensional tensor (for BSR and BSC layouts) where
//         ``K`` is the number of dense dimensions.
//     size (list, tuple, :class:`torch.Size`, optional): Size of the
//         sparse tensor: ``(*batchsize, nrows * blocksize[0], ncols *
//         blocksize[1], *densesize)`` where ``blocksize[0] ==
//         blocksize[1] == 1`` for CSR and CSC formats. If not provided,
//         the size will be inferred as the minimum size big enough to
//         hold all non-zero elements or blocks.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of
//         returned tensor.  Default: if None, infers data type from
//         :attr:`values`.
//     layout (:class:`torch.layout`, required): the desired layout of
//         returned tensor: :attr:`torch.sparse_csr`,
//         :attr:`torch.sparse_csc`, :attr:`torch.sparse_bsr`, or
//         :attr:`torch.sparse_bsc`.
//     device (:class:`torch.device`, optional): the desired device of
//         returned tensor.  Default: if None, uses the current device
//         for the default tensor type (see
//         :func:`torch.set_default_tensor_type`). :attr:`device` will be
//         the CPU for CPU tensor types and the current CUDA device for
//         CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     check_invariants (bool, optional): If sparse tensor invariants are checked.
//         Default: as returned by :func:`torch.sparse.check_sparse_tensor_invariants.is_enabled`,
//         initially False.
//
// Example::
//     >>> compressed_indices = [0, 2, 4]
//     >>> plain_indices = [0, 1, 0, 1]
//     >>> values = [1, 2, 3, 4]
//     >>> torch.sparse_compressed_tensor(torch.tensor(compressed_indices, dtype=torch.int64),
//     ...                                torch.tensor(plain_indices, dtype=torch.int64),
//     ...                                torch.tensor(values), dtype=torch.double, layout=torch.sparse_csr)
//     tensor(crow_indices=tensor([0, 2, 4]),
//            col_indices=tensor([0, 1, 0, 1]),
//            values=tensor([1., 2., 3., 4.]), size=(2, 2), nnz=4,
//            dtype=torch.float64, layout=torch.sparse_csr)
//
//
//go:linkname SparseCompressedTensor py.sparse_compressed_tensor
func SparseCompressedTensor(compressedIndices *py.Object, plainIndices *py.Object, values *py.Object, size *py.Object) *py.Object
// sparse_coo_tensor(indices, values, size=None, *, dtype=None, device=None, requires_grad=False, check_invariants=None, is_coalesced=None) -> Tensor
//
// Constructs a :ref:`sparse tensor in COO(rdinate) format
// <sparse-coo-docs>` with specified values at the given
// :attr:`indices`.
//
// .. note::
//
//    This function returns an :ref:`uncoalesced tensor
//    <sparse-uncoalesced-coo-docs>` when :attr:`is_coalesced` is
//    unspecified or ``None``.
//
// .. note::
//
//    If the ``device`` argument is not specified the device of the given
//    :attr:`values` and indices tensor(s) must match. If, however, the
//    argument is specified the input Tensors will be converted to the
//    given device and in turn determine the device of the constructed
//    sparse tensor.
//
// Args:
//     indices (array_like): Initial data for the tensor. Can be a list, tuple,
//         NumPy ``ndarray``, scalar, and other types. Will be cast to a :class:`torch.LongTensor`
//         internally. The indices are the coordinates of the non-zero values in the matrix, and thus
//         should be two-dimensional where the first dimension is the number of tensor dimensions and
//         the second dimension is the number of non-zero values.
//     values (array_like): Initial values for the tensor. Can be a list, tuple,
//         NumPy ``ndarray``, scalar, and other types.
//     size (list, tuple, or :class:`torch.Size`, optional): Size of the sparse tensor. If not
//         provided the size will be inferred as the minimum size big enough to hold all non-zero
//         elements.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if None, infers data type from :attr:`values`.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if None, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     check_invariants (bool, optional): If sparse tensor invariants are checked.
//         Default: as returned by :func:`torch.sparse.check_sparse_tensor_invariants.is_enabled`,
//         initially False.
//     is_coalesced (bool, optional): When``True``, the caller is
//         responsible for providing tensor indices that correspond to a
//         coalesced tensor.  If the :attr:`check_invariants` flag is
//         False, no error will be raised if the prerequisites are not
//         met and this will lead to silently incorrect results. To force
//         coalescion please use :meth:`coalesce` on the resulting
//         Tensor.
//         Default: None: except for trivial cases (e.g. nnz < 2) the
//         resulting Tensor has is_coalesced set to ``False```.
//
// Example::
//
//     >>> i = torch.tensor([[0, 1, 1],
//     ...                   [2, 0, 2]])
//     >>> v = torch.tensor([3, 4, 5], dtype=torch.float32)
//     >>> torch.sparse_coo_tensor(i, v, [2, 4])
//     tensor(indices=tensor([[0, 1, 1],
//                            [2, 0, 2]]),
//            values=tensor([3., 4., 5.]),
//            size=(2, 4), nnz=3, layout=torch.sparse_coo)
//
//     >>> torch.sparse_coo_tensor(i, v)  # Shape inference
//     tensor(indices=tensor([[0, 1, 1],
//                            [2, 0, 2]]),
//            values=tensor([3., 4., 5.]),
//            size=(2, 3), nnz=3, layout=torch.sparse_coo)
//
//     >>> torch.sparse_coo_tensor(i, v, [2, 4],
//     ...                         dtype=torch.float64,
//     ...                         device=torch.device('cuda:0'))
//     tensor(indices=tensor([[0, 1, 1],
//                            [2, 0, 2]]),
//            values=tensor([3., 4., 5.]),
//            device='cuda:0', size=(2, 4), nnz=3, dtype=torch.float64,
//            layout=torch.sparse_coo)
//
//     # Create an empty sparse tensor with the following invariants:
//     #   1. sparse_dim + dense_dim = len(SparseTensor.shape)
//     #   2. SparseTensor._indices().shape = (sparse_dim, nnz)
//     #   3. SparseTensor._values().shape = (nnz, SparseTensor.shape[sparse_dim:])
//     #
//     # For instance, to create an empty sparse tensor with nnz = 0, dense_dim = 0 and
//     # sparse_dim = 1 (hence indices is a 2D tensor of shape = (1, 0))
//     >>> S = torch.sparse_coo_tensor(torch.empty([1, 0]), [], [1])
//     tensor(indices=tensor([], size=(1, 0)),
//            values=tensor([], size=(0,)),
//            size=(1,), nnz=0, layout=torch.sparse_coo)
//
//     # and to create an empty sparse tensor with nnz = 0, dense_dim = 1 and
//     # sparse_dim = 1
//     >>> S = torch.sparse_coo_tensor(torch.empty([1, 0]), torch.empty([0, 2]), [1, 2])
//     tensor(indices=tensor([], size=(1, 0)),
//            values=tensor([], size=(0, 2)),
//            size=(1, 2), nnz=0, layout=torch.sparse_coo)
//
// .. _torch.sparse: https://pytorch.org/docs/stable/sparse.html
//
//
//go:linkname SparseCooTensor py.sparse_coo_tensor
func SparseCooTensor(indices *py.Object, values *py.Object, size *py.Object) *py.Object
// sparse_csc_tensor(ccol_indices, row_indices, values, size=None, *, dtype=None, device=None, requires_grad=False, check_invariants=None) -> Tensor
//
// Constructs a :ref:`sparse tensor in CSC (Compressed Sparse Column)
// <sparse-csc-docs>` with specified values at the given
// :attr:`ccol_indices` and :attr:`row_indices`. Sparse matrix
// multiplication operations in CSC format are typically faster than that
// for sparse tensors in COO format. Make you have a look at :ref:`the
// note on the data type of the indices <sparse-csc-docs>`.
//
// .. note::
//
//    If the ``device`` argument is not specified the device of the given
//    :attr:`values` and indices tensor(s) must match. If, however, the
//    argument is specified the input Tensors will be converted to the
//    given device and in turn determine the device of the constructed
//    sparse tensor.
//
// Args:
//     ccol_indices (array_like): (B+1)-dimensional array of size
//         ``(*batchsize, ncols + 1)``.  The last element of each batch
//         is the number of non-zeros. This tensor encodes the index in
//         values and row_indices depending on where the given column
//         starts. Each successive number in the tensor subtracted by the
//         number before it denotes the number of elements in a given
//         column.
//     row_indices (array_like): Row co-ordinates of each element in
//         values. (B+1)-dimensional tensor with the same length as
//         values.
//     values (array_list): Initial values for the tensor. Can be a list,
//         tuple, NumPy ``ndarray``, scalar, and other types that
//         represents a (1+K)-dimensional tensor where ``K`` is the number
//         of dense dimensions.
//     size (list, tuple, :class:`torch.Size`, optional): Size of the
//         sparse tensor: ``(*batchsize, nrows, ncols, *densesize)``. If
//         not provided, the size will be inferred as the minimum size
//         big enough to hold all non-zero elements.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of
//         returned tensor.  Default: if None, infers data type from
//         :attr:`values`.
//     device (:class:`torch.device`, optional): the desired device of
//         returned tensor.  Default: if None, uses the current device
//         for the default tensor type (see
//         :func:`torch.set_default_tensor_type`). :attr:`device` will be
//         the CPU for CPU tensor types and the current CUDA device for
//         CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     check_invariants (bool, optional): If sparse tensor invariants are checked.
//         Default: as returned by :func:`torch.sparse.check_sparse_tensor_invariants.is_enabled`,
//         initially False.
//
// Example::
//     >>> ccol_indices = [0, 2, 4]
//     >>> row_indices = [0, 1, 0, 1]
//     >>> values = [1, 2, 3, 4]
//     >>> torch.sparse_csc_tensor(torch.tensor(ccol_indices, dtype=torch.int64),
//     ...                         torch.tensor(row_indices, dtype=torch.int64),
//     ...                         torch.tensor(values), dtype=torch.double)
//     tensor(ccol_indices=tensor([0, 2, 4]),
//            row_indices=tensor([0, 1, 0, 1]),
//            values=tensor([1., 2., 3., 4.]), size=(2, 2), nnz=4,
//            dtype=torch.float64, layout=torch.sparse_csc)
//
//
//go:linkname SparseCscTensor py.sparse_csc_tensor
func SparseCscTensor(ccolIndices *py.Object, rowIndices *py.Object, values *py.Object, size *py.Object) *py.Object
// sparse_csr_tensor(crow_indices, col_indices, values, size=None, *, dtype=None, device=None, requires_grad=False, check_invariants=None) -> Tensor
//
// Constructs a :ref:`sparse tensor in CSR (Compressed Sparse Row) <sparse-csr-docs>` with specified
// values at the given :attr:`crow_indices` and :attr:`col_indices`. Sparse matrix multiplication operations
// in CSR format are typically faster than that for sparse tensors in COO format. Make you have a look
// at :ref:`the note on the data type of the indices <sparse-csr-docs>`.
//
// .. note::
//
//    If the ``device`` argument is not specified the device of the given
//    :attr:`values` and indices tensor(s) must match. If, however, the
//    argument is specified the input Tensors will be converted to the
//    given device and in turn determine the device of the constructed
//    sparse tensor.
//
// Args:
//     crow_indices (array_like): (B+1)-dimensional array of size
//         ``(*batchsize, nrows + 1)``.  The last element of each batch
//         is the number of non-zeros. This tensor encodes the index in
//         values and col_indices depending on where the given row
//         starts. Each successive number in the tensor subtracted by the
//         number before it denotes the number of elements in a given
//         row.
//     col_indices (array_like): Column co-ordinates of each element in
//         values. (B+1)-dimensional tensor with the same length
//         as values.
//     values (array_list): Initial values for the tensor. Can be a list,
//         tuple, NumPy ``ndarray``, scalar, and other types that
//         represents a (1+K)-dimensional tensor where ``K`` is the number
//         of dense dimensions.
//     size (list, tuple, :class:`torch.Size`, optional): Size of the
//         sparse tensor: ``(*batchsize, nrows, ncols, *densesize)``. If
//         not provided, the size will be inferred as the minimum size
//         big enough to hold all non-zero elements.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of
//         returned tensor.  Default: if None, infers data type from
//         :attr:`values`.
//     device (:class:`torch.device`, optional): the desired device of
//         returned tensor.  Default: if None, uses the current device
//         for the default tensor type (see
//         :func:`torch.set_default_tensor_type`). :attr:`device` will be
//         the CPU for CPU tensor types and the current CUDA device for
//         CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     check_invariants (bool, optional): If sparse tensor invariants are checked.
//         Default: as returned by :func:`torch.sparse.check_sparse_tensor_invariants.is_enabled`,
//         initially False.
//
// Example::
//     >>> crow_indices = [0, 2, 4]
//     >>> col_indices = [0, 1, 0, 1]
//     >>> values = [1, 2, 3, 4]
//     >>> torch.sparse_csr_tensor(torch.tensor(crow_indices, dtype=torch.int64),
//     ...                         torch.tensor(col_indices, dtype=torch.int64),
//     ...                         torch.tensor(values), dtype=torch.double)
//     tensor(crow_indices=tensor([0, 2, 4]),
//            col_indices=tensor([0, 1, 0, 1]),
//            values=tensor([1., 2., 3., 4.]), size=(2, 2), nnz=4,
//            dtype=torch.float64, layout=torch.sparse_csr)
//
//
//go:linkname SparseCsrTensor py.sparse_csr_tensor
func SparseCsrTensor(crowIndices *py.Object, colIndices *py.Object, values *py.Object, size *py.Object) *py.Object
// Splits the tensor into chunks. Each chunk is a view of the original tensor.
//
//     If :attr:`split_size_or_sections` is an integer type, then :attr:`tensor` will
//     be split into equally sized chunks (if possible). Last chunk will be smaller if
//     the tensor size along the given dimension :attr:`dim` is not divisible by
//     :attr:`split_size`.
//
//     If :attr:`split_size_or_sections` is a list, then :attr:`tensor` will be split
//     into ``len(split_size_or_sections)`` chunks with sizes in :attr:`dim` according
//     to :attr:`split_size_or_sections`.
//
//     Args:
//         tensor (Tensor): tensor to split.
//         split_size_or_sections (int) or (list(int)): size of a single chunk or
//             list of sizes for each chunk
//         dim (int): dimension along which to split the tensor.
//
//     Example::
//
//         >>> a = torch.arange(10).reshape(5, 2)
//         >>> a
//         tensor([[0, 1],
//                 [2, 3],
//                 [4, 5],
//                 [6, 7],
//                 [8, 9]])
//         >>> torch.split(a, 2)
//         (tensor([[0, 1],
//                  [2, 3]]),
//          tensor([[4, 5],
//                  [6, 7]]),
//          tensor([[8, 9]]))
//         >>> torch.split(a, [1, 4])
//         (tensor([[0, 1]]),
//          tensor([[2, 3],
//                  [4, 5],
//                  [6, 7],
//                  [8, 9]]))
//
//
//go:linkname Split py.split
func Split(tensor *py.Object, splitSizeOrSections *py.Object, dim *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.split`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname SplitCopy py.split_copy
func SplitCopy(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname SplitWithSizes py.split_with_sizes
func SplitWithSizes(__llgo_va_list ...interface{}) *py.Object
//
// Performs the same operation as :func:`torch.split_with_sizes`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname SplitWithSizesCopy py.split_with_sizes_copy
func SplitWithSizesCopy(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Spmm py.spmm
func Spmm(__llgo_va_list ...interface{}) *py.Object
//
// sqrt(input, *, out=None) -> Tensor
//
// Returns a new tensor with the square-root of the elements of :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \sqrt{\text{input}_{i}}
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([-2.0755,  1.0226,  0.0831,  0.4806])
//     >>> torch.sqrt(a)
//     tensor([    nan,  1.0112,  0.2883,  0.6933])
//
//
//go:linkname Sqrt py.sqrt
func Sqrt(input *py.Object) *py.Object
// None
//
//go:linkname Sqrt_ py.sqrt_
func Sqrt_(__llgo_va_list ...interface{}) *py.Object
//
// square(input, *, out=None) -> Tensor
//
// Returns a new tensor with the square of the elements of :attr:`input`.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([-2.0755,  1.0226,  0.0831,  0.4806])
//     >>> torch.square(a)
//     tensor([ 4.3077,  1.0457,  0.0069,  0.2310])
//
//
//go:linkname Square py.square
func Square(input *py.Object) *py.Object
// None
//
//go:linkname Square_ py.square_
func Square_(__llgo_va_list ...interface{}) *py.Object
//
// squeeze(input, dim=None) -> Tensor
//
// Returns a tensor with all specified dimensions of :attr:`input` of size `1` removed.
//
// For example, if `input` is of shape:
// :math:`(A \times 1 \times B \times C \times 1 \times D)` then the `input.squeeze()`
// will be of shape: :math:`(A \times B \times C \times D)`.
//
// When :attr:`dim` is given, a squeeze operation is done only in the given
// dimension(s). If `input` is of shape: :math:`(A \times 1 \times B)`,
// ``squeeze(input, 0)`` leaves the tensor unchanged, but ``squeeze(input, 1)``
// will squeeze the tensor to the shape :math:`(A \times B)`.
//
// .. note:: The returned tensor shares the storage with the input tensor,
//           so changing the contents of one will change the contents of the other.
//
// .. warning:: If the tensor has a batch dimension of size 1, then `squeeze(input)`
//           will also remove the batch dimension, which can lead to unexpected
//           errors. Consider specifying only the dims you wish to be squeezed.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int or tuple of ints, optional): if given, the input will be squeezed
//            only in the specified dimensions.
//
//         .. versionchanged:: 2.0
//            :attr:`dim` now accepts tuples of dimensions.
//
// Example::
//
//     >>> x = torch.zeros(2, 1, 2, 1, 2)
//     >>> x.size()
//     torch.Size([2, 1, 2, 1, 2])
//     >>> y = torch.squeeze(x)
//     >>> y.size()
//     torch.Size([2, 2, 2])
//     >>> y = torch.squeeze(x, 0)
//     >>> y.size()
//     torch.Size([2, 1, 2, 1, 2])
//     >>> y = torch.squeeze(x, 1)
//     >>> y.size()
//     torch.Size([2, 2, 1, 2])
//     >>> y = torch.squeeze(x, (1, 2, 3))
//     torch.Size([2, 2, 2])
//
//
//go:linkname Squeeze py.squeeze
func Squeeze(input *py.Object, dim *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.squeeze`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname SqueezeCopy py.squeeze_copy
func SqueezeCopy(__llgo_va_list ...interface{}) *py.Object
//
// sspaddmm(input, mat1, mat2, *, beta=1, alpha=1, out=None) -> Tensor
//
// Matrix multiplies a sparse tensor :attr:`mat1` with a dense tensor
// :attr:`mat2`, then adds the sparse tensor :attr:`input` to the result.
//
// Note: This function is equivalent to :func:`torch.addmm`, except
// :attr:`input` and :attr:`mat1` are sparse.
//
// Args:
//     input (Tensor): a sparse matrix to be added
//     mat1 (Tensor): a sparse matrix to be matrix multiplied
//     mat2 (Tensor): a dense matrix to be matrix multiplied
//
// Keyword args:
//     beta (Number, optional): multiplier for :attr:`mat` (:math:`\beta`)
//     alpha (Number, optional): multiplier for :math:`mat1 @ mat2` (:math:`\alpha`)
//     out (Tensor, optional): the output tensor.
//
//
//go:linkname Sspaddmm py.sspaddmm
func Sspaddmm(input *py.Object, mat1 *py.Object, mat2 *py.Object) *py.Object
//
// stack(tensors, dim=0, *, out=None) -> Tensor
//
// Concatenates a sequence of tensors along a new dimension.
//
// All tensors need to be of the same size.
//
// .. seealso::
//
//     :func:`torch.cat` concatenates the given sequence along an existing dimension.
//
// Arguments:
//     tensors (sequence of Tensors): sequence of tensors to concatenate
//     dim (int): dimension to insert. Has to be between 0 and the number
//         of dimensions of concatenated tensors (inclusive)
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
//
//go:linkname Stack py.stack
func Stack(tensors *py.Object, dim *py.Object) *py.Object
//
// std(input, dim=None, *, correction=1, keepdim=False, out=None) -> Tensor
//
// Calculates the standard deviation over the dimensions specified by :attr:`dim`.
// :attr:`dim` can be a single dimension, list of dimensions, or ``None`` to
// reduce over all dimensions.
//
// The standard deviation (:math:`\sigma`) is calculated as
//
// .. math:: \sigma = \sqrt{\frac{1}{\max(0,~N - \delta N)}\sum_{i=0}^{N-1}(x_i-\bar{x})^2}
//
// where :math:`x` is the sample set of elements, :math:`\bar{x}` is the
// sample mean, :math:`N` is the number of samples and :math:`\delta N` is
// the :attr:`correction`.
//
//
//
// If :attr:`keepdim` is ``True``, the output tensor is of the same size
// as :attr:`input` except in the dimension(s) :attr:`dim` where it is of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in the
// output tensor having 1 (or ``len(dim)``) fewer dimension(s).
//
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int or tuple of ints): the dimension or dimensions to reduce.
//
// Keyword args:
//     correction (int): difference between the sample size and sample degrees of freedom.
//         Defaults to `Bessel's correction`_, ``correction=1``.
//
//         .. versionchanged:: 2.0
//             Previously this argument was called ``unbiased`` and was a boolean
//             with ``True`` corresponding to ``correction=1`` and ``False`` being
//             ``correction=0``.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//     out (Tensor, optional): the output tensor.
//
// Example:
//
//     >>> a = torch.tensor(
//     ...     [[ 0.2035,  1.2959,  1.8101, -0.4644],
//     ...      [ 1.5027, -0.3270,  0.5905,  0.6538],
//     ...      [-1.5745,  1.3330, -0.5596, -0.6548],
//     ...      [ 0.1264, -0.5080,  1.6420,  0.1992]])
//     >>> torch.std(a, dim=1, keepdim=True)
//     tensor([[1.0311],
//             [0.7477],
//             [1.2204],
//             [0.9087]])
//
// .. _Bessel's correction: https://en.wikipedia.org/wiki/Bessel%27s_correction
//
//
//
//go:linkname Std py.std
func Std(input *py.Object, dim *py.Object) *py.Object
//
// std_mean(input, dim=None, *, correction=1, keepdim=False, out=None) -> (Tensor, Tensor)
//
// Calculates the standard deviation and mean over the dimensions specified by
// :attr:`dim`. :attr:`dim` can be a single dimension, list of dimensions, or
// ``None`` to reduce over all dimensions.
//
// The standard deviation (:math:`\sigma`) is calculated as
//
// .. math:: \sigma = \sqrt{\frac{1}{\max(0,~N - \delta N)}\sum_{i=0}^{N-1}(x_i-\bar{x})^2}
//
// where :math:`x` is the sample set of elements, :math:`\bar{x}` is the
// sample mean, :math:`N` is the number of samples and :math:`\delta N` is
// the :attr:`correction`.
//
//
//
//
// If :attr:`keepdim` is ``True``, the output tensor is of the same size
// as :attr:`input` except in the dimension(s) :attr:`dim` where it is of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in the
// output tensor having 1 (or ``len(dim)``) fewer dimension(s).
//
//
// Args:
//     input (Tensor): the input tensor.
//
//     dim (int or tuple of ints, optional): the dimension or dimensions to reduce.
//         If ``None``, all dimensions are reduced.
//
//
// Keyword args:
//     correction (int): difference between the sample size and sample degrees of freedom.
//         Defaults to `Bessel's correction`_, ``correction=1``.
//
//         .. versionchanged:: 2.0
//             Previously this argument was called ``unbiased`` and was a boolean
//             with ``True`` corresponding to ``correction=1`` and ``False`` being
//             ``correction=0``.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//     out (Tensor, optional): the output tensor.
//
// Returns:
//     A tuple (std, mean) containing the standard deviation and mean.
//
// Example:
//
//     >>> a = torch.tensor(
//     ...     [[ 0.2035,  1.2959,  1.8101, -0.4644],
//     ...      [ 1.5027, -0.3270,  0.5905,  0.6538],
//     ...      [-1.5745,  1.3330, -0.5596, -0.6548],
//     ...      [ 0.1264, -0.5080,  1.6420,  0.1992]])
//     >>> torch.std_mean(a, dim=0, keepdim=True)
//     (tensor([[1.2620, 1.0028, 1.0957, 0.6038]]),
//      tensor([[ 0.0645,  0.4485,  0.8707, -0.0665]]))
//
// .. _Bessel's correction: https://en.wikipedia.org/wiki/Bessel%27s_correction
//
//
//
//go:linkname StdMean py.std_mean
func StdMean(input *py.Object, dim *py.Object) *py.Object
// Short-time Fourier transform (STFT).
//
//     .. warning::
//         From version 1.8.0, :attr:`return_complex` must always be given
//         explicitly for real inputs and `return_complex=False` has been
//         deprecated. Strongly prefer `return_complex=True` as in a future
//         pytorch release, this function will only return complex tensors.
//
//         Note that :func:`torch.view_as_real` can be used to recover a real
//         tensor with an extra last dimension for real and imaginary components.
//
//     .. warning::
//         From version 2.1, a warning will be provided if a :attr:`window` is
//         not specified. In a future release, this attribute will be required.
//         Not providing a window currently defaults to using a rectangular window,
//         which may result in undesirable artifacts. Consider using tapered windows,
//         such as :func:`torch.hann_window`.
//
//     The STFT computes the Fourier transform of short overlapping windows of the
//     input. This giving frequency components of the signal as they change over
//     time. The interface of this function is modeled after (but *not* a drop-in
//     replacement for) librosa_ stft function.
//
//     .. _librosa: https://librosa.org/doc/latest/generated/librosa.stft.html
//
//     Ignoring the optional batch dimension, this method computes the following
//     expression:
//
//     .. math::
//         X[\omega, m] = \sum_{k = 0}^{\text{win\_length-1}}%
//                             \text{window}[k]\ \text{input}[m \times \text{hop\_length} + k]\ %
//                             \exp\left(- j \frac{2 \pi \cdot \omega k}{\text{n\_fft}}\right),
//
//     where :math:`m` is the index of the sliding window, and :math:`\omega` is
//     the frequency :math:`0 \leq \omega < \text{n\_fft}` for ``onesided=False``,
//     or :math:`0 \leq \omega < \lfloor \text{n\_fft} / 2 \rfloor + 1` for ``onesided=True``.
//
//     * :attr:`input` must be either a 1-D time sequence or a 2-D batch of time
//       sequences.
//
//     * If :attr:`hop_length` is ``None`` (default), it is treated as equal to
//       ``floor(n_fft / 4)``.
//
//     * If :attr:`win_length` is ``None`` (default), it is treated as equal to
//       :attr:`n_fft`.
//
//     * :attr:`window` can be a 1-D tensor of size :attr:`win_length`, e.g., from
//       :meth:`torch.hann_window`. If :attr:`window` is ``None`` (default), it is
//       treated as if having :math:`1` everywhere in the window. If
//       :math:`\text{win\_length} < \text{n\_fft}`, :attr:`window` will be padded on
//       both sides to length :attr:`n_fft` before being applied.
//
//     * If :attr:`center` is ``True`` (default), :attr:`input` will be padded on
//       both sides so that the :math:`t`-th frame is centered at time
//       :math:`t \times \text{hop\_length}`. Otherwise, the :math:`t`-th frame
//       begins at time  :math:`t \times \text{hop\_length}`.
//
//     * :attr:`pad_mode` determines the padding method used on :attr:`input` when
//       :attr:`center` is ``True``. See :meth:`torch.nn.functional.pad` for
//       all available options. Default is ``"reflect"``.
//
//     * If :attr:`onesided` is ``True`` (default for real input), only values for
//       :math:`\omega` in :math:`\left[0, 1, 2, \dots, \left\lfloor
//       \frac{\text{n\_fft}}{2} \right\rfloor + 1\right]` are returned because
//       the real-to-complex Fourier transform satisfies the conjugate symmetry,
//       i.e., :math:`X[m, \omega] = X[m, \text{n\_fft} - \omega]^*`.
//       Note if the input or window tensors are complex, then :attr:`onesided`
//       output is not possible.
//
//     * If :attr:`normalized` is ``True`` (default is ``False``), the function
//       returns the normalized STFT results, i.e., multiplied by :math:`(\text{frame\_length})^{-0.5}`.
//
//     * If :attr:`return_complex` is ``True`` (default if input is complex), the
//       return is a ``input.dim() + 1`` dimensional complex tensor. If ``False``,
//       the output is a ``input.dim() + 2`` dimensional real tensor where the last
//       dimension represents the real and imaginary components.
//
//     Returns either a complex tensor of size :math:`(* \times N \times T)` if
//     :attr:`return_complex` is true, or a real tensor of size :math:`(* \times N
//     \times T \times 2)`. Where :math:`*` is the optional batch size of
//     :attr:`input`, :math:`N` is the number of frequencies where STFT is applied
//     and :math:`T` is the total number of frames used.
//
//     .. warning::
//       This function changed signature at version 0.4.1. Calling with the
//       previous signature may cause error or return incorrect result.
//
//     Args:
//         input (Tensor): the input tensor of shape `(B?, L)` where `B?` is an optional
//             batch dimension
//         n_fft (int): size of Fourier transform
//         hop_length (int, optional): the distance between neighboring sliding window
//             frames. Default: ``None`` (treated as equal to ``floor(n_fft / 4)``)
//         win_length (int, optional): the size of window frame and STFT filter.
//             Default: ``None``  (treated as equal to :attr:`n_fft`)
//         window (Tensor, optional): the optional window function.
//             Shape must be 1d and `<= n_fft`
//             Default: ``None`` (treated as window of all :math:`1` s)
//         center (bool, optional): whether to pad :attr:`input` on both sides so
//             that the :math:`t`-th frame is centered at time :math:`t \times \text{hop\_length}`.
//             Default: ``True``
//         pad_mode (str, optional): controls the padding method used when
//             :attr:`center` is ``True``. Default: ``"reflect"``
//         normalized (bool, optional): controls whether to return the normalized STFT results
//              Default: ``False``
//         onesided (bool, optional): controls whether to return half of results to
//             avoid redundancy for real inputs.
//             Default: ``True`` for real :attr:`input` and :attr:`window`, ``False`` otherwise.
//         return_complex (bool, optional): whether to return a complex tensor, or
//             a real tensor with an extra last dimension for the real and
//             imaginary components.
//
//             .. versionchanged:: 2.0
//                ``return_complex`` is now a required argument for real inputs,
//                as the default is being transitioned to ``True``.
//
//             .. deprecated:: 2.0
//                ``return_complex=False`` is deprecated, instead use ``return_complex=True``
//                Note that calling :func:`torch.view_as_real` on the output will
//                recover the deprecated output format.
//
//     Returns:
//         Tensor: A tensor containing the STFT result with shape `(B?, N, T, C?)` where
//            - `B?` is an optional batch dimnsion from the input
//            - `N` is the number of frequency samples, `(n_fft // 2) + 1` for
//              `onesided=True`, or otherwise `n_fft`.
//            - `T` is the number of frames, `1 + L // hop_length`
//              for `center=True`, or `1 + (L - n_fft) // hop_length` otherwise.
//            - `C?` is an optional length-2 dimension of real and imaginary
//              components, present when `return_complex=False`.
//
//
//
//go:linkname Stft py.stft
func Stft(input *py.Object, nFft *py.Object, hopLength *py.Object, winLength *py.Object, window *py.Object, center *py.Object, padMode *py.Object, normalized *py.Object, onesided *py.Object, returnComplex *py.Object) *py.Object
//
// sub(input, other, *, alpha=1, out=None) -> Tensor
//
// Subtracts :attr:`other`, scaled by :attr:`alpha`, from :attr:`input`.
//
// .. math::
//     \text{{out}}_i = \text{{input}}_i - \text{{alpha}} \times \text{{other}}_i
//
//
// Supports :ref:`broadcasting to a common shape <broadcasting-semantics>`,
// :ref:`type promotion <type-promotion-doc>`, and integer, float, and complex inputs.
//
// Args:
//     input (Tensor): the input tensor.
//     other (Tensor or Number): the tensor or number to subtract from :attr:`input`.
//
// Keyword args:
//     alpha (Number): the multiplier for :attr:`other`.
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor((1, 2))
//     >>> b = torch.tensor((0, 1))
//     >>> torch.sub(a, b, alpha=2)
//     tensor([1, 0])
//
//
//go:linkname Sub py.sub
func Sub(input *py.Object, other *py.Object) *py.Object
//
// subtract(input, other, *, alpha=1, out=None) -> Tensor
//
// Alias for :func:`torch.sub`.
//
//
//go:linkname Subtract py.subtract
func Subtract(input *py.Object, other *py.Object) *py.Object
//
// sum(input, *, dtype=None) -> Tensor
//
// Returns the sum of all elements in the :attr:`input` tensor.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         If specified, the input tensor is casted to :attr:`dtype` before the operation
//         is performed. This is useful for preventing data type overflows. Default: None.
//
// Example::
//
//     >>> a = torch.randn(1, 3)
//     >>> a
//     tensor([[ 0.1133, -0.9567,  0.2958]])
//     >>> torch.sum(a)
//     tensor(-0.5475)
//
// .. function:: sum(input, dim, keepdim=False, *, dtype=None) -> Tensor
//    :noindex:
//
// Returns the sum of each row of the :attr:`input` tensor in the given
// dimension :attr:`dim`. If :attr:`dim` is a list of dimensions,
// reduce over all of them.
//
//
// If :attr:`keepdim` is ``True``, the output tensor is of the same size
// as :attr:`input` except in the dimension(s) :attr:`dim` where it is of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in the
// output tensor having 1 (or ``len(dim)``) fewer dimension(s).
//
//
// Args:
//     input (Tensor): the input tensor.
//
//     dim (int or tuple of ints, optional): the dimension or dimensions to reduce.
//         If ``None``, all dimensions are reduced.
//
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         If specified, the input tensor is casted to :attr:`dtype` before the operation
//         is performed. This is useful for preventing data type overflows. Default: None.
//
// Example::
//
//     >>> a = torch.randn(4, 4)
//     >>> a
//     tensor([[ 0.0569, -0.2475,  0.0737, -0.3429],
//             [-0.2993,  0.9138,  0.9337, -1.6864],
//             [ 0.1132,  0.7892, -0.1003,  0.5688],
//             [ 0.3637, -0.9906, -0.4752, -1.5197]])
//     >>> torch.sum(a, 1)
//     tensor([-0.4598, -0.1381,  1.3708, -2.6217])
//     >>> b = torch.arange(4 * 5 * 6).view(4, 5, 6)
//     >>> torch.sum(b, (2, 1))
//     tensor([  435.,  1335.,  2235.,  3135.])
//
//
//go:linkname Sum py.sum
func Sum(input *py.Object) *py.Object
//
// svd(input, some=True, compute_uv=True, *, out=None) -> (Tensor, Tensor, Tensor)
//
// Computes the singular value decomposition of either a matrix or batch of
// matrices :attr:`input`. The singular value decomposition is represented as a
// namedtuple `(U, S, V)`, such that :attr:`input` :math:`= U \text{diag}(S) V^{\text{H}}`.
// where :math:`V^{\text{H}}` is the transpose of `V` for real inputs,
// and the conjugate transpose of `V` for complex inputs.
// If :attr:`input` is a batch of matrices, then `U`, `S`, and `V` are also
// batched with the same batch dimensions as :attr:`input`.
//
// If :attr:`some` is `True` (default), the method returns the reduced singular
// value decomposition. In this case, if the last two dimensions of :attr:`input` are
// `m` and `n`, then the returned `U` and `V` matrices will contain only
// `min(n, m)` orthonormal columns.
//
// If :attr:`compute_uv` is `False`, the returned `U` and `V` will be
// zero-filled matrices of shape `(m, m)` and `(n, n)`
// respectively, and the same device as :attr:`input`. The argument :attr:`some`
// has no effect when :attr:`compute_uv` is `False`.
//
// Supports :attr:`input` of float, double, cfloat and cdouble data types.
// The dtypes of `U` and `V` are the same as :attr:`input`'s. `S` will
// always be real-valued, even if :attr:`input` is complex.
//
// .. warning::
//
//     :func:`torch.svd` is deprecated in favor of :func:`torch.linalg.svd`
//     and will be removed in a future PyTorch release.
//
//     ``U, S, V = torch.svd(A, some=some, compute_uv=True)`` (default) should be replaced with
//
//     .. code:: python
//
//         U, S, Vh = torch.linalg.svd(A, full_matrices=not some)
//         V = Vh.mH
//
//     ``_, S, _ = torch.svd(A, some=some, compute_uv=False)`` should be replaced with
//
//     .. code:: python
//
//         S = torch.linalg.svdvals(A)
//
// .. note:: Differences with :func:`torch.linalg.svd`:
//
//              * :attr:`some` is the opposite of
//                :func:`torch.linalg.svd`'s :attr:`full_matrices`. Note that
//                default value for both is `True`, so the default behavior is
//                effectively the opposite.
//              * :func:`torch.svd` returns `V`, whereas :func:`torch.linalg.svd` returns
//                `Vh`, that is, :math:`V^{\text{H}}`.
//              * If :attr:`compute_uv` is `False`, :func:`torch.svd` returns zero-filled
//                tensors for `U` and `Vh`, whereas :func:`torch.linalg.svd` returns
//                empty tensors.
//
// .. note:: The singular values are returned in descending order. If :attr:`input` is a batch of matrices,
//           then the singular values of each matrix in the batch are returned in descending order.
//
// .. note:: The `S` tensor can only be used to compute gradients if :attr:`compute_uv` is `True`.
//
// .. note:: When :attr:`some` is `False`, the gradients on `U[..., :, min(m, n):]`
//           and `V[..., :, min(m, n):]` will be ignored in the backward pass, as those vectors
//           can be arbitrary bases of the corresponding subspaces.
//
// .. note:: The implementation of :func:`torch.linalg.svd` on CPU uses LAPACK's routine `?gesdd`
//           (a divide-and-conquer algorithm) instead of `?gesvd` for speed. Analogously,
//           on GPU, it uses cuSOLVER's routines `gesvdj` and `gesvdjBatched` on CUDA 10.1.243
//           and later, and MAGMA's routine `gesdd` on earlier versions of CUDA.
//
// .. note:: The returned `U` will not be contiguous. The matrix (or batch of matrices) will
//           be represented as a column-major matrix (i.e. Fortran-contiguous).
//
// .. warning:: The gradients with respect to `U` and `V` will only be finite when the input does not
//              have zero nor repeated singular values.
//
// .. warning:: If the distance between any two singular values is close to zero, the gradients with respect to
//              `U` and `V` will be numerically unstable, as they depends on
//              :math:`\frac{1}{\min_{i \neq j} \sigma_i^2 - \sigma_j^2}`. The same happens when the matrix
//              has small singular values, as these gradients also depend on `S⁻¹`.
//
// .. warning:: For complex-valued :attr:`input` the singular value decomposition is not unique,
//              as `U` and `V` may be multiplied by an arbitrary phase factor :math:`e^{i \phi}` on every column.
//              The same happens when :attr:`input` has repeated singular values, where one may multiply
//              the columns of the spanning subspace in `U` and `V` by a rotation matrix
//              and `the resulting vectors will span the same subspace`_.
//              Different platforms, like NumPy, or inputs on different device types,
//              may produce different `U` and `V` tensors.
//
// Args:
//     input (Tensor): the input tensor of size `(*, m, n)` where `*` is zero or more
//                     batch dimensions consisting of `(m, n)` matrices.
//     some (bool, optional): controls whether to compute the reduced or full decomposition, and
//                            consequently, the shape of returned `U` and `V`. Default: `True`.
//     compute_uv (bool, optional): controls whether to compute `U` and `V`. Default: `True`.
//
// Keyword args:
//     out (tuple, optional): the output tuple of tensors
//
// Example::
//
//     >>> a = torch.randn(5, 3)
//     >>> a
//     tensor([[ 0.2364, -0.7752,  0.6372],
//             [ 1.7201,  0.7394, -0.0504],
//             [-0.3371, -1.0584,  0.5296],
//             [ 0.3550, -0.4022,  1.5569],
//             [ 0.2445, -0.0158,  1.1414]])
//     >>> u, s, v = torch.svd(a)
//     >>> u
//     tensor([[ 0.4027,  0.0287,  0.5434],
//             [-0.1946,  0.8833,  0.3679],
//             [ 0.4296, -0.2890,  0.5261],
//             [ 0.6604,  0.2717, -0.2618],
//             [ 0.4234,  0.2481, -0.4733]])
//     >>> s
//     tensor([2.3289, 2.0315, 0.7806])
//     >>> v
//     tensor([[-0.0199,  0.8766,  0.4809],
//             [-0.5080,  0.4054, -0.7600],
//             [ 0.8611,  0.2594, -0.4373]])
//     >>> torch.dist(a, torch.mm(torch.mm(u, torch.diag(s)), v.t()))
//     tensor(8.6531e-07)
//     >>> a_big = torch.randn(7, 5, 3)
//     >>> u, s, v = torch.svd(a_big)
//     >>> torch.dist(a_big, torch.matmul(torch.matmul(u, torch.diag_embed(s)), v.mT))
//     tensor(2.6503e-06)
//
// .. _the resulting vectors will span the same subspace:
//        (https://en.wikipedia.org/wiki/Singular_value_decomposition#Singular_values,_singular_vectors,_and_their_relation_to_the_SVD)
//
//
//go:linkname Svd py.svd
func Svd(input *py.Object, some *py.Object, computeUv *py.Object) *py.Object
//
// swapaxes(input, axis0, axis1) -> Tensor
//
// Alias for :func:`torch.transpose`.
//
// This function is equivalent to NumPy's swapaxes function.
//
// Examples::
//
//     >>> x = torch.tensor([[[0,1],[2,3]],[[4,5],[6,7]]])
//     >>> x
//     tensor([[[0, 1],
//             [2, 3]],
//
//             [[4, 5],
//             [6, 7]]])
//     >>> torch.swapaxes(x, 0, 1)
//     tensor([[[0, 1],
//             [4, 5]],
//
//             [[2, 3],
//             [6, 7]]])
//     >>> torch.swapaxes(x, 0, 2)
//     tensor([[[0, 4],
//             [2, 6]],
//
//             [[1, 5],
//             [3, 7]]])
//
//
//go:linkname Swapaxes py.swapaxes
func Swapaxes(input *py.Object, axis0 *py.Object, axis1 *py.Object) *py.Object
//
// swapdims(input, dim0, dim1) -> Tensor
//
// Alias for :func:`torch.transpose`.
//
// This function is equivalent to NumPy's swapaxes function.
//
// Examples::
//
//     >>> x = torch.tensor([[[0,1],[2,3]],[[4,5],[6,7]]])
//     >>> x
//     tensor([[[0, 1],
//             [2, 3]],
//
//             [[4, 5],
//             [6, 7]]])
//     >>> torch.swapdims(x, 0, 1)
//     tensor([[[0, 1],
//             [4, 5]],
//
//             [[2, 3],
//             [6, 7]]])
//     >>> torch.swapdims(x, 0, 2)
//     tensor([[[0, 4],
//             [2, 6]],
//
//             [[1, 5],
//             [3, 7]]])
//
//
//go:linkname Swapdims py.swapdims
func Swapdims(input *py.Object, dim0 *py.Object, dim1 *py.Object) *py.Object
// None
//
//go:linkname SymConstrainRange py.sym_constrain_range
func SymConstrainRange(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname SymConstrainRangeForSize py.sym_constrain_range_for_size
func SymConstrainRangeForSize(__llgo_va_list ...interface{}) *py.Object
//
// t(input) -> Tensor
//
// Expects :attr:`input` to be <= 2-D tensor and transposes dimensions 0
// and 1.
//
// 0-D and 1-D tensors are returned as is. When input is a 2-D tensor this
// is equivalent to ``transpose(input, 0, 1)``.
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> x = torch.randn(())
//     >>> x
//     tensor(0.1995)
//     >>> torch.t(x)
//     tensor(0.1995)
//     >>> x = torch.randn(3)
//     >>> x
//     tensor([ 2.4320, -0.4608,  0.7702])
//     >>> torch.t(x)
//     tensor([ 2.4320, -0.4608,  0.7702])
//     >>> x = torch.randn(2, 3)
//     >>> x
//     tensor([[ 0.4875,  0.9158, -0.5872],
//             [ 0.3938, -0.6929,  0.6932]])
//     >>> torch.t(x)
//     tensor([[ 0.4875,  0.3938],
//             [ 0.9158, -0.6929],
//             [-0.5872,  0.6932]])
//
// See also :func:`torch.transpose`.
//
//
//go:linkname T py.t
func T(input *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.t`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname TCopy py.t_copy
func TCopy(__llgo_va_list ...interface{}) *py.Object
//
// take(input, index) -> Tensor
//
// Returns a new tensor with the elements of :attr:`input` at the given indices.
// The input tensor is treated as if it were viewed as a 1-D tensor. The result
// takes the same shape as the indices.
//
// Args:
//     input (Tensor): the input tensor.
//     index (LongTensor): the indices into tensor
//
// Example::
//
//     >>> src = torch.tensor([[4, 3, 5],
//     ...                     [6, 7, 8]])
//     >>> torch.take(src, torch.tensor([0, 2, 5]))
//     tensor([ 4,  5,  8])
//
//
//go:linkname Take py.take
func Take(input *py.Object, index *py.Object) *py.Object
//
// take_along_dim(input, indices, dim=None, *, out=None) -> Tensor
//
// Selects values from :attr:`input` at the 1-dimensional indices from :attr:`indices` along the given :attr:`dim`.
//
// If :attr:`dim` is None, the input array is treated as if it has been flattened to 1d.
//
// Functions that return indices along a dimension, like :func:`torch.argmax` and :func:`torch.argsort`,
// are designed to work with this function. See the examples below.
//
// .. note::
//     This function is similar to NumPy's `take_along_axis`.
//     See also :func:`torch.gather`.
//
// Args:
//     input (Tensor): the input tensor.
//     indices (tensor): the indices into :attr:`input`. Must have long dtype.
//     dim (int, optional): dimension to select along.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> t = torch.tensor([[10, 30, 20], [60, 40, 50]])
//     >>> max_idx = torch.argmax(t)
//     >>> torch.take_along_dim(t, max_idx)
//     tensor([60])
//     >>> sorted_idx = torch.argsort(t, dim=1)
//     >>> torch.take_along_dim(t, sorted_idx, dim=1)
//     tensor([[10, 20, 30],
//             [40, 50, 60]])
//
//
//go:linkname TakeAlongDim py.take_along_dim
func TakeAlongDim(input *py.Object, indices *py.Object, dim *py.Object) *py.Object
//
// tan(input, *, out=None) -> Tensor
//
// Returns a new tensor with the tangent of the elements of :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \tan(\text{input}_{i})
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([-1.2027, -1.7687,  0.4412, -1.3856])
//     >>> torch.tan(a)
//     tensor([-2.5930,  4.9859,  0.4722, -5.3366])
//
//
//go:linkname Tan py.tan
func Tan(input *py.Object) *py.Object
// None
//
//go:linkname Tan_ py.tan_
func Tan_(__llgo_va_list ...interface{}) *py.Object
//
// tanh(input, *, out=None) -> Tensor
//
// Returns a new tensor with the hyperbolic tangent of the elements
// of :attr:`input`.
//
// .. math::
//     \text{out}_{i} = \tanh(\text{input}_{i})
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([ 0.8986, -0.7279,  1.1745,  0.2611])
//     >>> torch.tanh(a)
//     tensor([ 0.7156, -0.6218,  0.8257,  0.2553])
//
//
//go:linkname Tanh py.tanh
func Tanh(input *py.Object) *py.Object
// None
//
//go:linkname Tanh_ py.tanh_
func Tanh_(__llgo_va_list ...interface{}) *py.Object
//
// tensor(data, *, dtype=None, device=None, requires_grad=False, pin_memory=False) -> Tensor
//
// Constructs a tensor with no autograd history (also known as a "leaf tensor", see :doc:`/notes/autograd`) by copying :attr:`data`.
//
// .. warning::
//
//     When working with tensors prefer using :func:`torch.Tensor.clone`,
//     :func:`torch.Tensor.detach`, and :func:`torch.Tensor.requires_grad_` for
//     readability. Letting `t` be a tensor, ``torch.tensor(t)`` is equivalent to
//     ``t.clone().detach()``, and ``torch.tensor(t, requires_grad=True)``
//     is equivalent to ``t.clone().detach().requires_grad_(True)``.
//
// .. seealso::
//
//     :func:`torch.as_tensor` preserves autograd history and avoids copies where possible.
//     :func:`torch.from_numpy` creates a tensor that shares storage with a NumPy array.
//
// Args:
//     data (array_like): Initial data for the tensor. Can be a list, tuple,
//         NumPy ``ndarray``, scalar, and other types.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, infers data type from :attr:`data`.
//     device (:class:`torch.device`, optional): the device of the constructed tensor. If None and data is a tensor
//         then the device of data is used. If None and data is not a tensor then
//         the result tensor is constructed on the current device.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     pin_memory (bool, optional): If set, returned tensor would be allocated in
//         the pinned memory. Works only for CPU tensors. Default: ``False``.
//
//
// Example::
//
//     >>> torch.tensor([[0.1, 1.2], [2.2, 3.1], [4.9, 5.2]])
//     tensor([[ 0.1000,  1.2000],
//             [ 2.2000,  3.1000],
//             [ 4.9000,  5.2000]])
//
//     >>> torch.tensor([0, 1])  # Type inference on data
//     tensor([ 0,  1])
//
//     >>> torch.tensor([[0.11111, 0.222222, 0.3333333]],
//     ...              dtype=torch.float64,
//     ...              device=torch.device('cuda:0'))  # creates a double tensor on a CUDA device
//     tensor([[ 0.1111,  0.2222,  0.3333]], dtype=torch.float64, device='cuda:0')
//
//     >>> torch.tensor(3.14159)  # Create a zero-dimensional (scalar) tensor
//     tensor(3.1416)
//
//     >>> torch.tensor([])  # Create an empty tensor (of size (0,))
//     tensor([])
//
//
//go:linkname Tensor py.tensor
func Tensor(data *py.Object) *py.Object
//
// tensor_split(input, indices_or_sections, dim=0) -> List of Tensors
//
// Splits a tensor into multiple sub-tensors, all of which are views of :attr:`input`,
// along dimension :attr:`dim` according to the indices or number of sections specified
// by :attr:`indices_or_sections`. This function is based on NumPy's
// :func:`numpy.array_split`.
//
// Args:
//     input (Tensor): the tensor to split
//     indices_or_sections (Tensor, int or list or tuple of ints):
//         If :attr:`indices_or_sections` is an integer ``n`` or a zero dimensional long tensor
//         with value ``n``, :attr:`input` is split into ``n`` sections along dimension :attr:`dim`.
//         If :attr:`input` is divisible by ``n`` along dimension :attr:`dim`, each
//         section will be of equal size, :code:`input.size(dim) / n`. If :attr:`input`
//         is not divisible by ``n``, the sizes of the first :code:`int(input.size(dim) % n)`
//         sections will have size :code:`int(input.size(dim) / n) + 1`, and the rest will
//         have size :code:`int(input.size(dim) / n)`.
//
//         If :attr:`indices_or_sections` is a list or tuple of ints, or a one-dimensional long
//         tensor, then :attr:`input` is split along dimension :attr:`dim` at each of the indices
//         in the list, tuple or tensor. For instance, :code:`indices_or_sections=[2, 3]` and :code:`dim=0`
//         would result in the tensors :code:`input[:2]`, :code:`input[2:3]`, and :code:`input[3:]`.
//
//         If :attr:`indices_or_sections` is a tensor, it must be a zero-dimensional or one-dimensional
//         long tensor on the CPU.
//
//     dim (int, optional): dimension along which to split the tensor. Default: ``0``
//
// Example::
//
//     >>> x = torch.arange(8)
//     >>> torch.tensor_split(x, 3)
//     (tensor([0, 1, 2]), tensor([3, 4, 5]), tensor([6, 7]))
//
//     >>> x = torch.arange(7)
//     >>> torch.tensor_split(x, 3)
//     (tensor([0, 1, 2]), tensor([3, 4]), tensor([5, 6]))
//     >>> torch.tensor_split(x, (1, 6))
//     (tensor([0]), tensor([1, 2, 3, 4, 5]), tensor([6]))
//
//     >>> x = torch.arange(14).reshape(2, 7)
//     >>> x
//     tensor([[ 0,  1,  2,  3,  4,  5,  6],
//             [ 7,  8,  9, 10, 11, 12, 13]])
//     >>> torch.tensor_split(x, 3, dim=1)
//     (tensor([[0, 1, 2],
//             [7, 8, 9]]),
//      tensor([[ 3,  4],
//             [10, 11]]),
//      tensor([[ 5,  6],
//             [12, 13]]))
//     >>> torch.tensor_split(x, (1, 6), dim=1)
//     (tensor([[0],
//             [7]]),
//      tensor([[ 1,  2,  3,  4,  5],
//             [ 8,  9, 10, 11, 12]]),
//      tensor([[ 6],
//             [13]]))
//
//
//go:linkname TensorSplit py.tensor_split
func TensorSplit(input *py.Object, indicesOrSections *py.Object, dim *py.Object) *py.Object
// Returns a contraction of a and b over multiple dimensions.
//
//     :attr:`tensordot` implements a generalized matrix product.
//
//     Args:
//       a (Tensor): Left tensor to contract
//       b (Tensor): Right tensor to contract
//       dims (int or Tuple[List[int], List[int]] or List[List[int]] containing two lists or Tensor): number of dimensions to
//          contract or explicit lists of dimensions for :attr:`a` and
//          :attr:`b` respectively
//
//     When called with a non-negative integer argument :attr:`dims` = :math:`d`, and
//     the number of dimensions of :attr:`a` and :attr:`b` is :math:`m` and :math:`n`,
//     respectively, :func:`~torch.tensordot` computes
//
//     .. math::
//         r_{i_0,...,i_{m-d}, i_d,...,i_n}
//           = \sum_{k_0,...,k_{d-1}} a_{i_0,...,i_{m-d},k_0,...,k_{d-1}} \times b_{k_0,...,k_{d-1}, i_d,...,i_n}.
//
//     When called with :attr:`dims` of the list form, the given dimensions will be contracted
//     in place of the last :math:`d` of :attr:`a` and the first :math:`d` of :math:`b`. The sizes
//     in these dimensions must match, but :func:`~torch.tensordot` will deal with broadcasted
//     dimensions.
//
//     Examples::
//
//         >>> a = torch.arange(60.).reshape(3, 4, 5)
//         >>> b = torch.arange(24.).reshape(4, 3, 2)
//         >>> torch.tensordot(a, b, dims=([1, 0], [0, 1]))
//         tensor([[4400., 4730.],
//                 [4532., 4874.],
//                 [4664., 5018.],
//                 [4796., 5162.],
//                 [4928., 5306.]])
//
//         >>> # xdoctest: +REQUIRES(env:TORCH_DOCTEST_CUDA)
//         >>> a = torch.randn(3, 4, 5, device='cuda')
//         >>> b = torch.randn(4, 5, 6, device='cuda')
//         >>> c = torch.tensordot(a, b, dims=2).cpu()
//         tensor([[ 8.3504, -2.5436,  6.2922,  2.7556, -1.0732,  3.2741],
//                 [ 3.3161,  0.0704,  5.0187, -0.4079, -4.3126,  4.8744],
//                 [ 0.8223,  3.9445,  3.2168, -0.2400,  3.4117,  1.7780]])
//
//         >>> a = torch.randn(3, 5, 4, 6)
//         >>> b = torch.randn(6, 4, 5, 3)
//         >>> torch.tensordot(a, b, dims=([2, 1, 3], [1, 2, 0]))
//         tensor([[  7.7193,  -2.4867, -10.3204],
//                 [  1.5513, -14.4737,  -6.5113],
//                 [ -0.2850,   4.2573,  -3.5997]])
//
//
//go:linkname Tensordot py.tensordot
func Tensordot(a *py.Object, b *py.Object, dims *py.Object, out *py.Object) *py.Object
// None
//
//go:linkname Threshold py.threshold
func Threshold(__llgo_va_list ...interface{}) *py.Object
//
// threshold_(input, threshold, value) -> Tensor
//
// In-place version of :func:`~threshold`.
//
//
//go:linkname Threshold_ py.threshold_
func Threshold_(input *py.Object, threshold *py.Object, value *py.Object) *py.Object
//
// tile(input, dims) -> Tensor
//
// Constructs a tensor by repeating the elements of :attr:`input`.
// The :attr:`dims` argument specifies the number of repetitions
// in each dimension.
//
// If :attr:`dims` specifies fewer dimensions than :attr:`input` has, then
// ones are prepended to :attr:`dims` until all dimensions are specified.
// For example, if :attr:`input` has shape (8, 6, 4, 2) and :attr:`dims`
// is (2, 2), then :attr:`dims` is treated as (1, 1, 2, 2).
//
// Analogously, if :attr:`input` has fewer dimensions than :attr:`dims`
// specifies, then :attr:`input` is treated as if it were unsqueezed at
// dimension zero until it has as many dimensions as :attr:`dims` specifies.
// For example, if :attr:`input` has shape (4, 2) and :attr:`dims`
// is (3, 3, 2, 2), then :attr:`input` is treated as if it had the
// shape (1, 1, 4, 2).
//
// .. note::
//
//     This function is similar to NumPy's tile function.
//
// Args:
//     input (Tensor): the tensor whose elements to repeat.
//     dims (tuple): the number of repetitions per dimension.
//
// Example::
//
//     >>> x = torch.tensor([1, 2, 3])
//     >>> x.tile((2,))
//     tensor([1, 2, 3, 1, 2, 3])
//     >>> y = torch.tensor([[1, 2], [3, 4]])
//     >>> torch.tile(y, (2, 2))
//     tensor([[1, 2, 1, 2],
//             [3, 4, 3, 4],
//             [1, 2, 1, 2],
//             [3, 4, 3, 4]])
//
//
//go:linkname Tile py.tile
func Tile(input *py.Object, dims *py.Object) *py.Object
//
// topk(input, k, dim=None, largest=True, sorted=True, *, out=None) -> (Tensor, LongTensor)
//
// Returns the :attr:`k` largest elements of the given :attr:`input` tensor along
// a given dimension.
//
// If :attr:`dim` is not given, the last dimension of the `input` is chosen.
//
// If :attr:`largest` is ``False`` then the `k` smallest elements are returned.
//
// A namedtuple of `(values, indices)` is returned with the `values` and
// `indices` of the largest `k` elements of each row of the `input` tensor in the
// given dimension `dim`.
//
// The boolean option :attr:`sorted` if ``True``, will make sure that the returned
// `k` elements are themselves sorted
//
// Args:
//     input (Tensor): the input tensor.
//     k (int): the k in "top-k"
//     dim (int, optional): the dimension to sort along
//     largest (bool, optional): controls whether to return largest or
//            smallest elements
//     sorted (bool, optional): controls whether to return the elements
//            in sorted order
//
// Keyword args:
//     out (tuple, optional): the output tuple of (Tensor, LongTensor) that can be
//         optionally given to be used as output buffers
//
// Example::
//
//     >>> x = torch.arange(1., 6.)
//     >>> x
//     tensor([ 1.,  2.,  3.,  4.,  5.])
//     >>> torch.topk(x, 3)
//     torch.return_types.topk(values=tensor([5., 4., 3.]), indices=tensor([4, 3, 2]))
//
//
//go:linkname Topk py.topk
func Topk(input *py.Object, k *py.Object, dim *py.Object, largest *py.Object, sorted *py.Object) *py.Object
//
// trace(input) -> Tensor
//
// Returns the sum of the elements of the diagonal of the input 2-D matrix.
//
// Example::
//
//     >>> x = torch.arange(1., 10.).view(3, 3)
//     >>> x
//     tensor([[ 1.,  2.,  3.],
//             [ 4.,  5.,  6.],
//             [ 7.,  8.,  9.]])
//     >>> torch.trace(x)
//     tensor(15.)
//
//
//go:linkname Trace py.trace
func Trace(input *py.Object) *py.Object
//
// transpose(input, dim0, dim1) -> Tensor
//
// Returns a tensor that is a transposed version of :attr:`input`.
// The given dimensions :attr:`dim0` and :attr:`dim1` are swapped.
//
// If :attr:`input` is a strided tensor then the resulting :attr:`out`
// tensor shares its underlying storage with the :attr:`input` tensor, so
// changing the content of one would change the content of the other.
//
// If :attr:`input` is a :ref:`sparse tensor <sparse-docs>` then the
// resulting :attr:`out` tensor *does not* share the underlying storage
// with the :attr:`input` tensor.
//
// If :attr:`input` is a :ref:`sparse tensor <sparse-docs>` with compressed
// layout (SparseCSR, SparseBSR, SparseCSC or SparseBSC) the arguments
// :attr:`dim0` and :attr:`dim1` must be both batch dimensions, or must
// both be sparse dimensions. The batch dimensions of a sparse tensor are the
// dimensions preceding the sparse dimensions.
//
// .. note::
//     Transpositions which interchange the sparse dimensions of a `SparseCSR`
//     or `SparseCSC` layout tensor will result in the layout changing between
//     the two options. Transposition of the sparse dimensions of a ` SparseBSR`
//     or `SparseBSC` layout tensor will likewise generate a result with the
//     opposite layout.
//
//
// Args:
//     input (Tensor): the input tensor.
//     dim0 (int): the first dimension to be transposed
//     dim1 (int): the second dimension to be transposed
//
// Example::
//
//     >>> x = torch.randn(2, 3)
//     >>> x
//     tensor([[ 1.0028, -0.9893,  0.5809],
//             [-0.1669,  0.7299,  0.4942]])
//     >>> torch.transpose(x, 0, 1)
//     tensor([[ 1.0028, -0.1669],
//             [-0.9893,  0.7299],
//             [ 0.5809,  0.4942]])
//
// See also :func:`torch.t`.
//
//
//go:linkname Transpose py.transpose
func Transpose(input *py.Object, dim0 *py.Object, dim1 *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.transpose`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname TransposeCopy py.transpose_copy
func TransposeCopy(__llgo_va_list ...interface{}) *py.Object
//
// trapezoid(y, x=None, *, dx=None, dim=-1) -> Tensor
//
// Computes the `trapezoidal rule <https://en.wikipedia.org/wiki/Trapezoidal_rule>`_ along
// :attr:`dim`. By default the spacing between elements is assumed to be 1, but
// :attr:`dx` can be used to specify a different constant spacing, and :attr:`x` can be
// used to specify arbitrary spacing along :attr:`dim`.
//
//
// Assuming :attr:`y` is a one-dimensional tensor with elements :math:`{y_0, y_1, ..., y_n}`,
// the default computation is
//
// .. math::
//     \begin{aligned}
//         \sum_{i = 1}^{n-1} \frac{1}{2} (y_i + y_{i-1})
//     \end{aligned}
//
// When :attr:`dx` is specified the computation becomes
//
// .. math::
//     \begin{aligned}
//         \sum_{i = 1}^{n-1} \frac{\Delta x}{2} (y_i + y_{i-1})
//     \end{aligned}
//
// effectively multiplying the result by :attr:`dx`. When :attr:`x` is specified,
// assuming :attr:`x` is also a one-dimensional tensor with
// elements :math:`{x_0, x_1, ..., x_n}`, the computation becomes
//
// .. math::
//     \begin{aligned}
//         \sum_{i = 1}^{n-1} \frac{(x_i - x_{i-1})}{2} (y_i + y_{i-1})
//     \end{aligned}
//
// When :attr:`x` and :attr:`y` have the same size, the computation is as described above and no broadcasting is needed.
// The broadcasting behavior of this function is as follows when their sizes are different. For both :attr:`x`
// and :attr:`y`, the function computes the difference between consecutive elements along
// dimension :attr:`dim`. This effectively creates two tensors, `x_diff` and `y_diff`, that have
// the same shape as the original tensors except their lengths along the dimension :attr:`dim` is reduced by 1.
// After that, those two tensors are broadcast together to compute final output as part of the trapezoidal rule.
// See the examples below for details.
//
// .. note::
//     The trapezoidal rule is a technique for approximating the definite integral of a function
//     by averaging its left and right Riemann sums. The approximation becomes more accurate as
//     the resolution of the partition increases.
//
// Arguments:
//     y (Tensor): Values to use when computing the trapezoidal rule.
//     x (Tensor): If specified, defines spacing between values as specified above.
//
// Keyword arguments:
//     dx (float): constant spacing between values. If neither :attr:`x` or :attr:`dx`
//         are specified then this defaults to 1. Effectively multiplies the result by its value.
//     dim (int): The dimension along which to compute the trapezoidal rule.
//         The last (inner-most) dimension by default.
//
// Examples::
//
//     >>> # Computes the trapezoidal rule in 1D, spacing is implicitly 1
//     >>> y = torch.tensor([1, 5, 10])
//     >>> torch.trapezoid(y)
//     tensor(10.5)
//
//     >>> # Computes the same trapezoidal rule directly to verify
//     >>> (1 + 10 + 10) / 2
//     10.5
//
//     >>> # Computes the trapezoidal rule in 1D with constant spacing of 2
//     >>> # NOTE: the result is the same as before, but multiplied by 2
//     >>> torch.trapezoid(y, dx=2)
//     21.0
//
//     >>> # Computes the trapezoidal rule in 1D with arbitrary spacing
//     >>> x = torch.tensor([1, 3, 6])
//     >>> torch.trapezoid(y, x)
//     28.5
//
//     >>> # Computes the same trapezoidal rule directly to verify
//     >>> ((3 - 1) * (1 + 5) + (6 - 3) * (5 + 10)) / 2
//     28.5
//
//     >>> # Computes the trapezoidal rule for each row of a 3x3 matrix
//     >>> y = torch.arange(9).reshape(3, 3)
//     tensor([[0, 1, 2],
//             [3, 4, 5],
//             [6, 7, 8]])
//     >>> torch.trapezoid(y)
//     tensor([ 2., 8., 14.])
//
//     >>> # Computes the trapezoidal rule for each column of the matrix
//     >>> torch.trapezoid(y, dim=0)
//     tensor([ 6., 8., 10.])
//
//     >>> # Computes the trapezoidal rule for each row of a 3x3 ones matrix
//     >>> #   with the same arbitrary spacing
//     >>> y = torch.ones(3, 3)
//     >>> x = torch.tensor([1, 3, 6])
//     >>> torch.trapezoid(y, x)
//     array([5., 5., 5.])
//
//     >>> # Computes the trapezoidal rule for each row of a 3x3 ones matrix
//     >>> #   with different arbitrary spacing per row
//     >>> y = torch.ones(3, 3)
//     >>> x = torch.tensor([[1, 2, 3], [1, 3, 5], [1, 4, 7]])
//     >>> torch.trapezoid(y, x)
//     array([2., 4., 6.])
//
//
//go:linkname Trapezoid py.trapezoid
func Trapezoid(y *py.Object, x *py.Object) *py.Object
//
// trapz(y, x, *, dim=-1) -> Tensor
//
// Alias for :func:`torch.trapezoid`.
//
//
//go:linkname Trapz py.trapz
func Trapz(y *py.Object, x *py.Object) *py.Object
//
// triangular_solve(b, A, upper=True, transpose=False, unitriangular=False, *, out=None) -> (Tensor, Tensor)
//
// Solves a system of equations with a square upper or lower triangular invertible matrix :math:`A`
// and multiple right-hand sides :math:`b`.
//
// In symbols, it solves :math:`AX = b` and assumes :math:`A` is square upper-triangular
// (or lower-triangular if :attr:`upper`\ `= False`) and does not have zeros on the diagonal.
//
// `torch.triangular_solve(b, A)` can take in 2D inputs `b, A` or inputs that are
// batches of 2D matrices. If the inputs are batches, then returns
// batched outputs `X`
//
// If the diagonal of :attr:`A` contains zeros or elements that are very close to zero and
// :attr:`unitriangular`\ `= False` (default) or if the input matrix is badly conditioned,
// the result may contain `NaN` s.
//
// Supports input of float, double, cfloat and cdouble data types.
//
// .. warning::
//
//     :func:`torch.triangular_solve` is deprecated in favor of :func:`torch.linalg.solve_triangular`
//     and will be removed in a future PyTorch release.
//     :func:`torch.linalg.solve_triangular` has its arguments reversed and does not return a
//     copy of one of the inputs.
//
//     ``X = torch.triangular_solve(B, A).solution`` should be replaced with
//
//     .. code:: python
//
//         X = torch.linalg.solve_triangular(A, B)
//
// Args:
//     b (Tensor): multiple right-hand sides of size :math:`(*, m, k)` where
//                 :math:`*` is zero of more batch dimensions
//     A (Tensor): the input triangular coefficient matrix of size :math:`(*, m, m)`
//                 where :math:`*` is zero or more batch dimensions
//     upper (bool, optional): whether :math:`A` is upper or lower triangular. Default: ``True``.
//     transpose (bool, optional): solves `op(A)X = b` where `op(A) = A^T` if this flag is ``True``,
//                                 and `op(A) = A` if it is ``False``. Default: ``False``.
//     unitriangular (bool, optional): whether :math:`A` is unit triangular.
//         If True, the diagonal elements of :math:`A` are assumed to be
//         1 and not referenced from :math:`A`. Default: ``False``.
//
// Keyword args:
//     out ((Tensor, Tensor), optional): tuple of two tensors to write
//         the output to. Ignored if `None`. Default: `None`.
//
// Returns:
//     A namedtuple `(solution, cloned_coefficient)` where `cloned_coefficient`
//     is a clone of :math:`A` and `solution` is the solution :math:`X` to :math:`AX = b`
//     (or whatever variant of the system of equations, depending on the keyword arguments.)
//
// Examples::
//
//     >>> A = torch.randn(2, 2).triu()
//     >>> A
//     tensor([[ 1.1527, -1.0753],
//             [ 0.0000,  0.7986]])
//     >>> b = torch.randn(2, 3)
//     >>> b
//     tensor([[-0.0210,  2.3513, -1.5492],
//             [ 1.5429,  0.7403, -1.0243]])
//     >>> torch.triangular_solve(b, A)
//     torch.return_types.triangular_solve(
//     solution=tensor([[ 1.7841,  2.9046, -2.5405],
//             [ 1.9320,  0.9270, -1.2826]]),
//     cloned_coefficient=tensor([[ 1.1527, -1.0753],
//             [ 0.0000,  0.7986]]))
//
//
//go:linkname TriangularSolve py.triangular_solve
func TriangularSolve(b *py.Object, A *py.Object, upper *py.Object, transpose *py.Object, unitriangular *py.Object) *py.Object
//
// tril(input, diagonal=0, *, out=None) -> Tensor
//
// Returns the lower triangular part of the matrix (2-D tensor) or batch of matrices
// :attr:`input`, the other elements of the result tensor :attr:`out` are set to 0.
//
// The lower triangular part of the matrix is defined as the elements on and
// below the diagonal.
//
// The argument :attr:`diagonal` controls which diagonal to consider. If
// :attr:`diagonal` = 0, all elements on and below the main diagonal are
// retained. A positive value includes just as many diagonals above the main
// diagonal, and similarly a negative value excludes just as many diagonals below
// the main diagonal. The main diagonal are the set of indices
// :math:`\lbrace (i, i) \rbrace` for :math:`i \in [0, \min\{d_{1}, d_{2}\} - 1]` where
// :math:`d_{1}, d_{2}` are the dimensions of the matrix.
//
// Args:
//     input (Tensor): the input tensor.
//     diagonal (int, optional): the diagonal to consider
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(3, 3)
//     >>> a
//     tensor([[-1.0813, -0.8619,  0.7105],
//             [ 0.0935,  0.1380,  2.2112],
//             [-0.3409, -0.9828,  0.0289]])
//     >>> torch.tril(a)
//     tensor([[-1.0813,  0.0000,  0.0000],
//             [ 0.0935,  0.1380,  0.0000],
//             [-0.3409, -0.9828,  0.0289]])
//
//     >>> b = torch.randn(4, 6)
//     >>> b
//     tensor([[ 1.2219,  0.5653, -0.2521, -0.2345,  1.2544,  0.3461],
//             [ 0.4785, -0.4477,  0.6049,  0.6368,  0.8775,  0.7145],
//             [ 1.1502,  3.2716, -1.1243, -0.5413,  0.3615,  0.6864],
//             [-0.0614, -0.7344, -1.3164, -0.7648, -1.4024,  0.0978]])
//     >>> torch.tril(b, diagonal=1)
//     tensor([[ 1.2219,  0.5653,  0.0000,  0.0000,  0.0000,  0.0000],
//             [ 0.4785, -0.4477,  0.6049,  0.0000,  0.0000,  0.0000],
//             [ 1.1502,  3.2716, -1.1243, -0.5413,  0.0000,  0.0000],
//             [-0.0614, -0.7344, -1.3164, -0.7648, -1.4024,  0.0000]])
//     >>> torch.tril(b, diagonal=-1)
//     tensor([[ 0.0000,  0.0000,  0.0000,  0.0000,  0.0000,  0.0000],
//             [ 0.4785,  0.0000,  0.0000,  0.0000,  0.0000,  0.0000],
//             [ 1.1502,  3.2716,  0.0000,  0.0000,  0.0000,  0.0000],
//             [-0.0614, -0.7344, -1.3164,  0.0000,  0.0000,  0.0000]])
//
//
//go:linkname Tril py.tril
func Tril(input *py.Object, diagonal *py.Object) *py.Object
//
// tril_indices(row, col, offset=0, *, dtype=torch.long, device='cpu', layout=torch.strided) -> Tensor
//
// Returns the indices of the lower triangular part of a :attr:`row`-by-
// :attr:`col` matrix in a 2-by-N Tensor, where the first row contains row
// coordinates of all indices and the second row contains column coordinates.
// Indices are ordered based on rows and then columns.
//
// The lower triangular part of the matrix is defined as the elements on and
// below the diagonal.
//
// The argument :attr:`offset` controls which diagonal to consider. If
// :attr:`offset` = 0, all elements on and below the main diagonal are
// retained. A positive value includes just as many diagonals above the main
// diagonal, and similarly a negative value excludes just as many diagonals below
// the main diagonal. The main diagonal are the set of indices
// :math:`\lbrace (i, i) \rbrace` for :math:`i \in [0, \min\{d_{1}, d_{2}\} - 1]`
// where :math:`d_{1}, d_{2}` are the dimensions of the matrix.
//
// .. note::
//     When running on CUDA, ``row * col`` must be less than :math:`2^{59}` to
//     prevent overflow during calculation.
//
// Args:
//     row (``int``): number of rows in the 2-D matrix.
//     col (``int``): number of columns in the 2-D matrix.
//     offset (``int``): diagonal offset from the main diagonal.
//         Default: if not provided, 0.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, ``torch.long``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     layout (:class:`torch.layout`, optional): currently only support ``torch.strided``.
//
// Example::
//
//     >>> a = torch.tril_indices(3, 3)
//     >>> a
//     tensor([[0, 1, 1, 2, 2, 2],
//             [0, 0, 1, 0, 1, 2]])
//
//     >>> a = torch.tril_indices(4, 3, -1)
//     >>> a
//     tensor([[1, 2, 2, 3, 3, 3],
//             [0, 0, 1, 0, 1, 2]])
//
//     >>> a = torch.tril_indices(4, 3, 1)
//     >>> a
//     tensor([[0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3],
//             [0, 1, 0, 1, 2, 0, 1, 2, 0, 1, 2]])
//
//
//go:linkname TrilIndices py.tril_indices
func TrilIndices(row *py.Object, col *py.Object, offset *py.Object) *py.Object
// None
//
//go:linkname TripletMarginLoss py.triplet_margin_loss
func TripletMarginLoss(__llgo_va_list ...interface{}) *py.Object
//
// triu(input, diagonal=0, *, out=None) -> Tensor
//
// Returns the upper triangular part of a matrix (2-D tensor) or batch of matrices
// :attr:`input`, the other elements of the result tensor :attr:`out` are set to 0.
//
// The upper triangular part of the matrix is defined as the elements on and
// above the diagonal.
//
// The argument :attr:`diagonal` controls which diagonal to consider. If
// :attr:`diagonal` = 0, all elements on and above the main diagonal are
// retained. A positive value excludes just as many diagonals above the main
// diagonal, and similarly a negative value includes just as many diagonals below
// the main diagonal. The main diagonal are the set of indices
// :math:`\lbrace (i, i) \rbrace` for :math:`i \in [0, \min\{d_{1}, d_{2}\} - 1]` where
// :math:`d_{1}, d_{2}` are the dimensions of the matrix.
//
// Args:
//     input (Tensor): the input tensor.
//     diagonal (int, optional): the diagonal to consider
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(3, 3)
//     >>> a
//     tensor([[ 0.2309,  0.5207,  2.0049],
//             [ 0.2072, -1.0680,  0.6602],
//             [ 0.3480, -0.5211, -0.4573]])
//     >>> torch.triu(a)
//     tensor([[ 0.2309,  0.5207,  2.0049],
//             [ 0.0000, -1.0680,  0.6602],
//             [ 0.0000,  0.0000, -0.4573]])
//     >>> torch.triu(a, diagonal=1)
//     tensor([[ 0.0000,  0.5207,  2.0049],
//             [ 0.0000,  0.0000,  0.6602],
//             [ 0.0000,  0.0000,  0.0000]])
//     >>> torch.triu(a, diagonal=-1)
//     tensor([[ 0.2309,  0.5207,  2.0049],
//             [ 0.2072, -1.0680,  0.6602],
//             [ 0.0000, -0.5211, -0.4573]])
//
//     >>> b = torch.randn(4, 6)
//     >>> b
//     tensor([[ 0.5876, -0.0794, -1.8373,  0.6654,  0.2604,  1.5235],
//             [-0.2447,  0.9556, -1.2919,  1.3378, -0.1768, -1.0857],
//             [ 0.4333,  0.3146,  0.6576, -1.0432,  0.9348, -0.4410],
//             [-0.9888,  1.0679, -1.3337, -1.6556,  0.4798,  0.2830]])
//     >>> torch.triu(b, diagonal=1)
//     tensor([[ 0.0000, -0.0794, -1.8373,  0.6654,  0.2604,  1.5235],
//             [ 0.0000,  0.0000, -1.2919,  1.3378, -0.1768, -1.0857],
//             [ 0.0000,  0.0000,  0.0000, -1.0432,  0.9348, -0.4410],
//             [ 0.0000,  0.0000,  0.0000,  0.0000,  0.4798,  0.2830]])
//     >>> torch.triu(b, diagonal=-1)
//     tensor([[ 0.5876, -0.0794, -1.8373,  0.6654,  0.2604,  1.5235],
//             [-0.2447,  0.9556, -1.2919,  1.3378, -0.1768, -1.0857],
//             [ 0.0000,  0.3146,  0.6576, -1.0432,  0.9348, -0.4410],
//             [ 0.0000,  0.0000, -1.3337, -1.6556,  0.4798,  0.2830]])
//
//
//go:linkname Triu py.triu
func Triu(input *py.Object, diagonal *py.Object) *py.Object
//
// triu_indices(row, col, offset=0, *, dtype=torch.long, device='cpu', layout=torch.strided) -> Tensor
//
// Returns the indices of the upper triangular part of a :attr:`row` by
// :attr:`col` matrix in a 2-by-N Tensor, where the first row contains row
// coordinates of all indices and the second row contains column coordinates.
// Indices are ordered based on rows and then columns.
//
// The upper triangular part of the matrix is defined as the elements on and
// above the diagonal.
//
// The argument :attr:`offset` controls which diagonal to consider. If
// :attr:`offset` = 0, all elements on and above the main diagonal are
// retained. A positive value excludes just as many diagonals above the main
// diagonal, and similarly a negative value includes just as many diagonals below
// the main diagonal. The main diagonal are the set of indices
// :math:`\lbrace (i, i) \rbrace` for :math:`i \in [0, \min\{d_{1}, d_{2}\} - 1]`
// where :math:`d_{1}, d_{2}` are the dimensions of the matrix.
//
// .. note::
//     When running on CUDA, ``row * col`` must be less than :math:`2^{59}` to
//     prevent overflow during calculation.
//
// Args:
//     row (``int``): number of rows in the 2-D matrix.
//     col (``int``): number of columns in the 2-D matrix.
//     offset (``int``): diagonal offset from the main diagonal.
//         Default: if not provided, 0.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, ``torch.long``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     layout (:class:`torch.layout`, optional): currently only support ``torch.strided``.
//
// Example::
//
//     >>> a = torch.triu_indices(3, 3)
//     >>> a
//     tensor([[0, 0, 0, 1, 1, 2],
//             [0, 1, 2, 1, 2, 2]])
//
//     >>> a = torch.triu_indices(4, 3, -1)
//     >>> a
//     tensor([[0, 0, 0, 1, 1, 1, 2, 2, 3],
//             [0, 1, 2, 0, 1, 2, 1, 2, 2]])
//
//     >>> a = torch.triu_indices(4, 3, 1)
//     >>> a
//     tensor([[0, 0, 1],
//             [1, 2, 2]])
//
//
//go:linkname TriuIndices py.triu_indices
func TriuIndices(row *py.Object, col *py.Object, offset *py.Object) *py.Object
//
// true_divide(dividend, divisor, *, out) -> Tensor
//
// Alias for :func:`torch.div` with ``rounding_mode=None``.
//
//
//go:linkname TrueDivide py.true_divide
func TrueDivide(dividend *py.Object, divisor *py.Object) *py.Object
//
// trunc(input, *, out=None) -> Tensor
//
// Returns a new tensor with the truncated integer values of
// the elements of :attr:`input`.
//
// For integer inputs, follows the array-api convention of returning a
// copy of the input tensor.
//
// Args:
//     input (Tensor): the input tensor.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.randn(4)
//     >>> a
//     tensor([ 3.4742,  0.5466, -0.8008, -0.9079])
//     >>> torch.trunc(a)
//     tensor([ 3.,  0., -0., -0.])
//
//
//go:linkname Trunc py.trunc
func Trunc(input *py.Object) *py.Object
// None
//
//go:linkname Trunc_ py.trunc_
func Trunc_(__llgo_va_list ...interface{}) *py.Object
//
// unbind(input, dim=0) -> seq
//
// Removes a tensor dimension.
//
// Returns a tuple of all slices along a given dimension, already without it.
//
// Arguments:
//     input (Tensor): the tensor to unbind
//     dim (int): dimension to remove
//
// Example::
//
//     >>> torch.unbind(torch.tensor([[1, 2, 3],
//     >>>                            [4, 5, 6],
//     >>>                            [7, 8, 9]]))
//     (tensor([1, 2, 3]), tensor([4, 5, 6]), tensor([7, 8, 9]))
//
//
//go:linkname Unbind py.unbind
func Unbind(input *py.Object, dim *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.unbind`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname UnbindCopy py.unbind_copy
func UnbindCopy(__llgo_va_list ...interface{}) *py.Object
//
// unflatten(input, dim, sizes) -> Tensor
//
// Expands a dimension of the input tensor over multiple dimensions.
//
// .. seealso::
//
//     :func:`torch.flatten` the inverse of this function. It coalesces several dimensions into one.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int): Dimension to be unflattened, specified as an index into
//          ``input.shape``.
//     sizes (Tuple[int]): New shape of the unflattened dimension.
//          One of its elements can be `-1` in which case the corresponding output
//          dimension is inferred. Otherwise, the product of ``sizes`` *must*
//          equal ``input.shape[dim]``.
//
// Returns:
//     A View of input with the specified dimension unflattened.
//
// Examples::
//     >>> torch.unflatten(torch.randn(3, 4, 1), 1, (2, 2)).shape
//     torch.Size([3, 2, 2, 1])
//     >>> torch.unflatten(torch.randn(3, 4, 1), 1, (-1, 2)).shape
//     torch.Size([3, 2, 2, 1])
//     >>> torch.unflatten(torch.randn(5, 12, 3), -2, (2, 2, 3, 1, 1)).shape
//     torch.Size([5, 2, 2, 3, 1, 1, 3])
//
//
//go:linkname Unflatten py.unflatten
func Unflatten(input *py.Object, dim *py.Object, sizes *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.unfold`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname UnfoldCopy py.unfold_copy
func UnfoldCopy(__llgo_va_list ...interface{}) *py.Object
// Eliminates all but the first element from every consecutive group of equivalent elements.
//
//     .. note:: This function is different from :func:`torch.unique` in the sense that this function
//         only eliminates consecutive duplicate values. This semantics is similar to `std::unique`
//         in C++.
//
//     Args:
//         input (Tensor): the input tensor
//         return_inverse (bool): Whether to also return the indices for where
//             elements in the original input ended up in the returned unique list.
//         return_counts (bool): Whether to also return the counts for each unique
//             element.
//         dim (int): the dimension to apply unique. If ``None``, the unique of the
//             flattened input is returned. default: ``None``
//
//     Returns:
//         (Tensor, Tensor (optional), Tensor (optional)): A tensor or a tuple of tensors containing
//
//             - **output** (*Tensor*): the output list of unique scalar elements.
//             - **inverse_indices** (*Tensor*): (optional) if
//               :attr:`return_inverse` is True, there will be an additional
//               returned tensor (same shape as input) representing the indices
//               for where elements in the original input map to in the output;
//               otherwise, this function will only return a single tensor.
//             - **counts** (*Tensor*): (optional) if
//               :attr:`return_counts` is True, there will be an additional
//               returned tensor (same shape as output or output.size(dim),
//               if dim was specified) representing the number of occurrences
//               for each unique value or tensor.
//
//     Example::
//
//         >>> x = torch.tensor([1, 1, 2, 2, 3, 1, 1, 2])
//         >>> output = torch.unique_consecutive(x)
//         >>> output
//         tensor([1, 2, 3, 1, 2])
//
//         >>> output, inverse_indices = torch.unique_consecutive(x, return_inverse=True)
//         >>> output
//         tensor([1, 2, 3, 1, 2])
//         >>> inverse_indices
//         tensor([0, 0, 1, 1, 2, 3, 3, 4])
//
//         >>> output, counts = torch.unique_consecutive(x, return_counts=True)
//         >>> output
//         tensor([1, 2, 3, 1, 2])
//         >>> counts
//         tensor([2, 2, 1, 2, 1])
//
//
//go:linkname UniqueConsecutive py.unique_consecutive
func UniqueConsecutive(__llgo_va_list ...interface{}) *py.Object
//
// unsafe_chunk(input, chunks, dim=0) -> List of Tensors
//
// Works like :func:`torch.chunk` but without enforcing the autograd restrictions
// on inplace modification of the outputs.
//
// .. warning::
//     This function is safe to use as long as only the input, or only the outputs
//     are modified inplace after calling this function. It is user's
//     responsibility to ensure that is the case. If both the input and one or more
//     of the outputs are modified inplace, gradients computed by autograd will be
//     silently incorrect.
//
//
//go:linkname UnsafeChunk py.unsafe_chunk
func UnsafeChunk(input *py.Object, chunks *py.Object, dim *py.Object) *py.Object
//
// unsafe_split(tensor, split_size_or_sections, dim=0) -> List of Tensors
//
// Works like :func:`torch.split` but without enforcing the autograd restrictions
// on inplace modification of the outputs.
//
// .. warning::
//     This function is safe to use as long as only the input, or only the outputs
//     are modified inplace after calling this function. It is user's
//     responsibility to ensure that is the case. If both the input and one or more
//     of the outputs are modified inplace, gradients computed by autograd will be
//     silently incorrect.
//
//
//go:linkname UnsafeSplit py.unsafe_split
func UnsafeSplit(tensor *py.Object, splitSizeOrSections *py.Object, dim *py.Object) *py.Object
// None
//
//go:linkname UnsafeSplitWithSizes py.unsafe_split_with_sizes
func UnsafeSplitWithSizes(__llgo_va_list ...interface{}) *py.Object
//
// unsqueeze(input, dim) -> Tensor
//
// Returns a new tensor with a dimension of size one inserted at the
// specified position.
//
// The returned tensor shares the same underlying data with this tensor.
//
// A :attr:`dim` value within the range ``[-input.dim() - 1, input.dim() + 1)``
// can be used. Negative :attr:`dim` will correspond to :meth:`unsqueeze`
// applied at :attr:`dim` = ``dim + input.dim() + 1``.
//
// Args:
//     input (Tensor): the input tensor.
//     dim (int): the index at which to insert the singleton dimension
//
// Example::
//
//     >>> x = torch.tensor([1, 2, 3, 4])
//     >>> torch.unsqueeze(x, 0)
//     tensor([[ 1,  2,  3,  4]])
//     >>> torch.unsqueeze(x, 1)
//     tensor([[ 1],
//             [ 2],
//             [ 3],
//             [ 4]])
//
//
//go:linkname Unsqueeze py.unsqueeze
func Unsqueeze(input *py.Object, dim *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.unsqueeze`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname UnsqueezeCopy py.unsqueeze_copy
func UnsqueezeCopy(__llgo_va_list ...interface{}) *py.Object
//
// Performs the same operation as :func:`torch.values`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname ValuesCopy py.values_copy
func ValuesCopy(__llgo_va_list ...interface{}) *py.Object
//
// vander(x, N=None, increasing=False) -> Tensor
//
// Generates a Vandermonde matrix.
//
// The columns of the output matrix are elementwise powers of the input vector :math:`x^{(N-1)}, x^{(N-2)}, ..., x^0`.
// If increasing is True, the order of the columns is reversed :math:`x^0, x^1, ..., x^{(N-1)}`. Such a
// matrix with a geometric progression in each row is named for Alexandre-Theophile Vandermonde.
//
// Arguments:
//     x (Tensor): 1-D input tensor.
//     N (int, optional): Number of columns in the output. If N is not specified,
//         a square array is returned :math:`(N = len(x))`.
//     increasing (bool, optional): Order of the powers of the columns. If True,
//         the powers increase from left to right, if False (the default) they are reversed.
//
// Returns:
//     Tensor: Vandermonde matrix. If increasing is False, the first column is :math:`x^{(N-1)}`,
//     the second :math:`x^{(N-2)}` and so forth. If increasing is True, the columns
//     are :math:`x^0, x^1, ..., x^{(N-1)}`.
//
// Example::
//
//     >>> x = torch.tensor([1, 2, 3, 5])
//     >>> torch.vander(x)
//     tensor([[  1,   1,   1,   1],
//             [  8,   4,   2,   1],
//             [ 27,   9,   3,   1],
//             [125,  25,   5,   1]])
//     >>> torch.vander(x, N=3)
//     tensor([[ 1,  1,  1],
//             [ 4,  2,  1],
//             [ 9,  3,  1],
//             [25,  5,  1]])
//     >>> torch.vander(x, N=3, increasing=True)
//     tensor([[ 1,  1,  1],
//             [ 1,  2,  4],
//             [ 1,  3,  9],
//             [ 1,  5, 25]])
//
//
//
//go:linkname Vander py.vander
func Vander(x *py.Object, N *py.Object, increasing *py.Object) *py.Object
//
// var(input, dim=None, *, correction=1, keepdim=False, out=None) -> Tensor
//
// Calculates the variance over the dimensions specified by :attr:`dim`. :attr:`dim`
// can be a single dimension, list of dimensions, or ``None`` to reduce over all
// dimensions.
//
// The variance (:math:`\sigma^2`) is calculated as
//
// .. math:: \sigma^2 = \frac{1}{\max(0,~N - \delta N)}\sum_{i=0}^{N-1}(x_i-\bar{x})^2
//
// where :math:`x` is the sample set of elements, :math:`\bar{x}` is the
// sample mean, :math:`N` is the number of samples and :math:`\delta N` is
// the :attr:`correction`.
//
//
//
// If :attr:`keepdim` is ``True``, the output tensor is of the same size
// as :attr:`input` except in the dimension(s) :attr:`dim` where it is of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in the
// output tensor having 1 (or ``len(dim)``) fewer dimension(s).
//
//
// Args:
//     input (Tensor): the input tensor.
//
//     dim (int or tuple of ints, optional): the dimension or dimensions to reduce.
//         If ``None``, all dimensions are reduced.
//
//
// Keyword args:
//     correction (int): difference between the sample size and sample degrees of freedom.
//         Defaults to `Bessel's correction`_, ``correction=1``.
//
//         .. versionchanged:: 2.0
//             Previously this argument was called ``unbiased`` and was a boolean
//             with ``True`` corresponding to ``correction=1`` and ``False`` being
//             ``correction=0``.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//     out (Tensor, optional): the output tensor.
//
// Example:
//
//     >>> a = torch.tensor(
//     ...     [[ 0.2035,  1.2959,  1.8101, -0.4644],
//     ...      [ 1.5027, -0.3270,  0.5905,  0.6538],
//     ...      [-1.5745,  1.3330, -0.5596, -0.6548],
//     ...      [ 0.1264, -0.5080,  1.6420,  0.1992]])
//     >>> torch.var(a, dim=1, keepdim=True)
//     tensor([[1.0631],
//             [0.5590],
//             [1.4893],
//             [0.8258]])
//
// .. _Bessel's correction: https://en.wikipedia.org/wiki/Bessel%27s_correction
//
//
//
//go:linkname Var py.var
func Var(input *py.Object, dim *py.Object) *py.Object
//
// var_mean(input, dim=None, *, correction=1, keepdim=False, out=None) -> (Tensor, Tensor)
//
// Calculates the variance and mean over the dimensions specified by :attr:`dim`.
// :attr:`dim` can be a single dimension, list of dimensions, or ``None`` to
// reduce over all dimensions.
//
// The variance (:math:`\sigma^2`) is calculated as
//
// .. math:: \sigma^2 = \frac{1}{\max(0,~N - \delta N)}\sum_{i=0}^{N-1}(x_i-\bar{x})^2
//
// where :math:`x` is the sample set of elements, :math:`\bar{x}` is the
// sample mean, :math:`N` is the number of samples and :math:`\delta N` is
// the :attr:`correction`.
//
//
//
// If :attr:`keepdim` is ``True``, the output tensor is of the same size
// as :attr:`input` except in the dimension(s) :attr:`dim` where it is of size 1.
// Otherwise, :attr:`dim` is squeezed (see :func:`torch.squeeze`), resulting in the
// output tensor having 1 (or ``len(dim)``) fewer dimension(s).
//
//
// Args:
//     input (Tensor): the input tensor.
//
//     dim (int or tuple of ints, optional): the dimension or dimensions to reduce.
//         If ``None``, all dimensions are reduced.
//
//
// Keyword args:
//     correction (int): difference between the sample size and sample degrees of freedom.
//         Defaults to `Bessel's correction`_, ``correction=1``.
//
//         .. versionchanged:: 2.0
//             Previously this argument was called ``unbiased`` and was a boolean
//             with ``True`` corresponding to ``correction=1`` and ``False`` being
//             ``correction=0``.
//     keepdim (bool): whether the output tensor has :attr:`dim` retained or not.
//     out (Tensor, optional): the output tensor.
//
// Returns:
//     A tuple (var, mean) containing the variance and mean.
//
// Example:
//
//     >>> a = torch.tensor(
//     ...     [[ 0.2035,  1.2959,  1.8101, -0.4644],
//     ...      [ 1.5027, -0.3270,  0.5905,  0.6538],
//     ...      [-1.5745,  1.3330, -0.5596, -0.6548],
//     ...      [ 0.1264, -0.5080,  1.6420,  0.1992]])
//     >>> torch.var_mean(a, dim=0, keepdim=True)
//     (tensor([[1.5926, 1.0056, 1.2005, 0.3646]]),
//      tensor([[ 0.0645,  0.4485,  0.8707, -0.0665]]))
//
// .. _Bessel's correction: https://en.wikipedia.org/wiki/Bessel%27s_correction
//
//
//
//go:linkname VarMean py.var_mean
func VarMean(input *py.Object, dim *py.Object) *py.Object
//
// vdot(input, other, *, out=None) -> Tensor
//
// Computes the dot product of two 1D vectors along a dimension.
//
// In symbols, this function computes
//
// .. math::
//
//     \sum_{i=1}^n \overline{x_i}y_i.
//
// where :math:`\overline{x_i}` denotes the conjugate for complex
// vectors, and it is the identity for real vectors.
//
// .. note::
//
//     Unlike NumPy's vdot, torch.vdot intentionally only supports computing the dot product
//     of two 1D tensors with the same number of elements.
//
// .. seealso::
//
//         :func:`torch.linalg.vecdot` computes the dot product of two batches of vectors along a dimension.
//
// Args:
//     input (Tensor): first tensor in the dot product, must be 1D. Its conjugate is used if it's complex.
//     other (Tensor): second tensor in the dot product, must be 1D.
//
// Keyword args:
//
// .. note:: out (Tensor, optional): the output tensor.
//
//
// Example::
//
//     >>> torch.vdot(torch.tensor([2, 3]), torch.tensor([2, 1]))
//     tensor(7)
//     >>> a = torch.tensor((1 +2j, 3 - 1j))
//     >>> b = torch.tensor((2 +1j, 4 - 0j))
//     >>> torch.vdot(a, b)
//     tensor([16.+1.j])
//     >>> torch.vdot(b, a)
//     tensor([16.-1.j])
//
//
//go:linkname Vdot py.vdot
func Vdot(input *py.Object, other *py.Object) *py.Object
//
// view_as_complex(input) -> Tensor
//
// Returns a view of :attr:`input` as a complex tensor. For an input complex
// tensor of :attr:`size` :math:`m1, m2, \dots, mi, 2`, this function returns a
// new complex tensor of :attr:`size` :math:`m1, m2, \dots, mi` where the last
// dimension of the input tensor is expected to represent the real and imaginary
// components of complex numbers.
//
// .. warning::
//     :func:`view_as_complex` is only supported for tensors with
//     :class:`torch.dtype` ``torch.float64`` and ``torch.float32``.  The input is
//     expected to have the last dimension of :attr:`size` 2. In addition, the
//     tensor must have a `stride` of 1 for its last dimension. The strides of all
//     other dimensions must be even numbers.
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> x=torch.randn(4, 2)
//     >>> x
//     tensor([[ 1.6116, -0.5772],
//             [-1.4606, -0.9120],
//             [ 0.0786, -1.7497],
//             [-0.6561, -1.6623]])
//     >>> torch.view_as_complex(x)
//     tensor([(1.6116-0.5772j), (-1.4606-0.9120j), (0.0786-1.7497j), (-0.6561-1.6623j)])
//
//
//go:linkname ViewAsComplex py.view_as_complex
func ViewAsComplex(input *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.view_as_complex`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname ViewAsComplexCopy py.view_as_complex_copy
func ViewAsComplexCopy(__llgo_va_list ...interface{}) *py.Object
//
// view_as_real(input) -> Tensor
//
// Returns a view of :attr:`input` as a real tensor. For an input complex tensor of
// :attr:`size` :math:`m1, m2, \dots, mi`, this function returns a new
// real tensor of size :math:`m1, m2, \dots, mi, 2`, where the last dimension of size 2
// represents the real and imaginary components of complex numbers.
//
// .. warning::
//     :func:`view_as_real` is only supported for tensors with ``complex dtypes``.
//
// Args:
//     input (Tensor): the input tensor.
//
// Example::
//
//     >>> x=torch.randn(4, dtype=torch.cfloat)
//     >>> x
//     tensor([(0.4737-0.3839j), (-0.2098-0.6699j), (0.3470-0.9451j), (-0.5174-1.3136j)])
//     >>> torch.view_as_real(x)
//     tensor([[ 0.4737, -0.3839],
//             [-0.2098, -0.6699],
//             [ 0.3470, -0.9451],
//             [-0.5174, -1.3136]])
//
//
//go:linkname ViewAsReal py.view_as_real
func ViewAsReal(input *py.Object) *py.Object
//
// Performs the same operation as :func:`torch.view_as_real`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname ViewAsRealCopy py.view_as_real_copy
func ViewAsRealCopy(__llgo_va_list ...interface{}) *py.Object
//
// Performs the same operation as :func:`torch.view`, but all output tensors
// are freshly created instead of aliasing the input.
//
//
//go:linkname ViewCopy py.view_copy
func ViewCopy(__llgo_va_list ...interface{}) *py.Object
//
// vsplit(input, indices_or_sections) -> List of Tensors
//
// Splits :attr:`input`, a tensor with two or more dimensions, into multiple tensors
// vertically according to :attr:`indices_or_sections`. Each split is a view of
// :attr:`input`.
//
// This is equivalent to calling torch.tensor_split(input, indices_or_sections, dim=0)
// (the split dimension is 0), except that if :attr:`indices_or_sections` is an integer
// it must evenly divide the split dimension or a runtime error will be thrown.
//
// This function is based on NumPy's :func:`numpy.vsplit`.
//
// Args:
//     input (Tensor): tensor to split.
//     indices_or_sections (int or list or tuple of ints): See argument in :func:`torch.tensor_split`.
//
// Example::
//     >>> t = torch.arange(16.0).reshape(4,4)
//     >>> t
//     tensor([[ 0.,  1.,  2.,  3.],
//             [ 4.,  5.,  6.,  7.],
//             [ 8.,  9., 10., 11.],
//             [12., 13., 14., 15.]])
//     >>> torch.vsplit(t, 2)
//     (tensor([[0., 1., 2., 3.],
//              [4., 5., 6., 7.]]),
//      tensor([[ 8.,  9., 10., 11.],
//              [12., 13., 14., 15.]]))
//     >>> torch.vsplit(t, [3, 6])
//     (tensor([[ 0.,  1.,  2.,  3.],
//              [ 4.,  5.,  6.,  7.],
//              [ 8.,  9., 10., 11.]]),
//      tensor([[12., 13., 14., 15.]]),
//      tensor([], size=(0, 4)))
//
//
//
//go:linkname Vsplit py.vsplit
func Vsplit(input *py.Object, indicesOrSections *py.Object) *py.Object
//
// vstack(tensors, *, out=None) -> Tensor
//
// Stack tensors in sequence vertically (row wise).
//
// This is equivalent to concatenation along the first axis after all 1-D tensors have been reshaped by :func:`torch.atleast_2d`.
//
// Args:
//     tensors (sequence of Tensors): sequence of tensors to concatenate
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Example::
//
//     >>> a = torch.tensor([1, 2, 3])
//     >>> b = torch.tensor([4, 5, 6])
//     >>> torch.vstack((a,b))
//     tensor([[1, 2, 3],
//             [4, 5, 6]])
//     >>> a = torch.tensor([[1],[2],[3]])
//     >>> b = torch.tensor([[4],[5],[6]])
//     >>> torch.vstack((a,b))
//     tensor([[1],
//             [2],
//             [3],
//             [4],
//             [5],
//             [6]])
//
//
//
//
//go:linkname Vstack py.vstack
func Vstack(tensors *py.Object) *py.Object
//
// where(condition, input, other, *, out=None) -> Tensor
//
// Return a tensor of elements selected from either :attr:`input` or :attr:`other`, depending on :attr:`condition`.
//
// The operation is defined as:
//
// .. math::
//     \text{out}_i = \begin{cases}
//         \text{input}_i & \text{if } \text{condition}_i \\
//         \text{other}_i & \text{otherwise} \\
//     \end{cases}
//
// .. note::
//     The tensors :attr:`condition`, :attr:`input`, :attr:`other` must be :ref:`broadcastable <broadcasting-semantics>`.
//
// Arguments:
//     condition (BoolTensor): When True (nonzero), yield input, otherwise yield other
//     input (Tensor or Scalar): value (if :attr:`input` is a scalar) or values selected at indices
//                           where :attr:`condition` is ``True``
//     other (Tensor or Scalar): value (if :attr:`other` is a scalar) or values selected at indices
//                           where :attr:`condition` is ``False``
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//
// Returns:
//     Tensor: A tensor of shape equal to the broadcasted shape of :attr:`condition`, :attr:`input`, :attr:`other`
//
// Example::
//
//     >>> x = torch.randn(3, 2)
//     >>> y = torch.ones(3, 2)
//     >>> x
//     tensor([[-0.4620,  0.3139],
//             [ 0.3898, -0.7197],
//             [ 0.0478, -0.1657]])
//     >>> torch.where(x > 0, 1.0, 0.0)
//     tensor([[0., 1.],
//             [1., 0.],
//             [1., 0.]])
//     >>> torch.where(x > 0, x, y)
//     tensor([[ 1.0000,  0.3139],
//             [ 0.3898,  1.0000],
//             [ 0.0478,  1.0000]])
//     >>> x = torch.randn(2, 2, dtype=torch.double)
//     >>> x
//     tensor([[ 1.0779,  0.0383],
//             [-0.8785, -1.1089]], dtype=torch.float64)
//     >>> torch.where(x > 0, x, 0.)
//     tensor([[1.0779, 0.0383],
//             [0.0000, 0.0000]], dtype=torch.float64)
//
// .. function:: where(condition) -> tuple of LongTensor
//    :noindex:
//
// ``torch.where(condition)`` is identical to
// ``torch.nonzero(condition, as_tuple=True)``.
//
// .. note::
//     See also :func:`torch.nonzero`.
//
//
//go:linkname Where py.where
func Where(condition *py.Object, input *py.Object, other *py.Object) *py.Object
//
// xlogy(input, other, *, out=None) -> Tensor
//
// Alias for :func:`torch.special.xlogy`.
//
//
//go:linkname Xlogy py.xlogy
func Xlogy(input *py.Object, other *py.Object) *py.Object
// None
//
//go:linkname Xlogy_ py.xlogy_
func Xlogy_(__llgo_va_list ...interface{}) *py.Object
// None
//
//go:linkname Zero_ py.zero_
func Zero_(__llgo_va_list ...interface{}) *py.Object
//
// zeros(*size, *, out=None, dtype=None, layout=torch.strided, device=None, requires_grad=False) -> Tensor
//
// Returns a tensor filled with the scalar value `0`, with the shape defined
// by the variable argument :attr:`size`.
//
// Args:
//     size (int...): a sequence of integers defining the shape of the output tensor.
//         Can be a variable number of arguments or a collection like a list or tuple.
//
// Keyword args:
//     out (Tensor, optional): the output tensor.
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned tensor.
//         Default: if ``None``, uses a global default (see :func:`torch.set_default_tensor_type`).
//     layout (:class:`torch.layout`, optional): the desired layout of returned Tensor.
//         Default: ``torch.strided``.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, uses the current device for the default tensor type
//         (see :func:`torch.set_default_tensor_type`). :attr:`device` will be the CPU
//         for CPU tensor types and the current CUDA device for CUDA tensor types.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//
// Example::
//
//     >>> torch.zeros(2, 3)
//     tensor([[ 0.,  0.,  0.],
//             [ 0.,  0.,  0.]])
//
//     >>> torch.zeros(5)
//     tensor([ 0.,  0.,  0.,  0.,  0.])
//
//
//go:linkname Zeros py.zeros
func Zeros(__llgo_va_list ...interface{}) *py.Object
//
// zeros_like(input, *, dtype=None, layout=None, device=None, requires_grad=False, memory_format=torch.preserve_format) -> Tensor
//
// Returns a tensor filled with the scalar value `0`, with the same size as
// :attr:`input`. ``torch.zeros_like(input)`` is equivalent to
// ``torch.zeros(input.size(), dtype=input.dtype, layout=input.layout, device=input.device)``.
//
// .. warning::
//     As of 0.4, this function does not support an :attr:`out` keyword. As an alternative,
//     the old ``torch.zeros_like(input, out=output)`` is equivalent to
//     ``torch.zeros(input.size(), out=output)``.
//
// Args:
//     input (Tensor): the size of :attr:`input` will determine size of the output tensor.
//
// Keyword args:
//     dtype (:class:`torch.dtype`, optional): the desired data type of returned Tensor.
//         Default: if ``None``, defaults to the dtype of :attr:`input`.
//     layout (:class:`torch.layout`, optional): the desired layout of returned tensor.
//         Default: if ``None``, defaults to the layout of :attr:`input`.
//     device (:class:`torch.device`, optional): the desired device of returned tensor.
//         Default: if ``None``, defaults to the device of :attr:`input`.
//     requires_grad (bool, optional): If autograd should record operations on the
//         returned tensor. Default: ``False``.
//     memory_format (:class:`torch.memory_format`, optional): the desired memory format of
//         returned Tensor. Default: ``torch.preserve_format``.
//
// Example::
//
//     >>> input = torch.empty(2, 3)
//     >>> torch.zeros_like(input)
//     tensor([[ 0.,  0.,  0.],
//             [ 0.,  0.,  0.]])
//
//
//go:linkname ZerosLike py.zeros_like
func ZerosLike(input *py.Object) *py.Object
// broadcast_shapes(*shapes) -> Size
//
//     Similar to :func:`broadcast_tensors` but for shapes.
//
//     This is equivalent to
//     ``torch.broadcast_tensors(*map(torch.empty, shapes))[0].shape``
//     but avoids the need create to intermediate tensors. This is useful for
//     broadcasting tensors of common batch shape but different rightmost shape,
//     e.g. to broadcast mean vectors with covariance matrices.
//
//     Example::
//
//         >>> torch.broadcast_shapes((2,), (3, 1), (1, 1, 1))
//         torch.Size([1, 3, 2])
//
//     Args:
//         \*shapes (torch.Size): Shapes of tensors.
//
//     Returns:
//         shape (torch.Size): A shape compatible with all input shapes.
//
//     Raises:
//         RuntimeError: If shapes are incompatible.
//
//
//go:linkname BroadcastShapes py.broadcast_shapes
func BroadcastShapes(__llgo_va_list ...interface{}) *py.Object
// Computes the LU factorization of a matrix or batches of matrices
//     :attr:`A`. Returns a tuple containing the LU factorization and
//     pivots of :attr:`A`.  Pivoting is done if :attr:`pivot` is set to
//     ``True``.
//
//     .. warning::
//
//         :func:`torch.lu` is deprecated in favor of :func:`torch.linalg.lu_factor`
//         and :func:`torch.linalg.lu_factor_ex`. :func:`torch.lu` will be removed in a
//         future PyTorch release.
//         ``LU, pivots, info = torch.lu(A, compute_pivots)`` should be replaced with
//
//         .. code:: python
//
//             LU, pivots = torch.linalg.lu_factor(A, compute_pivots)
//
//         ``LU, pivots, info = torch.lu(A, compute_pivots, get_infos=True)`` should be replaced with
//
//         .. code:: python
//
//             LU, pivots, info = torch.linalg.lu_factor_ex(A, compute_pivots)
//
//     .. note::
//         * The returned permutation matrix for every matrix in the batch is
//           represented by a 1-indexed vector of size ``min(A.shape[-2], A.shape[-1])``.
//           ``pivots[i] == j`` represents that in the ``i``-th step of the algorithm,
//           the ``i``-th row was permuted with the ``j-1``-th row.
//         * LU factorization with :attr:`pivot` = ``False`` is not available
//           for CPU, and attempting to do so will throw an error. However,
//           LU factorization with :attr:`pivot` = ``False`` is available for
//           CUDA.
//         * This function does not check if the factorization was successful
//           or not if :attr:`get_infos` is ``True`` since the status of the
//           factorization is present in the third element of the return tuple.
//         * In the case of batches of square matrices with size less or equal
//           to 32 on a CUDA device, the LU factorization is repeated for
//           singular matrices due to the bug in the MAGMA library
//           (see magma issue 13).
//         * ``L``, ``U``, and ``P`` can be derived using :func:`torch.lu_unpack`.
//
//     .. warning::
//         The gradients of this function will only be finite when :attr:`A` is full rank.
//         This is because the LU decomposition is just differentiable at full rank matrices.
//         Furthermore, if :attr:`A` is close to not being full rank,
//         the gradient will be numerically unstable as it depends on the computation of :math:`L^{-1}` and :math:`U^{-1}`.
//
//     Args:
//         A (Tensor): the tensor to factor of size :math:`(*, m, n)`
//         pivot (bool, optional): controls whether pivoting is done. Default: ``True``
//         get_infos (bool, optional): if set to ``True``, returns an info IntTensor.
//                                     Default: ``False``
//         out (tuple, optional): optional output tuple. If :attr:`get_infos` is ``True``,
//                                then the elements in the tuple are Tensor, IntTensor,
//                                and IntTensor. If :attr:`get_infos` is ``False``, then the
//                                elements in the tuple are Tensor, IntTensor. Default: ``None``
//
//     Returns:
//         (Tensor, IntTensor, IntTensor (optional)): A tuple of tensors containing
//
//             - **factorization** (*Tensor*): the factorization of size :math:`(*, m, n)`
//
//             - **pivots** (*IntTensor*): the pivots of size :math:`(*, \text{min}(m, n))`.
//               ``pivots`` stores all the intermediate transpositions of rows.
//               The final permutation ``perm`` could be reconstructed by
//               applying ``swap(perm[i], perm[pivots[i] - 1])`` for ``i = 0, ..., pivots.size(-1) - 1``,
//               where ``perm`` is initially the identity permutation of :math:`m` elements
//               (essentially this is what :func:`torch.lu_unpack` is doing).
//
//             - **infos** (*IntTensor*, *optional*): if :attr:`get_infos` is ``True``, this is a tensor of
//               size :math:`(*)` where non-zero values indicate whether factorization for the matrix or
//               each minibatch has succeeded or failed
//
//     Example::
//
//         >>> # xdoctest: +REQUIRES(env:TORCH_DOCTEST_LAPACK)
//         >>> # xdoctest: +IGNORE_WANT("non-deterministic")
//         >>> A = torch.randn(2, 3, 3)
//         >>> A_LU, pivots = torch.lu(A)
//         >>> A_LU
//         tensor([[[ 1.3506,  2.5558, -0.0816],
//                  [ 0.1684,  1.1551,  0.1940],
//                  [ 0.1193,  0.6189, -0.5497]],
//
//                 [[ 0.4526,  1.2526, -0.3285],
//                  [-0.7988,  0.7175, -0.9701],
//                  [ 0.2634, -0.9255, -0.3459]]])
//         >>> pivots
//         tensor([[ 3,  3,  3],
//                 [ 3,  3,  3]], dtype=torch.int32)
//         >>> A_LU, pivots, info = torch.lu(A, get_infos=True)
//         >>> if info.nonzero().size(0) == 0:
//         ...     print('LU factorization succeeded for all samples!')
//         LU factorization succeeded for all samples!
//
//
//go:linkname Lu py.lu
func Lu(__llgo_va_list ...interface{}) *py.Object
// Performs linear Principal Component Analysis (PCA) on a low-rank
//     matrix, batches of such matrices, or sparse matrix.
//
//     This function returns a namedtuple ``(U, S, V)`` which is the
//     nearly optimal approximation of a singular value decomposition of
//     a centered matrix :math:`A` such that :math:`A = U diag(S) V^T`.
//
//     .. note:: The relation of ``(U, S, V)`` to PCA is as follows:
//
//                 - :math:`A` is a data matrix with ``m`` samples and
//                   ``n`` features
//
//                 - the :math:`V` columns represent the principal directions
//
//                 - :math:`S ** 2 / (m - 1)` contains the eigenvalues of
//                   :math:`A^T A / (m - 1)` which is the covariance of
//                   ``A`` when ``center=True`` is provided.
//
//                 - ``matmul(A, V[:, :k])`` projects data to the first k
//                   principal components
//
//     .. note:: Different from the standard SVD, the size of returned
//               matrices depend on the specified rank and q
//               values as follows:
//
//                 - :math:`U` is m x q matrix
//
//                 - :math:`S` is q-vector
//
//                 - :math:`V` is n x q matrix
//
//     .. note:: To obtain repeatable results, reset the seed for the
//               pseudorandom number generator
//
//     Args:
//
//         A (Tensor): the input tensor of size :math:`(*, m, n)`
//
//         q (int, optional): a slightly overestimated rank of
//                            :math:`A`. By default, ``q = min(6, m,
//                            n)``.
//
//         center (bool, optional): if True, center the input tensor,
//                                  otherwise, assume that the input is
//                                  centered.
//
//         niter (int, optional): the number of subspace iterations to
//                                conduct; niter must be a nonnegative
//                                integer, and defaults to 2.
//
//     References::
//
//         - Nathan Halko, Per-Gunnar Martinsson, and Joel Tropp, Finding
//           structure with randomness: probabilistic algorithms for
//           constructing approximate matrix decompositions,
//           arXiv:0909.4061 [math.NA; math.PR], 2009 (available at
//           `arXiv <http://arxiv.org/abs/0909.4061>`_).
//
//
//
//go:linkname PcaLowrank py.pca_lowrank
func PcaLowrank(A *py.Object, q *py.Object, center *py.Object, niter *py.Object) *py.Object
// Return the singular value decomposition ``(U, S, V)`` of a matrix,
//     batches of matrices, or a sparse matrix :math:`A` such that
//     :math:`A \approx U diag(S) V^T`. In case :math:`M` is given, then
//     SVD is computed for the matrix :math:`A - M`.
//
//     .. note:: The implementation is based on the Algorithm 5.1 from
//               Halko et al, 2009.
//
//     .. note:: To obtain repeatable results, reset the seed for the
//               pseudorandom number generator
//
//     .. note:: The input is assumed to be a low-rank matrix.
//
//     .. note:: In general, use the full-rank SVD implementation
//               :func:`torch.linalg.svd` for dense matrices due to its 10-fold
//               higher performance characteristics. The low-rank SVD
//               will be useful for huge sparse matrices that
//               :func:`torch.linalg.svd` cannot handle.
//
//     Args::
//         A (Tensor): the input tensor of size :math:`(*, m, n)`
//
//         q (int, optional): a slightly overestimated rank of A.
//
//         niter (int, optional): the number of subspace iterations to
//                                conduct; niter must be a nonnegative
//                                integer, and defaults to 2
//
//         M (Tensor, optional): the input tensor's mean of size
//                               :math:`(*, 1, n)`.
//
//     References::
//         - Nathan Halko, Per-Gunnar Martinsson, and Joel Tropp, Finding
//           structure with randomness: probabilistic algorithms for
//           constructing approximate matrix decompositions,
//           arXiv:0909.4061 [math.NA; math.PR], 2009 (available at
//           `arXiv <https://arxiv.org/abs/0909.4061>`_).
//
//
//
//go:linkname SvdLowrank py.svd_lowrank
func SvdLowrank(A *py.Object, q *py.Object, niter *py.Object, M *py.Object) *py.Object
// unique(input, sorted=True, return_inverse=False, return_counts=False, dim=None) -> Tuple[Tensor, Tensor, Tensor]
//
//     Returns the unique elements of the input tensor.
//
//     .. note:: This function is different from :func:`torch.unique_consecutive` in the sense that
//         this function also eliminates non-consecutive duplicate values.
//
//     .. note:: Currently in the CUDA implementation and the CPU implementation,
//         `torch.unique` always sort the tensor at the beginning regardless of the `sort` argument.
//         Sorting could be slow, so if your input tensor is already sorted, it is recommended to use
//         :func:`torch.unique_consecutive` which avoids the sorting.
//
//     Args:
//         input (Tensor): the input tensor
//         sorted (bool): Whether to sort the unique elements in ascending order
//             before returning as output.
//         return_inverse (bool): Whether to also return the indices for where
//             elements in the original input ended up in the returned unique list.
//         return_counts (bool): Whether to also return the counts for each unique
//             element.
//         dim (int, optional): the dimension to operate upon. If ``None``, the
//             unique of the flattened input is returned. Otherwise, each of the
//             tensors indexed by the given dimension is treated as one of the
//             elements to apply the unique operation upon. See examples for more
//             details. Default: ``None``
//
//     Returns:
//         (Tensor, Tensor (optional), Tensor (optional)): A tensor or a tuple of tensors containing
//
//             - **output** (*Tensor*): the output list of unique scalar elements.
//             - **inverse_indices** (*Tensor*): (optional) if
//               :attr:`return_inverse` is True, there will be an additional
//               returned tensor (same shape as input) representing the indices
//               for where elements in the original input map to in the output;
//               otherwise, this function will only return a single tensor.
//             - **counts** (*Tensor*): (optional) if
//               :attr:`return_counts` is True, there will be an additional
//               returned tensor (same shape as output or output.size(dim),
//               if dim was specified) representing the number of occurrences
//               for each unique value or tensor.
//
//     Example::
//
//         >>> output = torch.unique(torch.tensor([1, 3, 2, 3], dtype=torch.long))
//         >>> output
//         tensor([1, 2, 3])
//
//         >>> output, inverse_indices = torch.unique(
//         ...     torch.tensor([1, 3, 2, 3], dtype=torch.long), sorted=True, return_inverse=True)
//         >>> output
//         tensor([1, 2, 3])
//         >>> inverse_indices
//         tensor([0, 2, 1, 2])
//
//         >>> output, inverse_indices = torch.unique(
//         ...     torch.tensor([[1, 3], [2, 3]], dtype=torch.long), sorted=True, return_inverse=True)
//         >>> output
//         tensor([1, 2, 3])
//         >>> inverse_indices
//         tensor([[0, 2],
//                 [1, 2]])
//
//         >>> a = torch.tensor([
//         ...     [
//         ...         [1, 1, 0, 0],
//         ...         [1, 1, 0, 0],
//         ...         [0, 0, 1, 1],
//         ...     ],
//         ...     [
//         ...         [0, 0, 1, 1],
//         ...         [0, 0, 1, 1],
//         ...         [1, 1, 1, 1],
//         ...     ],
//         ...     [
//         ...         [1, 1, 0, 0],
//         ...         [1, 1, 0, 0],
//         ...         [0, 0, 1, 1],
//         ...     ],
//         ... ])
//
//         >>> # If we call `torch.unique(a, dim=0)`, each of the tensors `a[idx, :, :]`
//         >>> # will be compared. We can see that `a[0, :, :]` and `a[2, :, :]` match
//         >>> # each other, so one of them will be removed.
//         >>> (a[0, :, :] == a[2, :, :]).all()
//         tensor(True)
//         >>> a_unique_dim0 = torch.unique(a, dim=0)
//         >>> a_unique_dim0
//         tensor([[[0, 0, 1, 1],
//                  [0, 0, 1, 1],
//                  [1, 1, 1, 1]],
//                 [[1, 1, 0, 0],
//                  [1, 1, 0, 0],
//                  [0, 0, 1, 1]]])
//
//         >>> # Notice which sub-tensors from `a` match with the sub-tensors from
//         >>> # `a_unique_dim0`:
//         >>> (a_unique_dim0[0, :, :] == a[1, :, :]).all()
//         tensor(True)
//         >>> (a_unique_dim0[1, :, :] == a[0, :, :]).all()
//         tensor(True)
//
//         >>> # For `torch.unique(a, dim=1)`, each of the tensors `a[:, idx, :]` are
//         >>> # compared. `a[:, 0, :]` and `a[:, 1, :]` match each other, so one of
//         >>> # them will be removed.
//         >>> (a[:, 0, :] == a[:, 1, :]).all()
//         tensor(True)
//         >>> torch.unique(a, dim=1)
//         tensor([[[0, 0, 1, 1],
//                  [1, 1, 0, 0]],
//                 [[1, 1, 1, 1],
//                  [0, 0, 1, 1]],
//                 [[0, 0, 1, 1],
//                  [1, 1, 0, 0]]])
//
//         >>> # For `torch.unique(a, dim=2)`, the tensors `a[:, :, idx]` are compared.
//         >>> # `a[:, :, 0]` and `a[:, :, 1]` match each other. Also, `a[:, :, 2]` and
//         >>> # `a[:, :, 3]` match each other as well. So in this case, two of the
//         >>> # sub-tensors will be removed.
//         >>> (a[:, :, 0] == a[:, :, 1]).all()
//         tensor(True)
//         >>> (a[:, :, 2] == a[:, :, 3]).all()
//         tensor(True)
//         >>> torch.unique(a, dim=2)
//         tensor([[[0, 1],
//                  [0, 1],
//                  [1, 0]],
//                 [[1, 0],
//                  [1, 0],
//                  [1, 1]],
//                 [[0, 1],
//                  [0, 1],
//                  [1, 0]]])
//
//
//go:linkname Unique py.unique
func Unique(input *py.Object, sorted *py.Object, returnInverse *py.Object, returnCounts *py.Object, dim *py.Object) *py.Object
// Converts a tensor of flat indices into a tuple of coordinate tensors that
//     index into an arbitrary tensor of the specified shape.
//
//     Args:
//         indices (Tensor): An integer tensor containing indices into the
//             flattened version of an arbitrary tensor of shape :attr:`shape`.
//             All elements must be in the range ``[0, prod(shape) - 1]``.
//
//         shape (int, sequence of ints, or torch.Size): The shape of the arbitrary
//             tensor. All elements must be non-negative.
//
//     Returns:
//         tuple of Tensors: Each ``i``-th tensor in the ouput corresponds with
//         dimension ``i`` of :attr:`shape`. Each tensor has the same shape as
//         ``indices`` and contains one index into dimension ``i`` for each of the
//         flat indices given by ``indices``.
//
//     Example::
//
//         >>> import torch
//         >>> torch.unravel_index(torch.tensor(4), (3, 2))
//         (tensor(2),
//          tensor(0))
//
//         >>> torch.unravel_index(torch.tensor([4, 1]), (3, 2))
//         (tensor([2, 0]),
//          tensor([0, 1]))
//
//         >>> torch.unravel_index(torch.tensor([0, 1, 2, 3, 4, 5]), (3, 2))
//         (tensor([0, 0, 1, 1, 2, 2]),
//          tensor([0, 1, 0, 1, 0, 1]))
//
//         >>> torch.unravel_index(torch.tensor([1234, 5678]), (10, 10, 10, 10))
//         (tensor([1, 5]),
//          tensor([2, 6]),
//          tensor([3, 7]),
//          tensor([4, 8]))
//
//         >>> torch.unravel_index(torch.tensor([[1234], [5678]]), (10, 10, 10, 10))
//         (tensor([[1], [5]]),
//          tensor([[2], [6]]),
//          tensor([[3], [7]]),
//          tensor([[4], [8]]))
//
//         >>> torch.unravel_index(torch.tensor([[1234], [5678]]), (100, 100))
//         (tensor([[12], [56]]),
//          tensor([[34], [78]]))
//
//
//go:linkname UnravelIndex py.unravel_index
func UnravelIndex(indices *py.Object, shape *py.Object) *py.Object
// Returns whether PyTorch was built with _GLIBCXX_USE_CXX11_ABI=1
//
//go:linkname CompiledWithCxx11Abi py.compiled_with_cxx11_abi
func CompiledWithCxx11Abi() *py.Object
// Find the k largest (or smallest) eigenvalues and the corresponding
//     eigenvectors of a symmetric positive definite generalized
//     eigenvalue problem using matrix-free LOBPCG methods.
//
//     This function is a front-end to the following LOBPCG algorithms
//     selectable via `method` argument:
//
//       `method="basic"` - the LOBPCG method introduced by Andrew
//       Knyazev, see [Knyazev2001]. A less robust method, may fail when
//       Cholesky is applied to singular input.
//
//       `method="ortho"` - the LOBPCG method with orthogonal basis
//       selection [StathopoulosEtal2002]. A robust method.
//
//     Supported inputs are dense, sparse, and batches of dense matrices.
//
//     .. note:: In general, the basic method spends least time per
//       iteration. However, the robust methods converge much faster and
//       are more stable. So, the usage of the basic method is generally
//       not recommended but there exist cases where the usage of the
//       basic method may be preferred.
//
//     .. warning:: The backward method does not support sparse and complex inputs.
//       It works only when `B` is not provided (i.e. `B == None`).
//       We are actively working on extensions, and the details of
//       the algorithms are going to be published promptly.
//
//     .. warning:: While it is assumed that `A` is symmetric, `A.grad` is not.
//       To make sure that `A.grad` is symmetric, so that `A - t * A.grad` is symmetric
//       in first-order optimization routines, prior to running `lobpcg`
//       we do the following symmetrization map: `A -> (A + A.t()) / 2`.
//       The map is performed only when the `A` requires gradients.
//
//     Args:
//
//       A (Tensor): the input tensor of size :math:`(*, m, m)`
//
//       B (Tensor, optional): the input tensor of size :math:`(*, m,
//                   m)`. When not specified, `B` is interpreted as
//                   identity matrix.
//
//       X (tensor, optional): the input tensor of size :math:`(*, m, n)`
//                   where `k <= n <= m`. When specified, it is used as
//                   initial approximation of eigenvectors. X must be a
//                   dense tensor.
//
//       iK (tensor, optional): the input tensor of size :math:`(*, m,
//                   m)`. When specified, it will be used as preconditioner.
//
//       k (integer, optional): the number of requested
//                   eigenpairs. Default is the number of :math:`X`
//                   columns (when specified) or `1`.
//
//       n (integer, optional): if :math:`X` is not specified then `n`
//                   specifies the size of the generated random
//                   approximation of eigenvectors. Default value for `n`
//                   is `k`. If :math:`X` is specified, the value of `n`
//                   (when specified) must be the number of :math:`X`
//                   columns.
//
//       tol (float, optional): residual tolerance for stopping
//                  criterion. Default is `feps ** 0.5` where `feps` is
//                  smallest non-zero floating-point number of the given
//                  input tensor `A` data type.
//
//       largest (bool, optional): when True, solve the eigenproblem for
//                  the largest eigenvalues. Otherwise, solve the
//                  eigenproblem for smallest eigenvalues. Default is
//                  `True`.
//
//       method (str, optional): select LOBPCG method. See the
//                  description of the function above. Default is
//                  "ortho".
//
//       niter (int, optional): maximum number of iterations. When
//                  reached, the iteration process is hard-stopped and
//                  the current approximation of eigenpairs is returned.
//                  For infinite iteration but until convergence criteria
//                  is met, use `-1`.
//
//       tracker (callable, optional) : a function for tracing the
//                  iteration process. When specified, it is called at
//                  each iteration step with LOBPCG instance as an
//                  argument. The LOBPCG instance holds the full state of
//                  the iteration process in the following attributes:
//
//                    `iparams`, `fparams`, `bparams` - dictionaries of
//                    integer, float, and boolean valued input
//                    parameters, respectively
//
//                    `ivars`, `fvars`, `bvars`, `tvars` - dictionaries
//                    of integer, float, boolean, and Tensor valued
//                    iteration variables, respectively.
//
//                    `A`, `B`, `iK` - input Tensor arguments.
//
//                    `E`, `X`, `S`, `R` - iteration Tensor variables.
//
//                  For instance:
//
//                    `ivars["istep"]` - the current iteration step
//                    `X` - the current approximation of eigenvectors
//                    `E` - the current approximation of eigenvalues
//                    `R` - the current residual
//                    `ivars["converged_count"]` - the current number of converged eigenpairs
//                    `tvars["rerr"]` - the current state of convergence criteria
//
//                  Note that when `tracker` stores Tensor objects from
//                  the LOBPCG instance, it must make copies of these.
//
//                  If `tracker` sets `bvars["force_stop"] = True`, the
//                  iteration process will be hard-stopped.
//
//       ortho_iparams, ortho_fparams, ortho_bparams (dict, optional):
//                  various parameters to LOBPCG algorithm when using
//                  `method="ortho"`.
//
//     Returns:
//
//       E (Tensor): tensor of eigenvalues of size :math:`(*, k)`
//
//       X (Tensor): tensor of eigenvectors of size :math:`(*, m, k)`
//
//     References:
//
//       [Knyazev2001] Andrew V. Knyazev. (2001) Toward the Optimal
//       Preconditioned Eigensolver: Locally Optimal Block Preconditioned
//       Conjugate Gradient Method. SIAM J. Sci. Comput., 23(2),
//       517-541. (25 pages)
//       https://epubs.siam.org/doi/abs/10.1137/S1064827500366124
//
//       [StathopoulosEtal2002] Andreas Stathopoulos and Kesheng
//       Wu. (2002) A Block Orthogonalization Procedure with Constant
//       Synchronization Requirements. SIAM J. Sci. Comput., 23(6),
//       2165-2182. (18 pages)
//       https://epubs.siam.org/doi/10.1137/S1064827500370883
//
//       [DuerschEtal2018] Jed A. Duersch, Meiyue Shao, Chao Yang, Ming
//       Gu. (2018) A Robust and Efficient Implementation of LOBPCG.
//       SIAM J. Sci. Comput., 40(5), C655-C676. (22 pages)
//       https://epubs.siam.org/doi/abs/10.1137/17M1129830
//
//
//
//go:linkname Lobpcg py.lobpcg
func Lobpcg(A *py.Object, k *py.Object, B *py.Object, X *py.Object, n *py.Object, iK *py.Object, niter *py.Object, tol *py.Object, largest *py.Object, method *py.Object, tracker *py.Object, orthoIparams *py.Object, orthoFparams *py.Object, orthoBparams *py.Object) *py.Object
// from_dlpack(ext_tensor) -> Tensor
//
//     Converts a tensor from an external library into a ``torch.Tensor``.
//
//     The returned PyTorch tensor will share the memory with the input tensor
//     (which may have come from another library). Note that in-place operations
//     will therefore also affect the data of the input tensor. This may lead to
//     unexpected issues (e.g., other libraries may have read-only flags or
//     immutable data structures), so the user should only do this if they know
//     for sure that this is fine.
//
//     Args:
//         ext_tensor (object with ``__dlpack__`` attribute, or a DLPack capsule):
//             The tensor or DLPack capsule to convert.
//
//             If ``ext_tensor`` is a tensor (or ndarray) object, it must support
//             the ``__dlpack__`` protocol (i.e., have a ``ext_tensor.__dlpack__``
//             method). Otherwise ``ext_tensor`` may be a DLPack capsule, which is
//             an opaque ``PyCapsule`` instance, typically produced by a
//             ``to_dlpack`` function or method.
//
//     Examples::
//
//         >>> import torch.utils.dlpack
//         >>> t = torch.arange(4)
//
//         # Convert a tensor directly (supported in PyTorch >= 1.10)
//         >>> t2 = torch.from_dlpack(t)
//         >>> t2[:2] = -1  # show that memory is shared
//         >>> t2
//         tensor([-1, -1,  2,  3])
//         >>> t
//         tensor([-1, -1,  2,  3])
//
//         # The old-style DLPack usage, with an intermediate capsule object
//         >>> capsule = torch.utils.dlpack.to_dlpack(t)
//         >>> capsule
//         <capsule object "dltensor" at ...>
//         >>> t3 = torch.from_dlpack(capsule)
//         >>> t3
//         tensor([-1, -1,  2,  3])
//         >>> t3[0] = -9  # now we're sharing memory between 3 tensors
//         >>> t3
//         tensor([-9, -1,  2,  3])
//         >>> t2
//         tensor([-9, -1,  2,  3])
//         >>> t
//         tensor([-9, -1,  2,  3])
//
//
//
//go:linkname FromDlpack py.from_dlpack
func FromDlpack(extTensor *py.Object) *py.Object
// to_dlpack(tensor) -> PyCapsule
//
// Returns an opaque object (a "DLPack capsule") representing the tensor.
//
// .. note::
//   ``to_dlpack`` is a legacy DLPack interface. The capsule it returns
//   cannot be used for anything in Python other than use it as input to
//   ``from_dlpack``. The more idiomatic use of DLPack is to call
//   ``from_dlpack`` directly on the tensor object - this works when that
//   object has a ``__dlpack__`` method, which PyTorch and most other
//   libraries indeed have now.
//
// .. warning::
//   Only call ``from_dlpack`` once per capsule produced with ``to_dlpack``.
//   Behavior when a capsule is consumed multiple times is undefined.
//
// Args:
//     tensor: a tensor to be exported
//
// The DLPack capsule shares the tensor's memory.
//
//
//go:linkname ToDlpack py.to_dlpack
func ToDlpack(tensor *py.Object) *py.Object
// None
//
//go:linkname MatrixRank py.matrix_rank
func MatrixRank(input *py.Object, tol *py.Object, symmetric *py.Object) *py.Object
// None
//
//go:linkname Eig py.eig
func Eig(self *py.Object, eigenvectors *py.Object) *py.Object
// None
//
//go:linkname Solve py.solve
func Solve(input *py.Object, A *py.Object) *py.Object
// None
//
//go:linkname Lstsq py.lstsq
func Lstsq(input *py.Object, A *py.Object) *py.Object
// None
//
//go:linkname Symeig py.symeig
func Symeig(input *py.Object, eigenvectors *py.Object, upper *py.Object) *py.Object
//
//     Optimizes given model/function using TorchDynamo and specified backend.
//
//     Concretely, for every frame executed within the compiled region, we will attempt
//     to compile it and cache the compiled result on the code object for future
//     use.  A single frame may be compiled multiple times if previous compiled
//     results are not applicable for subsequent calls (this is called a "guard
//     failure), you can use TORCH_LOGS=guards to debug these situations.
//     Multiple compiled results can be associated with a frame up to
//     ``torch._dynamo.config.cache_size_limit``, which defaults to 64; at which
//     point we will fall back to eager.  Note that compile caches are per
//     *code object*, not frame; if you dynamically create multiple copies of a
//     function, they will all share the same code cache.
//
//     Args:
//        model (Callable): Module/function to optimize
//        fullgraph (bool): If False (default), torch.compile attempts to discover compileable regions
//         in the function that it will optimize. If True, then we require that the entire function be
//         capturable into a single graph. If this is not possible (that is, if there are graph breaks),
//         then this will raise an error.
//        dynamic (bool or None): Use dynamic shape tracing.  When this is True, we will up-front attempt
//         to generate a kernel that is as dynamic as possible to avoid recompilations when
//         sizes change.  This may not always work as some operations/optimizations will
//         force specialization; use TORCH_LOGS=dynamic to debug overspecialization.
//         When this is False, we will NEVER generate dynamic kernels, we will always specialize.
//         By default (None), we automatically detect if dynamism has occurred and compile a more
//         dynamic kernel upon recompile.
//        backend (str or Callable): backend to be used
//
//         - "inductor" is the default backend, which is a good balance between performance and overhead
//
//         - Non experimental in-tree backends can be seen with `torch._dynamo.list_backends()`
//
//         - Experimental or debug in-tree backends can be seen with `torch._dynamo.list_backends(None)`
//
//         - To register an out-of-tree custom backend: https://pytorch.org/docs/main/compile/custom-backends.html
//        mode (str): Can be either "default", "reduce-overhead", "max-autotune" or "max-autotune-no-cudagraphs"
//
//         - "default" is the default mode, which is a good balance between performance and overhead
//
//         - "reduce-overhead" is a mode that reduces the overhead of python with CUDA graphs,
//           useful for small batches.  Reduction of overhead can come at the cost of more memory
//           usage, as we will cache the workspace memory required for the invocation so that we
//           do not have to reallocate it on subsequent runs.  Reduction of overhead is not guaranteed
//           to work; today, we only reduce overhead for CUDA only graphs which do not mutate inputs.
//           There are other circumstances where CUDA graphs are not applicable; use TORCH_LOG=perf_hints
//           to debug.
//
//         - "max-autotune" is a mode that leverages Triton based matrix multiplications and convolutions
//           It enables CUDA graphs by default.
//
//         - "max-autotune-no-cudagraphs" is a mode similar to "max-autotune" but without CUDA graphs
//
//         - To see the exact configs that each mode sets you can call `torch._inductor.list_mode_options()`
//
//        options (dict): A dictionary of options to pass to the backend. Some notable ones to try out are
//
//         - `epilogue_fusion` which fuses pointwise ops into templates. Requires `max_autotune` to also be set
//
//         - `max_autotune` which will profile to pick the best matmul configuration
//
//         - `fallback_random` which is useful when debugging accuracy issues
//
//         - `shape_padding` which pads matrix shapes to better align loads on GPUs especially for tensor cores
//
//         - `triton.cudagraphs` which will reduce the overhead of python with CUDA graphs
//
//         - `trace.enabled` which is the most useful debugging flag to turn on
//
//         - `trace.graph_diagram` which will show you a picture of your graph after fusion
//
//         - For inductor you can see the full list of configs that it supports by calling `torch._inductor.list_options()`
//        disable (bool): Turn torch.compile() into a no-op for testing
//
//     Example::
//
//         @torch.compile(options={"triton.cudagraphs": True}, fullgraph=True)
//         def foo(x):
//             return torch.sin(x) + torch.cos(x)
//
//
//
//go:linkname Compile py.compile
func Compile(model *py.Object) *py.Object
//
//     Conditionally applies `true_fn` or `false_fn`.
//
//     .. warning::
//         `torch.cond` is a prototype feature in PyTorch. It has limited support for input and output types and
//         doesn't support training currently. Please look forward to a more stable implementation in a future version of PyTorch.
//         Read more about feature classification at: https://pytorch.org/blog/pytorch-feature-classification-changes/#prototype
//
//     `cond` is structured control flow operator. That is, it is like a Python if-statement,
//     but has restrictions on `true_fn`, `false_fn`, and `operands` that enable it to be
//     capturable using torch.compile and torch.export.
//
//     Assuming the constraints on `cond`'s arguments are met, `cond` is equivalent to the following::
//
//         def cond(pred, true_branch, false_branch, operands):
//             if pred:
//                 return true_branch(*operands)
//             else:
//                 return false_branch(*operands)
//
//     Args:
//         pred (Union[bool, torch.Tensor]): A boolean expression or a tensor with one element,
//           indicating which branch function to apply.
//
//         true_fn (Callable): A callable function (a -> b) that is within the
//           scope that is being traced.
//
//         false_fn (Callable): A callable function (a -> b) that is within the
//           scope that is being traced. The true branch and false branch must
//           have consistent input and outputs, meaning the inputs have to be
//           the same, and the outputs have to be the same type and shape.
//
//         operands (Tuple of possibly nested dict/list/tuple of torch.Tensor): A tuple of inputs to the true/false functions.
//
//     Example::
//
//         def true_fn(x: torch.Tensor):
//             return x.cos()
//         def false_fn(x: torch.Tensor):
//             return x.sin()
//         return cond(x.shape[0] > 4, true_fn, false_fn, (x,))
//
//     Restrictions:
//         - The conditional statement (aka `pred`) must meet one of the following constraints:
//
//           - It's a `torch.Tensor` with only one element, and torch.bool dtype
//
//           - It's a boolean expression, e.g. `x.shape[0] > 10` or `x.dim() > 1 and x.shape[1] > 10`
//
//         - The branch function (aka `true_fn`/`false_fn`) must meet all of the following constraints:
//
//           - The function signature must match with operands.
//
//           - The function must return a tensor with the same metadata, e.g. shape,
//             dtype, etc.
//
//           - The function cannot have in-place mutations on inputs or global variables.
//             (Note: in-place tensor operations such as `add_` for intermediate results
//             are allowed in a branch)
//
//     .. warning::
//         Temporal Limitations:
//
//         - `cond` only supports **inference** right now. Autograd will be supported in the future.
//
//         - The **output** of branches must be a **single Tensor**. Pytree of tensors will be supported in the future.
//
//
//
//go:linkname Cond py.cond
func Cond(pred *py.Object, trueFn *py.Object, falseFn *py.Object, operands *py.Object) *py.Object
//
//     vmap is the vectorizing map; ``vmap(func)`` returns a new function that
//     maps ``func`` over some dimension of the inputs. Semantically, vmap
//     pushes the map into PyTorch operations called by ``func``, effectively
//     vectorizing those operations.
//
//     vmap is useful for handling batch dimensions: one can write a function
//     ``func`` that runs on examples and then lift it to a function that can
//     take batches of examples with ``vmap(func)``. vmap can also be used to
//     compute batched gradients when composed with autograd.
//
//     .. note::
//         :func:`torch.vmap` is aliased to :func:`torch.func.vmap` for
//         convenience. Use whichever one you'd like.
//
//     Args:
//         func (function): A Python function that takes one or more arguments.
//             Must return one or more Tensors.
//         in_dims (int or nested structure): Specifies which dimension of the
//             inputs should be mapped over. ``in_dims`` should have a
//             structure like the inputs. If the ``in_dim`` for a particular
//             input is None, then that indicates there is no map dimension.
//             Default: 0.
//         out_dims (int or Tuple[int]): Specifies where the mapped dimension
//             should appear in the outputs. If ``out_dims`` is a Tuple, then
//             it should have one element per output. Default: 0.
//         randomness (str): Specifies whether the randomness in this
//             vmap should be the same or different across batches. If 'different',
//             the randomness for each batch will be different. If 'same', the
//             randomness will be the same across batches. If 'error', any calls to
//             random functions will error. Default: 'error'. WARNING: this flag
//             only applies to random PyTorch operations and does not apply to
//             Python's random module or numpy randomness.
//         chunk_size (None or int): If None (default), apply a single vmap over inputs.
//             If not None, then compute the vmap :attr:`chunk_size` samples at a time.
//             Note that :attr:`chunk_size=1` is equivalent to computing the vmap with a for-loop.
//             If you run into memory issues computing the vmap, please try a non-None chunk_size.
//
//     Returns:
//         Returns a new "batched" function. It takes the same inputs as
//         ``func``, except each input has an extra dimension at the index
//         specified by ``in_dims``. It takes returns the same outputs as
//         ``func``, except each output has an extra dimension at the index
//         specified by ``out_dims``.
//
//     .. warning:
//         :func:`vmap` works best with functional-style code. Please do not
//         perform any side-effects in ``func``, with the exception of
//         in-place PyTorch operations. Examples of side-effects include mutating
//         Python data structures and assigning values to variables not captured
//         in ``func``.
//
//     One example of using :func:`vmap` is to compute batched dot products. PyTorch
//     doesn't provide a batched ``torch.dot`` API; instead of unsuccessfully
//     rummaging through docs, use :func:`vmap` to construct a new function.
//
//         >>> torch.dot                            # [D], [D] -> []
//         >>> batched_dot = torch.func.vmap(torch.dot)  # [N, D], [N, D] -> [N]
//         >>> x, y = torch.randn(2, 5), torch.randn(2, 5)
//         >>> batched_dot(x, y)
//
//     :func:`vmap` can be helpful in hiding batch dimensions, leading to a simpler
//     model authoring experience.
//
//         >>> batch_size, feature_size = 3, 5
//         >>> weights = torch.randn(feature_size, requires_grad=True)
//         >>>
//         >>> def model(feature_vec):
//         >>>     # Very simple linear model with activation
//         >>>     return feature_vec.dot(weights).relu()
//         >>>
//         >>> examples = torch.randn(batch_size, feature_size)
//         >>> result = torch.vmap(model)(examples)
//
//     :func:`vmap` can also help vectorize computations that were previously difficult
//     or impossible to batch. One example is higher-order gradient computation.
//     The PyTorch autograd engine computes vjps (vector-Jacobian products).
//     Computing a full Jacobian matrix for some function f: R^N -> R^N usually
//     requires N calls to ``autograd.grad``, one per Jacobian row. Using :func:`vmap`,
//     we can vectorize the whole computation, computing the Jacobian in a single
//     call to ``autograd.grad``.
//
//         >>> # Setup
//         >>> N = 5
//         >>> f = lambda x: x ** 2
//         >>> x = torch.randn(N, requires_grad=True)
//         >>> y = f(x)
//         >>> I_N = torch.eye(N)
//         >>>
//         >>> # Sequential approach
//         >>> jacobian_rows = [torch.autograd.grad(y, x, v, retain_graph=True)[0]
//         >>>                  for v in I_N.unbind()]
//         >>> jacobian = torch.stack(jacobian_rows)
//         >>>
//         >>> # vectorized gradient computation
//         >>> def get_vjp(v):
//         >>>     return torch.autograd.grad(y, x, v)
//         >>> jacobian = torch.vmap(get_vjp)(I_N)
//
//     :func:`vmap` can also be nested, producing an output with multiple batched dimensions
//
//         >>> torch.dot                            # [D], [D] -> []
//         >>> batched_dot = torch.vmap(torch.vmap(torch.dot))  # [N1, N0, D], [N1, N0, D] -> [N1, N0]
//         >>> x, y = torch.randn(2, 3, 5), torch.randn(2, 3, 5)
//         >>> batched_dot(x, y) # tensor of size [2, 3]
//
//     If the inputs are not batched along the first dimension, ``in_dims`` specifies
//     the dimension that each inputs are batched along as
//
//         >>> torch.dot                            # [N], [N] -> []
//         >>> batched_dot = torch.vmap(torch.dot, in_dims=1)  # [N, D], [N, D] -> [D]
//         >>> x, y = torch.randn(2, 5), torch.randn(2, 5)
//         >>> batched_dot(x, y)   # output is [5] instead of [2] if batched along the 0th dimension
//
//     If there are multiple inputs each of which is batched along different dimensions,
//     ``in_dims`` must be a tuple with the batch dimension for each input as
//
//         >>> torch.dot                            # [D], [D] -> []
//         >>> batched_dot = torch.vmap(torch.dot, in_dims=(0, None))  # [N, D], [D] -> [N]
//         >>> x, y = torch.randn(2, 5), torch.randn(5)
//         >>> batched_dot(x, y) # second arg doesn't have a batch dim because in_dim[1] was None
//
//     If the input is a Python struct, ``in_dims`` must be a tuple containing a struct
//     matching the shape of the input:
//
//         >>> f = lambda dict: torch.dot(dict['x'], dict['y'])
//         >>> x, y = torch.randn(2, 5), torch.randn(5)
//         >>> input = {'x': x, 'y': y}
//         >>> batched_dot = torch.vmap(f, in_dims=({'x': 0, 'y': None},))
//         >>> batched_dot(input)
//
//     By default, the output is batched along the first dimension. However, it can be batched
//     along any dimension by using ``out_dims``
//
//         >>> f = lambda x: x ** 2
//         >>> x = torch.randn(2, 5)
//         >>> batched_pow = torch.vmap(f, out_dims=1)
//         >>> batched_pow(x) # [5, 2]
//
//     For any function that uses kwargs, the returned function will not batch the kwargs but will
//     accept kwargs
//
//         >>> x = torch.randn([2, 5])
//         >>> def fn(x, scale=4.):
//         >>>   return x * scale
//         >>>
//         >>> batched_pow = torch.vmap(fn)
//         >>> assert torch.allclose(batched_pow(x), x * 4)
//         >>> batched_pow(x, scale=x) # scale is not batched, output has shape [2, 2, 5]
//
//     .. note::
//         vmap does not provide general autobatching or handle variable-length
//         sequences out of the box.
//
//
//go:linkname Vmap py.vmap
func Vmap(func_ *py.Object, inDims *py.Object, outDims *py.Object, randomness *py.Object) *py.Object
