#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;
     
int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);
     
  int a,b,c;
  cin >> a >> b >> c;
     
  int ans = 0;
  ans = max (a + b + c, ans);
  ans = max ((a+b) * c, ans);
  ans = max ((a*b) + c, ans);
  ans = max (a * b * c, ans);
  cout << ans << eol;
  return 0;
}
