#include <iostream>
#include <vector>
#include <algorithm>
#include <functional>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int n;
  cin >> n;
  vector<int> score(n, 0);
  for (int i = 0; i < n; i++) {
    cin >> score[i];
  }
  sort(score.begin(), score.end(), greater<int>());

  int q, k;
  cin >> q;
  for (int i = 0; i < q; i++) {
    cin >> k;

    if (score[0] == 0) {
      cout << 0 << eol;
      continue;
    }

    if (k == 0) {
      cout << score[0] + 1 << eol;  
      continue;
    }
    if (k == n) {
      cout << 0 << eol;
      continue;
    }

    int x = score[k-1];
    if (x == score[k]) {
      cout << x + 1 << eol;
    } else {
      cout << x << eol;
    }
  }
  
  return 0;
}
