package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/seggga/csvquery/config"
	"github.com/seggga/csvquery/parse"
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

	// print binary's path and commit version
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
		fmt.Println("cannot read data ")
		return
	}
	logInfo.Infof("user's query is: %s", query)

	// CheckQuery checks the user's query for matching the pattern "SELECT-FROM-WHERE"
	err = parse.CheckQuery(query)
	if err != nil {
		logErr.Errorf("query is incorrect.\n%v\n", err)
		fmt.Println("query is incorrect.")
		return
	}

	// split query into lexemas
	queryLex := token.SplitQuery(query)
	// split lexemas into sections: SELECT, FROM and WHERE
	lexMachine, err := parse.NewLexMachine(queryLex)
	if err != nil {
		logErr.Errorln(err)
		fmt.Println(err)
		return
	}
	lexMachine.Where = rpn.ConvertToRPN(lexMachine.Where)

	// create a listener for SIGINT
	intChan := make(chan os.Signal, 1)
	signal.Notify(intChan, syscall.SIGINT)
	finishChan := make(chan struct{})

	// context to set timeout
	ctx := context.Context(context.Background())
	ctx, cancelFunc := context.WithTimeout(ctx, time.Duration(conf.Timeout))

	//

	// run scanner for csv-files
	//	go func(finishChan, ctx) {}()

	// watch for the interrupt signals and ctx closing because of timeout
	select {
	case <-intChan:
		logErr.Errorln("Program has been interrupted by user")
		fmt.Println("Program has been interrupted by user")
		cancelFunc()
	case <-ctx.Done():
		logErr.Errorln("there is no time left")
		fmt.Println("there is no time left")
	}

	// graceful shutdown to close opened csv-files
	timeOuter := time.NewTimer(time.Duration(conf.Graceful))
	select {
	case <-finishChan:
		logInfo.Println("all csv-files has been closed successfully")
	case <-timeOuter.C:
		logErr.Errorln("some csv-files has not been closed")
	}

	fmt.Println("Program exit")
	// currentSlice := rpn.InsertValues(valuesMap, queryTokens)

	// got, err := rpn.CalculateRPN(currentSlice)

	// fmt.Println("got:", got, "error:", err)
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
