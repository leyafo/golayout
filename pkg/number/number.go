package number

import (
	"math"
	"math/big"
)

func Bytes2bigInt(bytes []byte) *big.Int {
	bi := new(big.Int)
	bi.SetBytes(bytes)
	return bi
}

func ArrayBytes2BigInts(arrayBytes [][]byte) []*big.Int {
	var bis []*big.Int
	for _, bytes := range arrayBytes {
		bi := new(big.Int)
		bi.SetBytes(bytes)
		bis = append(bis, bi)
	}
	return bis
}

func BigInts2ArrayBytes(bisArray []*big.Int) [][]byte {
	bts := make([][]byte, len(bisArray))
	if bisArray == nil {
		return nil
	}
	for index, bis := range bisArray {
		bts[index] = BigInt2Bytes(bis)
	}
	return bts
}

func BigInt2Bytes(bi *big.Int) []byte {
	if bi == nil {
		return nil
	}
	return bi.Bytes()
}

func FloatStringToBigInt(amount string, decimals int)*big.Int{
	fAmount, _ := new(big.Float).SetString(amount)
	fi, _ := new(big.Float).Mul(fAmount, big.NewFloat(math.Pow10(decimals))).Int(nil)
	return fi
}

func BigIntToFloatString(bi *big.Int, decimals int)string{
	return BigIntToFloatStr(bi, decimals)
}

func StringNumbersToBigIntBytes(stringNumbers []string)[][]byte{
	var result [][]byte
	for _, s := range stringNumbers {
		bi := new(big.Int)
		bi.SetString(s, 10)
		result = append(result, bi.Bytes())
	}
	return result
}

func BigIntToFloatStr(bi *big.Int, decimal int)string{
	pow := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil)
	r := new(big.Rat)
	r.SetFrac(bi, pow)
	return r.FloatString(decimal)
}

func NumberWithDecimal(n int64, decimal int)*big.Int{
	bi := big.NewInt(n)
	return new(big.Int).Mul(bi, BigIntPow10(decimal))
}

func BigIntPow10(pow int)*big.Int{
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(pow)), nil)
}