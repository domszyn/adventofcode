var getBanks = () => new int[] { 14, 0, 15, 12, 11, 11, 3, 5, 1, 6, 8, 4, 9, 1, 8, 4 };
var getHash = (int[] banks) => string.Join("", banks.Select(i => i.ToString("X")));
var banks = getBanks();

var allocations = new HashSet<string>();
allocations.Add(getHash(banks));
var part1 = "";
for (var r = 1; ; r++)
{
    var max = banks.Max();
    var idx = Array.FindIndex(banks, _ => _ == max);

    banks[idx] = 0;
    for (var i = 0; i < max; i++)
    {
        banks[(idx + i + 1) % banks.Length]++;
    }

    var hash = getHash(banks);
    if (allocations.Contains(hash) && part1 == "")
    {
        part1 = hash;
        Console.WriteLine("Part 1: {0}", r);
        r = 0;
    }
    else if (hash == part1)
    {
        Console.WriteLine("Part 2: {0}", r);
        break;
    }
    else
    {
        allocations.Add(hash);
    }
}
