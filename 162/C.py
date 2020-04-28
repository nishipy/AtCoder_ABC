import math
from functools import reduce

def gcd_list(numbers):
    return reduce(math.gcd, numbers)

K = int(input())
ans = 0
for i in range(1, K+1):
    for j in range(1, K+1):
        for k in range(1, K+1):
            ans += gcd_list([i, j, k])

print(ans)