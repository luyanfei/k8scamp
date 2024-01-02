package main

import (
	"context"
	"httpserver/metrics"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
}

func main() {
	log.Println("Starting http server...")
	metrics.Register()

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthHandler)
	mux.Handle("/metrics", promhttp.Handler())
	serv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	go func() {
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown server failed: %s", err)
	}
	log.Println("Shutdown successfully.")
}

//当访问 localhost/healthz 时，应返回 200
func healthHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	delay := randInt(10, 2000)
	time.Sleep(time.Millisecond * time.Duration(delay))

	resHeader := w.Header()
	//接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		resHeader["x-request-"+k] = v
	}

	//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	version, ok := os.LookupEnv("VERSION")
	if ok {
		resHeader["x-VERSION"] = []string{version}
	} else {
		resHeader["x-VERSION"] = []string{"1.0"}
	}

	//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	log.Printf("Client address: %s, response status code: %d.\n", r.RemoteAddr, http.StatusOK)
}
