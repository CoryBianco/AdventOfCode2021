with open('input.txt') as f:
  lines = f.readlines()
  increased = 0
  previous = []
  previousSum = None
  for line in lines:
    if len(previous) == 3:
      if sum(previous) < (sum(previous) - previous[0] + int(line)):
        print(sum(previous) - (sum(previous) - previous[0] + int(line)))
        increased += 1
      previous.pop(0)
    previous.append(int(line))
  print(increased)
