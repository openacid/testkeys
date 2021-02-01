package testkeys

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {

	ta := require.New(t)

	cases := []struct {
		input string
		want  []string
	}{
		{
			"10vl5",
			[]string{
				``,
				`)`,
				`76{?G`,
				`@`,
				`K?L@`,
				`VC`,
				`]e6`,
				`^82%.`,
				`xcY`,
				`}>ih`,
			},
		},
	}

	for i, c := range cases {
		got, err := Asset(c.input)
		ta.Nil(err)

		want := []byte(strings.Join(c.want, "\n") + "\n")
		ta.Equal(want, got, "%d-th: case: %+v: Asset", i+1, c)

		{
			got := Load(c.input)
			ta.Equal(c.want, got, "%d-th: case: %+v compare strings", i+1, c)
		}
	}

	{ // not found
		_, err := Asset("foo")
		ta.NotNil(err)
	}

	names := AssetNames()
	want := []string{
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

		"870k_ip4_hex",

		"1mvl5_10",
	}

	sort.Strings(names)
	sort.Strings(want)

	ta.Equal(want, names, "check names")
}
