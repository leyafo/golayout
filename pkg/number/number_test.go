package number

import (
	"math/big"
	"strings"
	"testing"
)

func TestStringAmountToBigInt(t *testing.T){
	bi := FloatStringToBigInt("1", 5)
	c, _ := new(big.Int).SetString("100000",10)
	if bi.Cmp(c) != 0 {
		t.Error("convert failed")
	}
	t.Log(c, bi.String())

	numZero := 18
	bi = FloatStringToBigInt("1", numZero)
	c, _ = new(big.Int).SetString("1"+strings.Repeat("0", numZero),10)
	if bi.Cmp(c) != 0 {
		t.Error("convert failed")
	}
	t.Log(bi.Cmp(c), c.String(), bi.String())
}

func TestBigFloatToStr(t *testing.T) {
	bi := new(big.Int)
	bi.SetString("16999000000000000000", 10)
	str := BigIntToFloatStr(bi, 18)
	t.Log("convert value is: ", str)
	if str != "16.999000000000000000"{
		t.Error("calculate error")
	}

	bi.SetString("200000000000000000000", 10)
	str = BigIntToFloatStr(bi, 18)
	t.Log("convert value is: ", str)
	if str != "200.000000000000000000"{
		t.Error("calculate error")
	}

	bi.SetString("175046674540545800000000", 10)
	str = BigIntToFloatStr(bi, 18)
	t.Log("convert value is: ", str)
	if str != "175046.674540545800000000"{
		t.Error("calculate error")
	}

	bi.SetString("175046674540545887878721", 10)
	str = BigIntToFloatStr(bi, 18)
	t.Log("convert value is: ", str)
	if str != "175046.674540545887878721"{
		t.Error("calculate error")
	}
}