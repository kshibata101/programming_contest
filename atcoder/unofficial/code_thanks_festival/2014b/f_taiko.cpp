#include <iostream>
#define eol '\n';
#define DEVIDE 1000000007
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  string x,s,t;
  cin >> x >> s >> t;
  int n = x.size();
  int* dp = new int[n+1];
  dp[0] = 1;

  int sl = s.size();
  int tl = t.size();
  for (int i = 1; i <= n; i++) {
    if (i >= sl && x.substr(i-sl, sl) == s) {
      dp[i] += dp[i-sl];
    }
    if (i >= tl && x.substr(i-tl, tl) == t) {
      dp[i] += dp[i-tl];
    }
    dp[i] = dp[i] % DEVIDE;
  }

  cout << dp[n] << eol;

  return 0;
}
