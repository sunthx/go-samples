package main

import (
	"sync"
	"time"
)

//条件变量
//sync.Cond
//1. wait 		等待
//2. signal 	单发通知
//3. broadcase 	广播
func main(){
	var mailbox uint8
	var lock sync.RWMutex
	var wait sync.WaitGroup

	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(lock.RLocker())

	wait.Add(10)

	// send mail
	go func(){
		for i := 0 ; i < 10 ; i++{
			// sleep for a while
			time.Sleep(time.Microsecond * 500)
			lock.Lock()

			for mailbox == 1 {
				print("mailbox == 1 , waiting ...\n")
				sendCond.Wait()
			}

			mailbox = 1
			lock.Unlock()
			recvCond.Signal()
		}
	}()

	// receive mail
	go func(){
		for  i := 0 ; i < 10 ; i++{
			// sleep for a while
			time.Sleep(time.Microsecond * 500)
			lock.RLock()

			for mailbox == 0 {
				print("mailbox == 0 , waiting ...\n")
				recvCond.Wait()
			}

			wait.Done()
			mailbox = 0
			lock.RUnlock()
			sendCond.Signal()
		}
	}()

	wait.Wait()

	print("DOWN")
}