# Eliminacao de Gauss 

**Objetivo**: Compreender as características de Rust e Golang comparando programas implementados nestas linguagens com a implementação em C

Projetado sobre [Gaussian Elimination](https://github.com/gmendonca/gaussian-elimination-pthreads-openmp.git) 


## Componentes do grupo:
- Lucas Lemes <br/>
- Felipe Larangeira <br/>
- Gabriel de Souza Maciel  <br/>
- Frederico Larangeira <br/>


## Utilizacao do programa: 

> [!IMPORTANT] 
> Os programas foram executados e testados em maquina Windows

Para executar o codigo de `gauss.c`, digite:

```
    gcc gauss.c -o gauss.out 
    ./gauss.out <matrix_dimensions> [random seed]
```

A saida deve se parecer com:

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

Para executar o codigo de `gauss.go`, digite:

```
    go run gauss.go 
```

Apos iniciar a execucao, o programa ira pedir ao usuario para informar a dimensao da matriz entre 1 e 2000

A saida deve se parecer com isso:

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