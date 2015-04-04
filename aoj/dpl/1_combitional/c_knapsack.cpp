#include <iostream>
#define eol '\n';
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int N,W;
  cin >> N >> W;
  int* v = new int[N+1];
  int* w = new int[N+1];

  for (int i = 1; i <= N; i++) {
    cin >> v[i] >> w[i];
  }

  int** dp = new int*[N+1];
  for (int i = 0; i <= N; i++) {
    dp[i] = new int[W+1];
  }

  int ans = 0;
  for (int i = 1; i <= N; i++) {
    for (int j = 0; j <= W; j++) {
      dp[i][j] = dp[i-1][j];
      if (j >= w[i-1]) {
        dp[i][j] = max(dp[i][j], dp[i-1][j-w[i-1]] + v[i-1]);
      }
      if (j >= w[i]) {
        dp[i][j] = max(dp[i][j], dp[i][j-w[i]] + v[i]);
      }
      if (i == N) {
        ans = max(ans, dp[i][j]);
      }
    }
  }
  cout << ans << eol;
  return 0;
}
