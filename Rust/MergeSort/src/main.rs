use std::io::{self, Write};

fn merge_sort(a: &mut [i32], p: i32, r: i32) {
    if r - p > 1 {
        let q = (p + r) / 2;
        merge_sort(a, p, q);
        merge_sort(a, q, r);
        merge(a, p, q, r);
    }
}

fn merge(a: &mut [i32], p: i32, q: i32, e: i32) {
    let n1 = q - p;
    let n2 = e - q;
    let mut l: Vec<i32> = Vec::with_capacity((n1 + 1) as usize);
    let mut r: Vec<i32> = Vec::with_capacity((n2 + 1) as usize);
    for i in 0..n1 {
        l.push(i32::MAX);
        l[i as usize] = a[(p + i) as usize]
    }
    for j in 0..n2 {
        r.push(i32::MAX);
        r[j as usize] = a[(q + j) as usize]
    }
    println!("===============>*<===============");
    println!("Left array {:?}", l);
    println!("Right array {:?}", r);
    l.push(i32::MAX);
    r.push(i32::MAX);
    let mut i = 0;
    let mut j = 0;
    for k in p..e {
        if l[i] <= r[j] {
            a[k as usize] = l[i];
            i = i + 1;
        } else {
            a[k as usize] = r[j];
            j = j + 1;
        }
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
        println!("Array before Merge Sort {:?}", vec);
        let len: i32 = vec.len().try_into().unwrap();
        merge_sort(&mut vec, 0, len); // start, len
        println!("===============>*<===============");
        println!("Array after Merge Sort {:?}", vec);
    }
}
