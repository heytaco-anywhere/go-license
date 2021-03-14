package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/heytaco-anywhere/go-license/pkg/license"
)

var (
	company = flag.String("company", "", "The name of company.")
)

func main() {
	flag.Parse()

	// 1 year expiry.
	d := 365 * 24 * time.Hour
	s := &license.SigningData{
		Company:   *company,
		ExpiredAt: time.Now().Add(d),
	}

	signing, _ := json.Marshal(s)
	fmt.Printf("%s", signing)
}
