package ugc

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GTU Decode a GBK string to UTF-8
func GTU(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s),
		simplifiedchinese.GBK.NewDecoder())
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UTG Encode a UTF-8 string to GBK
func UTG(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s),
		simplifiedchinese.GBK.NewEncoder())
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}
