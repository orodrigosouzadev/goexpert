package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/orodrigosouzadev/goexpert/14-grpc/internal/database"
	"github.com/orodrigosouzadev/goexpert/14-grpc/internal/pb"
	"github.com/orodrigosouzadev/goexpert/14-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
