package validators

type CustomerValidator interface {
	ValidatePhoneNumber(number string) bool
}
