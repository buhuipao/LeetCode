type ParkingSystem struct {
    d map[int]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{d: map[int]int{1: big, 2: medium, 3: small}}
}


func (this *ParkingSystem) AddCar(carType int) bool {
    if this.d[carType] <= 0 {
        return false
    }
    this.d[carType] -= 1
    return true
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
