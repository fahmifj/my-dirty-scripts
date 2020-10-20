print("Please don't judge my python code")
print(".. also the variable naming")
filename = "dir.txt"
flists = []
keys = []
data = []

with open(filename) as f :
    content = f.read().splitlines()

for line in content:
    str = line[2:].replace('/', '')
    flists.append(str)

for flist in flists:
    keys.append(int(flist[1:]))
    data.append(flist[0])
    #print("{keys}->{data}\n".format(keys=keys,data=data))
    #print("%s --> %s" %(keys, data))
    #sort_it(key, data)

def sort_it(keys, data):
    for i in range(len(keys)):
        min_id = i
        for j in range(i+1, len(keys)):
            if int(keys[min_id]) > int(keys[j]):
                min_id = j
        keys[i], keys[min_id] = keys[min_id], keys[i]
        data[i], data[min_id] = data[min_id], data[i]
#print("Is it sorted")
sort_it(keys,data)

Flag = ""

for i in range (len(keys)):
    print("%d --> %s " %(keys[i], data[i]))
    #print("%s" %(data[i]))
    Flag = Flag + data[i]

import base64
print("Flag")
print(base64.b64decode(Flag))