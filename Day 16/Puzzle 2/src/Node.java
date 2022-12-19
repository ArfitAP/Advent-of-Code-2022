import java.util.LinkedList;
import java.util.List;

public class Node {
	public String name;
	public int flowRate;
	public List<String> tunnels;
	public boolean opened;
	
	public Node(Node node) {
		super();
		this.name = node.getName();
		this.flowRate = node.getFlowRate();
		this.tunnels = node.getTunnels();
		this.opened = node.isOpened();
	}
	
	public Node(String name, int flowRate, List<String> tunnels) {
		super();
		this.name = name;
		this.flowRate = flowRate;
		this.tunnels = tunnels;
		this.opened = false;
	}
	
	public Node() {
		super();
		tunnels = new LinkedList<String>();
		this.opened = false;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public int getFlowRate() {
		return flowRate;
	}

	public void setFlowRate(int flowRate) {
		this.flowRate = flowRate;
	}

	public List<String> getTunnels() {
		return tunnels;
	}

	public void setTunnels(List<String> tunnels) {
		this.tunnels = tunnels;
	}

	public boolean isOpened() {
		return opened;
	}

	public void setOpened(boolean opened) {
		this.opened = opened;
	}

	@Override
    public String toString() {
		String res = this.name + ": " + this.flowRate + " (";
		for(String tunnel : tunnels)
		{
			res += tunnel + " ";
		}
		res += ")";
        return res;
    }
	
	
	@Override
    public boolean equals(Object o) {
 
        if (o == this) {
            return true;
        }
 
        if (!(o instanceof Node)) {
            return false;
        }
         
        Node n = (Node) o;
         
        return name.equals(n.name);
    }
}


