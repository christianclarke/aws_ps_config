package main

import (
 "aws_ps_config/lib"
 "github.com/sirupsen/logrus"
 "fmt"
)

func main() {
	
	x, err := getconfig.Config("/c3/services/texttospeech", "eu-west-1")

	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(x)
}
