package abstract

import (
	"ProjectManagementService/internal/model"
)

type ICustomerProjectDal interface {
	GetByPK(projectId int64, industryId int8, productId int8) (*model.CustomerProject, error)

	UpdateByPK(data *model.CustomerProject) error

	Add(model *model.CustomerProject) error
}
