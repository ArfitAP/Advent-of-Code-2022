
public class Position {
	public int row;
	public int col;
	public int height;
	public int steps;
	
	public Position(int row, int col, int height, int steps) {
		super();
		this.row = row;
		this.col = col;
		this.height = height;
		this.steps = steps;
	}

	public int getRow() {
		return row;
	}

	public void setRow(int row) {
		this.row = row;
	}

	public int getCol() {
		return col;
	}

	public void setCol(int col) {
		this.col = col;
	}

	public int getHeight() {
		return height;
	}

	public void setHeight(int height) {
		this.height = height;
	}

	public int getSteps() {
		return steps;
	}

	public void setSteps(int steps) {
		this.steps = steps;
	}
	
	
	@Override
    public String toString() {
        return this.row + ", " + this.col;
    }
	
	@Override
    public int hashCode()
    {
        return this.row * 1000 + this.col;
    }
	
	@Override
    public boolean equals(Object o) {
 
        if (o == this) {
            return true;
        }
 
        if (!(o instanceof Position)) {
            return false;
        }
         
        Position c = (Position) o;
         
        return Integer.compare(row, c.row) == 0
                && Integer.compare(col, c.col) == 0;
    }
}
