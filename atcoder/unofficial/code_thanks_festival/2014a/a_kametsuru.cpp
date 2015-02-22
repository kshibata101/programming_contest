#include <iostream>
#define eol '\n';
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int a,b;
  cin >> a >> b;
  cout << a * 4 + b * 2 << eol;
  return 0;
}
