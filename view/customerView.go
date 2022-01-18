package main

import (
	"customerMange/model"
	"customerMange/service"
	"fmt"
)

type customerView struct {
	//接收用户的输入
	key string
	//表示受否循环的显示主菜单
	loop bool
	//增加一个customerService字段
	customerservice *service.CustomerService
}

//显示所有客户的信息
func (cv *customerView) list() {
	//首先获取当前所有的客户信息（都在切片里）
	customers := cv.customerservice.List()
	fmt.Println("---------------客户列表---------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i :=0; i < len(customers); i++{
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("---------------客户列表完成------------------")
	fmt.Println()
}

//添加客户信息
func (cv *customerView) add(){
	fmt.Println("---------------添加客户---------------------")
	fmt.Println("输入姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("输入性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("输入年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("输入电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("输入邮箱：")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的Customer实例
	//id号没有让用户输入，id有系统分配
	custommer := model.NewCustomerNoid(name, gender, age, phone, email)
	if cv.customerservice.Add(custommer){
		fmt.Println("---------------添加完成---------------------")
		fmt.Println()
	}
}

//修改客户信息
func (cv *customerView) change(){
	fmt.Println("---------------添加客户---------------------")
	fmt.Println("输入要修改的客户编号：")
	id := -1
		fmt.Scanln(&id)
	fmt.Println("把姓名改成：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("把性别改成：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("把年龄改成：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("输入电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("输入邮箱：")
	email := ""
	fmt.Scanln(&email)

	req := cv.customerservice.Change(id, name, gender, age, phone, email)
	if req{
		fmt.Println("---------------修改完成---------------------")
		fmt.Println()
	}
}

//删除客户
func (cv *customerView) delete(){
	fmt.Println("---------------删除客户---------------------")
	fmt.Println("输入要删除的客户编号：")
	id := -1
	fmt.Scanln(&id)
	if id == -1{
		//如果return出现在普通函数中，表示跳出该函数，即不再执行return后面的代码，也可以理解成终止函数
		return
	}
	fmt.Println("确认删除请输入（Y/y）：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y"{
		//调用customerService的Delete方法
		if cv.customerservice.Delete(id){
			fmt.Println("---------------已删除---------------------")
			fmt.Println()
		} else {
			fmt.Println("---------------客户id不存在---------------------")
		}
	}
}

func (c *customerView) mainMenu()  {
	for  {
		fmt.Println("---------------客户信息管理软件---------------")
		fmt.Println("               1 添 加 客 户")
		fmt.Println("               2 修 改 客 户")
		fmt.Println("               3 删 除 客 户")
		fmt.Println("               4 客 户 列 表")
		fmt.Println("               5 退      出")
		fmt.Println("请选择（1-5）：")

		fmt.Scanln(&c.key)
		switch c.key{
		case "1":
			c.add()
		case "2":
			c.change()
		case "3":
			c.delete()
		case "4":
			c.list()
		case "5":
			c.loop = false
		default:
			fmt.Println("输入有误")
		}
		if !c.loop{
			break
		}
	}
	fmt.Println("已退出客户管理系统")
}

func main()  {
	//在主函数中创建一个customerView实例，并运行主菜单
	customerview := customerView{
		key:  "",
		loop: true,
	}
	//这里对customerView结构体的customerService字段的初始化
	customerview.customerservice = service.NewCustomerService()
	customerview.mainMenu()
}
