package types

import "time"

type PerformanceNature string

const (
	PerformanceNaturePreview  PerformanceNature = "PREVIEW"
	PerformanceNaturePremiere PerformanceNature = "PREMIERE"
)

type Performance struct {
	ID int64
	// Play is the name of the performance: "The CICD - Corneille", "Les fourberies de Scala - Molière"
	Play              string
	StartTime         time.Time
	EndTime           time.Time
	PerformanceNature PerformanceNature
}
