package ngbanks

import (
	"encoding/json"
	"fmt" // Test purposes, perhaps ?
	"io/ioutil"
	"os"
)

/*
* Bank struct consisting of the four object values in the JSON.
* Name - Banks name
* Code - Bank's reference code
* Slug - Short tag i.e abbreviation
* Ussd - Bank's quick service ussd code
 */

type Bank struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Slug string `json:"slug"`
	Ussd Ussd   `json:"ussd"`
}

/*
 * Banks struct is a master array taking in parameters of the Bank type.
 */

type Banks struct {
	Banks []Bank `json:"banks"`
}

/*
 * Ussd type, a sub type of Bank's Ussd
 */

type Ussd struct {
	Code string `json:"code"`
}

/*
 * Open JSON file, read it and converts it to bytes before it's encoded into a variable.
 */

func RetrieveBanksJSON() []byte {
	// Revise json address.
	bankFile, err := os.Open("/home/Youngestev/go/src/ngbanks/db/banks.json")

	if err != nil {
		// Basically, it should return not print.
		fmt.Println(err)
	}

	banksJSONed, _ := ioutil.ReadAll(bankFile)
	return banksJSONed
}

/*
 * ngbank's GetBanks() method
 * returns an array of banks and their properties.
 */

func GetBanks() []Bank {
	var banks Banks
	def := []Bank{}

	// The below code should be refactored. I think.
	json.Unmarshal(RetrieveBanksJSON(), &banks)

	for i := 0; i < len(banks.Banks); i++ {
		return banks.Banks
	}
	return def
}

/*
 * ngbank's GetBank(slug) method
 * returns a bank's detail that matches the slug passed as an argument to the method.
 */
func GetBank(slug string) Bank {
	var bankSLug Banks
	defBank := Bank{}

	json.Unmarshal(RetrieveBanksJSON(), &bankSLug)

	// I'm feeling like a champion - LOL !!!!

	for i, s := range bankSLug.Banks {
		if slug == bankSLug.Banks[i].Slug {
			return s
		}
	}
	return defBank
}
