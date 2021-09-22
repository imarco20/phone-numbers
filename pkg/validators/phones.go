package validators

import (
	"regexp"
	"strings"
)

var CameroonRegex = regexp.MustCompile("\\(237\\)\\\n ?[2368]\\d{7,8}$\n")
var EthiopiaRegex = regexp.MustCompile("\\(251\\)\\\n ?[1-59]\\d{8}$\n")
var MoroccoRegex = regexp.MustCompile("\\(212\\)\\\n ?[5-9]\\d{8}$\n")
var MozambiqueRegex = regexp.MustCompile("\\(258\\)\\\n ?[28]\\d{7,8}$\n")
var UgandaRegex = regexp.MustCompile("\\(256\\)\\\n ?\\d{9}$\n")

type PhoneNumberValidator struct {
	exp *regexp.Regexp
}

func (n *PhoneNumberValidator) ValidatePhoneNumber(number string) bool {
	return n.exp.MatchString(number)
}

func PhoneNumberValidatorFactory(phoneNumber string) CustomerValidator {
	countryCode := strings.Split(phoneNumber, " ")[0]

	switch countryCode {
	case "(237)":
		return &PhoneNumberValidator{exp: regexp.MustCompile("\\(237\\)\\ ?[2368]\\d{7,8}")}
	case "(251)":
		return &PhoneNumberValidator{exp: regexp.MustCompile("\\(251\\)\\ ?[1-59]\\d{8}")}
	case "(212)":
		return &PhoneNumberValidator{exp: regexp.MustCompile("\\(212\\)\\ ?[5-9]\\d{8}")}
	case "(258)":
		return &PhoneNumberValidator{exp: regexp.MustCompile("\\(258\\)\\ ?[28]\\d{7,8}")}
	case "(256)":
		return &PhoneNumberValidator{exp: regexp.MustCompile("\\(256\\)\\ ?\\d{9}")}
	}
	return nil
}
