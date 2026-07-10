package authdomain

import "time"

const AuthChallengeLifeTime = 5 * time.Minute

type AuthChallenge struct {
	ID        string
	Provider  Provider
	Nonce     string
	ExpiredAt time.Time
}
