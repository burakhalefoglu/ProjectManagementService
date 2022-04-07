package concrete

import (
	"ProjectManagementService/internal/IoC"
	"ProjectManagementService/internal/model"
	"ProjectManagementService/internal/repository/abstract"
	JsonParser "ProjectManagementService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type customerProjectManager struct {
	Parser             *JsonParser.IJsonParser
	CustomerProjectDal *abstract.ICustomerProjectDal
}

func NewCustomerProjectManager() *customerProjectManager {
	return &customerProjectManager{Parser: &IoC.JsonParser,
		CustomerProjectDal: &IoC.CustomerProjectDal}
}

func (c *customerProjectManager) CreateOrUpdateCustomerProject(m *[]byte) (success bool, message string) {
	p := model.CustomerProject{}
	if err := (*c.Parser).DecodeJson(m, &p); err != nil {
		clogger.Error(&map[string]interface{}{
			"Json Parser Decode Err: ": err,
		})
		return false, err.Error()
	}

	var project, err = (*c.CustomerProjectDal).GetByPK(p.ProjectId, p.IndustryId, p.ProductId)
	if err != nil {
		clogger.Error(&map[string]interface{}{
			"CustomerProjectDal_GetByPK": err.Error(),
		})
		return false, err.Error()
	}
	defer clogger.Info(&map[string]interface{}{
		fmt.Sprintf("CustomerProject: %d", p.Id): "GetByPK",
	})

	if project != nil {
		if err := (*c.CustomerProjectDal).UpdateByPK(&p); err != nil {
			return false, err.Error()
		}
		return true, ""
	}
	if err := (*c.CustomerProjectDal).Add(&p); err != nil {
		return false, err.Error()
	}
	return true, ""
}
