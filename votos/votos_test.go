package votos

import (
	"testing"
)

func TestVotosValidos(t *testing.T) {
	eleicao, err := NovaEleicao(1000, 800, 150, 50)
	if err != nil {
		t.Fatalf("Erro inesperado ao criar eleição: %v", err)
	}

	actual := eleicao.PorcentagemVotosValidos()
	expected := 80.0
	if expected != actual {
		t.Errorf("Expected %f is different of actual %f", expected, actual)
	}
}

func TestVotosBrancos(t *testing.T) {
	eleicao, _ := NovaEleicao(1000, 800, 150, 50)
	actual := eleicao.PorcentagemVotosBrancos()
	expected := 15.0
	if expected != actual {
		t.Errorf("Expected %f is different of actual %f", expected, actual)
	}
}

func TestVotosNulos(t *testing.T) {
	eleicao, _ := NovaEleicao(1000, 800, 150, 50)
	actual := eleicao.PorcentagemVotosNulos()
	expected := 5.0
	if expected != actual {
		t.Errorf("Expected %f is different of actual %f", expected, actual)
	}

}

func TestErroVotos(t *testing.T) {
	var _, err = NovaEleicao(1000, 90, 150, 50)
	var expected = "Total de votos não correspondem"
	if err.Error() != expected {
		t.Errorf("Erro Inesperado: %v", err)
	}
}

func TestErroZeroVotos(t *testing.T) {
	var _, err = NovaEleicao(0, 0, 0, 0)
	var expected = "total de eleitores deve ser maior que zero"
	if err.Error() != expected {
		t.Errorf("Erro Inesperado: %v", err)
	}
}

func TestErroVotosNegativos(t *testing.T) {
	var _, err = NovaEleicao(100, -10, 60, 50)
	var expected = "valores de votos não podem ser negativos"
	if err.Error() != expected {
		t.Errorf("Erro Inesperado: %v", err)
	}
}
