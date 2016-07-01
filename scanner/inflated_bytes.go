package scanner

import "io"

// the 2 bytes indicating the beginning of the inflated data
var magicNumbers = []byte{0x1f, 0x8b}

// InflatedBytes finds and returns inflated bytes in a stream and reads until EOF
// of a null byte is encountered.
func InflatedBytes(r io.Reader) ([]byte, error) {
	magicFound := 0
	encodedB := []byte{}

	buf := make([]byte, 1)
	var err error
	for err == nil {
		_, err = r.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		switch buf[0] {
		case 0x0:
			break
		case magicNumbers[0]:
			magicFound = 1
		case magicNumbers[1]:
			if magicFound == 1 {
				magicFound = 2
			}
		default:
			if magicFound == 2 {
				encodedB = append(encodedB, buf[0])
				continue
			}
			magicFound = 0
		}
	}

	return encodedB, nil
}
