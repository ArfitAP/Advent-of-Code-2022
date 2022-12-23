using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Puzzle_2
{
    public class Blueprint
    {
        public int oreRobotCostOre { get; set; } = 0;
        public int clayRobotCostOre { get; set; } = 0;
        public int obsidianRobotCostOre { get; set; } = 0;
        public int obsidianRobotCostClay { get; set; } = 0;
        public int geodeRobotCostOre { get; set; } = 0;
        public int geodeRobotCostObsidian { get; set; } = 0;

        public int mostPossibleGeodes { get; set; } = 0;
        public int mostPossibleClay { get; set; } = 0;


        public void getMostPossibleGeodes(int minute, int oreRobots, int clayRobots, int obsidianRobots, int geodeRobots, int ore, int clay, int obsidian, int geodes)
        {
            if (minute == 0)
            {
                if (geodes > mostPossibleGeodes) mostPossibleGeodes = geodes;
                if (clay > mostPossibleClay) mostPossibleClay = clay;
                return;
            }

            //if (minute == 20) { Console.WriteLine(oreRobots); }

            int n = minute - 1;
            int maxPossibleScore = n * (n + 1) / 2 + geodeRobots * (n + 1) + geodes;
            if (maxPossibleScore <= mostPossibleGeodes) return;

            int nextore = ore + oreRobots;
            int nextclay = clay + clayRobots;
            int nextobsidian = obsidian + obsidianRobots;
            int nextgeodes = geodes + geodeRobots;

            if (ore >= geodeRobotCostOre && obsidian >= geodeRobotCostObsidian)
            {
                getMostPossibleGeodes(minute - 1, oreRobots, clayRobots, obsidianRobots, geodeRobots + 1, nextore - geodeRobotCostOre, nextclay, nextobsidian - geodeRobotCostObsidian, nextgeodes);
            }
            if (ore >= obsidianRobotCostOre && clay >= obsidianRobotCostClay)
            {
                getMostPossibleGeodes(minute - 1, oreRobots, clayRobots, obsidianRobots + 1, geodeRobots, nextore - obsidianRobotCostOre, nextclay - obsidianRobotCostClay, nextobsidian, nextgeodes);
            }
            if (ore >= clayRobotCostOre)
            {
                getMostPossibleGeodes(minute - 1, oreRobots, clayRobots + 1, obsidianRobots, geodeRobots, nextore - clayRobotCostOre, nextclay, nextobsidian, nextgeodes);
            }
            if (ore >= oreRobotCostOre)
            {
                getMostPossibleGeodes(minute - 1, oreRobots + 1, clayRobots, obsidianRobots, geodeRobots, nextore - oreRobotCostOre, nextclay, nextobsidian, nextgeodes);
            }

            getMostPossibleGeodes(minute - 1, oreRobots, clayRobots, obsidianRobots, geodeRobots, nextore, nextclay, nextobsidian, nextgeodes);

        }

    }
}
