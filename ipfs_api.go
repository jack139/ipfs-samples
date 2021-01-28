package main

import (
	"fmt"
	"strings"
	"os"
	"io"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	// Where your local node is running on localhost:5001
	sh := shell.NewShell("localhost:5001")

	// 添加内容
	cid, err := sh.Add(strings.NewReader("hello world!"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("added: %s\n", cid)

	// 直接读出成文件
	if err := sh.Get(cid, "temp.txt"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}


	// 获取文件内容
	data, err := sh.Cat(cid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer data.Close()

	// 使用缓存读出文件
	var dataBuf []byte
	longBuf := make([]byte, 5)

	for {
		sz, err := data.Read(longBuf)
		if err != nil {
			if err == io.EOF {
				if sz>0 { // EOF 此时有可能还读出了数据
					//fmt.Printf("EOF: %d\n", sz)
					dataBuf = append(dataBuf, longBuf[:sz]...)
				}
				break
			}
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		//fmt.Printf("%d %s\n", sz, longBuf)
		dataBuf = append(dataBuf, longBuf[:sz]...)
	}

	fmt.Printf("get: %d %s\n", len(dataBuf), dataBuf)

}