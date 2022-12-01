#define _CRT_SECURE_NO_DEPRECATE
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main()
{
    char const* const fileName = "../input.txt";
    FILE* file = fopen(fileName, "r");
    char line[256];
    int currsum = 0;
    int maxsum = 0;

    while (fgets(line, sizeof(line), file)) {

        //printf("%s", line);
        //printf("%i", strlen(line));

        if (strlen(line) > 1)
        {
            currsum += atoi(line);
        }
        else
        {
            if (currsum > maxsum) maxsum = currsum;
            currsum = 0;
        }
    }

    printf("%i", maxsum);

    fclose(file);

    return 0;
}
