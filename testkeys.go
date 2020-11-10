package testkeys

import (
	"bufio"
	"bytes"
)

func Load(fn string) []string {

	rst := []string{}

	bs, err := Asset(fn)
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
