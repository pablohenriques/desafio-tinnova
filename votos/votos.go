package votos

import "errors"

type Votos struct {
	votosTotais  int
	votosValidos int
	votosBrancos int
	votosNulos   int
}

func NovaEleicao(tVotos, vVotos, bVotos, nVotos int) (*Votos, error) {

	if err := checkVotos(tVotos, vVotos, bVotos, nVotos); err != nil {
		return nil, err
	}

	return &Votos{
		votosTotais:  tVotos,
		votosValidos: vVotos,
		votosBrancos: bVotos,
		votosNulos:   nVotos,
	}, nil
}

func (v Votos) PorcentagemVotosValidos() float64 {
	return calcularPercentual(float64(v.votosValidos), float64(v.votosTotais))
}

func (v Votos) PorcentagemVotosBrancos() float64 {
	return calcularPercentual(float64(v.votosBrancos), float64(v.votosTotais))
}

func (v Votos) PorcentagemVotosNulos() float64 {
	return calcularPercentual(float64(v.votosNulos), float64(v.votosTotais))
}

func calcularPercentual(numeroVotos, votosTotais float64) float64 {
	return (numeroVotos / votosTotais) * 100
}

func checkVotos(tVotos, vVotos, bVotos, nVotos int) error {
	if tVotos <= 0 {
		return errors.New("total de eleitores deve ser maior que zero")
	}

	if vVotos < 0 || bVotos < 0 || nVotos < 0 {
		return errors.New("valores de votos não podem ser negativos")
	}

	if (vVotos + bVotos + nVotos) != tVotos {
		return errors.New("Total de votos não correspondem")
	}

	return nil
}
