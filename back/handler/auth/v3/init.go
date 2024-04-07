package authv3

import (
	"log"
	"time"
)

func init() {
	go func() {
		// regenerate apple client secret once a month
		for {
			appleClientSecret, err := generateAppleClientSecret()
			if err != nil {
				log.Printf("failed to generate apple client secret, %+v", err)
			}
			appleOAuth2Config.ClientSecret = appleClientSecret
			<-time.After(30 * 24 * time.Hour)
		}
	}()
}
