package abstract

type ICustomerProjectService interface {
	CreateOrUpdateCustomerProject(model *[]byte) (success bool, message string)
}
