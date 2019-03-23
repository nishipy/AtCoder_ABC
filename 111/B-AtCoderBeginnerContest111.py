N = input()
if int(N) <= 111:
    print(111)
else:
    for m in range(0, 9):
        left = 112 + 111*m
        right = 222 + 111*m
        if (left<=int(N))&(int(N)<=right):
            print(right)
            break