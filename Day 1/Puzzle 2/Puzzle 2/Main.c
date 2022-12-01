#define _CRT_SECURE_NO_DEPRECATE
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void sortSums(int *top1, int *top2, int *top3);

int main()
{
    char const* const fileName = "../../input.txt";
    FILE* file = fopen(fileName, "r");
    char line[256];

    int currsum = 0;
    int top1sum = 0;
    int top2sum = 0;
    int top3sum = 0;

    while (fgets(line, sizeof(line), file)) {

        //printf("%s", line);
        //printf("%i", strlen(line));

        if (strlen(line) > 1)
        {
            currsum += atoi(line);
        }
        else
        {
            if (currsum > top3sum)
            {
                top3sum = currsum;
                sortSums(&top1sum, &top2sum, &top3sum);
            }
            currsum = 0;
        }
    }

    printf("%i", top1sum + top2sum + top3sum);

    fclose(file);

    return 0;
}

void sortSums(int *top1, int *top2, int *top3)
{
    if (*top3 >= *top2 && *top3 >= *top1)
    {
        int tmp = *top1;
        *top1 = *top3;
        *top3 = tmp;

        if (*top3 > *top2)
        {
            int tmp = *top2;
            *top2 = *top3;
            *top3 = tmp;
        }
    }
    else if (*top2 > *top1)
    {
        int tmp = *top1;
        *top1 = *top2;
        *top2 = tmp;

        if (*top3 > *top2)
        {
            int tmp = *top2;
            *top2 = *top3;
            *top3 = tmp;
        }
    }
    else
    {
        if (*top3 > *top2)
        {
            int tmp = *top2;
            *top2 = *top3;
            *top3 = tmp;
        }
    }
}