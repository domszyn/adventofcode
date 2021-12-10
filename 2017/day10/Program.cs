using System.Text;

var knotHash = (byte[] lengths, int rounds) =>
{
    var numbers = Enumerable.Range(0, 256).Select(_ => (byte)_).ToList();
    var currentPosition = 0;
    var skipSize = 0;

    for (var round = 0; round < rounds; round++)
    {
        foreach (var length in lengths)
        {
            var elements = new List<byte>();
            for (var i = 0; i < length; i++)
            {
                elements.Add(numbers[(i + currentPosition) % numbers.Count]);
            }
            elements.Reverse();
            for (var i = 0; i < length; i++)
            {
                numbers[(i + currentPosition) % numbers.Count] = elements[i];
            }
            currentPosition += length + skipSize++;
            currentPosition %= numbers.Count;
        }
    }

    return numbers;
};


var part1 = knotHash(new byte[] { 106, 118, 236, 1, 130, 0, 235, 254, 59, 205, 2, 87, 129, 25, 255, 118 }, 1);
Console.WriteLine("Part 1: {0}", part1[0] * part1[1]);

var input = Encoding.ASCII.GetBytes("106,118,236,1,130,0,235,254,59,205,2,87,129,25,255,118").ToList();
input.AddRange(new byte[] { 17, 31, 73, 47, 23 });

var sparseHash = knotHash(input.ToArray(), 64);
var denseHash = new byte[16];
for (var i = 0; i < 16; i++)
{
    var numbers = sparseHash.Skip(i * 16).Take(16).ToArray();
    for (var j = 0; j < 16; j++)
    {
        denseHash[i] ^= numbers[j];
    }
}

Console.WriteLine("Part 2: {0}", string.Join("", denseHash.Select(b => b.ToString("x2"))));
