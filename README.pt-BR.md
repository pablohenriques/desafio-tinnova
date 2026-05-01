<div align="right">
  <strong>🇧🇷 Português</strong> &nbsp;|&nbsp; <a href="README.md">🇺🇸 English</a>
</div>

# Desafios de Programação

> Série de exercícios progressivos cobrindo fundamentos de programação, algoritmos e desenvolvimento full-stack.

---

## Índice

- [Sobre](#sobre)
- [Persona Avaliadora — Ada](#persona-avaliadora--ada)
- [Exercícios](#exercícios)
  - [01 · Votos em Relação ao Total de Eleitores](#01--votos-em-relação-ao-total-de-eleitores)
  - [02 · Bubble Sort](#02--bubble-sort)
  - [03 · Fatorial](#03--fatorial)
  - [04 · Soma dos Múltiplos de 3 ou 5](#04--soma-dos-múltiplos-de-3-ou-5)
  - [05 · Cadastro de Veículos](#05--cadastro-de-veículos)
- [Tabela de Checkpoints](#tabela-de-checkpoints)
- [Critérios Gerais de Avaliação](#critérios-gerais-de-avaliação)

---

## Sobre

Este repositório reúne cinco desafios de programação que cobrem desde conceitos fundamentais de orientação a objetos e algoritmos clássicos até o desenvolvimento de uma aplicação web full-stack com API RESTful. Os exercícios são independentes entre si e implementados em **Go (Golang)**.

**Estrutura sugerida de diretórios:**

```
desafio-tinnova/
├── ex01-votos/
├── ex02-bubble-sort/
├── ex03-fatorial/
├── ex04-multiplos/
└── ex05-veiculos/
```

---

## Persona Avaliadora — Ada

**Ada** é uma engenheira de software sênior com 15 anos de experiência em back-end, arquitetura de sistemas e code review. Seu estilo é direto, técnico e construtivo — ela não apenas aponta erros, mas explica o *porquê* por trás de cada correção e garante que o conceito foi absorvido antes de seguir em frente.

Ada atua em três frentes:

| Papel | O que Ada faz |
|---|---|
| **Ensino** | Apresenta o conceito com exemplos práticos e analogias antes de cobrar a implementação. Explica a teoria por trás de cada exercício. |
| **Avaliação** | Revisa o código submetido contra os critérios da tabela de checkpoints. Sinaliza: `APROVADO`, `REVISAR` ou `REPROVADO` por item. |
| **Monitoria** | Após a entrega, Ada faz perguntas direcionadas para confirmar que o conceito foi internalizado, não apenas copiado. Exemplo: *"explique por que o Bubble Sort tem complexidade O(n²)"*. |

> **Princípio da Ada:** *"Código que funciona sem entendimento é dívida técnica disfarçada."*

---

## Exercícios

---

### 01 · Votos em Relação ao Total de Eleitores

#### Descrição

Dada a seguinte base de dados eleitoral:

```
total de eleitores = 1000
válidos            = 800
votos brancos      = 150
nulos              = 50
```

Implemente uma **classe** com três métodos que calculem:

- O percentual de votos **válidos** em relação ao total de eleitores
- O percentual de votos **brancos** em relação ao total de eleitores
- O percentual de votos **nulos** em relação ao total de eleitores

> **Dica:** "em relação ao total" significa que o divisor é sempre o total de eleitores.
> Exemplo: `percentualNulos = nulos / totalEleitores × 100`

Utilize **programação orientada a objetos**.

#### Conceitos Necessários

| # | Conceito | Por que é necessário |
|---|---|---|
| 1 | **Classes e instâncias** | O exercício exige encapsular os dados e métodos em uma classe |
| 2 | **Atributos de instância** | Armazenar `totalEleitores`, `validos`, `brancos` e `nulos` como estado do objeto |
| 3 | **Métodos de instância** | Cada cálculo de percentual deve ser um método separado |
| 4 | **Encapsulamento** | Proteger os dados e expor apenas a interface necessária |
| 5 | **Aritmética de ponto flutuante** | Divisão entre inteiros pode resultar em truncamento — entender `float` vs `int` |
| 6 | **Regra de três / percentagem** | Base matemática do cálculo: `(parte / total) × 100` |
| 7 | **Testes unitários** | Verificar os três métodos com os valores fornecidos |

---

### 02 · Bubble Sort

#### Descrição

Dado o vetor:

```
v = {5, 3, 2, 4, 7, 1, 0, 6}
```

Implemente o algoritmo **Bubble Sort** para ordenar o vetor em ordem crescente.

**Como o Bubble Sort funciona:**

1. Percorra o vetor comparando elementos adjacentes (dois a dois)
2. Troque as posições se o elemento da esquerda for maior que o da direita
3. Repita os passos acima `(n - 1)` vezes, onde `n` é o tamanho do vetor
4. A cada iteração, o maior elemento restante "sobe" para sua posição final — ele não precisa ser percorrido novamente

**Exemplo da primeira iteração:**

```
(5  3) 2  4  7  1  0  6   → troca  → 3  5  2  4  7  1  0  6
 3 (5  2) 4  7  1  0  6   → troca  → 3  2  5  4  7  1  0  6
 3  2 (5  4) 7  1  0  6   → troca  → 3  2  4  5  7  1  0  6
 3  2  4 (5  7) 1  0  6   → mantém → 3  2  4  5  7  1  0  6
 3  2  4  5 (7  1) 0  6   → troca  → 3  2  4  5  1  7  0  6
 3  2  4  5  1 (7  0) 6   → troca  → 3  2  4  5  1  0  7  6
 3  2  4  5  1  0 (7  6)  → troca  → 3  2  4  5  1  0  6 [7]
```

O `7` está em sua posição final. A segunda iteração não precisa incluí-lo.

#### Conceitos Necessários

| # | Conceito | Por que é necessário |
|---|---|---|
| 1 | **Arrays / Vetores** | Estrutura de dados base do exercício |
| 2 | **Laços de repetição aninhados** | Loop externo (n-1 iterações) + loop interno (comparações) |
| 3 | **Operação de swap (troca)** | Trocar dois elementos requer variável auxiliar ou destructuring |
| 4 | **Comparação entre elementos** | Condicionais para decidir se troca ou mantém |
| 5 | **Complexidade de algoritmos** | Entender O(n²) no pior caso e por que o algoritmo precisa de múltiplas passagens |
| 6 | **Otimização incremental** | Reduzir o tamanho da janela de comparação a cada iteração |
| 7 | **Índices e limites de array** | Evitar `IndexOutOfBounds` ao comparar pares adjacentes |

---

### 03 · Fatorial

#### Descrição

Implemente um programa que calcule o **fatorial** de qualquer número inteiro não-negativo.

**Definição formal:**

```
fatorial(n) = 1                      se n = 0
fatorial(n) = n × fatorial(n - 1)   se n > 0
```

**Casos de teste esperados:**

```
0! = 1
1! = 1
2! = 2
3! = 6
4! = 24
5! = 120
6! = 720
```

> **Atenção:** `0! = 1` porque o produto vazio (sem nenhum fator) é definido como 1.

#### Conceitos Necessários

| # | Conceito | Por que é necessário |
|---|---|---|
| 1 | **Recursão** | A definição matemática é naturalmente recursiva |
| 2 | **Caso base** | `fatorial(0) = 1` encerra a recursão; sem ele, há stack overflow |
| 3 | **Caso recursivo** | `n × fatorial(n - 1)` é a chamada que reduz o problema |
| 4 | **Call stack** | Entender como cada chamada empilha e desempilha na memória |
| 5 | **Funções puras** | O resultado depende apenas da entrada — sem efeitos colaterais |
| 6 | **Validação de entrada** | Tratar valores negativos e não-inteiros |
| 7 | **Testes unitários** | Cobrir todos os casos fornecidos, incluindo o edge case `0!` |

---

### 04 · Soma dos Múltiplos de 3 ou 5

#### Descrição

Implemente uma função que receba um número **X** como parâmetro e retorne a **soma de todos os números naturais abaixo de X que são múltiplos de 3 ou de 5**.

**Exemplo:**

Para `X = 10`, os múltiplos de 3 ou 5 abaixo de 10 são: `3, 5, 6, 9`

```
3 + 5 + 6 + 9 = 23
```

> **Atenção:** "abaixo de X" significa estritamente menor que X — o próprio X não é incluído.

#### Conceitos Necessários

| # | Conceito | Por que é necessário |
|---|---|---|
| 1 | **Operador módulo (`%`)** | Verificar divisibilidade: `n % 3 === 0` ou `n % 5 === 0` |
| 2 | **Laço de repetição** | Iterar de 1 até X-1 |
| 3 | **Condicional composta (OR)** | Um número é incluído se for múltiplo de 3 **ou** de 5 |
| 4 | **Acumulador** | Variável que soma os valores elegíveis |
| 5 | **Parâmetros de função** | A função deve ser genérica, não hardcoded para X=10 |
| 6 | **Edge cases** | X = 0, X = 1, X negativo — o que retornar? |
| 7 | **Funções puras** | Mesmo valor de entrada sempre gera o mesmo resultado |

---

### 05 · Cadastro de Veículos

#### Descrição

Desenvolva uma aplicação **full-stack** composta por:

- **Back-end:** API JSON RESTful escrita em **Go (Golang)** com todos os verbos HTTP
- **Front-end:** Single Page Application (SPA) consumindo a API

**Modelo de dados:**

```
veiculo:    string
marca:      string
ano:        integer
descricao:  text
vendido:    boolean
created:    datetime
updated:    datetime
```

**Funcionalidades obrigatórias:**

- Cadastrar, atualizar (total e parcial) e excluir veículos
- Exibir contagem de veículos **não vendidos**
- Distribuição de veículos por **década de fabricação** (ex: 1990 → 15 veículos)
- Distribuição de veículos por **fabricante** (ex: Ford → 14 veículos)
- Listar veículos registrados na **última semana**
- Validar consistência das marcas (rejeitar variações incorretas como *Forde*, *Volksvagen*)

**Endpoints da API:**

| Método | Rota | Descrição |
|---|---|---|
| `GET` | `/veiculos` | Retorna todos os veículos |
| `GET` | `/veiculos?marca={marca}&ano={ano}&cor={cor}` | Filtra por parâmetros |
| `GET` | `/veiculos/{id}` | Retorna detalhes de um veículo |
| `POST` | `/veiculos` | Cadastra um novo veículo |
| `PUT` | `/veiculos/{id}` | Atualiza todos os dados do veículo |
| `PATCH` | `/veiculos/{id}` | Atualiza parcialmente o veículo |
| `DELETE` | `/veiculos/{id}` | Remove o veículo |

#### Conceitos Necessários

| # | Conceito | Por que é necessário |
|---|---|---|
| 1 | **Arquitetura REST** | Todos os endpoints seguem as convenções RESTful |
| 2 | **Verbos HTTP** | GET, POST, PUT, PATCH e DELETE com semântica correta |
| 3 | **JSON** | Formato de troca de dados entre front e back |
| 4 | **CRUD** | As quatro operações básicas sobre o recurso veículo |
| 5 | **Banco de dados** | Persistência do modelo de dados (relacional ou NoSQL) |
| 6 | **ORM / Query Builder** | Mapeamento entre objetos e tabelas |
| 7 | **Validação de dados** | Marcas inválidas devem ser rejeitadas na camada de entrada |
| 8 | **Filtros e query params** | Endpoint de listagem com filtros opcionais |
| 9 | **Aggregations / GROUP BY** | Distribuição por década e por fabricante |
| 10 | **Filtro por data** | Veículos registrados na última semana (`created >= now - 7 days`) |
| 11 | **SPA (front-end)** | Comunicação assíncrona com a API, sem recarregamento de página |
| 12 | **Testes unitários** | Cobertura dos endpoints e regras de negócio |
| 13 | **Design Patterns** | Ex: Repository, Service Layer, MVC |
| 14 | **Clean Code** | Nomes expressivos, funções pequenas, sem duplicação |

---

## Tabela de Checkpoints

Use esta tabela para acompanhar o progresso em cada exercício. Ada avalia cada item como `✅ Aprovado`, `⚠️ Revisar` ou `❌ Reprovado`.

### Checkpoints por Exercício

| # | Checkpoint | Ex 01 | Ex 02 | Ex 03 | Ex 04 | Ex 05 |
|---|---|:---:|:---:|:---:|:---:|:---:|
| 1 | O código compila/executa sem erros | ☐ | ☐ | ☐ | ☐ | ☐ |
| 2 | O resultado está matematicamente correto | ☐ | ☐ | ☐ | ☐ | ☐ |
| 3 | Os casos de teste fornecidos passam | ☐ | ☐ | ☐ | ☐ | ☐ |
| 4 | Edge cases são tratados corretamente | ☐ | ☐ | ☐ | ☐ | ☐ |
| 5 | O código usa os conceitos exigidos pelo exercício | ☐ | ☐ | ☐ | ☐ | ☐ |
| 6 | Os nomes de variáveis e funções são expressivos | ☐ | ☐ | ☐ | ☐ | ☐ |
| 7 | Não há código duplicado desnecessário | ☐ | ☐ | ☐ | ☐ | ☐ |
| 8 | Há testes unitários com cobertura adequada | ☐ | ☐ | ☐ | ☐ | ☐ |
| 9 | O candidato consegue explicar a solução oralmente | ☐ | ☐ | ☐ | ☐ | ☐ |

### Checkpoints Específicos — Ex 01 (Votos)

| # | Critério | Status |
|---|---|:---:|
| 1.1 | Existe uma classe com atributos para os dados eleitorais | ☐ |
| 1.2 | Os três métodos de percentual estão implementados separadamente | ☐ |
| 1.3 | A divisão usa ponto flutuante (não trunca o resultado) | ☐ |
| 1.4 | Os resultados para os valores de exemplo são 80%, 15% e 5% | ☐ |

### Checkpoints Específicos — Ex 02 (Bubble Sort)

| # | Critério | Status |
|---|---|:---:|
| 2.1 | O algoritmo usa dois laços aninhados (ou equivalente) | ☐ |
| 2.2 | O swap é feito corretamente sem perder valores | ☐ |
| 2.3 | A janela de comparação reduz a cada iteração | ☐ |
| 2.4 | O vetor `{5,3,2,4,7,1,0,6}` resulta em `{0,1,2,3,4,5,6,7}` | ☐ |
| 2.5 | O candidato explica por que são necessárias `n-1` passagens | ☐ |

### Checkpoints Específicos — Ex 03 (Fatorial)

| # | Critério | Status |
|---|---|:---:|
| 3.1 | A solução usa recursão | ☐ |
| 3.2 | O caso base `fatorial(0) = 1` está implementado | ☐ |
| 3.3 | `5! = 120` e `6! = 720` são retornados corretamente | ☐ |
| 3.4 | Entradas inválidas (negativas) são tratadas | ☐ |
| 3.5 | O candidato explica o que acontece na call stack | ☐ |

### Checkpoints Específicos — Ex 04 (Múltiplos)

| # | Critério | Status |
|---|---|:---:|
| 4.1 | A função aceita X como parâmetro (não é hardcoded) | ☐ |
| 4.2 | O operador módulo é usado para verificar divisibilidade | ☐ |
| 4.3 | Para `X = 10`, o retorno é `23` | ☐ |
| 4.4 | Números múltiplos de ambos (ex: 15) não são somados duas vezes | ☐ |
| 4.5 | Edge cases com X ≤ 0 são tratados | ☐ |

### Checkpoints Específicos — Ex 05 (Cadastro de Veículos)

| # | Critério | Status |
|---|---|:---:|
| 5.1 | API retorna respostas em JSON com status HTTP corretos | ☐ |
| 5.2 | Todos os 7 endpoints estão implementados e funcionais | ☐ |
| 5.3 | PUT atualiza todos os campos; PATCH atualiza parcialmente | ☐ |
| 5.4 | Contagem de veículos não vendidos está correta | ☐ |
| 5.5 | Agrupamento por década de fabricação funciona corretamente | ☐ |
| 5.6 | Agrupamento por fabricante funciona corretamente | ☐ |
| 5.7 | Filtro de veículos da última semana funciona corretamente | ☐ |
| 5.8 | Marcas inválidas são rejeitadas com mensagem de erro clara | ☐ |
| 5.9 | O front-end SPA consome a API sem recarregar a página | ☐ |
| 5.10 | Testes unitários cobrem os endpoints e regras de negócio | ☐ |

---

## Critérios Gerais de Avaliação

Estes critérios se aplicam a todos os exercícios e são avaliados transversalmente por Ada:

| Critério | Descrição |
|---|---|
| **Facilidade de configuração** | O projeto deve rodar com poucos comandos; instruções claras no README |
| **Performance** | Soluções não devem ter complexidade desnecessariamente alta |
| **Código limpo** | Legível, sem código morto, sem magic numbers, bem nomeado |
| **Documentação de código** | Comentários onde a intenção não é óbvia pelo código |
| **Documentação do projeto** | README com instruções de instalação, execução e testes |
| **Boas práticas** | Git com commits coesos, `.gitignore`, sem credenciais no repositório |
| **Design Patterns** | Uso consciente de padrões onde aplicável (especialmente no Ex 05) |

---

> **Lembrete da Ada:** *"Antes de submeter, rode todos os seus testes, revise seus commits e verifique se alguém consegue rodar o projeto do zero seguindo apenas o seu README."*
