package henotest 

//引入包testing
import (
    "testing"
)

//注意方法名称
func TestAdd(t *testing.T) {
    cases := []struct{
        in1, in2, want int
    } {
        {-1, 1, 0},
        {100, 1, 10}, //弄一个错误试试
    }
    for _, c := range cases {
        got := Add(c.in1, c.in2)
        if got != c.want {
            //错误
            t.Errorf("Add(%d, %d) = want %d", c.in1, c.in2, c.want)
        }
    }
}
