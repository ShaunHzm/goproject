package test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var customers []map[string]interface{}

var customerid int

// 查找id
func findById(id int) int {
	index := -1
	for i := 0; i < len(customers); i++ {
		if customers[i]["cid"] == id {
			index = i
		}
	}
	return index
}

// 用户是否返回上一层
func isBack() bool {
	fmt.Println("请问是否返回上一层【Y/N】：")
	var backChoice string
	fmt.Scan(&backChoice)
	if strings.ToUpper(backChoice) == "Y" {
		return true
	} else {
		return false
	}
}

func inputInfo() (string, string, int8, string) {
	var name string
	fmt.Println("请输入客户姓名：")
	fmt.Scan(&name)

	var gender string
	fmt.Println("请输入性别：")
	fmt.Scan(&gender)

	var age int8
	fmt.Println("请输入客户年龄：")
	fmt.Scan(&age)

	var email string
	fmt.Println("请输入客户email：")
	fmt.Scan(&email)

	return name, gender, age, email
}

func addCustomer() {
	for {
		fmt.Printf("\033[1;35;40m%s\033[0m\n", "------------添加客户开始------------")
		name, gender, age, email := inputInfo()
		customerid++
		newCustomer := map[string]interface{}{
			"cid":    customerid,
			"name":   name,
			"gender": gender,
			"age":    age,
			"email":  email,
		}

		customers = append(customers, newCustomer)
		fmt.Printf("\033[1;35;40m%s\033[0m\n", "------------添加客户完成------------")

		b := isBack()
		if b {
			break
		}
	}
}

func listCustomer() {
	for {
		fmt.Printf("\033[1;32;40m%s\033[0m\n", "------------客户列表开始------------")
		for _, customer := range customers {
			fmt.Printf("编号：%-8d 姓名：%-8s 性别：%-8s 年龄：%-8d 邮箱：%-8s\n",
				customer["cid"], customer["name"], customer["gender"], customer["age"], customer["email"])
		}
		fmt.Printf("\033[1;32;40m%s\033[0m\n", "------------客户列表结束------------")
		b := isBack()
		if b {
			break
		}
	}
}

func updateCustomer() {
	for {
		fmt.Printf("\033[1;36;40m%s\033[0m\n", "------------客户修改开始------------")
		var updateCid int
		fmt.Printf("请输入更新客户编号")
		fmt.Scan(&updateCid)
		updateIndex := findById(updateCid)
		if updateIndex == -1 {
			fmt.Println("更新失败，该客户id不存在")
			continue
		}
		fmt.Println("请输入修改的客户信息")
		name, gender, age, email := inputInfo()

		customers[updateIndex]["name"] = name
		customers[updateIndex]["gender"] = gender
		customers[updateIndex]["age"] = age
		customers[updateIndex]["email"] = email
		fmt.Printf("\033[1;36;40m%s\033[0m\n", "------------客户修改完成------------")
		b := isBack()
		if b {
			break
		}
	}

}

func deleteCustomer() {
	fmt.Printf("\033[1;31;40m%s\033[0m\n", "------------删除客户开始------------")
	var deleteCid int
	fmt.Print("请输入删除客户ID:")
	fmt.Scan(&deleteCid)

	deleteIndex := findById(deleteCid)
	if deleteIndex == -1 {
		fmt.Println("删除失败：删除的客户ID不存在")
		return
	}

	customers = append(customers[:deleteIndex], customers[deleteIndex+1:]...)

	fmt.Printf("\033[1;31;40m%s\033[0m\n", "------------删除客户结束------------")
}

func MangerTest() {
	for {
		fmt.Printf("\033[1;33;40m%s\033[0m\n", `
----------------客户信息管理系统--------------
	1、添加客户
	2、查看客户
	3、更新客户
	4、删除客户
	5、退出
-------------------------------------------
`)
		var choice int
		fmt.Printf("\033[1;38;40m%s\033[0m", "请输入选择【1-5】:")
		stdin := bufio.NewReader(os.Stdin)
		fmt.Fscan(stdin, &choice)

		switch choice {
		case 1:
			addCustomer()
		case 2:
			listCustomer()
		case 3:
			updateCustomer()
		case 4:
			deleteCustomer()
		default:
			os.Exit(0)
		}
	}
}
