package helper

import "regexp"

func IsValidCpf(cpf string) bool {
	return MatchPattern(`^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$`, cpf)
}

func IsValidCnpj(cnpj string) bool {
	return MatchPattern(`^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`, cnpj)
}

func IsValidEmail(email string) bool {
	return MatchPattern(`^[a-z0-9+_.-]+@[a-z0-9.-]+$`, email)
}

func IsValidPhone(phone string) bool {
	return MatchPattern(`^((?:\+?55)?)([1-9][0-9])(9[0-9]{8})$`, phone)
}

func MatchPattern(pattern string, needle string) bool {
	return regexp.MustCompile(pattern).MatchString(needle)
}
