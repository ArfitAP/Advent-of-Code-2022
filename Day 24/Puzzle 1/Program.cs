using System.Linq;
using System.Runtime.CompilerServices;
using System.Xml.Linq;

namespace Puzzle_1
{
    internal class Program
    {

        static void Main(string[] args)
        {
            string line;

            List<Blizzard> blizzards = new();
             
            var startPosition = (0, 0);
            var endPosition = (0, 0);
            int row = -1;
            int col = 0;
            StreamReader reader;
            reader = new StreamReader(@"../../../../input.txt");
            do
            {
                line = reader.ReadLine()!;
                row++;

                if (row == 0)
                {
                    col = line.Length;
                    for (int i = 0; i < line.Length; i++)
                    {
                        if (line[i] == '.')
                        {
                            startPosition = (i, 0);
                        }
                    }
                    continue;
                }
                else if(line[1] == '#' || line[2] == '#')
                {
                    for (int i = 0; i < line.Length; i++)
                    {
                        if (line[i] == '.')
                        {
                            endPosition = (i, row);
                        }
                    }
                    continue;
                }

                for (int i = 0; i < line.Length; i++)
                {
                    if (line[i] == '<')
                    {
                        Blizzard tmp = new();
                        tmp.x = i;
                        tmp.y = row;
                        tmp.direction = Direction.Left;
                        blizzards.Add(tmp);
                    }
                    else if (line[i] == '>')
                    {
                        Blizzard tmp = new();
                        tmp.x = i;
                        tmp.y = row;
                        tmp.direction = Direction.Right;
                        blizzards.Add(tmp);
                    }
                    else if (line[i] == '^')
                    {
                        Blizzard tmp = new();
                        tmp.x = i;
                        tmp.y = row;
                        tmp.direction = Direction.Up;
                        blizzards.Add(tmp);
                    }
                    else if (line[i] == 'v')
                    {
                        Blizzard tmp = new();
                        tmp.x = i;
                        tmp.y = row;
                        tmp.direction = Direction.Down;
                        blizzards.Add(tmp);
                    }
                }             

            } while (!reader.EndOfStream);
            row++;

            char[,] grid = new char[row, col];

            foreach (Blizzard b in blizzards)
            {
                if (b.direction == Direction.Left) grid[b.y, b.x] = '<';
                else if (b.direction == Direction.Right) grid[b.y, b.x] = '>';
                else if (b.direction == Direction.Up) grid[b.y, b.x] = '^';
                else if (b.direction == Direction.Down) grid[b.y, b.x] = 'v';
            }

            Queue<PathDto> queue = new Queue<PathDto>();

            queue.Enqueue(new PathDto(startPosition.Item1, startPosition.Item2, 0));
            PathDto finalPos = null!;

            while (queue.Count != 0)
            {
                var p = queue.Dequeue();          

                var currPosition = (p.x, p.y);

                if (currPosition == endPosition)
                {
                    finalPos = p;
                    break;
                }
           
                if (queue.Where(q => q.x == currPosition.Item1 && q.y == currPosition.Item2 + 1 && q.step == p.step + 1).Any() == false && (currPosition.Item2 < row - 2 || currPosition.Item1 == endPosition.Item1) && anyBlizzardTargeting(grid, (currPosition.Item1, currPosition.Item2 + 1), p.step, row, col) == false)
                {
                    queue.Enqueue(new PathDto(currPosition.Item1, currPosition.Item2 + 1, p.step + 1));
                }
                if (queue.Where(q => q.x == currPosition.Item1 && q.y == currPosition.Item2 - 1 && q.step == p.step + 1).Any() == false && (currPosition.Item2 > 1) && anyBlizzardTargeting(grid, (currPosition.Item1, currPosition.Item2 - 1), p.step, row, col) == false)
                {
                    queue.Enqueue(new PathDto(currPosition.Item1, currPosition.Item2 - 1, p.step + 1));
                }
                if (queue.Where(q => q.x == currPosition.Item1 + 1 && q.y == currPosition.Item2 && q.step == p.step + 1).Any() == false && (currPosition.Item1 < col - 2 && currPosition.Item2 > 0) && anyBlizzardTargeting(grid, (currPosition.Item1 + 1, currPosition.Item2), p.step, row, col) == false)
                {
                    queue.Enqueue(new PathDto(currPosition.Item1 + 1, currPosition.Item2, p.step + 1));
                }
                if (queue.Where(q => q.x == currPosition.Item1 -1 && q.y == currPosition.Item2 && q.step == p.step + 1).Any() == false && (currPosition.Item1 > 1 && currPosition.Item2 > 0) && anyBlizzardTargeting(grid, (currPosition.Item1 - 1, currPosition.Item2), p.step, row, col) == false)
                {
                    queue.Enqueue(new PathDto(currPosition.Item1 - 1, currPosition.Item2, p.step + 1));
                }
                if (queue.Where(q => q.x == currPosition.Item1 && q.y == currPosition.Item2 && q.step == p.step + 1).Any() == false && anyBlizzardTargeting(grid, (currPosition.Item1, currPosition.Item2), p.step, row, col) == false)
                {
                    queue.Enqueue(new PathDto(currPosition.Item1, currPosition.Item2, p.step + 1));
                }
            }

            Console.WriteLine(finalPos.step);

            reader.Close();
            reader.Dispose();
        }

        public static bool anyBlizzardTargeting(char[,] grid, (int, int) position, int step, int rows, int cols)
        {
            
            if (grid[position.Item2, mod((position.Item1 - 1 - step - 1), (cols - 2)) + 1] == '>')
            {
                return true;
            }
            if (grid[position.Item2, mod((position.Item1 - 1 + step + 1), (cols - 2)) + 1] == '<')
            {
                return true;
            }
            if (grid[mod((position.Item2 - 1 + step + 1), (rows - 2)) + 1, position.Item1] == '^')
            {
                return true;
            }
            if (grid[mod((position.Item2 - 1 - step - 1), (rows - 2)) + 1, position.Item1] == 'v')
            {
                return true;
            }

            return false;
        }

        public static int mod(int x, int m)
        {
            return (x % m + m) % m;
        }
    }
}