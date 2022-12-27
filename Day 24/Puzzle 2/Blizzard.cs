using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Puzzle_2
{
    public class Blizzard
    {
        public int x { get; set; }
        public int y { get; set; }
        public Direction direction { get; set; }

        public void move(int rows, int cols)
        {
            if(direction == Direction.Left)
            {
                if (x == 1) x = cols - 2;
                else x--;
            }
            else if (direction == Direction.Right)
            {
                if (x == cols - 2) x = 1;
                else x++;
            }
            else if (direction == Direction.Up)
            {
                if (y == 1) y = rows - 2;
                else y--;
            }
            else if (direction == Direction.Down)
            {
                if (y == rows - 2) y = 1;
                else y++;
            }
        }
    }

    public enum Direction 
    {
        Left,
        Right,
        Up,
        Down
    }
}
