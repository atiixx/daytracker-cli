package types

type Question struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Answers      []string `json:"answers"`
	DefaultValue string   `json:"default_value"`
	CSVName      string   `json:"csv"`
}

type ConfigData struct {
	CSVFilepath string     `json:"csv_filepath"`
	CSVFilename string     `json:"csv_filename"`
	Questions   []Question `json:"questions"`
}
