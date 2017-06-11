/*
Author : Arindam Dasgupta
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
)

// RetailTradingChaincode is simple chaincode implementing logging to Blockchain


type RetailTradingChaincode struct {
}


type Assign struct {

		CustId string `json:"CustId"`
		Retailer string `json:"Retailer"`
		LoyaltyPoints string `json:"LoyaltyPoints"`
}


type MutualTrading struct {

		Customer1 string `json:"Customer1"`
		Retailer1 string `json:"Retailer1"`
		LoyaltyPoints1 string `json:"LoyaltyPoints1"`
		Customer2 string `json:"Customer2"`
		Retailer2 string `json:"Retailer2"`
		LoyaltyPoints2 string `json:"LoyaltyPoints2"`
		Status string `json:"Status"`
		TimeStamp string `json:"TimeStamp"`
}

type TRADEALL struct {

		TRADEDetail []MutualTrading `json:"TRADEDetail"`
}


// Init method will be called during deployment.

func (t *RetailTradingChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Init Chaincode...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}
	
	// Create RetailLoyaltyPointDetails table
	err := stub.CreateTable("RetailLoyaltyPointDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "CustId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Retailer", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "LoyaltyPoints", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating RetailLoyaltyPointDetails table.")
	}
	
	
	
	// Create table RetailTradingDetails
	
	// cust1+cust2+Timestamp is the key
	
	err1 := stub.CreateTable("RetailTradingDetails", []*shim.ColumnDefinition{
	    &shim.ColumnDefinition{Name: "Key", Type: shim.ColumnDefinition_STRING, Key: true}, 
		&shim.ColumnDefinition{Name: "Customer1", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Retailer1", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "LoyaltyPoints1", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Customer2", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Retailer2", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "LoyaltyPoints2", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "TimeStamp", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err1 != nil {
		return nil, errors.New("Failed creating RetailTradingDetails table.")
	}
	

	fmt.Println("Init Chaincode...done")

	return nil, nil
}


// assign function

func (t *RetailTradingChaincode) Assign(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("Assign invoke Begins...")
	
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}
	
	ok, err := stub.InsertRow("RetailLoyaltyPointDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: args[0]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
	      },
		  })
	  if err != nil {
			return nil, fmt.Errorf("insert Record operation failed. %s", err)
		}
		if !ok {
			return nil, errors.New("Customer Id is already assigned.")
		}
	
	fmt.Println("Assign invoke ends...")
	return nil, nil 
}


// update Function
// 3 arguments Id, LoyaltyPoints to Add or Substract, ADD or SUB indicator

func (t *RetailTradingChaincode) Update(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("Update invoke Begins...")
	
	
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}
	
	
	
    key := args[0]
	
    var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: key}}
	columns = append(columns, col1)

	row, err := stub.GetRow("RetailLoyaltyPointDetails", columns)
	if err != nil {
		fmt.Println("Failed retriving details of %s: %s", string(key), err)
		return nil, fmt.Errorf("Failed retriving details of %s: %s", string(key), err)
	}
	
    if len(row.Columns) != 0{
		
		// Customer1 := args[0]
		Retailer1 := row.Columns[1].GetString_()
		LoyaltyPoints1 := row.Columns[2].GetString_()
		
		
		LoyaltyPoints1_int, err2 := strconv.Atoi(LoyaltyPoints1)
		if err2 != nil {
		return nil, err2
	    }
		
		if args[2] == "ADD" {
		   
		   LoyaltyPointsADD_int, err := strconv.Atoi(args[1])
		   if err != nil {
					return nil, err
					}
		   
		   LoyaltyPoints1_int = LoyaltyPoints1_int + LoyaltyPointsADD_int
		   
		} else if args[2] == "SUB" {
		    LoyaltyPointsSUB_int, err := strconv.Atoi(args[1])
			if err != nil {
					return nil, err
					}
		    LoyaltyPoints1_int = LoyaltyPoints1_int - LoyaltyPointsSUB_int
			
			if LoyaltyPoints1_int <= 0 {
			    return nil, fmt.Errorf("Can not decrese Loyalty Points beyond 0")
			}
		
		} else { 
     		return nil, fmt.Errorf("Not a Valid Indicator to update")
		}
		
		LoyaltyPoints1 = strconv.Itoa(LoyaltyPoints1_int)
		
		ok, err := stub.ReplaceRow("RetailLoyaltyPointDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: args[0]}},
			&shim.Column{Value: &shim.Column_String_{String_: Retailer1}},
			&shim.Column{Value: &shim.Column_String_{String_: LoyaltyPoints1}},
	      },
		  })
	  if err != nil {
			return nil, fmt.Errorf("Update Record operation failed. %s", err)
		}
		if !ok {
			return nil, errors.New("Update Record operation failed")
		}
		
		}else{
		
	    fmt.Println("Customer Id : %s not assigned to anyone ", string(key))
		fmt.Println("Update RetailTrading Chaincode... end") 
		return nil, fmt.Errorf("Customer Id : %s not assigned to anyone ", string(key))
	}
	
	
	fmt.Println("Update invoke ends...")
	return nil, nil 
}


// LifeStyle to ShoppersStop Query

// 3 Arguments custId1 and custId2 and LoyaltyPoints to be trade by Customer1

func (t *RetailTradingChaincode) LSToSSQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key string
    var err error
	
    if len(args) != 3 {
        return nil, errors.New("Incorrect number of arguments. Expecting 3 arguments")
    }
	
	// custId1 and custId2 and LoyaltyPoints to be trade by Customer1

    key = args[0]
	
    var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: key}}
	columns = append(columns, col1)

	row, err := stub.GetRow("RetailLoyaltyPointDetails", columns)
	if err != nil {
		fmt.Println("Failed retriving details of %s: %s", string(key), err)
		return nil, fmt.Errorf("Failed retriving details of %s: %s", string(key), err)
	}
	
    if len(row.Columns) != 0{
		
		// Customer1 := args[0]
		Retailer1 := row.Columns[1].GetString_()
		LoyaltyPoints1 := row.Columns[2].GetString_()
		
		var Customer2, Retailer2, LoyaltyPoints2 string
		
		
		key = args[1]
	
		var columns2 []shim.Column
		col2 := shim.Column{Value: &shim.Column_String_{String_: key}}
		columns2 = append(columns2, col2)
		
		row1, err1 := stub.GetRow("RetailLoyaltyPointDetails", columns)
		if err1 != nil {
			fmt.Println("Failed retriving details of %s: %s", string(key), err1)
			return nil, fmt.Errorf("Failed retriving details of %s: %s", string(key), err1)
		}
		
		
		if len(row1.Columns) != 0{     
		    Customer2 = args[1]
			Retailer2 = row1.Columns[1].GetString_()
			LoyaltyPoints2 = row1.Columns[2].GetString_()
			
		}else{
	    fmt.Println("Customer Id : %s not assigned to anyone ", string(key))
		fmt.Println("Query RetailTrading Chaincode... end") 
		return nil, fmt.Errorf("Customer Id : %s not assigned to anyone ", string(key))
	    }
		
		
		LoyaltyPoints1_int, err := strconv.Atoi(LoyaltyPoints1)
		if err != nil {
		return nil, err
	    }
		
		LoyaltyPoints2_int, err := strconv.Atoi(LoyaltyPoints2)
		if err != nil {
		return nil, err
	    }
		
		
		CustomerTradingPoints_int, err := strconv.Atoi(args[2])
		if err != nil {
		return nil, err
	    }
		
		// Main Logic where 
		// Shopperstop Equivalent point amount is less for the same value of LifeStyle
		//Eg. If Lifestyle offers 100 then ShoppersStop equivalent is 90
		
		
		CustomerTradingPoints_int2 := CustomerTradingPoints_int - CustomerTradingPoints_int/10
		
		if (LoyaltyPoints1_int >= CustomerTradingPoints_int) &&  (LoyaltyPoints2_int > 0) && (LoyaltyPoints1_int > 0) && (LoyaltyPoints2_int >= CustomerTradingPoints_int2)  {
		
		    CustomerTradingPoints_str2 := strconv.Itoa(CustomerTradingPoints_int2)
		    str := `{"Customer1": "` + args[0] + `", "Retailer1": "` + Retailer1 + `", "TradePointsCust1": "` + args[2] + `", "Customer2": "` + Customer2 + `", "Retailer2": "` + Retailer2 + `", "TradePointsCust2": "` + CustomerTradingPoints_str2 + `"}`
			rowString := fmt.Sprintf("%s", str)
			fmt.Println("Query Done : Details :: %s", str)
			return []byte(rowString), nil		
		
		}else{
			fmt.Println("Trade not Possible")
			fmt.Println("Query RetailTrading Chaincode... end") 
			return nil, fmt.Errorf("Trade not Possible. Requested or Designated Customer has less or no Loyalty Points")
		}
	    
		
		
		
		}else{
		
	    fmt.Println("Customer Id : %s not assigned to anyone ", string(key))
		fmt.Println("Query RetailTrading Chaincode... end") 
		return nil, fmt.Errorf("Customer Id : %s not assigned to anyone ", string(key))
	}

}



func (t *RetailTradingChaincode) SSToLSQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key string
    var err error
	
	var Customer2, Retailer2, LoyaltyPoints2 string
	
    if len(args) != 3 {
        return nil, errors.New("Incorrect number of arguments. Expecting 3 arguments")
    }
	
	// custId1 and custId2 and LoyaltyPoints to be trade by Customer1

    key = args[0]
	
    var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: key}}
	columns = append(columns, col1)

	row, err := stub.GetRow("RetailLoyaltyPointDetails", columns)
	if err != nil {
		fmt.Println("Failed retriving details of %s: %s", string(key), err)
		return nil, fmt.Errorf("Failed retriving details of %s: %s", string(key), err)
	}
	
    if len(row.Columns) != 0{
		
		// Customer1 := args[0]
		Retailer1 := row.Columns[1].GetString_()
		LoyaltyPoints1 := row.Columns[2].GetString_()
		
		
		key = args[1]
	
		var columns2 []shim.Column
		col2 := shim.Column{Value: &shim.Column_String_{String_: key}}
		columns2 = append(columns2, col2)
		
		row1, err1 := stub.GetRow("RetailLoyaltyPointDetails", columns)
		if err1 != nil {
			fmt.Println("Failed retriving details of %s: %s", string(key), err1)
			return nil, fmt.Errorf("Failed retriving details of %s: %s", string(key), err1)
		}
		
		if len(row1.Columns) != 0{     
		    Customer2 = args[1]
			Retailer2 = row1.Columns[1].GetString_()
			LoyaltyPoints2 = row1.Columns[2].GetString_()
			
		}else{
	    fmt.Println("Customer Id : %s not assigned to anyone ", string(key))
		fmt.Println("Query RetailTrading Chaincode... end") 
		return nil, fmt.Errorf("Customer Id : %s not assigned to anyone ", string(key))
	    }
		
		
		LoyaltyPoints1_int, err := strconv.Atoi(LoyaltyPoints1)
		if err != nil {
		return nil, err
	    }
		
		LoyaltyPoints2_int, err := strconv.Atoi(LoyaltyPoints2)
		if err != nil {
		return nil, err
	    }
		
		CustomerTradingPoints_int, err := strconv.Atoi(args[2])
		if err != nil {
		return nil, err
	    }
		
		// Main Logic where 
		// Lifestyle Equivalent point amount is more for the same value of Shoppersstop
		//Eg. If ShopersStop offers 100 then LifeStyle equivalent is 110
		
		
		CustomerTradingPoints_int2 := CustomerTradingPoints_int + CustomerTradingPoints_int/10
		
		if (LoyaltyPoints1_int >= CustomerTradingPoints_int) &&  (LoyaltyPoints2_int > 0) && (LoyaltyPoints1_int > 0) && (LoyaltyPoints2_int >= CustomerTradingPoints_int2)  {
		
		    CustomerTradingPoints_str2 := strconv.Itoa(CustomerTradingPoints_int2)
		    str := `{"Customer1": "` + args[0] + `", "Retailer1": "` + Retailer1 + `", "TradePointsCust1": "` + args[2] + `", "Customer2": "` + Customer2 + `", "Retailer2": "` + Retailer2 + `", "TradePointsCust2": "` + CustomerTradingPoints_str2 + `"}`
			rowString := fmt.Sprintf("%s", str)
			fmt.Println("Query Done : Details :: %s", str)
			return []byte(rowString), nil		
		
		}else{
			fmt.Println("Trade not Possible")
			fmt.Println("Query RetailTrading Chaincode... end") 
			return nil, fmt.Errorf("Trade not Possible. Requested or Designated Customer has less or no Loyalty Points")
		}
	    
		
		
		
		}else{
		
	    fmt.Println("Customer Id : %s not assigned to anyone ", string(key))
		fmt.Println("Query RetailTrading Chaincode... end") 
		return nil, fmt.Errorf("Customer Id : %s not assigned to anyone ", string(key))
	}

}



// RetailTradingStatus invoke function
// 8 arguments like Structure of MutualTrading

func (t *RetailTradingChaincode) RetailTradingStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("RetailTradingStatus Information invoke Begins...")
	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 8")
	}
	
	
	// Update World State and the Table
	
	key := args[0]+args[3]+args[7]
	MutualTradingObj := MutualTrading{Customer1: args[0], Retailer1: args[1], LoyaltyPoints1: args[2], Customer2: args[3], Retailer2: args[4], LoyaltyPoints2: args[5], Status: args[6], TimeStamp: args[7]}
	res2F, _ := json.Marshal(MutualTradingObj)
    fmt.Println(string(res2F))
	err := stub.PutState(key,[]byte(string(res2F)))
			if err != nil {
				return nil, err
			}
			
	
    ok, errNew := stub.InsertRow("RetailTradingDetails", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: key}},
			&shim.Column{Value: &shim.Column_String_{String_: args[0]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[6]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[7]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[8]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[9]}},
	      },
		  })
	  if errNew != nil {
			return nil, fmt.Errorf("insert Record operation failed. %s", errNew)
		}
		if !ok {
			return nil, errors.New("insert Record operation failed for Trading.")
		}
		
	// If Trade is Accepted by Designated Customer then we have to update the loyalty points table	
		Status := args[6] 
	 if Status == "Trade_Accepted" {
	
		       
			   // Fetch Current Loyalty point details
			   
	            key = args[0]
				
				var LoyaltyPoints1_int int
				var LoyaltyPoints2_int int
				
				var columns []shim.Column
				col1 := shim.Column{Value: &shim.Column_String_{String_: key}}
				columns = append(columns, col1)

				row, err := stub.GetRow("RetailLoyaltyPointDetails", columns)
				if err != nil {
					fmt.Println("Failed retriving details of %s: %s", string(key), err)
					return nil, fmt.Errorf("Failed retriving details of %s: %s", string(key), err)
				}
				
				if len(row.Columns) != 0{
					
					// Customer1 := args[0]
					// Retailer1 := row.Columns[1].GetString_()
					LoyaltyPoints1 := row.Columns[2].GetString_()
					
					var LoyaltyPoints2 string
					
					
					key = args[3]
				
					var columns2 []shim.Column
					col2 := shim.Column{Value: &shim.Column_String_{String_: key}}
					columns2 = append(columns2, col2)
					
					row1, err1 := stub.GetRow("RetailLoyaltyPointDetails", columns)
					if err1 != nil {
						fmt.Println("Failed retriving details of %s: %s", string(key), err1)
						return nil, fmt.Errorf("Failed retriving details of %s: %s", string(key), err1)
					}
					
					if len(row1.Columns) != 0{     
						// Customer2 = args[1]
						// Retailer2 = row1.Columns[1].GetString_()
						LoyaltyPoints2 = row1.Columns[2].GetString_()
						
					}else{
					fmt.Println("Customer Id : %s not assigned to anyone ", string(key))
					fmt.Println("Query RetailTrading Chaincode... end") 
					return nil, fmt.Errorf("Customer Id : %s not assigned to anyone ", string(key))
					}
					
					
					LoyaltyPoints1_int, err = strconv.Atoi(LoyaltyPoints1)
					if err != nil {
					return nil, err
					}
					
					LoyaltyPoints2_int, err = strconv.Atoi(LoyaltyPoints2)
					if err != nil {
					return nil, err
					}
					
					
					}else{
					
					fmt.Println("Customer Id : %s not assigned to anyone ", string(key))
					fmt.Println("Query RetailTrading Chaincode... end") 
					return nil, fmt.Errorf("Customer Id : %s not assigned to anyone ", string(key))
				}
				
				// Adjust the points according to the trading 
			
				
             	Customer1Points, err := strconv.Atoi(args[2])
				if err != nil {
					return nil, err
					}
			    Customer2Points, err := strconv.Atoi(args[5])
				if err != nil {
					return nil, err
					}
				
				
				LoyaltyPoints1_int = LoyaltyPoints1_int - Customer1Points + Customer2Points
				LoyaltyPoints2_int = LoyaltyPoints2_int - Customer1Points + Customer1Points
				
				LoyaltyPoints1_str := strconv.Itoa(LoyaltyPoints1_int)
				LoyaltyPoints2_str := strconv.Itoa(LoyaltyPoints2_int)
				
				
				
				
			  // Update Table with New Values	
					
			   ok, errNew := stub.ReplaceRow("RetailLoyaltyPointDetails", shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: args[0]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
					&shim.Column{Value: &shim.Column_String_{String_: LoyaltyPoints1_str}},
				  },
				  })
			  if errNew != nil {
					return nil, fmt.Errorf("Update Record operation failed. %s", errNew)
				}
				if !ok {
					return nil, errors.New("Update Record operation failed for Trading.")
				}
				
				
				
				 ok1, errNew1 := stub.ReplaceRow("RetailLoyaltyPointDetails", shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
					&shim.Column{Value: &shim.Column_String_{String_: LoyaltyPoints2_str}},
				  },
				  })
			  if errNew1 != nil {
					return nil, fmt.Errorf("Update Record operation failed. %s", errNew)
				}
				if !ok1 {
					return nil, errors.New("Update Record operation failed for Trading.")
				}
				
				
				
				
				
				
				
	}	
		
	
	fmt.Println("RetailTradingStatus Information invoke ends...")
	return nil, nil 
}





// Customer Loyalty Points Query function

func (t *RetailTradingChaincode) LoyaltyPointsQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {


    fmt.Println("RetailLoyaltyPointsQuery Query Begins...")
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	key:= args[0];
	 
	 fmt.Println("key in query",key)
	 
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: key}}
	columns = append(columns, col1)

	row, err := stub.GetRow("RetailLoyaltyPointDetails", columns)
	if err != nil {
		fmt.Println("Failed retriving details of %s: %s", string(args[0]), err)
		return nil, fmt.Errorf("Failed retriving details of %s: %s", string(args[0]), err)
	}
	
	
    if len(row.Columns) != 0{
		
			CustId := row.Columns[0].GetString_()
			Retailer := row.Columns[1].GetString_()
			LoyaltyPoints := row.Columns[2].GetString_()
		
            Assignobj := Assign{CustId: CustId, Retailer: Retailer, LoyaltyPoints: LoyaltyPoints}
			res2F, _ := json.Marshal(Assignobj)
		    fmt.Println(string(res2F))
	
			fmt.Println("RetailLoyaltyPointsQuery Query ends...")
			return res2F, nil
		
   
     }
	 
	 return nil, fmt.Errorf("Failed retriving details of %s: %s", string(args[0]), err)

}




func (t *RetailTradingChaincode) GetALLTradingDetailsQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {


    fmt.Println("GetALLTradingDetailsQuery Query Begins...")
	
	var columns []shim.Column
	

	rows, err := stub.GetRows("RetailTradingDetails", columns)
	
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data \"}"
		return nil, errors.New(jsonResp)
	}
	


	 var tradeAll TRADEALL
	 
	 tradeAll.TRADEDetail = make([]MutualTrading, 0)
	 
	 for row := range rows {		
		
			Customer1 := row.Columns[1].GetString_()
			Retailer1 := row.Columns[2].GetString_()
			LoyaltyPoints1 := row.Columns[3].GetString_()
			Customer2 := row.Columns[4].GetString_()
			Retailer2 := row.Columns[5].GetString_()
			LoyaltyPoints2 := row.Columns[6].GetString_()
			Status := row.Columns[7].GetString_()
			TimeStamp := row.Columns[7].GetString_()
			
		
		
	        MutualTradingObj := MutualTrading{Customer1: Customer1, Retailer1: Retailer1, LoyaltyPoints1: LoyaltyPoints1, Customer2: Customer2, Retailer2: Retailer2, LoyaltyPoints2: LoyaltyPoints2, Status: Status, TimeStamp: TimeStamp}
			
			tradeAll.TRADEDetail = append(tradeAll.TRADEDetail, MutualTradingObj)
		
	}
		mapB, _ := json.Marshal(tradeAll)
        fmt.Println(string(mapB))
	
	    return mapB, nil

}






// Invoke Function

func (t *RetailTradingChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
      
	 fmt.Println("Invoke RetailTrading Chaincode... start") 

	if function == "Assign" {
		return t.Assign (stub, args)
	} else if function == "Update" {
		return t.Update(stub, args)
	} else if function == "RetailTradingStatus" {
		return t.RetailTradingStatus(stub, args)
	} else{
	    return nil, errors.New("Invalid function name. Expecting 'Assign' or 'Update' or 'RetailTradingStatus' but found '" + function + "'")
	}
	
	fmt.Println("Invoke RetailTrading Chaincode... end") 
	
	return nil,nil;
}




// Query to get CSP Service Details

func (t *RetailTradingChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query RetailTrading Chaincode... start") 

	
	if function == "LSToSSQuery" {
		return t.LSToSSQuery(stub, args)
	} 
	
	if function == "SSToLSQuery" {
		return t.SSToLSQuery(stub, args)
	} 
	
	if function == "LoyaltyPointsQuery" {
		return t.LoyaltyPointsQuery(stub, args)
	} 
	
	if function == "GetALLTradingDetailsQuery" {
		return t.GetALLTradingDetailsQuery(stub, args)
	} else{
	    return nil, errors.New("Invalid function name. Expecting 'LSToSSQuery' or 'SSToLSQuery' or 'LoyaltyPointsQuery' or 'GetALLTradingDetailsQuery' but found '" + function + "'")
	}
	
	
	fmt.Println("Query NumberPoratbility Chaincode... end") 
    return nil, nil 
  
	
}



func main() {
	err := shim.Start(new(RetailTradingChaincode))
	if err != nil {
		fmt.Println("Error starting RetailTradingChaincode: %s", err)
	}
}
