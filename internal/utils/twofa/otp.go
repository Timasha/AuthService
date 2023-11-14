package twofa

import "github.com/pquerna/otp/totp"

type DefaultOtp struct{}

func (d *DefaultOtp) GenerateKeys(login string) (string, string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      login,
		AccountName: login,
		SecretSize:  15,
	})
	return key.Secret(), key.URL(), err
}

func (d *DefaultOtp) ValidOtp(passcode string, key string) bool {
	return totp.Validate(passcode, key)
}
