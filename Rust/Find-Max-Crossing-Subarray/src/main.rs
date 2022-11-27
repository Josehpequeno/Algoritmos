use std::io::{self, Write};

struct Tupla {
    low: usize,
    high: usize,
    value: f64,
}

fn find_maximum_crossing_subarray(a: &[f64], low: usize, high: usize) -> Tupla {
    if low == high {
        return Tupla {
            low: low,
            high: low,
            value: a[low],
        };
    } else {
        let mid = (high + low) / 2;
        let left = find_maximum_crossing_subarray(a, low, mid);
        let right = find_maximum_crossing_subarray(a, mid + 1, high);
        let cross = find_max_crossing_subarray(a, low, mid, high);
        if left.value >= right.value && left.value >= cross.value {
            return left;
        } else if right.value >= left.value && right.value >= cross.value {
            return right;
        } else {
            return cross;
        }
    }
}

fn find_max_crossing_subarray(a: &[f64], low: usize, mid: usize, high: usize) -> Tupla {
    let mut left_sum = f64::MIN;
    let mut sum = 0 as f64;
    let mut max_left = 0;
    let mut max_right = 0;
    let mut i: i32 = mid as i32;
    while i >= (low as i32) {
        sum += a[i as usize];
        if sum > left_sum {
            left_sum = sum;
            max_left = i as usize;
        }
        i -= 1;
    }
    let mut right_sum = f64::MIN;
    sum = 0 as f64;
    for j in mid + 1..high + 1 {
        sum += a[j];
        if sum > right_sum {
            right_sum = sum;
            max_right = j;
        }
    }
    sum = left_sum + right_sum;
    return Tupla {
        low: max_left,
        high: max_right,
        value: sum,
    };
}

fn generate_changes_array(a: &mut [f64]) -> Vec<f64> {
    let mut changes: Vec<f64> = Vec::new();
    for i in 1..a.len() {
        let val = (a[i] - a[i - 1]) as f64;
        changes.push(val);
    }
    return changes;
}

fn main() {
    println!("Find Max Crossing Subarray");
    println!("Enter the prices separated by spaces\n* type clear to clear array, quit to quit");
    let mut vec: Vec<f64> = Vec::new();
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
            match trimmed.parse::<f64>() {
                //converting &str/string to int and checking letters and other characters
                Ok(i) => vec.push(i),
                Err(..) => continue,
            };
        }
        println!("Array to Find Max Crossing Subarray {:?}", vec);
        let len: i32 = vec.len().try_into().unwrap();
        let changes = generate_changes_array(&mut vec);
        println!("Array of changes:  {:?}", changes);
        if len > 1 {
            let high: usize = changes.len().try_into().unwrap();
            let result = find_maximum_crossing_subarray(&changes, 0, high - 1);
            println!("Buy in {} price: {}", result.low, vec[result.low]);
            println!(
                "Sale in {} price: {}",
                result.high + 1,
                vec[result.high + 1]
            );
            println!("Earnings per share: {}", result.value)
        }
        println!("-------------------------");
    }
}
