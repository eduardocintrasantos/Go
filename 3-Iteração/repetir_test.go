package repetir

import "testing"

func TestRepetir(t *testing.T) {
	resultado := Repetir("a")
	esperado := "aaaaa"

	if resultado != esperado {
		t.Errorf("esperado '%s', resultado '%s'", resultado, esperado)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a")
	}
}
