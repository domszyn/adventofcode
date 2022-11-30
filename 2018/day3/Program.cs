var fabric = new HashSet<int>[1000][];

for (var i = 0; i < fabric.Length; i++)
{
    fabric[i] = new HashSet<int>[fabric.Length];
    for (var j = 0; j < fabric.Length; j++)
    {
        fabric[i][j] = new HashSet<int>();
    }
}

var claims = Input.Claims;

foreach (var claim in claims)
{
    for (var i = claim.X; i < claim.X + claim.Width; i++)
    {
        for (var j = claim.Y; j < claim.Y + claim.Height; j++)
        {
            fabric[j][i].Add(claim.ID);
        }
    }
}

Console.WriteLine("Part 1: {0}", fabric.SelectMany(line => line).Count(_ => _.Count > 1));
var overlappingClaimIDs = fabric
    .SelectMany(line => line)
    .Where(_ => _.Count > 1)
    .SelectMany(_ => _.AsEnumerable())
    .Distinct();
Console.WriteLine("Part 2: {0}", claims.Select(_ => _.ID).Except(overlappingClaimIDs).Single());