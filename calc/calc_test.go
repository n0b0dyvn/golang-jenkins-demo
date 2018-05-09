package calc

import "testing"

func TestSum(t *testing.T) {
	for _,num := range []struct {
		a,b,c int
	}{
		{1,2,3},
		{3,2,5},
		{5,5,10},
	}{
		sum:=Sum(num.a,num.b)
		if num.c != sum {
			t.Errorf("Sum(%d,%d) == %d, want %d\n",num.a,num.b,sum,num.c)
		}
	}
}
