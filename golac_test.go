package golac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	cases := []struct {
		in   string
		want Document
	}{
		{
			"今天是个好日子",
			Document{
				Sentence{
					[]string{"今天", "是", "个", "好日子"},
					[]string{"TIME", "v", "q", "n"},
				},
			},
		},
		{
			"天气预报说今天要下雨",
			Document{
				Sentence{
					[]string{"天气预报", "说", "今天", "要", "下雨"},
					[]string{"n", "v", "TIME", "v", "v"},
				},
			},
		},
		{
			"下一班地铁马上就要到了",
			Document{
				Sentence{
					[]string{"下", "一班", "地铁", "马上", "就要", "到", "了"},
					[]string{"f", "m", "n", "d", "v", "v", "xc"},
				},
			},
		},
	}

	e := NewLocalExec(nil)

	for _, c := range cases {
		got, _ := e.Run(c.in)
		gotD, _ := Decode(got)

		// assert equality
		assert.Equal(t, gotD, c.want, "they should be equal")
	}
}
