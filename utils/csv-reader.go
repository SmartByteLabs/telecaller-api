package utils

import "encoding/csv"

import "io"

import "net/http"

// CSVReader is handling the file read part
type CSVReader struct {
	*csv.Reader
	header    []string
	headerLen int
}

// NewCSVReaderFromRequest create new csv reader from the http request.
func NewCSVReaderFromRequest(r *http.Request) (*CSVReader, error) {
	// read file recieved
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	return NewCSVReader(file)
}

// NewCSVReader create new csv reader correponds to the file.
func NewCSVReader(f io.Reader) (*CSVReader, error) {
	reader := new(CSVReader)

	// parse file
	reader.Reader = csv.NewReader(f)
	header, err := reader.Read()
	reader.header = StringHeaderCleaner(header)
	reader.headerLen = len(header)
	return reader, err
}

// ReadLine read csv line from the current position.
func (reader CSVReader) ReadLine() (Line, error) {
	line, err := reader.Read()
	if err != nil {
		if csvErr, _ := err.(*csv.ParseError); csvErr != nil && csvErr.Err == csv.ErrFieldCount {
			err = nil
		}
	}
	return reader.parseLine(line), err
}

func (reader CSVReader) parseLine(ss []string) Line {
	line := make(Line)
	for i, s := range ss {
		if i == reader.headerLen {
			break
		}

		line[reader.header[i]] = StringCleaner(s)

	}

	return line
}
