// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	// "fmt"
	"demo/demo_proto/biz/dal"
	"demo/demo_proto/biz/dal/mysql"
	"demo/demo_proto/biz/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	// CURD

	// Create
	mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "demo_passwd"})

	// // Update
	// mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("password", "22222222")

	// // Read
	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&row)

	// fmt.Printf("row: %+v\n", row)

	// // Delete
	// mysql.DB.Where("email = ?", "demo@example.com").Delete(&model.User{})

	// // 强制删除，区别于上面软删
	// mysql.DB.Unscoped().Where("email = ?", "demo@example.com").Delete(&model.User{})
}
