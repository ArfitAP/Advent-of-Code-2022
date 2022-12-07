
class Node(object):
    # Initializing to None
    def __init__(self):
        self.leafs = []
        self.folder = True
        self.size = 0
        self.path = ["/"]


def getNodeByPath(node, path):

    if node.path == path:
        return node
    else:
        newFolder = path[0:len(node.path) + 1]
        for leaf in node.leafs:
            if leaf.path == newFolder:
                return getNodeByPath(leaf, path)

        newleaf = Node()
        newleaf.path = newFolder
        node.leafs.append(newleaf)
        return newleaf

def getFolderSizes(node):

    if not node.folder:
        return node.size

    sum = 0
    for l in node.leafs:
        size = getFolderSizes(l)
        sum += size

    pathstring = getpathfromstack(node.path)
    if node.folder and (pathstring not in sizes.keys()):
        sizes[pathstring] = sum

    return sum

def getpathfromstack(pathStack):
    res = ""
    for s in pathStack:
        res += (s + "_")
    return res


file = open('../input.txt', 'r')

tree = Node()

currentNode = None
currentFolder = ""
pathStack = []

sizes = {}

while True:

    # Get next line from file
    line = file.readline()

    # if line is empty
    # end of file is reached
    if not line:
        break

    if line[0] == "$":
        if line.split(" ")[1] == "cd":
            folderCmd = line.split(" ")[2].strip()

            if folderCmd == "..":
                currentFolder = pathStack.pop()
            elif folderCmd == "/":
                currentFolder = "/"
                pathStack = ["/"]
            else:
                currentFolder = folderCmd
                pathStack.append(currentFolder)

        elif line.split(" ")[1].strip() == "ls":
            currentNode = getNodeByPath(tree, pathStack)

    else:
        added = False
        filePath = pathStack.copy()
        filePath.append(line.split(" ")[1].strip())

        for leaf in currentNode.leafs:
            if leaf.path == filePath:
                added = True

        if not added:
            newFolder = Node()
            newFolder.path = filePath
            if line.split(" ")[0] == "dir":
                newFolder.folder = True
            else:
                newFolder.folder = False
                newFolder.size = int(line.split(" ")[0])
            currentNode.leafs.append(newFolder)


getFolderSizes(tree)

filteredDict = dict()

for (key, value) in sizes.items():
    if value <= 100000:
        filteredDict[key] = value

sum = 0
for values in filteredDict.values():
    sum += values

print(sum)


file.close()
