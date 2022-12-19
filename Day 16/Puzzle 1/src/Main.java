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
	public static List<Node> allValves;
	public static int nodes = 0;
	public static Map<String, Map<String, Integer>> distances;
	
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
			for (Node i : allValves) {
				
				distances.put(i.getName(), new HashMap<String, Integer>());
				if(i.getFlowRate() == 0) 
				{
					openedValves.add(i);
					i.setOpened(true);
				}

			}
			
			getDistances();
			System.out.println("Distances calculated");
			
			
			getBestPathOnlyNodes(startPosition, 30, 0, openedValves, new LinkedList<Node>());
			
			System.out.println(best);
			
			
		} catch (IOException e) {
			System.out.println(e.getMessage());
			e.printStackTrace();
		}
	}
	
	public static void getBestPathOnlyNodes(Node node, int minute, int currPressureScore, List<Node> openedValves, List<Node> currentPath)
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
			if(currPressureScore > best) 
			{
				best = currPressureScore;
				//printList(currentPath);
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
				getBestPathOnlyNodes(next, minute - requiredminutes, currPressureScore + tmpScore, newOpenedValves, newPath);
			}
			
		}
		
		if(currPressureScore > best) 
		{
			best = currPressureScore;
			//printList(newPath);
		}

	}
	
	
	public static void getBestPathPermutting(int n, Node[] elements)
	{
		if(n == 1) 
		{
	        printArray(elements, ',');
	    } 
		else {
	        for(int i = 0; i < n-1; i++) {

	        	getBestPathPermutting(n - 1, elements);
	        	
	            if(n % 2 == 0) {
	                swap(elements, i, n-1);
	            } else {
	                swap(elements, 0, n-1);
	            }
	        }
	        getBestPathPermutting(n - 1, elements);
	    }
	}
	
	
	public static void getBestPath(Node node, int minute, int currPressureScore, List<Node> openedValves, List<Node> currentPath)
	{
		List<Node> newPath = new LinkedList<Node>(currentPath);
		newPath.add(node);
		
		//if(openedValves.size() == nodes) return;
		int numopened = 0;
		boolean allopened = true;
		Node firstopened = new Node();
		Node secondopened = new Node();
		for (Node i : allValves) {
			
			if(openedValves.contains(i) == false) 
			{
				allopened = false;
				numopened++;
				if(numopened == 1) {firstopened = new Node(i); firstopened.opened = i.isOpened();}
				else if(numopened == 2) {secondopened = new Node(i); secondopened.opened = i.isOpened();}
				else break;
			}
		}
	
		
		if(numopened == 2)
		{
			int requiredminutestofirst = distances.get(node.getName()).get(firstopened.getName());
			int requiredminutestosecond = distances.get(node.getName()).get(secondopened.getName());
			int requiredminutesfromfirsttosecond = distances.get(firstopened.getName()).get(secondopened.getName());
			
			if((requiredminutestofirst + requiredminutesfromfirsttosecond + 2) < minute && (requiredminutestosecond + requiredminutesfromfirsttosecond + 2) < minute)
			{
				int newScoreToFirst = currPressureScore + firstopened.getFlowRate() * (minute - requiredminutestofirst - 1) + secondopened.getFlowRate() * (minute - requiredminutestofirst - requiredminutesfromfirsttosecond - 2);
				int newScoreToSecond = currPressureScore + secondopened.getFlowRate() * (minute - requiredminutestosecond - 1) + firstopened.getFlowRate() * (minute - requiredminutestosecond - requiredminutesfromfirsttosecond - 2);
				
				if(newScoreToFirst > best)
				{
					best = newScoreToFirst;
				}
				else if(newScoreToSecond > best)
				{
					best = newScoreToSecond;
				}
				return;
			}
			else if((requiredminutestofirst + requiredminutesfromfirsttosecond + 2) < minute)
			{
				int newScoreToFirst = currPressureScore + firstopened.getFlowRate() * (minute - requiredminutestofirst - 1) + secondopened.getFlowRate() * (minute - requiredminutestofirst - requiredminutesfromfirsttosecond - 2);
				
				if(newScoreToFirst > best)
				{
					best = newScoreToFirst;
				}
				return;
			}
			else if((requiredminutestosecond + requiredminutesfromfirsttosecond + 2) < minute)
			{
				int newScoreToSecond = currPressureScore + secondopened.getFlowRate() * (minute - requiredminutestosecond - 1) + firstopened.getFlowRate() * (minute - requiredminutestosecond - requiredminutesfromfirsttosecond - 2);
				
				if(newScoreToSecond > best)
				{
					best = newScoreToSecond;
				}
				return;
			}
			else if(requiredminutestofirst + 1 < minute && requiredminutestosecond + 1 < minute)
			{
				int newScoreToFirst = currPressureScore + firstopened.getFlowRate() * (minute - requiredminutestofirst - 1);
				int newScoreToSecond = currPressureScore + secondopened.getFlowRate() * (minute - requiredminutestosecond - 1);
				
				if(newScoreToFirst > best)
				{
					best = newScoreToFirst;
				}
				else if(newScoreToSecond > best)
				{
					best = newScoreToSecond;
				}
				return;
			}
			else if(requiredminutestofirst + 1 < minute)
			{
				int newScoreToFirst = currPressureScore + firstopened.getFlowRate() * (minute - requiredminutestofirst - 1);
				
				if(newScoreToFirst > best)
				{
					best = newScoreToFirst;
				}
				return;
			}
			else if(requiredminutestosecond + 1 < minute)
			{
				int newScoreToSecond = currPressureScore + secondopened.getFlowRate() * (minute - requiredminutestosecond - 1);
				
				if(newScoreToSecond > best)
				{
					best = newScoreToSecond;
				}
				return;
			}
			else return;
		}
		
		if(numopened == 1)
		{
			int requiredminutes = distances.get(node.getName()).get(firstopened.getName()) + 1;
			if(requiredminutes < minute)
			{
				int newScore = currPressureScore + firstopened.getFlowRate() * (minute - requiredminutes);
				if(newScore > best) 
				{
					best = newScore;
				}
			}
			return;
		}
		
		if(minute == 0 || allopened)
		{
			if(currPressureScore > best) 
			{
				best = currPressureScore;
				/*for (Node i : newPath) {
					
					System.out.println(i);
				}
				System.out.println();
				
				for (Node i : openedValves) {
					
					System.out.println(i);
				}
				System.out.println();*/
			}
				
			return;
		}
		
		if(allopened) return;
		
		
		int possibleScore = currPressureScore;
		for (Node i : allValves) {
			
			if(openedValves.contains(i) == false)
			{
				//int requiredminutes = distances.get(node.getName()).get(i.getName()) + 1;
				possibleScore += ((minute - 1) * i.getFlowRate());
			}
		}
		if(possibleScore < 1750) return; // BEST !!!!!!!!!!!!!
	
			
		
		List<String> options = new LinkedList<String>();
		for(String tunel : node.getTunnels())
		{
			options.add(tunel);
		}
		
		boolean opened = openedValves.contains(node);
		if(opened == false && node.getFlowRate() > 0)
		{
			options.add("OPEN");
		}
		
		
		for(String s : options)
		{
			if(minute == 25) System.out.println(s + " " + best);
				
			if(s.equals("OPEN"))
			{
				//node.setOpened(true);
				List<Node> newOpenedValves = new LinkedList<Node>(openedValves);
				newOpenedValves.add(node);
				int tmpScore = (minute - 1) * node.getFlowRate();
				getBestPath(node, minute - 1, currPressureScore + tmpScore, newOpenedValves, newPath);
			}
			else
			{
				Node next = new Node();
				for (Node i : allValves) {			
					if(i.getName().equals(s))
					{
						next = new Node(i);
						break;
					}
				}

				getBestPath(next, minute - 1, currPressureScore, openedValves, newPath);
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
