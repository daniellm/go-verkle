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
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352600",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352601",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352602",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352603",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352604",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352640",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352680",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352681",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352682",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352683",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352684",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352685",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352686",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352687",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352688",
		"744f493648c83c5ede1726a0cfbe36d3830fd5b64a820b79ca77fe1593352689",
		"9661ae0db10ecdb9bea3ef0c5fb46bb233cb6ed7404b77e7b0732512ecc60100",
		"9661ae0db10ecdb9bea3ef0c5fb46bb233cb6ed7404b77e7b0732512ecc60101",
		"9661ae0db10ecdb9bea3ef0c5fb46bb233cb6ed7404b77e7b0732512ecc60102",
		"9661ae0db10ecdb9bea3ef0c5fb46bb233cb6ed7404b77e7b0732512ecc60103",
		"9661ae0db10ecdb9bea3ef0c5fb46bb233cb6ed7404b77e7b0732512ecc60104",
	}

	var keys [][]byte
	for _, s := range keyStrings {
		k, _ := hex.DecodeString(s)
		keys = append(keys, k)
	}
	tree := New().(*InternalNode)

	valStrings := []string{
		"0000000000000000000000000000000000000000000000000000000000000000",
		"0000000000000000000000000000000000000000000000000000000000000000",
		"0100000000000000000000000000000000000000000000000000000000000000",
		"f8811a5ee0d54eca4880eaee7b102eae4b3963ff343f50a024c0fd3d367cb8cc",
		"5001000000000000000000000000000000000000000000000000000000000000",
		"0000000000000000000000000000000000000000000000000000000000000000",
		"00608060405234801561001057600080fd5b50600436106100365760003560e0",
		"001c80632e64cec11461003b5780636057361d14610059575b600080fd5b6100",
		"0143610075565b60405161005091906100d9565b60405180910390f35b610073",
		"00600480360381019061006e919061009d565b61007e565b005b600080549050",
		"0090565b8060008190555050565b60008135905061009781610103565b929150",
		"0050565b6000602082840312156100b3576100b26100fe565b5b60006100c184",
		"00828501610088565b91505092915050565b6100d3816100f4565b8252505056",
		"005b60006020820190506100ee60008301846100ca565b92915050565b600081",
		"009050919050565b600080fd5b61010c816100f4565b811461011757600080fd",
		"005b5056fea2646970667358221220404e37f487a89a932dca5e77faaf6ca2de",
		"0000000000000000000000000000000000000000000000000000000000000000",
		"507a878a1965b9e2ce3f00000000000000000000000000000000000000000000",
		"0200000000000000000000000000000000000000000000000000000000000000",
		"c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470",
		"0000000000000000000000000000000000000000000000000000000000000000",
	}
	initialVals := map[string][]byte{}

	for i, s := range valStrings {
		if s == "" {
			continue
		}

		v, _ := hex.DecodeString(s)
		tree.Insert(keys[i], v, nil)

		initialVals[string(keys[i])] = v
	}

	tree.Commit()
	bytes, _ := tree.ToJSON()
	text := string(bytes)

	file, err := os.Create("compatibility.result")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(string(text))
	if err != nil {
		panic(err)
	}
}
