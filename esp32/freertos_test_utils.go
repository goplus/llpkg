package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type TestFunctionT func(c.Pointer)

/**
 * @brief Run a test function on each core
 *
 * This function will internally create a task pinned to each core, where each task will call the provided test
 * function. This function will block until all cores finish executing the test function.
 *
 * @param pxTestCode Test function
 * @param pvTestCodeArg Argument provided to test function
 * @param ulStackDepth Stack depth of the created tasks
 * @param uxPriority Priority of the created tasks
 */
//go:linkname VTestOnAllCores C.vTestOnAllCores
func VTestOnAllCores(pxTestCode TestFunctionT, pvTestCodeArg c.Pointer, ulStackDepth c.Uint32T, uxPriority UBaseTypeT)
