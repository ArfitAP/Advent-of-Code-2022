namespace Puzzle_1
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


            /*foreach(MonkeyYell monkey in monkeys)
            {
                Console.WriteLine(monkey.Name + ": " + monkey.Value + " " + monkey.FirstOperand + " " + monkey.Operator + " " + monkey.SecondOperand);
            }*/


            while(monkeys.Where(m => m.Name == "root").Single().Value.HasValue == false)
            {
                foreach (MonkeyYell monkey in monkeys)
                {
                    var firstMonkey = monkeys.Where(m => m.Name == monkey.FirstOperand).SingleOrDefault();
                    var secondMonkey = monkeys.Where(m => m.Name == monkey.SecondOperand).SingleOrDefault();
                    if (monkey.Value.HasValue == false && firstMonkey != null && firstMonkey.Value.HasValue && secondMonkey != null && secondMonkey.Value.HasValue)
                    {
                        monkey.calculate(firstMonkey.Value.Value, secondMonkey.Value.Value);
                    }
                }
            }

            Console.WriteLine(monkeys.Where(m => m.Name == "root").Single().Value!.Value);

            reader.Close();
            reader.Dispose();
        }
    }
}