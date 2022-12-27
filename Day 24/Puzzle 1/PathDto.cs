using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Puzzle_1
{
    public class PathDto
    {
        public int x { get; set; }
        public int y { get; set; }
        public int step { get; set; }   

        public PathDto(int x, int y, int step)
        {
            this.x = x;
            this.y = y;
            this.step = step;
        }
    }
}
