#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int n, pos;
  cin >> n;
  string p,q,r,ans;
  ans = "";
  for (int i = 0; i < n; i++) {
    cin >> p >> q >> r;
    
    if (p.compare("BEGINNING") == 0) {
      pos = 0;
    } else if (p.compare("MIDDLE") == 0) {
      pos = r.size() / 2;
    } else if (p.compare("END") == 0) {
      pos = r.size() - 1;
    }

    ans += r[pos];
  }

  cout << ans << eol;
  
  return 0;
}
