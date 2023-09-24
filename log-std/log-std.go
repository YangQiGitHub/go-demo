package logstd

import (
	"fmt"
	"log"
	"os"
)

func LogDemo() {
  // logPrint()
  // logPanic()
  // logFatal()
  // logConf()
  // logPrefix()
  log2File()
}

func logPrint() {
  log.Print("This is a log")
  log.Printf("This is a log: %d", 100)
  name := "zhangsan"
  age := 12
  log.Println(name, "is", age, "years old.")
}

func logPanic() {
  defer fmt.Println("发生 Panic 错误")
  log.Print("This is a log")
  log.Panic("This is a panic log.")
  log.Panicln("运行结束")
}

func logFatal() {
  defer fmt.Println("defer fatal...")
  log.Print("This is a log")
  log.Fatal("This is a fatal log.")
  log.Panicln("运行结束")
}

func logConf() {
  i := log.Flags()
  fmt.Println(i)
  log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
  log.Print("This is a new config")
  fmt.Println(log.Flags())
}

func logPrefix() {
  prefix := log.Prefix()
  fmt.Printf("日志的前缀是:%v\n", prefix)
  log.SetPrefix("[log-prefix] ")

  log.Print("set log prefix")
}

func log2File() {
  f, err := os.OpenFile("log-std/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
  if err != nil {
    log.Panic("打开日志文件异常")
  }
  log.SetPrefix("[abc prefix] ")
  log.SetOutput(f)
  log.Print("this is a file log...")
}
