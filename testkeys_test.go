package testkeys

import (
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
		got := Load(".", c.input)
		ta.Equal(c.want, got, "%d-th: case: %+v", i+1, c)
	}
}
