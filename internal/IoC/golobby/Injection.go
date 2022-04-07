package golobby

import (
	"ProjectManagementService/internal/IoC"
	repository "ProjectManagementService/internal/repository/abstract"
	"ProjectManagementService/internal/repository/concrete/Cassandra"
	service "ProjectManagementService/internal/service/abstract"
	manager "ProjectManagementService/internal/service/concrete"
	cassandra "ProjectManagementService/pkg/database/Cassandra"
	jsonParser "ProjectManagementService/pkg/jsonParser"
	"ProjectManagementService/pkg/jsonParser/gojson"
	"ProjectManagementService/pkg/kafka"
	"ProjectManagementService/pkg/kafka/kafkago"
	cache "ProjectManagementService/pkg/redis"
	rediscachev8 "ProjectManagementService/pkg/redis/redisv8"

	"github.com/golobby/container/v3"
)

type golobbyInjection struct{}

func InjectionConstructor() *golobbyInjection {
	return &golobbyInjection{}
}

func (i *golobbyInjection) Inject() {
	injectKafka()
	injectJsonParser()
	injectCache()

	injectCustomerProject()
}

func injectCustomerProject() {
	if err := container.Singleton(func() service.ICustomerProjectService {
		return manager.NewCustomerProjectManager()
	}); err != nil {
		panic(err)
	}
	if err := container.Singleton(func() repository.ICustomerProjectDal {
		return Cassandra.NewCustomerProjectDal(cassandra.CustomerProject)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.CustomerProjectDal); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.CustomerProjectService); err != nil {
		panic(err)
	}
}

func injectJsonParser() {
	if err := container.Singleton(func() jsonParser.IJsonParser {
		return gojson.GoJsonConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.JsonParser); err != nil {
		panic(err)
	}
}

func injectKafka() {
	if err := container.Singleton(func() kafka.IKafka {
		return kafkago.KafkaGoConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.Kafka); err != nil {
		panic(err)
	}
}

func injectCache() {
	if err := container.Singleton(func() cache.ICache {
		return rediscachev8.RedisCacheConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.RedisCache); err != nil {
		panic(err)
	}
}
