N, M = map(int,input().split())
favorite_lists = []
merged = []
count = 0
for _ in range(0, N):
    favorite_lists.append(input().split()[1:])
for l in favorite_lists:
    merged.extend(l)
for key in range(0, M+1):
    if merged.count(str(key))==N:
        count+=1
print(count)
