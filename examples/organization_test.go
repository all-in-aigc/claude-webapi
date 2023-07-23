package examples

import (
	"fmt"
	"log"
)

func ExampleGetOrganizations() {
	cli := getClient()

	res, err := cli.GetOrganizations()

	if err != nil {
		log.Fatalf("get orgs fail: %v\n", err)
	}

	for _, v := range res.Array() {
		orgid := v.Get("uuid").String()
		log.Printf("orgid: %s\n", orgid)
	}

	fmt.Println(len(res.Array()))
	// Output: 1
}
