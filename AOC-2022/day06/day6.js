fs = require('fs');
var data = fs.readFileSync('./input.txt', 'utf8');

const isMarker = (scope) => {
  const result = [...new Set(scope)];
  if (result.length === scope.length) {
    return true;
  }
  return false;
};

const compare = (origin, other1, other2, other3) => {
  if ((origin === other1) | (origin === other2) | (origin === other3)) {
    return true;
  } else {
    return false;
  }
};

scope = [data[0], data[1], data[2], data[3]];

for (var index = 4; index < data.length - 4; index++) {
  for (var j = 0; j < 14; j++) {
    scope[j] = data[index - (13 - j)];
  }

  console.log(scope);
  if (isMarker(scope)) {
    console.log(index + 1);
    break;
  }
}
