#include <iostream>
#define eol '\n';
using namespace std;

int main() {
  ios_base::sync_with_stdio(false);
  cin.tie(0);

  int h1,h2;
  cin >> h1 >> h2;
  cout << h1 - h2 << eol;
  return 0;
}
