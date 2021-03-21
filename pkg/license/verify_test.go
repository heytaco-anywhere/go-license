package license

import (
	"log"
	"testing"
)

func TestLicense_verifyKey(t *testing.T) {
	t.Run("verify the license", func(tt *testing.T) {
		const (
			// It was generated to test.
			key = "key/eyJlbWFpbCI6ImhleXRhY28uYW55d2hlcmVAZ21haWwuY29tIiwiZXhwaXJlZEF0IjoiMjAyMi0wMy0yMVQxNTozNTozMy45OTM5OTcrMDk6MDAifQ==.azjHS56rTd5EZdDTF8Q_UJjaX12DBLh8FYcF80Pl3tvrRdR3kGktWr0JenF0FxWbVAxX2ixpzc3ocleuu4S-4znVBw1zMPDgiqUKFuHrsH-Pzu-i9We2tOTqH78sAkAphtm3w3qAOPNm_cv6Frzpu21R4cbU7odVlf9-xPnaH__UYwWN81sxp1CBJ4fRk2GrWyaPuuyDPNUZ7wIsXbfrIKxLhjxYaS8cJ7RehRZoYDfWPA_Z2BiPRACgEpdNUsyihdfAEnj6yRQrFsN52AfqQ00Rb2uOHT5rFdQy2nW9UtBXRga5Mw52GbT0Y__nvBRpCH5cYjS9LSfIMdsfh5P1EA=="
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
