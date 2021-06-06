package test

import (
	model "github.com/phuockhanhle/familytree/model"
)

type TestCreateNode struct {
	info string
	toInsert model.Person,
	fromPerson: model.Person,
	refOutput: model.Person
}

type TestMatchNode struct {
	info: string,
	input: MatchProperties,
	refOutput: Person,
}

var createNodeTestSuite = []TestCreateNode [
	TestCreateNode{
		info: "insert first person"
		toInsert: ...
		fromPerson: ...
		refOutput: ...
	},
	TestCreateNode{
		info: "insert child from dad"
		toInsert: ...
		fromPerson: ...
		refOutput: ...
	},
	TestCreateNode{
		info: "insert child from mom"
		toInsert: ...
		fromPerson: ...
		refOutput: ...
	},
	TestCreateNode{
		info: "insert dad from child"
		toInsert: ...
		fromPerson: ...
		refOutput: ...
	},
	TestCreateNode{
		info: "insert mom from child"
		toInsert: ...
		fromPerson: ...
		refOutput: ...
	},
	TestCreateNode{
		info: "insert spouse (HasChild: true)"
		toInsert: ...
		fromPerson: ...
		refOutput: ...
	},
	TestCreateNode{
		info: "insert spouse (HasChild: true)"
		toInsert: ...
		fromPerson: ...
		refOutput: ...
	},
]

var matchNodeTestSuite = []TestMatchNode [
	TestMatchNode{
		info: "match people by idtree"
		input: ...
		refOutput: ...
	},
	TestMatchNode{
		info: "match person with ID"
		input: ...
		refOutput: ...
	},
]

func test(testData: TestCreateNode) {
	for _, data := createNodeTestSuite {
		output = insert(toInsert, fromPerson)
		assert(output == refOutput, info)
	}
}

