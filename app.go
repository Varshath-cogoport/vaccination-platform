package main

type User struct {
	Id       int
	Name     string
	Uid      int
	State    string
	District string
}

type Center struct {
	Id                 int
	CenterId           int
	State              string
	District           string
	DayCapacityMapping map[int]int
}

type Booking struct {
	Id       int
	CenterId int
	UserId   int
	day      int
	isActive bool
}

type Application struct {
	users    []User
	centers  []Center
	bookings []Booking
}

var app Application

func (a *Application) AddUser(u *User) {
	a.users = append(a.users, *u)
}

func (a *Application) AddCenter(c *Center) {
	c.DayCapacityMapping = make(map[int]int)
	a.centers = append(a.centers, *c)
}

func (a *Application) AddCenterCapacity(centerId, day, capacity int) {
	for i, _ := range a.centers {
		if a.centers[i].CenterId == centerId {
			a.centers[i].DayCapacityMapping[day] = capacity
		}
	}
}

func (a *Application) BookVaccination(b *Booking) {
	b.isActive = true
	if a.GetSlotsAvailable(b.CenterId, b.day) > 0 {
		a.bookings = append(a.bookings, *b)
	}
}

func (a *Application) CancelBooking(b *Booking) {
	for i, _ := range app.bookings {
		if app.bookings[i].Id == b.Id {
			app.bookings[i].isActive = false
		}
	}
}

func (a *Application) GetSlotsAvailable(centerId, day int) int {
	var bookings int
	for _, booking := range app.bookings {
		if booking.CenterId == centerId && booking.day == day && booking.isActive {
			bookings++
		}
	}

	for _, center := range app.centers {
		if center.CenterId == centerId {
			return center.DayCapacityMapping[day] - bookings
		}
	}

	return 0
}

func (a *Application) GetCenters() []Center {
	return app.centers
}

func (a *Application) GetBookings() []Booking {
	return app.bookings
}

func (a *Application) GetCentersForDistrict(districtName string) []Center {
	var centers []Center
	for _, center := range app.centers {
		if center.District == districtName {
			centers = append(centers, center)
		}
	}
	return centers
}
