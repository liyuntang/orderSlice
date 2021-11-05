package orderSlice

// import (
// 	"fmt"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

func GetSlice () []string {
	return []string{"aaa", "bbb", "ccc"}
	
}

// func CreateDb() {
// 	dsn := "mmmmmmm:XphJc@tcp(17.140.100.111:9000)/db_platform?charset=utf8mb4&parseTime=True&loc=Local"
// 	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("connect to mysql is bad, err is", err)
// 		return
// 	}
// 	fmt.Println("connect to mysql is ok")
// }
