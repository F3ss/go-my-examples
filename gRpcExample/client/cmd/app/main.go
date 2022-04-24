package main

import (
	"fmt"
	"grpcclient/pkg/api"

	"flag"
	"log"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatalln("Not enough values")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	checkErr(err)

	y, err := strconv.Atoi(flag.Arg(1))
	checkErr(err)

	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	checkErr(err)

	client := api.NewAdderClient(conn)

	res, err := client.Add(context.Background(), &api.AddRequest{X: int32(x), Y: int32(y)})
	checkErr(err)

	fmt.Println(res.GetResult())
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
