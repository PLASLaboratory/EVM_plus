package vm

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/common"
)

// FunctionNode represents node use for stack
type FunctionNode struct {
	FunctionName    string
	ContractName    string
	ContractAddress common.Address
	FunctionID      []byte
}

// Request represents printing method
func (n *FunctionNode) String() string {
	return fmt.Sprint(n.FunctionName, n.ContractName, string(n.ContractAddress.Hex()), string(n.FunctionID))
	// return "a"
}

// NewCallStack represents configuring new stack
func NewCallStack() *CallStack {
	return &CallStack{}
}

// CallStack represents structure for stack
type CallStack struct {
	functions []*FunctionNode
	count     int
}

// Push represents put nodes in stack
func (s *CallStack) Push(n *FunctionNode) {
	s.functions = append(s.functions[:s.count], n)
	s.count++
	fmt.Printf("In Stack function signature: %x\n", n.FunctionID)
	//fmt.Printf("contract address: %s\n",n.ContractAddress);
}

// Pop represents remove nodes in stack
func (s *CallStack) Pop() *FunctionNode {
	if s.count == 0 {
		fmt.Println("nothing to pop")
		return nil
	}
	s.count--

	fmt.Printf("Stack pop-> function signature: %x\n", s.functions[s.count].FunctionID)
	//fmt.Printf("Contract Address: %s\n",s.functions[s.count].ContractAddress);
	return s.functions[s.count]
}

// Find finds nodes in stack
// func (s *CallStack) Find(n *s.FunctionID) bool {
// 	for f := range s.functions {
// 		if s.functions[f].FunctionID == n { //같은게 있는경우
// 			return true
// 		}
// 	}
// }

//IsCallFromConstructor is return is function is call from constructor
func (s *CallStack) IsCallFromConstructor(a common.Address) bool {
	return s.checkPrevious(a)
}

// CheckPrevious check the previous nodes in stack
func (s *CallStack) checkPrevious(n common.Address) bool {
	var a = false
	if s.count == 0 || s.count == 1 {
		fmt.Println("nothing to pop")
	} else {
		for i := s.count; i > 0; i-- {
			if reflect.DeepEqual(s.functions[i].ContractAddress, n) { //같은게 있는경우
				a = true //return previos node's address
				break
			}
		}
	}
	return a
}

//GetFunctionID is get function ID from function signature
func (s *CallStack) GetFunctionID(input []byte) []byte {
	return getDataBig(input, big.NewInt(0), big.NewInt(4))
}

//ConvertAddresstoString is convert byte slice to string
// func (s *CallStack) ConvertAddresstoString(address common.Address) string {
// 	return common.ToHex(address)
// }

// func (s *CallStack) Remove(n *FunctionID) {
// 	for f in s.functions {

// 	}
// 	for i := range s.functions {

// 		if myconfig[i].Key == FunctionID {
// 			return true
// 		}
// 	}
// 	s.functions = append(s.functions[:s.count], n)
// 	s.count++
// 	fmt.Printf("In Stack function signature: %x\n", n.FunctionID)
// 	//fmt.Printf("contract address: %s\n",n.ContractAddress);
// }
