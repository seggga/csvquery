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

	// create a listener for SIGINT
	intChan := make(chan os.Signal, 1)
	signal.Notify(intChan, syscall.SIGINT)
	errorChan := make(chan error)
	finishChan := make(chan struct{})

	// context to set timeout
	ctx := context.Context(context.Background())
	ctx, cancelFunc := context.WithTimeout(ctx, time.Duration(conf.Timeout))

	// run scanner for csv-files
	go scanCSV(lexMachine, errorChan, finishChan, ctx)

	// watch for the interrupt signals and ctx closing because of timeout
	select {
	case err := <-errorChan:
		logErr.Errorln(err)
		fmt.Println("there is an error while reading csv-files")
		return
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
}
