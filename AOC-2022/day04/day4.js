fs = require('fs');
var data = fs.readFileSync('./input.txt', 'utf8');
data = data.split('\n');

data = data.map((line) => line.split(','));

console.log(data);

const contains = (range1, range2) => {
  if (range1[0] <= range2[0] && range1[1] >= range2[1]) {
    return true;
  } else {
    return false;
  }
};

//range1 isContained in range2 ?
const isContained = (range1, range2) => {
  if (contains(range1, range2) | contains(range2, range1)) {
    return true;
  } else {
    return false;
  }
};

var counter = 0;

for (var i = 0; i < data.length - 1; i++) {
  range1 = data[i][0].split('-').map((x) => parseInt(x));
  range2 = data[i][1].split('-').map((x) => parseInt(x));
  //   console.log(range1, range2);
  if (isContained(range1, range2)) {
    counter += 1;
  }
}

console.log(counter);
