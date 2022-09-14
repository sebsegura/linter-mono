package adder

import "testing"

type input struct {
	a int
	b int
}

func TestAdd(t *testing.T) {
	var (
		tests = []struct {
			name string
			give input
			want int
		}{
			{
				name: "should add two integers",
				give: input{
					a: 2,
					b: 2,
				},
				want: 4,
			},
		}
	)

	for _, tt := range tests {
		r := Add(tt.give.a, tt.give.b)

		if r != tt.want {
			t.Errorf("%d should be %d", r, tt.want)
		}
	}
}
