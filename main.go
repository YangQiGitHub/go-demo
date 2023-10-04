package main

import (
	fmtstd "github.com/YangQiGitHub/go-demo/fmt-std"
	// gormdemo "github.com/YangQiGitHub/go-demo/gorm-demo"
	logstd "github.com/YangQiGitHub/go-demo/log-std"
	osstd "github.com/YangQiGitHub/go-demo/os-std"
	// readyml "github.com/YangQiGitHub/go-demo/read-yml"
	timestd "github.com/YangQiGitHub/go-demo/time-std"
)

func main() {
  fmtstd.FmtDemoFun()
  osstd.OsDemoFun()
  osstd.FileDemo()
  timestd.OsTimeDemo()
  logstd.LogDemo()

  // readyml.ReadConf()
  // gormdemo.ConnectDatebase()

}


