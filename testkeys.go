package testkeys

import (
	"bufio"
	"bytes"
)

var Keys = []string{
	"empty",

	"10vl5",
	"11vl5",

	"10ll16k",

	"300vl50",

	"20kl10",
	"20kvl10",

	"50kl10",
	"50kvl10",

	"200kweb2",

	"1mvl5_10",
}

func Load(fn string) []string {

	rst := []string{}

	bs, err := Asset("data/" + fn)
	if err != nil {
		panic("Asset: " + fn + ":" + err.Error())
	}

	r := bytes.NewReader(bs)
	s := bufio.NewScanner(r)
	for s.Scan() {
		str := s.Text()
		rst = append(rst, str)
	}

	err = s.Err()
	if err != nil {
		panic("scan: " + fn + ":" + err.Error())
	}

	return rst
}
