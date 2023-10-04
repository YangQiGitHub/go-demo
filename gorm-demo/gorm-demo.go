package gormdemo

import (
	"fmt"
	"gorm.io/gorm"
  "gorm.io/driver/mysql"
)

type UserInfo struct {
  ID uint
  Name string
  Gender string
  Hobby string
}

func ConnectDatebase() {
  dsn := "root:password@tcp(127.0.0.1:3306)/gorm_db1?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic(err)
  }

  fmt.Printf("%v\n", db)
  db.AutoMigrate(&UserInfo{})

  // Create
  u1 := &UserInfo{ID: 1, Name:"aaa", Gender: "male", Hobby: "Sing"}
  u2 := &UserInfo{ID: 2, Name:"bbb", Gender: "male", Hobby: "Movie"}
  db.Create(u1)
  db.Create(u2)
  fmt.Println("已经写入")

  var user1 UserInfo
  db.First(&user1, 1)
  var user2 UserInfo
  db.First(&user2, 2)
  fmt.Printf("用户信息1: %+v\n", user1)
  fmt.Printf("用户信息2: %+v\n", user2)

  db.First(&user1, "name=?", "bbb")
  fmt.Printf("用户信息1---: %+v\n", user1)

  db.Model(&user1).Update("Hobby", "Run")
  db.Model(&user1).Updates(UserInfo{Name: "ccc", Gender: "fff"})

  db.Delete(&user1, 1)
}
