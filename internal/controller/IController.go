package IController

import "sync"

type IInsertController interface {
	StartListen(waitGroup *sync.WaitGroup)
}

func StartInsertListener(waitGroup *sync.WaitGroup, listener IInsertController){
	listener.StartListen(waitGroup)
}