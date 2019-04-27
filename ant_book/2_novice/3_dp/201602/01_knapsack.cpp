#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int n,W;
  cin >> n >> W;

  vector<int> w(n),v(n);
  for (int i = 0; i < n; i++) {
    cin >> w[i] >> v[i];
  }

  vector< vector<int> > dp(n+1, vector<int>(W+1,-1));
  dp[0][0] = 0;
  
  for (int i = 0; i < n; i++) {
    for (int j = 0; j <= W; j++) {
      if (dp[i][j] == -1) continue;
      
      dp[i+1][j] = max(dp[i+1][j], dp[i][j]);
      if (j+w[i] <= W) {
        dp[i+1][j+w[i]] = max(dp[i+1][j+w[i]], dp[i][j] + v[i]);
      }
    }
  }

  int ans = -1;
  for (int j = 0; j <= W; j++) {
    ans = max(ans, dp[n][j]);
  }
  cout << ans << eol;
  
  return 0;
}
