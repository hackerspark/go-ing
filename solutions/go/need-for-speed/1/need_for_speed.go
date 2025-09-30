package speed

// TODO: define the 'Car' type struct
type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}

type Track struct {
    distance int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{
        batteryDrain: batteryDrain,
        speed: speed,
        battery: 100,
        distance: 0,
    }
}

// TODO: define the 'Track' type struct

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    if car.battery >= car.batteryDrain {
        car.distance += car.speed
        car.battery -= car.batteryDrain
    }
    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	maxDistance := (car.battery / car.batteryDrain) * car.speed
    return maxDistance >= track.distance
}
