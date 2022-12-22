#include <stdio.h>
#include <string.h>
#include <stdlib.h>

struct Cube {
  int x;
  int y;
  int z;
};

struct Air {
  int x;
  int y;
  int z;
  int trapped;
};

int countlines(FILE* file);
int isAdjecent(struct Cube c1, struct Cube c2);
int isAirAdjecent(struct Air c1, struct Air c2);
int isAdjecentByCord(struct Cube c1, int x, int y, int z);
int isAir(struct Cube* cubes, int numOfCubes, int x, int y, int z);
int MarkAdjecent(struct Air* airs, int numOfAirs, struct Air a1);

int main()
{
    char const* const fileName = "../input.txt";
    FILE* file = fopen(fileName, "r");
    char line[256];

	int numOfCubes = countlines(file);
	int c = 0;
	int xmin = INT_MAX;
	int xmax = INT_MIN;
	int ymin = INT_MAX;
	int ymax = INT_MIN;
	int zmin = INT_MAX;
	int zmax = INT_MIN;
	
	struct Cube* cubes = malloc(numOfCubes * sizeof(struct Cube));
	
	fseek(file, 0, SEEK_SET);
	
	while (fgets(line, sizeof(line), file)) {
        
        char* token = strtok(line, ",");
        cubes[c].x = atoi(token);
        if(cubes[c].x > xmax) xmax = cubes[c].x;
        if(cubes[c].x < xmin) xmin = cubes[c].x;
        
        token = strtok(NULL, ",");
        cubes[c].y = atoi(token);
        if(cubes[c].y > ymax) ymax = cubes[c].y;
        if(cubes[c].y < ymin) ymin = cubes[c].y;
        
        token = strtok(NULL, ",");
        cubes[c].z = atoi(token);  
        if(cubes[c].z > zmax) zmax = cubes[c].z;
        if(cubes[c].z < zmin) zmin = cubes[c].z;
		
		c++;              
    }

	int i = 0;
	int j = 0;
	int k = 0;
	int cnum = 0;
	int totalSumOfFreeSides = 0;
	for(i = 0; i < numOfCubes; i++)
	{
		int freeSides = 6;
		for(j = 0; j < numOfCubes; j++)
		{
			if(i != j && isAdjecent(cubes[i], cubes[j]) == 1)
			{
				freeSides--;
			}
		}
		totalSumOfFreeSides += freeSides;
	}
	
	
	struct Air* airs = malloc((xmax-xmin+1)*(ymax-ymin+1)*(zmax-zmin+1) * sizeof(struct Air));
	int airNum = 0;
	for(i = xmin; i <= xmax; i++)
	{
		for(j = ymin; j <= ymax; j++)
		{
			for(k = zmin; k <= zmax; k++)
			{
				if(isAir(cubes, numOfCubes, i, j ,k) == 1)
				{
					
					if(i == xmin || i == xmax || j == ymin || j == ymax || k == zmin || k == zmax) 
					{								
						airs[airNum].x = i;
						airs[airNum].y = j;
						airs[airNum].z = k;
						airs[airNum].trapped = 0;
					}
					else
					{
						airs[airNum].x = i;
						airs[airNum].y = j;
						airs[airNum].z = k;
						airs[airNum].trapped = 1;
					}
					airNum++;
				}				
			}
		}
	}
	
	for(i = 0; i < airNum; i++)
	{
		if(airs[i].trapped == 0)
		{
			MarkAdjecent(airs, airNum, airs[i]);
		}
	}
	
	int totalSumOfAirSides = 0;
	for(i = 0; i < airNum; i++)
	{
		if(airs[i].trapped == 1)
		{
			int freeSides = 6;
			for(j = 0; j < airNum; j++)
			{
				if(i != j && isAirAdjecent(airs[i], airs[j]) == 1)
				{
					freeSides--;
				}
			}
			totalSumOfAirSides += freeSides;
		}
	}
	
	printf("%i \n", totalSumOfFreeSides - totalSumOfAirSides);

    fclose(file);

    return 0;
}

int MarkAdjecent(struct Air* airs, int numOfAirs, struct Air a1)
{
	int i = 0;
		
	for(i = 0; i < numOfAirs; i++)
	{
		if(isAirAdjecent(a1, airs[i]) && airs[i].trapped == 1)
		{
			airs[i].trapped = 0;
			MarkAdjecent(airs, numOfAirs, airs[i]);
		}
	}

}

int isAirAdjecent(struct Air c1, struct Air c2)
{
	if(c1.x == c2.x && c1.y == c2.y && (c1.z == c2.z - 1 || c1.z == c2.z + 1) )
	{
		return 1;
	}
	if(c1.x == c2.x && c1.z == c2.z && (c1.y == c2.y - 1 || c1.y == c2.y + 1) )
	{
		return 1;
	}
	if(c1.y == c2.y && c1.z == c2.z && (c1.x == c2.x - 1 || c1.x == c2.x + 1) )
	{
		return 1;
	}
	return 0;
}

int isAdjecent(struct Cube c1, struct Cube c2)
{
	if(c1.x == c2.x && c1.y == c2.y && (c1.z == c2.z - 1 || c1.z == c2.z + 1) )
	{
		return 1;
	}
	if(c1.x == c2.x && c1.z == c2.z && (c1.y == c2.y - 1 || c1.y == c2.y + 1) )
	{
		return 1;
	}
	if(c1.y == c2.y && c1.z == c2.z && (c1.x == c2.x - 1 || c1.x == c2.x + 1) )
	{
		return 1;
	}
	return 0;
}

int isAdjecentByCord(struct Cube c1, int x, int y, int z)
{
	struct Cube c2;
	c2.x = x;
	c2.y = y;
	c2.z = z;
	
	if(c1.x == c2.x && c1.y == c2.y && (c1.z == c2.z - 1 || c1.z == c2.z + 1) )
	{
		return 1;
	}
	if(c1.x == c2.x && c1.z == c2.z && (c1.y == c2.y - 1 || c1.y == c2.y + 1) )
	{
		return 1;
	}
	if(c1.y == c2.y && c1.z == c2.z && (c1.x == c2.x - 1 || c1.x == c2.x + 1) )
	{
		return 1;
	}
	return 0;
}

int isAir(struct Cube* cubes, int numOfCubes, int x, int y, int z)
{
	int i = 0;
	for(i = 0; i < numOfCubes; i++)
	{
		if(cubes[i].x == x && cubes[i].y == y && cubes[i].z == z)
		{
			return 0;	
		}
	}
	return 1;
}

int countlines(FILE* file)
{
  	int lines=0;
  	char ch;
  	
  	while(!feof(file))
	{
	  ch = fgetc(file);
	  if(ch == '\n')
	  {
	    lines++;
	  }
	}
	
	return lines;
}
