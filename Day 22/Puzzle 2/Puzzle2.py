
def getNextPositionTest():

    if currentPosition[2] == 0:
        if len(grid[currentPosition[0] - 1]) - 1 > currentPosition[1]:
            return currentPosition[0], currentPosition[1] + 1, currentPosition[2]
        elif 0 <= currentPosition[0] <= 4:
            return 13 - currentPosition[0], currentPosition[1] + 4, 2
        elif 5 <= currentPosition[0] <= 8:
            return 9, 8 - currentPosition[0] + 13, 1
        elif 9 <= currentPosition[0] <= 12:
            return 13 - currentPosition[0], currentPosition[1] - 4, 2
    elif currentPosition[2] == 1:
        if len(grid) > currentPosition[0] and len(grid[currentPosition[0]]) - 1 > currentPosition[1] - 1 and grid[currentPosition[0]][currentPosition[1] - 1] != " ":
            return currentPosition[0] + 1, currentPosition[1], currentPosition[2]
        elif 0 <= currentPosition[1] <= 4:
            return 12, 4 - currentPosition[1] + 9, 3
        elif 5 <= currentPosition[1] <= 8:
            return 8 - currentPosition[1] + 9, 9, 0
        elif 9 <= currentPosition[0] <= 12:
            return 8, 12 - currentPosition[1] + 1, 3
        elif 13 <= currentPosition[0] <= 16:
            return 16 - currentPosition[1] + 5, 1, 0
    if currentPosition[2] == 2:
        if currentPosition[1] > 1 and grid[currentPosition[0] - 1][currentPosition[1] - 2] != " ":
            return currentPosition[0], currentPosition[1] - 1, currentPosition[2]
        elif 0 <= currentPosition[0] <= 4:
            return 5, currentPosition[0] + 4, 1
        elif 5 <= currentPosition[0] <= 8:
            return 12, 8 - currentPosition[0] + 13, 3
        elif 9 <= currentPosition[0] <= 12:
            return 8, 12 - currentPosition[0] + 5, 3
    elif currentPosition[2] == 3:
        if currentPosition[0] > 1 and len(grid[currentPosition[0] - 2]) - 1 > currentPosition[1] - 1 and grid[currentPosition[0] - 2][currentPosition[1] - 1] != " ":
            return currentPosition[0] - 1, currentPosition[1], currentPosition[2]
        elif 0 <= currentPosition[1] <= 4:
            return currentPosition[1] - 4, 9, 1
        elif 5 <= currentPosition[1] <= 8:
            return currentPosition[1] - 4, 9, 0
        elif 9 <= currentPosition[0] <= 12:
            return 5, 12 - currentPosition[1] + 1, 1
        elif 13 <= currentPosition[0] <= 16:
            return 16 - currentPosition[1] + 5, 12, 2


def getNextPositionProd():

    if currentPosition[2] == 0:
        if len(grid[currentPosition[0] - 1]) - 1 > currentPosition[1]:
            return currentPosition[0], currentPosition[1] + 1, currentPosition[2]
        elif 0 <= currentPosition[0] <= 50:
            return 50 - currentPosition[0] + 101, 100, 2
        elif 51 <= currentPosition[0] <= 100:
            return 50, currentPosition[0] + 50, 3
        elif 101 <= currentPosition[0] <= 150:
            return 150 - currentPosition[0] + 1, 150, 2
        elif 151 <= currentPosition[0] <= 200:
            return 150, currentPosition[0] - 100, 3
    elif currentPosition[2] == 1:
        if len(grid) > currentPosition[0] and len(grid[currentPosition[0]]) - 1 > currentPosition[1] - 1 and grid[currentPosition[0]][currentPosition[1] - 1] != " ":
            return currentPosition[0] + 1, currentPosition[1], currentPosition[2]
        elif 0 <= currentPosition[1] <= 50:
            return 1, currentPosition[1] + 100, 1
        elif 51 <= currentPosition[1] <= 100:
            return currentPosition[1] + 100, 50, 2
        elif 101 <= currentPosition[1] <= 150:
            return currentPosition[1] - 50, 100, 2
    if currentPosition[2] == 2:
        if currentPosition[1] > 1 and grid[currentPosition[0] - 1][currentPosition[1] - 2] != " ":
            return currentPosition[0], currentPosition[1] - 1, currentPosition[2]
        elif 0 <= currentPosition[0] <= 50:
            return 50 - currentPosition[0] + 101, 1, 0
        elif 51 <= currentPosition[0] <= 100:
            return 101, currentPosition[0] - 50, 1
        elif 101 <= currentPosition[0] <= 150:
            return 150 - currentPosition[0] + 1, 51, 0
        elif 151 <= currentPosition[0] <= 200:
            return 1, currentPosition[0] - 100, 1
    elif currentPosition[2] == 3:
        if currentPosition[0] > 1 and len(grid[currentPosition[0] - 2]) - 1 > currentPosition[1] - 1 and grid[currentPosition[0] - 2][currentPosition[1] - 1] != " ":
            return currentPosition[0] - 1, currentPosition[1], currentPosition[2]
        elif 0 <= currentPosition[1] <= 50:
            return currentPosition[1] + 50, 51, 0
        elif 51 <= currentPosition[1] <= 100:
            return currentPosition[1] + 100, 1, 0
        elif 101 <= currentPosition[1] <= 150:
            return 200, currentPosition[1] - 100, 3


file = open('../input.txt', 'r')

mapRead = False
instructions = ""
instructionList = []
grid = []
currentPosition = (0, 0, -1)
nextPosition = (0, 0, -1)

while True:

    # Get next line from file
    line = file.readline()

    # if line is empty
    # end of file is reached
    if not line:
        break

    if len(line) <= 1:
        mapRead = True
        continue

    if not mapRead:
        grid.append(line)
    else:
        instructions = line

for i in range(len(grid[0])):
    if grid[0][i] == ".":
        currentPosition = (1, i+1, 0)
        break


tmpIndex = 0
for i in range(len(instructions)):
    if instructions[i] == "R" or instructions[i] == "L":
        instructionList.append(int(instructions[tmpIndex:i]))
        instructionList.append(instructions[i])
        tmpIndex = i + 1
instructionList.append(int(instructions[tmpIndex:]))


for i in range(len(instructionList)):

    if instructionList[i] == "R":
        currentPosition = (currentPosition[0], currentPosition[1], (currentPosition[2] + 1) % 4)
    elif instructionList[i] == "L":
        currentPosition = (currentPosition[0], currentPosition[1], (currentPosition[2] - 1) % 4)
    else:

        for j in range(instructionList[i]):
            # nextPosition = getNextPositionTest()
            nextPosition = getNextPositionProd()
            if grid[nextPosition[0] - 1][nextPosition[1] - 1] == "#":
                break
            else:
                currentPosition = nextPosition


print(1000 * currentPosition[0] + 4 * currentPosition[1] + currentPosition[2])

file.close()
