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
)

// NumberPortabilityChaincode is simple chaincode implementing logging to Blockchain


type NumberPortabilityChaincode struct {
}

type CGTADetails struct {

		Number string
		ServiceProvider string
		CustomerName string 
	    SSNNumber string
	    PortabilityIndicator string
}


type FinalPortInfo struct {

		Number string
		ServiceProvider string
		Plan string 
	    ServiceValidity string
	    TalktimeBalance string
		SMSbalance string
		DataBalance string
		CustomerAcceptance string
		
}


type CSPServiceDetails struct {

		Number string
		ServiceProviderOld string
		ServiceProviderNew string
		Plan string 
	    ServiceValidity string
	    TalktimeBalance string
		SMSbalance string
		DataBalance string
}

// Init method will be called during deployment.

func (t *NumberPortabilityChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Init Chaincode...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	fmt.Println("Init Chaincode...done")

	return nil, nil
}




// CGTA invoke function

func (t *NumberPortabilityChaincode) CGTAInformation(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("CGTA Information invoke Begins...")

	if len(args) != 5 {
		return nil, errors.New("Incorrect number of arguments. Expecting 5")
	}
	
	err := stub.PutState(args[0],[]byte(args[1]))
	if err != nil {
		return nil, err
	}
	fmt.Println("CGTA Information invoke ends...")
	return nil, nil 
}



// FinalPortInfo Invoke function

func (t *NumberPortabilityChaincode) FinalPortInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("FinalPortInfo Information invoke Begins...")

	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 8")
	}

	// Check the Customer Acceptance paramater, if true then update world state with new CSP value
	
	Acceptance := args[7]
	if(Acceptance == "true"){
	err := stub.PutState(args[0],[]byte(args[1]))
	if err != nil {
		return nil, err
	}
	}
	
	fmt.Println("FinalPortInfo Information invoke ends...")
	return nil, nil
}



// FinalPortInfo Invoke function

func (t *NumberPortabilityChaincode) CSPServiceDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("CSPServiceDetails Information invoke Begins...")


	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 8")
	}
	
	    var err error
		var ServiceValidity,TalktimeBalance,SMSbalance,DataBalance int
	    
	    Number := args[0]
		ServiceProviderOld := args[1]
		ServiceProviderNew := args[2]
		Plan := args[3]
		
	    ServiceValidity, err = strconv.Atoi(args[4])
		if err != nil {
		return nil, err
	    }
		
	    TalktimeBalance, err = strconv.Atoi(args[5])
		if err != nil {
		return nil, err
	    }
		
		SMSbalance, err = strconv.Atoi(args[6])
		if err != nil {
		return nil, err
	    }
		
		DataBalance, err = strconv.Atoi(args[7])
		if err != nil {
		return nil, err
	    }
		
		
		
		key := Number+ServiceProviderOld
		
            CSPServiceDetailsStructObjDonor := CSPServiceDetails{Number: args[0], ServiceProviderOld: args[1], ServiceProviderNew: args[2],  Plan: args[3], ServiceValidity: args[4], TalktimeBalance: args[5], SMSbalance: args[6], DataBalance: args[7]}
			fmt.Println("Donor Service Deatils Structure %s",CSPServiceDetailsStructObjDonor)
			err = stub.PutState(key,[]byte(fmt.Sprintf("%s",CSPServiceDetailsStructObjDonor)))
			if err != nil {
				return nil, err
			}
			
		// Calculate Acceptor Service
		
		if Plan == "PlanA"{
		    Plan = "PlanC"
			ServiceValidity = ServiceValidity - (ServiceValidity/5)
			TalktimeBalance = TalktimeBalance - (TalktimeBalance/5)
			SMSbalance = SMSbalance - (SMSbalance/5)
			DataBalance = DataBalance - (DataBalance/5)
		}	

       if Plan == "PlanB"{
		    Plan = "PlanA"
			ServiceValidity = ServiceValidity - (ServiceValidity/6)
			TalktimeBalance = TalktimeBalance - (TalktimeBalance/6)
			SMSbalance = SMSbalance - (SMSbalance/6)
			DataBalance = DataBalance - (DataBalance/6)
		}
		
       if Plan == "PlanC"{
		    Plan = "PlanB"
			ServiceValidity = ServiceValidity - (ServiceValidity/4)
			TalktimeBalance = TalktimeBalance - (TalktimeBalance/4)
			SMSbalance = SMSbalance - (SMSbalance/4)
			DataBalance = DataBalance - (DataBalance/4)
		}

         ServiceValidityNew := strconv.Itoa(ServiceValidity)
         TalktimeBalanceNew := strconv.Itoa(TalktimeBalance)
         SMSbalanceNew := strconv.Itoa(SMSbalance)
         DataBalanceNew := strconv.Itoa(DataBalance)
		 
		 // Put the state of Acceptor
		 
		 
       key = Number+ServiceProviderNew
		
            CSPServiceDetailsStructObjAcceptor := CSPServiceDetails{Number: args[0], ServiceProviderOld: args[1], ServiceProviderNew: args[2], Plan: args[3], ServiceValidity: ServiceValidityNew, TalktimeBalance: TalktimeBalanceNew, SMSbalance: SMSbalanceNew, DataBalance: DataBalanceNew}
            fmt.Println("Acceptor Service Deatils Structure %s",CSPServiceDetailsStructObjAcceptor)
			err = stub.PutState(key,[]byte(fmt.Sprintf("%s",CSPServiceDetailsStructObjAcceptor)))
			if err != nil {
				return nil, err
			}
	
	fmt.Println("CSPServiceDetails Information invoke ends...")
	return nil, nil
	
   
}






// Invoke Function

func (t *NumberPortabilityChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
      
	 fmt.Println("Invoke NumberPortability Chaincode... start") 

	
	// Handle different functions
	if function == "CGTAInformation" {
		return t.CGTAInformation(stub, args)
	} else if function == "CSPServiceDetails" {
		return t.CSPServiceDetails(stub, args)
	} else if function == "FinalPortInfo" {
		return t.FinalPortInfo(stub, args)
	}else{
	    return nil, errors.New("Invalid function name. Expecting 'CGTAInformation' or 'CSPServiceDetails' but found '" + function + "'")
	}
	
	
	fmt.Println("Invoke Numberportability Chaincode... end") 
	
	return nil,nil;
}




// Query to get CSP Service Details

func (t *NumberPortabilityChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query NumberPortability Chaincode... start") 

	if function != "QueryCSPDetails" {
		return nil, errors.New("Invalid query function name. Expecting 'query' but found '" + function + "'")
	}
	
	var key, jsonResp string
    var err error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }

    key = args[0]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }

	fmt.Println("Query NumberPoratbility Chaincode... end") 
    return valAsbytes, nil
	
    
  
	
}



func main() {
	err := shim.Start(new(NumberPortabilityChaincode))
	if err != nil {
		fmt.Println("Error starting NumberPortabilityChaincode: %s", err)
	}
}
