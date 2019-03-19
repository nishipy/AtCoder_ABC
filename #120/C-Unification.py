S = input()
n_zero = 0
n_one = 0
for i in S:
    if i == '0':
        n_zero += 1
    else:
        n_one += 1
print(min(n_one, n_zero)*2)