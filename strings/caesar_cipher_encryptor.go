package strings

// Category: strings
// Difficulty: Easy

// Given a string and a key, output the given ceaser cipher for that string

func KeyFromHashMapValue(hashMap map[byte]int, value int) byte {
	var returnValue byte
	for k, v := range hashMap {
		if value == v {
			returnValue = k
			break
		}
	}
	return returnValue
}

func CaesarCipherEncryptor(str string, key int) string {

	alphabetMap := map[byte]int{
		'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6,
		'g': 7, 'h': 8, 'i': 9, 'j': 10, 'k': 11, 'l': 12,
		'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18,
		's': 19, 't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26,
	}
	outPutString := []byte(str)

	for i := 0; i < len(str); i++ {
		offsetPosition := alphabetMap[str[i]] + key
		if offsetPosition > 26 {
			newCharPosition := offsetPosition % 26
			outPutString[i] = KeyFromHashMapValue(alphabetMap, newCharPosition)
		} else {
			outPutString[i] = KeyFromHashMapValue(alphabetMap, offsetPosition)
		}
	}
	return string(outPutString)
}
