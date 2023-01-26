use std::io::{ self, Write };

#[derive(Debug)]
struct Matrix {
    array: Vec<MatrixArray>,
    rows: usize,
    columns: usize,
}

#[derive(Debug)]
struct MatrixArray {
    values: Vec<f64>,
}

fn square_matrix_mutiply(a: &mut Matrix, b: &mut Matrix) -> Matrix {
    let n = a.rows;
    let mut c_array: Vec<MatrixArray> = Vec::new();
    for _ in 0..n {
        c_array.push(MatrixArray { values: Vec::new() });
    }
    let mut c = Matrix { array: c_array, rows: n, columns: n };
    if n == 1 {
        c.array[0].values.push(a.array[0].values[0] * b.array[0].values[0]);
    } else {
        let mut m1 = Matrix {
            array: vec![MatrixArray { values: vec![a.array[0].values[0]] }],
            rows: 1,
            columns: 1,
        };
        let mut m2 = Matrix {
            array: vec![MatrixArray { values: vec![b.array[0].values[0]] }],
            rows: 1,
            columns: 1,
        };
        let mut m3 = Matrix {
            array: vec![MatrixArray { values: vec![a.array[0].values[1]] }],
            rows: 1,
            columns: 1,
        };
        let mut m4 = Matrix {
            array: vec![MatrixArray { values: vec![b.array[1].values[0]] }],
            rows: 1,
            columns: 1,
        };
        c.array[0].values.push(
            square_matrix_mutiply(&mut m1, &mut m2).array[0].values[0] +
                square_matrix_mutiply(&mut m3, &mut m4).array[0].values[0]
        );

        m1 = Matrix {
            array: vec![MatrixArray { values: vec![a.array[0].values[0]] }],
            rows: 1,
            columns: 1,
        };
        m2 = Matrix {
            array: vec![MatrixArray { values: vec![b.array[0].values[1]] }],
            rows: 1,
            columns: 1,
        };
        m3 = Matrix {
            array: vec![MatrixArray { values: vec![a.array[0].values[1]] }],
            rows: 1,
            columns: 1,
        };
        m4 = Matrix {
            array: vec![MatrixArray { values: vec![b.array[1].values[1]] }],
            rows: 1,
            columns: 1,
        };
        c.array[0].values.push(
            square_matrix_mutiply(&mut m1, &mut m2).array[0].values[0] +
                square_matrix_mutiply(&mut m3, &mut m4).array[0].values[0]
        );

        m1 = Matrix {
            array: vec![MatrixArray { values: vec![a.array[1].values[0]] }],
            rows: 1,
            columns: 1,
        };
        m2 = Matrix {
            array: vec![MatrixArray { values: vec![b.array[0].values[0]] }],
            rows: 1,
            columns: 1,
        };
        m3 = Matrix {
            array: vec![MatrixArray { values: vec![a.array[1].values[1]] }],
            rows: 1,
            columns: 1,
        };
        m4 = Matrix {
            array: vec![MatrixArray { values: vec![b.array[1].values[0]] }],
            rows: 1,
            columns: 1,
        };
        c.array[1].values.push(
            square_matrix_mutiply(&mut m1, &mut m2).array[0].values[0] +
                square_matrix_mutiply(&mut m3, &mut m4).array[0].values[0]
        );

        m1 = Matrix {
            array: vec![MatrixArray { values: vec![a.array[1].values[0]] }],
            rows: 1,
            columns: 1,
        };
        m2 = Matrix {
            array: vec![MatrixArray { values: vec![b.array[0].values[1]] }],
            rows: 1,
            columns: 1,
        };
        m3 = Matrix {
            array: vec![MatrixArray { values: vec![a.array[1].values[1]] }],
            rows: 1,
            columns: 1,
        };
        m4 = Matrix {
            array: vec![MatrixArray { values: vec![b.array[1].values[1]] }],
            rows: 1,
            columns: 1,
        };
        c.array[1].values.push(
            square_matrix_mutiply(&mut m1, &mut m2).array[0].values[0] +
                square_matrix_mutiply(&mut m3, &mut m4).array[0].values[0]
        );
    }
    return c;
}

fn main() {
    println!("Square Matrix Mutiply 2x2");
    println!("Enter the number of square matrix separated by spaces\n* type quit to quit");
    let mut first: Vec<f64> = Vec::new();
    let mut second: Vec<f64> = Vec::new();
    let a_array = vec![MatrixArray { values: Vec::new() }, MatrixArray { values: Vec::new() }];
    let b_array = vec![MatrixArray { values: Vec::new() }, MatrixArray { values: Vec::new() }];
    let mut a_matrix = Matrix { array: a_array, columns: 0, rows: 0 };
    let mut b_matrix = Matrix { array: b_array, columns: 0, rows: 0 };
    'main_loop: loop {
        loop {
            print!("first array => ");
            let mut _result_flush = io::stdout().flush();
            let mut user_input = String::new();
            let stdin = io::stdin();
            let mut _result = stdin.read_line(&mut user_input);
            if user_input.trim() == "quit" {
                break 'main_loop;
            }
            let split = user_input.split(" ");
            for s in split {
                if first.len() >= 4 {
                    break;
                }
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
            let mut count = 0;
            let mut it = 0;
            if sqrt == trunc {
                a_matrix.rows = trunc as usize;
                a_matrix.columns = trunc as usize;
                for i in 0..first.len() {
                    if count >= (sqrt as i32) {
                        it += 1;
                        count = 0;
                    }
                    a_matrix.array[it].values.push(first[i]);
                    count += 1;
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
                break 'main_loop;
            }
            let split = user_input.split(" ");
            for s in split {
                if second.len() >= 4 {
                    break;
                }
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
            let mut count = 0;
            let mut it = 0;
            if sqrt == trunc {
                b_matrix.rows = trunc as usize;
                b_matrix.columns = trunc as usize;
                for i in 0..second.len() {
                    if count >= (sqrt as i32) {
                        it += 1;
                        count = 0;
                    }
                    b_matrix.array[it].values.push(second[i]);
                    count += 1;
                }
                break;
            } else {
                println!("Square matrix is not a square matrix {:?}", second);
            }
        }
        println!("First Matrix {:?}", a_matrix);
        println!("Second Matrix {:?}", b_matrix);
        if a_matrix.rows != b_matrix.rows {
            println!("Square matrices with different number of rows");
        } else if a_matrix.columns != b_matrix.columns {
            println!("Square matrices with different number of columns");
        } else {
            let c = square_matrix_mutiply(&mut a_matrix, &mut b_matrix);
            println!("Result Matrix {:?}", c);
        }
        first.clear();
        second.clear();
        a_matrix.array[0].values.clear();
        a_matrix.array[1].values.clear();
        b_matrix.array[0].values.clear();
        b_matrix.array[1].values.clear();
    }
}