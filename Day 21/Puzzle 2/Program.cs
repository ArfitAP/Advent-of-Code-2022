namespace Puzzle_2
{
    internal class Program
    {
        static void Main(string[] args)
        {
            string line;
            List<MonkeyYell> monkeys = new List<MonkeyYell>();

            StreamReader reader;
            reader = new StreamReader(@"../../../../input.txt");
            do
            {
                line = reader.ReadLine()!;

                MonkeyYell monkey = new MonkeyYell();

                monkey.Name = line.Split(":")[0];

                int value = 0;
                var isNumber = Int32.TryParse(line.Split(":")[1], out value);
                if(isNumber)
                {
                    monkey.Value = value;
                }
                else
                {
                    monkey.Value = null;
                    monkey.FirstOperand = line.Split(":")[1].Substring(1, 4);
                    monkey.SecondOperand = line.Split(":")[1].Substring(8, 4);
                    monkey.Operator = line.Split(":")[1].Substring(6, 1);
                }

                monkeys.Add(monkey);

            } while (!reader.EndOfStream);

            monkeys.Where(m => m.Name == "humn").Single().Value = null;
            monkeys.Where(m => m.Name == "root").Single().Operator = "=";


            while (true)
            {
                bool calculationDone = false;
                foreach (MonkeyYell monkey in monkeys)
                {
                    var firstMonkey = monkeys.Where(m => m.Name == monkey.FirstOperand).SingleOrDefault();
                    var secondMonkey = monkeys.Where(m => m.Name == monkey.SecondOperand).SingleOrDefault();
                    if (monkey.Value.HasValue == false && firstMonkey != null && firstMonkey.Value.HasValue && secondMonkey != null && secondMonkey.Value.HasValue)
                    {
                        monkey.calculate(firstMonkey.Value.Value, secondMonkey.Value.Value);
                        calculationDone = true;
                    }
                }

                if (calculationDone == false) break;
            }

            double rootOperandTarget = 0;
            string rootOperandName = string.Empty;
            if(monkeys.Where(m => m.Name == (monkeys.Where(m => m.Name == "root").Single().FirstOperand)).Single().Value.HasValue)
            {
                rootOperandTarget = monkeys.Where(m => m.Name == (monkeys.Where(m => m.Name == "root").Single().FirstOperand)).Single().Value!.Value;
                rootOperandName = monkeys.Where(m => m.Name == "root").Single().SecondOperand;
            }
            else
            {
                rootOperandTarget = monkeys.Where(m => m.Name == (monkeys.Where(m => m.Name == "root").Single().SecondOperand)).Single().Value!.Value;
                rootOperandName = monkeys.Where(m => m.Name == "root").Single().FirstOperand;
            }

            var equation = constructEquation(monkeys, rootOperandName);

            Console.WriteLine(printEquation(equation) + " = " + rootOperandTarget);


            while(equation.valueLeft != "humn" && equation.valueRight != "humn")
            {
                if(equation.Operator == "/")
                {
                    if(equation.valueLeft != string.Empty)
                    {
                        rootOperandTarget = double.Parse(equation.valueLeft) / rootOperandTarget;
                        equation = equation.opRight;
                    }
                    else if (equation.valueRight != string.Empty)
                    {
                        rootOperandTarget = double.Parse(equation.valueRight) * rootOperandTarget;
                        equation = equation.opLeft;
                    }
                }
                else if(equation.Operator == "*")
                {
                    if (equation.valueLeft != string.Empty)
                    {
                        rootOperandTarget = rootOperandTarget / double.Parse(equation.valueLeft);
                        equation = equation.opRight;
                    }
                    else if (equation.valueRight != string.Empty)
                    {
                        rootOperandTarget = rootOperandTarget / double.Parse(equation.valueRight);
                        equation = equation.opLeft;
                    }
                }
                else if (equation.Operator == "-")
                {
                    if (equation.valueLeft != string.Empty)
                    {
                        rootOperandTarget = double.Parse(equation.valueLeft) - rootOperandTarget;
                        equation = equation.opRight;
                    }
                    else if (equation.valueRight != string.Empty)
                    {
                        rootOperandTarget = rootOperandTarget + double.Parse(equation.valueRight);
                        equation = equation.opLeft;
                    }
                }
                else if (equation.Operator == "+")
                {
                    if (equation.valueLeft != string.Empty)
                    {
                        rootOperandTarget = rootOperandTarget - double.Parse(equation.valueLeft);
                        equation = equation.opRight;
                    }
                    else if (equation.valueRight != string.Empty)
                    {
                        rootOperandTarget = rootOperandTarget - double.Parse(equation.valueRight);
                        equation = equation.opLeft;
                    }
                }
            }

            Console.WriteLine(printEquation(equation) + " = " + rootOperandTarget);

            double humanResult = 0;
            if (equation.Operator == "/")
            {
                if (equation.valueLeft == "humn")
                {
                    humanResult = double.Parse(equation.valueRight) * rootOperandTarget;
                }
                else if (equation.valueRight == "humn")
                {
                    humanResult = double.Parse(equation.valueLeft) / rootOperandTarget;
                }
            }
            else if (equation.Operator == "*")
            {
                if (equation.valueLeft == "humn")
                {
                    humanResult = rootOperandTarget / double.Parse(equation.valueRight);
                }
                else if (equation.valueRight == "humn")
                {
                    humanResult = rootOperandTarget / double.Parse(equation.valueLeft);
                }
            }
            else if (equation.Operator == "-")
            {
                if (equation.valueLeft == "humn")
                {
                    humanResult = rootOperandTarget + double.Parse(equation.valueRight);
                }
                else if (equation.valueRight == "humn")
                {
                    humanResult = double.Parse(equation.valueLeft) - rootOperandTarget;
                }
            }
            else if (equation.Operator == "+")
            {
                if (equation.valueLeft == "humn")
                {
                    humanResult = rootOperandTarget - double.Parse(equation.valueRight);
                }
                else if (equation.valueRight == "humn")
                {
                    humanResult = rootOperandTarget - double.Parse(equation.valueLeft);
                }
            }

            Console.WriteLine("humn = " + humanResult);

            reader.Close();
            reader.Dispose();
        }


        public static string printEquation(Operation operation)
        {
            string res = "";

            if(operation.valueLeft != string.Empty && operation.valueRight != string.Empty)
            {
                res = operation.ToString();
            }
            else if(operation.valueLeft == string.Empty && operation.valueRight != string.Empty)
            {
                res = "(" + printEquation(operation.opLeft) + " " + operation.Operator + " " + operation.valueRight + ")";
            }
            else if (operation.valueLeft != string.Empty && operation.valueRight == string.Empty)
            {
                res = "(" + operation.valueLeft + " " + operation.Operator + " " + printEquation(operation.opRight) + ")";
            }
            else
            {
                res = "(" + printEquation(operation.opLeft) + " " + operation.Operator + " " + printEquation(operation.opRight) + ")";
            }

            return res;
        }

        public static Operation constructEquation(List<MonkeyYell> monkeys, string rootOperandName)
        {

            var operand = monkeys.Where(m => m.Name == rootOperandName).Single();

            var left = monkeys.Where(m => m.Name == operand.FirstOperand).Single();
            var right = monkeys.Where(m => m.Name == operand.SecondOperand).Single();

            Operation res = new();

            res.Operator = operand.Operator;

            if(left.Name == "humn")
            {
                res.valueLeft = "humn";
                res.opLeft = null!;
            }
            else if(left.Value.HasValue)
            {
                res.valueLeft = left.Value.Value.ToString();
                res.opLeft = null!;
            }
            else
            {
                res.opLeft = constructEquation(monkeys, left.Name);
                res.valueLeft = string.Empty;
            }

            if (right.Name == "humn")
            {
                res.valueRight = "humn";
                res.opRight = null!;
            }
            else if (right.Value.HasValue)
            {
                res.valueRight = right.Value.Value.ToString();
                res.opRight = null!;
            }
            else
            {
                res.opRight = constructEquation(monkeys, right.Name);
                res.valueRight = string.Empty;
            }

            return res;
        }
    }
}