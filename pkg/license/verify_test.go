package license

import (
	"log"
	"testing"
)

func TestLicense_verifyKey(t *testing.T) {
	t.Run("verify the license", func(tt *testing.T) {
		const (
			// It was generated to test.
			key = "key/eyJjb21wYW55IjoiQ0kiLCJleHBpcmVkQXQiOiIyMDIyLTAzLTE0VDEzOjAyOjU3LjI3MTc0NCswOTowMCJ9.P9dIQg4IaVCi5qOX-pgd7AZI04vtbIYlDT4vUlRKskjVMsztm4VMe_0JVR8ak6PQWVbRaf3pBQj8xwH3Nu2RsvNj5unt6UM6UiA_sPgjbzi-hGtsl9IsnrNIcPrgk5_3xEg_-jfXKIrgXdjBSMUTLFO8VgcQYDALMkoMexYGuHBj-v4mHF_55cMbXB-JWiZHFOzXEXaBOYyPYUpGH-XaYgCAA1L9TF20EntQ3a7Tcv8b3d34ApwrgmPHxDJ4J3kz5uQ1tKhwQ8O9o18IaygrS7czWo952AVC1HW9RUj9cz4THWXmdsI85i2cXTyq9OcbNNPdjKDC6fd6VOvNF4a2vA=="
		)

		signing, err := Verfiy(key)
		if err != nil {
			tt.Errorf("Verify returns error = %v", err)
			return
		}

		if signing.Company != "CI" {
			tt.Errorf("Verfiy() = %v, expeced %v", signing.Company, "CI")
			return
		}

		if signing.ExpiredAt.IsZero() {
			tt.Errorf("Verfiy() ExpiredAt is invalid: %v", signing.ExpiredAt)
		}

		log.Printf("signing data: %v", signing)
	})
}
