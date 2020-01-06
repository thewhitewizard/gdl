package gdl

import (
	"flag"
)


func getTestClient() (*GDLClient,error) {
	flag.Parse()
	client, err := New(nil)
	return client,err
}