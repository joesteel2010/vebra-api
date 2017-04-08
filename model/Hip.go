package model

type Hip struct {
	EnergyEfficiency    EnergyRatingPair `xml:"energyefficiency"`
	EnvironmentalImpact EnergyRatingPair `xml:"environmentalimpact"`
}
