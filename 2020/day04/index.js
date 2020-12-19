const fs = require("fs");

const scanPassport = (passport) =>
  passport.split(" ").reduce((pass, field) => {
    const [key, value] = field.split(":");
    return {
      ...pass,
      [key]: value,
    };
  }, {});

const passports = fs
  .readFileSync(__dirname + "/input.txt", "utf-8")
  .split("\n\n")
  .map((_) => _.replaceAll("\n", " "))
  .map(scanPassport);

const isValidByr = (byr, simpleCheck) =>
  !isNaN(+byr) && (simpleCheck || (+byr >= 1920 && +byr <= 2002));

const isValidIyr = (iyr, simpleCheck) =>
  !isNaN(+iyr) && (simpleCheck || (+iyr >= 2010 && +iyr <= 2020));

const isValidEyr = (eyr, simpleCheck) =>
  !isNaN(+eyr) && (simpleCheck || (+eyr >= 2020 && +eyr <= 2030));

const isValidHgt = (hgt, simpleCheck) => {
  if (simpleCheck || !hgt) return !!hgt;
  const height = +hgt.substr(0, hgt.length - 2);
  const unit = hgt.substr(hgt.length - 2);
  return (
    (unit === "cm" && height >= 150 && height <= 193) ||
    (unit === "in" && height >= 59 && height <= 76)
  );
};

const isValidHcl = (hcl, simpleCheck) =>
  !!hcl && (simpleCheck || /^#[0-9a-f]{6}$/.test(hcl));

const isValidEcl = (ecl, simpleCheck) =>
  !!ecl && (simpleCheck || /^amb|blu|brn|gry|grn|hzl|oth$/.test(ecl));

const isValidPid = (pid, simpleCheck) =>
  !!pid && (simpleCheck || /^\d{9}$/.test(pid));

const isValid = ({ byr, iyr, eyr, hgt, hcl, ecl, pid }, simpleCheck = true) =>
  isValidByr(byr, simpleCheck) &&
  isValidIyr(iyr, simpleCheck) &&
  isValidEyr(eyr, simpleCheck) &&
  isValidHgt(hgt, simpleCheck) &&
  isValidHcl(hcl, simpleCheck) &&
  isValidEcl(ecl, simpleCheck) &&
  isValidPid(pid, simpleCheck);

console.log("Part1: ", passports.filter(isValid).length);
console.log("Part2: ", passports.filter((p) => isValid(p, false)).length);
