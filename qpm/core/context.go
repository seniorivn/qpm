package core

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	msg "qpm.io/common/messages"
	"runtime"
)

const (
	Version       = "0.10.0"
	PackageFile   = "qpm.json"
	SignatureFile = "qpm.asc"
	Vendor        = "vendor"
	Address       = "pkg.qpm.io:7000"
	LicenseFile   = "LICENSE"
)

var UA = fmt.Sprintf("qpm/%v (%s; %s)", Version, runtime.GOOS, runtime.GOARCH)

type Context struct {
	Log    *log.Logger
	Client msg.QpmClient
}

func NewContext() *Context {
	log := log.New(os.Stderr, "QPM: ", log.LstdFlags)

	creds := credentials.NewClientTLSFromCert(nil, "")
	address := os.Getenv("SERVER")
	if address == "" {
		address = Address
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds), grpc.WithUserAgent(UA))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &Context{
		Log:    log,
		Client: msg.NewQpmClient(conn),
	}
}
