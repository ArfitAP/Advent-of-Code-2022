#include <stdio.h>
#include <string.h>
#include <stdlib.h>

struct Cube {
  int x;
  int y;
  int z;
};

int countlines(FILE* file);
int isAdjecent(struct Cube c1, struct Cube c2);

int main()
{
    char const* const fileName = "../input.txt";
    FILE* file = fopen(fileName, "r");
    char line[256];

	int numOfCubes = countlines(file);
	int c = 0;
	
	struct Cube* cubes = malloc(numOfCubes * sizeof(struct Cube));
	
	fseek(file, 0, SEEK_SET);
	
	while (fgets(line, sizeof(line), file)) {
        
        char* token = strtok(line, ",");
        cubes[c].x = atoi(token);
        
        token = strtok(NULL, ",");
        cubes[c].y = atoi(token);
        
        token = strtok(NULL, ",");
        cubes[c].z = atoi(token);  
		
		c++;              
    }

	int i = 0;
	int j = 0;
	/*for(i = 0; i < numOfCubes; i++)
	{
		printf("%i %i %i \n", cubes[i].x, cubes[i].y, cubes[i].z);
	}*/
	
	int totalSum = 0;
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
		totalSum += freeSides;
	}
	
	printf("%i \n", totalSum);

    fclose(file);

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
