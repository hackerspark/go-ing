package hamming

import "errors"

func Distance(a, b string) (int, error) {
    var length int
    
	if len(a) != len(b) {
        return 0, errors.New("Different lengths")
    }

    for i := range a {
        if(a[i] != b[i]) {
            length++
        }
    }

    return length, nil
}
