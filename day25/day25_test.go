package day25

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestConvertSnafuToInt(t *testing.T) {
	testlib.AssertEqual(t, ConvertSnafuToInt("1"), 1)
	testlib.AssertEqual(t, ConvertSnafuToInt("2"), 2)
	testlib.AssertEqual(t, ConvertSnafuToInt("1="), 3)
	testlib.AssertEqual(t, ConvertSnafuToInt("1-"), 4)
	testlib.AssertEqual(t, ConvertSnafuToInt("10"), 5)
	testlib.AssertEqual(t, ConvertSnafuToInt("11"), 6)
	testlib.AssertEqual(t, ConvertSnafuToInt("12"), 7)
	testlib.AssertEqual(t, ConvertSnafuToInt("2="), 8)
	testlib.AssertEqual(t, ConvertSnafuToInt("2-"), 9)
	testlib.AssertEqual(t, ConvertSnafuToInt("20"), 10)
	testlib.AssertEqual(t, ConvertSnafuToInt("1=0"), 15)
	testlib.AssertEqual(t, ConvertSnafuToInt("1-0"), 20)
	testlib.AssertEqual(t, ConvertSnafuToInt("1=11-2"), 2022)
	testlib.AssertEqual(t, ConvertSnafuToInt("1121-1110-1=0"), 314159265)
}

func TestConvertIntToSnafu(t *testing.T) {
	testlib.AssertEqual(t, ConvertIntToSnafu(1), "1")
	testlib.AssertEqual(t, ConvertIntToSnafu(2), "2")
	testlib.AssertEqual(t, ConvertIntToSnafu(3), "1=")
	testlib.AssertEqual(t, ConvertIntToSnafu(4), "1-")
	testlib.AssertEqual(t, ConvertIntToSnafu(5), "10")
	testlib.AssertEqual(t, ConvertIntToSnafu(6), "11")
	testlib.AssertEqual(t, ConvertIntToSnafu(7), "12")
	testlib.AssertEqual(t, ConvertIntToSnafu(8), "2=")
	testlib.AssertEqual(t, ConvertIntToSnafu(9), "2-")
	testlib.AssertEqual(t, ConvertIntToSnafu(10), "20")
	testlib.AssertEqual(t, ConvertIntToSnafu(15), "1=0")
	testlib.AssertEqual(t, ConvertIntToSnafu(20), "1-0")
	testlib.AssertEqual(t, ConvertIntToSnafu(2022), "1=11-2")
	testlib.AssertEqual(t, ConvertIntToSnafu(314159265), "1121-1110-1=0")
}

func TestPart1(t *testing.T) {
	result, _ := Run("test.txt")
	testlib.AssertEqual(t, result, "2=-1=0")
}

func TestPart2(t *testing.T) {
	_, result := Run("test.txt")
	testlib.AssertEqual(t, result, "")
}
