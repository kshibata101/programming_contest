#include <iostream>
#define eol '\n';
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int N,W;
  cin >> N >> W;
  int* v = new int[N];
  int* w = new int[N];

  for (int i = 0; i < N; i++) {
    cin >> v[i] >> w[i];
  }

  int** dp = new int*[N+1];
  for (int i = 0; i <= N; i++) {
    dp[i] = new int[W+1];
  }

  int ans = 0;
  for (int i = 0; i < N; i++) {
    for (int j = 0; j <= W; j++) {
      dp[i+1][j] = dp[i][j];
      if (j - w[i] >= 0) {
        dp[i+1][j] = max(dp[i+1][j], dp[i][j-w[i]] + v[i]);
      }
      if (i == N - 1) {
        ans = max(ans, dp[i+1][j]);
      }
    }
  }

  cout << ans << eol;

  return 0;
}
