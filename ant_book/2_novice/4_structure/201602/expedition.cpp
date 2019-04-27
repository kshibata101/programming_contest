#include <iostream>
#include <vector>
#include <queue>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  priority_queue<int> pque;
  int n,l,p;

  cin >> n >> l >> p;
  
  vector<int> a(n+1),b(n+1);
  for (int i = 0; i < n; i++) {
    cin >> a[i] >> b[i];
  }
  a[n] = l - a[n-1];
  b[n] = 0;

  int cnt = 0;
  int pos = 0;

  for (int i = 0; i <= n; i++) {
    p -= (a[i] - pos);
    while (p < 0) {
      if (pque.empty()) {
        break;
      }
      p += pque.top();
      cnt++;
      pque.pop();
    }

    // Ç³ÎÁÀÚ¤ì
    if (p < 0) {
      cnt = -1;
      break;
    }

    pque.push(b[i]);
    pos = a[i];
  }

  cout << cnt << eol;
  
  return 0;
}
