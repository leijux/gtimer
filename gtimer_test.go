package gtimer

import (
	"log"
	"testing"
	"time"
)

func Test_Task(t *testing.T) {
	timer:=New()
	timer.Task("leiju",time.Now().Add(2*time.Second),func(){
		log.Println("test")
	})
	timer.Task("sss",time.Now().Add(3*time.Second) , func(){
		log.Println("sssss")
	})
	time.Sleep(5 *time.Second)
}
