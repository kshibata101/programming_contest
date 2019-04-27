#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int n;
  string s, t;
  cin >> n >> s;

  t = "";

  int b = 0, e = n-1;
  for (int i = 0; i < n; i++) {
    bool left = true;
    for (int i = 0; i <= e-b; i++) {
      if (s[b+i] < s[e-i]) {
        left = true;
        break;
      } else if (s[b+i] > s[e-i]) {
        left = false;
        break;
      }
    }
    if (left) {
      t += s[b++];
    } else {
      t += s[e--];
    }
  }

  cout << t << eol;
  
  return 0;
}
