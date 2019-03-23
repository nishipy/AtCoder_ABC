import numpy as np
import collections
n = int(input())
a_n = np.array(input().split(),dtype='int')
diff = a_n[1:] - a_n[:-1]
if len(set(diff))==2:
    print(0)
elif len(set(diff))==1:
    print(int(n/2))
else:
    odd_th = collections.Counter(a_n[1::2]).most_common()
    eve_th = collections.Counter(a_n[0::2]).most_common()
    if odd_th[0][0]==eve_th[0][0]:
        if odd_th[0][1]>eve_th[0][1]:
            odd_ = odd_th[0][1]
            eve_ = eve_th[1][1]
        elif odd_th[0][1]<eve_th[0][1]:
            odd_ = odd_th[1][1]
            eve_ = eve_th[0][1]
        else:
            if odd_th[1][1]>=eve_th[1][1]:
                odd_ = odd_th[1][1]
                eve_ = eve_th[0][1]
            else:
                odd_ = odd_th[0][1]
                eve_ = eve_th[1][1] 
    else:
            odd_ = odd_th[0][1]
            eve_ = eve_th[0][1]
    print(n-odd_-eve_)
