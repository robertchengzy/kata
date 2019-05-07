package utils

import "regexp"

const (
	MOBILE_REG    = `^1[3456789][0-9]{9}$`
	TELEPHONE_REG = `^((010[0-9]{8})|(02[0-9]{9})|(0[3456789][0-9]{9,10}))$`
	NUMBER_REG    = `^((1[3456789][0-9]{9})|(010[0-9]{8})|(02[0-9]{9})|(0[3456789][0-9]{9,10}))$`
)

func CheckMobile(mobile string) bool {
	reg := regexp.MustCompile(MOBILE_REG)
	return reg.MatchString(mobile)
}

func CheckTelephone(telephone string) bool {
	reg := regexp.MustCompile(TELEPHONE_REG)
	return reg.MatchString(telephone)
}

func CheckNumbers(str string) bool {
	reg := regexp.MustCompile(NUMBER_REG)
	return reg.MatchString(str)
}
