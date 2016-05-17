package main

type FitnessTracker struct {
	avgChangeThreshold float64
	trackLength        int
	fits               []int
	lastFit            int
}

func (f *FitnessTracker) Add(fit int) {
	change := max(f.lastFit, fit) - min(f.lastFit, fit)
	f.lastFit = fit

	f.fits = append(f.fits, change)

	if excess := len(f.fits) - f.trackLength; excess > 0 {
		f.fits = f.fits[excess:]
	}
}

func (f *FitnessTracker) AvgChange() float64 {
	if len(f.fits) == 0 {
		return 0
	}

	var changeSum float64 = 0
	for _, i := range f.fits {
		changeSum += float64(i)
	}

	return changeSum / float64(len(f.fits))
}

func (f *FitnessTracker) IsImproving() bool {
	if len(f.fits) < f.trackLength {
		return true
	}

	if f.AvgChange() >= f.avgChangeThreshold {
		return true
	}

	return false
}
