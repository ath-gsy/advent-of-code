fs = require('fs');
const data = fs.readFileSync('./day01/input.txt', 'utf8');
var elves = data
  .split('\n\n')
  .map((line) => line.split('\n'))
  .map((tab) => tab.map((number) => parseInt(number)))
  .map((tabl) => tabl.reduce((acc, curr) => acc + curr));
console.log(Math.max(...elves.slice(0, elves.length - 1)));
