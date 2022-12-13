import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Set;

public class Main {

	public static void main(String[] args) {
		
		try {
			
			List<String> allLines = Files.readAllLines(Paths.get("../input.txt"));
			int rows = allLines.size();
			int cols = allLines.get(0).length();
			
			int[][] map = new int[rows][cols];
			
			Position startPos = new Position(0, 0, 0, 0);
			Position endPos = new Position(0, 0, 0, 0);
			
			int i = 0;

			for (String line : allLines) {
				
				for (int j = 0; j < line.length(); j++) {
					 
					char tmp = line.charAt(j);
					if(tmp == 'S')
					{
						map[i][j] = 0;
						startPos = new Position(i, j, 0, 0);
					}
					else if(tmp == 'E')
					{
						map[i][j] = 'z' - 'a';
						endPos = new Position(i, j, 'z' - 'a', 0);
					}
					else
					{
						map[i][j] = line.charAt(j) - 'a';
					}
		        }
				i++;
			}
			
			
			/*
			for (i = 0; i < rows; i++) {
				for (int j = 0; j < cols; j++) {
					System.out.print(map[i][j] + " ");
				}
				System.out.println();
			}
			*/
			
			Set<Position> visited = new HashSet<Position>();
			
	        LinkedList<Position> queue = new LinkedList<Position>();
	        
	        visited.add(startPos);
	        queue.add(startPos);
	        
	        while (queue.size() != 0)
	        {
	            Position s = queue.poll();	           
	            
	            if(s.equals(endPos))
	            {
	            	System.out.println(s.steps);
	            	break;
	            }
	            
	            if(s.getRow() > 0 && map[s.getRow()-1][s.getCol()] - 1 <= s.height)
	            {
	            	Position adjecent = new Position(s.getRow()-1, s.getCol(), map[s.getRow()-1][s.getCol()], s.steps + 1);
	            	if(visited.contains(adjecent) == false)
	            	{
	            		visited.add(adjecent);
	            		queue.add(adjecent);
	            	}
	            }
	            if(s.getRow() < rows - 1 && map[s.getRow()+1][s.getCol()] - 1 <= s.height)
	            {
	            	Position adjecent = new Position(s.getRow()+1, s.getCol(), map[s.getRow()+1][s.getCol()], s.steps + 1);
	            	if(visited.contains(adjecent) == false)
	            	{
	            		visited.add(adjecent);
	            		queue.add(adjecent);
	            	}
	            }
	            if(s.getCol() > 0 && map[s.getRow()][s.getCol() - 1] - 1 <= s.height)
	            {
	            	Position adjecent = new Position(s.getRow(), s.getCol() - 1, map[s.getRow()][s.getCol() - 1], s.steps + 1);
	            	if(visited.contains(adjecent) == false)
	            	{
	            		visited.add(adjecent);
	            		queue.add(adjecent);
	            	}
	            }
	            if(s.getCol() < cols - 1 && map[s.getRow()][s.getCol() + 1] - 1 <= s.height)
	            {
	            	Position adjecent = new Position(s.getRow(), s.getCol() + 1, map[s.getRow()][s.getCol() + 1], s.steps + 1);
	            	if(visited.contains(adjecent) == false)
	            	{
	            		visited.add(adjecent);
	            		queue.add(adjecent);
	            	}
	            }

	        }
	        
			
		} catch (IOException e) {
			System.out.println(e.getMessage());
			e.printStackTrace();
		}
	
	}
	

}


