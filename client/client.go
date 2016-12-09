package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/jnst/go-web/rpc/gen"
)

func main() {
	client := &http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := client.Post("http://localhost:8080/rpc", "text/plain", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}

	gachaResp := &pb.GachaResponse{}
	err = proto.Unmarshal(b, gachaResp)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(resp.StatusCode, gachaResp.GetCards()[0].Name)
}
