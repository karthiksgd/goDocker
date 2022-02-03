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

func RestCalls(w http.ResponseWriter, r *http.Request) {

	defaultTimer := 2

	w.Header().Set("content-type", "application/json")

	var custData CustData
	json.NewDecoder(r.Body).Decode(&custData)

	data, valid := sanitizeData(custData)

	fmt.Println("Sanitized Data : ", data, " ", valid)

	if valid == false {

		sanitData := data
		json.NewEncoder(w).Encode(sanitData)

	} else {

		// json.NewEncoder(w).Encode(custData)

		fmt.Println("Customer Data before loop : ", custData)

		if custData.OptIn == false {

			responseData := SkyData{false, "No Permission"}

			json.NewEncoder(w).Encode(responseData)

		} else if custData.InactivityTimer >= (defaultTimer * 60 * 60) {

			responseData := SkyData{false, "Session Timeout"}

			json.NewEncoder(w).Encode(responseData)

		} else if custData.CustomerId%7 == 0 {

			fmt.Println("Customer Data divisible by 7: ", custData)

			responseData := SkyData{false, "Linear Advert is targeted"}

			json.NewEncoder(w).Encode(responseData)

		} else if custData.InactivityTimer%9 == 0 {

			responseData := SkyData{false, "System is Offline"}

			json.NewEncoder(w).Encode(responseData)

		} else {

			responseData := SkyData{true, "All Good"}

			json.NewEncoder(w).Encode(responseData)
		}

	}

}

func sanitizeData(data CustData) (SkyData, bool) {

	var digits int
	var returnData SkyData
	var isValidated bool

	for data.CustomerId != 0 {
		data.CustomerId /= 10
		digits++
	}

	if digits != 12 {

		returnData = SkyData{false, "Customer Id Err"}
		isValidated = false
		fmt.Println(digits)

	} else {

		returnData = SkyData{}
		isValidated = true
	}

	return returnData, isValidated

}
