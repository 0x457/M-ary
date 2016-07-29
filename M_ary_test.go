package M_ary

import (
	"testing"
	"fmt"
)



func TestALL(t *testing.T) {
	for int10 := int64(1); int10 < 1000000; int10 = int10 + 1 {
		h36 := Dec2ThirtySix(int10)
		i10 := ThirtySix2Dec(h36)
		if i10 != int64(int10) {
			fmt.Printf("convert err %d => %s => %d \n", int10, h36, i10)
		}
		h36_add := Dec2ThirtySix(i10 + 1)
		fmt.Printf("convert  %s => %s \n", h36, h36_add)
	}
}

