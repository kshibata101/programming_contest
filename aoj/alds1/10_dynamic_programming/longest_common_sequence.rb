n = $stdin.gets.chomp.to_i
n.times do
  x = $stdin.gets.chomp
  y = $stdin.gets.chomp

  dp = Array.new(x.size+1).map{Array.new(y.size+1, 0)}
  for i in 0...x.size do
    for j in 0...y.size do
      dp[i+1][j+1] = [dp[i][j+1], dp[i+1][j]].max
      if x[i] == y[j]
        dp[i+1][j+1] = [dp[i+1][j+1], dp[i][j] + 1].max
      end
    end
  end
  puts dp[x.size][y.size]
end



