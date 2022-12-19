import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Iterator;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

public class Main {

	public static int best = 0;
	public static int myBest = 0;
	public static int elephsBest = 0;
	public static List<Node> allValves;
	public static int nodes = 0;
	public static Map<String, Map<String, Integer>> distances;
	public static Object[] myNodes = new Object[0];
	public static Object[] elephsNodes = new Object[0];
	public static List<List<Node>> selection = new LinkedList<List<Node>>();
	
	
	public static void main(String[] args) {

		try {
			List<String> allLines = Files.readAllLines(Paths.get("../input.txt"));

			Node startPosition = new Node();
			
			allValves = new LinkedList<>();
			for (String line : allLines) {
				
				Node tmp = new Node();
				tmp.setName(line.split(";")[0].split("has")[0].substring(6).trim());
				tmp.setFlowRate(Integer.parseInt(line.split(";")[0].split("=")[1].trim()));
				String[] tunnels = line.split(";")[1].substring(23).trim().split(",");
				for(String tunell : tunnels)
				{
					tmp.getTunnels().add(tunell.trim());
				}
				
				if(tmp.getName().equals("AA")) startPosition = tmp;
				allValves.add(tmp);		
			}
			
			distances = new HashMap<String, Map<String, Integer>>();
			nodes = allValves.size();
			
			List<Node> openedValves = new LinkedList<Node>();
			List<Node> closedValves = new LinkedList<Node>();
			for (Node i : allValves) {
				
				distances.put(i.getName(), new HashMap<String, Integer>());
				if(i.getFlowRate() == 0) 
				{
					openedValves.add(i);
					i.setOpened(true);
				}
				else
				{
					closedValves.add(i);
				}
			}
			
			
			getDistances();
			System.out.println("Distances calculated");
			
			
			List<Node> myOpenedValves = new LinkedList<Node>(openedValves);
			List<Node> elephsOpenedValves = new LinkedList<Node>(openedValves);
			
			for(int s = 1; s < closedValves.size(); s++)
			{
				System.out.println(s + "/" + (closedValves.size() - 1));
				
				selection = new LinkedList<List<Node>>();
				Object arr[] = closedValves.toArray();
				printCombination(arr, closedValves.size(), s);
				
				for(List<Node> currSelection : selection)
				{
					myOpenedValves = new LinkedList<Node>(openedValves);
					elephsOpenedValves = new LinkedList<Node>(openedValves);
					
					for (Node i : closedValves) {
						
						if(currSelection.contains(i)) 
						{
							elephsOpenedValves.add(i);
						}
						else
						{
							myOpenedValves.add(i);
						}
					}
					
					/*printList(myOpenedValves);
					System.out.println();
					printList(elephsOpenedValves);*/
					
					myBest = elephsBest = 0;
					getBestPathOnlyNodes(startPosition, 26, 0, myOpenedValves, new LinkedList<Node>(), false);
					getBestPathOnlyNodes(startPosition, 26, 0, elephsOpenedValves, new LinkedList<Node>(), true);
					
					int tmpScore = myBest + elephsBest;
					if(tmpScore > best) 
					{
						best = tmpScore;
					}
					
				}
			}
								
			System.out.println(best);
			
			
		} catch (IOException e) {
			System.out.println(e.getMessage());
			e.printStackTrace();
		}
	}
	
	public static void getBestPathOnlyNodes(Node node, int minute, int currPressureScore, List<Node> openedValves, List<Node> currentPath, boolean eleph)
	{
		List<Node> newPath = new LinkedList<Node>(currentPath);
		newPath.add(node);
		
		boolean allopened = true;
		for (Node i : allValves) {
			
			if(openedValves.contains(i) == false) 
			{
				allopened = false;
				break;
			}
		}
		
		if(minute == 0 || allopened)
		{
			if(eleph)
			{
				if(currPressureScore > elephsBest) 
				{
					elephsBest = currPressureScore;
				}
			}
			else
			{
				if(currPressureScore > myBest) 
				{
					myBest = currPressureScore;
				}
			}
				
			return;
		}
		
		if(allopened) return;
		
		List<String> options = new LinkedList<String>();
		for (Node i : allValves) {
			
			if(openedValves.contains(i) == false) 
			{
				options.add(i.getName());
			}
		}
		
		for(String s : options)
		{
			//if(minute == 25) System.out.println(s + " " + best);
				
			Node next = new Node();
			for (Node i : allValves) {			
				if(i.getName().equals(s))
				{
					next = new Node(i);
					break;
				}
			}
			
			List<Node> newOpenedValves = new LinkedList<Node>(openedValves);
			newOpenedValves.add(next);
			int requiredminutes = distances.get(node.getName()).get(next.getName()) + 1;
			if(minute - requiredminutes > 0)
			{
				int tmpScore = (minute - requiredminutes) * next.getFlowRate();
				getBestPathOnlyNodes(next, minute - requiredminutes, currPressureScore + tmpScore, newOpenedValves, newPath, eleph);
			}
			
		}
		
		if(eleph)
		{
			if(currPressureScore > elephsBest) 
			{
				elephsBest = currPressureScore;
			}
		}
		else
		{
			if(currPressureScore > myBest) 
			{
				myBest = currPressureScore;
			}
		}
	}
	
	
	public static void getDistances()
	{
		for (Node from : allValves) {
			
			for (Node to : allValves) {
				
				distances.get(from.getName()).put(to.getName(), 0);
				
				int dist = 0;
				boolean visited[] = new boolean[nodes];
				LinkedList<NodeDistance> queue = new LinkedList<NodeDistance>();
				
				visited[allValves.indexOf(from)]=true;
		        queue.add(new NodeDistance(from, 0));
		        
		        while (queue.size() != 0)
		        {
		        	NodeDistance n = queue.poll();
		        	
		        	if(n.node.equals(to))
		        	{
		        		distances.get(from.getName()).put(to.getName(), n.distance);
		        		break;
		        	}
		 
		            Iterator<String> i = n.node.getTunnels().listIterator();
		            while (i.hasNext())
		            {
		            	NodeDistance next = new NodeDistance(n.node, 0);
		            	String nextName = i.next();
						for (Node nn : allValves) {			
							if(nn.getName().equals(nextName))
							{
								next = new NodeDistance(nn, n.distance + 1);
								break;
							}
						}
						
		                if (!visited[allValves.indexOf(next.node)])
		                {
		                    visited[allValves.indexOf(next.node)] = true;
		                    queue.add(next);
		                }
		            }
		        }

			}

		}
		
	}
	
	static void printCombination(Object arr[], int n, int r)
    {
        // A temporary array to store all combination
        // one by one
		Object data[] = new Node[r];
 
        // Print all combination using temporary
        // array 'data[]'
        combinationUtil(arr, n, r, 0, data, 0);
    }
	
	static void combinationUtil(Object arr[], int n, int r,
            int index, Object data[], int i)
	{
		// Current combination is ready to be printed,
		// print it
		if (index == r) {
			selection.add(new LinkedList<Node>());
			for (int j = 0; j < r; j++)
			{
				selection.get(selection.size() - 1).add((Node)data[j]);
				//System.out.print(data[j] + " ");
			}
			
			//System.out.println("");
			return;
		}
		
		// When no more elements are there to put in data[]
		if (i >= n)
		return;
		
		// current is included, put next at next
		// location
		data[index] = arr[i];
		combinationUtil(arr, n, r, index + 1,
		                 data, i + 1);
		
		// current is excluded, replace it with
		// next (Note that i+1 is passed, but
		// index is not changed)
		combinationUtil(arr, n, r, index, data, i + 1);
	}
	
	private static void printArray(Node[] elements, char delimiter) {
	    String delimiterSpace = delimiter + " ";
	    for(int i = 0; i < elements.length; i++) {
	        System.out.print(elements[i] + delimiterSpace);
	    }
	    System.out.print('\n');
	}
	
	private static void swap(Node[] elements, int a, int b) {
		Node tmp = elements[a];
	    elements[a] = elements[b];
	    elements[b] = tmp;
	}
	
	private static void printList(List<Node> nodes)
	{
		for (Node i : nodes) {
			
			System.out.println(i);
		}
		System.out.println();
	}
	
	private static class NodeDistance
	{
		public Node node;
		public int distance;
		
		public NodeDistance(Node node, int distance) {
			super();
			this.node = node;
			this.distance = distance;
		}
			
	}

}
