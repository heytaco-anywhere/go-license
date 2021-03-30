package license

import (
	"log"
	"testing"
)

func TestLicense_verifyKey(t *testing.T) {
	t.Run("verify the license", func(tt *testing.T) {
		const (
			// It was generated to test.
			key = "key/eyJlbWFpbCI6ImhleXRhY28uYW55d2hlcmVAZ21haWwuY29tIiwiZXhwaXJlZEF0IjoiMjAyMS0wNC0yMlQwODoyMjo0OC43MjY1NiswOTowMCJ9.aaPIotQI3tJmdlRviyNDYfUsJw0lhhoYm_8uzlv_ma0abL0gYZW0L6CmV4uV3y1SsdaoXKne_fitK-zGLueIFKH4We4Xjetl0q8Qn_snGUQ3aZDG5dj6KlCXw5DFmF36o2520kjH6gES9yWOMngTTv4Wwj8F4ME4uPYMQFQzw8QZfY3C0fNQcaEKu2cn3cU5p5ZQKoXBsyl4zjnkWTkKo2MLY-ez8ORE9E76hbiW69-ST2Lb6yVd1iLpTeh2hfGh4JWyhW7v3rnjO90bvJ2P0tixoSY-LykBlEn_RRgy3eezMPNFwsZQx1GpfJyXOl9aBT3JBsQxjlKjOO_aKel97A=="
		)

		signing, err := Verfiy(key)
		if err != nil {
			tt.Errorf("Verify returns error = %v", err)
			return
		}

		if signing.Email != "heytaco.anywhere@gmail.com" {
			tt.Errorf("Verfiy() = %v, expeced %v", signing.Email, "CI")
			return
		}

		if signing.ExpiredAt.IsZero() {
			tt.Errorf("Verfiy() ExpiredAt is invalid: %v", signing.ExpiredAt)
		}

		log.Printf("signing data: %v", signing)
	})
}
