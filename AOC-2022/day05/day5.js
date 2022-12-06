fs = require('fs');
var data = fs.readFileSync('./input.txt', 'utf8');
data = data.split('\n\n');

// console.log(data);

class Pile {
  constructor(initialStack) {
    this.stack = initialStack;
    this.height = initialStack.length;
  }

  isEmpty() {
    if (this.height <= 0) {
      return true;
    }
    return false;
  }
  pop() {
    if (!this.isEmpty()) {
      this.stack = this.stack.slice(0, this.height - 1);
      this.height--;
    } else {
      return false;
    }
  }

  add(element) {
    this.stack.push(element);
    this.heigth++;
  }

  get() {
    return this.stack;
  }
}

console.log(data[0]);

initialConfig = data[0].split('\n');
numberLines = initialConfig.length;

initialConfig = initialConfig.map((line) => line.split(''));
console.log(initialConfig);

const findIndex = (pileNumber) => {
  for (var i = 0; i < initialConfig[numberLines - 1].length - 1; i++) {
    if (pileNumber === parseInt(initialConfig[numberLines - 1][i])) {
      return (indexPile = i);
    }
  }
};

// console.log(findIndex(1));

const initStack = (pileNumber) => {
  indexPile = findIndex(pileNumber);
  stack = [];

  for (var i = numberLines - 1; i >= 0; i--) {
    console.log(initialConfig[i][indexPile]);
    if (initialConfig[i][indexPile] === '') {
      return new Pile(stack);
    } else {
      stack.push(initialConfig[i][indexPile]);
    }
  }
};

pile1 = initStack(1);
console.log(typeof pile1);

pile1.get();
