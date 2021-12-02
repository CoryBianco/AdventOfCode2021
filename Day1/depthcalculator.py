with open('input.txt') as f:
  lines = f.readlines()
  increased = 0
  previous = None
  for line in lines:
    if previous:
      if previous < int(line):
        increased += 1
    previous = int(line)
  print(increased)
