// Copyright 2018 The dexon-consensus Authors
// This file is part of the dexon-consensus library.
//
// The dexon-consensus library is free software: you can redistribute it
// and/or modify it under the terms of the GNU Lesser General Public License as
// published by the Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.
//
// The dexon-consensus library is distributed in the hope that it will be
// useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser
// General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the dexon-consensus library. If not, see
// <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/dexon-foundation/dexon/accounts/abi"
)

var betContract = string("0x60806040523480156200001157600080fd5b5060405160208062000ee083398101806040528101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36200010b8162000112640100000000026401000000009004565b50620003f5565b60006200011e62000364565b6000620001396200029f640100000000026401000000009004565b15156200014557600080fd5b606484101515620001e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001807f4578706563746174696f6e2073686f756c64206265206c657373207468616e2081526020017f3130302e0000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b836001819055506200021061271085620002f664010000000002620008c4179091906401000000009004565b925060008260006064811015156200022457fe5b602002018181525050600190505b606481101562000285576200025f8184620003386401000000000262000902179091906401000000009004565b82826064811015156200026e57fe5b602002018181525050808060010191505062000232565b8160029060646200029892919062000388565b5050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b60008060008414156200030d576000915062000331565b82840290508284828115156200031f57fe5b041415156200032d57600080fd5b8091505b5092915050565b6000806000831115156200034b57600080fd5b82848115156200035757fe5b0490508091505092915050565b610c8060405190810160405280606490602082028038833980820191505090505090565b8260648101928215620003ba579160200282015b82811115620003b95782518255916020019190600101906200039c565b5b509050620003c99190620003cd565b5090565b620003f291905b80821115620003ee576000816000905550600101620003d4565b5090565b90565b610adb80620004056000396000f3006080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630c60e0c3146100a9578063379607f5146100d6578063715018a6146101035780637365870b1461011a5780638da5cb5b1461013a5780638f32d59b14610191578063e1152343146101c0578063ed88c68e14610201578063f2fde38b1461020b578063fc1c39dc1461024e575b600080fd5b3480156100b557600080fd5b506100d460048036038101908080359060200190929190505050610279565b005b3480156100e257600080fd5b50610101600480360381019080803590602001909291905050506103cb565b005b34801561010f57600080fd5b506101186104be565b005b61013860048036038101908080359060200190929190505050610590565b005b34801561014657600080fd5b5061014f610803565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561019d57600080fd5b506101a661082c565b604051808215151515815260200191505060405180910390f35b3480156101cc57600080fd5b506101eb60048036038101908080359060200190929190505050610883565b6040518082815260200191505060405180910390f35b61020961089d565b005b34801561021757600080fd5b5061024c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061089f565b005b34801561025a57600080fd5b506102636108be565b6040518082815260200191505060405180910390f35b6000610283610a26565b600061028d61082c565b151561029857600080fd5b606484101515610336576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001807f4578706563746174696f6e2073686f756c64206265206c657373207468616e2081526020017f3130302e0000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b83600181905550610352612710856108c490919063ffffffff16565b9250600082600060648110151561036557fe5b602002018181525050600190505b60648110156103b35761038f818461090290919063ffffffff16565b828260648110151561039d57fe5b6020020181815250508080600101915050610373565b8160029060646103c4929190610a4a565b5050505050565b6103d361082c565b15156103de57600080fd5b3073ffffffffffffffffffffffffffffffffffffffff1631811115151561046d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4e6f20656e6f756768206f6620746f6b656e20746f20636c61696d2e0000000081525060200191505060405180910390fd5b610475610803565b73ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156104ba573d6000803e3d6000fd5b5050565b6104c661082c565b15156104d157600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060018311151561060b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f5461726765742073686f756c64206265206269676765722e000000000000000081525060200191505060405180910390fd5b6001548311151515610685576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f5461726765742073686f756c6420626520736d616c6c65722e0000000000000081525060200191505060405180910390fd5b612710341115156106fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4d696e696d756d206265742069732031303030302e000000000000000000000081525060200191505060405180910390fd5b600160642f81151561070c57fe5b0601915060009050828210156107a05761075661271061074860026001870360648110151561073757fe5b0154346108c490919063ffffffff16565b61090290919063ffffffff16565b90503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561079e573d6000803e3d6000fd5b505b3373ffffffffffffffffffffffffffffffffffffffff167f97371a3349bea11f577edf6e64350a3dfb9de665d1154c7e6d08eb0805aa043084348460405180848152602001838152602001828152602001935050505060405180910390a2505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b60028160648110151561089257fe5b016000915090505481565b565b6108a761082c565b15156108b257600080fd5b6108bb8161092c565b50565b60015481565b60008060008414156108d957600091506108fb565b82840290508284828115156108ea57fe5b041415156108f757600080fd5b8091505b5092915050565b60008060008311151561091457600080fd5b828481151561091f57fe5b0490508091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561096857600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b610c8060405190810160405280606490602082028038833980820191505090505090565b8260648101928215610a79579160200282015b82811115610a78578251825591602001919060010190610a5d565b5b509050610a869190610a8a565b5090565b610aac91905b80821115610aa8576000816000905550600101610a90565b5090565b905600a165627a7a7230582087d19571ab3ae7207fb8f6182b6b891f2414f00e1904790341a82c957044b5330029")
var betConstructor = []string{"62"}
var betABIJSON = `
[
    {
      "constant": false,
      "inputs": [],
      "name": "renounceOwnership",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "owner",
      "outputs": [
        {
          "name": "",
          "type": "address"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "isOwner",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "name": "payout",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "transferOwnership",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "expectation",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "name": "_expectation",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "_addr",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "_target",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "_value",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "_pay",
          "type": "uint256"
        }
      ],
      "name": "Bet",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "previousOwner",
          "type": "address"
        },
        {
          "indexed": true,
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "OwnershipTransferred",
      "type": "event"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_expectation",
          "type": "uint256"
        }
      ],
      "name": "updateExpectation",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_target",
          "type": "uint256"
        }
      ],
      "name": "bet",
      "outputs": [],
      "payable": true,
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_amount",
          "type": "uint256"
        }
      ],
      "name": "claim",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [],
      "name": "donate",
      "outputs": [],
      "payable": true,
      "stateMutability": "payable",
      "type": "function"
    }
  ]
`
var betABI abi.ABI

var oneDEX *big.Int

func init() {
	oneDEX = new(big.Int).Exp(big.NewInt(10), big.NewInt(19), nil)
	var err error
	betABI, err = abi.JSON(strings.NewReader(betABIJSON))
	if err != nil {
		panic(err)
	}
}

func (m *Monkey) Gamble() {
	fmt.Println("Deploying contract ...")
	contract := m.deploy(m.source, betContract, betConstructor, new(big.Int), math.MaxUint64)
	fmt.Println("  Contract deployed: ", contract.String())
	fmt.Println("Donating ...")
	input, err := betABI.Pack("donate")
	if err != nil {
		panic(err)
	}
	m.call(m.source, contract, input, new(big.Int).Set(oneDEX), 0, math.MaxUint64)

	time.Sleep(5 * time.Second)

	input, err = betABI.Pack("bet", big.NewInt(50))
	if err != nil {
		panic(err)
	}

	nonce := uint64(0)
	for {
		fmt.Println("nonce", nonce)
		for _, key := range m.keys {
			m.call(key, contract, input, big.NewInt(100000), uint64(32740), nonce)
		}
		nonce += 1
		time.Sleep(500 * time.Millisecond)
	}
}
