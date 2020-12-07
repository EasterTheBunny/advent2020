package types

import (
	"fmt"
	"reflect"
	"regexp"
)

// Passport is a structure that contains passport data
type Passport struct {
	BirthYear      string `file:"byr"`
	IssueYear      string `file:"iyr"`
	ExpirationYear string `file:"eyr"`
	Height         string `file:"hgt"`
	HairColor      string `file:"hcl"`
	EyeColor       string `file:"ecl"`
	ID             string `file:"pid"`
	CountryID      string `file:"cid"`
}

// Validate returns whether a passport has valid field data
func (p *Passport) Validate() bool {
	if p.BirthYear == "" {
		return false
	}

	if p.IssueYear == "" {
		return false
	}

	if p.ExpirationYear == "" {
		return false
	}

	if p.Height == "" {
		return false
	}

	if p.HairColor == "" {
		return false
	}

	if p.EyeColor == "" {
		return false
	}

	if p.ID == "" {
		return false
	}

	/*
		if p.CountryID == "" {
			return false
		}
	*/

	return true
}

// Unmarshal reads bytes of data from byte source
func (p *Passport) Unmarshal(b []byte) {
	// the regex below needs a space at the end to make a match on the last
	// item. this ensures that there is a space
	c := fmt.Sprintf("%s ", b)

	v := reflect.ValueOf(p).Elem()
	typeOf := v.Type()
	tagReg := regexp.MustCompile(`file:"(.*)"`)

	for i := 0; i < v.NumField(); i++ {
		tMatch := tagReg.FindStringSubmatch(fmt.Sprintf("%s", typeOf.Field(i).Tag))

		if len(tMatch) > 1 {
			regex := regexp.MustCompile(fmt.Sprintf(`%s:([#\w]+)\s`, tMatch[1]))
			fMatch := regex.FindStringSubmatch(c)

			if len(fMatch) > 1 {
				newVal := fmt.Sprintf("%s", fMatch[1])
				v.Field(i).Set(reflect.ValueOf(newVal))
			}
		}
	}
}
