/*
Copyright 2020 The Tekton Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	// "context"
	// "fmt"
	// "net"
	"net/http"
	"path"
	// "time"

	// "k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/rest"

	"go.uber.org/zap"
	// "go.uber.org/zap/zapcore"

	// "github.com/golang-jwt/jwt/v4"
	// grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	// grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	// grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	// "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	// grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	// prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tektoncd/results/pkg/api/server/config"
	"github.com/tektoncd/results/pkg/api/server/logger"
	// v1alpha2 "github.com/tektoncd/results/pkg/api/server/v1alpha2"
	// "github.com/tektoncd/results/pkg/api/server/v1alpha2/auth"
	// v1alpha2pb "github.com/tektoncd/results/proto/v1alpha2/results_go_proto"
	_ "go.uber.org/automaxprocs"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/credentials/insecure"
	// "google.golang.org/grpc/health"
	// healthpb "google.golang.org/grpc/health/grpc_health_v1"
	// "google.golang.org/grpc/reflection"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"log"
)

func main() {
	log.Println("Test6")
	serverConfig := config.Get()

	log := logger.Get(serverConfig.LOG_LEVEL)
	defer log.Sync()

	if serverConfig.DB_USER == "" || serverConfig.DB_PASSWORD == "" {
		log.Fatal("Must provide both DB_USER and DB_PASSWORD")
	}
	// Connect to the database.
	// DSN derived from https://pkg.go.dev/gorm.io/driver/postgres

	// dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", serverConfig.DB_HOST, serverConfig.DB_USER, serverConfig.DB_PASSWORD, serverConfig.DB_NAME, serverConfig.DB_PORT, serverConfig.DB_SSLMODE)

	gormConfig := &gorm.Config{}
	if log.Level() != zap.DebugLevel {
		gormConfig.Logger = gormlogger.Default.LogMode(gormlogger.Silent)
	}
	// db, err := gorm.Open(postgres.Open(dbURI), gormConfig)
	// if err != nil {
	// 	log.Fatalf("Failed to open the results.db: %v", err)
	// }

	// Load TLS cert for server
	// creds, tlsError := credentials.NewServerTLSFromFile(path.Join(serverConfig.TLS_PATH, "tls.crt"), path.Join(serverConfig.TLS_PATH, "tls.key"))
	// if tlsError != nil {
	// 	log.Infof("Error loading TLS key pair for server: %v", tlsError)
	// 	log.Warn("Creating server without TLS")
	// 	creds = insecure.NewCredentials()
	// }

	// Register API server(s)
	// v1a2, err := v1alpha2.New(serverConfig, log, db)
	// if err != nil {
	// 	log.Fatalf("Failed to create server: %v", err)
	// }

	// s := grpc.NewServer(
	// 	grpc.Creds(creds),
	// )
	// v1alpha2pb.RegisterResultsServer(s, v1a2)

	// Load REST client TLS cert to connect to the gRPC server
	// if tlsError == nil {
	// 	creds, err = credentials.NewClientTLSFromFile(path.Join(serverConfig.TLS_PATH, "tls.crt"), serverConfig.TLS_HOSTNAME_OVERRIDE)
	// 	if err != nil {
	// 		log.Fatalf("Error loading TLS certificate for REST: %v", err)
	// 	}
	// }

	// Register gRPC server endpoint for gRPC gateway
	// ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// defer cancel()
	mux := runtime.NewServeMux()
	mux.Handle("/", runtime.Pattern{}, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		log.Info("request....")
		w.Write([]byte("Hello"))
	})

	// Start REST proxy server
	log.Infof("REST server Listening on: %s", serverConfig.REST_PORT)
	// if tlsError != nil {
	// 	log.Fatal(http.ListenAndServe(":"+serverConfig.REST_PORT, mux))
	// } else {
		log.Fatal(http.ListenAndServeTLS(":"+serverConfig.REST_PORT, path.Join(serverConfig.TLS_PATH, "tls.crt"), path.Join(serverConfig.TLS_PATH, "tls.key"), mux))
	// }

}
