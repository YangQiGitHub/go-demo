package osstd

import (
	"fmt"
	"os"
)

func OsDemoFun() {
  // osCreate()
  // osMkdir()
  // osMkdirAll()
  // osRemove()
  // osRemoveAll()
  // osGetwd()
  // osChdir()
  // osTempDir()
  // osRename()
  // osChmod()
}

// Create File
func osCreate() {
  f, err := os.Create("os-std/test.txt")
  if err != nil {
    panic(err)
  }
  defer f.Close()
  fmt.Printf("f: %v\n", f)
}


// Make Dir
func osMkdir() {
  err := os.Mkdir("os-std/testdir", os.ModePerm)
  if err != nil {
    panic(err)
  }
}

// 创建多级目录
func osMkdirAll() {
  err := os.MkdirAll("os-std/a/b/c", os.ModePerm)
  if err != nil {
    panic(err)
  }
}

// Remove removes the named file or (empty) directory.
func osRemove() {
  err := os.Remove("os-std/testdir")
  if err != nil {
    panic(err)
  }
}

// 强制删除目录以及目录中的文件
func osRemoveAll() {
  err := os.RemoveAll("os-std/a")
  if err != nil {
    panic(err)
  }
}

// 获取工作目录
func osGetwd() {
  dir, err := os.Getwd()
  if err != nil {
    fmt.Printf("err: %v\n", err)
  } else {
    fmt.Printf("%v\n", dir)
  }
}

// 改变工作目录
func osChdir() {
  err := os.Chdir("a/b")
  if err != nil {
    fmt.Printf("err: %v\n", err)
  } else {
    fmt.Println(os.Getwd())
  }
}

// 获取存放临时文件的临时目录
func osTempDir() {
  s := os.TempDir()
  fmt.Println(s)
}

// 重命名文件
func osRename() {
  // err1 := os.Rename("os-std/test.txt", "os-std/1.txt")
  // fmt.Printf("%v\n", err1)

  err2 := os.Rename("os-std/a", "os-std/b")
  fmt.Printf("%v\n", err2)
}

// 修改文件权限
func osChmod() {
  err := os.Chmod("os-std/1.txt", 0644)
  if err != nil {
    fmt.Println(err)
  }
}


