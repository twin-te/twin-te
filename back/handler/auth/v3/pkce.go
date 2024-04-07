package authv3

import "golang.org/x/oauth2"

var (
	verifier            = oauth2.GenerateVerifier()
	verifierOption      = oauth2.VerifierOption(verifier)
	s256ChallengeOption = oauth2.S256ChallengeOption(verifier)
)
