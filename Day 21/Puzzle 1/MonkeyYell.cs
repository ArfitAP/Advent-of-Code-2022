using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Puzzle_1
{
    public class MonkeyYell
    {
        public string Name { get; set; } = string.Empty;
        public double? Value { get; set; } = null;
        public string FirstOperand { get; set; } = string.Empty;
        public string SecondOperand { get; set; } = string.Empty;
        public string Operator { get; set; } = string.Empty;

        public void calculate(double firstOperand, double secondOperand)
        {
            if(Operator == "+")
            {
                Value = firstOperand + secondOperand;
            }
            else if(Operator == "-") 
            {
                Value = firstOperand - secondOperand;
            }
            else if(Operator == "*") 
            {   
                Value = firstOperand * secondOperand;
            }
            else if(Operator == "/")
            {
                Value = firstOperand / secondOperand;
            }
        }
    }
}
