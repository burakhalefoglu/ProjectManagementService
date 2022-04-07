package KafkaController

import (
	"ProjectManagementService/internal/IoC"
	"ProjectManagementService/internal/service/abstract"
	"ProjectManagementService/pkg/helper"
	"ProjectManagementService/pkg/kafka"
	"sync"
)

type insertController struct {
	Kafka                  *kafka.IKafka
	CustomerProjectService *abstract.ICustomerProjectService
}

func InsertControllerConstructor() *insertController {
	return &insertController{Kafka: &IoC.Kafka,
		CustomerProjectService: &IoC.CustomerProjectService,
	}
}

func (controller *insertController) StartListen(waitGroup *sync.WaitGroup) {
	waitGroup.Add(1)
	helper.CreateHealthFile()
	go (*controller.Kafka).Consume("CustomerProjectMetadata",
		"CustomerProjectMetadata_ConsumerGroup",
		waitGroup,
		(*controller.CustomerProjectService).CreateOrUpdateCustomerProject)
}
