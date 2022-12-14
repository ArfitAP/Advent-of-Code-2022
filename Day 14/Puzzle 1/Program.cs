namespace Puzzle_1
{
    internal class Program
    {
        static void Main(string[] args)
        {
            string line;

            char[,] grid = new char[500, 1001];

            for (int i = 0; i < 500; i++)
            {
                for (int j = 0; j < 1001; j++)
                {
                    grid[i, j] = '.';
                }
            }

            grid[0, 500] = '+';

            StreamReader reader;
            reader = new StreamReader(@"../../../../input.txt");
            do
            {
                line = reader.ReadLine()!;

                var pairs = line.Split("->");
                for(int i = 1; i < pairs.Length; i++)
                {
                    int col1 = Int32.Parse(pairs[i - 1].Trim().Split(",")[0]);
                    int row1 = Int32.Parse(pairs[i - 1].Trim().Split(",")[1]);

                    int col2 = Int32.Parse(pairs[i].Trim().Split(",")[0]);
                    int row2 = Int32.Parse(pairs[i].Trim().Split(",")[1]);

                    if(col1 == col2)
                    {
                        var startRow = Math.Min(row1, row2);
                        var endRow = Math.Max(row1, row2);
                        for (int r = startRow; r <= endRow; r++)
                        {
                            grid[r, col1] = '#';
                        }
                    }
                    else
                    {
                        var startCol = Math.Min(col1, col2);
                        var endCol = Math.Max(col1, col2);
                        for (int c = startCol; c <= endCol; c++)
                        {
                            grid[row1, c] = '#';
                        }
                    }
                }


            } while (!reader.EndOfStream);

            int sands = 0;
            bool rest = true;

            while (rest)
            {
                grid[0, 500] = '+';
                int currrow = 0;
                int currcol = 500;
                bool moved = false;

                do
                {
                    if(currrow == 499)
                    {
                        rest = false;
                        break;
                    }

                    if(grid[currrow + 1, currcol] == '.')
                    {
                        grid[currrow, currcol] = '.';
                        grid[currrow + 1, currcol] = '+';
                        currrow++;
                        moved = true;
                    }
                    else if(grid[currrow + 1, currcol - 1] == '.')
                    {
                        grid[currrow, currcol] = '.';
                        grid[currrow + 1, currcol - 1] = '+';
                        currrow++;
                        currcol--;
                        moved = true;
                    }
                    else if (grid[currrow + 1, currcol + 1] == '.')
                    {
                        grid[currrow, currcol] = '.';
                        grid[currrow + 1, currcol + 1] = '+';
                        currrow++;
                        currcol++;
                        moved = true;
                    }
                    else
                    {
                        grid[currrow, currcol] = 'o';
                        moved = false;
                        sands++;
                    }


                } while (moved);

            }

            Console.WriteLine(sands);

            /*for (int i = 0; i < 10; i++)
            {
                for (int j = 490; j < 510; j++)
                {
                    Console.Write(grid[i,j]);
                }
                Console.WriteLine();
            }*/

            reader.Close();
            reader.Dispose();
        }
    }
}