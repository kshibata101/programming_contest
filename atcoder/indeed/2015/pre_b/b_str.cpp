#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  string s,t;
  cin >> s >> t;

  int ans = -1;
  int size = s.size();
  for (int i = 0; i < size; i++) {
    if (s == t) {
      ans = i;
      break;
    }

    s = (s[size-1] + s).substr(0, size);
  }

  cout << ans << eol;
  
  return 0;
}
