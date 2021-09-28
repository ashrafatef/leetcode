let prompt = require("prompt-sync")();

const number = prompt("What is your name?");

const MIN_THRESHOLD = 2 ** 31 * -1;
const MAX_THRESHOLD = 2 ** 31 - 1;

function reverse(x) {
  let isNegative = false;
  if (x == 0) {
    return 0;
  }
  if (x < 0) {
    isNegative = true;
    x *= -1;
  }
  let reversedNumber = x.toString().split("").reverse();
  for (let index = 0; index < reversedNumber.length; index++) {
    const element = reversedNumber[index];
    if (parseInt(element, 10) == 0) {
      reversedNumber.shift();
      index -= 1;
    }
    if (parseInt(element, 10) != 0) {
      break;
    }
  }
  let finalNumber = parseInt(reversedNumber.join(""), 10);
  finalNumber = isNegative ? finalNumber * -1 : finalNumber;
  if (finalNumber <= MIN_THRESHOLD || finalNumber >= MAX_THRESHOLD) {
    return 0;
  }
  return finalNumber;
}
//-65090

reverse(number);
