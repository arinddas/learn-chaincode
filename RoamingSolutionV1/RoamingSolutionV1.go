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
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
)

// RoamingSolutionChaincode is simple chaincode implementing logging to Blockchain


type RoamingSolutionChaincode struct {
}


type CDR struct {

		Number string `json:"Number"`
		CallDuration string `json:"CallDuration"`
		CallCost string `json:"CallCost"`
	    DataDuration string `json:"DataDuration"`
	    DataCost string `json:"DataCost"`
		Status string `json:"Status"`
}

type Subscriber struct {

		Number string `json:"Number"`
		CallDuration string `json:"CallDuration"`
	    DataDuration string `json:"DataDuration"`
		Status string `json:"Status"`
}







// Init method will be called during deployment.

func (t *RoamingSolutionChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	fmt.Println("Init Chaincode...")
	
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	fmt.Println("Init Chaincode...done")

	return nil, nil
}



// EntitlementFromHPMN Invoke function

func (t *RoamingSolutionChaincode) EntitlementFromHPMN(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("EntitlementFromHPMN invoke Begins...")


	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	
	    var err error
		
		Status1 := "CDRApprovedByHPMN"
		key := args[0]
		
		valAsbytes, err := stub.GetState(key)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
			return nil, errors.New(jsonResp)
		} else if len(valAsbytes) == 0{
			jsonResp := "{\"Error\":\"Failed to get Query for " + key + "\"}"
			return nil, errors.New(jsonResp)
		}
		
		res := CDR{}
        json.Unmarshal(valAsbytes, &res)
		
		CallDuration := res.CallDuration
		CallCost := res.CallCost
		DataDuration := res.DataDuration
		DataCost := res.DataCost
		
            CDRobj := CDR{Number: args[0], CallDuration: CallDuration, CallCost: CallCost, DataDuration: DataDuration, DataCost: DataCost, Status: Status1}
			res2F, _ := json.Marshal(CDRobj)
		    fmt.Println(string(res2F))
		    err = stub.PutState(key,[]byte(string(res2F)))
			if err != nil {
				return nil, err
			}
			
	
		fmt.Println("EntitlementFromHPMN invoke ends...")
		return nil, nil
		
   
}

// EntitlementFromVPMN Invoke function

func (t *RoamingSolutionChaincode) EntitlementFromVPMN(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
      
        if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	    }
		
		 var CallDurationint,CallCostint,DataDurationint,DataCostint int
		
	
		Status1 := "SubscriberDetailsReceived"
		key := args[0]
		
		
		Subscriberobj := Subscriber{Number: args[0], CallDuration: args[1], DataDuration: args[2], Status: Status1}
	    res2F, _ := json.Marshal(Subscriberobj)
        fmt.Println(string(res2F))
	    err := stub.PutState(key,[]byte(string(res2F)))
			if err != nil {
				return nil, err
			}
	
		
		valAsbytes, err := stub.GetState(key)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
			return nil, errors.New(jsonResp)
		} else if len(valAsbytes) == 0{
			jsonResp := "{\"Error\":\"Failed to get Query for " + key + "\"}"
			return nil, errors.New(jsonResp)
		}
		
		res := Subscriber{}
        json.Unmarshal(valAsbytes, &res)
        
		
		
	    fmt.Println("Subscriber Details Structure",res)
	    
	   
		
		
	    CallDurationint, err = strconv.Atoi(res.CallDuration)
		if err != nil {
		return nil, err
	    }
		
		DataDurationint, err = strconv.Atoi(res.DataDuration)
		if err != nil {
		return nil, err
	    }
		
			
		// Calculate cost details for VPMN Service
		
		if CallDurationint >= 1 && CallDurationint <= 300{
		   CallCostint = CallDurationint * 2;
		}
		
		if CallDurationint > 300 && CallDurationint <= 1500{
		   CallCostint = CallDurationint * 3;
		}

        if CallDurationint > 1500 {
		   CallCostint = CallDurationint * 4;
		}
		
		
		if DataDurationint >= 1 && DataDurationint <= 300{
		   DataCostint = DataDurationint * 3;
		}
		
		if DataDurationint > 300 && DataDurationint <= 1500{
		   DataCostint = DataDurationint * 4;
		}

        if DataDurationint > 1500 {
		   DataCostint = DataDurationint * 5;
		}
		
 
         CallCost := strconv.Itoa(CallCostint)
         DataCost := strconv.Itoa(DataCostint)
         CallDuration := strconv.Itoa(CallDurationint)
         DataDuration := strconv.Itoa(DataDurationint)
		 
		 
		 // Put the state of CDR
		 
          Status1 = "CDRApprovalPending"
		
            CDRobj := CDR{Number: args[0], CallDuration: CallDuration, CallCost: CallCost, DataDuration: DataDuration, DataCost: DataCost, Status: Status1}
			res2F, _ = json.Marshal(CDRobj)
		    fmt.Println(string(res2F))
		    err = stub.PutState(key,[]byte(string(res2F)))
			if err != nil {
				return nil, err
			}
		
		
		
		
		fmt.Println("CDR Details Structure",CDRobj)
			
			
		fmt.Println("Invoke EntitlementFromHPMN Chaincode... end") 
		return nil,nil
	
	

}



// Invoke Function

func (t *RoamingSolutionChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
      
	 fmt.Println("Invoke RoamingSolution Chaincode... start") 

	
	// Handle different functions UserAcceptance
	if function == "EntitlementFromVPMN" {
		return t.EntitlementFromVPMN (stub, args)
	} else if function == "EntitlementFromHPMN" {
		return t.EntitlementFromHPMN(stub, args)
	} else{
	    return nil, errors.New("Invalid function name. Expecting 'EntitlementFromHPMN' or 'EntitlementFromVPMN' but found '" + function + "'")
	}
	
	fmt.Println("Invoke RoamingSolution Chaincode... end") 
	
	return nil,nil;
}




// Query to get CSP Service Details

func (t *RoamingSolutionChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query RoamingSolution Chaincode... start") 

	key := args[0]

    
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp := "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query NumberPoratbility Chaincode... end") 
    return valAsbytes, nil 
  
	
}



func main() {
	err := shim.Start(new(RoamingSolutionChaincode))
	if err != nil {
		fmt.Println("Error starting RoamingSolutionChaincode: %s", err)
	}
}
