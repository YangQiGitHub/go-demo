package osstd

import (
	"fmt"
	"io"
	"os"
)

func FileDemo() {
  // fileStatus()
  // readFile()
  // readFileAt()
  // readDir()
  // fileSeek()
  // fileWrite()
  fileWriteAt()
}

func fileStatus() {
  f, err1 := os.OpenFile("os-std/1.txt", os.O_RDWR | os.O_CREATE, 0775)
  if err1 != nil {
    fmt.Println(err1)
  }
  defer f.Close()

  fileInfo, err2 := f.Stat()
  if err2 != nil {
    panic(err2)
  }
  fmt.Printf("%v\n", fileInfo.Size())
}


func readFile() {
  f, err1 := os.OpenFile("os-std/1.txt", os.O_RDWR | os.O_CREATE, 0775)
  if err1 != nil {
    fmt.Println(err1)
  }
  defer f.Close()

  var body []byte
  for {
    buf := make([]byte, 4)
    n, err2 := f.Read(buf)
    if err2 == io.EOF {
      break
    }

    fmt.Printf("Read size: %d \n", n)
    body = append(body, buf[:n]...)
  }
  fmt.Printf("Read Content: %s \n", body)
}


func readFileAt() {
  f, err1 := os.OpenFile("os-std/1.txt", os.O_RDWR | os.O_CREATE, 0775)
  if err1 != nil {
    fmt.Println(err1)
  }
  defer f.Close()

  buf := make([]byte, 3)
  n, err2 := f.ReadAt(buf, 6)
  if err2 != nil {
    panic(err2)
  }
  fmt.Printf("read content: %s\n", buf[:n])
}

func readDir() {
  f, err := os.Open("os-std")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  dirs, err := f.ReadDir(-1)
  for _, v := range dirs {
    fmt.Println(v.Name(), ":", v.IsDir())
  }
}

func fileSeek() {
  f, err1 := os.OpenFile("os-std/1.txt", os.O_RDWR | os.O_CREATE, 0775)
  if err1 != nil {
    fmt.Println(err1)
  }
  defer f.Close()

  f.Seek(3, 0)
  buf := make([]byte, 4)
  n, _ := f.Read(buf)
  fmt.Printf("读取内容:%s\n", buf[:n])
}

func fileWrite() {
  f, err1 := os.OpenFile("os-std/1.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0755)
  if err1 != nil {
    panic(err1)
  }
  defer f.Close()

  f.Write([]byte("hello golang~\n"))
  f.WriteString("Hello Golang\n")
}

func fileWriteAt() {
  f, err1 := os.OpenFile("os-std/1.txt", os.O_RDWR | os.O_CREATE, 0755)
  if err1 != nil {
    panic(err1)
  }
  defer f.Close()
  _, err := f.WriteAt([]byte("insert content"), 5)
  fmt.Println(err)
}

