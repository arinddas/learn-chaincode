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
	"strings"
)

// NumberPortabilityChaincode is simple chaincode implementing logging to Blockchain


type NumberPortabilityChaincode struct {
}

type EligibilityConfirm struct {

		Number string
		ServiceProviderOld string
		ServiceProviderNew string
		CustomerName string 
	    SSNNumber string
	    PortabilityIndicator string
		status string
}



type UserAcceptance struct {

		Number string
		ServiceProviderOld string
		PlanOld string 
	    ServiceValidityOld string
	    TalktimeBalanceOld string
		SMSbalanceOld string
		DataBalanceOld string
		ServiceProviderNew string
		PlanNew string 
	    ServiceValidityNew string
	    TalktimeBalanceNew string
		SMSbalanceNew string
		DataBalanceNew string
		CustomerAcceptance string
		status string
		
}


type UsageDetailsFromDonorandAcceptorCSP struct {

		Number string
		ServiceProviderOld string
		PlanOld string 
	    ServiceValidityOld string
	    TalktimeBalanceOld string
		SMSbalanceOld string
		DataBalanceOld string
		ServiceProviderNew string
		PlanNew string 
	    ServiceValidityNew string
	    TalktimeBalanceNew string
		SMSbalanceNew string
		DataBalanceNew string
		status string
		
}


type UsageDetailsFromCSP struct {

		Number string
		ServiceProviderOld string
		ServiceProviderNew string
		Plan string 
	    ServiceValidity string
	    TalktimeBalance string
		SMSbalance string
		DataBalance string
		status string
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

func (t *NumberPortabilityChaincode) EligibilityConfirm(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("EligibilityConfirm  Information invoke Begins...")
	
     //VP0
	
	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 6")
	}
	status1 := "EligibilityConfirmed"
	key := args[0]+args[1]+args[2]
	EligibilityConfirmObj := EligibilityConfirm{Number: args[0], ServiceProviderOld: args[1], ServiceProviderNew: args[2], CustomerName: args[3], SSNNumber: args[4], PortabilityIndicator: args[5], status: status1}
	err := stub.PutState(key,[]byte(fmt.Sprintf("%s",EligibilityConfirmObj)))
			if err != nil {
				return nil, err
			}
	
	
	fmt.Println("EligibilityConfirm  Information invoke ends...")
	return nil, nil 
}



func (t *NumberPortabilityChaincode) ConfirmationOfMNPRequest(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("ConfirmationOfMNPRequest  Information invoke Begins...")
	
     //VP0
	
	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 6")
	}
	
	status1 := "InitiationConfirmed"
	key := args[0]+args[1]+args[2]
	EligibilityConfirmObj := EligibilityConfirm{Number: args[0], ServiceProviderOld: args[1], ServiceProviderNew: args[2], CustomerName: args[3], SSNNumber: args[4], PortabilityIndicator: args[5], status: status1}
	err := stub.PutState(key,[]byte(fmt.Sprintf("%s",EligibilityConfirmObj)))
			if err != nil {
				return nil, err
			}
	
	fmt.Println("ConfirmationOfMNPRequest  Information invoke ends...")
	return nil, nil 
}




// UserAcceptance Invoke function

func (t *NumberPortabilityChaincode) UserAcceptance(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("UserAcceptance Information invoke Begins...")

	if len(args) != 14 {
		return nil, errors.New("Incorrect number of arguments. Expecting 14")
	}

	// Check the User Acceptance paramater, if true then update world state with new status
	
	var status1 string
	key := args[0]+args[1]+args[7]
	Acceptance := args[13]
	if(Acceptance == "true"){
	 status1 = "CustomerAccepted"
	} else {
	  status1 = "CustomerRejected"
	}
	
	UserAcceptanceObj := UserAcceptance{Number: args[0], ServiceProviderOld: args[1], PlanOld: args[2], ServiceValidityOld: args[3], TalktimeBalanceOld: args[4], SMSbalanceOld: args[5], DataBalanceOld: args[6], ServiceProviderNew: args[7], PlanNew: args[8], ServiceValidityNew: args[9], TalktimeBalanceNew: args[10], SMSbalanceNew: args[11], DataBalanceNew: args[12], CustomerAcceptance: args[13], status: status1}
    fmt.Println("UserAcceptance Details Structure ",UserAcceptanceObj)
	err := stub.PutState(key,[]byte(fmt.Sprintf("%s",UserAcceptanceObj)))
	if err != nil {
		return nil, err
	}
	
	fmt.Println("UserAcceptance Information invoke ends...")
	return nil, nil
}



// FinalPortInfo Invoke function

func (t *NumberPortabilityChaincode) UsageDetailsFromDonorCSP(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("UsageDetailsFromDonorCSP Information invoke Begins...")


	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 8")
	}
	
	    var err error
		
		status1 := "DonorApproved"
		key := args[0]+args[1]+args[2]
		
            UsageDetailsFromDonorCSPObj := UsageDetailsFromCSP{Number: args[0], ServiceProviderOld: args[1], ServiceProviderNew: args[2], Plan: args[3], ServiceValidity: args[4], TalktimeBalance: args[5], SMSbalance: args[6], DataBalance: args[7], status: status1}
			fmt.Println("Donor Service Details Structure ",UsageDetailsFromDonorCSPObj)
			err = stub.PutState(key,[]byte(fmt.Sprintf("%s",UsageDetailsFromDonorCSPObj)))
			if err != nil {
				return nil, err
			}
			
	
		fmt.Println("UsageDetailsFromDonorCSP Information invoke ends...")
		return nil, nil
		
   
}

// in args this will take three values - number and OldCSP (DonorCSP) and newCSP
func (t *NumberPortabilityChaincode) EntitlementFromRecipientCSP(stub shim.ChaincodeStubInterface, argsOld []string) ([]byte, error) {
      
        if len(argsOld) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	    }
		
		var ServiceValidity,TalktimeBalance,SMSbalance,DataBalance int
		
	
		key := argsOld[0]+argsOld[1]+argsOld[2]
		valAsbytes, err := stub.GetState(key)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
			return nil, errors.New(jsonResp)
		} else if len(valAsbytes) == 0{
			jsonResp := "{\"Error\":\"Failed to get Query for " + key + "\"}"
			return nil, errors.New(jsonResp)
		}
		
		donorDetails := fmt.Sprintf("%s", valAsbytes)
		
		donorDetails = strings.Trim(donorDetails,"{")
		donorDetails = strings.Trim(donorDetails,"}")
		donorDetails = strings.Trim(donorDetails,"[")
		donorDetails = strings.Trim(donorDetails,"]")
		
		args := strings.Split(donorDetails, " ")
		
	    fmt.Println("Donor Service Details Structure",args)
	   
	    
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
		 
        status1 := "AcceptorApproved"
		
            UsageDetailsFromAcceptorCSPObj := UsageDetailsFromCSP{Number: argsOld[0], ServiceProviderOld: argsOld[1], ServiceProviderNew: argsOld[2], Plan: Plan, ServiceValidity: ServiceValidityNew, TalktimeBalance: TalktimeBalanceNew, SMSbalance: SMSbalanceNew, DataBalance: DataBalanceNew, status: status1}
			fmt.Println("Acceptor Service Details Structure",UsageDetailsFromAcceptorCSPObj)
			err = stub.PutState(key,[]byte(fmt.Sprintf("%s",UsageDetailsFromAcceptorCSPObj)))
			if err != nil {
				return nil, err
			}
		
	
	    valAsbytesNew, errNew := stub.GetState(key)
		if errNew != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
			return nil, errors.New(jsonResp)
		} else if len(valAsbytes) == 0{
			jsonResp := "{\"Error\":\"Failed to get Query for " + key + "\"}"
			return nil, errors.New(jsonResp)
		}
		
		
		acceptorDetails := fmt.Sprintf("%s", valAsbytesNew)
		
		acceptorDetails = strings.Trim(acceptorDetails,"{")
		acceptorDetails = strings.Trim(acceptorDetails,"}")
		acceptorDetails = strings.Trim(acceptorDetails,"[")
		acceptorDetails = strings.Trim(acceptorDetails,"]")
		
		argsNew := strings.Split(acceptorDetails, " ")
		
		fmt.Println("Acceptor Service Details Structure",argsNew)
		
		
		
		
		UsageDetailsFromDonorandAcceptorCSPObj := UsageDetailsFromDonorandAcceptorCSP{Number: args[0], ServiceProviderOld: args[1], PlanOld: args[3], ServiceValidityOld: args[4], TalktimeBalanceOld: args[5], SMSbalanceOld: args[6], DataBalanceOld: args[7], ServiceProviderNew: argsNew[2], PlanNew: argsNew[3], ServiceValidityNew: argsNew[4], TalktimeBalanceNew: argsNew[5], SMSbalanceNew: argsNew[6], DataBalanceNew: argsNew[7], status: status1}
        
		fmt.Println("Donor+Acceptor Service Details Structure",UsageDetailsFromDonorandAcceptorCSPObj)
		// put the value for Regulator Query in future
		err = stub.PutState(key,[]byte(fmt.Sprintf("%s",UsageDetailsFromDonorandAcceptorCSPObj)))
			if err != nil {
				return nil, err
			}
			
			
		fmt.Println("Invoke EntitlementFromRecipientCSP Chaincode... end") 
		return nil,nil
	
	

}

// args should be Number, serviceProviderOld, serviceProviderNew

func (t *NumberPortabilityChaincode) RegulatorQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err error

    if len(args) != 3 {
        return nil, errors.New("Incorrect number of arguments. Expecting 3 arguments")
    }

    key = args[0]+args[1]+args[2]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp = "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query NumberPortability Chaincode... end") 
    return valAsbytes, nil 

}


// args should be Number, serviceProviderOld, serviceProviderNew

func (t *NumberPortabilityChaincode) EntitlementFromRecipientCSPQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err error

    if len(args) != 3 {
        return nil, errors.New("Incorrect number of arguments. Expecting 3 arguments")
    }

    key = args[0]+args[1]+args[2]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp = "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query EntitlementFromRecipientCSPQuery ... end") 
    return valAsbytes, nil 

}



// Invoke Function

func (t *NumberPortabilityChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
      
	 fmt.Println("Invoke NumberPortability Chaincode... start") 

	
	// Handle different functions UserAcceptance
	if function == "EligibilityConfirm" {
		return t.EligibilityConfirm (stub, args)
	} else if function == "UsageDetailsFromDonorCSP" {
		return t.UsageDetailsFromDonorCSP(stub, args)
	}else if function == "EntitlementFromRecipientCSP" {
		return t.EntitlementFromRecipientCSP(stub, args)
	}else if function == "UserAcceptance" {
		return t.UserAcceptance(stub, args)
	}else if function == "ConfirmationOfMNPRequest" {
		return t.ConfirmationOfMNPRequest(stub, args)
	} else{
	    return nil, errors.New("Invalid function name. Expecting 'EligibilityConfirm' or 'UsageDetailsFromDonorCSP' or 'EntitlementFromRecipientCSP' but found '" + function + "'")
	}
	
	
	fmt.Println("Invoke Numberportability Chaincode... end") 
	
	return nil,nil;
}




// Query to get CSP Service Details

func (t *NumberPortabilityChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query NumberPortability Chaincode... start") 

	
	if function == "EntitlementFromRecipientCSPQuery" {
		return t.EntitlementFromRecipientCSPQuery(stub, args)
	} 
	
	if function == "RegulatorQuery" {
		return t.RegulatorQuery(stub, args)
	} 
	
	// else We can query WorldState to fetch value
	
	var key, jsonResp string
    var err error

    if len(args) < 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }
	fmt.Println(len(args))
	if len(args) == 3 {
	   key = args[0]+args[1]+args[2]
	} else if len(args) == 2 {
	   key = args[0]+args[1]
	} else {
	   key = args[0]
	}

    
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp = "{\"Error\":\"Failed to get Query for " + key + "\"}"
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
