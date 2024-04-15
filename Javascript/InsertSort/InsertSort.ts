import * as readline from "readline";

function insertSort(array: number[]): void {
  for (let j = 1; j < array.length; j++) {
    let key = array[j];
    let i = j - 1;
    while (i >= 0 && array[i] > key) {
      array[i + 1] = array[i];
      i--;
    }
    array[i + 1] = key;
  }
}

function main(): void {
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
  });

  console.log("Insert Sort");
  console.log(
    "Enter the numbers separated by spaces\n* type clear to clear array, quit to quit"
  );
  console.log("---------------------");
  let array: number[] = [];

  rl.on("line", (input) => {
    const text = input.trim();
    if (text === "clear") {
      array = [];
      console.log("Array empty");
      return;
    }
    if (text === "quit") {
      rl.close();
      return;
    }

    const numbers = text.split(/\s+/);
    for (let number of numbers) {
      const num = parseInt(number);
      if (!isNaN(num)) {
        array.push(num);
      }
    }

    console.log("----------------------");
    console.log("Array before Insert Sort: ", array);
    insertSort(array);
    console.log("Array after Insert Sort: ", array);
    console.log("----------------------");
  });
}

main();
