package day13

import (
	"advent-of-code-2022/testlib"
	"testing"
)

func TestParseWithEmptyString(t *testing.T) {
	root := Node{children: []*Node{}}
	result := Parse(&root, "")
	testlib.AssertEqual(t, &root, result)
	testlib.AssertEqual(t, root.value, 0)
	testlib.AssertEqual(t, len(root.children), 0)
}

func TestParseWithSingleValue(t *testing.T) {
	root := Node{children: []*Node{}}
	Parse(&root, "4")
	testlib.AssertEqual(t, root.value, 0)
	testlib.AssertEqual(t, len(root.children), 1)
	testlib.AssertEqual(t, root.children[0].value, 4)
}

func TestParseWithMultipleValues(t *testing.T) {
	root := Node{children: []*Node{}}
	Parse(&root, "4,9")
	testlib.AssertEqual(t, root.value, 0)
	testlib.AssertEqual(t, len(root.children), 2)
	testlib.AssertEqual(t, root.children[0].value, 4)
	testlib.AssertEqual(t, root.children[1].value, 9)
}

func TestParseWithListValue(t *testing.T) {
	root := Node{children: []*Node{}}
	Parse(&root, "[4,9]")
	testlib.AssertEqual(t, root.value, 0)
	testlib.AssertEqual(t, len(root.children), 1)
	testlib.AssertEqual(t, len(root.children[0].children), 2)
	testlib.AssertEqual(t, root.children[0].children[0].value, 4)
	testlib.AssertEqual(t, root.children[0].children[1].value, 9)
}

func TestParseWithMultipleLists(t *testing.T) {
	root := Node{children: []*Node{}}
	Parse(&root, "1,[4,9],[[7,8,22,23,24],10,11],16")
	testlib.AssertEqual(t, root.value, 0)
	testlib.AssertEqual(t, len(root.children), 4)
	testlib.AssertEqual(t, root.children[0].value, 1)
	testlib.AssertEqual(t, len(root.children[1].children), 2)
	testlib.AssertEqual(t, root.children[1].children[0].value, 4)
	testlib.AssertEqual(t, root.children[1].children[1].value, 9)
	testlib.AssertEqual(t, len(root.children[2].children), 3)
	testlib.AssertEqual(t, len(root.children[2].children[0].children), 5)
	testlib.AssertEqual(t, root.children[2].children[0].children[0].value, 7)
	testlib.AssertEqual(t, root.children[2].children[0].children[1].value, 8)
	testlib.AssertEqual(t, root.children[2].children[0].children[2].value, 22)
	testlib.AssertEqual(t, root.children[2].children[0].children[3].value, 23)
	testlib.AssertEqual(t, root.children[2].children[0].children[4].value, 24)
	testlib.AssertEqual(t, root.children[2].children[1].value, 10)
	testlib.AssertEqual(t, root.children[2].children[2].value, 11)
	testlib.AssertEqual(t, root.children[3].value, 16)
}

func TestPart1(t *testing.T) {
	result, _ := Run("test.txt")
	testlib.AssertEqual(t, result, 13)
}

func TestPart2(t *testing.T) {
	_, result := Run("test.txt")
	testlib.AssertEqual(t, result, 140)
}
