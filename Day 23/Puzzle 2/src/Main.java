import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.LinkedList;
import java.util.List;

public class Main {


	public static void main(String[] args) {

		try {
			List<String> allLines = Files.readAllLines(Paths.get("../input.txt"));
			int firstConsideredMove = 0;
			
			List<Elf> elves = new LinkedList<>();
			
			int y = allLines.size();
			for (String line : allLines) {
				
				for(int c = 0; c < line.length(); c++)
				{
					if(line.charAt(c) == '#')
					{
						elves.add(new Elf(c, y));
					}	
				}
				y--;
			}
			
			boolean anyMoveDone = true;
			int roundNumber = 0;
			while(anyMoveDone == true)
			{
				roundNumber++;
				
				for (Elf e : elves) 
				{
					boolean canNorth = true;
					boolean canSouth = true;
					boolean canWest = true;
					boolean canEast = true;
					e.setNoMove(true);
					
					for (Elf adjE : elves) {		
						if((e.x == adjE.x && e.y == adjE.y - 1) || (e.x == adjE.x - 1 && e.y == adjE.y - 1) || (e.x == adjE.x + 1 && e.y == adjE.y - 1))
						{
							canNorth = false;
							e.setNoMove(false);
						}
						if((e.x == adjE.x && e.y == adjE.y + 1) || (e.x == adjE.x - 1 && e.y == adjE.y + 1) || (e.x == adjE.x + 1 && e.y == adjE.y + 1))
						{
							canSouth = false;
							e.setNoMove(false);
						}
						if((e.y == adjE.y && e.x == adjE.x + 1) || (e.y == adjE.y - 1 && e.x == adjE.x + 1) || (e.y == adjE.y + 1 && e.x == adjE.x + 1))
						{
							canWest = false;
							e.setNoMove(false);
						}
						if((e.y == adjE.y && e.x == adjE.x - 1) || (e.y == adjE.y - 1 && e.x == adjE.x - 1) || (e.y == adjE.y + 1 && e.x == adjE.x - 1))
						{
							canEast = false;
							e.setNoMove(false);
						}
					}
					
					if(!canNorth && !canSouth && !canWest && !canEast) e.setNoMove(true);
					if(e.isNoMove()) continue;
					
					for(int i = 0; i < 4; i++)
					{
						int move = (firstConsideredMove + i) % 4;
						
						if(move == 0 && canNorth)
						{
							e.propose(e.x, e.y + 1);
							break;
						}
						else if(move == 1 && canSouth)
						{
							e.propose(e.x, e.y - 1);
							break;
						}
						else if(move == 2 && canWest)
						{
							e.propose(e.x - 1, e.y);
							break;
						}
						else if(move == 3 && canEast)
						{
							e.propose(e.x + 1, e.y);
							break;
						}
					}				
				}			
				
				anyMoveDone = false;
				for (Elf e : elves) {		
					
					if(e.noMove) continue;
						
					boolean canMove = true;
					for (Elf otherE : elves) {	
						
						if(otherE.x == e.x && otherE.y == e.y) continue;
						
						if(otherE.isNoMove() == false && otherE.proposedX == e.proposedX && otherE.proposedY == e.proposedY)
						{
							canMove = false;
							break;
						}								
					}

					if(canMove)
					{
						e.move();
						anyMoveDone = true;
					}
				}
				
				firstConsideredMove = (firstConsideredMove + 1) % 4;
			}
				
					
			System.out.println(roundNumber);			
					
			
		} catch (IOException e) {
			System.out.println(e.getMessage());
			e.printStackTrace();
		}
	}
	

}
