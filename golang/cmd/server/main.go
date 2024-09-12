package main

import (
	// "context"

	"log"
	// "net/http"
	// "os"
	// "os/signal"

	// "strconv"
	// "syscall"
	// "time"

	// pb "english-ai-full/ecomm-grpc/proto"
"english-ai-full/ecomm-api/websocket/websocket_repository"
"english-ai-full/ecomm-api/websocket/websocket_service"
	comment_api "english-ai-full/ecomm-api/comment-api"
	reading_api "english-ai-full/ecomm-api/reading-api"
	user_api "english-ai-full/ecomm-api/user-api"
	websocket_handler "english-ai-full/ecomm-api/websocket/websocket_handler"
	"english-ai-full/ecomm-grpc/config"
	pb "english-ai-full/ecomm-grpc/proto"
	pb_comment "english-ai-full/ecomm-grpc/proto/comment"
	pb_reading "english-ai-full/ecomm-grpc/proto/reading"

	// "github.com/go-chi/chi"

	"github.com/go-chi/chi"
	"github.com/ianschenck/envflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"english-ai-full/ecomm-api/handler"
	"english-ai-full/ecomm-api/route"
)

const minSecretKeySize = 32

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	var (
		secretKey = envflag.String("SECRET_KEY", "01234567890123456789012345678901", "secret key for JWT signing")
		svcAddr   = envflag.String("GRPC_SVC_ADDR", cfg.GRPCAddress, "address where the ecomm-grpc service is listening on")
	)
	envflag.Parse()

	if len(*secretKey) < minSecretKeySize {
		log.Fatalf("SECRET_KEY must be at least %d characters", minSecretKeySize)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(*svcAddr, opts...)
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()



	r := chi.NewRouter()


// new reading handler
// conn_reading, err := grpc.NewClient(*svcAddr, opts...)
// if err != nil {
// 	log.Fatalf("failed to connect to server: %v", err)
// }
// defer conn.Close()

// reading 
client_reading := pb_reading.NewEcommReadingClient(conn)

hdl_reading := reading_api.NewReadingHandler(client_reading, *secretKey)
reading_api.RegisterReadingRoutes(r, hdl_reading)
//  user handler
client := pb.NewEcommUserClient(conn)
	hdl := handler.NewHandler(client, *secretKey)
	
	route.RegisterRoutes(r, hdl)


	// comment 
	client_comment := pb_comment.NewCommentServiceClient(conn)

	hdl_comment := comment_api.NewCommentHandler(client_comment, *secretKey)
	comment_api.RegisterCommentRoutes(r, hdl_comment)


// new user  handler
hdl_NewUser := user_api.NewHandlerUser(client, *secretKey)

user_api.RegisterRoutesUser(r, hdl_NewUser)
// start user
websockrepo := websocket_repository.NewInMemoryMessageRepository()
websocketService := websocket_service.NewWebSocketService(websockrepo)
go websocketService.Run()

websocketHandler := websocket_handler.NewWebSocketHandler(websocketService)

// http.HandleFunc("/ws", websocketHandler.HandleWebSocket)

// wsHandler := websocket.NewHandler()
r.Get("/ws", websocketHandler.HandleWebSocket)





route.Start(":8888", r)


}


	// go func() {
	// 	log.Printf("Starting HTTP server on %s", cfg.HTTPAddress)
	// 	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("Failed to start Chi server: %v", err)
	// 	}
	// }()

	// //  Graceful shutdown
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	// <-quit

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// if err := server.Shutdown(ctx); err != nil {
	// 	log.Fatalf("Server forced to shutdown: %v", err)
	// }

	// log.Println("HTTP server stopped")