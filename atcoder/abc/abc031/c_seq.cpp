#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int solve(int N, vector<int> a) {
  int taka_max_max = -10000;
  
  for (int i = 0; i < N; i++) {
    int taka_max, aoki_max;
    taka_max = aoki_max = -10000;
    for (int j = 0; j < N; j++) {
      if (i == j) continue;
      int s = min(i,j);
      int e = max(i,j);

      int taka, aoki;
      taka = aoki = 0;
      for (int k = s; k <= e; k++) {
        if ((k-s) % 2 == 0) {
          taka += a[k];
        } else {
          aoki += a[k];
        }
      }

      if (aoki > aoki_max) {
        aoki_max = aoki;
        taka_max = taka;
      }
    }
    taka_max_max = max(taka_max_max, taka_max);
  }
  return taka_max_max;
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int N;
  cin >> N;
  vector<int> a(N);
  for (int i = 0; i < N; i++) {
    cin >> a[i];
  }

  cout << solve(N, a) << eol;

  return 0;
}
