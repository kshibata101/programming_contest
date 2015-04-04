#include <iostream>
#include <iomanip>
#define eol '\n';
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);
  int N, K;
  cin >> N >> K;
  int* p = new int[N+1];
  for (int i = 1; i <= N; i++) {
    cin >> p[i];
  }

  double*** dp = new double**[N+1];
  for (int n = 1; n <= N; n++) {
    dp[n] = new double*[K];
    for (int i = 0; i < K; i++) {
      dp[n][i] = new double[K+1];
    }
  }

  // init
  dp[1][0][0] = 1.0;

  for (int n = 1; n < N; n++) {
    for (int i = 0; i < K; i++) {
      for (int j = 0; j < K; j++) {
        if (j > 0) {
          dp[n+1][i][j-1] += dp[n][i][j] * p[n+1] / 100.0;
        } else if (i + 1 < K) {
          dp[n+1][i+1][j] += dp[n][i][j] * p[n+1] / 100.0;
        } else {
          dp[n+1][i][j] += dp[n][i][j] * p[n+1] / 100.0;
        }
        if (i + 2 < K && j + 1 < K) {
          dp[n+1][i+2][j+1] += dp[n][i][j] * (100 - p[n+1]) / 100.0;
        } else {
          dp[n+1][i][j] += dp[n][i][j] * (100 - p[n+1]) / 100.0;
        }
      }
    }
  }

  // expectation value
  double e = 0.0;
  for (int i = 0; i < K; i++) {
    int seat = K - 1 - i;
    for (int j = 0; j < K; j++) {
      e += (seat + j) * dp[N][i][j];
    }
  }
  cout << setprecision(8) << e << eol;
  return 0;
}
