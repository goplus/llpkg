package main

import (
	"fmt"
	"unsafe"

	"github.com/goplus/llgo/c" // 提供 C I/O 函数等封装
	"github.com/goplus/llpkg/bzip2"
	// 假设 bzip2 包在同一项目的 bzip2 目录下
)

func CompressFile(inPath, outPath string) error {
	// 打开输入文件 (只读)
	inPathC := c.AllocaCStr(inPath)

	inFile := c.Fopen(inPathC, c.Str("rb"))
	if inFile == nil {
		return fmt.Errorf("failed to open input file: %s", inPath)
	}
	defer c.Fclose(inFile)

	// 打开输出文件 (只写)
	outPathC := c.AllocaCStr(outPath)
	outFile := c.Fopen(outPathC, c.Str("wb"))
	if outFile == nil {
		return fmt.Errorf("failed to open output file: %s", outPath)
	}
	defer c.Fclose(outFile)

	// 创建 bzip2 压缩流
	var bzerr c.Int
	// blockSize100k=9 表示最大压缩，verbosity=0，workFactor=30
	bzfile := bzip2.BzWriteOpen(&bzerr, outFile, 9, 0, 30)
	if bzfile == nil || bzerr != bzip2.BZ_OK {
		return fmt.Errorf("BzWriteOpen error, code=%d", bzerr)
	}

	// 循环读取输入文件并写入到压缩流
	buf := make([]byte, 4096)
	for {
		n := c.Fread(unsafe.Pointer(&buf[0]), 1, uintptr(len(buf)), inFile)
		if n == 0 {
			break // 读取完毕或发生错误
		}

		bzip2.BzWrite(&bzerr, bzfile, unsafe.Pointer(&buf[0]), c.Int(n))
		if bzerr != bzip2.BZ_OK {
			return fmt.Errorf("BzWrite error, code=%d", bzerr)
		}

		// 如果实际读取到的数据比缓冲区小，说明文件读完了
		if n < uintptr(len(buf)) {
			break
		}
	}

	// 关闭 bzip2 压缩流
	bzip2.BzWriteClose(&bzerr, bzfile, 0, nil, nil)
	if bzerr != bzip2.BZ_OK {
		return fmt.Errorf("BzWriteClose error, code=%d", bzerr)
	}

	return nil
}

func main() {
	CompressFile("test.txt", "../decompress/test.bz2")
}
