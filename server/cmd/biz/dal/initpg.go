package dal

import (
	"fmt"
	"github.com/renxingdawang/SDataHK/biz/dal/db"
)

func Init() {
	db.Init()
	fmt.Println("success init db")
}
