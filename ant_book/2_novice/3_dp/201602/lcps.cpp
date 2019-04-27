#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int n,m;
  string s,t;
  cin >> n >> m >> s >> t;

  vector< vector<int> > dp(n+1,vector<int>(m+1,0));

  for (int i = 0; i < n; i++) {
    for (int j = 0; j < m; j++) {
      dp[i+1][j] = max(dp[i+1][j], dp[i][j]);
      dp[i][j+1] = max(dp[i][j+1], dp[i][j]);
      if (s[i] == t[j]) {
        dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j]+1);  
      } else {
        dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j]);
      }
    }
  }

  cout << dp[n][m] << eol;
  
  return 0;
}
