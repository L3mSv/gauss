use rand::Rng;
use std::io;
use std::time::Instant;

//impl <T> SliceIndex<[T]> for i32;
// asjdhaiosu
//const MAX:usize = 100;

fn printa_matrizes(a: &mut Vec<Vec<f32>>, b: &mut Vec<f32>, x: &mut Vec<f32>, size: usize) {
    println!("A:");
    for i in 0..size {
        for j in 0..size {
            print!("{} ", a[i][j]);
        }
        println!("");
    }

    println!("B:");
    for i in 0..size {
        print!("{} ", b[i]);
    }
    println!("");

    println!("X:");
    for i in 0..size {
        print!("{} ", x[i]);
    }
    println!("");
}

fn main() {
    println!("Insira o tamanho da matriz");
    let mut size = String::new();

    io::stdin().read_line(&mut size).expect("Erro ao ler input");

    let size: usize = size.trim().parse().expect("Digite um numero");

    let start = Instant::now();

    //println!("{:.2?}", start.elapsed());

    //define outras variaveis
    /*let mut norm:usize;
    let mut row:usize;
    let mut col:usize;*/

    let mut mult:f32;

    //define as matrizes
    //A * X = B, X sera o output do programa
    let mut a: Vec<Vec<f32>> = vec![vec![0.0; size]; size];
    let mut b: Vec<f32> = vec![0.0; size];
    let mut x: Vec<f32> = vec![0.0; size];

    //inicializa as matrizes
    for i in 0..size {
        for j in 0..size {
            a[i][j] = rand::rng().random_range(1..=10) as f32;
        }

        b[i] = rand::rng().random_range(1..=10) as f32;
    }

    printa_matrizes(&mut a, &mut b, &mut x, size);

    //executa o gauss
    for norm in 0..size - 1 {
        for row in norm + 1..size {
            mult = a[row][norm] / a[norm][norm];
            for col in norm..size {
                a[row][col] -= a[norm][col] * mult;
            }
            b[row] -= b[norm] * mult;
        }
    }

    printa_matrizes(&mut a, &mut b, &mut x, size);

    for row in (0..size-1).rev() {
        x[row] = b[row];
        for col in (row..size-1).rev() {
            x[row] -= a[row][col] * x[col];
        }
        x[row] /= a[row][row];
    }

    printa_matrizes(&mut a, &mut b, &mut x, size);

    println!("{:.2?}", start.elapsed());
}
