package dependencies

type OtpGenerator interface {
	GenerateKeys(login string) (string, string, error)
	ValidOtp(passcode string, key string) bool
}
