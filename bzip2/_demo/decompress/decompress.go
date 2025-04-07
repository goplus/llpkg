package main

import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/bzip2"
)

func DecompressFile(inPath, outPath string) error {
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

	// 创建 bzip2 解压流
	var bzerr c.Int
	bzfile := bzip2.BzReadOpen(&bzerr, inFile, 0, 0, nil, 0)
	if bzfile == nil || bzerr != bzip2.BZ_OK {
		return fmt.Errorf("BzReadOpen error, code=%d", bzerr)
	}

	buf := make([]byte, 4096)
	for {
		n := bzip2.BzRead(&bzerr, bzfile, unsafe.Pointer(&buf[0]), c.Int(len(buf)))
		if bzerr == bzip2.BZ_STREAM_END {
			// 表示解压结束
			if n > 0 {
				// 理论上不会有数据了，但如果有，就写一下
				c.Fwrite(unsafe.Pointer(&buf[0]), 1, uintptr(n), outFile)
			}
			break
		}

		if bzerr != bzip2.BZ_OK && bzerr != bzip2.BZ_STREAM_END {
			return fmt.Errorf("BzRead error, code=%d", bzerr)
		}
		// 写入解压后的数据
		if n > 0 {
			c.Fwrite(unsafe.Pointer(&buf[0]), 1, uintptr(n), outFile)
		} else {
			// 如果没有数据读出，就退出
			break
		}
	}

	bzip2.BzReadClose(&bzerr, bzfile)
	if bzerr != bzip2.BZ_OK {
		return fmt.Errorf("BzReadClose error, code=%d", bzerr)
	}

	return nil
}

func main() {
	err := DecompressFile("test.bz2", "ttt.test")
	if err != nil {
		panic(err)
	}

	b, err := os.ReadFile("ttt.test")
	if err != nil {
		panic(err)
	}

	log.Println(string(b))
}
