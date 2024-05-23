package twofa

import (
	"github.com/Timasha/AuthService/utils/consts"
	"github.com/pquerna/otp/totp"
)

type DefaultOtp struct {
	cfg Config
}

func New(cfg Config) (d *DefaultOtp) {
	if cfg.OrganizationName == "" {
		cfg.OrganizationName = "defaultOrganization"
	}

	return &DefaultOtp{
		cfg: cfg,
	}
}

func (d *DefaultOtp) GenerateKeys(login string) (secret string, link string, err error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      d.cfg.OrganizationName,
		AccountName: login,
		SecretSize:  consts.OTPSecretSize,
	})

	if err != nil {
		return "", "", err
	}

	return key.Secret(), key.URL(), err
}

func (d *DefaultOtp) ValidOtp(passcode string, key string) bool {
	return totp.Validate(passcode, key)
}
