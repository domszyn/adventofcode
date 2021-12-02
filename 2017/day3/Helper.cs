public static class Spiral
{
    public static int SumAdjacent(int[,] grid, int x, int y)
    {
        var sum = 0;
        for (int i = Math.Max(x - 1, 0); i < Math.Min(grid.Length, x + 2); i++)
        {
            for (int j = Math.Max(y - 1, 0); j < Math.Min(grid.Length, y + 2); j++)
            {
                sum += grid[j, i];
            }
        }
        return sum;
    }
}