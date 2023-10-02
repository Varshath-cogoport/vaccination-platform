package main

import "fmt"

func main() {
	//onboard user
	u1 := User{
		Id:       1,
		Name:     "harry",
		Uid:      99,
		State:    "Maharashtra",
		District: "Mumbai",
	}
	app.AddUser(&u1)

	//onboard center
	c1 := Center{
		Id:       1,
		CenterId: 22,
		State:    "maharashtra",
		District: "Mumbai",
	}
	app.AddCenter(&c1)

	//onboard center
	c2 := Center{
		Id:       1,
		CenterId: 23,
		State:    "Karnataka",
		District: "Banglore",
	}
	app.AddCenter(&c2)

	//add capacity per center
	app.AddCenterCapacity(22, 1, 1)

	//book a vaccination
	b1 := Booking{
		Id:       1,
		CenterId: 22,
		UserId:   1,
		day:      1,
	}
	app.BookVaccination(&b1)

	//cancel a booking
	app.CancelBooking(&b1)

	//getBookings
	fmt.Println(app.bookings)

	//getCenters
	fmt.Println(app.GetCenters())

	//getCentersByName
	fmt.Println(app.GetCentersForDistrict("Mumbai"))
}
