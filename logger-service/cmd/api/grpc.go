package main

import (
	"context"
	"fmt"
	"github.com/matizaj/go-app/log-service/data"
	log "github.com/matizaj/go-app/log-service/logs"
	"google.golang.org/grpc"
	"net"
)

type LogServer struct {
	log.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *log.LogRequest) (*log.LogResponse, error) {
	// get the inputa data: name, data
	input := req.GetLogEntry()

	// write the log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		fmt.Println(err)
		res := &log.LogResponse{Result: "failure"}
		return res, err
	}
	res := &log.LogResponse{Result: "logged!"}
	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		fmt.Println("Failure to start grpc ", err)
	}
	s := grpc.NewServer()
	log.RegisterLogServiceServer(s, &LogServer{Models: app.Models})
	fmt.Printf("grpc server started on port %s", gRpcPort)
	err = s.Serve(lis)
	if err != nil {
		fmt.Println("grpc server failed ", err)
	}
}
