package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/heytaco-anywhere/go-license/pkg/license"
)

var (
	email = flag.String("email", "heytaco.anywhere@gmail.com", "The email of customer.")
)

func main() {
	flag.Parse()

	// 1 year expiry.
	d := 365 * 24 * time.Hour
	s := &license.SigningData{
		Email:   *email,
		ExpiredAt: time.Now().Add(d),
	}

	signing, _ := json.Marshal(s)
	fmt.Printf("%s", signing)
}
