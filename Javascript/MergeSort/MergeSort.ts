import * as readline from "readline";

function merge(array: number[], p: number, q: number, r: number): void {
  let n1 = q - p;
  let n2 = r - q;

  let L = new Array(n1 + 1);
  let R = new Array(n2 + 1);
  for (let i = 0; i < n1; i++) {
    L[i] = array[p + i];
  }
  for (let j = 0; j < n2; j++) {
    R[j] = array[q + j];
  }
  console.log("===============>*<===============");
  console.log("Left array", L.slice(0, L.length - 1));
  console.log("Right array", R.slice(0, R.length - 1));
  L[n1] = Number.POSITIVE_INFINITY;
  R[n2] = Number.POSITIVE_INFINITY;

  let i = 0;
  let j = 0;
  for (let k = p; k < r; k++) {
    if (L[i] <= R[j]) {
      array[k] = L[i];
      i++;
    } else {
      array[k] = R[j];
      j++;
    }
  }
}

function mergeSort(array: number[], p: number, r: number): void {
  if (p < r - 1) {
    let q = Math.floor((p + r) / 2.0);
    mergeSort(array, p, q);
    mergeSort(array, q, r);
    merge(array, p, q, r);
  }
}

function main(): void {
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
  });

  console.log("Merge Sort");
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
    console.log("Array before Merge Sort: ", array);
    mergeSort(array, 0, array.length); // array, start, length
    console.log("Array after Merge Sort: ", array);
    console.log("----------------------");
  });
}

main();
