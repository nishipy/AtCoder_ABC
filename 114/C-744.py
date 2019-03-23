import itertools
import numpy as np
N = input()
digit = len(N)
ans = 0
if int(N) < 357:
    print(ans)
else:
    for i in range(3, len(N)):
        ans += 3**i - 3 * 2**i + 3
    for v in itertools.product('753', repeat=digit):
        if (len(set(list(v)))==3) & (int(''.join(list(v))) <= int(N)):
            ans +=1
    print(ans)
