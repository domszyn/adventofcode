import '../array.js';
import { almanac } from "./input.js";

function parseAlmanac() {
    const [seedLine, ...rest] = almanac.split('\n');
    const seeds = seedLine.split(" ").slice(1).map(s => parseInt(s));
    let maps = [];

    for (let i = 0; i < rest.length; i++) {
        const line = rest[i];
        if (line == "") {
            maps.push([]);
            i++;
            continue;
        } else {
            const [destinationStart, sourceStart, length] = line.split(' ').map(n => parseInt(n));
            maps[maps.length - 1].push({
                sourceStart,
                destinationStart,
                length
            });
        }
    }

    maps = maps.map(m => m.sort((a, b) => a.sourceStart - b.sourceStart));

    return [seeds, maps];
}

const [seeds, maps] = parseAlmanac();

function translate(map, source) {
    const mapping = map.find(({ sourceStart, length }) => source >= sourceStart && source < sourceStart + length);
    return mapping ? mapping.destinationStart - mapping.sourceStart + source : source;
}

const part1 = seeds.map(seed => maps.reduce((s, map) => translate(map, s), seed)).min();
console.log("Part 1", part1);

function translateRange(map, sourceStart, length) {
    const mapping = map.find(m => sourceStart >= m.sourceStart && sourceStart < m.sourceStart + m.length);
    if (!mapping) {
        return [{ sourceStart, length }];
    }

    let ranges = [{
        sourceStart: mapping.destinationStart + sourceStart - mapping.sourceStart,
        length: Math.min(sourceStart + length, mapping.sourceStart + mapping.length) - sourceStart
    }];

    if (sourceStart + length > mapping.sourceStart + mapping.length) {
        ranges = [
            ...ranges,
            ...translateRange(
                map,
                mapping.sourceStart + mapping.length,
                sourceStart + length - mapping.sourceStart - mapping.length
            )
        ];
    }

    return ranges;
}

let input = [];
for (let i = 0; i < seeds.length; i += 2) {
    input.push({
        destinationStart: 0,
        sourceStart: seeds[i],
        length: seeds[i + 1]
    });
}
input.sort((a, b) => a.sourceStart - b.sourceStart);

const minLocation = maps.reduce(
    (input, output) => 
        input.map(({ sourceStart, length }) => translateRange(output, sourceStart, length)).flatten(),
    input
).map(({ sourceStart }) => sourceStart).min();

console.log("Part 2", minLocation);
