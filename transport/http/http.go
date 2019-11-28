package http

import (
	"log"
	"net/http"

	"github.com/phungvandat/onemilion/domain"
	"github.com/phungvandat/onemilion/service"
	"github.com/phungvandat/onemilion/util/constants"
	jq "github.com/phungvandat/onemilion/util/jobqueue"
)

// NewHTTPHandler func
func NewHTTPHandler(jobqueue jq.JobQueue, s service.Service) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", logRouteMiddleware(http.HandlerFunc(index)))
	mux.Handle("/test", logRouteMiddleware(webhook(jobqueue)))
	mux.HandleFunc("/other", other(s))
	return mux
}

func other(s service.Service) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != constants.GETMethod.String() {
			http.Error(res, "GET method only", http.StatusMethodNotAllowed)
			return
		}
		num := "other"
		input := domain.Payload{
			Num: num,
		}
		s.Test.Test(input)
		res.Write([]byte("XOng roi ne"))
	}
}

func webhook(jobQueue jq.JobQueue) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != constants.POSTMethod.String() &&
			req.Method != constants.GETMethod.String() {
			http.Error(res, "GET, POST method only", http.StatusMethodNotAllowed)
			return
		}
		if req.Method == constants.POSTMethod.String() {
			var num string = "0"
			nums, ok := req.URL.Query()["num"]
			if ok && len(nums[0]) >= 1 {
				num = nums[0]
			}
			job := jq.Job{
				Payload: domain.Payload{
					Num: num,
				},
			}
			jobQueue <- job
		}
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	if req.Method != constants.GETMethod.String() {
		http.Error(res, "GET method only", http.StatusMethodNotAllowed)
		return
	}
	res.Write([]byte("API"))
}

func logRouteMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("Method:%v Route:%v", req.Method, req.URL.String())
		next.ServeHTTP(res, req)
	})
}
