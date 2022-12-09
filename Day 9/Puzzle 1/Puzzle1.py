
file = open('../input.txt', 'r')


hPos = (0, 0)
tPos = (0, 0)
visited = [(0, 0)]


while True:

    # Get next line from file
    line = file.readline()

    # if line is empty
    # end of file is reached
    if not line:
        break

    direction, step = line.split(" ")

    for i in range(int(step)):

        if direction == "R":
            hPos = (hPos[0] + 1, hPos[1])
            if (hPos[0]-tPos[0]) ** 2 + (hPos[1]-tPos[1]) ** 2 > 2:
                tPos = (hPos[0] - 1, hPos[1])
                visited.append(tPos)
        elif direction == "L":
            hPos = (hPos[0] - 1, hPos[1])
            if (hPos[0] - tPos[0]) ** 2 + (hPos[1] - tPos[1]) ** 2 > 2:
                tPos = (hPos[0] + 1, hPos[1])
                visited.append(tPos)
        elif direction == "U":
            hPos = (hPos[0], hPos[1] + 1)
            if (hPos[0] - tPos[0]) ** 2 + (hPos[1] - tPos[1]) ** 2 > 2:
                tPos = (hPos[0], hPos[1] - 1)
                visited.append(tPos)
        elif direction == "D":
            hPos = (hPos[0], hPos[1] - 1)
            if (hPos[0] - tPos[0]) ** 2 + (hPos[1] - tPos[1]) ** 2 > 2:
                tPos = (hPos[0], hPos[1] + 1)
                visited.append(tPos)


print(len(set(visited)))

file.close()
