package gdl

import "testing"
import "fmt"

func TestAllParkings(t *testing.T) {

	client, errC := getTestClient()
	if errC != nil {
		t.Fatal(errC)
	}
	parkings, err := client.GetParkings()
	
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(parkings)
}

 