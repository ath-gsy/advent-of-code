fs = require('fs');
// const rawData = fs.readFileSync('./AOC-2022/day14/input.txt', 'utf8');
const rawData = fs.readFileSync('./input.txt', 'utf8');

///Initializing the cave plan
var cavePlan = new Array(200);
for (var i = 0; i < cavePlan.length; i++) {
  cavePlan[i] = new Array(1000);
}
for (var i = 0; i < cavePlan.length; i++) {
  for (var j = 0; j < cavePlan[0].length; j++) {
    cavePlan[i][j] = '.';
  }
}
// cavePlan[0][499] = '+';
///
///Functions
const cavePlanPrinter = (cavePlan) => {
  totalString = '';
  for (var i = 0; i < cavePlan.length; i++) {
    totalString += cavePlan[i].join('') + '\n';
  }
  //   console.log(cavePlan[i]);
  fs.writeFile('./cavePlan.txt', totalString, (err) => {
    if (err) {
      console.error(err);
    }
    // file written successfully
  });
};

const drawLine = (start, end) => {
  //   console.log(start, end);
  if (start.x === end.x) {
    for (var j = Math.min(start.y, end.y); j <= Math.max(start.y, end.y); j++) {
      cavePlan[j - 1][start.x - 1] = '#';
    }
  } else {
    for (var j = Math.min(start.x, end.x); j <= Math.max(start.x, end.x); j++) {
      cavePlan[start.y - 1][j - 1] = '#';
    }
  }
};
///
///Repérage des lignes de rochers
var lines = rawData.split('\n');

for (var i = 0; i < lines.length; i++) {
  var line = lines[i].split(' -> ');
  line = [
    ...line.map(
      (coord) =>
        (coord = {
          x: parseInt(coord.split(',')[0]),
          y: parseInt(coord.split(',')[1]),
        })
    ),
  ];

  for (var l = 0; l < line.length - 1; l++) {
    drawLine(line[l], line[l + 1]);
  }
}
///1er print de la cave (je suis quelqu'un de très visuel)
cavePlanPrinter(cavePlan);
///
///Fonction récursive de falling des grains
const fallingSand = (startingPoint) => {
  for (var i = startingPoint.y; i < cavePlan.length - 1; i++) {
    if (i === cavePlan.length - 2) {
      console.log('no bound');
      return [false];
    }

    if (cavePlan[i + 1][startingPoint.x] === 'o' || cavePlan[i + 1][startingPoint.x] === '#') {
      if (cavePlan[i + 1][startingPoint.x - 1] === 'o' || cavePlan[i + 1][startingPoint.x - 1] === '#') {
        if (cavePlan[i + 1][startingPoint.x + 1] === 'o' || cavePlan[i + 1][startingPoint.x + 1] === '#') {
          cavePlan[i][startingPoint.x] = 'o';

          // cavePlanPrinter(cavePlan);

          return [true, { x: startingPoint.x, y: i }];
        } else {
          // console.log('right free');
          return fallingSand({ x: startingPoint.x + 1, y: i });
        }
      } else {
        // console.log('left free');
        return fallingSand({ x: startingPoint.x - 1, y: i });
      }
    }
  }
};
///
/// COUNTING
// COUNTER = 0;
// var inside = true;
// while (inside) {
//   COUNTER++;
//   inside = fallingSand({ x: 499, y: 1 })[0];
//   if (COUNTER % 10 == 0) {
//     cavePlanPrinter(cavePlan);
//   }
// }
// console.log(COUNTER - 1);
///

//// PART 2 ////
drawLine({ x: 0, y: 164 }, { x: cavePlan[0].length, y: 164 });
cavePlanPrinter(cavePlan);

const fallingSand2 = (startingPoint) => {
  for (var i = startingPoint.y; i < cavePlan.length - 1; i++) {
    if (cavePlan[0][499] === 'o') {
      console.log('full');
      return [false];
    }

    if (cavePlan[i + 1][startingPoint.x] === 'o' || cavePlan[i + 1][startingPoint.x] === '#') {
      if (cavePlan[i + 1][startingPoint.x - 1] === 'o' || cavePlan[i + 1][startingPoint.x - 1] === '#') {
        if (cavePlan[i + 1][startingPoint.x + 1] === 'o' || cavePlan[i + 1][startingPoint.x + 1] === '#') {
          cavePlan[i][startingPoint.x] = 'o';

          // cavePlanPrinter(cavePlan);

          return [true, { x: startingPoint.x, y: i }];
        } else {
          // console.log('right free');
          return fallingSand({ x: startingPoint.x + 1, y: i });
        }
      } else {
        // console.log('left free');
        return fallingSand({ x: startingPoint.x - 1, y: i });
      }
    }
  }
};

COUNTER = 0;
var notBlocked = true;
while (notBlocked) {
  COUNTER++;
  notBlocked = fallingSand2({ x: 499, y: 0 })[0];
  // if (cavePlan[0][499] === 'o') {
  //   notBlocked = false;
  // }
}
cavePlanPrinter(cavePlan);
console.log(COUNTER - 1);

// cavePlanPrinter(cavePlan);
// for (var i = 0; i < 24727; i++) {
//   fallingSand2({ x: 499, y: 0 });
// }
// cavePlanPrinter(cavePlan);
