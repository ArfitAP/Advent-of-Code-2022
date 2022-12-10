import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.HashMap;
import java.util.List;

public class Main {

	public static void main(String[] args) {

		try {
			List<String> allLines = Files.readAllLines(Paths.get("../input.txt"));
			String[] spritePos = new String[40];
			String[][] screen = new String[6][40];
			
			for(int i = 0; i < 40; i++)
			{
				if(i < 3) spritePos[i] = "#";
				else spritePos[i] = ".";
			}
			
			for(int i = 0; i < 6; i++)
			{
				for(int j = 0; j < 40; j++)
				{
					screen[i][j] = " ";
				}
			}
			
			int clock = 0;
			int register = 1;
			int crtRow = 0;
			int crtCol = 0;
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
						
						if(spritePos[crtCol].equals("#")) screen[crtRow][crtCol] = "#";
						else screen[crtRow][crtCol] = ".";
						
						crtCol++;
						if(crtCol >= 40)
						{
							crtRow++;
							crtCol = 0;
						}
												
						if(t == 1)
						{
							register += inc;
							
							for(int i = 0; i < 40; i++)
							{
								if(i >= register - 1 && i <= register + 1) spritePos[i] = "#";
								else spritePos[i] = ".";
							}
						}
					}
				}
				else if (line.split(" ")[0].equals("noop"))
				{
					clock++;
					
					if(clock == 20 || (clock - 20) % 40 == 0)
					{
						values.put(clock, register);
					}
					
					if(spritePos[crtCol].equals("#")) screen[crtRow][crtCol] = "#";
					else screen[crtRow][crtCol] = ".";
					
					crtCol++;
					if(crtCol >= 40)
					{
						crtRow++;
						crtCol = 0;
					}
																
				}
				

			}
			
			
			for(int i = 0; i < 6; i++)
			{
				for(int j = 0; j < 40; j++)
				{
					System.out.print(screen[i][j]);
				}
				System.out.println("");
			}

			
		} catch (IOException e) {
			System.out.println(e.getMessage());
			e.printStackTrace();
		}
	}

}
