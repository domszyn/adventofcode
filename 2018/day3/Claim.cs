using System.Text.RegularExpressions;
public struct Claim
{
    public int ID { get; set; } = 0;
    public int X { get; set; } = 0;
    public int Y { get; set; } = 0;
    public int Width { get; set; } = 0;
    public int Height { get; set; } = 0;

    public Claim(string claimDescription)
    {
        var rx = new Regex(@"^#(?<id>\d+) @ (?<x>\d+),(?<y>\d+): (?<w>\d+)x(?<h>\d+)$");
        var matches = rx.Matches(claimDescription);
        if (matches.Count == 1)
        {
            ID = Convert.ToInt32(matches[0].Groups["id"].Value);
            X = Convert.ToInt32(matches[0].Groups["x"].Value);
            Y = Convert.ToInt32(matches[0].Groups["y"].Value);
            Width = Convert.ToInt32(matches[0].Groups["w"].Value);
            Height = Convert.ToInt32(matches[0].Groups["h"].Value);
        }
    }
}