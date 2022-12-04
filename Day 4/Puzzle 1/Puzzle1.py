
file = open('../input.txt', 'r')
count = 0

while True:

    # Get next line from file
    line = file.readline()

    # if line is empty
    # end of file is reached
    if not line:
        break

    firstElf, secondElf = line.split(",")

    firstElfList = [i for i in range(int(firstElf.split("-")[0]), int(firstElf.split("-")[1]) + 1)]
    secondElfList = [i for i in range(int(secondElf.split("-")[0]), int(secondElf.split("-")[1]) + 1)]

    if set(firstElfList).issubset(set(secondElfList)) or set(secondElfList).issubset(set(firstElfList)):
        count += 1

print(count)

file.close()
