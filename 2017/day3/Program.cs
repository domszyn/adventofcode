
const int Input = 347991;

var location = 1;
for (int i = 2; i < 10000; i++)
{
    location += 8 * (i - 1);

    if (location >= Input)
    {
        var sl = i * 2 - 2;
        var corners = new[] { location, location - sl, location - 2 * sl, location - 3 * sl };
        Console.WriteLine("Part 1: {0}", sl + Input - corners.Last(c => c > Input));
        break;
    }
}

var matrix = new int[11, 11];
int x = 5, y = 5;
matrix[y, x] = 1;

for (int i = 1; i < 5; i++)
{
    do
    {
        x++;
        matrix[y, x] = Spiral.SumAdjacent(matrix, x, y);
        if (matrix[y, x] > Input) {
            Console.WriteLine("Part 2: {0}", matrix[y, x]);
            return;
        }
    } while (matrix[y - 1, x] != 0);

    do
    {
        y--;
        matrix[y, x] = Spiral.SumAdjacent(matrix, x, y);
        if (matrix[y, x] > Input) {
            Console.WriteLine("Part 2: {0}", matrix[y, x]);
            return;
        }
    } while (y > 0 && matrix[y, x - 1] != 0);

    do
    {
        x--;
        matrix[y, x] = Spiral.SumAdjacent(matrix, x, y);
        if (matrix[y, x] > Input) {
            Console.WriteLine("Part 2: {0}", matrix[y, x]);
            return;
        }
    } while (matrix[y + 1, x] != 0);

    do
    {
        y++;
        matrix[y, x] = Spiral.SumAdjacent(matrix, x, y);
        if (matrix[y, x] > Input) {
            Console.WriteLine("Part 2: {0}", matrix[y, x]);
            return;
        }
    } while (matrix[y, x + 1] != 0);
}