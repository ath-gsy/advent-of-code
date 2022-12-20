// Implémentation de l'algorithme de Dijkstra : https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
// A partir de sa decription (ie sans copier du code ou pseudo-code)

fs = require('fs');
// const rawData = fs.readFileSync('./AOC-2022/day12/input.txt', 'utf8');
const rawData = fs.readFileSync('./input.txt', 'utf8');

////
////
////CLASS

class Point {
  constructor(x, y) {
    this.x = x;
    this.y = y;
  }

  //moving methods
  up() {
    if (this.x - 1 > 0) {
      this.x -= 1;
    } else {
      console.log('upper bound reached');
    }
  }

  down() {
    if (this.x + 1 < height) {
      this.x += 1;
    } else {
      console.log('lower bound reached');
    }
  }

  left() {
    if (this.y - 1 > 0) {
      this.y -= 1;
    } else {
      console.log('left bound reached');
    }
  }

  right() {
    if (this.y + 1 < length) {
      this.y += 1;
    } else {
      console.log('right bound reached');
    }
  }
}

class NodeTable {
  constructor(array) {
    this.board = array;
    this.rowLimit = array.length - 1;
    this.columnLimit = array[0].length - 1;
  }

  setPoint(point, value) {
    if ((point.x >= 0) & (point.x <= this.rowLimit) & (point.y >= 0) & (point.y <= this.columnLimit)) {
      this.board[point.x][point.y] = value;
    } else {
      console.log("point can't be set : out of bound");
    }
  }

  findVoisins(point) {
    var currentValue = data[point.x][point.y].charCodeAt(0);
    var voisins = [];

    if (point.x != this.rowLimit && data[point.x + 1][point.y].charCodeAt(0) <= currentValue + 1 && this.board[point.x + 1][point.y] == 10000) {
      voisins.push({ diffValue: data[point.x + 1][point.y].charCodeAt(0) - currentValue, x: point.x + 1, y: point.y });
      // console.log(getPoint(voisins[voisins.length - 1]));
    }
    if (point.x != 0 && data[point.x - 1][point.y].charCodeAt(0) <= currentValue + 1 && this.board[point.x - 1][point.y] == 10000) {
      voisins.push({ diffValue: data[point.x - 1][point.y].charCodeAt(0) - currentValue, x: point.x - 1, y: point.y });
      // console.log(getPoint(voisins[voisins.length - 1]));
    }
    if (point.y != this.columnLimit && data[point.x][point.y + 1].charCodeAt(0) <= currentValue + 1 && this.board[point.x][point.y + 1] == 10000) {
      voisins.push({ diffValue: data[point.x][point.y + 1].charCodeAt(0) - currentValue, x: point.x, y: point.y + 1 });
      // console.log(getPoint(voisins[voisins.length - 1]));
    }
    if (point.y != 0 && data[point.x][point.y - 1].charCodeAt(0) <= currentValue + 1 && this.board[point.x][point.y - 1] == 10000) {
      voisins.push({ diffValue: data[point.x][point.y - 1].charCodeAt(0) - currentValue, x: point.x, y: point.y - 1 });
      // console.log(getPoint(voisins[voisins.length - 1]));
    }

    voisins.sort(function (a, b) {
      return b.diffValue - a.diffValue;
    });

    // if (point.x != this.rowLimit) {
    //   voisins.push(data[point.x - 1][point.y]);
    // }
    // if (point.x != 0) {
    //   voisins.push(data[point.x + 1][point.y]);
    // }
    // if (point.y != this.columnLimit) {
    //   voisins.push(data[point.x][point.y + 1]);
    // }
    // if (point.y != 0) {
    //   voisins.push(data[point.x][point.y - 1]);
    // }

    return voisins;
  }

  nodesPrinter() {
    var totalString = '';
    for (var i = 0; i < this.board.length; i++) {
      totalString += this.board[i].join('') + '\n';
    }
    fs.writeFile('./nodesTable.txt', totalString, (err) => {
      if (err) {
        console.error(err);
      }
    });
  }
}

var data = rawData.split('\n');

////
////Initialisation
////

// const heightTable = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'];

///Création du tableau contenant les nodes (la "map")
const height = data.length - 1;
const length = data[0].length - 1;

var nodes = data.slice(0).map((line) => new Array(length + 1).fill(10000));
nodes = new NodeTable(nodes);

const S = new Point(20, 0);
const E = new Point(20, 43);
nodes.setPoint(S, 0);
nodes.setPoint(E, 0);

nodes.nodesPrinter();

data = data.map((line) => line.split(''));
data[S.x][S.y] = 'a';
data[E.x][E.y] = 'z';

///
const getPoint = (point) => {
  return data[point.x][point.y];
};
///

////
///////////////////
////

currentNode = new Point(S.x, S.y);

const test = new Point(5, 2);

// console.log(getPoint(test));

console.log(nodes.findVoisins(test));

// const stepAction = (point) => {

// }

var tries = new Array(500);

var stepList = new Array(100);
const pathTry = (stepList) => {};

// for (var voisinsIndex = 0; voisinsIndex < voisins.length; voisinsIndex++) {
//   return algoBestPath(voisins[voisinsIndex], currentDistance);
// }

// if () {

// } else {
//   for (var voisinsIndex = 0; voisinsIndex < voisins.length; voisinsIndex++) {
//     return algoBestPath(voisins[voisinsIndex], currentDistance);
//   }
// }
