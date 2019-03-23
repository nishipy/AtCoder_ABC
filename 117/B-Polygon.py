N = int(input())
length_list = input().split()
length_list = [int(i) for i in length_list]
length_list.sort()
if length_list[-1] < sum(length_list[0:-1]):
    print('Yes')
else:
    print('No')