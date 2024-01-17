// This is free and unencumbered software released into the public domain.
//
// Anyone is free to copy, modify, publish, use, compile, sell, or
// distribute this software, either in source code form or as a compiled
// binary, for any purpose, commercial or non-commercial, and by any
// means.
//
// In jurisdictions that recognize copyright laws, the author or authors
// of this software dedicate any and all copyright interest in the
// software to the public domain. We make this dedication for the benefit
// of the public at large and to the detriment of our heirs and
// successors. We intend this dedication to be an overt act of
// relinquishment in perpetuity of all present and future rights to this
// software under copyright law.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.
//
// For more information, please refer to <https://unlicense.org>

package verkle

import (
	"encoding/hex"
	"os"
	"testing"
)

// This test is used to compare the state of the tree with other verkle implementations.
// More tests will be added in the future.
// The state of the tree is written to "compatibility.result"
// TODO: standardize the format of the tree dump so it can be easily compared

func TestCompatibility(t *testing.T) {
	keyStrings := []string{
		"0000000000000000000000000000000000000000000000000000000000000000",
		"000102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f",
		"1100000000000000000000000000000000000000000000000000000000000000",
		"2200000000000000000000000000000000000000000000000000000000000000",
		"2211000000000000000000000000000000000000000000000000000000000000",
		"3300000000000000000000000000000000000000000000000000000000000000",
		"3300000000000000000000000000000000000000000000000000000000000001",
		"33000000000000000000000000000000000000000000000000000000000000ff",
		"4400000000000000000000000000000000000000000000000000000000000000",
		"4400000011000000000000000000000000000000000000000000000000000000",
		"5500000000000000000000000000000000000000000000000000000000000000",
		"5500000000000000000000000000000000000000000000000000000000001100",
	}

	var keys [][]byte
	for _, s := range keyStrings {
		k, _ := hex.DecodeString(s)
		keys = append(keys, k)
	}
	tree := New().(*InternalNode)

	valStrings := []string{
		"000000000000000000000000000000000123456789abcdef0123456789abcdef",
		"0000000000000000000000000000000000000000000000000000000000000002",
		"0000000000000000000000000000000000000000000000000000000000000003",
		"0000000000000000000000000000000000000000000000000000000000000004",
		"0000000000000000000000000000000000000000000000000000000000000005",
		"0000000000000000000000000000000000000000000000000000000000000006",
		"0000000000000000000000000000000000000000000000000000000000000007",
		"0000000000000000000000000000000000000000000000000000000000000008",
		"0000000000000000000000000000000000000000000000000000000000000009",
		"000000000000000000000000000000000000000000000000000000000000000a",
		"000000000000000000000000000000000000000000000000000000000000000b",
		"000000000000000000000000000000000000000000000000000000000000000c",
	}

	for i, s := range valStrings {
		v, _ := hex.DecodeString(s)
		tree.Insert(keys[i], v, nil)
	}

	tree.Commit()
	bytes, _ := tree.ToJSON()
	text := string(bytes)

	file, err := os.Create("testResults/compatibility1.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(string(text))
	if err != nil {
		panic(err)
	}

	//////////////////////////////////////////////////////////////////////////////

	keyStrings = []string{
		"0000000000000000000000000000000000000000000000000000000000000000",
		"1100000000000000000000000000000000000000000000000000000000010000",
		"4400000011000000000000000000000000000000000000000000000000000001",
	}

	keys = nil
	for _, s := range keyStrings {
		k, _ := hex.DecodeString(s)
		keys = append(keys, k)
	}

	valStrings = []string{
		"0000000000000000000000000000000000000000000000000000000000000011",
		"0000000000000000000000000000000000000000000000000000000000000012",
		"0000000000000000000000000000000000000000000000000000000000000013",
	}

	for i, s := range valStrings {
		v, _ := hex.DecodeString(s)
		tree.Insert(keys[i], v, nil)
	}

	tree.Commit()
	bytes, _ = tree.ToJSON()
	text = string(bytes)

	file, err = os.Create("testResults/compatibility2.after-update.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(string(text))
	if err != nil {
		panic(err)
	}

	//////////////////////////////////////////////////////////////////////////////

	keyStrings = []string{
		"1100000000000000000000000000000000000000000000000000000000010000",
		"2211000000000000000000000000000000000000000000000000000000000000",
		"5500000000000000000000000000000000000000000000000000000000001100",
	}

	for _, s := range keyStrings {
		k, _ := hex.DecodeString(s)
		tree.Delete(k, nil)
	}

	tree.Commit()
	bytes, _ = tree.ToJSON()
	text = string(bytes)

	file, err = os.Create("testResults/compatibility3.after-delete.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(string(text))
	if err != nil {
		panic(err)
	}

}
