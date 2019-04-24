package main

type hotel struct {
	Name    string
	Address string
	Zip     string
	Region  string
}

func main() {
	califoniaHotels := []hotel{
		hotel{
			Name:    "Hilton",
			Address: "20, Bonstreet, Papa",
			Zip:     "2050",
			Region:  "Papa",
		},
		hotel{
			Name:    "Sheraton",
			Address: "20, Bonstreet, Papa",
			Zip:     "2050",
			Region:  "Papa",
		},
	}
}
