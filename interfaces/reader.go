package interfaces

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func TestReader() {
	reader.Validate(MyReader{})
}

