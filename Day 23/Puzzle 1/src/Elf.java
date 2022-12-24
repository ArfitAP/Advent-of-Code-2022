
public class Elf {
	public int x;
	public int y;
	public int proposedX;
	public int proposedY;
	public boolean noMove;
	
	public Elf(int x, int y) {
		this.x = x;
		this.y = y;
		this.proposedX = 0;
		this.proposedY = 0;
		this.noMove = false;
	}
	
	public void propose(int x, int y) {
		this.proposedX = x;
		this.proposedY = y;
		this.noMove = false;
	}
	
	public void move() {
		this.x = proposedX;
		this.y = proposedY;
	}

	public int getX() {
		return x;
	}

	public void setX(int x) {
		this.x = x;
	}

	public int getY() {
		return y;
	}

	public void setY(int y) {
		this.y = y;
	}

	public int getProposedX() {
		return proposedX;
	}

	public void setProposedX(int proposedX) {
		this.proposedX = proposedX;
	}

	public int getProposedY() {
		return proposedY;
	}

	public void setProposedY(int proposedY) {
		this.proposedY = proposedY;
	}
	
	public boolean isNoMove() {
		return noMove;
	}

	public void setNoMove(boolean noMove) {
		this.noMove = noMove;
	}

	@Override
	public String toString() {
		return "(" + x + ", " + y + ")";
	}
	
	
}


