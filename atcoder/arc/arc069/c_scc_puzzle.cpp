#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  long long n,m;
  cin >> n >> m;

  long ans = 0;
  if (n < m / 2) {
    ans = n + (m - 2*n) / 4;
  } else {
    ans = m / 2;
  }

  cout << ans << eol;
  
  return 0;
}
