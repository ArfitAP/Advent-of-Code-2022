#include <stdio.h>
#include <string.h>
#include <stdlib.h>

struct Number {
  int value;
  int order;
};

int countlines(FILE* file);
int swap(struct Number* sequence, int index1, int index2);
void move(struct Number* sequence, int numOfNumbers, int orderNumber);
void moveLeft(struct Number* sequence, int numOfNumbers, int orderNumber);
void moveRight(struct Number* sequence, int numOfNumbers, int orderNumber);
int nthElement(struct Number* sequence, int numOfNumbers, int n);

int main()
{
    char const* const fileName = "../input.txt";
    FILE* file = fopen(fileName, "r");
    char line[256];

	int numOfNumbers = countlines(file);
	
	struct Number* sequence = malloc(numOfNumbers * sizeof(struct Number));
	
	fseek(file, 0, SEEK_SET);
	
	int index = 0;
	while (fgets(line, sizeof(line), file)) {
        
        sequence[index].value = atoi(line);  
		sequence[index].order = index;
		
		index++;              
    }

	int i = 0;
	for(i = 0; i < numOfNumbers; i++)
	{
		move(sequence, numOfNumbers, i);
	}
	
	
	/*for(i = 0; i < numOfNumbers; i++)
	{
		printf("%i: %i \n", sequence[i].order, sequence[i].value);
	}*/
	
	
	printf("%i\n", nthElement(sequence, numOfNumbers, 1000) + nthElement(sequence, numOfNumbers, 2000) + nthElement(sequence, numOfNumbers, 3000));
		

    fclose(file);

    return 0;
}

int nthElement(struct Number* sequence, int numOfNumbers, int n)
{
	int i = 0;
	int currIndex = 0;
	for(i = 0; i < numOfNumbers; i++)
	{
		if(sequence[i].value == 0)
		{
			currIndex = i;
			break;
		}
	}

	int requestedIndex = currIndex;
	for(i = 0; i < n; i++)
	{
		if(requestedIndex == numOfNumbers - 1)
		{
			requestedIndex = 0;
		}
		else
		{
			requestedIndex++;
		}
	}
	
	return sequence[requestedIndex].value;
}

void move(struct Number* sequence, int numOfNumbers, int orderNumber)
{
	int i = 0;
	for(i = 0; i < numOfNumbers; i++)
	{
		if(sequence[i].order == orderNumber)
		{
			if(sequence[i].value < 0)
			{
				int n = 0;
				int numOfSwaps = -sequence[i].value;
				int ind1 = i;
				for(n = 0; n < numOfSwaps; n++)
				{				
					int ind2 = ind1 == 0 ? numOfNumbers - 1 : ind1 - 1;
					ind1 = swap(sequence, ind1, ind2);
				}
			}
			else if(sequence[i].value > 0)
			{
				int n = 0;
				int numOfSwaps = sequence[i].value;
				int ind1 = i;
				for(n = 0; n < numOfSwaps; n++)
				{				
					int ind2 = ind1 == numOfNumbers - 1 ? 0 : ind1 + 1;
					ind1 = swap(sequence, ind1, ind2);
				}
			}
				
			return;
		}
	}
}

void moveLeft(struct Number* sequence, int numOfNumbers, int orderNumber)
{
	int i = 0;
	for(i = 0; i < numOfNumbers; i++)
	{
		if(sequence[i].order == orderNumber)
		{
			int n = 0;
			int ind1 = i;
			for(n = 0; n < -sequence[i].value; n++)
			{				
				int ind2 = ind1 == 0 ? numOfNumbers - 1 : ind1 - 1;
				ind1 = swap(sequence, ind1, ind2);
			}
			
			return;
		}
	}
}

void moveRight(struct Number* sequence, int numOfNumbers, int orderNumber)
{
	int i = 0;
	for(i = 0; i < numOfNumbers; i++)
	{
		if(sequence[i].order == orderNumber)
		{
			int n = 0;
			int ind1 = i;
			for(n = 0; n < sequence[i].value; n++)
			{				
				int ind2 = ind1 == numOfNumbers - 1 ? 0 : ind1 + 1;
				ind1 = swap(sequence, ind1, ind2);
			}
			
			return;
		}
	}
}


int swap(struct Number* sequence, int index1, int index2)
{
	struct Number n1 = sequence[index1];
	sequence[index1] = sequence[index2];
	sequence[index2] = n1;
	
	return index2;
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

