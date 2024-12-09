import { makeArray } from "../utils.js";
import { input } from "./input.js";
import '../array.js';
const lengths = [...input].map(Number);

const readDisk = (arr) => {
    const disk = makeArray(lengths.sum());
    const index = new Map();
    for (let i = 0, j = 0, id = 0; i < arr.length; i++) {
        if (i % 2 === 1) {
            j += lengths[i];
            continue
        } else {
            index.set(id, j);
            for (let k = j; k < j + lengths[i]; k++) {
                disk[k] = id;
            }
            id++;
            j += lengths[i]
        }

    }
    return [disk, index];
}

const checksum = (arr) => arr.map((id, idx) => {
    if (id === undefined) return 0;
    return id * idx;
}).sum();

const [disk] = readDisk(lengths);

for (let i = disk.length - 1; i >= 0; i--) {
    if (disk[i] == undefined) {
        continue;
    }

    const eidx = disk.indexOf(undefined);
    if (eidx > i) {
        break;
    }

    disk[eidx] = disk[i];
    disk[i] = undefined;
}

console.log('Part 1', checksum(disk));

const findEmptyBlock = (arr, len, end) => {
    for (let i = 0; i < end;) {
        let idx = arr.indexOf(undefined, i);

        let allEmpty = true;
        for (let j = idx; j < idx + len; j++) {
            if (arr[j] !== undefined) {
                i = j;
                allEmpty = false;
                break;
            }
        }

        if (allEmpty) {
            return idx;
        }
    }

    return -1;
}

const [disk2, diskIndex] = readDisk(lengths);
const unmovableSizes = new Set();

for (let i = input.length - 1; i > 0; i -= 2) {
    let id = i / 2;
    let pos = diskIndex.get(id);
    let fileSize = +input[i];

    if (unmovableSizes.has(fileSize)) {
        continue
    };
    
    let emptyBlock = findEmptyBlock(disk2, +input[i], pos);
    if (emptyBlock > 0 && emptyBlock < pos) {
        for (let j = 0; j < input[i]; j++) {
            disk2[emptyBlock + j] = id;
            disk2[pos + j] = undefined;
        }

        diskIndex.set(id, emptyBlock);
    } else {
        unmovableSizes.add(fileSize);
    }
}

console.log('Part 2', checksum(disk2));
