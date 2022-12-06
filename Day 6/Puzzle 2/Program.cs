namespace Puzzle_2
{
    internal class Program
    {
        static void Main(string[] args)
        {
            char ch;
            int count = 0;
            List<char> window = new();

            StreamReader reader;
            reader = new StreamReader(@"../../../../input.txt");
            do
            {
                ch = (char)reader.Read();

                count++;

                if (window.Count < 14) window.Add(ch);
                else
                {
                    window.RemoveAt(0);
                    window.Add(ch);
                }

                if (window.Count >= 14 && checkAllDifferent(window)) break;


            } while (!reader.EndOfStream);

            reader.Close();
            reader.Dispose();

            Console.WriteLine(count);
            Console.ReadLine();
        }

        static bool checkAllDifferent(List<char> chars)
        {
            return chars.Distinct().Count() == chars.Count;
        }
    }
}