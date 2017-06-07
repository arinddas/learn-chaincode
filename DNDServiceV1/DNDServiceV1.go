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
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
)

// DNDServiceChaincode is simple chaincode implementing logging to Blockchain


type DNDServiceChaincode struct {
}


type DNDinfo struct {

		Subscriber string `json:"Subscriber"`
		DNDActivator string `json:"DNDActivator"`
		StatusMessage string `json:"StatusMessage"`
		DNDStatus string `json:"DNDStatus"`
}




// Init method will be called during deployment.

func (t *DNDServiceChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Init Chaincode...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}
	
	fmt.Println("Init Chaincode...done")

	return nil, nil
}



//





// CGTA invoke function

func (t *DNDServiceChaincode) DNDServiceActivation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("DNDServiceActivation  Information invoke Begins...")
	
     //VP0
	
	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}
	key := args[0]+args[1]
	DNDinfoObj := DNDinfo{Subscriber: args[0], DNDActivator: args[1], StatusMessage: args[2], DNDStatus: args[3]}
	res2F, _ := json.Marshal(DNDinfoObj)
    fmt.Println(string(res2F))
	err := stub.PutState(key,[]byte(string(res2F)))
			if err != nil {
				return nil, err
			}
	
	
	fmt.Println("DNDServiceActivation  Information invoke ends...")
	return nil, nil 
}


func (t *DNDServiceChaincode) DNDServiceDeactivation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("DNDServiceDeactivation  Information invoke Begins...")
	
     //VP0
	
	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}
	key := args[0]+args[1]
	DNDinfoObj := DNDinfo{Subscriber: args[0], DNDActivator: args[1], StatusMessage: args[2], DNDStatus: args[3]}
	res2F, _ := json.Marshal(DNDinfoObj)
    fmt.Println(string(res2F))
	err := stub.PutState(key,[]byte(string(res2F)))
			if err != nil {
				return nil, err
			}
	
	
	fmt.Println("DNDServiceDeactivation  Information invoke ends...")
	return nil, nil 
}





// args should be Number, serviceProviderOld, serviceProviderNew

func (t *DNDServiceChaincode) DNDQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err error

    if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2 arguments")
    }

    key = args[0]+args[1]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp = "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query DNDService Chaincode... end") 
    return valAsbytes, nil 

}



// Invoke Function

func (t *DNDServiceChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
      
	 fmt.Println("Invoke DNDService Chaincode... start") 

	
	// Handle different functions UserAcceptance
	if function == "DNDServiceActivation" {
		return t.DNDServiceActivation (stub, args)
	} else if function == "DNDServiceDeactivation" {
		return t.DNDServiceDeactivation(stub, args)
	} else{
	    return nil, errors.New("Invalid function name. Expecting 'DNDServiceActivation' or 'DNDServiceDeactivation' but found '" + function + "'")
	}
	
	fmt.Println("Invoke DNDService Chaincode... end") 
	
	return nil,nil;
}




// Query to get CSP Service Details

func (t *DNDServiceChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query DNDService Chaincode... start") 

	
	if function == "DNDQuery" {
		return t.DNDQuery(stub, args)
	} 

    if len(args) < 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }
	
	key := args[0]+args[1]
	

    
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp := "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query DNDService Chaincode... end") 
    return valAsbytes, nil 
  
	
}



func main() {
	err := shim.Start(new(DNDServiceChaincode))
	if err != nil {
		fmt.Println("Error starting DNDServiceChaincode: %s", err)
	}
}
