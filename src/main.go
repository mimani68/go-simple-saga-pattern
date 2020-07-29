package main

import (
	"fmt"
)

func main() {
	fmt.Println("salam")

	// A ) campaign object
	// B ) config object
	// C ) attachment

	UpdateCampaign := func () (b bool, e error)  {
		// updating that failed
		return false, nil
	}

	CompensateCampaign := func (v int) error {
		fmt.Printf("\ntransaction goes back to %v\n" ,v)
		return nil
	}

	baseValue := 100  // get from server CampaignRepository.getSingleById( id )
	// newValue := 120
	type successState struct {
		Campaign bool
	}
	var s successState
	var err error
	s.Campaign, err= UpdateCampaign()
	if s.Campaign == false {
		CompensateCampaign(baseValue)
		fmt.Printf("\nerror is %v\n",err)
	} else {
		fmt.Println("\nupdate successful\n")
	}
}
