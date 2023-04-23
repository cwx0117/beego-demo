package test

import (
	"demo/utils/fileUtils"
	"fmt"
	"testing"
)

func TestToZip(t *testing.T) {
	flag, err := fileUtils.CompressDirToZip("D:\\uoload\\个人信息", "C:\\Users\\34528\\Desktop\\zip\\archive.zip")
	if err != nil {
		return
	}
	fmt.Println(flag)
}
