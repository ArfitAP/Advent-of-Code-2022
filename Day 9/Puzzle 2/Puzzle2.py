
file = open('../input.txt', 'r')

knotPos = [(0, 0), (0, 0), (0, 0), (0, 0), (0, 0), (0, 0), (0, 0), (0, 0), (0, 0), (0, 0)]
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
            knotPos[0] = (knotPos[0][0] + 1, knotPos[0][1])
        elif direction == "L":
            knotPos[0] = (knotPos[0][0] - 1, knotPos[0][1])
        elif direction == "U":
            knotPos[0] = (knotPos[0][0], knotPos[0][1] + 1)
        elif direction == "D":
            knotPos[0] = (knotPos[0][0], knotPos[0][1] - 1)

        for k in range(9):
            if (knotPos[k][0]-knotPos[k+1][0]) ** 2 + (knotPos[k][1]-knotPos[k+1][1]) ** 2 == 4:
                if knotPos[k][0] < knotPos[k+1][0]:
                    knotPos[k+1] = (knotPos[k+1][0] - 1, knotPos[k+1][1])
                elif knotPos[k][0] > knotPos[k+1][0]:
                    knotPos[k+1] = (knotPos[k+1][0] + 1, knotPos[k+1][1])
                elif knotPos[k][1] < knotPos[k+1][1]:
                    knotPos[k+1] = (knotPos[k+1][0], knotPos[k+1][1] - 1)
                elif knotPos[k][1] > knotPos[k+1][1]:
                    knotPos[k+1] = (knotPos[k+1][0], knotPos[k+1][1] + 1)

            elif (knotPos[k][0] - knotPos[k + 1][0]) ** 2 + (knotPos[k][1] - knotPos[k + 1][1]) ** 2 >= 5:
                if knotPos[k][0] < knotPos[k+1][0] and knotPos[k][1] < knotPos[k+1][1]:
                    knotPos[k + 1] = (knotPos[k + 1][0] - 1, knotPos[k + 1][1] - 1)
                elif knotPos[k][0] < knotPos[k+1][0] and knotPos[k][1] > knotPos[k+1][1]:
                    knotPos[k + 1] = (knotPos[k + 1][0] - 1, knotPos[k + 1][1] + 1)
                elif knotPos[k][0] > knotPos[k+1][0] and knotPos[k][1] < knotPos[k+1][1]:
                    knotPos[k + 1] = (knotPos[k + 1][0] + 1, knotPos[k + 1][1] - 1)
                elif knotPos[k][0] > knotPos[k+1][0] and knotPos[k][1] > knotPos[k+1][1]:
                    knotPos[k + 1] = (knotPos[k + 1][0] + 1, knotPos[k + 1][1] + 1)

            if k == 8:
                visited.append(knotPos[k + 1])


print(len(set(visited)))

file.close()
