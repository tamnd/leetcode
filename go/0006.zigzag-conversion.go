import "bytes"

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	ret := bytes.Buffer{}
	p := numRows*2 - 2

	for i := 0; i < len(s); i += p {
		ret.WriteByte(s[i])
	}

	for i := 1; i <= numRows-2; i++ {
		ret.WriteByte(s[i])

		for j := p; j-i < len(s); j += p {
			ret.WriteByte(s[j-i])
			if j+i < len(s) {
				ret.WriteByte(s[j+i])
			}
		}
	}

	for i := numRows - 1; i < len(s); i += p {
		ret.WriteByte(s[i])
	}

	return ret.String()
}
