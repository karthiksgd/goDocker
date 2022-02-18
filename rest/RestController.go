package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CustData struct {
	CustomerId      int64 `json: "customerId"`
	OptIn           bool  `json: "optIn"`
	InactivityTimer int   `json: "inactivityTimer"`
}

type SkyData struct {
	Response bool   `json:"response"`
	Message  string `json:"message"`
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	customers := GetCustomerData()

	fmt.Println(customers)

	json.NewEncoder(w).Encode(customers)

}

func CustomerData(c CustData) SkyData {

	defaultTimer := 2

	// var custData CustData

	var responseData SkyData
	// json.NewDecoder(r.Body).Decode(&custData)

	data, valid := sanitizeData(c)

	fmt.Println("Sanitized Data : ", data, " ", valid)

	if valid == false {

		responseData = data

		// json.NewEncoder(w).Encode(sanitData)

	} else {

		// json.NewEncoder(w).Encode(custData)

		if c.OptIn == false {

			responseData = SkyData{false, "No Permission"}

			// json.NewEncoder(w).Encode(responseData)

		} else if c.InactivityTimer >= (defaultTimer * 60 * 60) {

			responseData = SkyData{false, "Session Timeout"}

			// json.NewEncoder(w).Encode(responseData)

		} else if c.CustomerId%7 == 0 {

			fmt.Println("Customer Data divisible by 7: ", c)

			responseData = SkyData{false, "Linear Advert is targeted"}

			// json.NewEncoder(w).Encode(responseData)

		} else if c.InactivityTimer%9 == 0 {

			responseData = SkyData{false, "System is Offline"}

			// json.NewEncoder(w).Encode(responseData)

		} else {

			responseData = SkyData{true, "All Good"}

			// json.NewEncoder(w).Encode(responseData)
		}

	}
	return responseData

}

func sanitizeData(data CustData) (SkyData, bool) {

	var digits int
	var returnData SkyData
	var isValidated bool

	fmt.Println("Test Data", data)

	for data.CustomerId != 0 {
		data.CustomerId /= 10
		digits++
	}

	if digits != 12 {

		returnData = SkyData{false, "Customer Id Err"}
		isValidated = false
		fmt.Println(digits)

	} else if data.OptIn == false {

		returnData = SkyData{false, "Customer hasn't Opted In"}
		isValidated = false

	} else {

		returnData = SkyData{}
		isValidated = true
	}

	return returnData, isValidated

}

func CheckCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	var customer CustData

	json.NewDecoder(r.Body).Decode(&customer)

	responseData := CustomerData(customer)

	json.NewEncoder(w).Encode(responseData)

	// fmt.Println(responseData)

}
