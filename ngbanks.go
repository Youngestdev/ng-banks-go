package ngbanks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Bank defines it's structure.
type Bank struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Slug string `json:"slug"`
	Ussd Ussd   `json:"ussd"`
}

// Banks which is an array of bank has a data type of Bank.
type Banks struct {
	Banks []Bank `json:"banks"`
}

// Ussd struct of type string.
type Ussd struct {
	Code string `json:"code"`
}

// retrieveBanksJSON() returns json data in the fle below to byte.
func retrieveBanksJSON() []byte {
	// Revise json address.
	bankFile, err := os.Open("/home/Youngestev/go/src/ngbanks/db/banks.json")

	if err != nil {
		// Basically, it should return not print.
		fmt.Println(err)
	}

	banksJSONed, _ := ioutil.ReadAll(bankFile)
	return banksJSONed
}

// GetBanks returns an array of banks.
func GetBanks() []Bank {
	var banks Banks
	def := []Bank{}

	// The below code should be refactored. I think.
	json.Unmarshal(retrieveBanksJSON(), &banks)

	for i := 0; i < len(banks.Banks); i++ {
		return banks.Banks
	}
	return def
}

// GetBank returns a bank details depending on the slug passed in.
func GetBank(slug string) Bank {
	var bankSLug Banks
	defBank := Bank{}

	json.Unmarshal(retrieveBanksJSON(), &bankSLug)

	// I'm feeling like a champion - LOL !!!!

	for i, s := range bankSLug.Banks {
		if slug == bankSLug.Banks[i].Slug {
			return s
		}
	}
	return defBank
}
