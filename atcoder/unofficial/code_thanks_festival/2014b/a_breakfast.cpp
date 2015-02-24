#include <iostream>
#include <vector>
#define eol '\n'
using namespace std;
     
int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);
     
  int a,b;
  cin >> a >> b;
  cout << max(a, b) << eol;
  return 0;
}
