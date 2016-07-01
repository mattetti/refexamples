package scanner

import "io"

// the 2 bytes indicating the beginning of the hidden data
var magicBytes = []byte{0x4d, 0x41}

// SecretBytes finds and returns secret bytes in a stream and reads until EOF
// or until a null byte is encountered.
func SecretBytes(r io.Reader) ([]byte, error) {
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
		case magicBytes[0]:
			magicFound = 1
		case magicBytes[1]:
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
