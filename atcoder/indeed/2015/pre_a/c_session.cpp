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

    int x;

    if (k == 0) {
      x = score[0] + 1;
    } else {
      if (k == n) {
        x = 0;
      } else {
        x = score[k-1];
        if (x == score[k]) {
          x += 1;
        } else {
          x = score[k] + 1;
        }
      }
    }

    if (x == 1) {
      x = 0;
    }

    cout << x << eol;
  }
  return 0;
}
