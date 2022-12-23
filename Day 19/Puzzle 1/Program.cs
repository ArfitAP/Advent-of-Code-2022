namespace Puzzle_1
{
    internal class Program
    {
        static void Main(string[] args)
        {
            string line;
            List<Blueprint> blueprints = new List<Blueprint>();

            StreamReader reader;
            reader = new StreamReader(@"../../../../input.txt");
            do
            {
                line = reader.ReadLine()!;

                Blueprint bp = new Blueprint();

                bp.oreRobotCostOre = Int32.Parse(line.Split("costs")[1].Split("ore")[0].Trim());
                bp.clayRobotCostOre = Int32.Parse(line.Split("costs")[2].Split("ore")[0].Trim());
                bp.obsidianRobotCostOre = Int32.Parse(line.Split("costs")[3].Split("ore")[0].Trim());
                bp.obsidianRobotCostClay = Int32.Parse(line.Split("and")[1].Split("clay")[0].Trim());
                bp.geodeRobotCostOre = Int32.Parse(line.Split("costs")[4].Split("ore")[0].Trim());
                bp.geodeRobotCostObsidian = Int32.Parse(line.Split("and")[2].Split("obsidian")[0].Trim());

                blueprints.Add(bp);

            } while (!reader.EndOfStream);

            int scoreSum = 0;
            foreach(Blueprint bp in blueprints)
            {
                bp.getMostPossibleGeodes(24, 1, 0, 0, 0, 0, 0, 0, 0);
                scoreSum += (blueprints.IndexOf(bp) + 1) * bp.mostPossibleGeodes;
            }

            Console.WriteLine(scoreSum);

            reader.Close();
            reader.Dispose();

        }

    }
}