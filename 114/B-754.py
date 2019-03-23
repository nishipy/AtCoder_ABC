S = input()
diff = 753
for i in range(0, len(S)-2):
    if abs(753 - int(S[i:i+3])) < diff:
        diff = abs(753 - int(S[i:i+3]))
print(diff)
