using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Xml.Linq;

namespace Puzzle_2
{
    public class Operation
    {
        public Operation opLeft { get; set; } = null!;
        public string valueLeft { get; set; } = string.Empty;

        public string Operator { get; set; } = string.Empty;

        public Operation opRight { get; set; } = null!;
        public string valueRight { get; set; } = string.Empty;

        public override string ToString()
        {
            return "(" + valueLeft + " " + Operator + " " + valueRight + ")";
        }
    }
}
