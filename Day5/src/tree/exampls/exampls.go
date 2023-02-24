package exampls

import "day5/tree"

var Root1 = &tree.TreeNode{HasToy: true,
	Left: &tree.TreeNode{HasToy: true,
		Left:  &tree.TreeNode{HasToy: false},
		Right: &tree.TreeNode{HasToy: true}},
	Right: &tree.TreeNode{HasToy: true,
		Left:  &tree.TreeNode{HasToy: true},
		Right: &tree.TreeNode{HasToy: false}}}

var Root2 = &tree.TreeNode{HasToy: true,
	Left: &tree.TreeNode{HasToy: true,
		Left: &tree.TreeNode{HasToy: false}},
	Right: &tree.TreeNode{HasToy: true,
		Left:  &tree.TreeNode{HasToy: true},
		Right: &tree.TreeNode{HasToy: true}}}

var Root3 = &tree.TreeNode{HasToy: false,
	Left: &tree.TreeNode{HasToy: true,
		Left:  nil,
		Right: nil},
	Right: &tree.TreeNode{HasToy: false,
		Left: &tree.TreeNode{HasToy: true,
			Left:  nil,
			Right: nil},
		Right: &tree.TreeNode{HasToy: false,
			Left:  nil,
			Right: nil}}}

var Root4 = &tree.TreeNode{HasToy: true,
	Left: &tree.TreeNode{HasToy: true,
		Left:  &tree.TreeNode{HasToy: false},
		Right: &tree.TreeNode{HasToy: true}},
	Right: &tree.TreeNode{HasToy: true,
		Left:  &tree.TreeNode{HasToy: true},
		Right: &tree.TreeNode{HasToy: false}}}

var Root5 = &tree.TreeNode{HasToy: true,
	Left: &tree.TreeNode{HasToy: true,
		Left: &tree.TreeNode{HasToy: false}},
	Right: &tree.TreeNode{HasToy: true,
		Left:  &tree.TreeNode{HasToy: true},
		Right: &tree.TreeNode{HasToy: true}}}

var Root6 = &tree.TreeNode{HasToy: true,
	Left: &tree.TreeNode{HasToy: true,
		Left: &tree.TreeNode{HasToy: false}},
	Right: &tree.TreeNode{HasToy: false,
		Left:  nil,
		Right: nil}}

var Root7 = &tree.TreeNode{HasToy: true, Left: nil, Right: nil}

var Root8 = &tree.TreeNode{HasToy: true,
	Left: &tree.TreeNode{HasToy: true,
		Left:  &tree.TreeNode{HasToy: true},
		Right: &tree.TreeNode{HasToy: false}},
	Right: &tree.TreeNode{HasToy: false,
		Left:  &tree.TreeNode{HasToy: true},
		Right: &tree.TreeNode{HasToy: true}}}

var Root9 = &tree.TreeNode{HasToy: true,
	Left:  &tree.TreeNode{HasToy: true},
	Right: &tree.TreeNode{HasToy: false}}

var Root10 = &tree.TreeNode{HasToy: false,
	Left: &tree.TreeNode{HasToy: true,
		Right: &tree.TreeNode{HasToy: true}},
	Right: &tree.TreeNode{HasToy: false,
		Right: &tree.TreeNode{HasToy: true}}}

var Root11 = &tree.TreeNode{HasToy: true,
	Left: &tree.TreeNode{HasToy: true,
		Left: &tree.TreeNode{HasToy: false,
			Left: &tree.TreeNode{HasToy: true,
				Left: &tree.TreeNode{HasToy: true}},
			Right: &tree.TreeNode{HasToy: true}},
		Right: &tree.TreeNode{HasToy: true,
			Left:  &tree.TreeNode{HasToy: true},
			Right: &tree.TreeNode{HasToy: false}}},
	Right: &tree.TreeNode{HasToy: true,
		Left: &tree.TreeNode{HasToy: true,
			Left:  &tree.TreeNode{HasToy: true},
			Right: &tree.TreeNode{HasToy: false}},
		Right: &tree.TreeNode{HasToy: false,
			Left: &tree.TreeNode{HasToy: true,
				Left: &tree.TreeNode{HasToy: true}},
			Right: &tree.TreeNode{HasToy: true}}}}

var Root12 = &tree.TreeNode{HasToy: true,
	Left: &tree.TreeNode{HasToy: true,
		Left: &tree.TreeNode{HasToy: false,
			Left: &tree.TreeNode{HasToy: true,
				Left: &tree.TreeNode{HasToy: true}},
			Right: &tree.TreeNode{HasToy: true}},
		Right: &tree.TreeNode{HasToy: true,
			Left:  &tree.TreeNode{HasToy: true},
			Right: &tree.TreeNode{HasToy: false}}},
	Right: &tree.TreeNode{HasToy: true,
		Left: &tree.TreeNode{HasToy: true,
			Left: &tree.TreeNode{HasToy: true}},
		Right: &tree.TreeNode{HasToy: false,
			Left: &tree.TreeNode{HasToy: true,
				Left: &tree.TreeNode{HasToy: false}},
			Right: &tree.TreeNode{HasToy: true}}}}
