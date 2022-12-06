
file = open('../input.txt', 'r')
count = 0
stackRead = False
stack = []
move = 0
fromStack = 0
toStack = 0

while True:

    # Get next line from file
    line = file.readline()

    # if line is empty
    # end of file is reached
    if not line:
        break

    if len(line) == 1:
        stackRead = True
        continue

    if len(stack) == 0:
        for i in range(int(len(line) / 4)):
            stack.append([])

    if stackRead is False:
        s = 0
        for i in range(1, len(line) - 1, 4):
            if line[i] != " ":
                stack[s].insert(0, line[i])
            s += 1
    else:
        move = int(line[4:].split("from")[0].strip())
        fromStack = int(line.split("from")[1].split("to")[0].strip())
        toStack = int(line.split("from")[1].split("to")[1].strip())

        for i in range(move):
            tmp = stack[fromStack - 1].pop()
            stack[toStack - 1].append(tmp)

res = ""
for s in stack:
    res += s[len(s) - 1]

print(res)

file.close()
