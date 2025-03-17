# Eliminação de Gauss 

**Objetivo**: Compreender as características de Rust e Golang comparando programas implementados nestas linguagens com a implementação em C.

*Projetado sobre [Gaussian Elimination](https://github.com/gmendonca/gaussian-elimination-pthreads-openmp.git)*


## *Componentes do grupo:*
- Lucas Lemes <br/>
- Felipe Larangeira <br/>
- Gabriel de Souza Maciel  <br/>
- Frederico Larangeira <br/>


## *Utilização do programa:* 

> [!IMPORTANT] 
> Os programas foram executados e testados em **máquina Windows**.


## *Em C:*

Para executar o código de `gauss.c`, digite:

```
    gcc gauss.c -o gauss.out 
    ./gauss.out <matrix_dimensions> [random seed]
```

A saída no intervalo de **[1, 10]** deve se parecer com:

```
    Matrix dimension N = 5.

    Initializing...

    A =
            0.45,  0.70,  0.39,  0.02,  0.16;
            0.47,  0.67,  0.51,  0.23,  0.23;
            0.80,  0.90,  0.42,  0.90,  0.82;
            0.66,  0.76,  0.68,  0.77,  0.75;
            0.09,  0.07,  0.99,  0.75,  0.21;

    B = [ 0.61;  0.25;  0.25;  0.04;  0.57]

    Starting clock.
    Computing Serially.
    Stopped clock.

    Elapsed time = 0.0916 ms.

    X = [-74.89; 51.54; -7.39; 11.85;  7.83]
```

A partir de números maiores a saída deve ser:

```
Matrix dimension N = 10.

Initializing...

Starting clock.
Computing Serially.
Stopped clock.

Elapsed time = 0.0838 ms.

X = [-3.75;  1.31;  0.75; -0.31;  2.15; -1.47;  1.93;  0.80;  4.84; -4.57]
```

## *Em Golang:*

> [!WARNING] 
> No intervalo de valores **[10,90]** para a dimensão da matriz, o tempo de execução do programa é impresso como 0 segundos!
> Infelizmente não encontramos solução, peço desculpas. 

Para executar o código de `gauss.go`, digite:

```
    go run gauss.go 
```

Após iniciar a execução, o programa irá pedir ao usuário para informar a dimensão da matriz entre **1** e **2000**.

A saída deve se parecer com isso:

```
Enter matrix dimension N (1 to 2000): 5

Initializing...

Matrix A:
40.49, 62.98, 79.01, 29.03, 13.38;
39.15, 46.83,  3.62, 67.78, 18.07;
7.25, 84.64, 83.44, 21.44, 97.68;
55.11, 12.64,  0.40, 92.67, 47.76;
38.51, 49.31,  6.61, 29.23, 49.03;

Vector B:
[97.37; 68.52; 79.67; 28.52; 20.30]
Computing Serially.

Elapsed time = 0 ms.

Solution Vector X:
[-0.59;  0.91;  0.60;  0.86; -0.63]
```

Para números maiores que **10** a saída deve ser:

```
Enter matrix dimension N (1 to 2000): 645

Initializing...
Computing Serially.

Elapsed time = 58 ms.
```

## *Em Rust:*

> [!NOTE]
> Para executar o arquivc ```gauss.rs```, este deve estar dentro de uma pasta junto ao arquivo ```cargo.toml```.

Para executar o arquivo `gauss.rs`:

```
cargo build 
cargo run

```

A saida deve se parecer com:

```
Insira o tamanho da matriz
5
A:
1 2 6 6 7 
3 9 3 3 9 
4 7 4 9 5 
5 2 7 5 8 
9 9 4 9 6 
B:
5 8 4 2 1 
X:
0 0 0 0 0 
A:
1 2 6 6 7 
0 3 -15 -15 -12 
0 0 -25 -20 -27 
0 0 0 -14.599998 9.040001 
0 0 0 0 0.9315033 
B:
5 -7 -18.333334 4.533333 0.31963968 
X:
0 0 0 0 0 
A:
1 2 6 6 7 
0 3 -15 -15 -12 
0 0 -25 -20 -27 
0 0 0 -14.599998 9.040001 
0 0 0 0 0.9315033
B:
5 -7 -18.333334 4.533333 0.31963968
X:
0 -1145.0795 119.81845 -4.8438354 0
2.47ms

```
