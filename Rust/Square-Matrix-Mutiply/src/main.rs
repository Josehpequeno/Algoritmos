use std::io::{self, Write};

struct Matrix {
    array: Vec<f64>,
    rows: usize,
    colums: usize
}

fn find_max_crossing_subarray (a: Matrix, b: Matrix ) {
    
}

fn main() {
    println!("Square Matrix Mutiply 2x2");
    println!("Enter the number of square matrix separated by spaces\n* type quit to quit");
    let mut first: Vec<f64> = Vec::new();
    let mut second: Vec<f64> = Vec::new();
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
                Err(..) => continue,
            };
        }
        print!("second array => ");
        _result_flush = io::stdout().flush();
        user_input = String::new();
        _result = stdin.read_line(&mut user_input);
        if user_input.trim() == "quit" {
            break;
        }
        split = user_input.split(" ");
        for s in split {
            let trimmed = s.trim();
            match trimmed.parse::<f64>() {
                Ok(i) => second.push(i),
                Err(..) => continue,
            };
        }
        println!("First Matrix {:?}", first);
        println!("Second Matrix {:?}", second);

    }
}
