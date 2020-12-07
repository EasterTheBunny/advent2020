package types

import "testing"

func TestUnmarshal(t *testing.T) {
	p := Passport{}
	data := `iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929 `
	b := []byte(data)

	p.Unmarshal(b)

	if p.IssueYear != "2013" {
		t.Errorf("unexpected value for IssueYear '%s': expected '%s'", p.IssueYear, "2013")
	}

	if p.EyeColor != "amb" {
		t.Errorf("unexpected value for EyeColor '%s': expected %s", p.EyeColor, "amb")
	}

	if p.CountryID != "350" {
		t.Errorf("unexpected value for CountryID '%s': expected %s", p.CountryID, "350")
	}

	if p.ExpirationYear != "2023" {
		t.Errorf("unexpected value for ExpirationYear '%s': expected %s", p.ExpirationYear, "2023")
	}

	if p.ID != "028048884" {
		t.Errorf("unexpected value for ID '%s': expected %s", p.ID, "028048884")
	}

	if p.HairColor != "#cfa07d" {
		t.Errorf("unexpected value for HairColor '%s': expected %s", p.HairColor, "#cfa07d")
	}

	if p.BirthYear != "1929" {
		t.Errorf("unexpected value for BirthYear '%s': expected %s", p.BirthYear, "1929")
	}
}
