// Code generated by ethgo/abigen. DO NOT EDIT.
// Hash: 3d1ecdf4aa6a2c578e0c3bbb14cc28ae2c8ebc4495f7d6128959f961afd0f635
// Version: 0.1.1
package ens

import (
	"fmt"
	"math/big"

	"github.com/cloudwalk/ethgo"
	"github.com/cloudwalk/ethgo/contract"
	"github.com/cloudwalk/ethgo/jsonrpc"
)

var (
	_ = big.NewInt
	_ = jsonrpc.NewClient
)

// Resolver is a solidity contract
type Resolver struct {
	c *contract.Contract
}

// DeployResolver deploys a new Resolver contract
func DeployResolver(provider *jsonrpc.Client, from ethgo.Address, args []interface{}, opts ...contract.ContractOption) (contract.Txn, error) {
	return contract.DeployContract(abiResolver, binResolver, args, opts...)
}

// NewResolver creates a new instance of the contract at a specific address
func NewResolver(addr ethgo.Address, opts ...contract.ContractOption) *Resolver {
	return &Resolver{c: contract.NewContract(addr, abiResolver, opts...)}
}

// calls

// ABI calls the ABI method in the solidity contract
func (r *Resolver) ABI(node [32]byte, contentTypes *big.Int, block ...ethgo.BlockNumber) (retval0 *big.Int, retval1 []byte, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = r.c.Call("ABI", ethgo.EncodeBlock(block...), node, contentTypes)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["contentType"].(*big.Int)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	retval1, ok = out["data"].([]byte)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 1")
		return
	}
	
	return
}

// Addr calls the addr method in the solidity contract
func (r *Resolver) Addr(node [32]byte, block ...ethgo.BlockNumber) (retval0 ethgo.Address, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = r.c.Call("addr", ethgo.EncodeBlock(block...), node)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["ret"].(ethgo.Address)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Content calls the content method in the solidity contract
func (r *Resolver) Content(node [32]byte, block ...ethgo.BlockNumber) (retval0 [32]byte, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = r.c.Call("content", ethgo.EncodeBlock(block...), node)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["ret"].([32]byte)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Name calls the name method in the solidity contract
func (r *Resolver) Name(node [32]byte, block ...ethgo.BlockNumber) (retval0 string, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = r.c.Call("name", ethgo.EncodeBlock(block...), node)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["ret"].(string)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// Pubkey calls the pubkey method in the solidity contract
func (r *Resolver) Pubkey(node [32]byte, block ...ethgo.BlockNumber) (retval0 [32]byte, retval1 [32]byte, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = r.c.Call("pubkey", ethgo.EncodeBlock(block...), node)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["x"].([32]byte)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	retval1, ok = out["y"].([32]byte)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 1")
		return
	}
	
	return
}

// SupportsInterface calls the supportsInterface method in the solidity contract
func (r *Resolver) SupportsInterface(interfaceID [4]byte, block ...ethgo.BlockNumber) (retval0 bool, err error) {
	var out map[string]interface{}
	var ok bool

	out, err = r.c.Call("supportsInterface", ethgo.EncodeBlock(block...), interfaceID)
	if err != nil {
		return
	}

	// decode outputs
	retval0, ok = out["0"].(bool)
	if !ok {
		err = fmt.Errorf("failed to encode output at index 0")
		return
	}
	
	return
}

// txns

// SetABI sends a setABI transaction in the solidity contract
func (r *Resolver) SetABI(node [32]byte, contentType *big.Int, data []byte) (contract.Txn, error) {
	return r.c.Txn("setABI", node, contentType, data)
}

// SetAddr sends a setAddr transaction in the solidity contract
func (r *Resolver) SetAddr(node [32]byte, addr ethgo.Address) (contract.Txn, error) {
	return r.c.Txn("setAddr", node, addr)
}

// SetContent sends a setContent transaction in the solidity contract
func (r *Resolver) SetContent(node [32]byte, hash [32]byte) (contract.Txn, error) {
	return r.c.Txn("setContent", node, hash)
}

// SetName sends a setName transaction in the solidity contract
func (r *Resolver) SetName(node [32]byte, name string) (contract.Txn, error) {
	return r.c.Txn("setName", node, name)
}

// SetPubkey sends a setPubkey transaction in the solidity contract
func (r *Resolver) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (contract.Txn, error) {
	return r.c.Txn("setPubkey", node, x, y)
}

// events

func (r *Resolver) ABIChangedEventSig() ethgo.Hash {
	return r.c.GetABI().Events["ABIChanged"].ID()
}

func (r *Resolver) AddrChangedEventSig() ethgo.Hash {
	return r.c.GetABI().Events["AddrChanged"].ID()
}

func (r *Resolver) ContentChangedEventSig() ethgo.Hash {
	return r.c.GetABI().Events["ContentChanged"].ID()
}

func (r *Resolver) NameChangedEventSig() ethgo.Hash {
	return r.c.GetABI().Events["NameChanged"].ID()
}

func (r *Resolver) PubkeyChangedEventSig() ethgo.Hash {
	return r.c.GetABI().Events["PubkeyChanged"].ID()
}
