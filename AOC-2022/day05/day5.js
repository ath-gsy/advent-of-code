fs = require('fs');
var data = fs.readFileSync('./input.txt', 'utf8');
data = data.split('\n\n');

///CLASS
class Pile {
  constructor(initialStack) {
    this.stack = initialStack;
    this.height = this.stack.length;
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

  set(table) {
    this.stack = table;
    this.height = table.length;
  }
}
///

initialConfig = data[0].split('\n');
numberLines = initialConfig.length;

initialConfig = initialConfig.map((line) => line.split(''));

///renvoie le numéro du caractère qu'il faut chercher sur
///chaque ligne pour créer les piles (on itère sur la dernière ligne)
const findIndex = (pileNumber) => {
  for (var i = 0; i < initialConfig[numberLines - 1].length - 1; i++) {
    if (pileNumber === parseInt(initialConfig[numberLines - 1][i])) {
      return (indexPile = i);
    }
  }
};

///renvoie la pile au numéro désigné
const initStack = (pileNumber) => {
  indexPile = findIndex(pileNumber);
  stack = [];

  for (var i = numberLines - 2; i >= 0; i--) {
    if (initialConfig[i][indexPile] === ' ') {
      return stack;
    } else {
      stack.push(initialConfig[i][indexPile]);
    }
  }
  return stack;
};

///Création des piles
cratesHolder = [];
for (var i = 1; i < 10; i++) {
  pile = new Pile([]);
  pile.set(initStack(i));
  cratesHolder.push(pile);
}

///instruction parser
const instructionParsing = (instruction) => {
  tab = instruction.split(' ');
  return { quantity: tab[1], origin: tab[3], destination: tab[5] };
};

instructionText = data[1].split('\n');
instructions = [];

for (var i = 0; i < instructionText.length; i++) {
  instructions.push(instructionParsing(instructionText[i]));
}

const crateMovement = (origin, destination) => {
  crate = cratesHolder[origin - 1].stack.pop();
  cratesHolder[destination - 1].stack.push(crate);
};

for (var i = 0; i < instructions.length; i++) {
  for (var j = 0; j < instructions[i].quantity; j++) {
    crateMovement(instructions[i].origin, instructions[i].destination);
  }
}

s = '';
for (var i = 0; i < cratesHolder.length; i++) {
  s += cratesHolder[i].stack[cratesHolder[i].stack.length - 1];
}

console.log(s);
