fs = require('fs');
// var data = fs.readFileSync('./AOC-2022/day07/input.txt', 'utf8');
var data = fs.readFileSync('./input.txt', 'utf8');
// var data = fs.readFileSync('./AOC-2022/day07/exemple.txt', 'utf8');
// var data = fs.readFileSync('./exemple.txt', 'utf8');

///CLASS
class Directory {
  constructor(dirName, Tree = []) {
    this.name = dirName;
    this.directories = [];
    this.files = [];
    this.parentTree = Tree;
  }
  addDirectory(dirName, dirIndex) {
    this.directories.push({ dirName: dirName, dirIndex: dirIndex });
  }
  addFile(size, fileName) {
    var file = { size: size, fileName: fileName };
    this.files.push(file);
  }
}

//////////////////////////////////////commande////////////////////////////////////////
const commande = (cmd) => {
  if (cmd.slice(0, 2) == 'ls') {
    // console.log('ls command');
    return false;
  }
  if (cmd.slice(0, 2) == 'cd') {
    if (cmd.slice(3, 5) == '..') {
      if (currentDirectory.parentTree.length <= 1) {
        return false;
      }
      console.log('\\\\');
      currentDirectory.parentTree.pop();
      currentDirectory.dirIndex = currentDirectory.parentTree[currentDirectory.parentTree.length - 1].index;
    } else {
      //cas de base: $ cd 'nom_dossier'
      var dirToFind = cmd.slice(3, cmd.length);

      //   console.log('//' + dirToFind);

      //   var dirIndex = directories.findIndex((value) => {
      //     return value.name === dirToFind && JSON.stringify(value.parentTree) === JSON.stringify(value.parentTree);
      //   });

      if (dirToFind === '/') {
        var dirIndex = directories.findIndex((value) => {
          return value.name === dirToFind;
        });
      } else {
        var indexes = [],
          i;
        for (i = 0; i < directories.length; i++) {
          //   console.log(directories[i].parentTree);
          //   console.log(currentDirectory.parentTree);
          //   console.log(JSON.stringify(directories[i].parentTree) === JSON.stringify(currentDirectory.parentTree));
          if (directories[i].name === dirToFind && JSON.stringify(directories[i].parentTree) === JSON.stringify(currentDirectory.parentTree)) {
            indexes.push(i);
          }
        }
        dirIndex = indexes[0];
        if (indexes.length > 1) {
          console.log('MORE THAN ONE');
        }
      }

      //   var dirIndex = directories.findIndex((value) => {
      //     return value.name === dirToFind;
      //   });
      if (dirIndex === -1) {
        console.log('__________ERROR : unknown file named: "' + dirToFind + '"___________');
        return false;
      }
      currentDirectory.dirIndex = dirIndex;
      currentDirectory.parentTree.push({ name: dirToFind, index: dirIndex });
    }
  } else {
    console.log('__________undefined command: "' + cmd + '"__________');
  }
};

////////////////////////////new dir / file//////////////////////////////////////////
const newDir = (dirName) => {
  //ajout du répertoire dans la liste des répertoires du répertoire parent
  directories[currentDirectory.dirIndex].addDirectory(dirName, directories.length);

  //ajout du répertoire à la liste de tous les répertoires
  directories.push(new Directory(dirName, [...currentDirectory.parentTree]));
};

const newFile = (fileSize, fileName) => {
  directories[currentDirectory.dirIndex].addFile(fileSize, fileName);
};

/////////////////////////////////parser//////////////////////////////////////////////
const lineParser = (ligne) => {
  if (ligne[0] === '$') {
    return commande(ligne.slice(2, ligne.length));
  }
  if (ligne[0] === 'd') {
    // console.log(ligne);
    return newDir(ligne.slice(4, ligne.length));
  }
  //   if (typeof parseInt(ligne[0]) === 'number') {
  if (Number.isInteger(parseInt(ligne[0]))) {
    // console.log(ligne);
    var fileInfo = ligne.split(' ');
    return newFile(fileInfo[0], fileInfo[1]);
  }
  console.log('_____________ERROR : unauthorized line: "' + ligne + '"_____________');
  return false;
};

/////////////////////////////tableaux globaux/////////////////////////////////////////
/////
//Tableau contenant le chemin actuel
var currentDirectory = { dirIndex: 0, parentTree: [] };

//Tableau contenant tous les dossiers [tableau contenant des objets Directory]
directories = [new Directory('/', [])];
//////
//////

/////DATA PROCESSING/////
// text = '$ cd /\ndir gg\n$ cd gg\ndir toto\ndir nouv\n$ cd toto\ndir toy\n1254 gea\ncd ..\ncd ..\n';
// text = '$ cd /\ndir frgg\ndir aag\n293559 jztrccm.hvd\n559 jztd\n$ cd frgg\n1475 aaab\n$ cd ..\n$ cd aag\n400 dag';
text = data;
//  /\n$ ls\ndir bfqzjjct\n293559 jztrccm.hvd';
data = text.split('\n');
/////

/////PARSING
for (var ligne = 0; ligne < data.length; ligne++) {
  //   console.log(data[ligne]);
  var parsing = lineParser(data[ligne]);
  //   console.log('_____________');

  //   console.log(directories.map((dir) => dir.files));
}

console.log('_____________');
// console.log(currentDirectory);

////
////Size count:
////

const sizeOfDirectoryFiles = (directory) => {
  sizes = directory.files.map((file) => file.size);
  totalSize = sizes.reduce((acc, cur) => acc + parseInt(cur), 0);
  return totalSize;
};

const sizeOfDirectory = (directory, totalSize = 0) => {
  var sizeOfFiles = sizeOfDirectoryFiles(directory);

  if (directory.directories.length === 0) {
    return sizeOfFiles;
  } else {
    var sizeOfDirectories = 0;
    for (var i = 0; i < directory.directories.length; i++) {
      //   console.log(directory.directories);

      const a = sizeOfDirectory(directories[directory.directories[i].dirIndex]);

      //   console.log(a);
      sizeOfDirectories += a;
      //   console.log(sizeOfDirectories);
    }
    return sizeOfDirectories + sizeOfFiles;
  }
};

// console.log(sizeOfDirectory(directories[0]));

var agregator = 0;
for (var i = 0; i < directories.length; i++) {
  var size = sizeOfDirectory(directories[i]);
  if (size <= 100000) {
    // console.log(directories[i]);
    // console.log('/____________/');
    agregator += size;
  }
}
console.log(agregator);
