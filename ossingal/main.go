package main

import (
	"fmt"
	"os"
	"os/signal"
)

func singal1(){
	ch:=make(chan os.Signal,1)

	//第一个参数是监听的信号
	signal.Notify(ch,os.Kill,os.Interrupt)
	fmt.Print("ch wait beigin /r/n")
	s:=<-ch
	fmt.Printf("kill %v",s)
}

func singal2(){
	ch:=make(chan os.Signal,1)
	signal.Notify(ch)
	signal.Stop(ch)
	<-ch
	fmt.Print("xxxxx")

}

func main(){
	singal2()
}



