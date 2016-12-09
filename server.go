package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	stats "github.com/fukata/golang-stats-api-handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/proto"
	pb "github.com/jnst/go-web/rpc/gen"
)

// RPCHandler handles rpc request.
func RPCHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	gachaResp := createResponse()
	data, err := proto.Marshal(gachaResp)
	if err != nil {
		log.Fatalln(err.Error())
	}

	w.Write(data)
}

// NotFoundHandler handles 404 Not Found.
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func createResponse() *pb.GachaResponse {
	c := &pb.Card{
		CardId: proto.Int32(1),
		Name:   proto.String("Long Sword"),
		Type:   pb.Type_TROOP.Enum(),
		Rarity: pb.Rarity_COMMON.Enum(),
		Cost:   proto.Int32(3),
		Parameter: &pb.Parameter{
			Hp:       proto.Int32(1200),
			Damage:   proto.Int32(420),
			HitSpeed: proto.Float32(1.5),
			Range:    proto.Int32(2),
			Radius:   proto.Float32(2.2),
			Lifetime: proto.Int32(30),
		},
	}
	resp := &pb.GachaResponse{
		Cards: []*pb.Card{c},
	}
	return resp
}

// AddContext returns context added handler.
func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)
		cookie, _ := r.Cookie("username")
		if cookie != nil {
			ctx := context.WithValue(r.Context(), "Username", cookie.Value)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func main() {
	stats.PrettyPrintEnabled()

	db, err := sql.Open("mysql", "root:root@tcp(castle.local.infra:3306)/test-castle")
	if err != nil {
		panic(err.Error())
	}
	if db.Ping() != nil {
		panic(err.Error())
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", NotFoundHandler)
	mux.HandleFunc("/api/stats", stats.Handler)
	mux.HandleFunc("/rpc", RPCHandler)

	log.Println("Start server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
