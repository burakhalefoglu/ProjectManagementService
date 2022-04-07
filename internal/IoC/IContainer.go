package IoC

import (
	repository "ProjectManagementService/internal/repository/abstract"
	service "ProjectManagementService/internal/service/abstract"
	jsonParser "ProjectManagementService/pkg/jsonParser"
	"ProjectManagementService/pkg/kafka"
	cache "ProjectManagementService/pkg/redis"
)

type IContainer interface {
	Inject()
}

func InjectContainers(container IContainer) {
	container.Inject()
}

var RedisCache cache.ICache
var Kafka kafka.IKafka
var JsonParser jsonParser.IJsonParser

var CustomerProjectService service.ICustomerProjectService
var CustomerProjectDal repository.ICustomerProjectDal
