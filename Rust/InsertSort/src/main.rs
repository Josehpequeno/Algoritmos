use std::io;

fn insert_sort(a: &mut [i32]) {
    for j in 1..a.len() {
        let chave = a[j];
        let mut i = j; //i e j são usize assim não podendo serem negativados ...
        while i > 0 && a[i-1] > chave {
            a.swap(i, i - 1); // por isso essas alterações no algoritmo
            i -= 1;
        }
        a[i] = chave; // i+1 passa a ser i 
    }
}

fn main() {
    println!("Insert Sort");
    println!("Digite os números separados por espaços");
    let mut vec: Vec<i32> = Vec::new();
    loop {
        let mut user_input = String::new();
        let stdin = io::stdin();
        let _result = stdin.read_line(&mut user_input); // lendo as entradas
        let split = user_input.split(" ");
        for s in split {
            println!("{}", s.trim().parse::<i32>().unwrap());
            vec.push(s.trim().parse::<i32>().unwrap()); //retirando espaços e convertendo &str/string to int
        }
        println!("Array antes do Insert Sort {:?}", vec);
        insert_sort(&mut vec);
        println!("Array depois do Insert Sort {:?}", vec);
    }
}
