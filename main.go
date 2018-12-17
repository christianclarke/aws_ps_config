package main

import (
	"aws_ps_config/lib"
	"flag"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

var awsParamStorePath string
var awsRegion string
var executable string

func init() {
	flag.StringVar(&awsParamStorePath, "awsParamStorePath", "", "the aws param store path to get")
	flag.StringVar(&awsRegion, "awsRegion", "", "the aws region param store")
	flag.StringVar(&executable, "executable", "", "The program to execute")
	flag.Parse()

	if awsParamStorePath == "" || awsRegion == "" || executable == "" {
		logrus.Error("The Aws Parameter Store Path, the Aws Region or the executable were not provided.")
		os.Exit(1)
	}
}

func main() {
	config, err := lib.GetAwsParamStoreData(awsParamStorePath, awsRegion)

	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	lib.SetEnvironmentVars(config)
	command := exec.Command(executable)
	err = command.Run()

	if (err != nil) {
		logrus.Fatal(err)
		os.Exit(1)
	}
}
