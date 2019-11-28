package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/joho/godotenv"
	svc "github.com/phungvandat/onemilion/service"
	testSvc "github.com/phungvandat/onemilion/service/test"
	httpTransport "github.com/phungvandat/onemilion/transport/http"
	mongoDB "github.com/phungvandat/onemilion/util/config/db/mongo"
	envConfig "github.com/phungvandat/onemilion/util/config/env"
	jq "github.com/phungvandat/onemilion/util/jobqueue"
)

func main() {
	var isProduction = envConfig.GetENV() == "production"
	if !isProduction {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("failed to load .env by error: %v", err))
		}
	}

	mongoDB, closeMongoDB := mongoDB.NewDB(envConfig.GetMogoDBName(), envConfig.GetMongoURI())
	defer closeMongoDB()

	testSvc := testSvc.NewTestService(mongoDB)
	svc := svc.Service{
		Test: testSvc,
	}

	var maxQueue int = 900
	if val, err := strconv.Atoi(envConfig.GetMaxQueue()); err == nil {
		maxQueue = val
	}

	jobQueue := jq.NewJobQueue(maxQueue)

	var maxWorker int = 100
	if val, err := strconv.Atoi(envConfig.GetMaxWorker()); err == nil {
		maxWorker = val
	}

	dispatcher := jq.NewDispatcher(maxWorker)

	dispatcher.Run(svc, jobQueue)

	errs := make(chan error)

	httpPort := "3003"
	if envConfig.GetBackendPort() != "" {
		httpPort = envConfig.GetBackendPort()
	}
	httpAddr := fmt.Sprintf(":%v", httpPort)

	var httpHandler http.Handler
	{
		httpHandler = httpTransport.NewHTTPHandler(jobQueue, svc)
	}

	go func() {
		log.Printf("transport:HTTP addr:%v", httpAddr)
		errs <- http.ListenAndServe(httpAddr, httpHandler)
	}()

	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-ch)
	}()

	log.Println("exit", <-errs)
}
