#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;
     
int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);
     
  int n, f, cnt = 0;
  cin >> n;
  vector<int> v(n);
  for (int i = 0; i < n; i++) {
    cin >> v[i];
  }
  for (int i = 0; i < n; i++) {
    cin >> f;
    if (f > v[i] / 2) cnt++;
  }
     
  cout << cnt << eol;
  return 0;
}
