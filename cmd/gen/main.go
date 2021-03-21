package main

import (
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
	d := license.Year
	s := &license.SigningData{
		Email:     *email,
		ExpiredAt: time.Now().Add(d),
	}

	fmt.Printf("%s", s.JSON())
}
