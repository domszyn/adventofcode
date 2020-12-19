const fs = require("fs");

const program = fs.readFileSync(__dirname + "/input.txt", "utf-8");

const run = (program, terminate) => {
  const instructions = program
    .split("\n")
    .map((_) => _.split(" "))
    .map(([instruction, value]) => ({ instruction, value: +value }));

  let acc = 0;
  const executed = new Set();

  const execute = (address) => {
    if (address > instructions.length) {
      return acc;
    }

    const { instruction, value } = instructions[address - 1];

    if (executed.has(address)) {
      return terminate ? acc : undefined;
    }

    executed.add(address);
    switch (instruction) {
      case "nop":
        return execute(address + 1);
      case "acc":
        acc += +value;
        return execute(address + 1);
      case "jmp":
        return execute(address + +value);
      default:
        console.log("Unknown instruction", instruction);
        throw new Error();
    }
  };

  return execute(1);
};

console.log("Part1: ", run(program, true));

const replaceInstruction = (search, replace) => {
  let returnAcc;
  let startIndex = 0;
  do {
    startIndex = program.indexOf(search, startIndex + 1);
    if (startIndex < 0) {
      break;
    }
    const patchedProgram =
      program.substring(0, startIndex) +
      replace +
      program.substring(startIndex + 3);
    returnAcc = run(patchedProgram);
  } while (returnAcc === undefined && startIndex < program.length);

  return returnAcc;
};

console.log(
  "Part2: ",
  replaceInstruction("nop", "jmp") || replaceInstruction("jmp", "nop")
);
