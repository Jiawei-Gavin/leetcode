package leetcode

type ParkingSystem struct {
	parking []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		parking: []int{big, medium, small},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if (carType == 1 && this.parking[0] > 0) || (carType == 2 && this.parking[1] > 0) || (carType == 3 && this.parking[2] > 0) {
		this.parking[carType-1]--
		return true
	}
	return false
}
