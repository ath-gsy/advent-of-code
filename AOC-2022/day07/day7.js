fs = require('fs');
var data = fs.readFileSync('./input.txt', 'utf8');

///CLASS
class Directory {
  constructor(dirName, Tree = []) {
    this.name = dirName;
    this.directories = [];
    this.files = [];
    this.parentTree = Tree;
  }
  addDirectory(dirName) {
    this.directories.push(new Directory(dirName));
  }
  addFile(size, fileName) {
    var file = { size: size, fileName: fileName };
    this.files.push(file);
  }
}

///////////

// var currentDirectory = { parentTree: [] };
var currentDirectory = { dirIndex: 0, parentTree: [] };

const commande = (cmd) => {
  if (cmd.slice(0, 2) == 'cd') {
    if (cmd.slice(3, 5) == '..') {
      if (currentDirectory.parentTree.length === 0) {
        return false;
      }
      currentDirectory.parentTree.pop();
      currentDirectory.dirIndex = currentDirectory.parentTree[currentDirectory.parentTree.length - 1].index;
    } else {
      dirToFind = cmd.slice(3, cmd.length);

      var dirIndex = directories.findIndex((value) => {
        return value.name === dirToFind;
      });
      currentDirectory.dirIndex = dirIndex;
      currentDirectory.parentTree.push({ name: dirToFind, index: dirIndex });
    }
  } else {
  }
};

const newDir = (dirName) => {
  directories[currentDirectory.dirIndex].addDirectory(dirName);

  directories.push(new Directory(dirName, [...currentDirectory.parentTree]));
};

const newFile = (fileSize, fileName) => {
  directories[currentDirectory.dirIndex].addFile(fileSize, fileName);
};

const lineParser = (ligne) => {
  if (ligne[0] === '$') {
    return commande(ligne.slice(2, ligne.length));
  }
  if (ligne[0] === 'd') {
    return newDir(ligne.slice(4, ligne.length));
  }
  if (typeof parseInt(ligne[0]) === 'number') {
    var fileInfo = ligne.split(' ');
    return newFile(fileInfo[0], fileInfo[1]);
  }
  console.log('_____________ERROR_____________');
  return false;
};

//////
//////
directories = [new Directory('/')];
//////
//////

text = '$ cd /\ndir gg\n$ cd gg\ndir toto\n$ cd toto\ndir toy\n1254 gea';
//  /\n$ ls\ndir bfqzjjct\n293559 jztrccm.hvd';
data = text.split('\n');

for (var ligne = 0; ligne < data.length; ligne++) {
  //   console.log(data[ligne]);
  var parsing = lineParser(data[ligne]);
  console.log('_____________');
  console.log(directories.map((dir) => dir.files));
  console.log(currentDirectory);
}
