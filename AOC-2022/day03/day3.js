fs = require('fs');
var data = fs.readFileSync('./input.txt', 'utf8');
data = data.split('\n');

var items = [];

const compareCompartments = (comp1, comp2) => {
  for (var i = 0; i < comp1.length; i++) {
    for (var j = 0; j < comp2.length; j++) {
      if (comp1[i] === comp2[j]) {
        return items.push(comp1[i]);
      }
    }
  }
};

// console.log(compareCompartments('nsmsVQVZVQmTRmddrL', 'gfFhfLgLDzgvffDhbF'));
// console.log(items);

for (var i = 0; i < data.length; i++) {
  comp1 = data[i].slice(0, data[i].length / 2);
  comp2 = data[i].slice(data[i].length / 2, data[i].length);
  //   console.log(comp1, comp2);
  compareCompartments(comp1, comp2);
}

console.log(items);
// console.log(items[199]);

const priorities = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'];

const priority = (items) => {
  var priorityScore = 0;
  for (var i = 0; i < items.length; i++) {
    for (var j = 0; j < priorities.length; j++) {
      if (items[i] === priorities[j]) {
        priorityScore += j + 1;
        console.log(priorityScore);
      }
    }
  }
  return priorityScore;
};

console.log(priority(items));
