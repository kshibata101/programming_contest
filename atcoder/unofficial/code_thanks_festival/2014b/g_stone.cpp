#include <iostream>
#define eol '\n';
using namespace std;

int** dp;

bool is_sente(int n, int p) {
  if (p <= 0) return false;
  if (n <= p) return true;
  if (dp[n][p] == 1) return true;
  if (dp[n][p] == -1) return false;

  bool res = false;
  if (is_sente(n, p-1)) {
    res = true;
  } else {
    res = !is_sente(n-p, p+1);
  }

  if (res) 
    dp[n][p] = 1;
  else
    dp[n][p] = -1;

  return res;
}

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int N,P;
  cin >> N >> P;
  dp = new int*[N+1];
  for (int i = 0; i <= N; i++) {
    dp[i] = new int[P+1];
  }

  if (is_sente(N, P)) {
    cout << "first" << eol;
  } else {
    cout << "second" << eol;
  }
  
  return 0;
}
