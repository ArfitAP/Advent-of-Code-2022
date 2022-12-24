
def moveUp(steps):
    lastPos = currentPosition
    while steps > 0:
        for j in range(len(grid)):

            if len(grid[(currentPosition[0] - 2 - j) % len(grid)]) - 1 < currentPosition[1]:
                continue

            elem = grid[(currentPosition[0] - 2 - j) % len(grid)][currentPosition[1] - 1]
            if elem == " ":
                continue
            elif elem == ".":
                steps = steps - 1
                lastPos = (currentPosition[0] - 2 - j) % len(grid) + 1, currentPosition[1], currentPosition[2]
            elif elem == "#":
                return lastPos

            if steps == 0:
                return (currentPosition[0] - 2 - j) % len(grid) + 1, currentPosition[1], currentPosition[2]


def moveDown(steps):
    lastPos = currentPosition
    while steps > 0:
        for j in range(len(grid)):

            if len(grid[(currentPosition[0] + j) % len(grid)]) - 1 < currentPosition[1]:
                continue

            elem = grid[(currentPosition[0] + j) % len(grid)][currentPosition[1] - 1]
            if elem == " ":
                continue
            elif elem == ".":
                steps = steps - 1
                lastPos = (currentPosition[0] + j) % len(grid) + 1, currentPosition[1], currentPosition[2]
            elif elem == "#":
                return lastPos

            if steps == 0:
                return (currentPosition[0] + j) % len(grid) + 1, currentPosition[1], currentPosition[2]


def moveRight(steps):
    lastPos = currentPosition
    while steps > 0:
        for j in range(len(grid[currentPosition[0] - 1]) - 1):
            elem = grid[currentPosition[0] - 1][(currentPosition[1] + j) % (len(grid[currentPosition[0] - 1]) - 1)]
            #print((currentPosition[1] + j) % (len(grid[currentPosition[0]]) - 1))
            if elem == " ":
                continue
            elif elem == ".":
                steps = steps - 1
                lastPos = currentPosition[0], (currentPosition[1] + j) % (len(grid[currentPosition[0] - 1]) - 1) + 1, currentPosition[2]
            elif elem == "#":
                return lastPos

            if steps == 0:
                return currentPosition[0], (currentPosition[1] + j) % (len(grid[currentPosition[0] - 1]) - 1) + 1, currentPosition[2]

def moveLeft(steps):

    lastPos = currentPosition
    while steps > 0:
        for j in range(len(grid[currentPosition[0] - 1]) - 1):
            elem = grid[currentPosition[0] - 1][(currentPosition[1] - 2 - j) % (len(grid[currentPosition[0] - 1]) - 1)]
            #print((currentPosition[1] - 2 - j) % (len(grid[currentPosition[0]]) - 1))
            if elem == " ":
                continue
            elif elem == ".":
                steps = steps - 1
                lastPos = currentPosition[0], (currentPosition[1] - 2 - j) % (len(grid[currentPosition[0] - 1]) - 1) + 1, currentPosition[2]
            elif elem == "#":
                return lastPos

            if steps == 0:
                return currentPosition[0], (currentPosition[1] - 2 - j) % (len(grid[currentPosition[0] - 1]) - 1) + 1, currentPosition[2]


file = open('../input.txt', 'r')

mapRead = False
instructions = ""
instructionList = []
grid = []
currentPosition = (0, 0, -1)

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


#for i in range(len(grid)):
#    print(grid[i])

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
        if currentPosition[2] == 0:
            currentPosition = moveRight(instructionList[i])
        elif currentPosition[2] == 1:
            currentPosition = moveDown(instructionList[i])
        elif currentPosition[2] == 2:
            currentPosition = moveLeft(instructionList[i])
        elif currentPosition[2] == 3:
            currentPosition = moveUp(instructionList[i])

#print(currentPosition)

print(1000 * currentPosition[0] + 4 * currentPosition[1] + currentPosition[2])

file.close()
