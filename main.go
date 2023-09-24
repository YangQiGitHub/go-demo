package main

import (
	fmtstd "github.com/YangQiGitHub/go-demo/fmt-std"
	logstd "github.com/YangQiGitHub/go-demo/log-std"
	osstd "github.com/YangQiGitHub/go-demo/os-std"
	timestd "github.com/YangQiGitHub/go-demo/time-std"
)

func main() {
  fmtstd.FmtDemoFun()
  osstd.OsDemoFun()
  osstd.FileDemo()
  timestd.OsTimeDemo()
  logstd.LogDemo()
}

