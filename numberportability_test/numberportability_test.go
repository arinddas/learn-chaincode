/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)

// AssetManagementChaincode is simple chaincode implementing a basic Asset Management system
// with access control enforcement at chaincode level.
// Look here for more information on how to implement access control at chaincode level:
// https://github.com/hyperledger/fabric/blob/master/docs/tech/application-ACL.md
// An asset is simply represented by a string.

type AssetManagementChaincode struct {
}

// Init method will be called during deployment.



func (t *AssetManagementChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Init Chaincode...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	// Create ownership table
	err := stub.CreateTable("AssetsOwnership", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "mobileNumber", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "name", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "idNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		
	})
	if err != nil {
		return nil, errors.New("Failed creating AssetsOwnership table.")
	}

	fmt.Println("Init Chaincode...done")

	return nil, nil
}

func (t *AssetManagementChaincode) assign(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("In Assign...")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	mobileNumber := args[0]
	name := args[1]
	address := args[2]
	idNumber := args[3]
	

	// Register assignment
	fmt.Println("New owner of [%s] is [%s]", mobileNumber, name)

	ok, err = stub.InsertRow("AssetsOwnership", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: mobileNumber}},
			&shim.Column{Value: &shim.Column_String_{String_: name}},
			&shim.Column{Value: &shim.Column_String_{String_: address}},
			&shim.Column{Value: &shim.Column_String_{String_: idNumber}},
	})

	if !ok && err == nil {
		return nil, errors.New("MobileNumber is already assigned.")
	}

	fmt.Println("Assign...done!")

	return nil, err
}


func (t *AssetManagementChaincode) transfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("Transfer Begins...")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}


	mobileNumber := args[0]
	name := args[1]
	address := args[2]
	idNumber := args[3]
	

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: mobileNumber}}
	columns = append(columns, col1)

	row, err := stub.GetRow("AssetsOwnership", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed retrieving asset [%s]: [%s]", mobileNumber, err)
	}

	prvOwner := row.Columns[1].GetString()
	fmt.Println("Previous owener of [%s] is [%s]", mobileNumber, prvOwner)
	
	if len(prvOwner) == 0 {
		return nil, fmt.Errorf("Invalid previous owner. Nil")
	}

	err = stub.DeleteRow(
		"AssetsOwnership",
		[]shim.Column{shim.Column{Value: &shim.Column_String_{String_: mobileNumber}}},
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

	_, err = stub.InsertRow(
		"AssetsOwnership",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: mobileNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: name}},
			    &shim.Column{Value: &shim.Column_String_{String_: address}},
			    &shim.Column{Value: &shim.Column_String_{String_: idNumber}},
			},
		})
	if err != nil {
		return nil, errors.New("Failed inserting row.")
	}

	fmt.Println("New owner of [%s] is [%s]", mobileNumber, name)

	fmt.Println("Transfer...done")

	return nil, nil
}


// Invoke will be called for every transaction.
// Supported functions are the following:
// "assign(asset, owner)": to assign ownership of assets. An asset can be owned by a single entity.
// Only an administrator can call this function.
// "transfer(asset, newOwner)": to transfer the ownership of an asset. Only the owner of the specific

func (t *AssetManagementChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Handle different functions
	if function == "assign" {
		// Assign ownership
		return t.assign(stub, args)
	} else if function == "transfer" {
		// Transfer ownership
		return t.transfer(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

// Query callback representing the query of a chaincode
// Supported functions are the following:
// "query(asset)": returns the owner of the asset.
// Anyone can invoke this function.
func (t *AssetManagementChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query [%s]", function)

	if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting 'query' but found '" + function + "'")
	}

	var err error

	if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting name of an asset to query")
		return nil, errors.New("Incorrect number of arguments. Expecting mobileNumber to query")
	}

	// Who is the owner of the asset?
	mobileNumber := args[0]

	fmt.Println("Arg [%s]", string(mobileNumber))

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: mobileNumber}}
	columns = append(columns, col1)

	row, err := stub.GetRow("AssetsOwnership", columns)
	if err != nil {
		fmt.Println("Failed retriving details of [%s]: [%s]", string(mobileNumber), err)
		return nil, fmt.Errorf("Failed retriving details of [%s]: [%s]", string(mobileNumber), err)
	}
   if row == nil {
		fmt.Println("MobileNumber : [%s] not assigned to anyone ", string(mobileNumber))
		return nil, fmt.Errorf("MobileNumber : [%s] not assigned to anyone ", string(mobileNumber))
	}
	fmt.Println("Query done : Owner of the number :: [%s]", row.Columns[1].GetString()))

	return row.Columns[], nil
}

func main() {
	err := shim.Start(new(AssetManagementChaincode))
	if err != nil {
		fmt.Println("Error starting AssetManagementChaincode: %s", err)
	}
}
