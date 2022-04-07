package main

import (
	"ProjectManagementService/internal/IoC"
	"ProjectManagementService/internal/IoC/golobby"
	IController "ProjectManagementService/internal/controller"
	contorller "ProjectManagementService/internal/controller/kafka"
	"ProjectManagementService/pkg/helper"
	"log"
	"runtime"
	"sync"
	"time"

	logger "github.com/appneuroncompany/light-logger"

	"github.com/joho/godotenv"
)

func main() {
	defer helper.DeleteHealthFile()
	logger.Log.App = "ProjectManagementServiceWorker"
	runtime.MemProfileRate = 0
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	for {
		startConsumer()
		time.Sleep(time.Second * 5)
	}
}

func startConsumer() {
	wg := sync.WaitGroup{}
	IoC.InjectContainers(golobby.InjectionConstructor())
	IController.StartInsertListener(&wg, contorller.InsertControllerConstructor())
	wg.Wait()
}
