let tape = new Map();
let cursor = 0;
let state = 'A';

for (let i = 0; i < 12794428; i++) {
    if (!tape.has(cursor)) {
        tape.set(cursor, 0);
    }
    if (state == 'A') {
        if (tape.get(cursor) == 0) {
            tape.set(cursor, 1);
            cursor++;
            state = 'B';
        } else {
            tape.set(cursor, 0);
            cursor--;
            state = 'F';
        }
    } else if (state == 'B') {
        if (tape.get(cursor) == 0) {
            tape.set(cursor, 0);
            cursor++;
            state = 'C';
        } else {
            tape.set(cursor, 0);
            cursor++;
            state = 'D';
        }
    } else if (state == 'C') {
        if (tape.get(cursor) == 0) {
            tape.set(cursor, 1);
            cursor--;
            state = 'D';
        } else {
            tape.set(cursor, 1);
            cursor++;
            state = 'E';
        }
    } else if (state == 'D') {
        if (tape.get(cursor) == 0) {
            tape.set(cursor, 0);
            cursor--;
            state = 'E';
        } else {
            tape.set(cursor, 0);
            cursor--;
            state = 'D';
        }
    } else if (state == 'E') {
        if (tape.get(cursor) == 0) {
            tape.set(cursor, 0);
            cursor++;
            state = 'A';
        } else {
            tape.set(cursor, 1);
            cursor++;
            state = 'C';
        }
    } else if (state == 'F') {
        if (tape.get(cursor) == 0) {
            tape.set(cursor, 1);
            cursor--;
            state = 'A';
        } else {
            tape.set(cursor, 1);
            cursor++;
            state = 'A';
        }
    }
}

console.log([...tape.values()].filter(v => v == 1).length);