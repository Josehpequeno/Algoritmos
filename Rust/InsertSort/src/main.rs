use std::io::{self, Write};

fn insert_sort(a: &mut [i32]) {
    for j in 1..a.len() {
        let chave = a[j];
        let mut i = j; //i and j are usize, so they can't be negative...
        while i > 0 && a[i - 1] > chave {
            a.swap(i, i - 1); // so these changes in the algorithm
            i -= 1;
        }
        a[i] = chave; // where is i+1 becomes i and so on
    }
}

fn main() {
    println!("Insert Sort");
    println!("Enter the numbers separated by spaces\n* type clear to clear array, quit to quit");
    let mut vec: Vec<i32> = Vec::new();
    loop {
        print!("=> ");
        let _result_flush = io::stdout().flush();
        let mut user_input = String::new();
        let stdin = io::stdin();
        let _result = stdin.read_line(&mut user_input); // reading the entries
        if user_input.trim() == "clear" {
            vec.clear();
            continue;
        }
        if user_input.trim() == "quit" {
            break;
        }
        let split = user_input.split(" ");
        for s in split {
            let trimmed = s.trim(); //removing spaces
            match trimmed.parse::<i32>() {
                //converting &str/string to int and checking letters and other characters
                Ok(i) => vec.push(i),
                Err(..) => continue,
            };
        }
        println!("Array before Insert Sort {:?}", vec);
        insert_sort(&mut vec);
        println!("Array after Insert Sort {:?}", vec);
    }
}
