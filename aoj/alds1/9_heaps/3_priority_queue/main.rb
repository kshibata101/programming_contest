def max_heapify(a, i)
  h = a.size
  l = 2*(i+1)-1
  r = 2*(i+1)

  if l < h && a[l] > a[i]
    largest = l
  else
    largest = i
  end
  largest = r if r < h && a[r] > a[largest]

  if largest != i
    a[i], a[largest] = a[largest], a[i]
    max_heapify(a, largest)
  end
end

def insert(s, k)
  s.push(k)
  k_index = s.size - 1
  while true
    j = (k_index - 1) / 2
    if k_index > 0 && s[k_index] > s[j]
      s[k_index], s[j] = s[j], s[k_index]
      k_index = j
    else
      break
    end
  end
end

def extract_max(s)
  puts s[0]
  s[0] = s.pop
  max_heapify(s, 0)
end

s = []
index = 0

loop do
  input = gets.split

  if input[0] == 'insert'
    insert(s, input[1].to_i)
  elsif input[0] == 'extract'
    extract_max(s)
  else
    break
  end
end
