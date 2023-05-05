package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(1, 2)

	if total != 3 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 3)
	}

	total = SomaX(1, 2)

	if total != 4 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 4)
	}

	total = SomaXX(1, 2)

	if total != 5 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 5)
	}

	total = SomaXXX(1, 2)

	if total != 6 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 6)
	}

	total = SomaXXXX(1, 2)

	if total != 7 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 7)
	}

	total = Sub(3, 2)

	if total != 1 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 1)
	}
}
