fs = require('fs');
var data = fs.readFileSync('./input.txt', 'utf8');

data = data.split('\n');

const plays = [
  [4, 8, 3],
  [1, 5, 9],
  [7, 2, 6],
];

const play = (p1, p2) => {
  return plays[letterToNumber(p1)][letterToNumber(p2)];
};

const letterToNumber = (letter) => {
  if ((letter === 'A') | (letter === 'X')) {
    return 0;
  } else if ((letter === 'B') | (letter === 'Y')) {
    return 1;
  } else if ((letter === 'C') | (letter === 'Z')) {
    return 2;
  }
};

// var score = 0;

// for (i = 0; i < data.length - 1; i++) {
//   ligne = data[i].split(' ');
//   score += play(ligne[0], ligne[1]);
// }

// console.log(score);

// Part 2

const choices = [
  ['Z', 'X', 'Y'],
  ['X', 'Y', 'Z'],
  ['Y', 'Z', 'X'],
];

const choiceMaker = (p1, choice) => {
  return choices[p1][choice];
};

const play2 = (p1, choice) => {
  var move = choiceMaker(letterToNumber(p1), letterToNumber(choice));

  return plays[letterToNumber(p1)][letterToNumber(move)];
};

var score2 = 0;

for (i = 0; i < data.length - 1; i++) {
  ligne = data[i].split(' ');
  score2 += play2(ligne[0], ligne[1]);
}

console.log(score2);
