package donationusecase

import (
	"context"
	"log"
	"sync"
	"time"

	authmodule "github.com/twin-te/twinte-back/module/auth"
	donationmodule "github.com/twin-te/twinte-back/module/donation"
	donationport "github.com/twin-te/twinte-back/module/donation/port"
)

var _ donationmodule.UseCase = (*impl)(nil)

type impl struct {
	a authmodule.AccessController
	f donationport.Factory
	g donationport.Gateway
	r donationport.Repository

	contributorsCache      []donationmodule.Contributor
	contributorsCacheMutex sync.RWMutex

	totalAmountCache      int
	totalAmountCacheMutex sync.RWMutex
}

func New(a authmodule.AccessController, f donationport.Factory, g donationport.Gateway, r donationport.Repository) *impl {
	uc := &impl{a: a, f: f, g: g, r: r}

	go func() {
		for {
			log.Println("update contributors cache")
			if err := uc.updateContributorsCache(context.Background()); err != nil {
				log.Printf("failed to update contributors cache, %+v", err)
			}
			<-time.After(time.Hour)
		}
	}()

	go func() {
		for {
			log.Println("update total amount cache")
			if err := uc.updateTotalAmountCache(context.Background()); err != nil {
				log.Printf("failed to update total amount cache, %+v", err)
			}
			<-time.After(time.Hour)
		}
	}()

	return uc
}
