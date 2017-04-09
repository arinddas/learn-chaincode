/*
Copyright IBM Corp. 2017 All Rights Reserved.

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
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
)

// IOTBlockchain is simple chaincode implementing logging to Blockchain


type IOTBlockchain struct {
}


type DeviceInformation struct {

		Id string `json:"Id"`
		DeviceTemperature string `json:"DeviceTemperature"`
		AtmosphericTemperature string `json:"AtmosphericTemperature"`
	    Humidity string `json:"Humidity"`
}



// Init method will be called during deployment.

func (t *IOTBlockchain) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Init Chaincode...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}
	
	fmt.Println("Init Chaincode...done")

	return nil, nil
}


// DeviceInformation function

func (t *IOTBlockchain) DeviceInformation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("DeviceInformation invoke Begins...")
	
	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}
	
	key := args[0]
	DeviceInformationObj := DeviceInformation{Id: args[0], DeviceTemperature: args[1], AtmosphericTemperature: args[2], Humidity: args[3]}
	res2F, _ := json.Marshal(DeviceInformationObj)
    fmt.Println(string(res2F))
	err := stub.PutState(key,[]byte(string(res2F)))
			if err != nil {
				return nil, err
			}
	
	
	fmt.Println("DeviceInformation invoke ends...")
	return nil, nil 
}




// Invoke Function

func (t *IOTBlockchain) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
      
	 fmt.Println("Invoke IOTBlockchain Chaincode... start") 

	
	// Handle different functions UserAcceptance
	if function == "DeviceInformation" {
		return t.DeviceInformation (stub, args)
	} else{
	    return nil, errors.New("Invalid function name. Expecting 'DeviceInformation' but found '" + function + "'")
	}
	
	fmt.Println("Invoke IOTBlockchain Chaincode... end") 
	
	return nil,nil;
}




// Query to get CSP Service Details

func (t *IOTBlockchain) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	    fmt.Println("Query IOTBlockchain Chaincode... start") 
	    key := args[0]
    
		valAsbytes, err := stub.GetState(key)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
			return nil, errors.New(jsonResp)
		} else if len(valAsbytes) == 0{
			jsonResp := "{\"Error\":\"Failed to get Query for " + key + "\"}"
			return nil, errors.New(jsonResp)
		}

	fmt.Println("Query IOTBlockchain Chaincode... end") 
    return valAsbytes, nil 
  
	
}



func main() {
	err := shim.Start(new(IOTBlockchain))
	if err != nil {
		fmt.Println("Error starting IOTBlockchain: %s", err)
	}
}
