package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/seggga/csvquery/config"
	"github.com/seggga/csvquery/rpn"
	"github.com/seggga/csvquery/token"
	"github.com/sirupsen/logrus"
)

var (
	gitCommit string                // print git commit version when program starts
	logFile   string = "access.log" // log-file to hold user's querys
	errFile   string = "error.log"  // error-log file, holds interrupt, timeout issues and invalid user's query
)

func main() {
	// load configuration from file
	conf, err := config.GetConfig("config/config.toml")
	if err != nil {
		panic(err)
	}

	// create info logger
	logInfo := logrus.New()
	fileInfo, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileInfo.Close()
	initLogInfo(logInfo, fileInfo, conf)

	// create error logger
	logErr := logrus.New()
	fileErr, err := os.OpenFile(errFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileErr.Close()
	initLogErr(logErr, fileErr, conf)

	// print binary path and commit version
	if printBinaryData() != nil {
		fmt.Println(err)
		return
	}

	// load users query
	fmt.Println("Please, enter the query : ")
	reader := bufio.NewReader(os.Stdin)
	query, err := reader.ReadString('\n')
	if err != nil {
		logErr.Errorf("There is an error entering data.\n%v\n", err)
		return
	}
	logInfo.Infof("user's query is: %s", query)
	//
	//
	//
	//
	//
	//
	//
	//
	//

	//query := `age > 40 AND (city_name == "Tokyo" OR new_issues <= 1000)`

	queryTokens := token.SplitQuery(query)
	queryTokens = rpn.ConvertToRPN(queryTokens)

	valuesMap := map[string]string{
		"age":        "30",
		"city_name":  "Moscow",
		"new_issues": "1000",
	}
	currentSlice := rpn.InsertValues(valuesMap, queryTokens)

	got, err := rpn.CalculateRPN(currentSlice)

	fmt.Println("got:", got, "error:", err)
}

// initLogInfo - initializes info logger
func initLogInfo(log *logrus.Logger, file *os.File, conf *config.ConfigType) {

	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(file)
	log.Info("logging started")
	log.Info("command-line parameters:")
	log.Infof("timeout: %s", conf.Timeout)
}

// initLogErr - initializes error logger
func initLogErr(log *logrus.Logger, file *os.File, conf *config.ConfigType) {
	log.SetLevel(logrus.ErrorLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(file)
}

// printBinaryData prints data about the binary
func printBinaryData() error {

	// print current directory path
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Printf("binary path: %s\n", path)

	// print commit
	fmt.Printf("commit version: %s\n", gitCommit)

	return nil
}
