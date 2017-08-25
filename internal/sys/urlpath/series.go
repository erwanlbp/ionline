package urlpath

// ===== The series page =====

// SeriesPath returns the path to the series listing with the possible params
func SeriesPath() string {
	return seriesBasePath()
}

// SeriesClientURL returns the URL of the series listing with the params replaced
func SeriesClientURL() string {
	return seriesBasePath()
}

func seriesBasePath() string {
	return Series
}

// ===== Add a serie=====

// AddSeriePath returns the path to add a serie with the possible params
func AddSeriePath() string {
	return addSerieBasePath()
}

// AddSerieClientURL returns the URL to add a serie with the params replaced
func AddSerieClientURL() string {
	return addSerieBasePath()
}

func addSerieBasePath() string {
	return Series
}

// ===== Delete a serie=====

// DeleteSeriePath returns the path to delete a serie with the possible params
func DeleteSeriePath() string {
	return deleteSerieBasePath(IDPathDef)
}

// DeleteSerieClientURL returns the URL to delete a serie with the params replaced
func DeleteSerieClientURL(ID string) string {
	return deleteSerieBasePath(ID)
}

func deleteSerieBasePath(ID string) string {
	return Series + "/" + ID
}
