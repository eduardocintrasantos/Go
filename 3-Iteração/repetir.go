package repetir

func Repetir(caracter string) string {
	var repeticao string
	for i := 0; i < 5; i++ {
		repeticao += caracter
	}
	return repeticao
}
