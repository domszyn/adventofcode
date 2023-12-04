const input = ``;

const rows = input.split('\n');
const matrix = rows.map(s => s.split(''));
const rowLength = rows[0].length;

var sumOfPartNumbers = 0;
const numbers = new Set('0123456789.');

var gears = new Map();

function addGear(key, numberStr) {
    if (!gears.has(key)) {
        gears.set(key, []);
    }

    gears.set(key, [
        ...gears.get(key),
        parseInt(numberStr)
    ]);
}

function isPartNumber(numberStr, rowIndex, startIndex) {
    if (rowIndex > 0) {
        for (let i = startIndex - 1; i <= startIndex + numberStr.length; i++) {
            if (i < 0 || i >= rowLength) {
                continue
            }

            const symbol = matrix[rowIndex - 1][i];
            if (!numbers.has(symbol)) {
                if (symbol == '*') {
                    addGear((rowIndex - 1).toString() + ',' + i, numberStr)
                }
                return true;
            }
        }
    }

    if (startIndex > 0 && !numbers.has(matrix[rowIndex][startIndex - 1])) {
        if (matrix[rowIndex][startIndex - 1] == '*') {
            addGear(rowIndex + ',' + (startIndex - 1).toString(), numberStr)
        }
        return true;
    }

    if (startIndex + numberStr.length < rowLength && !numbers.has(matrix[rowIndex][startIndex + numberStr.length])) {
        if (matrix[rowIndex][startIndex + numberStr.length] == '*') {
            addGear(rowIndex + ',' + (startIndex + numberStr.length).toString(), numberStr);
        }
        return true;
    }

    if (rowIndex + 1 < rows.length) {
        for (let i = startIndex - 1; i <= startIndex + numberStr.length; i++) {
            if (i < 0 || i >= rowLength) {
                continue
            }

            const symbol = matrix[rowIndex + 1][i];
            if (!numbers.has(matrix[rowIndex + 1][i])) {
                if (symbol == '*') {
                    addGear((rowIndex + 1).toString() + ',' + i, numberStr);
                }

                return true;
            }
        }
    }

    return false;
}

function isDigit(char) {
    return char && char >= '0' && char <= '9';
}

for (let index = 0; index < rows.length; index++) {
    const row = rows[index];
    for (const match of row.matchAll(/\d+/g)) {
        var numberStr = match[0];
        if (isPartNumber(numberStr, index, match.index)) {
            sumOfPartNumbers += parseInt(numberStr)
        }
    }
}

var sumOfGearRatios = 0;

for (const g of gears.values()) {
    if (g.length == 2) {
        sumOfGearRatios += g[0]*g[1];
    }
}

console.log("Part 1", sumOfPartNumbers);
console.log("Part 2", sumOfGearRatios);