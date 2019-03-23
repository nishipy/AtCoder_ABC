sum_gift = 0
Gift_lists = []
N = int(input())
for i in range(0, N):
    Gift_lists.append([x.strip() for x in input().split()])

for l in Gift_lists:
    if l[1]=='JPY':
        sum_gift += int(l[0])
    else:
        sum_gift += float(l[0])*380000.0
print(sum_gift)