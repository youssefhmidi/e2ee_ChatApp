package services

import "github.com/youssefhmidi/E2E_encryptedConnection/models"

type JwtRequirement struct {
	Secret map[string]string
	Expiry map[string]int
}

func NewJwtService(Secrets map[string]string, Expiries map[string]int) models.JwtService {
	return &JwtRequirement{
		Secret: Secrets,
		Expiry: Expiries,
	}
}

func (jr *JwtRequirement) GetSecret(from string) string {
	return jr.Secret[from]
}

func (jr *JwtRequirement) GetExpiryTime(from string) int {
	return jr.Expiry[from]
}
