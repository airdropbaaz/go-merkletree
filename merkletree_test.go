// Copyright © 2018, 2019 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package merkletree

import (
	"encoding/hex"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wealdtech/go-merkletree/blake2b"
	"github.com/wealdtech/go-merkletree/keccak256"
)

// _byteArray is a helper to turn a string in to a byte array
func _byteArray(input string) []byte {
	x, _ := hex.DecodeString(input)
	return x
}

var tests = []struct {
	// hash type to use
	hashType HashType
	// data to create the node
	data [][]byte
	// expected error when attempting to create the tree
	createErr error
	// root hash after the tree has been created
	root []byte
	// pollards after the tree has been created
	pollards [][][]byte
	// DOT representation of tree
	dot string
	// salt the data?
	salt bool
	// saltedRoot hash after the tree has been created with the salt
	saltedRoot []byte
}{
	{ // 0
		hashType:  blake2b.New(),
		createErr: errors.New("tree must have at least 1 piece of data"),
	},
	{ // 1
		hashType:  blake2b.New(),
		data:      [][]byte{},
		createErr: errors.New("tree must have at least 1 piece of data"),
	},
	{ // 2
		hashType: blake2b.New(),
		data: [][]byte{
			[]byte("Foo"),
			[]byte("Bar"),
		},
		root: _byteArray("e9e0083e456539e9f6336164cd98700e668178f98af147ef750eb90afcf2f637"),
		pollards: [][][]byte{
			[][]byte{
				_byteArray("e9e0083e456539e9f6336164cd98700e668178f98af147ef750eb90afcf2f637"),
			},
		},
		dot:        "digraph MerkleTree {rankdir = TB;node [shape=rectangle margin=\"0.2,0.2\"];\"Foo\" [shape=oval];\"Foo\"->2 [label=\"+00000000\"];2 [label=\"2434…cfac\"];2->1;\"Bar\" [shape=oval];\"Bar\"->3 [label=\"+00000001\"];3 [label=\"f40e…406d\"];2->3 [style=invisible arrowhead=none];3->1;{rank=same;2;3};1 [label=\"8d18…354d\"];}",
		salt:       true,
		saltedRoot: _byteArray("8d18bf1d3ce24f418ded32b60e8881f60ce658bb33b0cfbbc152fcb3aac8354d"),
	},
	{ // 3
		hashType: keccak256.New(),
		data: [][]byte{
			[]byte("Foo"),
			[]byte("Bar"),
		},
		root:       _byteArray("fb6c3a47aacb11c3f7ee3717cfbd43e4ad08da66d2cb049358db7e056baaaeed"),
		dot:        "digraph MerkleTree {rankdir = TB;node [shape=rectangle margin=\"0.2,0.2\"];\"Foo\" [shape=oval];\"Foo\"->2 [label=\"+00000000\"];2 [label=\"0b34…c7b9\"];2->1;\"Bar\" [shape=oval];\"Bar\"->3 [label=\"+00000001\"];3 [label=\"5621…a66c\"];2->3 [style=invisible arrowhead=none];3->1;{rank=same;2;3};1 [label=\"e637…f3b6\"];}",
		salt:       true,
		saltedRoot: _byteArray("e6379b212ee745a62d259ecf0bcccd316d67782a159ded7a88ac788c64b3f3b6"),
	},
	{ // 4
		hashType: blake2b.New(),
		data: [][]byte{
			[]byte("Foo"),
		},
		root: _byteArray("7b506db718d5cce819ca4d33d2348065a5408cc89aa8b3f7ac70a0c186a2c81f"),
		dot:  "digraph MerkleTree {rankdir = TB;node [shape=rectangle margin=\"0.2,0.2\"];\"Foo\" [shape=oval];\"Foo\"->1;1 [label=\"7b50…c81f\"];{rank=same;1};}",
	},
	{ // 5
		hashType: blake2b.New(),
		data: [][]byte{
			[]byte("Foo"),
			[]byte("Bar"),
			[]byte("Baz"),
		},
		root: _byteArray("2c95331b1a38dba3600391a3e864f9418a271388936e54edecd916824bb54203"),
		dot:  "digraph MerkleTree {rankdir = TB;node [shape=rectangle margin=\"0.2,0.2\"];\"Foo\" [shape=oval];\"Foo\"->4;4 [label=\"7b50…c81f\"];4->2;\"Bar\" [shape=oval];\"Bar\"->5;5 [label=\"03c7…6406\"];4->5 [style=invisible arrowhead=none];5->2;\"Baz\" [shape=oval];\"Baz\"->6;6 [label=\"6d5f…2ae0\"];5->6 [style=invisible arrowhead=none];6->3;7 [label=\"0000…0000\"];6->7 [style=invisible arrowhead=none];7->3;{rank=same;4;5;6;7};3 [label=\"113f…1135\"];3->1;2 [label=\"e9e0…f637\"];2->1;1 [label=\"2c95…4203\"];}",
	},
	{ // 6
		hashType: blake2b.New(),
		data: [][]byte{
			[]byte("Foo"),
			[]byte("Bar"),
			[]byte("Baz"),
			[]byte("Qux"),
			[]byte("Quux"),
			[]byte("Quuz"),
		},
		root: _byteArray("9db41fa50e69f2d9ce73367bf8fd249fa960f6a416352f473693ea79540e516d"),
		dot:  "digraph MerkleTree {rankdir = TB;node [shape=rectangle margin=\"0.2,0.2\"];\"Foo\" [shape=oval];\"Foo\"->8;8 [label=\"7b50…c81f\"];8->4;\"Bar\" [shape=oval];\"Bar\"->9;9 [label=\"03c7…6406\"];8->9 [style=invisible arrowhead=none];9->4;\"Baz\" [shape=oval];\"Baz\"->10;10 [label=\"6d5f…2ae0\"];9->10 [style=invisible arrowhead=none];10->5;\"Qux\" [shape=oval];\"Qux\"->11;11 [label=\"d5d1…3cda\"];10->11 [style=invisible arrowhead=none];11->5;\"Quux\" [shape=oval];\"Quux\"->12;12 [label=\"2fec…1151\"];11->12 [style=invisible arrowhead=none];12->6;\"Quuz\" [shape=oval];\"Quuz\"->13;13 [label=\"aff2…62e5\"];12->13 [style=invisible arrowhead=none];13->6;14 [label=\"0000…0000\"];13->14 [style=invisible arrowhead=none];14->7;15 [label=\"0000…0000\"];14->15 [style=invisible arrowhead=none];15->7;{rank=same;8;9;10;11;12;13;14;15};7 [label=\"0eb9…9761\"];7->3;6 [label=\"3705…4377\"];6->3;5 [label=\"f277…7fd5\"];5->2;4 [label=\"e9e0…f637\"];4->2;3 [label=\"5082…d5f0\"];3->1;2 [label=\"7799…9592\"];2->1;1 [label=\"9db4…516d\"];}",
	},
	{ // 7
		hashType: blake2b.New(),
		data: [][]byte{
			[]byte("Foo"),
			[]byte("Bar"),
			[]byte("Baz"),
			[]byte("Qux"),
			[]byte("Quux"),
			[]byte("Quuz"),
			[]byte("FooBar"),
			[]byte("FooBaz"),
			[]byte("BarBaz"),
		},
		root:       _byteArray("bbd7d9d866170e36ea4b2f8a83f528fd28fc090505e1c6b235dbe07b7180c416"),
		salt:       true,
		saltedRoot: _byteArray("65309048c992735e2bac6e442652b32077b819072aa2b83004daa29d248c4dbd"),
		pollards: [][][]byte{
			[][]byte{
				_byteArray("bbd7d9d866170e36ea4b2f8a83f528fd28fc090505e1c6b235dbe07b7180c416"),
			},
			[][]byte{
				_byteArray("bbd7d9d866170e36ea4b2f8a83f528fd28fc090505e1c6b235dbe07b7180c416"),
				_byteArray("0d1f02496807682084e8d4617610883bf0ef53cd8d4a18dd8daa803e2bf4d49e"),
				_byteArray("f407ceb23bce5b1d8cffac7a69f8f93a04f70e705bbb3bfd80795aafdfe50f2b"),
			},
			[][]byte{
				_byteArray("bbd7d9d866170e36ea4b2f8a83f528fd28fc090505e1c6b235dbe07b7180c416"),
				_byteArray("0d1f02496807682084e8d4617610883bf0ef53cd8d4a18dd8daa803e2bf4d49e"),
				_byteArray("f407ceb23bce5b1d8cffac7a69f8f93a04f70e705bbb3bfd80795aafdfe50f2b"),
				_byteArray("7799922ba259c0529cdfb9f974024d45abef9b3190850bc23fc5145cf81c9592"),
				_byteArray("2845629f2f482d7e66c9d88f969825b2811744eb9f7a6119f48d7dd62200c279"),
				_byteArray("3d627dc70a6f885aabe95badbfadce488ba8c74fc012c3850e21ef449fdbc517"),
				_byteArray("85c09af929492a871e4fae32d9d5c36e352471cd659bcdb61de08f1722acc3b1"),
			},
			[][]byte{
				_byteArray("bbd7d9d866170e36ea4b2f8a83f528fd28fc090505e1c6b235dbe07b7180c416"),
				_byteArray("0d1f02496807682084e8d4617610883bf0ef53cd8d4a18dd8daa803e2bf4d49e"),
				_byteArray("f407ceb23bce5b1d8cffac7a69f8f93a04f70e705bbb3bfd80795aafdfe50f2b"),
				_byteArray("7799922ba259c0529cdfb9f974024d45abef9b3190850bc23fc5145cf81c9592"),
				_byteArray("2845629f2f482d7e66c9d88f969825b2811744eb9f7a6119f48d7dd62200c279"),
				_byteArray("3d627dc70a6f885aabe95badbfadce488ba8c74fc012c3850e21ef449fdbc517"),
				_byteArray("85c09af929492a871e4fae32d9d5c36e352471cd659bcdb61de08f1722acc3b1"),
				_byteArray("e9e0083e456539e9f6336164cd98700e668178f98af147ef750eb90afcf2f637"),
				_byteArray("f27788f150c5f45bb618f23034f12d3777f5348ec83ea75e3e81f467b9d67fd5"),
				_byteArray("3705db8dede3991c0846bae4f9de86a2c5957283cdd3434337ee1bb98b2d4377"),
				_byteArray("b12ad3438e85d70be08e4df96595e779f9df6e95fa0ae3026fb149b635f91342"),
				_byteArray("30a175e2105c86f4136aee0a55cf7665b7cff22c76527392bf76b10030414cc6"),
				_byteArray("0eb923b0cbd24df54401d998531feead35a47a99f4deed205de4af81120f9761"),
				_byteArray("0eb923b0cbd24df54401d998531feead35a47a99f4deed205de4af81120f9761"),
				_byteArray("0eb923b0cbd24df54401d998531feead35a47a99f4deed205de4af81120f9761"),
			},
		},
		dot: "digraph MerkleTree {rankdir = TB;node [shape=rectangle margin=\"0.2,0.2\"];\"Foo\" [shape=oval];\"Foo\"->16 [label=\"+00000000\"];16 [label=\"2434…cfac\"];16->8;\"Bar\" [shape=oval];\"Bar\"->17 [label=\"+00000001\"];17 [label=\"f40e…406d\"];16->17 [style=invisible arrowhead=none];17->8;\"Baz\" [shape=oval];\"Baz\"->18 [label=\"+00000002\"];18 [label=\"7c7c…5a1c\"];17->18 [style=invisible arrowhead=none];18->9;\"Qux\" [shape=oval];\"Qux\"->19 [label=\"+00000003\"];19 [label=\"b718…ea0d\"];18->19 [style=invisible arrowhead=none];19->9;\"Quux\" [shape=oval];\"Quux\"->20 [label=\"+00000004\"];20 [label=\"be71…083a\"];19->20 [style=invisible arrowhead=none];20->10;\"Quuz\" [shape=oval];\"Quuz\"->21 [label=\"+00000005\"];21 [label=\"73a8…bc0f\"];20->21 [style=invisible arrowhead=none];21->10;\"FooBar\" [shape=oval];\"FooBar\"->22 [label=\"+00000006\"];22 [label=\"e8b3…5f20\"];21->22 [style=invisible arrowhead=none];22->11;\"FooBaz\" [shape=oval];\"FooBaz\"->23 [label=\"+00000007\"];23 [label=\"970f…8911\"];22->23 [style=invisible arrowhead=none];23->11;\"BarBaz\" [shape=oval];\"BarBaz\"->24 [label=\"+00000008\"];24 [label=\"cb70…bc43\"];23->24 [style=invisible arrowhead=none];24->12;25 [label=\"0000…0000\"];24->25 [style=invisible arrowhead=none];25->12;26 [label=\"0000…0000\"];25->26 [style=invisible arrowhead=none];26->13;27 [label=\"0000…0000\"];26->27 [style=invisible arrowhead=none];27->13;28 [label=\"0000…0000\"];27->28 [style=invisible arrowhead=none];28->14;29 [label=\"0000…0000\"];28->29 [style=invisible arrowhead=none];29->14;30 [label=\"0000…0000\"];29->30 [style=invisible arrowhead=none];30->15;31 [label=\"0000…0000\"];30->31 [style=invisible arrowhead=none];31->15;{rank=same;16;17;18;19;20;21;22;23;24;25;26;27;28;29;30;31};15 [label=\"0eb9…9761\"];15->7;14 [label=\"0eb9…9761\"];14->7;13 [label=\"0eb9…9761\"];13->6;12 [label=\"e9d0…3bad\"];12->6;11 [label=\"b3c6…0cd2\"];11->5;10 [label=\"7473…a3a4\"];10->5;9 [label=\"6dc5…fd2b\"];9->4;8 [label=\"8d18…354d\"];8->4;7 [label=\"85c0…c3b1\"];7->3;6 [label=\"fb16…ac5b\"];6->3;5 [label=\"4847…84bf\"];5->2;4 [label=\"622b…1133\"];4->2;3 [label=\"ee7e…8174\"];3->1;2 [label=\"286d…fea8\"];2->1;1 [label=\"6530…4dbd\"];}",
	},
}

func TestNew(t *testing.T) {
	for i, test := range tests {
		tree, err := NewUsing(test.data, test.hashType, false)
		if test.createErr != nil {
			assert.Equal(t, test.createErr, err, fmt.Sprintf("expected error at test %d", i))
		} else {
			assert.Nil(t, err, fmt.Sprintf("failed to create tree at test %d", i))
			assert.Equal(t, test.root, tree.Root(), fmt.Sprintf("unexpected root at test %d", i))
		}
	}
}

func TestString(t *testing.T) {
	for i, test := range tests {
		if test.createErr == nil {
			tree, err := NewUsing(test.data, test.hashType, false)
			assert.Nil(t, err, fmt.Sprintf("failed to create tree at test %d", i))
			assert.Equal(t, fmt.Sprintf("%x", test.root), tree.String(), fmt.Sprintf("incorrect string representation at test %d", i))
		}
	}
}

func TestDOT(t *testing.T) {
	for i, test := range tests {
		if test.createErr == nil {
			tree, err := NewUsing(test.data, test.hashType, test.salt)
			assert.Nil(t, err, fmt.Sprintf("failed to create tree at test %d", i))
			assert.Equal(t, test.dot, tree.DOT(new(StringFormatter), nil), fmt.Sprintf("incorrect DOT representation at test %d", i))
		}
	}
}

func TestFormatter(t *testing.T) {
	tree, err := New(tests[5].data)
	assert.Nil(t, err, "failed to create tree")
	assert.Equal(t, "digraph MerkleTree {rankdir = TB;node [shape=rectangle margin=\"0.2,0.2\"];\"466f…6f6f\" [shape=oval];\"466f…6f6f\"->4;4 [label=\"7b50…c81f\"];4->2;\"4261…6172\" [shape=oval];\"4261…6172\"->5;5 [label=\"03c7…6406\"];4->5 [style=invisible arrowhead=none];5->2;\"4261…617a\" [shape=oval];\"4261…617a\"->6;6 [label=\"6d5f…2ae0\"];5->6 [style=invisible arrowhead=none];6->3;7 [label=\"0000…0000\"];6->7 [style=invisible arrowhead=none];7->3;{rank=same;4;5;6;7};3 [label=\"113f…1135\"];3->1;2 [label=\"e9e0…f637\"];2->1;1 [label=\"2c95…4203\"];}", tree.DOT(nil, nil), "incorrect default representation")
	assert.Equal(t, "digraph MerkleTree {rankdir = TB;node [shape=rectangle margin=\"0.2,0.2\"];\"466f6f\" [shape=oval];\"466f6f\"->4;4 [label=\"7b506db718d5cce819ca4d33d2348065a5408cc89aa8b3f7ac70a0c186a2c81f\"];4->2;\"426172\" [shape=oval];\"426172\"->5;5 [label=\"03c70c07424c7d85174bf8e0dbd4600a4bd21c00ce34dea7ab57c83c398e6406\"];4->5 [style=invisible arrowhead=none];5->2;\"42617a\" [shape=oval];\"42617a\"->6;6 [label=\"6d5fd2391f8abb79469edf404fd1751a74056ce54ee438c128bba9e680242ae0\"];5->6 [style=invisible arrowhead=none];6->3;7 [label=\"0000000000000000000000000000000000000000000000000000000000000000\"];6->7 [style=invisible arrowhead=none];7->3;{rank=same;4;5;6;7};3 [label=\"113f21ad3be5252e487795473d5e0e221fddf3daee6b5596635428e5feaa1135\"];3->1;2 [label=\"e9e0083e456539e9f6336164cd98700e668178f98af147ef750eb90afcf2f637\"];2->1;1 [label=\"2c95331b1a38dba3600391a3e864f9418a271388936e54edecd916824bb54203\"];}", tree.DOT(new(HexFormatter), new(HexFormatter)), "incorrect default representation")
}
