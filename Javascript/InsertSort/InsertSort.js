const readline = require("readline");

function insertSort(array) {
  for (let j = 1; j < array.length; j++) {
    let key = array[j];
    i = j - 1;
    while (i >= 0 && array[i] > key) {
      array[i + 1] = array[i];
      i--;
    }
    array[i + 1] = key;
  }
}

function main() {
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
  });

  console.log("Insert Sort");
  console.log(
    "Enter the numbers separated by spaces\n* type clear to clear array, quit to quit"
  );
  console.log("---------------------");
  let array = [];

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

    const numbers = text.split(/\s+/); // Break input into numbers
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
