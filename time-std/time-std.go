package timestd

import (
	"fmt"
	"time"
)

func OsTimeDemo() {
  // timeNow()
  // timeUnix()
  // parseTime()
  // timeFormat()
  // timeDuration()
  // timeEqual()
  // timeTicker()
}

func timeNow()  {
  now := time.Now()
  fmt.Printf("%T\n %#[1]v\n %+[1]v\n", now)
  fmt.Println("-------------------------")
  fmt.Println("year:", now.Year())
  fmt.Println("month:", now.Month())
  fmt.Println("day:", now.Day())
  fmt.Println("hour:", now.Hour())
  fmt.Println("minute:", now.Minute())
  fmt.Println("second:", now.Second())
}

func timeUnix() {
  now := time.Now()
  unixSec := now.Unix()
  fmt.Printf("unix sec values: %d\n", unixSec)
  unixMilli := now.UnixMilli()
  fmt.Printf("unix milli values: %d\n", unixMilli)

  nowTime := time.UnixMilli(unixMilli)
  fmt.Printf("current time: %v", nowTime)
}

func parseTime() {
  ti, err := time.Parse("2006/01/02 15:04:05", "2023/08/01 13:00:00")
  if err != nil {
    panic(err)
  }
  fmt.Printf("%v\n", ti)

  loc, err := time.LoadLocation("Asia/Shanghai")
  if err != nil {
    panic(err)
  }
  locTime, err := time.ParseInLocation("2006/01/02 15:04:05", "2023/08/01 13:00:00", loc)
  if err != nil {
    panic(err)
  }
  fmt.Printf("time loc: %v\n", locTime)
}

func timeFormat() {
  now := time.Now()
  timeStr := now.Format("2006-01-02 15:04:05")
  fmt.Printf("Format time str: %v\n", timeStr)

  time12Str := now.Format("2006-01-02 03:04:05")
  fmt.Printf("Format time 12 str: %v\n", time12Str)
}

func timeDuration() {
  fiveMinutes := 5 * time.Second

  now := time.Now()
  fmt.Printf("time now:\n %v\n", now)

  nextTime := now.Add(fiveMinutes)
  fmt.Printf("time add:\n %v\n", nextTime)

  sub := nextTime.Sub(now)
  fmt.Printf("time sub: %v\n", sub)
}

func timeEqual() {
  now := time.Now()
  now1 := time.Now()
  fmt.Println(now.Equal(now1))
}

func timeTicker() {
  // tick := time.Tick(5 * time.Second)
  // fmt.Println(time.Now())
  // for t2 := range tick {
  //   fmt.Println(t2)
  // }

  // time.AfterFunc(time.Second * 5, func() {
  //   fmt.Println("5 second run")
  // })

  fmt.Println(time.Now())
  time.Sleep(5 * time.Second)
  fmt.Println(time.Now())
}


