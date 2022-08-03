package sql

import (
	"crypto/tls"
	"fmt"

	"github.com/textileio/textile/api/common"
	"golang.org/x/vuln/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	query = `tableland create "CREATE TABLE AssetTracker (id INT PRIMARY KEY, name TEXT, type TEXT, cid TEXT, provider TEXT, url TEXT);" --description="Filebase Asset Tracker`
)

func setup() {
	creds := credentials.NewTLS(&tls.Config{})
	auth := common.Credentials{}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(auth)}
	cli, err := client.NewClient("api.hub.textile.io:443", opts...)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success!")
}
