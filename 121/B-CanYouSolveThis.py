N_M_C = [int(x.strip()) for x in input().split()]
B = [int(x.strip()) for x in input().split()]
A = []
for i in range(0, N_M_C [0]):
    A.append([int(x.strip()) for x in input().split()])
count = 0
for row_A in A:
    result = 0
    for i in range(0, len(B)):
        result += row_A[i]*B[i]
    if result > - N_M_C[2]:
        count += 1
print(count)
