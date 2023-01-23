use std::io::{ self, Write };

#[derive(Debug)]
struct Matrix {
    array: Vec<f64>,
    rows: usize,
    colums: usize,
}

fn square_matrix_mutiply(a: &mut Matrix, b: &mut Matrix) {
    println!("{:?} {:?}", a, b)
}

fn main() {
    println!("Square Matrix Mutiply 2x2");
    println!("Enter the number of square matrix separated by spaces\n* type quit to quit");
    let mut first: Vec<f64> = Vec::new();
    let mut second: Vec<f64> = Vec::new();
    let mut a_matrix = Matrix { array: Vec::new(), colums: 0, rows: 0 };
    let mut b_matrix = Matrix { array: Vec::new(), colums: 0, rows: 0 };
    loop {
        loop {
            print!("first array => ");
            let mut _result_flush = io::stdout().flush();
            let mut user_input = String::new();
            let stdin = io::stdin();
            let mut _result = stdin.read_line(&mut user_input);
            if user_input.trim() == "quit" {
                break;
            }
            let mut split = user_input.split(" ");
            for s in split {
                let trimmed = s.trim();
                match trimmed.parse::<f64>() {
                    Ok(i) => first.push(i),
                    Err(..) => {
                        continue;
                    }
                }
            }
            let sqrt = (first.len() as f64).sqrt();
            let trunc = sqrt.trunc();
            if sqrt == trunc {
                a_matrix.rows = trunc as usize;
                a_matrix.colums = trunc as usize;
                for i in 0..first.len() {
                    a_matrix.array.push(first[i]);
                }
                break;
            } else {
                println!("Square matrix is not a square matrix {:?}", first);
            }
        }

        loop {
            print!("second array => ");
            let mut _result_flush = io::stdout().flush();
            let mut user_input = String::new();
            let stdin = io::stdin();
            let mut _result = stdin.read_line(&mut user_input);
            if user_input.trim() == "quit" {
                break;
            }
            let mut split = user_input.split(" ");
            for s in split {
                let trimmed = s.trim();
                match trimmed.parse::<f64>() {
                    Ok(i) => second.push(i),
                    Err(..) => {
                        continue;
                    }
                }
            }
            let sqrt = (second.len() as f64).sqrt();
            let trunc = sqrt.trunc();
            if sqrt == trunc {
                b_matrix.rows = trunc as usize;
                b_matrix.colums = trunc as usize;
                for i in 0..second.len() {
                    b_matrix.array.push(second[i]);
                }
                break;
            } else {
                println!("Square matrix is not a square matrix {:?}", second);
            }
        }
        println!("First Matrix {:?}", first);
        println!("Second Matrix {:?}", second);
        square_matrix_mutiply(&mut a_matrix, &mut b_matrix);
        first.clear();
        second.clear();
    }
}