import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.HashMap;
import java.util.List;

public class Main {

	public static void main(String[] args) {

		try {
			List<String> allLines = Files.readAllLines(Paths.get("../input.txt"));
			int clock = 0;
			int register = 1;
			HashMap<Integer, Integer> values = new HashMap<Integer, Integer>();

			for (String line : allLines) {
				if(line.split(" ")[0].equals("addx"))
				{
					int inc = Integer.parseInt(line.split(" ")[1]);
					for(int t = 0; t < 2; t++)
					{
						clock++;
						
						if(clock == 20 || (clock - 20) % 40 == 0)
						{
							values.put(clock, register);
						}						
						
												
						if(t == 1) register += inc;
					}
				}
				else if (line.split(" ")[0].equals("noop"))
				{
					clock++;
					
					if(clock == 20 || (clock - 20) % 40 == 0)
					{
						values.put(clock, register);
					}
																
				}
				

			}
			
			int strength = 0;
			for (Integer i : values.keySet()) {
				
				strength += i * values.get(i);
			}
			System.out.println(strength);
			
		} catch (IOException e) {
			System.out.println(e.getMessage());
			e.printStackTrace();
		}
	}

}
