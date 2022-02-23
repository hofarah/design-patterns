package main

func main() {
	airportMg := newAirportManager()
	airbus310 := &airbus{
		name:     "airbus310",
		mediator: airportMg,
	}
	airbus200 := &airbus{
		name:     "airbus200",
		mediator: airportMg,
	}
	airbus210 := &airbus{
		name:     "airbus210",
		mediator: airportMg,
	}
	airbus310.landingReq()
	airbus200.landingReq()
	airbus210.landingReq()
	airbus310.taxi()
	airbus200.taxi()
	airbus210.taxi()
}
