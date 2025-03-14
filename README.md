# Eliminacao de Gauss 

**Objetivo**: Compreender as características de Rust e Golang comparando programas implementados nestas linguagens com a implementação em C

Projetado sobre ([Gaussian Elimination](https://github.com/gmendonca/gaussian-elimination-pthreads-openmp.git))


## Componentes do grupo:
- Lucas Lemes <br/>
- Felipe Laranjeira <br/>
- Gabriel de Souza Maciel  <br/>
- Frederico Laranjeira <br/>


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
    Enter matrix dimension N (1 to 2000): 23

    Initializing...
    Computing Serially.

    Elapsed time = 0 ms.

    Solution Vector X:
    [ 1.01; -0.53; -0.23;  0.49; -1.40; -0.04; -0.64; -0.10; -0.70; -0.79; -1.79;  1.09;  0.86; -0.18;  1.22; -1.18;  1.04;  1.30; -1.13;  0.96;  0.13;  0.08;  0.66]
```