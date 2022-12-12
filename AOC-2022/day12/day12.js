// Implémentation de l'algorithme de Dijkstra : https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
// A partir de sa decription (ie sans copier du code)

fs = require('fs');
// var data = fs.readFileSync('./AOC-2022/day12/input.txt', 'utf8');
const rawData = fs.readFileSync('./input.txt', 'utf8');

////
////
////CLASS

class Point {
  constructor(x, y) {
    this.x = x;
    this.y = y;
  }
}

class NodeTable {
  constructor(array) {
    this.board = array;
    this.columnLimit = array.length - 1;
    this.rowLimit = array[0].length - 1;
  }

  setPoint(point, value) {
    this.board[point.x][point.y] = value;
  }

  findVoisins(point) {
    value = this.board[point.x][point.y];

    if (this.board[point.x])
      //source : stack exchange
      for (let x = Math.max(0, point.x - 1); x <= Math.min(point.x + 1, rowLimit); x++) {
        for (let y = Math.max(0, point.y - 1); y <= Math.min(point.y + 1, columnLimit); y++) {
          if (x !== point.x || y !== point.y) {
            console.log(myArray[x][y]);
            sum += myArray[x][y];
          }
        }
      }
  }
}

var data = rawData.split('\n');

////
////Initialisation
////

const heightTable = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'];

///Création du tableau contenant les nodes (la "map")
const height = data.length;
const length = data[0].length;

var nodes = data.slice(0).map((line) => new Array(length).fill(-1));
nodes = new NodeTable(nodes);

const S = new Point(20, 0);
const E = new Point(20, 43);
nodes.setPoint(S, 0);
nodes.setPoint(E, 0);

data = data.map((line) => line.split(''));
data[S.x][S.y] = 'a';
data[E.x][E.y] = 'z';

currentNode = new Point(S.x, S.y);
