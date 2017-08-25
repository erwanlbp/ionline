package types

// Quality represents a quality
type Quality int

// The differents qualities
const (
	QualityHDTV = Quality(iota)
	Quality720
	Quality1080
)

// Strings of the qualities
const (
	QualityHDTVString = "HDTV"
	Quality720String  = "720p"
	Quality1080String = "1080p"
)

// QualityName maps Quality -> String
var QualityName = map[Quality]string{
	QualityHDTV: QualityHDTVString,
	Quality720:  Quality720String,
	Quality1080: Quality1080String,
}

// QualityValue maps String -> Quality
var QualityValue = map[string]Quality{
	QualityHDTVString: QualityHDTV,
	Quality720String:  Quality720,
	Quality1080String: Quality1080,
}
