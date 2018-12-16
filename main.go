package main

import (
 "aws_ps_config/lib"
 "github.com/sirupsen/logrus"
 "fmt"
)

func main() {
	var path = "path"
	var region = "eu-west-1"
	
	x, err := getconfig.Config(path, region)

	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(x)
}
