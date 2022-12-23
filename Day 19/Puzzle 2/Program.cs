using static System.Net.Mime.MediaTypeNames;

namespace Puzzle_2
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


            
            int scoreMul = 1;
            
            foreach (Blueprint bp in blueprints)
            {
                if (blueprints.IndexOf(bp) == 3) break;

                //Console.WriteLine(blueprints.IndexOf(bp) + 1 + "/3");
                bp.getMostPossibleGeodes(32, 1, 0, 0, 0, 0, 0, 0, 0);
                scoreMul *= bp.mostPossibleGeodes;
            }
            Console.WriteLine(scoreMul);


            //Blueprint test = blueprints.First();
            //test.getMostPossibleGeodes(32, 1, 0, 0, 0, 0, 0, 0, 0);
            //Console.WriteLine(test.mostPossibleGeodes);

            reader.Close();
            reader.Dispose();
        }
    }
}