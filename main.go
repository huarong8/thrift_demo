package main

import (
	"flag"
)

func main(){
	var f string
	flag.StringVar(&f, "flag", "client", "client or server ?")
	flag.Parse()

	if f == "client"{
		Client()
	}else{
		Serve()
	}
}
