package service

import "customerMange/model"

type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片已经有多少个客户
	//该字段后面还可以作为新客户的id+1
	customerNum int
}

//返回*CustomerService
func NewCustomerService() *CustomerService  {
	//为了方便查看，我们先初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (CS *CustomerService) List() []model.Customer  {
	return CS.customers
}

//添加客户到customers切片
func (CS *CustomerService) Add(customer model.Customer) bool  {
	//确定一个生成id的规则，就是添加的顺序
	CS.customerNum++
	customer.Id = CS.customerNum
	CS.customers = append(CS.customers, customer)
	return true
}

//根据id查找客户在切片中对应的索引，如果无，返回-1
func (CS *CustomerService) FindByid(id int) int{
	index := -1
	for i, v := range CS.customers{
		if v.Id == id{
			index = i
			break
		}
	}
	return index
}

//删除客户
func (CS *CustomerService) Delete(id int) bool{
	index := CS.FindByid(id)
	if index == -1{
		return false
	}
	//从切片中删除一个元素，用append()追加一个切片，追加的切片参数后面要有三个点 ...  ，固定格式
	CS.customers = append(CS.customers[:index], CS.customers[index+1:]...)
	return true
}
