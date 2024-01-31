func mySqrt(x int) int {
    z := 1.0
	previousZ := 0.0

    // newtons method solution
	for i := 0; i < 100; i++ { 
		z -= (z*z - float64(x)) / (2 * z)
		// consider if change in values is small enough
        if math.Abs(z-previousZ) < 0.05 {
			break
		}

		// save the current value of z to be the previous for the next iteration
		previousZ = z
	}

	return int(z)
}