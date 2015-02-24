#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;
     
int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);
     
  int N,T;
  cin >> N >> T;
  vector<int> a(N);
     
  for (int i = 0; i < N; i++) {
    cin >> a[i];
  }
     
  int ans = 0;
  for (int t = 1; t <= T; t++) {
    int cnt = 0;
    for (int i = 0; i < N; i++) {
      if (t % a[i] == 0) {
        cnt++;
      }
    }
    ans = max(ans, cnt);
  }
  cout << ans << eol;
  return 0;
}

