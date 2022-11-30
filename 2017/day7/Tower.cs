public class Tower
{
    private List<Tower> disc = new List<Tower>();
    public string Name { get; set; } = "";
    public int OwnWeight { get; set; } = 0;
    public IEnumerable<Tower> Disc => disc.AsReadOnly();

    public int TotalWeight => OwnWeight + disc.Select(t => t.TotalWeight).Sum();

    public Tower? UnbalancedSubTower()
    {
        if (disc.Count < 2)
        {
            return null;
        }

        var outliers = disc.Skip(1).Where( t => t.TotalWeight != disc[0].TotalWeight);

        return outliers.Count() > 1 ? disc[0] : outliers.SingleOrDefault();
    }

    public void AddTower(Tower tower)
    {
        disc.Add(tower);
    }
}